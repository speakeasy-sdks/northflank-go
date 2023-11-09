# UpdateRegistryType2


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Description`                                                                                            | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Description of the credentials.                                                                          | This is a set of saved credentials.                                                                      |
| `Password`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Password, Personal Access Token, or API key for the container registry.                                  | password1234                                                                                             |
| `RegistryURL`                                                                                            | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Custom url for the container registry. Only usable (and required) when `provider` is `custom`.           | https://example.com                                                                                      |
| `Restrictions`                                                                                           | [*shared.UpdateRegistryType2Restrictions](../../../pkg/models/shared/updateregistrytype2restrictions.md) | :heavy_minus_sign:                                                                                       | Data about whether the credentials are restricted to certain projects.                                   |                                                                                                          |
| `Username`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Username for the container registry.                                                                     | test-user                                                                                                |