# HTTPLogSinkSinkData

Details about the HTTP log sink.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Auth`                                                                  | [shared.HTTPLogSinkAuth](../../../pkg/models/shared/httplogsinkauth.md) | :heavy_check_mark:                                                      | N/A                                                                     |                                                                         |
| `Encoding`                                                              | [*shared.Encoding](../../../pkg/models/shared/encoding.md)              | :heavy_minus_sign:                                                      | Encoding options                                                        |                                                                         |
| `URI`                                                                   | *string*                                                                | :heavy_check_mark:                                                      | Uri to send logs to.                                                    | my.log-collector.com                                                    |