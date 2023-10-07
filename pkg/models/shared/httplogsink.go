// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
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
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Bearer token authentication strategy.
	Strategy HTTPLogSinkSinkDataAuth3Strategy `json:"strategy"`
	// Token for bearer token authentication.
	Token *string `json:"token,omitempty"`
}

func (h HTTPLogSinkSinkDataAuth3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPLogSinkSinkDataAuth3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *HTTPLogSinkSinkDataAuth3) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Password for basic http authentication.
	Password string `json:"password"`
	// Basic HTTP authentication strategy.
	Strategy HTTPLogSinkSinkDataAuth2Strategy `json:"strategy"`
	// Username for basic http authentication.
	User *string `json:"user,omitempty"`
}

func (h HTTPLogSinkSinkDataAuth2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPLogSinkSinkDataAuth2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *HTTPLogSinkSinkDataAuth2) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// No authentication strategy
	Strategy HTTPLogSinkSinkDataAuth1Strategy `json:"strategy"`
}

func (h HTTPLogSinkSinkDataAuth1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPLogSinkSinkDataAuth1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *HTTPLogSinkSinkDataAuth1) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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

	httpLogSinkSinkDataAuth1 := new(HTTPLogSinkSinkDataAuth1)
	if err := utils.UnmarshalJSON(data, &httpLogSinkSinkDataAuth1, "", true, true); err == nil {
		u.HTTPLogSinkSinkDataAuth1 = httpLogSinkSinkDataAuth1
		u.Type = HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth1
		return nil
	}

	httpLogSinkSinkDataAuth3 := new(HTTPLogSinkSinkDataAuth3)
	if err := utils.UnmarshalJSON(data, &httpLogSinkSinkDataAuth3, "", true, true); err == nil {
		u.HTTPLogSinkSinkDataAuth3 = httpLogSinkSinkDataAuth3
		u.Type = HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth3
		return nil
	}

	httpLogSinkSinkDataAuth2 := new(HTTPLogSinkSinkDataAuth2)
	if err := utils.UnmarshalJSON(data, &httpLogSinkSinkDataAuth2, "", true, true); err == nil {
		u.HTTPLogSinkSinkDataAuth2 = httpLogSinkSinkDataAuth2
		u.Type = HTTPLogSinkSinkDataAuthTypeHTTPLogSinkSinkDataAuth2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u HTTPLogSinkSinkDataAuth) MarshalJSON() ([]byte, error) {
	if u.HTTPLogSinkSinkDataAuth1 != nil {
		return utils.MarshalJSON(u.HTTPLogSinkSinkDataAuth1, "", true)
	}

	if u.HTTPLogSinkSinkDataAuth2 != nil {
		return utils.MarshalJSON(u.HTTPLogSinkSinkDataAuth2, "", true)
	}

	if u.HTTPLogSinkSinkDataAuth3 != nil {
		return utils.MarshalJSON(u.HTTPLogSinkSinkDataAuth3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
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
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Codec to encode logs in
	Codec HTTPLogSinkSinkDataEncodingCodec `json:"codec"`
}

func (h HTTPLogSinkSinkDataEncoding) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPLogSinkSinkDataEncoding) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HTTPLogSinkSinkDataEncoding) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *HTTPLogSinkSinkDataEncoding) GetCodec() HTTPLogSinkSinkDataEncodingCodec {
	if o == nil {
		return HTTPLogSinkSinkDataEncodingCodec("")
	}
	return o.Codec
}

// HTTPLogSinkSinkData - Details about the HTTP log sink.
type HTTPLogSinkSinkData struct {
	AdditionalProperties map[string]interface{}  `additionalProperties:"true" json:"-"`
	Auth                 HTTPLogSinkSinkDataAuth `json:"auth"`
	// Encoding options
	Encoding *HTTPLogSinkSinkDataEncoding `json:"encoding,omitempty"`
	// Uri to send logs to.
	URI string `json:"uri"`
}

func (h HTTPLogSinkSinkData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPLogSinkSinkData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HTTPLogSinkSinkData) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
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
	// Details about the HTTP log sink.
	SinkData HTTPLogSinkSinkData `json:"sinkData"`
	// The type of the log sink.
	SinkType HTTPLogSinkSinkType `json:"sinkType"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `default:"false" json:"useCustomLabels"`
}

func (h HTTPLogSink) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPLogSink) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *HTTPLogSink) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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
