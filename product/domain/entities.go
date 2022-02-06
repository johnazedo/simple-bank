package domain

type Measurement string

const (
	PerKilo Measurement = "/kg"
	PerUnit             = "/unidade"
)

type Product struct {
	UUID        string      `json:"uuid"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Price       float32     `json:"price"`
	Description string      `json:"description"`
	Measurement Measurement `json:"measurement"`
	Amount      int         `json:"amount"`
}

func NewProduct(name string, image string, price float32,
	description string, measurement Measurement, amount int) *Product {

	product := &Product{
		Name:        name,
		Image:       image,
		Price:       price,
		Description: description,
		Measurement: measurement,
		Amount:      amount,
	}
	return product
}

func (p *Product) CheckValidAmount() {
	if p.Amount < 0 {
		p.Amount = 0
	}
}

func (p *Product) CheckValidPrice() error {
	if p.Price <= 0 {
		return InvalidPriceError{}
	}
	return nil
}
