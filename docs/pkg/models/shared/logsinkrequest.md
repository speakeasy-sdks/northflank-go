# LogSinkRequest


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Projects`                                                                                               | []*string*                                                                                               | :heavy_minus_sign:                                                                                       | If `restricted` is `true`, only logs from these projects will be sent to the log sink.                   | default-project                                                                                          |
| `Restricted`                                                                                             | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If `true`, only logs from the projects in `projects` will be sent to the log sink.                       | true                                                                                                     |
| `ResumeLogSink`                                                                                          | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If `true`, and the log sink is currently paused, the log sink will be resumed after updating.            | false                                                                                                    |
| `SinkData`                                                                                               | [*shared.LogSinkRequestSinkData](../../../pkg/models/shared/logsinkrequestsinkdata.md)                   | :heavy_minus_sign:                                                                                       | Data about the log sink.                                                                                 |                                                                                                          |
| `UseCustomLabels`                                                                                        | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels | true                                                                                                     |