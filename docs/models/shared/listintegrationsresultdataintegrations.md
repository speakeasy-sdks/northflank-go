# ListIntegrationsResultDataIntegrations

An integration object.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `AdditionalProperties`                                   | map[string]*interface{}*                                 | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `CreatedAt`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | The time the integration was created.                    | 2021-01-20T11:19:53.175Z                                 |
| `Description`                                            | **string*                                                | :heavy_minus_sign:                                       | A short description of the integration.                  | The integration description                              |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | Identifier for the integration.                          | gcp-integration                                          |
| `Name`                                                   | *string*                                                 | :heavy_check_mark:                                       | The name of the integration.                             | GCP integration                                          |
| `Provider`                                               | **string*                                                | :heavy_minus_sign:                                       | The cloud provider to which this integration belongs to. | gcp                                                      |