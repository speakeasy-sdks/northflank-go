# LogSinkDetails1

Data about the log sink.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Auth`                                                                          | [*shared.LogSinkDetailsAuth](../../models/shared/logsinkdetailsauth.md)         | :heavy_minus_sign:                                                              | Object containing authentication data for the log sink.                         |                                                                                 |
| `Encoding`                                                                      | [*shared.LogSinkDetailsEncoding](../../models/shared/logsinkdetailsencoding.md) | :heavy_minus_sign:                                                              | Encoding options                                                                |                                                                                 |
| `Endpoint`                                                                      | *string*                                                                        | :heavy_check_mark:                                                              | The endpoint of the Loki log sink.                                              | https://logs.example.com                                                        |