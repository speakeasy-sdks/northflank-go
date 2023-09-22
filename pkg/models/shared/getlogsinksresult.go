// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
	"time"
)

// GetLogSinksResultDataLogSinksSinkType - The type of the log sink.
type GetLogSinksResultDataLogSinksSinkType string

const (
	GetLogSinksResultDataLogSinksSinkTypeLoki        GetLogSinksResultDataLogSinksSinkType = "loki"
	GetLogSinksResultDataLogSinksSinkTypeDatadogLogs GetLogSinksResultDataLogSinksSinkType = "datadog_logs"
	GetLogSinksResultDataLogSinksSinkTypePapertrail  GetLogSinksResultDataLogSinksSinkType = "papertrail"
	GetLogSinksResultDataLogSinksSinkTypeHTTP        GetLogSinksResultDataLogSinksSinkType = "http"
	GetLogSinksResultDataLogSinksSinkTypeAwsS3       GetLogSinksResultDataLogSinksSinkType = "aws_s3"
	GetLogSinksResultDataLogSinksSinkTypeLogdna      GetLogSinksResultDataLogSinksSinkType = "logdna"
	GetLogSinksResultDataLogSinksSinkTypeCoralogix   GetLogSinksResultDataLogSinksSinkType = "coralogix"
	GetLogSinksResultDataLogSinksSinkTypeLogtail     GetLogSinksResultDataLogSinksSinkType = "logtail"
	GetLogSinksResultDataLogSinksSinkTypeHoneycomb   GetLogSinksResultDataLogSinksSinkType = "honeycomb"
	GetLogSinksResultDataLogSinksSinkTypeLogzio      GetLogSinksResultDataLogSinksSinkType = "logzio"
)

func (e GetLogSinksResultDataLogSinksSinkType) ToPointer() *GetLogSinksResultDataLogSinksSinkType {
	return &e
}

func (e *GetLogSinksResultDataLogSinksSinkType) UnmarshalJSON(data []byte) error {
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
		*e = GetLogSinksResultDataLogSinksSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetLogSinksResultDataLogSinksSinkType: %v", v)
	}
}

// GetLogSinksResultDataLogSinks - A log sink object.
type GetLogSinksResultDataLogSinks struct {
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
	// The type of the log sink.
	SinkType GetLogSinksResultDataLogSinksSinkType `json:"sinkType"`
	// Timestamp of when the log sink was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `default:"false" json:"useCustomLabels"`
}

func (g GetLogSinksResultDataLogSinks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLogSinksResultDataLogSinks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLogSinksResultDataLogSinks) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GetLogSinksResultDataLogSinks) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GetLogSinksResultDataLogSinks) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetLogSinksResultDataLogSinks) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetLogSinksResultDataLogSinks) GetProjects() []string {
	if o == nil {
		return []string{}
	}
	return o.Projects
}

func (o *GetLogSinksResultDataLogSinks) GetRestricted() bool {
	if o == nil {
		return false
	}
	return o.Restricted
}

func (o *GetLogSinksResultDataLogSinks) GetSinkType() GetLogSinksResultDataLogSinksSinkType {
	if o == nil {
		return GetLogSinksResultDataLogSinksSinkType("")
	}
	return o.SinkType
}

func (o *GetLogSinksResultDataLogSinks) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *GetLogSinksResultDataLogSinks) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}

// GetLogSinksResultData - Result data.
type GetLogSinksResultData struct {
	// An array of log sinks added to this account.
	LogSinks []GetLogSinksResultDataLogSinks `json:"logSinks"`
}

func (o *GetLogSinksResultData) GetLogSinks() []GetLogSinksResultDataLogSinks {
	if o == nil {
		return []GetLogSinksResultDataLogSinks{}
	}
	return o.LogSinks
}

// GetLogSinksResultPagination - Data about the endpoint pagination.
type GetLogSinksResultPagination struct {
	// The number of results returned by this request.
	Count float32 `json:"count"`
	// The cursor to access the next page of results.
	Cursor *string `json:"cursor,omitempty"`
	// Is there another page of results available?
	HasNextPage bool `json:"hasNextPage"`
}

func (o *GetLogSinksResultPagination) GetCount() float32 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *GetLogSinksResultPagination) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetLogSinksResultPagination) GetHasNextPage() bool {
	if o == nil {
		return false
	}
	return o.HasNextPage
}

// GetLogSinksResult - Response object.
type GetLogSinksResult struct {
	// Result data.
	Data GetLogSinksResultData `json:"data"`
	// Data about the endpoint pagination.
	Pagination GetLogSinksResultPagination `json:"pagination"`
}

func (o *GetLogSinksResult) GetData() GetLogSinksResultData {
	if o == nil {
		return GetLogSinksResultData{}
	}
	return o.Data
}

func (o *GetLogSinksResult) GetPagination() GetLogSinksResultPagination {
	if o == nil {
		return GetLogSinksResultPagination{}
	}
	return o.Pagination
}
