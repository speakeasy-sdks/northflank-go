// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// LogSinkRequestSinkData8 - Honeycomb Sink Schema.
type LogSinkRequestSinkData8 struct {
	// Honeycomb API Key
	APIKey string `json:"api_key"`
	// Name of the dataset
	Dataset string `json:"dataset"`
}

func (o *LogSinkRequestSinkData8) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *LogSinkRequestSinkData8) GetDataset() string {
	if o == nil {
		return ""
	}
	return o.Dataset
}

// LogSinkRequestSinkData7 - LogDNA Sink Schema.
type LogSinkRequestSinkData7 struct {
	// Ingestion Key
	APIKey string `json:"api_key"`
}

func (o *LogSinkRequestSinkData7) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

// LogSinkRequestSinkData6 - Logtail Sink Schema.
type LogSinkRequestSinkData6 struct {
	// Logtail Source Token
	Token string `json:"token"`
}

func (o *LogSinkRequestSinkData6) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// LogSinkRequestSinkData5Auth - Authentication object.
type LogSinkRequestSinkData5Auth struct {
	// Access key id for the bucket.
	AccessKeyID string `json:"accessKeyId"`
	// Secret access key for the bucket.
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *LogSinkRequestSinkData5Auth) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *LogSinkRequestSinkData5Auth) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// LogSinkRequestSinkData5Compression - Log file compression method.
type LogSinkRequestSinkData5Compression string

const (
	LogSinkRequestSinkData5CompressionGzip LogSinkRequestSinkData5Compression = "gzip"
	LogSinkRequestSinkData5CompressionNone LogSinkRequestSinkData5Compression = "none"
)

func (e LogSinkRequestSinkData5Compression) ToPointer() *LogSinkRequestSinkData5Compression {
	return &e
}

func (e *LogSinkRequestSinkData5Compression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gzip":
		fallthrough
	case "none":
		*e = LogSinkRequestSinkData5Compression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData5Compression: %v", v)
	}
}

// LogSinkRequestSinkData5Region - Region of the S3 bucket.
type LogSinkRequestSinkData5Region string

const (
	LogSinkRequestSinkData5RegionEuWest1    LogSinkRequestSinkData5Region = "eu-west-1"
	LogSinkRequestSinkData5RegionEuWest2    LogSinkRequestSinkData5Region = "eu-west-2"
	LogSinkRequestSinkData5RegionEuWest3    LogSinkRequestSinkData5Region = "eu-west-3"
	LogSinkRequestSinkData5RegionEuCentral1 LogSinkRequestSinkData5Region = "eu-central-1"
	LogSinkRequestSinkData5RegionEuSouth1   LogSinkRequestSinkData5Region = "eu-south-1"
	LogSinkRequestSinkData5RegionEuNorth1   LogSinkRequestSinkData5Region = "eu-north-1"
	LogSinkRequestSinkData5RegionUsWest1    LogSinkRequestSinkData5Region = "us-west-1"
	LogSinkRequestSinkData5RegionUsWest2    LogSinkRequestSinkData5Region = "us-west-2"
	LogSinkRequestSinkData5RegionUsEast1    LogSinkRequestSinkData5Region = "us-east-1"
	LogSinkRequestSinkData5RegionUsEast2    LogSinkRequestSinkData5Region = "us-east2"
)

func (e LogSinkRequestSinkData5Region) ToPointer() *LogSinkRequestSinkData5Region {
	return &e
}

func (e *LogSinkRequestSinkData5Region) UnmarshalJSON(data []byte) error {
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
		*e = LogSinkRequestSinkData5Region(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData5Region: %v", v)
	}
}

// LogSinkRequestSinkData5 - AWS S3 or compatible API Sink Schema.
type LogSinkRequestSinkData5 struct {
	// Authentication object.
	Auth *LogSinkRequestSinkData5Auth `json:"auth,omitempty"`
	// Name of the S3 Bucket.
	Bucket string `json:"bucket"`
	// Log file compression method.
	Compression LogSinkRequestSinkData5Compression `json:"compression"`
	// Endpoint for the AWS S3 or compatible API bucket.
	Endpoint string `json:"endpoint"`
	// Region of the S3 bucket.
	Region LogSinkRequestSinkData5Region `json:"region"`
}

func (o *LogSinkRequestSinkData5) GetAuth() *LogSinkRequestSinkData5Auth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *LogSinkRequestSinkData5) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *LogSinkRequestSinkData5) GetCompression() LogSinkRequestSinkData5Compression {
	if o == nil {
		return LogSinkRequestSinkData5Compression("")
	}
	return o.Compression
}

func (o *LogSinkRequestSinkData5) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *LogSinkRequestSinkData5) GetRegion() LogSinkRequestSinkData5Region {
	if o == nil {
		return LogSinkRequestSinkData5Region("")
	}
	return o.Region
}

// LogSinkRequestSinkData4Auth3Strategy - Bearer token authentication strategy.
type LogSinkRequestSinkData4Auth3Strategy string

const (
	LogSinkRequestSinkData4Auth3StrategyBearer LogSinkRequestSinkData4Auth3Strategy = "bearer"
)

func (e LogSinkRequestSinkData4Auth3Strategy) ToPointer() *LogSinkRequestSinkData4Auth3Strategy {
	return &e
}

func (e *LogSinkRequestSinkData4Auth3Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bearer":
		*e = LogSinkRequestSinkData4Auth3Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData4Auth3Strategy: %v", v)
	}
}

// LogSinkRequestSinkData4Auth3 - Authenticate with a bearer token strategy.
type LogSinkRequestSinkData4Auth3 struct {
	// Bearer token authentication strategy.
	Strategy LogSinkRequestSinkData4Auth3Strategy `json:"strategy"`
	// Token for bearer token authentication.
	Token *string `json:"token,omitempty"`
}

func (o *LogSinkRequestSinkData4Auth3) GetStrategy() LogSinkRequestSinkData4Auth3Strategy {
	if o == nil {
		return LogSinkRequestSinkData4Auth3Strategy("")
	}
	return o.Strategy
}

func (o *LogSinkRequestSinkData4Auth3) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// LogSinkRequestSinkData4Auth2Strategy - Basic HTTP authentication strategy.
type LogSinkRequestSinkData4Auth2Strategy string

const (
	LogSinkRequestSinkData4Auth2StrategyBasic LogSinkRequestSinkData4Auth2Strategy = "basic"
)

func (e LogSinkRequestSinkData4Auth2Strategy) ToPointer() *LogSinkRequestSinkData4Auth2Strategy {
	return &e
}

func (e *LogSinkRequestSinkData4Auth2Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = LogSinkRequestSinkData4Auth2Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData4Auth2Strategy: %v", v)
	}
}

// LogSinkRequestSinkData4Auth2 - Authenticate with a basic http strategy.
type LogSinkRequestSinkData4Auth2 struct {
	// Password for basic http authentication.
	Password string `json:"password"`
	// Basic HTTP authentication strategy.
	Strategy LogSinkRequestSinkData4Auth2Strategy `json:"strategy"`
	// Username for basic http authentication.
	User *string `json:"user,omitempty"`
}

func (o *LogSinkRequestSinkData4Auth2) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *LogSinkRequestSinkData4Auth2) GetStrategy() LogSinkRequestSinkData4Auth2Strategy {
	if o == nil {
		return LogSinkRequestSinkData4Auth2Strategy("")
	}
	return o.Strategy
}

func (o *LogSinkRequestSinkData4Auth2) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// LogSinkRequestSinkData4Auth1Strategy - No authentication strategy
type LogSinkRequestSinkData4Auth1Strategy string

const (
	LogSinkRequestSinkData4Auth1StrategyNone LogSinkRequestSinkData4Auth1Strategy = "none"
)

func (e LogSinkRequestSinkData4Auth1Strategy) ToPointer() *LogSinkRequestSinkData4Auth1Strategy {
	return &e
}

func (e *LogSinkRequestSinkData4Auth1Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = LogSinkRequestSinkData4Auth1Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData4Auth1Strategy: %v", v)
	}
}

// LogSinkRequestSinkData4Auth1 - No authentication strategy
type LogSinkRequestSinkData4Auth1 struct {
	// No authentication strategy
	Strategy LogSinkRequestSinkData4Auth1Strategy `json:"strategy"`
}

func (o *LogSinkRequestSinkData4Auth1) GetStrategy() LogSinkRequestSinkData4Auth1Strategy {
	if o == nil {
		return LogSinkRequestSinkData4Auth1Strategy("")
	}
	return o.Strategy
}

type LogSinkRequestSinkData4AuthType string

const (
	LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth1 LogSinkRequestSinkData4AuthType = "LogSinkRequest_sinkData_4_auth_1"
	LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth2 LogSinkRequestSinkData4AuthType = "LogSinkRequest_sinkData_4_auth_2"
	LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth3 LogSinkRequestSinkData4AuthType = "LogSinkRequest_sinkData_4_auth_3"
)

type LogSinkRequestSinkData4Auth struct {
	LogSinkRequestSinkData4Auth1 *LogSinkRequestSinkData4Auth1
	LogSinkRequestSinkData4Auth2 *LogSinkRequestSinkData4Auth2
	LogSinkRequestSinkData4Auth3 *LogSinkRequestSinkData4Auth3

	Type LogSinkRequestSinkData4AuthType
}

func CreateLogSinkRequestSinkData4AuthLogSinkRequestSinkData4Auth1(logSinkRequestSinkData4Auth1 LogSinkRequestSinkData4Auth1) LogSinkRequestSinkData4Auth {
	typ := LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth1

	return LogSinkRequestSinkData4Auth{
		LogSinkRequestSinkData4Auth1: &logSinkRequestSinkData4Auth1,
		Type:                         typ,
	}
}

func CreateLogSinkRequestSinkData4AuthLogSinkRequestSinkData4Auth2(logSinkRequestSinkData4Auth2 LogSinkRequestSinkData4Auth2) LogSinkRequestSinkData4Auth {
	typ := LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth2

	return LogSinkRequestSinkData4Auth{
		LogSinkRequestSinkData4Auth2: &logSinkRequestSinkData4Auth2,
		Type:                         typ,
	}
}

func CreateLogSinkRequestSinkData4AuthLogSinkRequestSinkData4Auth3(logSinkRequestSinkData4Auth3 LogSinkRequestSinkData4Auth3) LogSinkRequestSinkData4Auth {
	typ := LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth3

	return LogSinkRequestSinkData4Auth{
		LogSinkRequestSinkData4Auth3: &logSinkRequestSinkData4Auth3,
		Type:                         typ,
	}
}

func (u *LogSinkRequestSinkData4Auth) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	logSinkRequestSinkData4Auth1 := new(LogSinkRequestSinkData4Auth1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData4Auth1); err == nil {
		u.LogSinkRequestSinkData4Auth1 = logSinkRequestSinkData4Auth1
		u.Type = LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth1
		return nil
	}

	logSinkRequestSinkData4Auth2 := new(LogSinkRequestSinkData4Auth2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData4Auth2); err == nil {
		u.LogSinkRequestSinkData4Auth2 = logSinkRequestSinkData4Auth2
		u.Type = LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth2
		return nil
	}

	logSinkRequestSinkData4Auth3 := new(LogSinkRequestSinkData4Auth3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData4Auth3); err == nil {
		u.LogSinkRequestSinkData4Auth3 = logSinkRequestSinkData4Auth3
		u.Type = LogSinkRequestSinkData4AuthTypeLogSinkRequestSinkData4Auth3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkRequestSinkData4Auth) MarshalJSON() ([]byte, error) {
	if u.LogSinkRequestSinkData4Auth1 != nil {
		return json.Marshal(u.LogSinkRequestSinkData4Auth1)
	}

	if u.LogSinkRequestSinkData4Auth2 != nil {
		return json.Marshal(u.LogSinkRequestSinkData4Auth2)
	}

	if u.LogSinkRequestSinkData4Auth3 != nil {
		return json.Marshal(u.LogSinkRequestSinkData4Auth3)
	}

	return nil, nil
}

// LogSinkRequestSinkData4EncodingCodec - Codec to encode logs in
type LogSinkRequestSinkData4EncodingCodec string

const (
	LogSinkRequestSinkData4EncodingCodecText LogSinkRequestSinkData4EncodingCodec = "text"
	LogSinkRequestSinkData4EncodingCodecJSON LogSinkRequestSinkData4EncodingCodec = "json"
)

func (e LogSinkRequestSinkData4EncodingCodec) ToPointer() *LogSinkRequestSinkData4EncodingCodec {
	return &e
}

func (e *LogSinkRequestSinkData4EncodingCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json":
		*e = LogSinkRequestSinkData4EncodingCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData4EncodingCodec: %v", v)
	}
}

// LogSinkRequestSinkData4Encoding - Encoding options
type LogSinkRequestSinkData4Encoding struct {
	// Codec to encode logs in
	Codec LogSinkRequestSinkData4EncodingCodec `json:"codec"`
}

func (o *LogSinkRequestSinkData4Encoding) GetCodec() LogSinkRequestSinkData4EncodingCodec {
	if o == nil {
		return LogSinkRequestSinkData4EncodingCodec("")
	}
	return o.Codec
}

// LogSinkRequestSinkData4 - HTTP Sink Schema.
type LogSinkRequestSinkData4 struct {
	Auth LogSinkRequestSinkData4Auth `json:"auth"`
	// Encoding options
	Encoding *LogSinkRequestSinkData4Encoding `json:"encoding,omitempty"`
	// Uri to send logs to.
	URI string `json:"uri"`
}

func (o *LogSinkRequestSinkData4) GetAuth() LogSinkRequestSinkData4Auth {
	if o == nil {
		return LogSinkRequestSinkData4Auth{}
	}
	return o.Auth
}

func (o *LogSinkRequestSinkData4) GetEncoding() *LogSinkRequestSinkData4Encoding {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *LogSinkRequestSinkData4) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// LogSinkRequestSinkData32AuthenticationStrategy - The authentication strategy.
type LogSinkRequestSinkData32AuthenticationStrategy string

const (
	LogSinkRequestSinkData32AuthenticationStrategyToken LogSinkRequestSinkData32AuthenticationStrategy = "token"
)

func (e LogSinkRequestSinkData32AuthenticationStrategy) ToPointer() *LogSinkRequestSinkData32AuthenticationStrategy {
	return &e
}

func (e *LogSinkRequestSinkData32AuthenticationStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "token":
		*e = LogSinkRequestSinkData32AuthenticationStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData32AuthenticationStrategy: %v", v)
	}
}

// LogSinkRequestSinkData32 - Authenticate with a token.
type LogSinkRequestSinkData32 struct {
	// The authentication strategy.
	AuthenticationStrategy LogSinkRequestSinkData32AuthenticationStrategy `json:"authenticationStrategy"`
	// The HTTP Token for the Papertrail log destination.
	Token string `json:"token"`
	// The uri for the Papertrail log destination.
	URI string `json:"uri"`
}

func (o *LogSinkRequestSinkData32) GetAuthenticationStrategy() LogSinkRequestSinkData32AuthenticationStrategy {
	if o == nil {
		return LogSinkRequestSinkData32AuthenticationStrategy("")
	}
	return o.AuthenticationStrategy
}

func (o *LogSinkRequestSinkData32) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *LogSinkRequestSinkData32) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// LogSinkRequestSinkData31AuthenticationStrategy - The authentication strategy.
type LogSinkRequestSinkData31AuthenticationStrategy string

const (
	LogSinkRequestSinkData31AuthenticationStrategyPort LogSinkRequestSinkData31AuthenticationStrategy = "port"
)

func (e LogSinkRequestSinkData31AuthenticationStrategy) ToPointer() *LogSinkRequestSinkData31AuthenticationStrategy {
	return &e
}

func (e *LogSinkRequestSinkData31AuthenticationStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "port":
		*e = LogSinkRequestSinkData31AuthenticationStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData31AuthenticationStrategy: %v", v)
	}
}

// LogSinkRequestSinkData31 - Authenticate with a host/port
type LogSinkRequestSinkData31 struct {
	// The authentication strategy.
	AuthenticationStrategy LogSinkRequestSinkData31AuthenticationStrategy `json:"authenticationStrategy"`
	// The host for the Papertrail log destination.
	Host string `json:"host"`
	// The port for the Papertrail log destination.
	Port float32 `json:"port"`
}

func (o *LogSinkRequestSinkData31) GetAuthenticationStrategy() LogSinkRequestSinkData31AuthenticationStrategy {
	if o == nil {
		return LogSinkRequestSinkData31AuthenticationStrategy("")
	}
	return o.AuthenticationStrategy
}

func (o *LogSinkRequestSinkData31) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *LogSinkRequestSinkData31) GetPort() float32 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

type LogSinkRequestSinkData3Type string

const (
	LogSinkRequestSinkData3TypeLogSinkRequestSinkData31 LogSinkRequestSinkData3Type = "LogSinkRequest_sinkData_3_1"
	LogSinkRequestSinkData3TypeLogSinkRequestSinkData32 LogSinkRequestSinkData3Type = "LogSinkRequest_sinkData_3_2"
)

type LogSinkRequestSinkData3 struct {
	LogSinkRequestSinkData31 *LogSinkRequestSinkData31
	LogSinkRequestSinkData32 *LogSinkRequestSinkData32

	Type LogSinkRequestSinkData3Type
}

func CreateLogSinkRequestSinkData3LogSinkRequestSinkData31(logSinkRequestSinkData31 LogSinkRequestSinkData31) LogSinkRequestSinkData3 {
	typ := LogSinkRequestSinkData3TypeLogSinkRequestSinkData31

	return LogSinkRequestSinkData3{
		LogSinkRequestSinkData31: &logSinkRequestSinkData31,
		Type:                     typ,
	}
}

func CreateLogSinkRequestSinkData3LogSinkRequestSinkData32(logSinkRequestSinkData32 LogSinkRequestSinkData32) LogSinkRequestSinkData3 {
	typ := LogSinkRequestSinkData3TypeLogSinkRequestSinkData32

	return LogSinkRequestSinkData3{
		LogSinkRequestSinkData32: &logSinkRequestSinkData32,
		Type:                     typ,
	}
}

func (u *LogSinkRequestSinkData3) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	logSinkRequestSinkData31 := new(LogSinkRequestSinkData31)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData31); err == nil {
		u.LogSinkRequestSinkData31 = logSinkRequestSinkData31
		u.Type = LogSinkRequestSinkData3TypeLogSinkRequestSinkData31
		return nil
	}

	logSinkRequestSinkData32 := new(LogSinkRequestSinkData32)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData32); err == nil {
		u.LogSinkRequestSinkData32 = logSinkRequestSinkData32
		u.Type = LogSinkRequestSinkData3TypeLogSinkRequestSinkData32
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkRequestSinkData3) MarshalJSON() ([]byte, error) {
	if u.LogSinkRequestSinkData31 != nil {
		return json.Marshal(u.LogSinkRequestSinkData31)
	}

	if u.LogSinkRequestSinkData32 != nil {
		return json.Marshal(u.LogSinkRequestSinkData32)
	}

	return nil, nil
}

// LogSinkRequestSinkData2Region - The Datadog region.
type LogSinkRequestSinkData2Region string

const (
	LogSinkRequestSinkData2RegionEu  LogSinkRequestSinkData2Region = "eu"
	LogSinkRequestSinkData2RegionUs  LogSinkRequestSinkData2Region = "us"
	LogSinkRequestSinkData2RegionUs3 LogSinkRequestSinkData2Region = "us3"
	LogSinkRequestSinkData2RegionUs5 LogSinkRequestSinkData2Region = "us5"
)

func (e LogSinkRequestSinkData2Region) ToPointer() *LogSinkRequestSinkData2Region {
	return &e
}

func (e *LogSinkRequestSinkData2Region) UnmarshalJSON(data []byte) error {
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
		*e = LogSinkRequestSinkData2Region(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData2Region: %v", v)
	}
}

// LogSinkRequestSinkData2 - Datadog Sink Schema.
type LogSinkRequestSinkData2 struct {
	// The Datadog API key.
	DefaultAPIKey *string `json:"default_api_key,omitempty"`
	// The Datadog region.
	Region *LogSinkRequestSinkData2Region `json:"region,omitempty"`
}

func (o *LogSinkRequestSinkData2) GetDefaultAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.DefaultAPIKey
}

func (o *LogSinkRequestSinkData2) GetRegion() *LogSinkRequestSinkData2Region {
	if o == nil {
		return nil
	}
	return o.Region
}

// LogSinkRequestSinkData1AuthStrategy - The authentication method.
type LogSinkRequestSinkData1AuthStrategy string

const (
	LogSinkRequestSinkData1AuthStrategyBasic LogSinkRequestSinkData1AuthStrategy = "basic"
)

func (e LogSinkRequestSinkData1AuthStrategy) ToPointer() *LogSinkRequestSinkData1AuthStrategy {
	return &e
}

func (e *LogSinkRequestSinkData1AuthStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = LogSinkRequestSinkData1AuthStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogSinkRequestSinkData1AuthStrategy: %v", v)
	}
}

// LogSinkRequestSinkData1Auth - Object containing authentication data for the log sink.
type LogSinkRequestSinkData1Auth struct {
	// The password for the log sink.
	Password *string `json:"password,omitempty"`
	// The authentication method.
	Strategy *LogSinkRequestSinkData1AuthStrategy `json:"strategy,omitempty"`
	// The username for the log sink.
	User *string `json:"user,omitempty"`
}

func (o *LogSinkRequestSinkData1Auth) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LogSinkRequestSinkData1Auth) GetStrategy() *LogSinkRequestSinkData1AuthStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *LogSinkRequestSinkData1Auth) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// LogSinkRequestSinkData1 - Loki Sink Schema.
type LogSinkRequestSinkData1 struct {
	// Object containing authentication data for the log sink.
	Auth *LogSinkRequestSinkData1Auth `json:"auth,omitempty"`
	// The endpoint of the Loki log sink.
	Endpoint *string `json:"endpoint,omitempty"`
}

func (o *LogSinkRequestSinkData1) GetAuth() *LogSinkRequestSinkData1Auth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *LogSinkRequestSinkData1) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

type LogSinkRequestSinkDataType string

const (
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData1 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_1"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData2 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_2"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData3 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_3"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData4 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_4"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData5 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_5"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData6 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_6"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData7 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_7"
	LogSinkRequestSinkDataTypeLogSinkRequestSinkData8 LogSinkRequestSinkDataType = "LogSinkRequest_sinkData_8"
)

type LogSinkRequestSinkData struct {
	LogSinkRequestSinkData1 *LogSinkRequestSinkData1
	LogSinkRequestSinkData2 *LogSinkRequestSinkData2
	LogSinkRequestSinkData3 *LogSinkRequestSinkData3
	LogSinkRequestSinkData4 *LogSinkRequestSinkData4
	LogSinkRequestSinkData5 *LogSinkRequestSinkData5
	LogSinkRequestSinkData6 *LogSinkRequestSinkData6
	LogSinkRequestSinkData7 *LogSinkRequestSinkData7
	LogSinkRequestSinkData8 *LogSinkRequestSinkData8

	Type LogSinkRequestSinkDataType
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData1(logSinkRequestSinkData1 LogSinkRequestSinkData1) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData1

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData1: &logSinkRequestSinkData1,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData2(logSinkRequestSinkData2 LogSinkRequestSinkData2) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData2

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData2: &logSinkRequestSinkData2,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData3(logSinkRequestSinkData3 LogSinkRequestSinkData3) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData3

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData3: &logSinkRequestSinkData3,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData4(logSinkRequestSinkData4 LogSinkRequestSinkData4) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData4

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData4: &logSinkRequestSinkData4,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData5(logSinkRequestSinkData5 LogSinkRequestSinkData5) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData5

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData5: &logSinkRequestSinkData5,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData6(logSinkRequestSinkData6 LogSinkRequestSinkData6) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData6

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData6: &logSinkRequestSinkData6,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData7(logSinkRequestSinkData7 LogSinkRequestSinkData7) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData7

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData7: &logSinkRequestSinkData7,
		Type:                    typ,
	}
}

func CreateLogSinkRequestSinkDataLogSinkRequestSinkData8(logSinkRequestSinkData8 LogSinkRequestSinkData8) LogSinkRequestSinkData {
	typ := LogSinkRequestSinkDataTypeLogSinkRequestSinkData8

	return LogSinkRequestSinkData{
		LogSinkRequestSinkData8: &logSinkRequestSinkData8,
		Type:                    typ,
	}
}

func (u *LogSinkRequestSinkData) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	logSinkRequestSinkData1 := new(LogSinkRequestSinkData1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData1); err == nil {
		u.LogSinkRequestSinkData1 = logSinkRequestSinkData1
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData1
		return nil
	}

	logSinkRequestSinkData2 := new(LogSinkRequestSinkData2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData2); err == nil {
		u.LogSinkRequestSinkData2 = logSinkRequestSinkData2
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData2
		return nil
	}

	logSinkRequestSinkData3 := new(LogSinkRequestSinkData3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData3); err == nil {
		u.LogSinkRequestSinkData3 = logSinkRequestSinkData3
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData3
		return nil
	}

	logSinkRequestSinkData4 := new(LogSinkRequestSinkData4)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData4); err == nil {
		u.LogSinkRequestSinkData4 = logSinkRequestSinkData4
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData4
		return nil
	}

	logSinkRequestSinkData5 := new(LogSinkRequestSinkData5)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData5); err == nil {
		u.LogSinkRequestSinkData5 = logSinkRequestSinkData5
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData5
		return nil
	}

	logSinkRequestSinkData6 := new(LogSinkRequestSinkData6)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData6); err == nil {
		u.LogSinkRequestSinkData6 = logSinkRequestSinkData6
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData6
		return nil
	}

	logSinkRequestSinkData7 := new(LogSinkRequestSinkData7)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData7); err == nil {
		u.LogSinkRequestSinkData7 = logSinkRequestSinkData7
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData7
		return nil
	}

	logSinkRequestSinkData8 := new(LogSinkRequestSinkData8)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&logSinkRequestSinkData8); err == nil {
		u.LogSinkRequestSinkData8 = logSinkRequestSinkData8
		u.Type = LogSinkRequestSinkDataTypeLogSinkRequestSinkData8
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LogSinkRequestSinkData) MarshalJSON() ([]byte, error) {
	if u.LogSinkRequestSinkData1 != nil {
		return json.Marshal(u.LogSinkRequestSinkData1)
	}

	if u.LogSinkRequestSinkData2 != nil {
		return json.Marshal(u.LogSinkRequestSinkData2)
	}

	if u.LogSinkRequestSinkData3 != nil {
		return json.Marshal(u.LogSinkRequestSinkData3)
	}

	if u.LogSinkRequestSinkData4 != nil {
		return json.Marshal(u.LogSinkRequestSinkData4)
	}

	if u.LogSinkRequestSinkData5 != nil {
		return json.Marshal(u.LogSinkRequestSinkData5)
	}

	if u.LogSinkRequestSinkData6 != nil {
		return json.Marshal(u.LogSinkRequestSinkData6)
	}

	if u.LogSinkRequestSinkData7 != nil {
		return json.Marshal(u.LogSinkRequestSinkData7)
	}

	if u.LogSinkRequestSinkData8 != nil {
		return json.Marshal(u.LogSinkRequestSinkData8)
	}

	return nil, nil
}

// LogSinkRequest - Request body
type LogSinkRequest struct {
	// If `restricted` is `true`, only logs from these projects will be sent to the log sink.
	Projects []string `json:"projects,omitempty"`
	// If `true`, only logs from the projects in `projects` will be sent to the log sink.
	Restricted *bool `json:"restricted,omitempty"`
	// If `true`, and the log sink is currently paused, the log sink will be resumed after updating.
	ResumeLogSink *bool `json:"resumeLogSink,omitempty"`
	// Data about the log sink.
	SinkData *LogSinkRequestSinkData `json:"sinkData,omitempty"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `json:"useCustomLabels,omitempty"`
}

func (o *LogSinkRequest) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *LogSinkRequest) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *LogSinkRequest) GetResumeLogSink() *bool {
	if o == nil {
		return nil
	}
	return o.ResumeLogSink
}

func (o *LogSinkRequest) GetSinkData() *LogSinkRequestSinkData {
	if o == nil {
		return nil
	}
	return o.SinkData
}

func (o *LogSinkRequest) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}
