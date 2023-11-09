# Service


## Fields

| Field                                                                                                                              | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `Duration`                                                                                                                         | *float32*                                                                                                                          | :heavy_check_mark:                                                                                                                 | Duration services have been running in this billing period, in seconds.                                                            | 172800                                                                                                                             |
| `Price`                                                                                                                            | [*shared.InvoiceDetailsResultSchemasDataProjectsPrice](../../../pkg/models/shared/invoicedetailsresultschemasdataprojectsprice.md) | :heavy_minus_sign:                                                                                                                 | Details about the price for all services in this project, broken down by resource type.                                            |                                                                                                                                    |