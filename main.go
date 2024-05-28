package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"github.com/ankurparanjpe/orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
