package aftership_test

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v3"
)

func ExampleClient_GetNotification() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the notification
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := cli.GetNotification(context.Background(), param)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleClient_AddNotification() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Add notification receivers to a tracking number.
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	data := aftership.Notification{
		Emails: []string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
		SMSes:  []string{"+85291239123", "+85261236123", "Invalid Mobile Phone Number"},
	}

	result, err := cli.AddNotification(context.Background(), param, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleClient_RemoveNotification() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Remove notification receivers from a tracking number.
	param := aftership.SlugTrackingNumber{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	data := aftership.Notification{
		Emails: []string{"user2@gmail.com"},
		SMSes:  []string{"+85261236123"},
	}

	result, err := cli.RemoveNotification(context.Background(), param, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
