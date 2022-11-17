package aftership_test

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aftership/aftership-sdk-go/v2"
)

func ExampleClient_CreateTracking() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a tracking
	trackingNumber := strconv.FormatInt(time.Now().Unix(), 10)
	newTracking := aftership.CreateTrackingParams{
		TrackingNumber: trackingNumber,
		Slug:           "dhl",
		Title:          "Title Name",
		SMSes: []string{
			"+18555072509",
			"+18555072501",
		},
		Emails: []string{
			"email@yourdomain.com",
			"another_email@yourdomain.com",
		},
		OrderID: "ID 1234",
		CustomFields: map[string]string{
			"product_name":  "iPhone Case",
			"product_price": "USD19.99",
		},
		Language:                  "en",
		OrderPromisedDeliveryDate: "2019-05-20",
		DeliveryType:              "pickup_at_store",
		PickupLocation:            "Flagship Store",
		PickupNote:                "Reach out to our staffs when you arrive our stores for shipment pickup",
		ShipmentType:              "test_type",
		ShipmentTags:              []string{"test_tag1", "test_tag2"},
	}

	result, err := cli.CreateTracking(context.Background(), newTracking)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleClient_DeleteTracking() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Delete a tracking
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1234567890",
	}

	result, err := cli.DeleteTracking(context.Background(), param)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleClient_GetTrackings() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get tracking results of multiple trackings.
	multiParams := aftership.GetTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	multiResults, err := cli.GetTrackings(context.Background(), multiParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(multiResults)
}

func ExampleClient_GetTracking() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get tracking results of a single tracking.
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := cli.GetTracking(context.Background(), param, aftership.GetTrackingParams{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Get tracking results of a single tracking by id.
	paramID := aftership.TrackingID("rymq9l34ztbvvk9md2ync00r")

	result, err = cli.GetTracking(context.Background(), paramID, aftership.GetTrackingParams{
		Fields: "tracking_postal_code,title,order_id,shipment_tags",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func ExampleClient_UpdateTracking() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Update a tracking.
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	updateReq := aftership.UpdateTrackingParams{
		Title:        "New Title",
		ShipmentType: "test_type2",
	}

	result, err := cli.UpdateTracking(context.Background(), param, updateReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleClient_RetrackTracking() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Retrack an expired tracking.
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := cli.RetrackTracking(context.Background(), param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func ExampleClient_MarkTrackingAsCompleted() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	param := aftership.SlugTrackingNumber{
		Slug:           "USPS",
		TrackingNumber: "1587721393824",
	}

	result, err := cli.MarkTrackingAsCompleted(context.Background(), param, aftership.TrackingCompletedStatusDelivered)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
