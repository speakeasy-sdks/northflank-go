# ListAddonTypesResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AddonTypesResult`                                                         | [*shared.AddonTypesResult](../../../pkg/models/shared/addontypesresult.md) | :heavy_minus_sign:                                                         | A list of the available addon types.                                       |
| `ContentType`                                                              | *string*                                                                   | :heavy_check_mark:                                                         | HTTP response content type for this operation                              |
| `StatusCode`                                                               | *int*                                                                      | :heavy_check_mark:                                                         | HTTP response status code for this operation                               |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_check_mark:                                                         | Raw HTTP response; suitable for custom response parsing                    |