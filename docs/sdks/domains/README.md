# Domains

### Available Operations

* [Add](#add) - Add subdomain
* [Assign](#assign) - Assign service to subdomain
* [Create](#create) - Create new domain
* [Delete](#delete) - Delete domain
* [DeleteCdn](#deletecdn) - Remove CDN from a subdomain
* [DeleteSubdomain](#deletesubdomain) - Delete subdomain
* [Enable](#enable) - Enable CDN on a subdomain
* [Get](#get) - Get domain
* [GetSubdomain](#getsubdomain) - Get subdomain
* [ListDomains](#listdomains) - List domains
* [Unassign](#unassign) - Unassign subdomain
* [Verify](#verify) - Verify subdomain
* [VerifyDomain](#verifydomain) - Verify domain

## Add

Adds a new subdomain to the domain.

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
    addSubDomainRequest := shared.AddSubDomainRequest{
        Cdn: &shared.AddSubDomainRequestCdn{
            Cloudfront: &shared.AddSubDomainRequestCdnCloudfront{
                Enabled: false,
            },
        },
        Subdomain: "site",
    }
    domain := "example.com"

    ctx := context.Background()
    res, err := s.Domains.Add(ctx, addSubDomainRequest, domain)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddSubDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `addSubDomainRequest`                                                    | [shared.AddSubDomainRequest](../../models/shared/addsubdomainrequest.md) | :heavy_check_mark:                                                       | Request body                                                             |                                                                          |
| `domain`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      | example.com                                                              |


### Response

**[*operations.AddSubDomainResponse](../../models/operations/addsubdomainresponse.md), error**


## Assign

Assigns a service port to the given subdomain

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
    assignSubDomainRequest := shared.AssignSubDomainRequest{
        PortName: "port-1",
        ProjectID: "default-project",
        ServiceID: "example-service",
    }
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.Assign(ctx, assignSubDomainRequest, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `assignSubDomainRequest`                                                       | [shared.AssignSubDomainRequest](../../models/shared/assignsubdomainrequest.md) | :heavy_check_mark:                                                             | Request body                                                                   |                                                                                |
| `domain`                                                                       | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | example.com                                                                    |
| `subdomain`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | app                                                                            |


### Response

**[*operations.AssignSubDomainResponse](../../models/operations/assignsubdomainresponse.md), error**


## Create

Registers a new domain

### Example Usage

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
    res, err := s.Domains.Create(ctx, shared.CreateDomainRequest{
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateDomainRequest](../../models/shared/createdomainrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateDomainResponse](../../models/operations/createdomainresponse.md), error**


## Delete

Deletes a domain and each of its registered subdomains.

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
    domain := "example.com"

    ctx := context.Background()
    res, err := s.Domains.Delete(ctx, domain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |


### Response

**[*operations.DeleteDomainResponse](../../models/operations/deletedomainresponse.md), error**


## DeleteCdn

Removes the CDN integration from the given subdomain

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
    cdnRequest := shared.CDNRequest{
        Provider: "cloudfront",
    }
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.DeleteCdn(ctx, cdnRequest, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |                                                        |
| `cdnRequest`                                           | [shared.CDNRequest](../../models/shared/cdnrequest.md) | :heavy_check_mark:                                     | Request body                                           |                                                        |
| `domain`                                               | *string*                                               | :heavy_check_mark:                                     | N/A                                                    | example.com                                            |
| `subdomain`                                            | *string*                                               | :heavy_check_mark:                                     | N/A                                                    | app                                                    |


### Response

**[*operations.DeleteCDNResponse](../../models/operations/deletecdnresponse.md), error**


## DeleteSubdomain

Removes a subdomain from a domain.

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
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.DeleteSubdomain(ctx, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |
| `subdomain`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | app                                                   |


### Response

**[*operations.DeleteSubDomainResponse](../../models/operations/deletesubdomainresponse.md), error**


## Enable

Enables a CDN integration on the given subdomain

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
    cdnRequest := shared.CDNRequest{
        Provider: "cloudfront",
    }
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.Enable(ctx, cdnRequest, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |                                                        |
| `cdnRequest`                                           | [shared.CDNRequest](../../models/shared/cdnrequest.md) | :heavy_check_mark:                                     | Request body                                           |                                                        |
| `domain`                                               | *string*                                               | :heavy_check_mark:                                     | N/A                                                    | example.com                                            |
| `subdomain`                                            | *string*                                               | :heavy_check_mark:                                     | N/A                                                    | app                                                    |


### Response

**[*operations.EnableCDNResponse](../../models/operations/enablecdnresponse.md), error**


## Get

Get the details about a domain

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
    domain := "example.com"

    ctx := context.Background()
    res, err := s.Domains.Get(ctx, domain)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |


### Response

**[*operations.GetDomainResponse](../../models/operations/getdomainresponse.md), error**


## GetSubdomain

Gets details about the given subdomain

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
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.GetSubdomain(ctx, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSubDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |
| `subdomain`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | app                                                   |


### Response

**[*operations.GetSubDomainResponse](../../models/operations/getsubdomainresponse.md), error**


## ListDomains

Lists available domains

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
    cursor := "perferendis"
    page := 1
    perPage := 50

    ctx := context.Background()
    res, err := s.Domains.ListDomains(ctx, cursor, page, perPage)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDomainsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `page`                                                | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 1                                                     |
| `perPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 50                                                    |


### Response

**[*operations.ListDomainsResponse](../../models/operations/listdomainsresponse.md), error**


## Unassign

Removes a subdomain from its assigned service

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
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.Unassign(ctx, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |
| `subdomain`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | app                                                   |


### Response

**[*operations.UnassignSubDomainResponse](../../models/operations/unassignsubdomainresponse.md), error**


## Verify

Gets details about the given subdomain

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
    domain := "example.com"
    subdomain := "app"

    ctx := context.Background()
    res, err := s.Domains.Verify(ctx, domain, subdomain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |
| `subdomain`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | app                                                   |


### Response

**[*operations.VerifySubDomainResponse](../../models/operations/verifysubdomainresponse.md), error**


## VerifyDomain

Attempts to verify the domain

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
    domain := "example.com"

    ctx := context.Background()
    res, err := s.Domains.VerifyDomain(ctx, domain)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `domain`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example.com                                           |


### Response

**[*operations.VerifyDomainResponse](../../models/operations/verifydomainresponse.md), error**

