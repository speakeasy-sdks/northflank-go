# CloudProviders

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
    res, err := s.CloudProviders.Create(ctx, shared.CreateIntegrationRequest{
        Credentials: shared.CreateIntegrationRequestCredentials{
            AccessKey: northflank.String("corrupti"),
            APIKey: northflank.String("provident"),
            KeyfileJSON: northflank.String("distinctio"),
            SecretKey: northflank.String("quibusdam"),
        },
        Description: northflank.String("This is a new cloud provider integration."),
        Gcp: &shared.CreateIntegrationRequestGcp{
            ProjectID: "unde",
        },
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateIntegrationRequest](../../models/shared/createintegrationrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateIntegrationResponse](../../models/operations/createintegrationresponse.md), error**


## CreateCluster

Creates a new cluster.

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
    res, err := s.CloudProviders.CreateCluster(ctx, shared.CreateClusterRequest{
        Description: northflank.String("This is a new cluster."),
        Gcp: &shared.CreateClusterRequestGcp{
            ProjectID: "example-project-id",
        },
        Integration: &shared.CreateClusterRequestIntegration{
            AccessKey: northflank.String("nulla"),
            APIKey: northflank.String("corrupti"),
            KeyfileJSON: northflank.String("illum"),
            SecretKey: northflank.String("vel"),
        },
        IntegrationID: northflank.String("gcp-integration"),
        KubernetesVersion: "1.23.8",
        Name: "GCP Cluster 1",
        NodePools: []shared.CreateClusterRequestNodePools{
            shared.CreateClusterRequestNodePools{
                Autoscaling: &shared.CreateClusterRequestNodePoolsAutoscaling{
                    Enabled: northflank.Bool(true),
                    Max: northflank.Int64(10),
                    Min: northflank.Int64(1),
                },
                AvailabilityZones: []string{
                    "error",
                },
                DiskSize: 100,
                DiskType: northflank.String("deserunt"),
                Labels: &shared.CreateClusterRequestNodePoolsLabels{},
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflank.Bool(false),
                SystemPool: northflank.Bool(false),
            },
        },
        Provider: shared.CreateClusterRequestProviderGcp,
        Region: "europe-west2",
        Settings: shared.CreateClusterRequestSettings{
            Builds: &shared.CreateClusterRequestSettingsBuilds{
                ClusterID: northflank.String("build-cluster"),
                Mode: shared.CreateClusterRequestSettingsBuildsModeInternal.ToPointer(),
                Plan: northflank.String("nf-compute-200"),
            },
            Logging: &shared.CreateClusterRequestSettingsLogging{
                Loki: &shared.CreateClusterRequestSettingsLoggingLoki{},
                Mode: shared.CreateClusterRequestSettingsLoggingModePaas.ToPointer(),
            },
            Registry: &shared.CreateClusterRequestSettingsRegistry{
                Mode: shared.CreateClusterRequestSettingsRegistryModePaas.ToPointer(),
                RegistryID: northflank.String("my-registry-credentials"),
            },
        },
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.CreateClusterRequest](../../models/shared/createclusterrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CreateClusterResponse](../../models/operations/createclusterresponse.md), error**


## DeleteCluster

Delete the given cluster. Fails if the cluster has associated projects.

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
    clusterID := "gcp-cluster-1"

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

**[*operations.DeleteClusterResponse](../../models/operations/deleteclusterresponse.md), error**


## DeleteIntegration

Delete the given integration. Fails if the integration is associated with existing clusters.

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
    integrationID := "gcp-integration-1"

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

**[*operations.DeleteIntegrationResponse](../../models/operations/deleteintegrationresponse.md), error**


## Get

Lists supported cloud providers

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

**[*operations.GetCloudProvidersResponse](../../models/operations/getcloudprovidersresponse.md), error**


## GetCluster

Get information about the given cluster

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
    clusterID := "gcp-cluster-1"

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

**[*operations.GetClusterResponse](../../models/operations/getclusterresponse.md), error**


## GetIntegration

Get information about the given integration

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
    integrationID := "gcp-integration-1"

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

**[*operations.GetIntegrationResponse](../../models/operations/getintegrationresponse.md), error**


## ListClusters

Lists clusters for the authenticated user or team.

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
    cursor := "24"
    page := 1
    perPage := 50

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

**[*operations.GetClustersResponse](../../models/operations/getclustersresponse.md), error**


## ListIntegrations

Lists integrations for the authenticated user or team.

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
    cursor := "debitis"
    page := 1
    perPage := 50

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

**[*operations.GetIntegrationsResponse](../../models/operations/getintegrationsresponse.md), error**


## UpdateCluster

Update an existing cluster.

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
    updateClusterRequest := shared.UpdateClusterRequest{
        Description: northflank.String("This is an updated description."),
        NodePools: []shared.UpdateClusterRequestNodePools{
            shared.UpdateClusterRequestNodePools{
                Autoscaling: &shared.UpdateClusterRequestNodePoolsAutoscaling{
                    Enabled: northflank.Bool(true),
                    Max: northflank.Int64(10),
                    Min: northflank.Int64(1),
                },
                AvailabilityZones: []string{
                    "ipsa",
                },
                DiskSize: 100,
                DiskType: northflank.String("delectus"),
                ID: northflank.String("6aa96121-0345-43ad-bade-af36d540c222"),
                Labels: &shared.UpdateClusterRequestNodePoolsLabels{},
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflank.Bool(false),
                SystemPool: northflank.Bool(false),
            },
        },
        Settings: &shared.UpdateClusterRequestSettings{
            Builds: &shared.UpdateClusterRequestSettingsBuilds{
                ClusterID: northflank.String("build-cluster"),
                Mode: shared.UpdateClusterRequestSettingsBuildsModePaas.ToPointer(),
                Plan: northflank.String("nf-compute-200"),
            },
            Logging: &shared.UpdateClusterRequestSettingsLogging{
                Loki: &shared.UpdateClusterRequestSettingsLoggingLoki{},
                Mode: shared.UpdateClusterRequestSettingsLoggingModePaas.ToPointer(),
            },
            Registry: &shared.UpdateClusterRequestSettingsRegistry{
                Mode: shared.UpdateClusterRequestSettingsRegistryModePaas.ToPointer(),
                RegistryID: northflank.String("my-registry-credentials"),
            },
        },
    }
    clusterID := "gcp-cluster-1"

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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `updateClusterRequest`                                                     | [shared.UpdateClusterRequest](../../models/shared/updateclusterrequest.md) | :heavy_check_mark:                                                         | Request body                                                               |                                                                            |
| `clusterID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        | gcp-cluster-1                                                              |


### Response

**[*operations.UpdateClusterResponse](../../models/operations/updateclusterresponse.md), error**


## UpdateIntegration

Update information about the given integration

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
    updateIntegrationRequest := shared.UpdateIntegrationRequest{
        Credentials: shared.UpdateIntegrationRequestCredentials{
            AccessKey: northflank.String("minus"),
            APIKey: northflank.String("placeat"),
            KeyfileJSON: northflank.String("voluptatum"),
            SecretKey: northflank.String("iusto"),
        },
        Description: northflank.String("This is a new description."),
    }
    integrationID := "gcp-integration-1"

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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `updateIntegrationRequest`                                                         | [shared.UpdateIntegrationRequest](../../models/shared/updateintegrationrequest.md) | :heavy_check_mark:                                                                 | Request body                                                                       |                                                                                    |
| `integrationID`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | gcp-integration-1                                                                  |


### Response

**[*operations.UpdateIntegrationResponse](../../models/operations/updateintegrationresponse.md), error**

