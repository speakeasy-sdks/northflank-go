<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
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
    res, err := s.Addons.ListAddonTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddonTypesResult != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->