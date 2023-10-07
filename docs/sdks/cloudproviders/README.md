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
	"context"
	"log"
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.CloudProviders.Create(ctx, shared.CreateIntegrationRequest{
        AdditionalProperties: map[string]interface{}{
            "online": "Configuration",
        },
        Credentials: shared.CreateIntegrationRequestCredentials{
            AdditionalProperties: map[string]interface{}{
                "Money": "blue",
            },
        },
        Description: northflankgo.String("This is a new cloud provider integration."),
        Gcp: &shared.CreateIntegrationRequestGcp{
            AdditionalProperties: map[string]interface{}{
                "shred": "abnormally",
            },
            ProjectID: "orange Northwest",
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
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.CloudProviders.CreateCluster(ctx, shared.CreateClusterRequest{
        AdditionalProperties: map[string]interface{}{
            "Darmstadtium": "enormously",
        },
        Description: northflankgo.String("This is a new cluster."),
        Gcp: &shared.CreateClusterRequestGcp{
            AdditionalProperties: map[string]interface{}{
                "Minnesota": "male",
            },
            ProjectID: "example-project-id",
        },
        Integration: &shared.CreateClusterRequestIntegration{
            AdditionalProperties: map[string]interface{}{
                "Oregon": "Lodi",
            },
        },
        IntegrationID: northflankgo.String("gcp-integration"),
        KubernetesVersion: "1.23.8",
        Name: "GCP Cluster 1",
        NodePools: []shared.CreateClusterRequestNodePools{
            shared.CreateClusterRequestNodePools{
                AdditionalProperties: map[string]interface{}{
                    "Towels": "JBOD",
                },
                Autoscaling: &shared.CreateClusterRequestNodePoolsAutoscaling{
                    AdditionalProperties: map[string]interface{}{
                        "male": "US",
                    },
                    Enabled: northflankgo.Bool(true),
                    Max: northflankgo.Int64(10),
                    Min: northflankgo.Int64(1),
                },
                AvailabilityZones: []string{
                    "strategize",
                },
                DiskSize: 100,
                Labels: map[string]interface{}{
                    "Ford": "Estate",
                },
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflankgo.Bool(false),
            },
        },
        Provider: shared.CreateClusterRequestProviderGcp,
        Region: "europe-west2",
        Settings: shared.CreateClusterRequestSettings{
            AdditionalProperties: map[string]interface{}{
                "Southeast": "solutions",
            },
            Builds: &shared.CreateClusterRequestSettingsBuilds{
                AdditionalProperties: map[string]interface{}{
                    "integrate": "Montana",
                },
                ClusterID: northflankgo.String("build-cluster"),
                Plan: northflankgo.String("nf-compute-200"),
            },
            Logging: &shared.CreateClusterRequestSettingsLogging{
                AdditionalProperties: map[string]interface{}{
                    "Cambridgeshire": "um",
                },
                Loki: shared.CreateCreateClusterRequestSettingsLoggingLokiCreateClusterRequestSettingsLoggingLoki1(
                        shared.CreateClusterRequestSettingsLoggingLoki1{
                            AdditionalProperties: map[string]interface{}{
                                "matrices": "Carolina",
                            },
                            S3AccessKey: "payment",
                            S3BucketName: "capacity",
                            S3Region: "protocol paragraph",
                            S3SecretKey: "blah",
                        },
                ),
            },
            Registry: &shared.CreateClusterRequestSettingsRegistry{
                AdditionalProperties: map[string]interface{}{
                    "invoice": "gloomy",
                },
                RegistryID: northflankgo.String("my-registry-credentials"),
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
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
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

**[*operations.DeleteClusterResponse](../../models/operations/deleteclusterresponse.md), error**


## DeleteIntegration

Delete the given integration. Fails if the integration is associated with existing clusters.

### Example Usage

```go
package main

import(
	"context"
	"log"
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
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

**[*operations.DeleteIntegrationResponse](../../models/operations/deleteintegrationresponse.md), error**


## Get

Lists supported cloud providers

### Example Usage

```go
package main

import(
	"context"
	"log"
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
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
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
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

**[*operations.GetClusterResponse](../../models/operations/getclusterresponse.md), error**


## GetIntegration

Get information about the given integration

### Example Usage

```go
package main

import(
	"context"
	"log"
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
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

**[*operations.GetIntegrationResponse](../../models/operations/getintegrationresponse.md), error**


## ListClusters

Lists clusters for the authenticated user or team.

### Example Usage

```go
package main

import(
	"context"
	"log"
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )
    var cursor *string = "24"
    var page *int64 = 1
    var perPage *int64 = 50

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
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )
    var cursor *string = "South"
    var page *int64 = 1
    var perPage *int64 = 50

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
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )
    updateClusterRequest := shared.UpdateClusterRequest{
        AdditionalProperties: map[string]interface{}{
            "Fish": "Sausages",
        },
        Description: northflankgo.String("This is an updated description."),
        NodePools: []shared.UpdateClusterRequestNodePools{
            shared.UpdateClusterRequestNodePools{
                AdditionalProperties: map[string]interface{}{
                    "defect": "sievert",
                },
                Autoscaling: &shared.UpdateClusterRequestNodePoolsAutoscaling{
                    AdditionalProperties: map[string]interface{}{
                        "Corporate": "Southeast",
                    },
                    Enabled: northflankgo.Bool(true),
                    Max: northflankgo.Int64(10),
                    Min: northflankgo.Int64(1),
                },
                AvailabilityZones: []string{
                    "Metrics",
                },
                DiskSize: 100,
                ID: northflankgo.String("6aa96121-0345-43ad-bade-af36d540c222"),
                Labels: map[string]interface{}{
                    "Communications": "Solutions",
                },
                NodeCount: 3,
                NodeType: "n2-standard-8",
                Preemptible: northflankgo.Bool(false),
            },
        },
        Settings: &shared.UpdateClusterRequestSettings{
            AdditionalProperties: map[string]interface{}{
                "Southeast": "Cambridgeshire",
            },
            Builds: &shared.UpdateClusterRequestSettingsBuilds{
                AdditionalProperties: map[string]interface{}{
                    "Roentgenium": "Associate",
                },
                ClusterID: northflankgo.String("build-cluster"),
                Plan: northflankgo.String("nf-compute-200"),
            },
            Logging: &shared.UpdateClusterRequestSettingsLogging{
                AdditionalProperties: map[string]interface{}{
                    "actuating": "Frozen",
                },
                Loki: shared.CreateUpdateClusterRequestSettingsLoggingLokiUpdateClusterRequestSettingsLoggingLoki1(
                        shared.UpdateClusterRequestSettingsLoggingLoki1{
                            AdditionalProperties: map[string]interface{}{
                                "Senior": "exactly",
                            },
                            S3AccessKey: "Gabon Croatian Brand",
                            S3BucketName: "Minivan Fresh Neon",
                            S3Region: "South Polestar monitor",
                            S3SecretKey: "bypass Outdoors invoice",
                        },
                ),
            },
            Registry: &shared.UpdateClusterRequestSettingsRegistry{
                AdditionalProperties: map[string]interface{}{
                    "Chino": "Bedfordshire",
                },
                RegistryID: northflankgo.String("my-registry-credentials"),
            },
        },
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
	northflankgo "github.com/speakeasy-sdks/northflank-go"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
)

func main() {
    s := northflankgo.New(
        northflankgo.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )
    updateIntegrationRequest := shared.UpdateIntegrationRequest{
        AdditionalProperties: map[string]interface{}{
            "before": "Bronze",
        },
        Credentials: shared.UpdateIntegrationRequestCredentials{
            AdditionalProperties: map[string]interface{}{
                "Hendersonville": "interactive",
            },
        },
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `updateIntegrationRequest`                                                         | [shared.UpdateIntegrationRequest](../../models/shared/updateintegrationrequest.md) | :heavy_check_mark:                                                                 | Request body                                                                       |                                                                                    |
| `integrationID`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | gcp-integration-1                                                                  |


### Response

**[*operations.UpdateIntegrationResponse](../../models/operations/updateintegrationresponse.md), error**

