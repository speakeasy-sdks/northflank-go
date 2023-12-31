# HTTPLogSink2

Authenticate with a basic http strategy.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Password`                                                                      | *string*                                                                        | :heavy_check_mark:                                                              | Password for basic http authentication.                                         | secret-password                                                                 |
| `Strategy`                                                                      | [shared.HTTPLogSinkStrategy](../../../pkg/models/shared/httplogsinkstrategy.md) | :heavy_check_mark:                                                              | Basic HTTP authentication strategy.                                             | basic                                                                           |
| `User`                                                                          | **string*                                                                       | :heavy_minus_sign:                                                              | Username for basic http authentication.                                         | my-user                                                                         |