<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
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