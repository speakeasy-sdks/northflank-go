# ListIntegrationsResultPagination

Data about the endpoint pagination.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `AdditionalProperties`                          | map[string]*interface{}*                        | :heavy_minus_sign:                              | N/A                                             |                                                 |
| `Count`                                         | *float32*                                       | :heavy_check_mark:                              | The number of results returned by this request. | 1                                               |
| `Cursor`                                        | **string*                                       | :heavy_minus_sign:                              | The cursor to access the next page of results.  |                                                 |
| `HasNextPage`                                   | *bool*                                          | :heavy_check_mark:                              | Is there another page of results available?     | false                                           |