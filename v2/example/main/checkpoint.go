package main

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/checkpoint"
	"github.com/aftership/aftership-sdk-go/v2/conf"
)

func main() {
	aftership := aftership.NewAfterShip(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})

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
