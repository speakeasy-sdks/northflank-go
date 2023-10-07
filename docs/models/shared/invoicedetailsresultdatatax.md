# InvoiceDetailsResultDataTax

Details about the tax on the invoice.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `AdditionalProperties`                       | map[string]*interface{}*                     | :heavy_minus_sign:                           | N/A                                          |                                              |
| `Percent`                                    | *float32*                                    | :heavy_check_mark:                           | Percentage of subtotal to be applied as tax. | 20                                           |
| `Value`                                      | *float32*                                    | :heavy_check_mark:                           | Value of the tax on the invoice.             | 200                                          |