# GetInvoiceDetailsResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ContentType`                                                               | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `InvoiceDetailsResult`                                                      | [*shared.InvoiceDetailsResult](../../models/shared/invoicedetailsresult.md) | :heavy_minus_sign:                                                          | Details about an invoice.                                                   |
| `StatusCode`                                                                | *int*                                                                       | :heavy_check_mark:                                                          | N/A                                                                         |
| `RawResponse`                                                               | [*http.Response](https://pkg.go.dev/net/http#Response)                      | :heavy_minus_sign:                                                          | N/A                                                                         |