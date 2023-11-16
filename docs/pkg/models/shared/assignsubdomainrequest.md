# AssignSubDomainRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `PortName`                                                   | *string*                                                     | :heavy_check_mark:                                           | The name of the port that will be assigned to the subdomain. | port-1                                                       |
| `ProjectID`                                                  | *string*                                                     | :heavy_check_mark:                                           | The ID of the project the service belongs to.                | default-project                                              |
| `ServiceID`                                                  | *string*                                                     | :heavy_check_mark:                                           | The ID of the service to assign the subdomain to.            | example-service                                              |