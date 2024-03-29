# Billing
(*Billing*)

### Available Operations

* [Get](#get) - List invoices
* [GetDetails](#getdetails) - Get invoice details

## Get

Get a list of past invoices

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"context"
	"log"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )


    var cursor *string = northflankgo.String("24")

    var page *int64 = northflankgo.Int64(1)

    var perPage *int64 = northflankgo.Int64(50)

    ctx := context.Background()
    res, err := s.Billing.Get(ctx, cursor, page, perPage)
    if err != nil {
        log.Fatal(err)
    }
    if res.PastInvoicesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | 24                                                    |
| `page`                                                | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 1                                                     |
| `perPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 50                                                    |


### Response

**[*operations.GetPastInvoicesResponse](../../pkg/models/operations/getpastinvoicesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDetails

Get details about an invoice. If `timestamp` is passed in as a query parameter, this endpoint returns details about the invoice containing that timestamp. Otherwise, returns a preview invoice displaying billing data from after the most recent invoice.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"context"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Billing.GetDetails(ctx, operations.GetInvoiceDetailsRequest{
        AddonID: northflankgo.String("default-addon"),
        JobID: northflankgo.String("example-JobId"),
        ProjectID: northflankgo.String("default-project"),
        ServiceID: northflankgo.String("example-service"),
        Timestamp: northflankgo.Int64(1657206215),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InvoiceDetailsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetInvoiceDetailsRequest](../../pkg/models/operations/getinvoicedetailsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetInvoiceDetailsResponse](../../pkg/models/operations/getinvoicedetailsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
