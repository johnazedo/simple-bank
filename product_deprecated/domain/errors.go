package domain

import "errors"
var InvalidPriceError = errors.New("Product price must be greater than 0")
