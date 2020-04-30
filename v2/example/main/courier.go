package main

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/conf"
)

func main() {
	aftership := aftership.NewAfterShip(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})

	result, err := aftership.Courier.GetCouriers()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
