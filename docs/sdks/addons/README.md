# Addons
(*Addons*)

### Available Operations

* [ListAddonTypes](#listaddontypes) - List addon types

## ListAddonTypes

Gets information about the available addon types

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

    ctx := context.Background()
    res, err := s.Addons.ListAddonTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddonTypesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListAddonTypesResponse](../../pkg/models/operations/listaddontypesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
