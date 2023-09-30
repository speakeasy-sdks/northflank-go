# CreateDomainResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `APIErrorResult`                                                        | [*shared.APIErrorResult](../../models/shared/apierrorresult.md)         | :heavy_minus_sign:                                                      | The domain is not valid, possibly because it is too long.               |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | HTTP response content type for this operation                           |
| `CreateDomainResult`                                                    | [*shared.CreateDomainResult](../../models/shared/createdomainresult.md) | :heavy_minus_sign:                                                      | Details about the newly added domain.                                   |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | HTTP response status code for this operation                            |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | Raw HTTP response; suitable for custom response parsing                 |