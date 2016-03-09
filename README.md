aftership-go 
============

### Introduction:

[Aftership](https://aftership.com) provides an API to Track & Notify of shipments from hundreds of couriers worldwide. aftership-go is a SDK to develop Apps using Aftership API in go-lang. All endpoints including couriers, tracking and notification are supported.

You will need to create an account at [Aftership](https://aftership.com) and obtain an API key first to access Aftership APIs using aftership-go SDK.

### Installation
If you already have go installed locally then do,
````
go get github.com/vimukthi-git/aftership-go
````
### Example:

- Getting couriers already added to your account,

```go
package main

import (
        "fmt"
        "github.com/vimukthi-git/aftership-go/apiV4"
        "github.com/vimukthi-git/aftership-go/impl"
)

func main() {
        var api apiV4.CourierHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
        }
        res, meta := api.GetCouriers()
        if (meta.Code == 200) {
            fmt.Print(res)
        }
        
}

```

- Posting a tracking to the API,

````go
package main

import (
        "fmt"
        "github.com/vimukthi-git/aftership-go/apiV4"
        "github.com/vimukthi-git/aftership-go/impl"
)

func main() {
        var api apiV4.TrackingsHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
        }
        res, meta := api.CreateTracking(apiV4.NewTracking{
                "1Z9999999999999998",
                nil,
                "",
                "",
                "",
                "",
                "",
                nil,
                nil,
                nil,
                nil,
                "",
                "",
                "",
                "",
                "",
                nil,
        })
        
        if (meta.Code == 200) {
            fmt.Print(res)
        }
}

````

- Getting Checkpoints and other tracking information for a tracking number,

````go
package main

import (
        "fmt"
        "github.com/vimukthi-git/aftership-go/apiV4"
        "github.com/vimukthi-git/aftership-go/impl"
)

func main() {
        var api apiV4.TrackingsHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
        }
        res, meta := api.GetTracking(
            apiV4.TrackingId{
                "",
                "xq-express",
                "LS404494276CN",
            }, 
            "",
            "",
        )
        
        if (meta.Code == 200) {
            fmt.Print(res)
        }
}

````

Check `./impl/impl_test.go` for examples on using all endpoints.