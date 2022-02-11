package tests

import (
	"context"
	"errors"
	"fmt"
	postgres2 "github.com/JohnAzedo/eCommerce/product/infra/postgres"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type PostgresContainer struct {
	testcontainers.Container
	Database *gorm.DB
}

func SetupContainer(ctx context.Context) (*PostgresContainer, error) {
	const dbname = "products_test"
	const port = "5432/tcp"
	var env = map[string]string{
		"POSTGRES_PASSWORD": "lemonade",
		"POSTGRES_USER": "lemonade",
		"POSTGRES_DB": dbname,
	}

	req := testcontainers.ContainerRequest {
		Image: "postgres:latest",
		ExposedPorts: []string{port},
		Env: env,
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections"),
			wait.ForListeningPort(port),
		),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started: true,
	})

	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, port)
	if err != nil {
		return nil, err
	}

	db, err := SetupDatabase(dbname, ip, mappedPort.Port())
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{Container: container, Database: db}, nil
}


func SetupDatabase(dbname string, host string, port string) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=lemonade dbname=%s password=lemonade sslmode=disable", host, port, dbname)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&postgres2.ProductModel{})
	if err != nil {
		return nil, errors.New("Error on ProductModel: " + err.Error())
	}

	return db, nil
}