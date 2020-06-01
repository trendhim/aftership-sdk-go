package aftership_test

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
)

func GetCouriers() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get couriers
	result, err := cli.GetCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func GetAllCouriers() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all couriers
	result, err := cli.GetAllCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func DetectCouriers() {
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

	list, err := cli.DetectCouriers(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list)
}
