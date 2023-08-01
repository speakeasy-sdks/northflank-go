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

    ctx := context.Background()
    res, err := s.Billing.Get(ctx, operations.GetPastInvoicesRequest{
        Cursor: northflank.String("24"),
        Page: northflank.Int64(1),
        PerPage: northflank.Int64(50),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PastInvoicesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetPastInvoicesRequest](../../models/operations/getpastinvoicesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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

