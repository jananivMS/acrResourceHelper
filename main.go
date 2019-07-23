package main

import (
	"context"
	"fmt"

	"github.com/jananiv/resourcehelper/acr"
	"github.com/jananiv/resourcehelper/config"
)

func main() {
	config.LoadSettings()

	ctx := context.Background()

	//_, err := acr.CreateRegistry(ctx, "Janani-testRG", "JananivAcr123", "southcentralus", "Standard", false)
	_,err := acr.DeleteRegistry(ctx, "Janani-testRG", "JananivAcr123")
	if err != nil {
		fmt.Println("error is %s", err)
	}
}
