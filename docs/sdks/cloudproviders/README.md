# CloudProviders
(*CloudProviders*)

### Available Operations

* [Create](#create) - Create integration
* [CreateCluster](#createcluster) - Create cluster
* [DeleteCluster](#deletecluster) - Delete cluster
* [DeleteIntegration](#deleteintegration) - Delete integration
* [Get](#get) - List providers
* [GetCluster](#getcluster) - Get cluster
* [GetIntegration](#getintegration) - Get integration
* [ListClusters](#listclusters) - List clusters
* [ListIntegrations](#listintegrations) - List integrations
* [UpdateCluster](#updatecluster) - Update cluster
* [UpdateIntegration](#updateintegration) - Update integration

## Create

Creates a new integration.

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
    res, err := s.CloudProviders.Create(ctx, shared.CreateIntegrationRequest{
        Credentials: shared.Credentials{},
        Description: northflankgo.String("This is a new cloud provider integration."),
        Name: "New Integration",
        Provider: shared.CreateIntegrationRequestProviderGcp,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateIntegrationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.CreateIntegrationRequest](../../pkg/models/shared/createintegrationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateIntegrationResponse](../../pkg/models/operations/createintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCluster

Creates a new cluster.

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
    res, err := s.CloudProviders.CreateCluster(ctx, shared.CreateClusterRequest{
        Description: northflankgo.String("This is a new cluster."),
        IntegrationID: northflankgo.String("gcp-integration"),
        KubernetesVersion: "1.23.8",
        Name: "GCP Cluster 1",
        NodePools: []shared.NodePools{
            shared.NodePools{
                DiskSize: 100,
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflankgo.Bool(false),
            },
        },
        Provider: shared.CreateClusterRequestProviderGcp,
        Region: "europe-west2",
        Settings: shared.Settings{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateClusterResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateClusterRequest](../../pkg/models/shared/createclusterrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateClusterResponse](../../pkg/models/operations/createclusterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteCluster

Delete the given cluster. Fails if the cluster has associated projects.

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


    var clusterID string = "gcp-cluster-1"

    ctx := context.Background()
    res, err := s.CloudProviders.DeleteCluster(ctx, clusterID)
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
| `clusterID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | gcp-cluster-1                                         |


### Response

**[*operations.DeleteClusterResponse](../../pkg/models/operations/deleteclusterresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 409                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## DeleteIntegration

Delete the given integration. Fails if the integration is associated with existing clusters.

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


    var integrationID string = "gcp-integration-1"

    ctx := context.Background()
    res, err := s.CloudProviders.DeleteIntegration(ctx, integrationID)
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
| `integrationID`                                       | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | gcp-integration-1                                     |


### Response

**[*operations.DeleteIntegrationResponse](../../pkg/models/operations/deleteintegrationresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.APIErrorResult | 409                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## Get

Lists supported cloud providers

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
    res, err := s.CloudProviders.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.CloudProvidersResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCloudProvidersResponse](../../pkg/models/operations/getcloudprovidersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCluster

Get information about the given cluster

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


    var clusterID string = "gcp-cluster-1"

    ctx := context.Background()
    res, err := s.CloudProviders.GetCluster(ctx, clusterID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ClusterDetailsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `clusterID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | gcp-cluster-1                                         |


### Response

**[*operations.GetClusterResponse](../../pkg/models/operations/getclusterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetIntegration

Get information about the given integration

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


    var integrationID string = "gcp-integration-1"

    ctx := context.Background()
    res, err := s.CloudProviders.GetIntegration(ctx, integrationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetIntegrationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `integrationID`                                       | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | gcp-integration-1                                     |


### Response

**[*operations.GetIntegrationResponse](../../pkg/models/operations/getintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListClusters

Lists clusters for the authenticated user or team.

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


    var cursor *string = northflankgo.String("24")

    var page *int64 = northflankgo.Int64(1)

    var perPage *int64 = northflankgo.Int64(50)

    ctx := context.Background()
    res, err := s.CloudProviders.ListClusters(ctx, cursor, page, perPage)
    if err != nil {
        log.Fatal(err)
    }

    if res.ClustersResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | 24                                                    |
| `page`                                                | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 1                                                     |
| `perPage`                                             | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 50                                                    |


### Response

**[*operations.GetClustersResponse](../../pkg/models/operations/getclustersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListIntegrations

Lists integrations for the authenticated user or team.

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

    var page *int64 = northflankgo.Int64(1)

    var perPage *int64 = northflankgo.Int64(50)

    ctx := context.Background()
    res, err := s.CloudProviders.ListIntegrations(ctx, cursor, page, perPage)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListIntegrationsResult != nil {
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

**[*operations.GetIntegrationsResponse](../../pkg/models/operations/getintegrationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCluster

Update an existing cluster.

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


    updateClusterRequest := shared.UpdateClusterRequest{
        Description: northflankgo.String("This is an updated description."),
    }

    var clusterID string = "gcp-cluster-1"

    ctx := context.Background()
    res, err := s.CloudProviders.UpdateCluster(ctx, updateClusterRequest, clusterID)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateClusterResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `updateClusterRequest`                                                         | [shared.UpdateClusterRequest](../../pkg/models/shared/updateclusterrequest.md) | :heavy_check_mark:                                                             | Request body                                                                   |                                                                                |
| `clusterID`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            | gcp-cluster-1                                                                  |


### Response

**[*operations.UpdateClusterResponse](../../pkg/models/operations/updateclusterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateIntegration

Update information about the given integration

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


    updateIntegrationRequest := shared.UpdateIntegrationRequest{
        Credentials: shared.UpdateIntegrationRequestCredentials{},
        Description: northflankgo.String("This is a new description."),
    }

    var integrationID string = "gcp-integration-1"

    ctx := context.Background()
    res, err := s.CloudProviders.UpdateIntegration(ctx, updateIntegrationRequest, integrationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateIntegrationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `updateIntegrationRequest`                                                             | [shared.UpdateIntegrationRequest](../../pkg/models/shared/updateintegrationrequest.md) | :heavy_check_mark:                                                                     | Request body                                                                           |                                                                                        |
| `integrationID`                                                                        | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | gcp-integration-1                                                                      |


### Response

**[*operations.UpdateIntegrationResponse](../../pkg/models/operations/updateintegrationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
