package aftership_test

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
)

func CheckpointsEndpoint_getLastCheckpoint() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get last checkpopint
	param := aftership.SlugTrackingNumber{
		Slug:           "ups",
		TrackingNumber: "1234567890",
	}

	result, err := cli.Checkpoint.GetLastCheckpoint(context.Background(), param, aftership.GetCheckpointParams{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
