# LogSinkRequestSinkData31

Authenticate with a host/port


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `AuthenticationStrategy`                     | *string*                                     | :heavy_check_mark:                           | The authentication strategy.                 | port                                         |
| `Host`                                       | *string*                                     | :heavy_check_mark:                           | The host for the Papertrail log destination. | logs1.papertrailapp.com:                     |
| `Port`                                       | *float32*                                    | :heavy_check_mark:                           | The port for the Papertrail log destination. | 8000                                         |