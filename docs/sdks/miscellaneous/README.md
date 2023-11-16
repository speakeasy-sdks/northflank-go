# Miscellaneous
(*Miscellaneous*)

### Available Operations

* [GetDNSID](#getdnsid) - Get DNS ID
* [HealthCheck](#healthcheck) - Health check

## GetDNSID

Returns the partially random string used when generating host names for the authenticated account.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v2"
	"context"
	"log"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Miscellaneous.GetDNSID(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DNSIDResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDNSIDResponse](../../pkg/models/operations/getdnsidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## HealthCheck

Returns api service status

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v2"
	"context"
	"log"
	"net/http"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Miscellaneous.HealthCheck(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetHealthCheckResponse](../../pkg/models/operations/gethealthcheckresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
