# InvoiceDetailsResultPrice

Details about the price of a object, broken down by resource type.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `CPU`                                             | *float32*                                         | :heavy_check_mark:                                | Price of CPU usage for this object, in cents.     | 200                                               |
| `Memory`                                          | *float32*                                         | :heavy_check_mark:                                | Price of memory usage for this object, in cents.  | 200                                               |
| `Storage`                                         | *float32*                                         | :heavy_check_mark:                                | Price of storage usage for this object, in cents. | 200                                               |
| `Total`                                           | *float32*                                         | :heavy_check_mark:                                | Total price for this object, in cents.            | 600                                               |