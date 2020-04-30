package main

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/checkpoint"
	"github.com/aftership/aftership-sdk-go/v2/conf"
)

func main() {
	aftership, err := aftership.NewAfterShip(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get last checkpopint
	param := checkpoint.SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1234567890",
	}

	result, err := aftership.LastCheckpoint.GetLastCheckpoint(param, "", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
