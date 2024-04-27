package main

import (
	"booking-management/bootstrap"
	_ "booking-management/docs"
	"context"
)

func main() {
	ctx := context.Background()

	bootstrap.Start(ctx)
}