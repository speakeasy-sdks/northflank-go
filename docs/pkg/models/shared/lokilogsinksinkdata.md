# LokiLogSinkSinkData

Details about the Loki log sink.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Auth`                                                                           | [*shared.LokiLogSinkAuth](../../../pkg/models/shared/lokilogsinkauth.md)         | :heavy_minus_sign:                                                               | Object containing authentication data for the log sink.                          |                                                                                  |
| `Encoding`                                                                       | [*shared.LokiLogSinkEncoding](../../../pkg/models/shared/lokilogsinkencoding.md) | :heavy_minus_sign:                                                               | Encoding options                                                                 |                                                                                  |
| `Endpoint`                                                                       | *string*                                                                         | :heavy_check_mark:                                                               | The endpoint of the Loki log sink.                                               | https://logs.example.com                                                         |