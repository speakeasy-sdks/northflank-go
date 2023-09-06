# Integrations

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
    res, err := s.Integrations.Add(ctx, operations.AddRegistryRequestBody{})
    if err != nil {
        log.Fatal(err)
    }

    if res.RegistriesResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AddRegistryRequestBody](../../models/operations/addregistryrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.AddRegistryResponse](../../models/operations/addregistryresponse.md), error**


## Create

Creates a new log sink.

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
    res, err := s.Integrations.Create(ctx, operations.CreateLogSinkRequestBody{})
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateLogSinkResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateLogSinkRequestBody](../../models/operations/createlogsinkrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateLogSinkResponse](../../models/operations/createlogsinkresponse.md), error**


## Delete

Deletes a log sink.

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
    logSinkID := "example-log-sink"

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

**[*operations.DeleteLogSinkResponse](../../models/operations/deletelogsinkresponse.md), error**


## DeleteRegistry

Deletes a set of registry credential data.

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
    credentialID := "example-credentials"

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

**[*operations.DeleteRegistryResponse](../../models/operations/deleteregistryresponse.md), error**


## GenerateVCSToken

Generate a token for a specific VCS link.

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
    customVCSID := "cdb3d41f-0dd8-49ad-92d5-7544c98c490b"
    vcsLinkID := "63ebb6ce2ccc6c7affdbf253"

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

**[*operations.GenerateVCSTokenResponse](../../models/operations/generatevcstokenresponse.md), error**


## Get

Gets details about a given log sink.

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
    logSinkID := "example-log-sink"

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

**[*operations.GetLogSinkResponse](../../models/operations/getlogsinkresponse.md), error**


## GetBranches

Gets a list of branches for the repo

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
    res, err := s.Integrations.GetBranches(ctx, operations.GetBranchesRequest{
        Cursor: northflank.String("nisi"),
        PerPage: northflank.Int64(50),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetBranchesRequest](../../models/operations/getbranchesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetBranchesResponse](../../models/operations/getbranchesresponse.md), error**


## GetRegistry

Views a set of registry credential data.

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
    credentialID := "example-credentials"

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

**[*operations.GetRegistryResponse](../../models/operations/getregistryresponse.md), error**


## GetRepos

Gets a list of repositories accessible to this account

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
    res, err := s.Integrations.GetRepos(ctx, operations.GetRepositoriesRequest{
        AccountLogin: northflank.String("example-user"),
        Cursor: northflank.String("recusandae"),
        PerPage: northflank.Int64(50),
        SelfHostedVcsID: northflank.String("temporibus"),
        VcsLinkID: northflank.String("ab"),
        VcsService: operations.GetRepositoriesVcsServiceGithub.ToPointer(),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetRepositoriesRequest](../../models/operations/getrepositoriesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetRepositoriesResponse](../../models/operations/getrepositoriesresponse.md), error**


## ListLogSinks

Gets a list of log sinks added to this account.

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
    cursor := "quis"
    perPage := 50

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

**[*operations.GetLogSinksResponse](../../models/operations/getlogsinksresponse.md), error**


## ListRegistries

Lists the container registry credentials saved to this account. Does not display secrets.

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
    cursor := "veritatis"
    perPage := 50

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

**[*operations.GetRegistriesResponse](../../models/operations/getregistriesresponse.md), error**


## ListVcsProviders

Lists linked version control providers

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

**[*operations.GetVCSProvidersResponse](../../models/operations/getvcsprovidersresponse.md), error**


## Pause

Pauses a given log sink.

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
    logSinkID := "example-log-sink"

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

**[*operations.PauseLogSinkResponse](../../models/operations/pauselogsinkresponse.md), error**


## Resume

Resumes a paused log sink.

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
    logSinkID := "example-log-sink"

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

**[*operations.ResumeLogSinkResponse](../../models/operations/resumelogsinkresponse.md), error**


## Update

Updates the settings for a log sink.

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
    logSinkRequest := shared.LogSinkRequest{
        Projects: []string{
            "default-project",
        },
        Restricted: northflank.Bool(true),
        ResumeLogSink: northflank.Bool(false),
        SinkData: &shared.LogSinkRequestSinkData{},
        UseCustomLabels: northflank.Bool(true),
    }
    logSinkID := "example-log-sink"

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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `logSinkRequest`                                               | [shared.LogSinkRequest](../../models/shared/logsinkrequest.md) | :heavy_check_mark:                                             | Request body                                                   |                                                                |
| `logSinkID`                                                    | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            | example-log-sink                                               |


### Response

**[*operations.UpdateLogSinkResponse](../../models/operations/updatelogsinkresponse.md), error**


## UpdateRegistry

Updates a set of registry credential data.

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
    requestBody := operations.UpdateRegistryRequestBody{}
    credentialID := "example-credentials"

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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `requestBody`                                                                                | [operations.UpdateRegistryRequestBody](../../models/operations/updateregistryrequestbody.md) | :heavy_check_mark:                                                                           | Request body                                                                                 |                                                                                              |
| `credentialID`                                                                               | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          | example-credentials                                                                          |


### Response

**[*operations.UpdateRegistryResponse](../../models/operations/updateregistryresponse.md), error**

