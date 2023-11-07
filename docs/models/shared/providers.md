# Providers

A provider object


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `DiskSizes`                                                     | []*float32*                                                     | :heavy_check_mark:                                              | An array of supported node disk sizes                           | 100                                                             |
| `Engine`                                                        | *string*                                                        | :heavy_check_mark:                                              | The kubernetes engine used.                                     | Elastic Kubernetes Service (EKS)                                |
| `Flags`                                                         | [shared.Flags](../../models/shared/flags.md)                    | :heavy_check_mark:                                              | An object with feature flags to indicate (un)supported features |                                                                 |
| `ID`                                                            | *string*                                                        | :heavy_check_mark:                                              | The ID of the provider.                                         | aws                                                             |
| `KubernetesVersions`                                            | []*string*                                                      | :heavy_check_mark:                                              | An array of available kubernetes versions                       | 1.21                                                            |
| `Name`                                                          | *string*                                                        | :heavy_check_mark:                                              | The name of the provider.                                       | Amazon Web Services                                             |
| `NodeTypes`                                                     | []*string*                                                      | :heavy_check_mark:                                              | An array of supported node types                                | m5.2xlarge                                                      |
| `Regions`                                                       | []*string*                                                      | :heavy_check_mark:                                              | An array of available regions                                   | eu-west-2                                                       |