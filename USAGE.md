<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
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
	res, err := s.Miscellaneous.GetDNSID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.DNSIDResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->