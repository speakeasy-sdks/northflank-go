# GetVCSProvidersResponse


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ContentType`                                                                  | *string*                                                                       | :heavy_check_mark:                                                             | HTTP response content type for this operation                                  |
| `StatusCode`                                                                   | *int*                                                                          | :heavy_check_mark:                                                             | HTTP response status code for this operation                                   |
| `RawResponse`                                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)                         | :heavy_minus_sign:                                                             | Raw HTTP response; suitable for custom response parsing                        |
| `VCSProvidersResult`                                                           | [*shared.VCSProvidersResult](../../../pkg/models/shared/vcsprovidersresult.md) | :heavy_minus_sign:                                                             | Details about the version control providers available for use.                 |