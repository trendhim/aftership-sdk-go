package aftership_test

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
)

func ExampleNewClient() {
	cli, err := aftership.NewClient(aftership.Config{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cli)
}
