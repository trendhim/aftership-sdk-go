package aftership

import (
	"context"
	"fmt"
)

func ExampleClient_BatchPredictEstimatedDeliveryDate() {
	cli, err := NewClient(Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	params := []EstimatedDeliveryDate{
		{
			Slug:            "fedex",
			ServiceTypeName: "FEDEX HOME DELIVERY",
			OriginAddress: &Address{
				Country:     "USA",
				State:       "WA",
				PostalCode:  "98108",
				RawLocation: "Seattle, Washington, 98108, USA, United States",
			},
			DestinationAddress: &Address{
				Country:     "USA",
				State:       "CA",
				PostalCode:  "92019",
				RawLocation: "El Cajon, California, 92019, USA, United States",
			},
			Weight: &Weight{
				Unit:  "kg",
				Value: 11,
			},
			PackageCount: 1,
			PickupTime:   "2021-07-01 15:00:00",
			EstimatedPickup: &EstimatedPickup{
				OrderTime:       "2021-07-01 15:04:05",
				OrderCutoffTime: "20:00:00",
				BusinessDays:    []int64{1, 2, 3, 4, 5, 6, 7},
				OrderProcessingTime: &OrderProcessingTime{
					Unit:  "day",
					Value: 0,
				},
			},
		},
	}

	list, err := cli.BatchPredictEstimatedDeliveryDate(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list)
}
