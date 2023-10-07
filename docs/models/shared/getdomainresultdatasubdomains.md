# GetDomainResultDataSubdomains

Details about a subdomain.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AdditionalProperties`                                    | map[string]*interface{}*                                  | :heavy_minus_sign:                                        | N/A                                                       |                                                           |
| `FullName`                                                | *string*                                                  | :heavy_check_mark:                                        | The full domain including the subdomain.                  | app.example.com                                           |
| `Name`                                                    | *string*                                                  | :heavy_check_mark:                                        | The subdomain added, or -default for the empty subdomain. | app                                                       |