package aftership_test

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
)

func CouriersEndpoint_GetCouriers() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get couriers
	result, err := cli.Courier.GetCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func CouriersEndpoint_GetAllCouriers() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all couriers
	result, err := cli.Courier.GetAllCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func CouriersEndpoint_DetectCouriers() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Detect courier
	req := aftership.CourierDetectionParams{
		TrackingNumber: "906587618687",
	}

	list, err := cli.Courier.DetectCouriers(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list)
}
