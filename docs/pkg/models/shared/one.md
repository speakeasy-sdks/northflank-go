# One

Loki Sink Schema.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Auth`                                                                         | [*shared.LogSinkRequestAuth](../../../pkg/models/shared/logsinkrequestauth.md) | :heavy_minus_sign:                                                             | Object containing authentication data for the log sink.                        |                                                                                |
| `Endpoint`                                                                     | **string*                                                                      | :heavy_minus_sign:                                                             | The endpoint of the Loki log sink.                                             | https://logs.example.com                                                       |