# APIErrorResult

API Error response model


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Code`                                                   | *string*                                                 | :heavy_check_mark:                                       | Error code                                               |
| `Data`                                                   | [*sdkerrors.Data](../../../pkg/models/sdkerrors/data.md) | :heavy_minus_sign:                                       | N/A                                                      |
| `Message`                                                | *string*                                                 | :heavy_check_mark:                                       | Error description                                        |
| `Status`                                                 | *float32*                                                | :heavy_check_mark:                                       | Status code                                              |