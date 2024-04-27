package main

import (
	"booking-management/bootstrap"
	"context"
)

func main() {
	ctx := context.Background()

	bootstrap.Start(ctx)
}