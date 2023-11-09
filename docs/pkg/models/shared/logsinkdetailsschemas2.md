# LogSinkDetailsSchemas2

Authenticate with a token.


## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `AuthenticationStrategy`                                                                                          | [shared.LogSinkDetailsAuthenticationStrategy](../../../pkg/models/shared/logsinkdetailsauthenticationstrategy.md) | :heavy_check_mark:                                                                                                | The authentication strategy.                                                                                      | token                                                                                                             |
| `Token`                                                                                                           | *string*                                                                                                          | :heavy_check_mark:                                                                                                | The HTTP Token for the Papertrail log destination.                                                                | 123abcdefghijklmnopqrstuvwxy                                                                                      |
| `URI`                                                                                                             | *string*                                                                                                          | :heavy_check_mark:                                                                                                | The uri for the Papertrail log destination.                                                                       | https://logs.collector.solarwinds.com/v1/log                                                                      |