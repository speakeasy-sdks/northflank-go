# Integrations
(*Integrations*)

### Available Operations

* [Add](#add) - Add registry
* [Create](#create) - Create log sink
* [Delete](#delete) - Delete log sink
* [DeleteRegistry](#deleteregistry) - Delete registry
* [GenerateVCSToken](#generatevcstoken) - Generate VCS token
* [Get](#get) - Get log sink details
* [GetBranches](#getbranches) - List branches
* [GetRegistry](#getregistry) - Get registry
* [GetRepos](#getrepos) - List repositories
* [ListLogSinks](#listlogsinks) - List log sinks
* [ListRegistries](#listregistries) - List registries
* [ListVcsProviders](#listvcsproviders) - List VCS providers
* [Pause](#pause) - Pause log sink
* [Resume](#resume) - Resume log sink
* [Update](#update) - Update log sink
* [UpdateRegistry](#updateregistry) - Update registry

## Add

Adds a new set of container registry credentials to this account.

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
    res, err := s.Integrations.Add(ctx, operations.CreateAddRegistryRequestBodyAddRegistryType2(
            shared.AddRegistryType2{
                Description: "This is a set of saved credentials.",
                Keyfile: shared.Keyfile{},
                Name: "Example Credentials",
                Provider: shared.AddRegistryType2ProviderDockerhub,
                RegistryURL: northflankgo.String("https://example.com"),
            },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.RegistriesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.AddRegistryRequestBody](../../pkg/models/operations/addregistryrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.AddRegistryResponse](../../pkg/models/operations/addregistryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Creates a new log sink.

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
    res, err := s.Integrations.Create(ctx, operations.CreateCreateLogSinkRequestBodyHTTPLogSink(
            shared.HTTPLogSink{
                Description: northflankgo.String("This is an example log sink."),
                ForwardAccessLogs: northflankgo.Bool(true),
                Name: "example-log-sink",
                Restricted: northflankgo.Bool(true),
                SinkData: shared.HTTPLogSinkSinkData{
                    Auth: shared.CreateHTTPLogSinkAuthHTTPLogSink2(
                            shared.HTTPLogSink2{
                                Password: "secret-password",
                                Strategy: shared.HTTPLogSinkStrategyBasic,
                                User: northflankgo.String("my-user"),
                            },
                    ),
                    URI: "my.log-collector.com",
                },
                SinkType: shared.HTTPLogSinkSinkTypeHTTP,
                UseCustomLabels: northflankgo.Bool(true),
            },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateLogSinkResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLogSinkRequestBody](../../pkg/models/operations/createlogsinkrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CreateLogSinkResponse](../../pkg/models/operations/createlogsinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

Deletes a log sink.

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


    var logSinkID string = "example-log-sink"

    ctx := context.Background()
    res, err := s.Integrations.Delete(ctx, logSinkID)
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
| `logSinkID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example-log-sink                                      |


### Response

**[*operations.DeleteLogSinkResponse](../../pkg/models/operations/deletelogsinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteRegistry

Deletes a set of registry credential data.

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


    var credentialID string = "example-credentials"

    ctx := context.Background()
    res, err := s.Integrations.DeleteRegistry(ctx, credentialID)
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
| `credentialID`                                        | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example-credentials                                   |


### Response

**[*operations.DeleteRegistryResponse](../../pkg/models/operations/deleteregistryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GenerateVCSToken

Generate a token for a specific VCS link.

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


    var customVCSID string = "cdb3d41f-0dd8-49ad-92d5-7544c98c490b"

    var vcsLinkID string = "63ebb6ce2ccc6c7affdbf253"

    ctx := context.Background()
    res, err := s.Integrations.GenerateVCSToken(ctx, customVCSID, vcsLinkID)
    if err != nil {
        log.Fatal(err)
    }
    if res.VCSTokenResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `customVCSID`                                         | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | cdb3d41f-0dd8-49ad-92d5-7544c98c490b                  |
| `vcsLinkID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | 63ebb6ce2ccc6c7affdbf253                              |


### Response

**[*operations.GenerateVCSTokenResponse](../../pkg/models/operations/generatevcstokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Gets details about a given log sink.

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


    var logSinkID string = "example-log-sink"

    ctx := context.Background()
    res, err := s.Integrations.Get(ctx, logSinkID)
    if err != nil {
        log.Fatal(err)
    }
    if res.LogSinkDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `logSinkID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example-log-sink                                      |


### Response

**[*operations.GetLogSinkResponse](../../pkg/models/operations/getlogsinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetBranches

Gets a list of branches for the repo

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"context"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/operations"
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
    res, err := s.Integrations.GetBranches(ctx, operations.GetBranchesRequest{
        PerPage: northflankgo.Int64(50),
        RepositoryName: "next-js-example",
        RepositoryOwner: "northflank-examples",
        VcsService: "github",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BranchesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetBranchesRequest](../../pkg/models/operations/getbranchesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetBranchesResponse](../../pkg/models/operations/getbranchesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRegistry

Views a set of registry credential data.

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


    var credentialID string = "example-credentials"

    ctx := context.Background()
    res, err := s.Integrations.GetRegistry(ctx, credentialID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegistryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `credentialID`                                        | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example-credentials                                   |


### Response

**[*operations.GetRegistryResponse](../../pkg/models/operations/getregistryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRepos

Gets a list of repositories accessible to this account

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"context"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/operations"
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
    res, err := s.Integrations.GetRepos(ctx, operations.GetRepositoriesRequest{
        AccountLogin: northflankgo.String("example-user"),
        PerPage: northflankgo.Int64(50),
        VcsService: operations.VcsServiceGithub.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RepositoriesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetRepositoriesRequest](../../pkg/models/operations/getrepositoriesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetRepositoriesResponse](../../pkg/models/operations/getrepositoriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListLogSinks

Gets a list of log sinks added to this account.

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


    var cursor *string = northflankgo.String("<value>")

    var perPage *int64 = northflankgo.Int64(50)

    ctx := context.Background()
    res, err := s.Integrations.ListLogSinks(ctx, cursor, perPage)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLogSinksResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `perPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 50                                                    |


### Response

**[*operations.GetLogSinksResponse](../../pkg/models/operations/getlogsinksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListRegistries

Lists the container registry credentials saved to this account. Does not display secrets.

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


    var cursor *string = northflankgo.String("<value>")

    var perPage *int64 = northflankgo.Int64(50)

    ctx := context.Background()
    res, err := s.Integrations.ListRegistries(ctx, cursor, perPage)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegistriesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `perPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 50                                                    |


### Response

**[*operations.GetRegistriesResponse](../../pkg/models/operations/getregistriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListVcsProviders

Lists linked version control providers

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
    res, err := s.Integrations.ListVcsProviders(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.VCSProvidersResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetVCSProvidersResponse](../../pkg/models/operations/getvcsprovidersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Pause

Pauses a given log sink.

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


    var logSinkID string = "example-log-sink"

    ctx := context.Background()
    res, err := s.Integrations.Pause(ctx, logSinkID)
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
| `logSinkID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example-log-sink                                      |


### Response

**[*operations.PauseLogSinkResponse](../../pkg/models/operations/pauselogsinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Resume

Resumes a paused log sink.

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


    var logSinkID string = "example-log-sink"

    ctx := context.Background()
    res, err := s.Integrations.Resume(ctx, logSinkID)
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
| `logSinkID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | example-log-sink                                      |


### Response

**[*operations.ResumeLogSinkResponse](../../pkg/models/operations/resumelogsinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Updates the settings for a log sink.

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


    logSinkRequest := shared.LogSinkRequest{
        Restricted: northflankgo.Bool(true),
        ResumeLogSink: northflankgo.Bool(false),
        UseCustomLabels: northflankgo.Bool(true),
    }

    var logSinkID string = "example-log-sink"

    ctx := context.Background()
    res, err := s.Integrations.Update(ctx, logSinkRequest, logSinkID)
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `logSinkRequest`                                                   | [shared.LogSinkRequest](../../pkg/models/shared/logsinkrequest.md) | :heavy_check_mark:                                                 | Request body                                                       |                                                                    |
| `logSinkID`                                                        | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                | example-log-sink                                                   |


### Response

**[*operations.UpdateLogSinkResponse](../../pkg/models/operations/updatelogsinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateRegistry

Updates a set of registry credential data.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	northflankgo "github.com/speakeasy-sdks/northflank-go/v3"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/operations"
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


    var requestBody operations.UpdateRegistryRequestBody = operations.CreateUpdateRegistryRequestBodyUpdateRegistryType2(
            shared.UpdateRegistryType2{
                Description: northflankgo.String("This is a set of saved credentials."),
                Password: "password1234",
                RegistryURL: northflankgo.String("https://example.com"),
                Username: "test-user",
            },
    )

    var credentialID string = "example-credentials"

    ctx := context.Background()
    res, err := s.Integrations.UpdateRegistry(ctx, requestBody, credentialID)
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `requestBody`                                                                                    | [operations.UpdateRegistryRequestBody](../../pkg/models/operations/updateregistryrequestbody.md) | :heavy_check_mark:                                                                               | Request body                                                                                     |                                                                                                  |
| `credentialID`                                                                                   | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              | example-credentials                                                                              |


### Response

**[*operations.UpdateRegistryResponse](../../pkg/models/operations/updateregistryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
