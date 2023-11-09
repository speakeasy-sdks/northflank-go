# LogioLogSink

Create a log sink using Logz.io


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Description`                                                                                            | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Description of the log sink.                                                                             | This is an example log sink.                                                                             |
| `ForwardAccessLogs`                                                                                      | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If `true` your network access logs will be forwarded with your workload logs                             | true                                                                                                     |
| `Name`                                                                                                   | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Name of the log sink.                                                                                    | example-log-sink                                                                                         |
| `Projects`                                                                                               | []*string*                                                                                               | :heavy_minus_sign:                                                                                       | If `restricted` is `true`, only logs from these projects will be sent to the log sink.                   | default-project                                                                                          |
| `Restricted`                                                                                             | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If `true`, only logs from the projects in `projects` will be sent to the log sink.                       | true                                                                                                     |
| `SinkData`                                                                                               | [shared.LogioLogSinkSinkData](../../../pkg/models/shared/logiologsinksinkdata.md)                        | :heavy_check_mark:                                                                                       | Details about the Logz.io log sink.                                                                      |                                                                                                          |
| `SinkType`                                                                                               | [shared.LogioLogSinkSinkType](../../../pkg/models/shared/logiologsinksinktype.md)                        | :heavy_check_mark:                                                                                       | The type of the log sink.                                                                                | logzio                                                                                                   |
| `UseCustomLabels`                                                                                        | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels | true                                                                                                     |