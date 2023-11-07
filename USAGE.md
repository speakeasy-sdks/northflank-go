<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v2"
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
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
<!-- End SDK Example Usage -->