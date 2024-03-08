<div align="center">
    <img src="https://github.com/speakeasy-sdks/northflank-go/assets/6267663/f8f4b777-94b1-4603-93e5-64d48fd2aa54" width="300">
    <h1>Go SDK</h1>
   <p>The comprehensive developer platform to build and scale microservices, jobs and managed databases with a powerful UI, API & CLI.</p>
   <a href="https://northflank.com/docs"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=5444e4&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/northflank-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/northflank-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/northflank-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/northflank-go?sort=semver&style=for-the-badge" /></a>
</div>


<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/northflank-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Miscellaneous](docs/sdks/miscellaneous/README.md)

* [GetDNSID](docs/sdks/miscellaneous/README.md#getdnsid) - Get DNS ID
* [HealthCheck](docs/sdks/miscellaneous/README.md#healthcheck) - Health check

### [Addons](docs/sdks/addons/README.md)

* [ListAddonTypes](docs/sdks/addons/README.md#listaddontypes) - List addon types

### [Billing](docs/sdks/billing/README.md)

* [Get](docs/sdks/billing/README.md#get) - List invoices
* [GetDetails](docs/sdks/billing/README.md#getdetails) - Get invoice details

### [CloudProviders](docs/sdks/cloudproviders/README.md)

* [Create](docs/sdks/cloudproviders/README.md#create) - Create integration
* [CreateCluster](docs/sdks/cloudproviders/README.md#createcluster) - Create cluster
* [DeleteCluster](docs/sdks/cloudproviders/README.md#deletecluster) - Delete cluster
* [DeleteIntegration](docs/sdks/cloudproviders/README.md#deleteintegration) - Delete integration
* [Get](docs/sdks/cloudproviders/README.md#get) - List providers
* [GetCluster](docs/sdks/cloudproviders/README.md#getcluster) - Get cluster
* [GetIntegration](docs/sdks/cloudproviders/README.md#getintegration) - Get integration
* [ListClusters](docs/sdks/cloudproviders/README.md#listclusters) - List clusters
* [ListIntegrations](docs/sdks/cloudproviders/README.md#listintegrations) - List integrations
* [UpdateCluster](docs/sdks/cloudproviders/README.md#updatecluster) - Update cluster
* [UpdateIntegration](docs/sdks/cloudproviders/README.md#updateintegration) - Update integration

### [Domains](docs/sdks/domains/README.md)

* [Add](docs/sdks/domains/README.md#add) - Add subdomain
* [Assign](docs/sdks/domains/README.md#assign) - Assign service to subdomain
* [Create](docs/sdks/domains/README.md#create) - Create new domain
* [Delete](docs/sdks/domains/README.md#delete) - Delete domain
* [DeleteCdn](docs/sdks/domains/README.md#deletecdn) - Remove CDN from a subdomain
* [DeleteSubdomain](docs/sdks/domains/README.md#deletesubdomain) - Delete subdomain
* [Enable](docs/sdks/domains/README.md#enable) - Enable CDN on a subdomain
* [Get](docs/sdks/domains/README.md#get) - Get domain
* [GetSubdomain](docs/sdks/domains/README.md#getsubdomain) - Get subdomain
* [ListDomains](docs/sdks/domains/README.md#listdomains) - List domains
* [Unassign](docs/sdks/domains/README.md#unassign) - Unassign subdomain
* [Verify](docs/sdks/domains/README.md#verify) - Verify subdomain
* [VerifyDomain](docs/sdks/domains/README.md#verifydomain) - Verify domain

### [Integrations](docs/sdks/integrations/README.md)

* [Add](docs/sdks/integrations/README.md#add) - Add registry
* [Create](docs/sdks/integrations/README.md#create) - Create log sink
* [Delete](docs/sdks/integrations/README.md#delete) - Delete log sink
* [DeleteRegistry](docs/sdks/integrations/README.md#deleteregistry) - Delete registry
* [GenerateVCSToken](docs/sdks/integrations/README.md#generatevcstoken) - Generate VCS token
* [Get](docs/sdks/integrations/README.md#get) - Get log sink details
* [GetBranches](docs/sdks/integrations/README.md#getbranches) - List branches
* [GetRegistry](docs/sdks/integrations/README.md#getregistry) - Get registry
* [GetRepos](docs/sdks/integrations/README.md#getrepos) - List repositories
* [ListLogSinks](docs/sdks/integrations/README.md#listlogsinks) - List log sinks
* [ListRegistries](docs/sdks/integrations/README.md#listregistries) - List registries
* [ListVcsProviders](docs/sdks/integrations/README.md#listvcsproviders) - List VCS providers
* [Pause](docs/sdks/integrations/README.md#pause) - Pause log sink
* [Resume](docs/sdks/integrations/README.md#resume) - Resume log sink
* [Update](docs/sdks/integrations/README.md#update) - Update log sink
* [UpdateRegistry](docs/sdks/integrations/README.md#updateregistry) - Update registry
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 409                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

### Example

```go
package main

import (
	"context"
	"errors"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/sdkerrors"
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

	var clusterID string = "gcp-cluster-1"

	ctx := context.Background()
	res, err := s.CloudProviders.DeleteCluster(ctx, clusterID)
	if err != nil {

		var e *sdkerrors.APIErrorResult
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.northflank.com` | None |

#### Example

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
		northflankgo.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		northflankgo.WithServerURL("https://api.northflank.com"),
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BasicAuth`  | http         | HTTP Basic   |
| `BearerAuth` | http         | HTTP Bearer  |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
