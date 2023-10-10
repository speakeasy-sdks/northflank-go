// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
)

// DatadogLogSinkSinkDataRegion - The Datadog region.
type DatadogLogSinkSinkDataRegion string

const (
	DatadogLogSinkSinkDataRegionEu  DatadogLogSinkSinkDataRegion = "eu"
	DatadogLogSinkSinkDataRegionUs  DatadogLogSinkSinkDataRegion = "us"
	DatadogLogSinkSinkDataRegionUs3 DatadogLogSinkSinkDataRegion = "us3"
	DatadogLogSinkSinkDataRegionUs5 DatadogLogSinkSinkDataRegion = "us5"
)

func (e DatadogLogSinkSinkDataRegion) ToPointer() *DatadogLogSinkSinkDataRegion {
	return &e
}

func (e *DatadogLogSinkSinkDataRegion) UnmarshalJSON(data []byte) error {
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
		*e = DatadogLogSinkSinkDataRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatadogLogSinkSinkDataRegion: %v", v)
	}
}

// DatadogLogSinkSinkData - Details about the Datadog log sink.
type DatadogLogSinkSinkData struct {
	// The Datadog API key.
	DefaultAPIKey string `json:"default_api_key"`
	// The Datadog region.
	Region DatadogLogSinkSinkDataRegion `json:"region"`
}

func (o *DatadogLogSinkSinkData) GetDefaultAPIKey() string {
	if o == nil {
		return ""
	}
	return o.DefaultAPIKey
}

func (o *DatadogLogSinkSinkData) GetRegion() DatadogLogSinkSinkDataRegion {
	if o == nil {
		return DatadogLogSinkSinkDataRegion("")
	}
	return o.Region
}

// DatadogLogSinkSinkType - The type of the log sink.
type DatadogLogSinkSinkType string

const (
	DatadogLogSinkSinkTypeDatadogLogs DatadogLogSinkSinkType = "datadog_logs"
)

func (e DatadogLogSinkSinkType) ToPointer() *DatadogLogSinkSinkType {
	return &e
}

func (e *DatadogLogSinkSinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "datadog_logs":
		*e = DatadogLogSinkSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DatadogLogSinkSinkType: %v", v)
	}
}

// DatadogLogSink - Create a log sink using Datadog
type DatadogLogSink struct {
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
	// Details about the Datadog log sink.
	SinkData DatadogLogSinkSinkData `json:"sinkData"`
	// The type of the log sink.
	SinkType DatadogLogSinkSinkType `json:"sinkType"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `default:"false" json:"useCustomLabels"`
}

func (d DatadogLogSink) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DatadogLogSink) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DatadogLogSink) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DatadogLogSink) GetForwardAccessLogs() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardAccessLogs
}

func (o *DatadogLogSink) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DatadogLogSink) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *DatadogLogSink) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *DatadogLogSink) GetSinkData() DatadogLogSinkSinkData {
	if o == nil {
		return DatadogLogSinkSinkData{}
	}
	return o.SinkData
}

func (o *DatadogLogSink) GetSinkType() DatadogLogSinkSinkType {
	if o == nil {
		return DatadogLogSinkSinkType("")
	}
	return o.SinkType
}

func (o *DatadogLogSink) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}
