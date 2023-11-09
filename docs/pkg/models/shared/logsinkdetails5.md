# LogSinkDetails5

AWS S3 or compatible API Sink Schema.


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `Auth`                                                                                               | [*shared.LogSinkDetailsSchemasDataAuth](../../../pkg/models/shared/logsinkdetailsschemasdataauth.md) | :heavy_minus_sign:                                                                                   | Authentication object.                                                                               |                                                                                                      |
| `Bucket`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | Name of the S3 Bucket.                                                                               | northflank-logs                                                                                      |
| `Compression`                                                                                        | [shared.LogSinkDetailsCompression](../../../pkg/models/shared/logsinkdetailscompression.md)          | :heavy_check_mark:                                                                                   | Log file compression method.                                                                         | gzip                                                                                                 |
| `Endpoint`                                                                                           | *string*                                                                                             | :heavy_check_mark:                                                                                   | Endpoint for the AWS S3 or compatible API bucket.                                                    | my.bucket.com                                                                                        |
| `Region`                                                                                             | [shared.LogSinkDetailsSchemasRegion](../../../pkg/models/shared/logsinkdetailsschemasregion.md)      | :heavy_check_mark:                                                                                   | Region of the S3 bucket.                                                                             | eu-west-2                                                                                            |