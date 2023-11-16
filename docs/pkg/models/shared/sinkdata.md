# SinkData

Details about the AWS S3 log sink.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Auth`                                                          | [*shared.Auth](../../../pkg/models/shared/auth.md)              | :heavy_minus_sign:                                              | Authentication object.                                          |                                                                 |
| `Bucket`                                                        | *string*                                                        | :heavy_check_mark:                                              | Name of the S3 Bucket.                                          | northflank-logs                                                 |
| `Compression`                                                   | [shared.Compression](../../../pkg/models/shared/compression.md) | :heavy_check_mark:                                              | Log file compression method.                                    | gzip                                                            |
| `Endpoint`                                                      | *string*                                                        | :heavy_check_mark:                                              | Endpoint for the AWS S3 or compatible API bucket.               | my.bucket.com                                                   |
| `Region`                                                        | [shared.Region](../../../pkg/models/shared/region.md)           | :heavy_check_mark:                                              | Region of the S3 bucket.                                        | eu-west-2                                                       |