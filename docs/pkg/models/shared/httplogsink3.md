# HTTPLogSink3

Authenticate with a bearer token strategy.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Strategy`                                                                                    | [shared.HTTPLogSinkSchemasStrategy](../../../pkg/models/shared/httplogsinkschemasstrategy.md) | :heavy_check_mark:                                                                            | Bearer token authentication strategy.                                                         | bearer                                                                                        |
| `Token`                                                                                       | **string*                                                                                     | :heavy_minus_sign:                                                                            | Token for bearer token authentication.                                                        | my-token                                                                                      |