# InvoiceDetailsResultData

Result data.


## Fields

| Field                                                                                                                                           | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     | Example                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `Currency`                                                                                                                                      | *string*                                                                                                                                        | :heavy_check_mark:                                                                                                                              | Currency code for the currency the invoice is billed in.                                                                                        | usd                                                                                                                                             |
| `Paid`                                                                                                                                          | **bool*                                                                                                                                         | :heavy_minus_sign:                                                                                                                              | If `timestamp` is passed in, whether the invoice has been paid.                                                                                 |                                                                                                                                                 |
| `Period`                                                                                                                                        | [shared.Period](../../../pkg/models/shared/period.md)                                                                                           | :heavy_check_mark:                                                                                                                              | Information about the billing period of the invoice.                                                                                            |                                                                                                                                                 |
| `Projects`                                                                                                                                      | [][shared.Projects](../../../pkg/models/shared/projects.md)                                                                                     | :heavy_check_mark:                                                                                                                              | An array of projects billed in this invoice. If `projectId` is passed in, only projects with a `projectId` matching the value will be returned. |                                                                                                                                                 |
| `SubTotal`                                                                                                                                      | *float32*                                                                                                                                       | :heavy_check_mark:                                                                                                                              | Total cost of the invoice, in cents, excluding tax.                                                                                             | 1000                                                                                                                                            |
| `Tax`                                                                                                                                           | [shared.Tax](../../../pkg/models/shared/tax.md)                                                                                                 | :heavy_check_mark:                                                                                                                              | Details about the tax on the invoice.                                                                                                           |                                                                                                                                                 |
| `Total`                                                                                                                                         | *float32*                                                                                                                                       | :heavy_check_mark:                                                                                                                              | Total cost of the invoice, in cents, including tax.                                                                                             | 1200                                                                                                                                            |