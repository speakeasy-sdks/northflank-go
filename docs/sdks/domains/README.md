# Domains
(*Domains*)

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


    addSubDomainRequest := shared.AddSubDomainRequest{
        Cdn: &shared.Cdn{
            Cloudfront: &shared.Cloudfront{
                Enabled: false,
            },
        },
        Subdomain: "site",
    }

    var domain string = "example.com"

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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `addSubDomainRequest`                                                        | [shared.AddSubDomainRequest](../../pkg/models/shared/addsubdomainrequest.md) | :heavy_check_mark:                                                           | Request body                                                                 |                                                                              |
| `domain`                                                                     | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          | example.com                                                                  |


### Response

**[*operations.AddSubDomainResponse](../../pkg/models/operations/addsubdomainresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 400,409                  | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## Assign

Assigns a service port to the given subdomain

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


    assignSubDomainRequest := shared.AssignSubDomainRequest{
        PortName: "port-1",
        ProjectID: "default-project",
        ServiceID: "example-service",
    }

    var domain string = "example.com"

    var subdomain string = "app"

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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `assignSubDomainRequest`                                                           | [shared.AssignSubDomainRequest](../../pkg/models/shared/assignsubdomainrequest.md) | :heavy_check_mark:                                                                 | Request body                                                                       |                                                                                    |
| `domain`                                                                           | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | example.com                                                                        |
| `subdomain`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | app                                                                                |


### Response

**[*operations.AssignSubDomainResponse](../../pkg/models/operations/assignsubdomainresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Registers a new domain

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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateDomainRequest](../../pkg/models/shared/createdomainrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateDomainResponse](../../pkg/models/operations/createdomainresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 400,409                  | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## Delete

Deletes a domain and each of its registered subdomains.

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


    var domain string = "example.com"

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

**[*operations.DeleteDomainResponse](../../pkg/models/operations/deletedomainresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCdn

Removes the CDN integration from the given subdomain

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


    cdnRequest := shared.CDNRequest{
        Provider: "cloudfront",
    }

    var domain string = "example.com"

    var subdomain string = "app"

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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `cdnRequest`                                               | [shared.CDNRequest](../../pkg/models/shared/cdnrequest.md) | :heavy_check_mark:                                         | Request body                                               |                                                            |
| `domain`                                                   | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | example.com                                                |
| `subdomain`                                                | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | app                                                        |


### Response

**[*operations.DeleteCDNResponse](../../pkg/models/operations/deletecdnresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteSubdomain

Removes a subdomain from a domain.

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


    var domain string = "example.com"

    var subdomain string = "app"

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

**[*operations.DeleteSubDomainResponse](../../pkg/models/operations/deletesubdomainresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 400,404                  | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## Enable

Enables a CDN integration on the given subdomain

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


    cdnRequest := shared.CDNRequest{
        Provider: "cloudfront",
    }

    var domain string = "example.com"

    var subdomain string = "app"

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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `cdnRequest`                                               | [shared.CDNRequest](../../pkg/models/shared/cdnrequest.md) | :heavy_check_mark:                                         | Request body                                               |                                                            |
| `domain`                                                   | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | example.com                                                |
| `subdomain`                                                | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | app                                                        |


### Response

**[*operations.EnableCDNResponse](../../pkg/models/operations/enablecdnresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Get the details about a domain

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


    var domain string = "example.com"

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

**[*operations.GetDomainResponse](../../pkg/models/operations/getdomainresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSubdomain

Gets details about the given subdomain

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


    var domain string = "example.com"

    var subdomain string = "app"

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

**[*operations.GetSubDomainResponse](../../pkg/models/operations/getsubdomainresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListDomains

Lists available domains

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


    var cursor *string = "string"

    var page *int64 = 1

    var perPage *int64 = 50

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

**[*operations.ListDomainsResponse](../../pkg/models/operations/listdomainsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Unassign

Removes a subdomain from its assigned service

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


    var domain string = "example.com"

    var subdomain string = "app"

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

**[*operations.UnassignSubDomainResponse](../../pkg/models/operations/unassignsubdomainresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Verify

Gets details about the given subdomain

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


    var domain string = "example.com"

    var subdomain string = "app"

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

**[*operations.VerifySubDomainResponse](../../pkg/models/operations/verifysubdomainresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 400                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## VerifyDomain

Attempts to verify the domain

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


    var domain string = "example.com"

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

**[*operations.VerifyDomainResponse](../../pkg/models/operations/verifydomainresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 400                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
