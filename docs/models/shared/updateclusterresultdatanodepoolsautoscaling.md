# UpdateClusterResultDataNodePoolsAutoscaling

Auto scaling settings to use for the node pool. Requires that the cloud provider supports this feature.


## Fields

| Field              | Type               | Required           | Description        | Example            |
| ------------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| `Enabled`          | **bool*            | :heavy_minus_sign: | N/A                | true               |
| `Max`              | **int64*           | :heavy_minus_sign: | N/A                | 10                 |
| `Min`              | **int64*           | :heavy_minus_sign: | N/A                | 1                  |