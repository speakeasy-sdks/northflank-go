// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
	"time"
)

// LogSinkDetailsDataSinkData8 - Honeycomb Sink Schema.
type LogSinkDetailsDataSinkData8 struct {
	// Honeycomb API Key
	APIKey string `json:"api_key"`
	// Name of the dataset
	Dataset string `json:"dataset"`
}

func (o *LogSinkDetailsDataSinkData8) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *LogSinkDetailsDataSinkData8) GetDataset() string {
	if o == nil {
		return ""
	}
	return o.Dataset
}

// LogSinkDetailsDataSinkData7 - LogDNA Sink Schema.
type LogSinkDetailsDataSinkData7 struct {
	// Ingestion Key
	APIKey string `json:"api_key"`
}

func (o *LogSinkDetailsDataSinkData7) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

// LogSinkDetailsDataSinkData6 - Logtail Sink Schema.
type LogSinkDetailsDataSinkData6 struct {
	// Logtail Source Token
	Token string `json:"token"`
}

func (o *LogSinkDetailsDataSinkData6) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// LogSinkDetailsDataSinkData5Auth - Authentication object.
type LogSinkDetailsDataSinkData5Auth struct {
	// Access key id for the bucket.
	AccessKeyID string `json:"accessKeyId"`
	// Secret access key for the bucket.
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *LogSinkDetailsDataSinkData5Auth) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *LogSinkDetailsDataSinkData5Auth) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// LogSinkDetailsDataSinkData5Compression - Log file compression method.
type LogSinkDetailsDataSinkData5Compression string

const (
	LogSinkDetailsDataSinkData5CompressionGzip LogSinkDetailsDataSinkData5Compression = "gzip"
	LogSinkDetailsDataSinkData5CompressionNone LogSinkDetailsDataSinkData5Compression = "none"
)

func (e LogSinkDetailsDataSinkData5Compression) ToPointer() *LogSinkDetailsDataSinkData5Compression {
	return &e
}

func (e *LogSinkDetailsDataSinkData5Compression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gzip":
		fallthrough
	case "none":
		*e = LogSinkDetailsDataSinkData5Compression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData5Compression: %v", v)
	}
}

// LogSinkDetailsDataSinkData5Region - Region of the S3 bucket.
type LogSinkDetailsDataSinkData5Region string

const (
	LogSinkDetailsDataSinkData5RegionEuWest1    LogSinkDetailsDataSinkData5Region = "eu-west-1"
	LogSinkDetailsDataSinkData5RegionEuWest2    LogSinkDetailsDataSinkData5Region = "eu-west-2"
	LogSinkDetailsDataSinkData5RegionEuWest3    LogSinkDetailsDataSinkData5Region = "eu-west-3"
	LogSinkDetailsDataSinkData5RegionEuCentral1 LogSinkDetailsDataSinkData5Region = "eu-central-1"
	LogSinkDetailsDataSinkData5RegionEuSouth1   LogSinkDetailsDataSinkData5Region = "eu-south-1"
	LogSinkDetailsDataSinkData5RegionEuNorth1   LogSinkDetailsDataSinkData5Region = "eu-north-1"
	LogSinkDetailsDataSinkData5RegionUsWest1    LogSinkDetailsDataSinkData5Region = "us-west-1"
	LogSinkDetailsDataSinkData5RegionUsWest2    LogSinkDetailsDataSinkData5Region = "us-west-2"
	LogSinkDetailsDataSinkData5RegionUsEast1    LogSinkDetailsDataSinkData5Region = "us-east-1"
	LogSinkDetailsDataSinkData5RegionUsEast2    LogSinkDetailsDataSinkData5Region = "us-east2"
)

func (e LogSinkDetailsDataSinkData5Region) ToPointer() *LogSinkDetailsDataSinkData5Region {
	return &e
}

func (e *LogSinkDetailsDataSinkData5Region) UnmarshalJSON(data []byte) error {
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
		*e = LogSinkDetailsDataSinkData5Region(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData5Region: %v", v)
	}
}

// LogSinkDetailsDataSinkData5 - AWS S3 or compatible API Sink Schema.
type LogSinkDetailsDataSinkData5 struct {
	// Authentication object.
	Auth *LogSinkDetailsDataSinkData5Auth `json:"auth,omitempty"`
	// Name of the S3 Bucket.
	Bucket string `json:"bucket"`
	// Log file compression method.
	Compression LogSinkDetailsDataSinkData5Compression `json:"compression"`
	// Endpoint for the AWS S3 or compatible API bucket.
	Endpoint string `json:"endpoint"`
	// Region of the S3 bucket.
	Region LogSinkDetailsDataSinkData5Region `json:"region"`
}

func (o *LogSinkDetailsDataSinkData5) GetAuth() *LogSinkDetailsDataSinkData5Auth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *LogSinkDetailsDataSinkData5) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *LogSinkDetailsDataSinkData5) GetCompression() LogSinkDetailsDataSinkData5Compression {
	if o == nil {
		return LogSinkDetailsDataSinkData5Compression("")
	}
	return o.Compression
}

func (o *LogSinkDetailsDataSinkData5) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *LogSinkDetailsDataSinkData5) GetRegion() LogSinkDetailsDataSinkData5Region {
	if o == nil {
		return LogSinkDetailsDataSinkData5Region("")
	}
	return o.Region
}

// LogSinkDetailsDataSinkData4Auth3Strategy - Bearer token authentication strategy.
type LogSinkDetailsDataSinkData4Auth3Strategy string

const (
	LogSinkDetailsDataSinkData4Auth3StrategyBearer LogSinkDetailsDataSinkData4Auth3Strategy = "bearer"
)

func (e LogSinkDetailsDataSinkData4Auth3Strategy) ToPointer() *LogSinkDetailsDataSinkData4Auth3Strategy {
	return &e
}

func (e *LogSinkDetailsDataSinkData4Auth3Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bearer":
		*e = LogSinkDetailsDataSinkData4Auth3Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData4Auth3Strategy: %v", v)
	}
}

// LogSinkDetailsDataSinkData4Auth3 - Authenticate with a bearer token strategy.
type LogSinkDetailsDataSinkData4Auth3 struct {
	// Bearer token authentication strategy.
	Strategy LogSinkDetailsDataSinkData4Auth3Strategy `json:"strategy"`
	// Token for bearer token authentication.
	Token *string `json:"token,omitempty"`
}

func (o *LogSinkDetailsDataSinkData4Auth3) GetStrategy() LogSinkDetailsDataSinkData4Auth3Strategy {
	if o == nil {
		return LogSinkDetailsDataSinkData4Auth3Strategy("")
	}
	return o.Strategy
}

func (o *LogSinkDetailsDataSinkData4Auth3) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// LogSinkDetailsDataSinkData4Auth2Strategy - Basic HTTP authentication strategy.
type LogSinkDetailsDataSinkData4Auth2Strategy string

const (
	LogSinkDetailsDataSinkData4Auth2StrategyBasic LogSinkDetailsDataSinkData4Auth2Strategy = "basic"
)

func (e LogSinkDetailsDataSinkData4Auth2Strategy) ToPointer() *LogSinkDetailsDataSinkData4Auth2Strategy {
	return &e
}

func (e *LogSinkDetailsDataSinkData4Auth2Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = LogSinkDetailsDataSinkData4Auth2Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData4Auth2Strategy: %v", v)
	}
}

// LogSinkDetailsDataSinkData4Auth2 - Authenticate with a basic http strategy.
type LogSinkDetailsDataSinkData4Auth2 struct {
	// Password for basic http authentication.
	Password string `json:"password"`
	// Basic HTTP authentication strategy.
	Strategy LogSinkDetailsDataSinkData4Auth2Strategy `json:"strategy"`
	// Username for basic http authentication.
	User *string `json:"user,omitempty"`
}

func (o *LogSinkDetailsDataSinkData4Auth2) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *LogSinkDetailsDataSinkData4Auth2) GetStrategy() LogSinkDetailsDataSinkData4Auth2Strategy {
	if o == nil {
		return LogSinkDetailsDataSinkData4Auth2Strategy("")
	}
	return o.Strategy
}

func (o *LogSinkDetailsDataSinkData4Auth2) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// LogSinkDetailsDataSinkData4Auth1Strategy - No authentication strategy
type LogSinkDetailsDataSinkData4Auth1Strategy string

const (
	LogSinkDetailsDataSinkData4Auth1StrategyNone LogSinkDetailsDataSinkData4Auth1Strategy = "none"
)

func (e LogSinkDetailsDataSinkData4Auth1Strategy) ToPointer() *LogSinkDetailsDataSinkData4Auth1Strategy {
	return &e
}

func (e *LogSinkDetailsDataSinkData4Auth1Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = LogSinkDetailsDataSinkData4Auth1Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData4Auth1Strategy: %v", v)
	}
}

// LogSinkDetailsDataSinkData4Auth1 - No authentication strategy
type LogSinkDetailsDataSinkData4Auth1 struct {
	// No authentication strategy
	Strategy LogSinkDetailsDataSinkData4Auth1Strategy `json:"strategy"`
}

func (o *LogSinkDetailsDataSinkData4Auth1) GetStrategy() LogSinkDetailsDataSinkData4Auth1Strategy {
	if o == nil {
		return LogSinkDetailsDataSinkData4Auth1Strategy("")
	}
	return o.Strategy
}

type LogSinkDetailsDataSinkData4AuthType string

const (
	LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth1 LogSinkDetailsDataSinkData4AuthType = "LogSinkDetails_data_sinkData_4_auth_1"
	LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth2 LogSinkDetailsDataSinkData4AuthType = "LogSinkDetails_data_sinkData_4_auth_2"
	LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth3 LogSinkDetailsDataSinkData4AuthType = "LogSinkDetails_data_sinkData_4_auth_3"
)

type LogSinkDetailsDataSinkData4Auth struct {
	LogSinkDetailsDataSinkData4Auth1 *LogSinkDetailsDataSinkData4Auth1
	LogSinkDetailsDataSinkData4Auth2 *LogSinkDetailsDataSinkData4Auth2
	LogSinkDetailsDataSinkData4Auth3 *LogSinkDetailsDataSinkData4Auth3

	Type LogSinkDetailsDataSinkData4AuthType
}

func CreateLogSinkDetailsDataSinkData4AuthLogSinkDetailsDataSinkData4Auth1(logSinkDetailsDataSinkData4Auth1 LogSinkDetailsDataSinkData4Auth1) LogSinkDetailsDataSinkData4Auth {
	typ := LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth1

	return LogSinkDetailsDataSinkData4Auth{
		LogSinkDetailsDataSinkData4Auth1: &logSinkDetailsDataSinkData4Auth1,
		Type:                             typ,
	}
}

func CreateLogSinkDetailsDataSinkData4AuthLogSinkDetailsDataSinkData4Auth2(logSinkDetailsDataSinkData4Auth2 LogSinkDetailsDataSinkData4Auth2) LogSinkDetailsDataSinkData4Auth {
	typ := LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth2

	return LogSinkDetailsDataSinkData4Auth{
		LogSinkDetailsDataSinkData4Auth2: &logSinkDetailsDataSinkData4Auth2,
		Type:                             typ,
	}
}

func CreateLogSinkDetailsDataSinkData4AuthLogSinkDetailsDataSinkData4Auth3(logSinkDetailsDataSinkData4Auth3 LogSinkDetailsDataSinkData4Auth3) LogSinkDetailsDataSinkData4Auth {
	typ := LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth3

	return LogSinkDetailsDataSinkData4Auth{
		LogSinkDetailsDataSinkData4Auth3: &logSinkDetailsDataSinkData4Auth3,
		Type:                             typ,
	}
}

func (u *LogSinkDetailsDataSinkData4Auth) UnmarshalJSON(data []byte) error {

	logSinkDetailsDataSinkData4Auth1 := new(LogSinkDetailsDataSinkData4Auth1)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData4Auth1, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData4Auth1 = logSinkDetailsDataSinkData4Auth1
		u.Type = LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth1
		return nil
	}

	logSinkDetailsDataSinkData4Auth3 := new(LogSinkDetailsDataSinkData4Auth3)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData4Auth3, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData4Auth3 = logSinkDetailsDataSinkData4Auth3
		u.Type = LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth3
		return nil
	}

	logSinkDetailsDataSinkData4Auth2 := new(LogSinkDetailsDataSinkData4Auth2)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData4Auth2, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData4Auth2 = logSinkDetailsDataSinkData4Auth2
		u.Type = LogSinkDetailsDataSinkData4AuthTypeLogSinkDetailsDataSinkData4Auth2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkDetailsDataSinkData4Auth) MarshalJSON() ([]byte, error) {
	if u.LogSinkDetailsDataSinkData4Auth1 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData4Auth1, "", true)
	}

	if u.LogSinkDetailsDataSinkData4Auth2 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData4Auth2, "", true)
	}

	if u.LogSinkDetailsDataSinkData4Auth3 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData4Auth3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// LogSinkDetailsDataSinkData4EncodingCodec - Codec to encode logs in
type LogSinkDetailsDataSinkData4EncodingCodec string

const (
	LogSinkDetailsDataSinkData4EncodingCodecText LogSinkDetailsDataSinkData4EncodingCodec = "text"
	LogSinkDetailsDataSinkData4EncodingCodecJSON LogSinkDetailsDataSinkData4EncodingCodec = "json"
)

func (e LogSinkDetailsDataSinkData4EncodingCodec) ToPointer() *LogSinkDetailsDataSinkData4EncodingCodec {
	return &e
}

func (e *LogSinkDetailsDataSinkData4EncodingCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json":
		*e = LogSinkDetailsDataSinkData4EncodingCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData4EncodingCodec: %v", v)
	}
}

// LogSinkDetailsDataSinkData4Encoding - Encoding options
type LogSinkDetailsDataSinkData4Encoding struct {
	// Codec to encode logs in
	Codec LogSinkDetailsDataSinkData4EncodingCodec `json:"codec"`
}

func (o *LogSinkDetailsDataSinkData4Encoding) GetCodec() LogSinkDetailsDataSinkData4EncodingCodec {
	if o == nil {
		return LogSinkDetailsDataSinkData4EncodingCodec("")
	}
	return o.Codec
}

// LogSinkDetailsDataSinkData4 - HTTP Sink Schema.
type LogSinkDetailsDataSinkData4 struct {
	Auth LogSinkDetailsDataSinkData4Auth `json:"auth"`
	// Encoding options
	Encoding *LogSinkDetailsDataSinkData4Encoding `json:"encoding,omitempty"`
	// Uri to send logs to.
	URI string `json:"uri"`
}

func (o *LogSinkDetailsDataSinkData4) GetAuth() LogSinkDetailsDataSinkData4Auth {
	if o == nil {
		return LogSinkDetailsDataSinkData4Auth{}
	}
	return o.Auth
}

func (o *LogSinkDetailsDataSinkData4) GetEncoding() *LogSinkDetailsDataSinkData4Encoding {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *LogSinkDetailsDataSinkData4) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// LogSinkDetailsDataSinkData32AuthenticationStrategy - The authentication strategy.
type LogSinkDetailsDataSinkData32AuthenticationStrategy string

const (
	LogSinkDetailsDataSinkData32AuthenticationStrategyToken LogSinkDetailsDataSinkData32AuthenticationStrategy = "token"
)

func (e LogSinkDetailsDataSinkData32AuthenticationStrategy) ToPointer() *LogSinkDetailsDataSinkData32AuthenticationStrategy {
	return &e
}

func (e *LogSinkDetailsDataSinkData32AuthenticationStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "token":
		*e = LogSinkDetailsDataSinkData32AuthenticationStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData32AuthenticationStrategy: %v", v)
	}
}

// LogSinkDetailsDataSinkData32 - Authenticate with a token.
type LogSinkDetailsDataSinkData32 struct {
	// The authentication strategy.
	AuthenticationStrategy LogSinkDetailsDataSinkData32AuthenticationStrategy `json:"authenticationStrategy"`
	// The HTTP Token for the Papertrail log destination.
	Token string `json:"token"`
	// The uri for the Papertrail log destination.
	URI string `json:"uri"`
}

func (o *LogSinkDetailsDataSinkData32) GetAuthenticationStrategy() LogSinkDetailsDataSinkData32AuthenticationStrategy {
	if o == nil {
		return LogSinkDetailsDataSinkData32AuthenticationStrategy("")
	}
	return o.AuthenticationStrategy
}

func (o *LogSinkDetailsDataSinkData32) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *LogSinkDetailsDataSinkData32) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// LogSinkDetailsDataSinkData31AuthenticationStrategy - The authentication strategy.
type LogSinkDetailsDataSinkData31AuthenticationStrategy string

const (
	LogSinkDetailsDataSinkData31AuthenticationStrategyPort LogSinkDetailsDataSinkData31AuthenticationStrategy = "port"
)

func (e LogSinkDetailsDataSinkData31AuthenticationStrategy) ToPointer() *LogSinkDetailsDataSinkData31AuthenticationStrategy {
	return &e
}

func (e *LogSinkDetailsDataSinkData31AuthenticationStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "port":
		*e = LogSinkDetailsDataSinkData31AuthenticationStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData31AuthenticationStrategy: %v", v)
	}
}

// LogSinkDetailsDataSinkData31 - Authenticate with a host/port
type LogSinkDetailsDataSinkData31 struct {
	// The authentication strategy.
	AuthenticationStrategy LogSinkDetailsDataSinkData31AuthenticationStrategy `json:"authenticationStrategy"`
	// The host for the Papertrail log destination.
	Host string `json:"host"`
	// The port for the Papertrail log destination.
	Port float32 `json:"port"`
}

func (o *LogSinkDetailsDataSinkData31) GetAuthenticationStrategy() LogSinkDetailsDataSinkData31AuthenticationStrategy {
	if o == nil {
		return LogSinkDetailsDataSinkData31AuthenticationStrategy("")
	}
	return o.AuthenticationStrategy
}

func (o *LogSinkDetailsDataSinkData31) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *LogSinkDetailsDataSinkData31) GetPort() float32 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

type LogSinkDetailsDataSinkData3Type string

const (
	LogSinkDetailsDataSinkData3TypeLogSinkDetailsDataSinkData31 LogSinkDetailsDataSinkData3Type = "LogSinkDetails_data_sinkData_3_1"
	LogSinkDetailsDataSinkData3TypeLogSinkDetailsDataSinkData32 LogSinkDetailsDataSinkData3Type = "LogSinkDetails_data_sinkData_3_2"
)

type LogSinkDetailsDataSinkData3 struct {
	LogSinkDetailsDataSinkData31 *LogSinkDetailsDataSinkData31
	LogSinkDetailsDataSinkData32 *LogSinkDetailsDataSinkData32

	Type LogSinkDetailsDataSinkData3Type
}

func CreateLogSinkDetailsDataSinkData3LogSinkDetailsDataSinkData31(logSinkDetailsDataSinkData31 LogSinkDetailsDataSinkData31) LogSinkDetailsDataSinkData3 {
	typ := LogSinkDetailsDataSinkData3TypeLogSinkDetailsDataSinkData31

	return LogSinkDetailsDataSinkData3{
		LogSinkDetailsDataSinkData31: &logSinkDetailsDataSinkData31,
		Type:                         typ,
	}
}

func CreateLogSinkDetailsDataSinkData3LogSinkDetailsDataSinkData32(logSinkDetailsDataSinkData32 LogSinkDetailsDataSinkData32) LogSinkDetailsDataSinkData3 {
	typ := LogSinkDetailsDataSinkData3TypeLogSinkDetailsDataSinkData32

	return LogSinkDetailsDataSinkData3{
		LogSinkDetailsDataSinkData32: &logSinkDetailsDataSinkData32,
		Type:                         typ,
	}
}

func (u *LogSinkDetailsDataSinkData3) UnmarshalJSON(data []byte) error {

	logSinkDetailsDataSinkData31 := new(LogSinkDetailsDataSinkData31)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData31, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData31 = logSinkDetailsDataSinkData31
		u.Type = LogSinkDetailsDataSinkData3TypeLogSinkDetailsDataSinkData31
		return nil
	}

	logSinkDetailsDataSinkData32 := new(LogSinkDetailsDataSinkData32)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData32, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData32 = logSinkDetailsDataSinkData32
		u.Type = LogSinkDetailsDataSinkData3TypeLogSinkDetailsDataSinkData32
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkDetailsDataSinkData3) MarshalJSON() ([]byte, error) {
	if u.LogSinkDetailsDataSinkData31 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData31, "", true)
	}

	if u.LogSinkDetailsDataSinkData32 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData32, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// LogSinkDetailsDataSinkData2Region - The Datadog region.
type LogSinkDetailsDataSinkData2Region string

const (
	LogSinkDetailsDataSinkData2RegionEu  LogSinkDetailsDataSinkData2Region = "eu"
	LogSinkDetailsDataSinkData2RegionUs  LogSinkDetailsDataSinkData2Region = "us"
	LogSinkDetailsDataSinkData2RegionUs3 LogSinkDetailsDataSinkData2Region = "us3"
	LogSinkDetailsDataSinkData2RegionUs5 LogSinkDetailsDataSinkData2Region = "us5"
)

func (e LogSinkDetailsDataSinkData2Region) ToPointer() *LogSinkDetailsDataSinkData2Region {
	return &e
}

func (e *LogSinkDetailsDataSinkData2Region) UnmarshalJSON(data []byte) error {
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
		*e = LogSinkDetailsDataSinkData2Region(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData2Region: %v", v)
	}
}

// LogSinkDetailsDataSinkData2 - Data about the log sink.
type LogSinkDetailsDataSinkData2 struct {
	// The Datadog API key.
	DefaultAPIKey string `json:"default_api_key"`
	// The Datadog region.
	Region LogSinkDetailsDataSinkData2Region `json:"region"`
}

func (o *LogSinkDetailsDataSinkData2) GetDefaultAPIKey() string {
	if o == nil {
		return ""
	}
	return o.DefaultAPIKey
}

func (o *LogSinkDetailsDataSinkData2) GetRegion() LogSinkDetailsDataSinkData2Region {
	if o == nil {
		return LogSinkDetailsDataSinkData2Region("")
	}
	return o.Region
}

// LogSinkDetailsDataSinkData1AuthStrategy - The authentication strategy.
type LogSinkDetailsDataSinkData1AuthStrategy string

const (
	LogSinkDetailsDataSinkData1AuthStrategyBasic LogSinkDetailsDataSinkData1AuthStrategy = "basic"
)

func (e LogSinkDetailsDataSinkData1AuthStrategy) ToPointer() *LogSinkDetailsDataSinkData1AuthStrategy {
	return &e
}

func (e *LogSinkDetailsDataSinkData1AuthStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = LogSinkDetailsDataSinkData1AuthStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData1AuthStrategy: %v", v)
	}
}

// LogSinkDetailsDataSinkData1Auth - Object containing authentication data for the log sink.
type LogSinkDetailsDataSinkData1Auth struct {
	// The password for the log sink.
	Password *string `json:"password,omitempty"`
	// The authentication strategy.
	Strategy *LogSinkDetailsDataSinkData1AuthStrategy `json:"strategy,omitempty"`
	// The username for the log sink.
	User *string `json:"user,omitempty"`
}

func (o *LogSinkDetailsDataSinkData1Auth) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LogSinkDetailsDataSinkData1Auth) GetStrategy() *LogSinkDetailsDataSinkData1AuthStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *LogSinkDetailsDataSinkData1Auth) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// LogSinkDetailsDataSinkData1EncodingCodec - Codec to encode logs in
type LogSinkDetailsDataSinkData1EncodingCodec string

const (
	LogSinkDetailsDataSinkData1EncodingCodecText LogSinkDetailsDataSinkData1EncodingCodec = "text"
	LogSinkDetailsDataSinkData1EncodingCodecJSON LogSinkDetailsDataSinkData1EncodingCodec = "json"
)

func (e LogSinkDetailsDataSinkData1EncodingCodec) ToPointer() *LogSinkDetailsDataSinkData1EncodingCodec {
	return &e
}

func (e *LogSinkDetailsDataSinkData1EncodingCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json":
		*e = LogSinkDetailsDataSinkData1EncodingCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkData1EncodingCodec: %v", v)
	}
}

// LogSinkDetailsDataSinkData1Encoding - Encoding options
type LogSinkDetailsDataSinkData1Encoding struct {
	// Codec to encode logs in
	Codec LogSinkDetailsDataSinkData1EncodingCodec `json:"codec"`
}

func (o *LogSinkDetailsDataSinkData1Encoding) GetCodec() LogSinkDetailsDataSinkData1EncodingCodec {
	if o == nil {
		return LogSinkDetailsDataSinkData1EncodingCodec("")
	}
	return o.Codec
}

// LogSinkDetailsDataSinkData1 - Data about the log sink.
type LogSinkDetailsDataSinkData1 struct {
	// Object containing authentication data for the log sink.
	Auth *LogSinkDetailsDataSinkData1Auth `json:"auth,omitempty"`
	// Encoding options
	Encoding *LogSinkDetailsDataSinkData1Encoding `json:"encoding,omitempty"`
	// The endpoint of the Loki log sink.
	Endpoint string `json:"endpoint"`
}

func (o *LogSinkDetailsDataSinkData1) GetAuth() *LogSinkDetailsDataSinkData1Auth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *LogSinkDetailsDataSinkData1) GetEncoding() *LogSinkDetailsDataSinkData1Encoding {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *LogSinkDetailsDataSinkData1) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

type LogSinkDetailsDataSinkDataType string

const (
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData1 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_1"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData2 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_2"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData3 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_3"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData4 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_4"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData5 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_5"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData6 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_6"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData7 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_7"
	LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData8 LogSinkDetailsDataSinkDataType = "LogSinkDetails_data_sinkData_8"
)

type LogSinkDetailsDataSinkData struct {
	LogSinkDetailsDataSinkData1 *LogSinkDetailsDataSinkData1
	LogSinkDetailsDataSinkData2 *LogSinkDetailsDataSinkData2
	LogSinkDetailsDataSinkData3 *LogSinkDetailsDataSinkData3
	LogSinkDetailsDataSinkData4 *LogSinkDetailsDataSinkData4
	LogSinkDetailsDataSinkData5 *LogSinkDetailsDataSinkData5
	LogSinkDetailsDataSinkData6 *LogSinkDetailsDataSinkData6
	LogSinkDetailsDataSinkData7 *LogSinkDetailsDataSinkData7
	LogSinkDetailsDataSinkData8 *LogSinkDetailsDataSinkData8

	Type LogSinkDetailsDataSinkDataType
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData1(logSinkDetailsDataSinkData1 LogSinkDetailsDataSinkData1) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData1

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData1: &logSinkDetailsDataSinkData1,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData2(logSinkDetailsDataSinkData2 LogSinkDetailsDataSinkData2) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData2

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData2: &logSinkDetailsDataSinkData2,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData3(logSinkDetailsDataSinkData3 LogSinkDetailsDataSinkData3) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData3

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData3: &logSinkDetailsDataSinkData3,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData4(logSinkDetailsDataSinkData4 LogSinkDetailsDataSinkData4) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData4

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData4: &logSinkDetailsDataSinkData4,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData5(logSinkDetailsDataSinkData5 LogSinkDetailsDataSinkData5) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData5

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData5: &logSinkDetailsDataSinkData5,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData6(logSinkDetailsDataSinkData6 LogSinkDetailsDataSinkData6) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData6

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData6: &logSinkDetailsDataSinkData6,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData7(logSinkDetailsDataSinkData7 LogSinkDetailsDataSinkData7) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData7

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData7: &logSinkDetailsDataSinkData7,
		Type:                        typ,
	}
}

func CreateLogSinkDetailsDataSinkDataLogSinkDetailsDataSinkData8(logSinkDetailsDataSinkData8 LogSinkDetailsDataSinkData8) LogSinkDetailsDataSinkData {
	typ := LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData8

	return LogSinkDetailsDataSinkData{
		LogSinkDetailsDataSinkData8: &logSinkDetailsDataSinkData8,
		Type:                        typ,
	}
}

func (u *LogSinkDetailsDataSinkData) UnmarshalJSON(data []byte) error {

	logSinkDetailsDataSinkData6 := new(LogSinkDetailsDataSinkData6)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData6, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData6 = logSinkDetailsDataSinkData6
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData6
		return nil
	}

	logSinkDetailsDataSinkData7 := new(LogSinkDetailsDataSinkData7)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData7, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData7 = logSinkDetailsDataSinkData7
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData7
		return nil
	}

	logSinkDetailsDataSinkData2 := new(LogSinkDetailsDataSinkData2)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData2, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData2 = logSinkDetailsDataSinkData2
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData2
		return nil
	}

	logSinkDetailsDataSinkData8 := new(LogSinkDetailsDataSinkData8)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData8, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData8 = logSinkDetailsDataSinkData8
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData8
		return nil
	}

	logSinkDetailsDataSinkData1 := new(LogSinkDetailsDataSinkData1)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData1, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData1 = logSinkDetailsDataSinkData1
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData1
		return nil
	}

	logSinkDetailsDataSinkData4 := new(LogSinkDetailsDataSinkData4)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData4, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData4 = logSinkDetailsDataSinkData4
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData4
		return nil
	}

	logSinkDetailsDataSinkData5 := new(LogSinkDetailsDataSinkData5)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData5, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData5 = logSinkDetailsDataSinkData5
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData5
		return nil
	}

	logSinkDetailsDataSinkData3 := new(LogSinkDetailsDataSinkData3)
	if err := utils.UnmarshalJSON(data, &logSinkDetailsDataSinkData3, "", true, true); err == nil {
		u.LogSinkDetailsDataSinkData3 = logSinkDetailsDataSinkData3
		u.Type = LogSinkDetailsDataSinkDataTypeLogSinkDetailsDataSinkData3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkDetailsDataSinkData) MarshalJSON() ([]byte, error) {
	if u.LogSinkDetailsDataSinkData1 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData1, "", true)
	}

	if u.LogSinkDetailsDataSinkData2 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData2, "", true)
	}

	if u.LogSinkDetailsDataSinkData3 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData3, "", true)
	}

	if u.LogSinkDetailsDataSinkData4 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData4, "", true)
	}

	if u.LogSinkDetailsDataSinkData5 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData5, "", true)
	}

	if u.LogSinkDetailsDataSinkData6 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData6, "", true)
	}

	if u.LogSinkDetailsDataSinkData7 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData7, "", true)
	}

	if u.LogSinkDetailsDataSinkData8 != nil {
		return utils.MarshalJSON(u.LogSinkDetailsDataSinkData8, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// LogSinkDetailsDataSinkType - The type of the log sink.
type LogSinkDetailsDataSinkType string

const (
	LogSinkDetailsDataSinkTypeLoki        LogSinkDetailsDataSinkType = "loki"
	LogSinkDetailsDataSinkTypeDatadogLogs LogSinkDetailsDataSinkType = "datadog_logs"
	LogSinkDetailsDataSinkTypePapertrail  LogSinkDetailsDataSinkType = "papertrail"
	LogSinkDetailsDataSinkTypeHTTP        LogSinkDetailsDataSinkType = "http"
	LogSinkDetailsDataSinkTypeAwsS3       LogSinkDetailsDataSinkType = "aws_s3"
	LogSinkDetailsDataSinkTypeLogdna      LogSinkDetailsDataSinkType = "logdna"
	LogSinkDetailsDataSinkTypeCoralogix   LogSinkDetailsDataSinkType = "coralogix"
	LogSinkDetailsDataSinkTypeLogtail     LogSinkDetailsDataSinkType = "logtail"
	LogSinkDetailsDataSinkTypeHoneycomb   LogSinkDetailsDataSinkType = "honeycomb"
	LogSinkDetailsDataSinkTypeLogzio      LogSinkDetailsDataSinkType = "logzio"
)

func (e LogSinkDetailsDataSinkType) ToPointer() *LogSinkDetailsDataSinkType {
	return &e
}

func (e *LogSinkDetailsDataSinkType) UnmarshalJSON(data []byte) error {
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
		*e = LogSinkDetailsDataSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkDetailsDataSinkType: %v", v)
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
	SinkData *LogSinkDetailsDataSinkData `json:"sinkData,omitempty"`
	// The type of the log sink.
	SinkType LogSinkDetailsDataSinkType `json:"sinkType"`
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

func (o *LogSinkDetailsData) GetSinkData() *LogSinkDetailsDataSinkData {
	if o == nil {
		return nil
	}
	return o.SinkData
}

func (o *LogSinkDetailsData) GetSinkType() LogSinkDetailsDataSinkType {
	if o == nil {
		return LogSinkDetailsDataSinkType("")
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
