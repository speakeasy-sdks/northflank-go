# CreateDomainResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `APIErrorResult`                                                        | [*shared.APIErrorResult](../../models/shared/apierrorresult.md)         | :heavy_minus_sign:                                                      | The domain is not valid, possibly because it is too long.               |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `CreateDomainResult`                                                    | [*shared.CreateDomainResult](../../models/shared/createdomainresult.md) | :heavy_minus_sign:                                                      | Details about the newly added domain.                                   |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |