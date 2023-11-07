# InvoiceDetailsResultSchemasPrice

Details about the price for all addons in this project, broken down by resource type.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CPU`                                                            | *float32*                                                        | :heavy_check_mark:                                               | Price of CPU usage for all addons in this project, in cents.     | 200                                                              |
| `Memory`                                                         | *float32*                                                        | :heavy_check_mark:                                               | Price of memory usage for all addons in this project, in cents.  | 200                                                              |
| `Storage`                                                        | *float32*                                                        | :heavy_check_mark:                                               | Price of storage usage for all addons in this project, in cents. | 200                                                              |
| `Total`                                                          | *float32*                                                        | :heavy_check_mark:                                               | Total price for all addons in this project, in cents.            | 600                                                              |