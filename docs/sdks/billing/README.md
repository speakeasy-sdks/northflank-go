# Billing

### Available Operations

* [Get](#get) - List invoices
* [GetDetails](#getdetails) - Get invoice details

## Get

Get a list of past invoices

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/operations"
)

func main() {
    s := northflank.New(
        northflank.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )
    cursor := "24"
    page := 1
    perPage := 50

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

**[*operations.GetPastInvoicesResponse](../../models/operations/getpastinvoicesresponse.md), error**


## GetDetails

Get details about an invoice. If `timestamp` is passed in as a query parameter, this endpoint returns details about the invoice containing that timestamp. Otherwise, returns a preview invoice displaying billing data from after the most recent invoice.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/operations"
)

func main() {
    s := northflank.New(
        northflank.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Billing.GetDetails(ctx, operations.GetInvoiceDetailsRequest{
        AddonID: northflank.String("default-addon"),
        JobID: northflank.String("example-JobId"),
        ProjectID: northflank.String("default-project"),
        ServiceID: northflank.String("example-service"),
        Timestamp: northflank.Int64(1657206215),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetInvoiceDetailsRequest](../../models/operations/getinvoicedetailsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetInvoiceDetailsResponse](../../models/operations/getinvoicedetailsresponse.md), error**

