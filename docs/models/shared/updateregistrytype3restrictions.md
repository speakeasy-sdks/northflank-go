# UpdateRegistryType3Restrictions

Data about whether the credentials are restricted to certain projects.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AdditionalProperties`                                                 | map[string]*interface{}*                                               | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |
| `Projects`                                                             | []*string*                                                             | :heavy_minus_sign:                                                     | An array of projects the credentials are restricted to, if applicable. | default-project                                                        |
| `Restricted`                                                           | *bool*                                                                 | :heavy_check_mark:                                                     | Whether the credentials are restricted to specific projects.           | true                                                                   |