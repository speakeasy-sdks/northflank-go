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

    ctx := context.Background()
    res, err := s.Domains.Add(ctx, operations.AddSubDomainRequest{
        AddSubDomainRequest: shared.AddSubDomainRequest{
            Cdn: &shared.AddSubDomainRequestCdn{
                Cloudfront: &shared.AddSubDomainRequestCdnCloudfront{
                    Enabled: false,
                },
            },
            Subdomain: "site",
        },
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddSubDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.AddSubDomainRequest](../../models/operations/addsubdomainrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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

    ctx := context.Background()
    res, err := s.Domains.Assign(ctx, operations.AssignSubDomainRequest{
        AssignSubDomainRequest: shared.AssignSubDomainRequest{
            PortName: "port-1",
            ProjectID: "default-project",
            ServiceID: "example-service",
        },
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AssignSubDomainRequest](../../models/operations/assignsubdomainrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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

    ctx := context.Background()
    res, err := s.Domains.Delete(ctx, operations.DeleteDomainRequest{
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteDomainRequest](../../models/operations/deletedomainrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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

    ctx := context.Background()
    res, err := s.Domains.DeleteCdn(ctx, operations.DeleteCDNRequest{
        CDNRequest: shared.CDNRequest{
            Provider: "cloudfront",
        },
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeleteCDNRequest](../../models/operations/deletecdnrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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

    ctx := context.Background()
    res, err := s.Domains.DeleteSubdomain(ctx, operations.DeleteSubDomainRequest{
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteSubDomainRequest](../../models/operations/deletesubdomainrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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

    ctx := context.Background()
    res, err := s.Domains.Enable(ctx, operations.EnableCDNRequest{
        CDNRequest: shared.CDNRequest{
            Provider: "cloudfront",
        },
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.EnableCDNRequest](../../models/operations/enablecdnrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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

    ctx := context.Background()
    res, err := s.Domains.Get(ctx, operations.GetDomainRequest{
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetDomainRequest](../../models/operations/getdomainrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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

    ctx := context.Background()
    res, err := s.Domains.GetSubdomain(ctx, operations.GetSubDomainRequest{
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSubDomainResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetSubDomainRequest](../../models/operations/getsubdomainrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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

    ctx := context.Background()
    res, err := s.Domains.ListDomains(ctx, operations.ListDomainsRequest{
        Cursor: northflank.String("deleniti"),
        Page: northflank.Int64(1),
        PerPage: northflank.Int64(50),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDomainsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListDomainsRequest](../../models/operations/listdomainsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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

    ctx := context.Background()
    res, err := s.Domains.Unassign(ctx, operations.UnassignSubDomainRequest{
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UnassignSubDomainRequest](../../models/operations/unassignsubdomainrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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

    ctx := context.Background()
    res, err := s.Domains.Verify(ctx, operations.VerifySubDomainRequest{
        Domain: "example.com",
        Subdomain: "app",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.VerifySubDomainRequest](../../models/operations/verifysubdomainrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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

    ctx := context.Background()
    res, err := s.Domains.VerifyDomain(ctx, operations.VerifyDomainRequest{
        Domain: "example.com",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.VerifyDomainRequest](../../models/operations/verifydomainrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.VerifyDomainResponse](../../models/operations/verifydomainresponse.md), error**

