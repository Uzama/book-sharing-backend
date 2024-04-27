package handler

import (
	"booking-management/http/validators"
	"booking-management/utils/container"
)

type Handler struct {
	validator validators.Validator
}

func NewHandler(ctr container.Containers) Handler {
	handler := Handler{
		validator: validators.NewValidator(),
	}

	return handler
}