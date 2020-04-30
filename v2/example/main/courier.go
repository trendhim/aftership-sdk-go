package main

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/courier"
)

func main() {
	aftership := aftership.NewAfterShip(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})

	// Get couriers
	result, err := aftership.Courier.GetCouriers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Get all couriers
	result, err = aftership.Courier.GetAllCouriers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Detect courier
	req := courier.DetectCourierRequest{
		Tracking: courier.DetectParam{
			TrackingNumber: "906587618687",
		},
	}

	list, err := aftership.Courier.DetectCouriers(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}
}
