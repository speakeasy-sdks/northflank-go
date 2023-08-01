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
                    "suscipit",
                    "iure",
                    "magnam",
                },
                DiskSize: 100,
                DiskType: northflank.String("debitis"),
                Labels: &shared.CreateClusterRequestNodePoolsLabels{},
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflank.Bool(false),
                SystemPool: northflank.Bool(false),
            },
            shared.CreateClusterRequestNodePools{
                Autoscaling: &shared.CreateClusterRequestNodePoolsAutoscaling{
                    Enabled: northflank.Bool(true),
                    Max: northflank.Int64(10),
                    Min: northflank.Int64(1),
                },
                AvailabilityZones: []string{
                    "delectus",
                },
                DiskSize: 100,
                DiskType: northflank.String("tempora"),
                Labels: &shared.CreateClusterRequestNodePoolsLabels{},
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflank.Bool(false),
                SystemPool: northflank.Bool(false),
            },
            shared.CreateClusterRequestNodePools{
                Autoscaling: &shared.CreateClusterRequestNodePoolsAutoscaling{
                    Enabled: northflank.Bool(true),
                    Max: northflank.Int64(10),
                    Min: northflank.Int64(1),
                },
                AvailabilityZones: []string{
                    "molestiae",
                    "minus",
                },
                DiskSize: 100,
                DiskType: northflank.String("placeat"),
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
                Mode: shared.CreateClusterRequestSettingsRegistryModeSelfHosted.ToPointer(),
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

    ctx := context.Background()
    res, err := s.CloudProviders.DeleteCluster(ctx, operations.DeleteClusterRequest{
        ClusterID: "gcp-cluster-1",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteClusterRequest](../../models/operations/deleteclusterrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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

    ctx := context.Background()
    res, err := s.CloudProviders.DeleteIntegration(ctx, operations.DeleteIntegrationRequest{
        IntegrationID: "gcp-integration-1",
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
| `request`                                                                                  | [operations.DeleteIntegrationRequest](../../models/operations/deleteintegrationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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

    ctx := context.Background()
    res, err := s.CloudProviders.GetCluster(ctx, operations.GetClusterRequest{
        ClusterID: "gcp-cluster-1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ClusterDetailsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetClusterRequest](../../models/operations/getclusterrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


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

    ctx := context.Background()
    res, err := s.CloudProviders.GetIntegration(ctx, operations.GetIntegrationRequest{
        IntegrationID: "gcp-integration-1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetIntegrationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetIntegrationRequest](../../models/operations/getintegrationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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

    ctx := context.Background()
    res, err := s.CloudProviders.ListClusters(ctx, operations.GetClustersRequest{
        Cursor: northflank.String("24"),
        Page: northflank.Int64(1),
        PerPage: northflank.Int64(50),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ClustersResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetClustersRequest](../../models/operations/getclustersrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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

    ctx := context.Background()
    res, err := s.CloudProviders.ListIntegrations(ctx, operations.GetIntegrationsRequest{
        Cursor: northflank.String("nisi"),
        Page: northflank.Int64(1),
        PerPage: northflank.Int64(50),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListIntegrationsResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetIntegrationsRequest](../../models/operations/getintegrationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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

    ctx := context.Background()
    res, err := s.CloudProviders.UpdateCluster(ctx, operations.UpdateClusterRequest{
        UpdateClusterRequest: shared.UpdateClusterRequest{
            Description: northflank.String("This is an updated description."),
            NodePools: []shared.UpdateClusterRequestNodePools{
                shared.UpdateClusterRequestNodePools{
                    Autoscaling: &shared.UpdateClusterRequestNodePoolsAutoscaling{
                        Enabled: northflank.Bool(true),
                        Max: northflank.Int64(10),
                        Min: northflank.Int64(1),
                    },
                    AvailabilityZones: []string{
                        "ab",
                        "quis",
                        "veritatis",
                        "deserunt",
                    },
                    DiskSize: 100,
                    DiskType: northflank.String("perferendis"),
                    ID: northflank.String("6aa96121-0345-43ad-bade-af36d540c222"),
                    Labels: &shared.UpdateClusterRequestNodePoolsLabels{},
                    NodeCount: 3,
                    NodeType: "n2-standard-8",
                    Preemptible: northflank.Bool(false),
                    SystemPool: northflank.Bool(false),
                },
                shared.UpdateClusterRequestNodePools{
                    Autoscaling: &shared.UpdateClusterRequestNodePoolsAutoscaling{
                        Enabled: northflank.Bool(true),
                        Max: northflank.Int64(10),
                        Min: northflank.Int64(1),
                    },
                    AvailabilityZones: []string{
                        "repellendus",
                        "sapiente",
                    },
                    DiskSize: 100,
                    DiskType: northflank.String("quo"),
                    ID: northflank.String("6aa96121-0345-43ad-bade-af36d540c222"),
                    Labels: &shared.UpdateClusterRequestNodePoolsLabels{},
                    NodeCount: 3,
                    NodeType: "n2-standard-8",
                    Preemptible: northflank.Bool(false),
                    SystemPool: northflank.Bool(false),
                },
                shared.UpdateClusterRequestNodePools{
                    Autoscaling: &shared.UpdateClusterRequestNodePoolsAutoscaling{
                        Enabled: northflank.Bool(true),
                        Max: northflank.Int64(10),
                        Min: northflank.Int64(1),
                    },
                    AvailabilityZones: []string{
                        "at",
                    },
                    DiskSize: 100,
                    DiskType: northflank.String("at"),
                    ID: northflank.String("6aa96121-0345-43ad-bade-af36d540c222"),
                    Labels: &shared.UpdateClusterRequestNodePoolsLabels{},
                    NodeCount: 3,
                    NodeType: "n2-standard-8",
                    Preemptible: northflank.Bool(false),
                    SystemPool: northflank.Bool(false),
                },
                shared.UpdateClusterRequestNodePools{
                    Autoscaling: &shared.UpdateClusterRequestNodePoolsAutoscaling{
                        Enabled: northflank.Bool(true),
                        Max: northflank.Int64(10),
                        Min: northflank.Int64(1),
                    },
                    AvailabilityZones: []string{
                        "molestiae",
                        "quod",
                        "quod",
                        "esse",
                    },
                    DiskSize: 100,
                    DiskType: northflank.String("totam"),
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
                    Mode: shared.UpdateClusterRequestSettingsBuildsModeBuildCluster.ToPointer(),
                    Plan: northflank.String("nf-compute-200"),
                },
                Logging: &shared.UpdateClusterRequestSettingsLogging{
                    Loki: &shared.UpdateClusterRequestSettingsLoggingLoki{},
                    Mode: shared.UpdateClusterRequestSettingsLoggingModeLoki.ToPointer(),
                },
                Registry: &shared.UpdateClusterRequestSettingsRegistry{
                    Mode: shared.UpdateClusterRequestSettingsRegistryModePaas.ToPointer(),
                    RegistryID: northflank.String("my-registry-credentials"),
                },
            },
        },
        ClusterID: "gcp-cluster-1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateClusterResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateClusterRequest](../../models/operations/updateclusterrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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

    ctx := context.Background()
    res, err := s.CloudProviders.UpdateIntegration(ctx, operations.UpdateIntegrationRequest{
        UpdateIntegrationRequest: shared.UpdateIntegrationRequest{
            Credentials: shared.UpdateIntegrationRequestCredentials{
                AccessKey: northflank.String("nam"),
                APIKey: northflank.String("officia"),
                KeyfileJSON: northflank.String("occaecati"),
                SecretKey: northflank.String("fugit"),
            },
            Description: northflank.String("This is a new description."),
        },
        IntegrationID: "gcp-integration-1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateIntegrationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateIntegrationRequest](../../models/operations/updateintegrationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateIntegrationResponse](../../models/operations/updateintegrationresponse.md), error**

