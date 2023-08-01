# VerifySubDomainResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `APIErrorResult`                                                | [*shared.APIErrorResult](../../models/shared/apierrorresult.md) | :heavy_minus_sign:                                              | Failed to verify the subdomain.                                 |
| `ContentType`                                                   | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `StatusCode`                                                    | *int*                                                           | :heavy_check_mark:                                              | N/A                                                             |
| `RawResponse`                                                   | [*http.Response](https://pkg.go.dev/net/http#Response)          | :heavy_minus_sign:                                              | N/A                                                             |
| `SuccessResult`                                                 | [*shared.SuccessResult](../../models/shared/successresult.md)   | :heavy_minus_sign:                                              | The operation was performed successfully.                       |