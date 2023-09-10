// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// HTTPLogSinkSinkDataAuth3Strategy - Bearer token authentication strategy.
type HTTPLogSinkSinkDataAuth3Strategy string

const (
	HTTPLogSinkSinkDataAuth3StrategyBearer HTTPLogSinkSinkDataAuth3Strategy = "bearer"
)

func (e HTTPLogSinkSinkDataAuth3Strategy) ToPointer() *HTTPLogSinkSinkDataAuth3Strategy {
	return &e
}

func (e *HTTPLogSinkSinkDataAuth3Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bearer":
		*e = HTTPLogSinkSinkDataAuth3Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HTTPLogSinkSinkDataAuth3Strategy: %v", v)
	}
}

// HTTPLogSinkSinkDataAuth3 - Authenticate with a bearer token strategy.
type HTTPLogSinkSinkDataAuth3 struct {
	// Bearer token authentication strategy.
	Strategy HTTPLogSinkSinkDataAuth3Strategy `json:"strategy"`
	// Token for bearer token authentication.
	Token *string `json:"token,omitempty"`
}

func (o *HTTPLogSinkSinkDataAuth3) GetStrategy() HTTPLogSinkSinkDataAuth3Strategy {
	if o == nil {
		return HTTPLogSinkSinkDataAuth3Strategy("")
	}
	return o.Strategy
}

func (o *HTTPLogSinkSinkDataAuth3) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// HTTPLogSinkSinkDataAuth2Strategy - Basic HTTP authentication strategy.
type HTTPLogSinkSinkDataAuth2Strategy string

const (
	HTTPLogSinkSinkDataAuth2StrategyBasic HTTPLogSinkSinkDataAuth2Strategy = "basic"
)

func (e HTTPLogSinkSinkDataAuth2Strategy) ToPointer() *HTTPLogSinkSinkDataAuth2Strategy {
	return &e
}

func (e *HTTPLogSinkSinkDataAuth2Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = HTTPLogSinkSinkDataAuth2Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HTTPLogSinkSinkDataAuth2Strategy: %v", v)
	}
}

// HTTPLogSinkSinkDataAuth2 - Authenticate with a basic http strategy.
type HTTPLogSinkSinkDataAuth2 struct {
	// Password for basic http authentication.
	Password string `json:"password"`
	// Basic HTTP authentication strategy.
	Strategy HTTPLogSinkSinkDataAuth2Strategy `json:"strategy"`
	// Username for basic http authentication.
	User *string `json:"user,omitempty"`
}

func (o *HTTPLogSinkSinkDataAuth2) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *HTTPLogSinkSinkDataAuth2) GetStrategy() HTTPLogSinkSinkDataAuth2Strategy {
	if o == nil {
		return HTTPLogSinkSinkDataAuth2Strategy("")
	}
	return o.Strategy
}

func (o *HTTPLogSinkSinkDataAuth2) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

// HTTPLogSinkSinkDataAuth1Strategy - No authentication strategy
type HTTPLogSinkSinkDataAuth1Strategy string

const (
	HTTPLogSinkSinkDataAuth1StrategyNone HTTPLogSinkSinkDataAuth1Strategy = "none"
)

func (e HTTPLogSinkSinkDataAuth1Strategy) ToPointer() *HTTPLogSinkSinkDataAuth1Strategy {
	return &e
}

func (e *HTTPLogSinkSinkDataAuth1Strategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = HTTPLogSinkSinkDataAuth1Strategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HTTPLogSinkSinkDataAuth1Strategy: %v", v)
	}
}

// HTTPLogSinkSinkDataAuth1 - No authentication strategy
type HTTPLogSinkSinkDataAuth1 struct {
	// No authentication strategy
	Strategy HTTPLogSinkSinkDataAuth1Strategy `json:"strategy"`
}

func (o *HTTPLogSinkSinkDataAuth1) GetStrategy() HTTPLogSinkSinkDataAuth1Strategy {
	if o == nil {
		return HTTPLogSinkSinkDataAuth1Strategy("")
	}
	return o.Strategy
}

type HTTPLogSinkSinkDataAuthType string

const (
	HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth1 HTTPLogSinkSinkDataAuthType = "HTTPLogSink_sinkData_auth_1"
	HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth2 HTTPLogSinkSinkDataAuthType = "HTTPLogSink_sinkData_auth_2"
	HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth3 HTTPLogSinkSinkDataAuthType = "HTTPLogSink_sinkData_auth_3"
)

type HTTPLogSinkSinkDataAuth struct {
	HTTPLogSinkSinkDataAuth1 *HTTPLogSinkSinkDataAuth1
	HTTPLogSinkSinkDataAuth2 *HTTPLogSinkSinkDataAuth2
	HTTPLogSinkSinkDataAuth3 *HTTPLogSinkSinkDataAuth3

	Type HTTPLogSinkSinkDataAuthType
}

func CreateHTTPLogSinkSinkDataAuthHTTPLogSinkSinkDataAuth1(httpLogSinkSinkDataAuth1 HTTPLogSinkSinkDataAuth1) HTTPLogSinkSinkDataAuth {
	typ := HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth1

	return HTTPLogSinkSinkDataAuth{
		HTTPLogSinkSinkDataAuth1: &httpLogSinkSinkDataAuth1,
		Type:                     typ,
	}
}

func CreateHTTPLogSinkSinkDataAuthHTTPLogSinkSinkDataAuth2(httpLogSinkSinkDataAuth2 HTTPLogSinkSinkDataAuth2) HTTPLogSinkSinkDataAuth {
	typ := HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth2

	return HTTPLogSinkSinkDataAuth{
		HTTPLogSinkSinkDataAuth2: &httpLogSinkSinkDataAuth2,
		Type:                     typ,
	}
}

func CreateHTTPLogSinkSinkDataAuthHTTPLogSinkSinkDataAuth3(httpLogSinkSinkDataAuth3 HTTPLogSinkSinkDataAuth3) HTTPLogSinkSinkDataAuth {
	typ := HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth3

	return HTTPLogSinkSinkDataAuth{
		HTTPLogSinkSinkDataAuth3: &httpLogSinkSinkDataAuth3,
		Type:                     typ,
	}
}

func (u *HTTPLogSinkSinkDataAuth) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	httpLogSinkSinkDataAuth1 := new(HTTPLogSinkSinkDataAuth1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&httpLogSinkSinkDataAuth1); err == nil {
		u.HTTPLogSinkSinkDataAuth1 = httpLogSinkSinkDataAuth1
		u.Type = HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth1
		return nil
	}

	httpLogSinkSinkDataAuth3 := new(HTTPLogSinkSinkDataAuth3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&httpLogSinkSinkDataAuth3); err == nil {
		u.HTTPLogSinkSinkDataAuth3 = httpLogSinkSinkDataAuth3
		u.Type = HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth3
		return nil
	}

	httpLogSinkSinkDataAuth2 := new(HTTPLogSinkSinkDataAuth2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&httpLogSinkSinkDataAuth2); err == nil {
		u.HTTPLogSinkSinkDataAuth2 = httpLogSinkSinkDataAuth2
		u.Type = HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u HTTPLogSinkSinkDataAuth) MarshalJSON() ([]byte, error) {
	if u.HTTPLogSinkSinkDataAuth1 != nil {
		return json.Marshal(u.HTTPLogSinkSinkDataAuth1)
	}

	if u.HTTPLogSinkSinkDataAuth3 != nil {
		return json.Marshal(u.HTTPLogSinkSinkDataAuth3)
	}

	if u.HTTPLogSinkSinkDataAuth2 != nil {
		return json.Marshal(u.HTTPLogSinkSinkDataAuth2)
	}

	return nil, nil
}

// HTTPLogSinkSinkDataEncodingCodec - Codec to encode logs in
type HTTPLogSinkSinkDataEncodingCodec string

const (
	HTTPLogSinkSinkDataEncodingCodecText HTTPLogSinkSinkDataEncodingCodec = "text"
	HTTPLogSinkSinkDataEncodingCodecJSON HTTPLogSinkSinkDataEncodingCodec = "json"
)

func (e HTTPLogSinkSinkDataEncodingCodec) ToPointer() *HTTPLogSinkSinkDataEncodingCodec {
	return &e
}

func (e *HTTPLogSinkSinkDataEncodingCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "json":
		*e = HTTPLogSinkSinkDataEncodingCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HTTPLogSinkSinkDataEncodingCodec: %v", v)
	}
}

// HTTPLogSinkSinkDataEncoding - Encoding options
type HTTPLogSinkSinkDataEncoding struct {
	// Codec to encode logs in
	Codec HTTPLogSinkSinkDataEncodingCodec `json:"codec"`
}

func (o *HTTPLogSinkSinkDataEncoding) GetCodec() HTTPLogSinkSinkDataEncodingCodec {
	if o == nil {
		return HTTPLogSinkSinkDataEncodingCodec("")
	}
	return o.Codec
}

// HTTPLogSinkSinkData - Details about the HTTP log sink.
type HTTPLogSinkSinkData struct {
	Auth HTTPLogSinkSinkDataAuth `json:"auth"`
	// Encoding options
	Encoding *HTTPLogSinkSinkDataEncoding `json:"encoding,omitempty"`
	// Uri to send logs to.
	URI string `json:"uri"`
}

func (o *HTTPLogSinkSinkData) GetAuth() HTTPLogSinkSinkDataAuth {
	if o == nil {
		return HTTPLogSinkSinkDataAuth{}
	}
	return o.Auth
}

func (o *HTTPLogSinkSinkData) GetEncoding() *HTTPLogSinkSinkDataEncoding {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *HTTPLogSinkSinkData) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

// HTTPLogSinkSinkType - The type of the log sink.
type HTTPLogSinkSinkType string

const (
	HTTPLogSinkSinkTypeHTTP HTTPLogSinkSinkType = "http"
)

func (e HTTPLogSinkSinkType) ToPointer() *HTTPLogSinkSinkType {
	return &e
}

func (e *HTTPLogSinkSinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		*e = HTTPLogSinkSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HTTPLogSinkSinkType: %v", v)
	}
}

// HTTPLogSink - Create a log sink using HTTP
type HTTPLogSink struct {
	// Description of the log sink.
	Description *string `json:"description,omitempty"`
	// If `true` your network access logs will be forwarded with your workload logs
	ForwardAccessLogs *bool `json:"forwardAccessLogs,omitempty"`
	// Name of the log sink.
	Name string `json:"name"`
	// If `restricted` is `true`, only logs from these projects will be sent to the log sink.
	Projects []string `json:"projects,omitempty"`
	// If `true`, only logs from the projects in `projects` will be sent to the log sink.
	Restricted *bool `json:"restricted,omitempty"`
	// Details about the HTTP log sink.
	SinkData HTTPLogSinkSinkData `json:"sinkData"`
	// The type of the log sink.
	SinkType HTTPLogSinkSinkType `json:"sinkType"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `json:"useCustomLabels,omitempty"`
}

func (o *HTTPLogSink) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *HTTPLogSink) GetForwardAccessLogs() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardAccessLogs
}

func (o *HTTPLogSink) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *HTTPLogSink) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *HTTPLogSink) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *HTTPLogSink) GetSinkData() HTTPLogSinkSinkData {
	if o == nil {
		return HTTPLogSinkSinkData{}
	}
	return o.SinkData
}

func (o *HTTPLogSink) GetSinkType() HTTPLogSinkSinkType {
	if o == nil {
		return HTTPLogSinkSinkType("")
	}
	return o.SinkType
}

func (o *HTTPLogSink) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}
