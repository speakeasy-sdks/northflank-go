// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
)

// AWSLogSinkSinkDataAuth - Authentication object.
type AWSLogSinkSinkDataAuth struct {
	// Access key id for the bucket.
	AccessKeyID string `json:"accessKeyId"`
	// Secret access key for the bucket.
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *AWSLogSinkSinkDataAuth) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *AWSLogSinkSinkDataAuth) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// AWSLogSinkSinkDataCompression - Log file compression method.
type AWSLogSinkSinkDataCompression string

const (
	AWSLogSinkSinkDataCompressionGzip AWSLogSinkSinkDataCompression = "gzip"
	AWSLogSinkSinkDataCompressionNone AWSLogSinkSinkDataCompression = "none"
)

func (e AWSLogSinkSinkDataCompression) ToPointer() *AWSLogSinkSinkDataCompression {
	return &e
}

func (e *AWSLogSinkSinkDataCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gzip":
		fallthrough
	case "none":
		*e = AWSLogSinkSinkDataCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AWSLogSinkSinkDataCompression: %v", v)
	}
}

// AWSLogSinkSinkDataRegion - Region of the S3 bucket.
type AWSLogSinkSinkDataRegion string

const (
	AWSLogSinkSinkDataRegionEuWest1    AWSLogSinkSinkDataRegion = "eu-west-1"
	AWSLogSinkSinkDataRegionEuWest2    AWSLogSinkSinkDataRegion = "eu-west-2"
	AWSLogSinkSinkDataRegionEuWest3    AWSLogSinkSinkDataRegion = "eu-west-3"
	AWSLogSinkSinkDataRegionEuCentral1 AWSLogSinkSinkDataRegion = "eu-central-1"
	AWSLogSinkSinkDataRegionEuSouth1   AWSLogSinkSinkDataRegion = "eu-south-1"
	AWSLogSinkSinkDataRegionEuNorth1   AWSLogSinkSinkDataRegion = "eu-north-1"
	AWSLogSinkSinkDataRegionUsWest1    AWSLogSinkSinkDataRegion = "us-west-1"
	AWSLogSinkSinkDataRegionUsWest2    AWSLogSinkSinkDataRegion = "us-west-2"
	AWSLogSinkSinkDataRegionUsEast1    AWSLogSinkSinkDataRegion = "us-east-1"
	AWSLogSinkSinkDataRegionUsEast2    AWSLogSinkSinkDataRegion = "us-east2"
)

func (e AWSLogSinkSinkDataRegion) ToPointer() *AWSLogSinkSinkDataRegion {
	return &e
}

func (e *AWSLogSinkSinkDataRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east2":
		*e = AWSLogSinkSinkDataRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AWSLogSinkSinkDataRegion: %v", v)
	}
}

// AWSLogSinkSinkData - Details about the AWS S3 log sink.
type AWSLogSinkSinkData struct {
	// Authentication object.
	Auth *AWSLogSinkSinkDataAuth `json:"auth,omitempty"`
	// Name of the S3 Bucket.
	Bucket string `json:"bucket"`
	// Log file compression method.
	Compression AWSLogSinkSinkDataCompression `json:"compression"`
	// Endpoint for the AWS S3 or compatible API bucket.
	Endpoint string `json:"endpoint"`
	// Region of the S3 bucket.
	Region AWSLogSinkSinkDataRegion `json:"region"`
}

func (o *AWSLogSinkSinkData) GetAuth() *AWSLogSinkSinkDataAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *AWSLogSinkSinkData) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *AWSLogSinkSinkData) GetCompression() AWSLogSinkSinkDataCompression {
	if o == nil {
		return AWSLogSinkSinkDataCompression("")
	}
	return o.Compression
}

func (o *AWSLogSinkSinkData) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *AWSLogSinkSinkData) GetRegion() AWSLogSinkSinkDataRegion {
	if o == nil {
		return AWSLogSinkSinkDataRegion("")
	}
	return o.Region
}

// AWSLogSinkSinkType - The type of the log sink.
type AWSLogSinkSinkType string

const (
	AWSLogSinkSinkTypeAwsS3 AWSLogSinkSinkType = "aws_s3"
)

func (e AWSLogSinkSinkType) ToPointer() *AWSLogSinkSinkType {
	return &e
}

func (e *AWSLogSinkSinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aws_s3":
		*e = AWSLogSinkSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AWSLogSinkSinkType: %v", v)
	}
}

// AWSLogSink - Create a log sink using AWS S3
type AWSLogSink struct {
	// Description of the log sink.
	Description *string `json:"description,omitempty"`
	// If `true` your network access logs will be forwarded with your workload logs
	ForwardAccessLogs *bool `default:"false" json:"forwardAccessLogs"`
	// Name of the log sink.
	Name string `json:"name"`
	// If `restricted` is `true`, only logs from these projects will be sent to the log sink.
	Projects []string `json:"projects,omitempty"`
	// If `true`, only logs from the projects in `projects` will be sent to the log sink.
	Restricted *bool `default:"false" json:"restricted"`
	// Details about the AWS S3 log sink.
	SinkData AWSLogSinkSinkData `json:"sinkData"`
	// The type of the log sink.
	SinkType AWSLogSinkSinkType `json:"sinkType"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `default:"false" json:"useCustomLabels"`
}

func (a AWSLogSink) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AWSLogSink) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AWSLogSink) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AWSLogSink) GetForwardAccessLogs() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardAccessLogs
}

func (o *AWSLogSink) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AWSLogSink) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *AWSLogSink) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *AWSLogSink) GetSinkData() AWSLogSinkSinkData {
	if o == nil {
		return AWSLogSinkSinkData{}
	}
	return o.SinkData
}

func (o *AWSLogSink) GetSinkType() AWSLogSinkSinkType {
	if o == nil {
		return AWSLogSinkSinkType("")
	}
	return o.SinkType
}

func (o *AWSLogSink) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}
