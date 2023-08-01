# GetIntegrationResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ContentType`                                                               | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `GetIntegrationResult`                                                      | [*shared.GetIntegrationResult](../../models/shared/getintegrationresult.md) | :heavy_minus_sign:                                                          | Details about the given integration.                                        |
| `StatusCode`                                                                | *int*                                                                       | :heavy_check_mark:                                                          | N/A                                                                         |
| `RawResponse`                                                               | [*http.Response](https://pkg.go.dev/net/http#Response)                      | :heavy_minus_sign:                                                          | N/A                                                                         |