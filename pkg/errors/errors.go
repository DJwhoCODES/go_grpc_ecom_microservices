package errors

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrProductNotFound = errors.New("product not found")
	ErrCartNotFound    = errors.New("cart not found")
	ErrOrderNotFound   = errors.New("order not found")
	ErrPaymentFailed   = errors.New("payment failed")
)
