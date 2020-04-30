package main

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/notification"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
)

func main() {
	aftership, err := aftership.NewAfterShip(&conf.AfterShipConf{
		APIKey: "d655b36a-d268-4eb3-a0d6-8939c79da93e",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the notification
	param := tracking.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := aftership.Notification.GetNotification(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Add notification receivers to a tracking number.
	data := notification.Data{
		Notification: notification.Notification{
			Emails: []string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
			Smses:  []string{"+85291239123", "+85261236123", "Invalid Mobile Phone Number"},
		},
	}
	result, err = aftership.Notification.AddNotification(param, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Remove notification receivers from a tracking number.
	data = notification.Data{
		Notification: notification.Notification{
			Emails: []string{"user2@gmail.com"},
			Smses:  []string{"+85261236123"},
		},
	}
	result, err = aftership.Notification.RemoveNotification(param, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
