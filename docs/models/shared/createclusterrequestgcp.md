# CreateClusterRequestGcp

GCP specific data. Required when `provider` is `gcp`.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AdditionalProperties`                                    | map[string]*interface{}*                                  | :heavy_minus_sign:                                        | N/A                                                       |                                                           |
| `ProjectID`                                               | *string*                                                  | :heavy_check_mark:                                        | ID of the GCP project the cluster will be provisioned in. | example-project-id                                        |