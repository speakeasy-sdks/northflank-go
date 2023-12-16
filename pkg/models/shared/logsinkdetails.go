// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/utils"
	"time"
)

// LogSinkDetails8 - Honeycomb Sink Schema.
type LogSinkDetails8 struct {
	// Honeycomb API Key
	APIKey string `json:"api_key"`
	// Name of the dataset
	Dataset string `json:"dataset"`
}

func (o *LogSinkDetails8) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *LogSinkDetails8) GetDataset() string {
	if o == nil {
		return ""
	}
	return o.Dataset
}

// LogSinkDetails7 - LogDNA Sink Schema.
type LogSinkDetails7 struct {
	// Ingestion Key
	APIKey string `json:"api_key"`
}

func (o *LogSinkDetails7) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

// LogSinkDetails6 - Logtail Sink Schema.
type LogSinkDetails6 struct {
	// Logtail Source Token
	Token string `json:"token"`
}

func (o *LogSinkDetails6) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// LogSinkDetailsSchemasDataAuth - Authentication object.
type LogSinkDetailsSchemasDataAuth struct {
	// Access key id for the bucket.
	AccessKeyID string `json:"accessKeyId"`
	// Secret access key for the bucket.
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *LogSinkDetailsSchemasDataAuth) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *LogSinkDetailsSchemasDataAuth) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// LogSinkDetailsCompression - Log file compression method.
type LogSinkDetailsCompression string

const (
	LogSinkDetailsCompressionGzip LogSinkDetailsCompression = "gzip"
	LogSinkDetailsCompressionNone LogSinkDetailsCompression = "none"
)

func (e LogSinkDetailsCompression) ToPointer() *LogSinkDetailsCompression {
	return &e
}

func (e *LogSinkDetailsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gzip":
		fallthrough
	case "none":
		*e = LogSinkDetailsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsCompression: %v", v)
	}
}

// LogSinkDetailsSchemasRegion - Region of the S3 bucket.
type LogSinkDetailsSchemasRegion string

const (
	LogSinkDetailsSchemasRegionEuWest1    LogSinkDetailsSchemasRegion = "eu-west-1"
	LogSinkDetailsSchemasRegionEuWest2    LogSinkDetailsSchemasRegion = "eu-west-2"
	LogSinkDetailsSchemasRegionEuWest3    LogSinkDetailsSchemasRegion = "eu-west-3"
	LogSinkDetailsSchemasRegionEuCentral1 LogSinkDetailsSchemasRegion = "eu-central-1"
	LogSinkDetailsSchemasRegionEuSouth1   LogSinkDetailsSchemasRegion = "eu-south-1"
	LogSinkDetailsSchemasRegionEuNorth1   LogSinkDetailsSchemasRegion = "eu-north-1"
	LogSinkDetailsSchemasRegionUsWest1    LogSinkDetailsSchemasRegion = "us-west-1"
	LogSinkDetailsSchemasRegionUsWest2    LogSinkDetailsSchemasRegion = "us-west-2"
	LogSinkDetailsSchemasRegionUsEast1    LogSinkDetailsSchemasRegion = "us-east-1"
	LogSinkDetailsSchemasRegionUsEast2    LogSinkDetailsSchemasRegion = "us-east2"
)

func (e LogSinkDetailsSchemasRegion) ToPointer() *LogSinkDetailsSchemasRegion {
	return &e
}

func (e *LogSinkDetailsSchemasRegion) UnmarshalJSON(data []byte) error {
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
		*e = LogSinkDetailsSchemasRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSchemasRegion: %v", v)
	}
}

// LogSinkDetails5 - AWS S3 or compatible API Sink Schema.
type LogSinkDetails5 struct {
	// Authentication object.
	Auth *LogSinkDetailsSchemasDataAuth `json:"auth,omitempty"`
	// Name of the S3 Bucket.
	Bucket string `json:"bucket"`
	// Log file compression method.
	Compression LogSinkDetailsCompression `json:"compression"`
	// Endpoint for the AWS S3 or compatible API bucket.
	Endpoint string `json:"endpoint"`
	// Region of the S3 bucket.
	Region LogSinkDetailsSchemasRegion `json:"region"`
}

func (o *LogSinkDetails5) GetAuth() *LogSinkDetailsSchemasDataAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *LogSinkDetails5) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *LogSinkDetails5) GetCompression() LogSinkDetailsCompression {
	if o == nil {
		return LogSinkDetailsCompression("")
	}
	return o.Compression
}

func (o *LogSinkDetails5) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *LogSinkDetails5) GetRegion() LogSinkDetailsSchemasRegion {
	if o == nil {
		return LogSinkDetailsSchemasRegion("")
	}
	return o.Region
}

// LogSinkDetailsSchemasDataSinkDataStrategy - Bearer token authentication strategy.
type LogSinkDetailsSchemasDataSinkDataStrategy string

const (
	LogSinkDetailsSchemasDataSinkDataStrategyBearer LogSinkDetailsSchemasDataSinkDataStrategy = "bearer"
)

func (e LogSinkDetailsSchemasDataSinkDataStrategy) ToPointer() *LogSinkDetailsSchemasDataSinkDataStrategy {
	return &e
}

func (e *LogSinkDetailsSchemasDataSinkDataStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bearer":
		*e = LogSinkDetailsSchemasDataSinkDataStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSchemasDataSinkDataStrategy: %v", v)
	}
}

// LogSinkDetailsSchemas3 - Authenticate with a bearer token strategy.
type LogSinkDetailsSchemas3 struct {
	// Bearer token authentication strategy.
	Strategy LogSinkDetailsSchemasDataSinkDataStrategy `json:"strategy"`
	// Token for bearer token authentication.
	Token *string `json:"token,omitempty"`
}

func (o *LogSinkDetailsSchemas3) GetStrategy() LogSinkDetailsSchemasDataSinkDataStrategy {
	if o == nil {
		return LogSinkDetailsSchemasDataSinkDataStrategy("")
	}
	return o.Strategy
}

func (o *LogSinkDetailsSchemas3) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// LogSinkDetailsSchemasDataStrategy - Basic HTTP authentication strategy.
type LogSinkDetailsSchemasDataStrategy string

const (
	LogSinkDetailsSchemasDataStrategyBasic LogSinkDetailsSchemasDataStrategy = "basic"
)

func (e LogSinkDetailsSchemasDataStrategy) ToPointer() *LogSinkDetailsSchemasDataStrategy {
	return &e
}

func (e *LogSinkDetailsSchemasDataStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = LogSinkDetailsSchemasDataStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSchemasDataStrategy: %v", v)
	}
}

// LogSinkDetailsSchemasData2 - Authenticate with a basic http strategy.
type LogSinkDetailsSchemasData2 struct {
	// Password for basic http authentication.
	Password string `json:"password"`
	// Basic HTTP authentication strategy.
	Strategy LogSinkDetailsSchemasDataStrategy `json:"strategy"`
	// Username for basic http authentication.
	User *string `json:"user,omitempty"`
}

func (o *LogSinkDetailsSchemasData2) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *LogSinkDetailsSchemasData2) GetStrategy() LogSinkDetailsSchemasDataStrategy {
	if o == nil {
		return LogSinkDetailsSchemasDataStrategy("")
	}
	return o.Strategy
}

func (o *LogSinkDetailsSchemasData2) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// LogSinkDetailsSchemasStrategy - No authentication strategy
type LogSinkDetailsSchemasStrategy string

const (
	LogSinkDetailsSchemasStrategyNone LogSinkDetailsSchemasStrategy = "none"
)

func (e LogSinkDetailsSchemasStrategy) ToPointer() *LogSinkDetailsSchemasStrategy {
	return &e
}

func (e *LogSinkDetailsSchemasStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = LogSinkDetailsSchemasStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSchemasStrategy: %v", v)
	}
}

// LogSinkDetailsSchemasData1 - No authentication strategy
type LogSinkDetailsSchemasData1 struct {
	// No authentication strategy
	Strategy LogSinkDetailsSchemasStrategy `json:"strategy"`
}

func (o *LogSinkDetailsSchemasData1) GetStrategy() LogSinkDetailsSchemasStrategy {
	if o == nil {
		return LogSinkDetailsSchemasStrategy("")
	}
	return o.Strategy
}

type LogSinkDetailsSchemasAuthType string

const (
	LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemasData1 LogSinkDetailsSchemasAuthType = "LogSinkDetails_Schemas_data_1"
	LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemasData2 LogSinkDetailsSchemasAuthType = "LogSinkDetails_Schemas_data_2"
	LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemas3     LogSinkDetailsSchemasAuthType = "LogSinkDetails_Schemas_3"
)

type LogSinkDetailsSchemasAuth struct {
	LogSinkDetailsSchemasData1 *LogSinkDetailsSchemasData1
	LogSinkDetailsSchemasData2 *LogSinkDetailsSchemasData2
	LogSinkDetailsSchemas3     *LogSinkDetailsSchemas3

	Type LogSinkDetailsSchemasAuthType
}

func CreateLogSinkDetailsSchemasAuthLogSinkDetailsSchemasData1(logSinkDetailsSchemasData1 LogSinkDetailsSchemasData1) LogSinkDetailsSchemasAuth {
	typ := LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemasData1

	return LogSinkDetailsSchemasAuth{
		LogSinkDetailsSchemasData1: &logSinkDetailsSchemasData1,
		Type:                       typ,
	}
}

func CreateLogSinkDetailsSchemasAuthLogSinkDetailsSchemasData2(logSinkDetailsSchemasData2 LogSinkDetailsSchemasData2) LogSinkDetailsSchemasAuth {
	typ := LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemasData2

	return LogSinkDetailsSchemasAuth{
		LogSinkDetailsSchemasData2: &logSinkDetailsSchemasData2,
		Type:                       typ,
	}
}

func CreateLogSinkDetailsSchemasAuthLogSinkDetailsSchemas3(logSinkDetailsSchemas3 LogSinkDetailsSchemas3) LogSinkDetailsSchemasAuth {
	typ := LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemas3

	return LogSinkDetailsSchemasAuth{
		LogSinkDetailsSchemas3: &logSinkDetailsSchemas3,
		Type:                   typ,
	}
}

func (u *LogSinkDetailsSchemasAuth) UnmarshalJSON(data []byte) error {

	logSinkDetailsSchemasData1 := LogSinkDetailsSchemasData1{}
	if err := utils.UnmarshalJSON(data, &logSinkDetailsSchemasData1, "", true, true); err == nil {
		u.LogSinkDetailsSchemasData1 = &logSinkDetailsSchemasData1
		u.Type = LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemasData1
		return nil
	}

	logSinkDetailsSchemas3 := LogSinkDetailsSchemas3{}
	if err := utils.UnmarshalJSON(data, &logSinkDetailsSchemas3, "", true, true); err == nil {
		u.LogSinkDetailsSchemas3 = &logSinkDetailsSchemas3
		u.Type = LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemas3
		return nil
	}

	logSinkDetailsSchemasData2 := LogSinkDetailsSchemasData2{}
	if err := utils.UnmarshalJSON(data, &logSinkDetailsSchemasData2, "", true, true); err == nil {
		u.LogSinkDetailsSchemasData2 = &logSinkDetailsSchemasData2
		u.Type = LogSinkDetailsSchemasAuthTypeLogSinkDetailsSchemasData2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkDetailsSchemasAuth) MarshalJSON() ([]byte, error) {
	if u.LogSinkDetailsSchemasData1 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsSchemasData1, "", true)
	}

	if u.LogSinkDetailsSchemasData2 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsSchemasData2, "", true)
	}

	if u.LogSinkDetailsSchemas3 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsSchemas3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// LogSinkDetailsSchemasCodec - Codec to encode logs in
type LogSinkDetailsSchemasCodec string

const (
	LogSinkDetailsSchemasCodecText LogSinkDetailsSchemasCodec = "text"
	LogSinkDetailsSchemasCodecJSON LogSinkDetailsSchemasCodec = "json"
)

func (e LogSinkDetailsSchemasCodec) ToPointer() *LogSinkDetailsSchemasCodec {
	return &e
}

func (e *LogSinkDetailsSchemasCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json":
		*e = LogSinkDetailsSchemasCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSchemasCodec: %v", v)
	}
}

// LogSinkDetailsSchemasEncoding - Encoding options
type LogSinkDetailsSchemasEncoding struct {
	// Codec to encode logs in
	Codec LogSinkDetailsSchemasCodec `json:"codec"`
}

func (o *LogSinkDetailsSchemasEncoding) GetCodec() LogSinkDetailsSchemasCodec {
	if o == nil {
		return LogSinkDetailsSchemasCodec("")
	}
	return o.Codec
}

// LogSinkDetails4 - HTTP Sink Schema.
type LogSinkDetails4 struct {
	Auth LogSinkDetailsSchemasAuth `json:"auth"`
	// Encoding options
	Encoding *LogSinkDetailsSchemasEncoding `json:"encoding,omitempty"`
	// Uri to send logs to.
	URI string `json:"uri"`
}

func (o *LogSinkDetails4) GetAuth() LogSinkDetailsSchemasAuth {
	if o == nil {
		return LogSinkDetailsSchemasAuth{}
	}
	return o.Auth
}

func (o *LogSinkDetails4) GetEncoding() *LogSinkDetailsSchemasEncoding {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *LogSinkDetails4) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// LogSinkDetailsAuthenticationStrategy - The authentication strategy.
type LogSinkDetailsAuthenticationStrategy string

const (
	LogSinkDetailsAuthenticationStrategyToken LogSinkDetailsAuthenticationStrategy = "token"
)

func (e LogSinkDetailsAuthenticationStrategy) ToPointer() *LogSinkDetailsAuthenticationStrategy {
	return &e
}

func (e *LogSinkDetailsAuthenticationStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "token":
		*e = LogSinkDetailsAuthenticationStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsAuthenticationStrategy: %v", v)
	}
}

// LogSinkDetailsSchemas2 - Authenticate with a token.
type LogSinkDetailsSchemas2 struct {
	// The authentication strategy.
	AuthenticationStrategy LogSinkDetailsAuthenticationStrategy `json:"authenticationStrategy"`
	// The HTTP Token for the Papertrail log destination.
	Token string `json:"token"`
	// The uri for the Papertrail log destination.
	URI string `json:"uri"`
}

func (o *LogSinkDetailsSchemas2) GetAuthenticationStrategy() LogSinkDetailsAuthenticationStrategy {
	if o == nil {
		return LogSinkDetailsAuthenticationStrategy("")
	}
	return o.AuthenticationStrategy
}

func (o *LogSinkDetailsSchemas2) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *LogSinkDetailsSchemas2) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// LogSinkDetailsSchemasAuthenticationStrategy - The authentication strategy.
type LogSinkDetailsSchemasAuthenticationStrategy string

const (
	LogSinkDetailsSchemasAuthenticationStrategyPort LogSinkDetailsSchemasAuthenticationStrategy = "port"
)

func (e LogSinkDetailsSchemasAuthenticationStrategy) ToPointer() *LogSinkDetailsSchemasAuthenticationStrategy {
	return &e
}

func (e *LogSinkDetailsSchemasAuthenticationStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "port":
		*e = LogSinkDetailsSchemasAuthenticationStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSchemasAuthenticationStrategy: %v", v)
	}
}

// LogSinkDetailsSchemas1 - Authenticate with a host/port
type LogSinkDetailsSchemas1 struct {
	// The authentication strategy.
	AuthenticationStrategy LogSinkDetailsSchemasAuthenticationStrategy `json:"authenticationStrategy"`
	// The host for the Papertrail log destination.
	Host string `json:"host"`
	// The port for the Papertrail log destination.
	Port float32 `json:"port"`
}

func (o *LogSinkDetailsSchemas1) GetAuthenticationStrategy() LogSinkDetailsSchemasAuthenticationStrategy {
	if o == nil {
		return LogSinkDetailsSchemasAuthenticationStrategy("")
	}
	return o.AuthenticationStrategy
}

func (o *LogSinkDetailsSchemas1) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *LogSinkDetailsSchemas1) GetPort() float32 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

type LogSinkDetails3Type string

const (
	LogSinkDetails3TypeLogSinkDetailsSchemas1 LogSinkDetails3Type = "LogSinkDetails_Schemas_1"
	LogSinkDetails3TypeLogSinkDetailsSchemas2 LogSinkDetails3Type = "LogSinkDetails_Schemas_2"
)

// LogSinkDetails3 - Papertrail Sink Schema.
type LogSinkDetails3 struct {
	LogSinkDetailsSchemas1 *LogSinkDetailsSchemas1
	LogSinkDetailsSchemas2 *LogSinkDetailsSchemas2

	Type LogSinkDetails3Type
}

func CreateLogSinkDetails3LogSinkDetailsSchemas1(logSinkDetailsSchemas1 LogSinkDetailsSchemas1) LogSinkDetails3 {
	typ := LogSinkDetails3TypeLogSinkDetailsSchemas1

	return LogSinkDetails3{
		LogSinkDetailsSchemas1: &logSinkDetailsSchemas1,
		Type:                   typ,
	}
}

func CreateLogSinkDetails3LogSinkDetailsSchemas2(logSinkDetailsSchemas2 LogSinkDetailsSchemas2) LogSinkDetails3 {
	typ := LogSinkDetails3TypeLogSinkDetailsSchemas2

	return LogSinkDetails3{
		LogSinkDetailsSchemas2: &logSinkDetailsSchemas2,
		Type:                   typ,
	}
}

func (u *LogSinkDetails3) UnmarshalJSON(data []byte) error {

	logSinkDetailsSchemas1 := LogSinkDetailsSchemas1{}
	if err := utils.UnmarshalJSON(data, &logSinkDetailsSchemas1, "", true, true); err == nil {
		u.LogSinkDetailsSchemas1 = &logSinkDetailsSchemas1
		u.Type = LogSinkDetails3TypeLogSinkDetailsSchemas1
		return nil
	}

	logSinkDetailsSchemas2 := LogSinkDetailsSchemas2{}
	if err := utils.UnmarshalJSON(data, &logSinkDetailsSchemas2, "", true, true); err == nil {
		u.LogSinkDetailsSchemas2 = &logSinkDetailsSchemas2
		u.Type = LogSinkDetails3TypeLogSinkDetailsSchemas2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkDetails3) MarshalJSON() ([]byte, error) {
	if u.LogSinkDetailsSchemas1 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsSchemas1, "", true)
	}

	if u.LogSinkDetailsSchemas2 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsSchemas2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// LogSinkDetailsRegion - The Datadog region.
type LogSinkDetailsRegion string

const (
	LogSinkDetailsRegionEu  LogSinkDetailsRegion = "eu"
	LogSinkDetailsRegionUs  LogSinkDetailsRegion = "us"
	LogSinkDetailsRegionUs3 LogSinkDetailsRegion = "us3"
	LogSinkDetailsRegionUs5 LogSinkDetailsRegion = "us5"
)

func (e LogSinkDetailsRegion) ToPointer() *LogSinkDetailsRegion {
	return &e
}

func (e *LogSinkDetailsRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "eu":
		fallthrough
	case "us":
		fallthrough
	case "us3":
		fallthrough
	case "us5":
		*e = LogSinkDetailsRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsRegion: %v", v)
	}
}

type LogSinkDetails2 struct {
	// The Datadog API key.
	DefaultAPIKey string `json:"default_api_key"`
	// The Datadog region.
	Region LogSinkDetailsRegion `json:"region"`
}

func (o *LogSinkDetails2) GetDefaultAPIKey() string {
	if o == nil {
		return ""
	}
	return o.DefaultAPIKey
}

func (o *LogSinkDetails2) GetRegion() LogSinkDetailsRegion {
	if o == nil {
		return LogSinkDetailsRegion("")
	}
	return o.Region
}

// LogSinkDetailsStrategy - The authentication strategy.
type LogSinkDetailsStrategy string

const (
	LogSinkDetailsStrategyBasic LogSinkDetailsStrategy = "basic"
)

func (e LogSinkDetailsStrategy) ToPointer() *LogSinkDetailsStrategy {
	return &e
}

func (e *LogSinkDetailsStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = LogSinkDetailsStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsStrategy: %v", v)
	}
}

// LogSinkDetailsAuth - Object containing authentication data for the log sink.
type LogSinkDetailsAuth struct {
	// The password for the log sink.
	Password *string `json:"password,omitempty"`
	// The authentication strategy.
	Strategy *LogSinkDetailsStrategy `json:"strategy,omitempty"`
	// The username for the log sink.
	User *string `json:"user,omitempty"`
}

func (o *LogSinkDetailsAuth) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LogSinkDetailsAuth) GetStrategy() *LogSinkDetailsStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *LogSinkDetailsAuth) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// LogSinkDetailsCodec - Codec to encode logs in
type LogSinkDetailsCodec string

const (
	LogSinkDetailsCodecText LogSinkDetailsCodec = "text"
	LogSinkDetailsCodecJSON LogSinkDetailsCodec = "json"
)

func (e LogSinkDetailsCodec) ToPointer() *LogSinkDetailsCodec {
	return &e
}

func (e *LogSinkDetailsCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json":
		*e = LogSinkDetailsCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsCodec: %v", v)
	}
}

// LogSinkDetailsEncoding - Encoding options
type LogSinkDetailsEncoding struct {
	// Codec to encode logs in
	Codec LogSinkDetailsCodec `json:"codec"`
}

func (o *LogSinkDetailsEncoding) GetCodec() LogSinkDetailsCodec {
	if o == nil {
		return LogSinkDetailsCodec("")
	}
	return o.Codec
}

type LogSinkDetails1 struct {
	// Object containing authentication data for the log sink.
	Auth *LogSinkDetailsAuth `json:"auth,omitempty"`
	// Encoding options
	Encoding *LogSinkDetailsEncoding `json:"encoding,omitempty"`
	// The endpoint of the Loki log sink.
	Endpoint string `json:"endpoint"`
}

func (o *LogSinkDetails1) GetAuth() *LogSinkDetailsAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *LogSinkDetails1) GetEncoding() *LogSinkDetailsEncoding {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *LogSinkDetails1) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

type LogSinkDetailsSinkDataType string

const (
	LogSinkDetailsSinkDataTypeLogSinkDetails1 LogSinkDetailsSinkDataType = "LogSinkDetails_1"
	LogSinkDetailsSinkDataTypeLogSinkDetails2 LogSinkDetailsSinkDataType = "LogSinkDetails_2"
	LogSinkDetailsSinkDataTypeLogSinkDetails3 LogSinkDetailsSinkDataType = "LogSinkDetails_3"
	LogSinkDetailsSinkDataTypeLogSinkDetails4 LogSinkDetailsSinkDataType = "LogSinkDetails_4"
	LogSinkDetailsSinkDataTypeLogSinkDetails5 LogSinkDetailsSinkDataType = "LogSinkDetails_5"
	LogSinkDetailsSinkDataTypeLogSinkDetails6 LogSinkDetailsSinkDataType = "LogSinkDetails_6"
	LogSinkDetailsSinkDataTypeLogSinkDetails7 LogSinkDetailsSinkDataType = "LogSinkDetails_7"
	LogSinkDetailsSinkDataTypeLogSinkDetails8 LogSinkDetailsSinkDataType = "LogSinkDetails_8"
)

// LogSinkDetailsSinkData - Data about the log sink.
type LogSinkDetailsSinkData struct {
	LogSinkDetails1 *LogSinkDetails1
	LogSinkDetails2 *LogSinkDetails2
	LogSinkDetails3 *LogSinkDetails3
	LogSinkDetails4 *LogSinkDetails4
	LogSinkDetails5 *LogSinkDetails5
	LogSinkDetails6 *LogSinkDetails6
	LogSinkDetails7 *LogSinkDetails7
	LogSinkDetails8 *LogSinkDetails8

	Type LogSinkDetailsSinkDataType
}

func CreateLogSinkDetailsSinkDataLogSinkDetails1(logSinkDetails1 LogSinkDetails1) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails1

	return LogSinkDetailsSinkData{
		LogSinkDetails1: &logSinkDetails1,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails2(logSinkDetails2 LogSinkDetails2) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails2

	return LogSinkDetailsSinkData{
		LogSinkDetails2: &logSinkDetails2,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails3(logSinkDetails3 LogSinkDetails3) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails3

	return LogSinkDetailsSinkData{
		LogSinkDetails3: &logSinkDetails3,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails4(logSinkDetails4 LogSinkDetails4) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails4

	return LogSinkDetailsSinkData{
		LogSinkDetails4: &logSinkDetails4,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails5(logSinkDetails5 LogSinkDetails5) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails5

	return LogSinkDetailsSinkData{
		LogSinkDetails5: &logSinkDetails5,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails6(logSinkDetails6 LogSinkDetails6) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails6

	return LogSinkDetailsSinkData{
		LogSinkDetails6: &logSinkDetails6,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails7(logSinkDetails7 LogSinkDetails7) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails7

	return LogSinkDetailsSinkData{
		LogSinkDetails7: &logSinkDetails7,
		Type:            typ,
	}
}

func CreateLogSinkDetailsSinkDataLogSinkDetails8(logSinkDetails8 LogSinkDetails8) LogSinkDetailsSinkData {
	typ := LogSinkDetailsSinkDataTypeLogSinkDetails8

	return LogSinkDetailsSinkData{
		LogSinkDetails8: &logSinkDetails8,
		Type:            typ,
	}
}

func (u *LogSinkDetailsSinkData) UnmarshalJSON(data []byte) error {

	logSinkDetails6 := LogSinkDetails6{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails6, "", true, true); err == nil {
		u.LogSinkDetails6 = &logSinkDetails6
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails6
		return nil
	}

	logSinkDetails7 := LogSinkDetails7{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails7, "", true, true); err == nil {
		u.LogSinkDetails7 = &logSinkDetails7
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails7
		return nil
	}

	logSinkDetails2 := LogSinkDetails2{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails2, "", true, true); err == nil {
		u.LogSinkDetails2 = &logSinkDetails2
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails2
		return nil
	}

	logSinkDetails8 := LogSinkDetails8{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails8, "", true, true); err == nil {
		u.LogSinkDetails8 = &logSinkDetails8
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails8
		return nil
	}

	logSinkDetails1 := LogSinkDetails1{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails1, "", true, true); err == nil {
		u.LogSinkDetails1 = &logSinkDetails1
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails1
		return nil
	}

	logSinkDetails4 := LogSinkDetails4{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails4, "", true, true); err == nil {
		u.LogSinkDetails4 = &logSinkDetails4
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails4
		return nil
	}

	logSinkDetails5 := LogSinkDetails5{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails5, "", true, true); err == nil {
		u.LogSinkDetails5 = &logSinkDetails5
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails5
		return nil
	}

	logSinkDetails3 := LogSinkDetails3{}
	if err := utils.UnmarshalJSON(data, &logSinkDetails3, "", true, true); err == nil {
		u.LogSinkDetails3 = &logSinkDetails3
		u.Type = LogSinkDetailsSinkDataTypeLogSinkDetails3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkDetailsSinkData) MarshalJSON() ([]byte, error) {
	if u.LogSinkDetails1 != nil {
		return utils.MarshalJSON(u.LogSinkDetails1, "", true)
	}

	if u.LogSinkDetails2 != nil {
		return utils.MarshalJSON(u.LogSinkDetails2, "", true)
	}

	if u.LogSinkDetails3 != nil {
		return utils.MarshalJSON(u.LogSinkDetails3, "", true)
	}

	if u.LogSinkDetails4 != nil {
		return utils.MarshalJSON(u.LogSinkDetails4, "", true)
	}

	if u.LogSinkDetails5 != nil {
		return utils.MarshalJSON(u.LogSinkDetails5, "", true)
	}

	if u.LogSinkDetails6 != nil {
		return utils.MarshalJSON(u.LogSinkDetails6, "", true)
	}

	if u.LogSinkDetails7 != nil {
		return utils.MarshalJSON(u.LogSinkDetails7, "", true)
	}

	if u.LogSinkDetails8 != nil {
		return utils.MarshalJSON(u.LogSinkDetails8, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// LogSinkDetailsSinkType - The type of the log sink.
type LogSinkDetailsSinkType string

const (
	LogSinkDetailsSinkTypeLoki        LogSinkDetailsSinkType = "loki"
	LogSinkDetailsSinkTypeDatadogLogs LogSinkDetailsSinkType = "datadog_logs"
	LogSinkDetailsSinkTypePapertrail  LogSinkDetailsSinkType = "papertrail"
	LogSinkDetailsSinkTypeHTTP        LogSinkDetailsSinkType = "http"
	LogSinkDetailsSinkTypeAwsS3       LogSinkDetailsSinkType = "aws_s3"
	LogSinkDetailsSinkTypeLogdna      LogSinkDetailsSinkType = "logdna"
	LogSinkDetailsSinkTypeCoralogix   LogSinkDetailsSinkType = "coralogix"
	LogSinkDetailsSinkTypeLogtail     LogSinkDetailsSinkType = "logtail"
	LogSinkDetailsSinkTypeHoneycomb   LogSinkDetailsSinkType = "honeycomb"
	LogSinkDetailsSinkTypeLogzio      LogSinkDetailsSinkType = "logzio"
)

func (e LogSinkDetailsSinkType) ToPointer() *LogSinkDetailsSinkType {
	return &e
}

func (e *LogSinkDetailsSinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "loki":
		fallthrough
	case "datadog_logs":
		fallthrough
	case "papertrail":
		fallthrough
	case "http":
		fallthrough
	case "aws_s3":
		fallthrough
	case "logdna":
		fallthrough
	case "coralogix":
		fallthrough
	case "logtail":
		fallthrough
	case "honeycomb":
		fallthrough
	case "logzio":
		*e = LogSinkDetailsSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsSinkType: %v", v)
	}
}

// LogSinkDetailsData - Result data.
type LogSinkDetailsData struct {
	// Timestamp of when the log sink was created.
	CreatedAt time.Time `json:"createdAt"`
	// Description of the log sink.
	Description *string `json:"description,omitempty"`
	// Identifier for the Log Sink
	ID string `json:"id"`
	// Name of the log sink.
	Name string `json:"name"`
	// If `restricted` is `true`, only logs from these projects will be sent to the log sink.
	Projects []string `json:"projects"`
	// If `true`, only logs from the projects in `projects` will be sent to the log sink.
	Restricted bool `json:"restricted"`
	// Data about the log sink.
	SinkData *LogSinkDetailsSinkData `json:"sinkData,omitempty"`
	// The type of the log sink.
	SinkType LogSinkDetailsSinkType `json:"sinkType"`
	// Timestamp of when the log sink was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `default:"false" json:"useCustomLabels"`
}

func (l LogSinkDetailsData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LogSinkDetailsData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LogSinkDetailsData) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *LogSinkDetailsData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LogSinkDetailsData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LogSinkDetailsData) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *LogSinkDetailsData) GetProjects() []string {
	if o == nil {
		return []string{}
	}
	return o.Projects
}

func (o *LogSinkDetailsData) GetRestricted() bool {
	if o == nil {
		return false
	}
	return o.Restricted
}

func (o *LogSinkDetailsData) GetSinkData() *LogSinkDetailsSinkData {
	if o == nil {
		return nil
	}
	return o.SinkData
}

func (o *LogSinkDetailsData) GetSinkType() LogSinkDetailsSinkType {
	if o == nil {
		return LogSinkDetailsSinkType("")
	}
	return o.SinkType
}

func (o *LogSinkDetailsData) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *LogSinkDetailsData) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}

// LogSinkDetails - Response object.
type LogSinkDetails struct {
	// Result data.
	Data LogSinkDetailsData `json:"data"`
}

func (o *LogSinkDetails) GetData() LogSinkDetailsData {
	if o == nil {
		return LogSinkDetailsData{}
	}
	return o.Data
}
