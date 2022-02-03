package domain

type InvalidPriceError struct {
	Err error
}

func (e InvalidPriceError) Error() string {
	return "Product price must be greater than 0"
}

