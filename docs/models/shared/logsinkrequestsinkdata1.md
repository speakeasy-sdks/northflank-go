# LogSinkRequestSinkData1

Loki Sink Schema.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `AdditionalProperties`                                                             | map[string]*interface{}*                                                           | :heavy_minus_sign:                                                                 | N/A                                                                                |                                                                                    |
| `Auth`                                                                             | [*LogSinkRequestSinkData1Auth](../../models/shared/logsinkrequestsinkdata1auth.md) | :heavy_minus_sign:                                                                 | Object containing authentication data for the log sink.                            |                                                                                    |
| `Endpoint`                                                                         | **string*                                                                          | :heavy_minus_sign:                                                                 | The endpoint of the Loki log sink.                                                 | https://logs.example.com                                                           |