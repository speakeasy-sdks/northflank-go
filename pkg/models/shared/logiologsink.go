// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// LogioLogSinkSinkDataRegion - Your Logzio region code
type LogioLogSinkSinkDataRegion string

const (
	LogioLogSinkSinkDataRegionEu LogioLogSinkSinkDataRegion = "eu"
	LogioLogSinkSinkDataRegionUk LogioLogSinkSinkDataRegion = "uk"
	LogioLogSinkSinkDataRegionUs LogioLogSinkSinkDataRegion = "us"
	LogioLogSinkSinkDataRegionCa LogioLogSinkSinkDataRegion = "ca"
	LogioLogSinkSinkDataRegionAu LogioLogSinkSinkDataRegion = "au"
	LogioLogSinkSinkDataRegionNl LogioLogSinkSinkDataRegion = "nl"
	LogioLogSinkSinkDataRegionWa LogioLogSinkSinkDataRegion = "wa"
)

func (e LogioLogSinkSinkDataRegion) ToPointer() *LogioLogSinkSinkDataRegion {
	return &e
}

func (e *LogioLogSinkSinkDataRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "eu":
		fallthrough
	case "uk":
		fallthrough
	case "us":
		fallthrough
	case "ca":
		fallthrough
	case "au":
		fallthrough
	case "nl":
		fallthrough
	case "wa":
		*e = LogioLogSinkSinkDataRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogioLogSinkSinkDataRegion: %v", v)
	}
}

// LogioLogSinkSinkData - Details about the Logz.io log sink.
type LogioLogSinkSinkData struct {
	// Your Logzio region code
	Region LogioLogSinkSinkDataRegion `json:"region"`
	// The Log Shipping Token of the account you want to ship to
	Token string `json:"token"`
}

func (o *LogioLogSinkSinkData) GetRegion() LogioLogSinkSinkDataRegion {
	if o == nil {
		return LogioLogSinkSinkDataRegion("")
	}
	return o.Region
}

func (o *LogioLogSinkSinkData) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// LogioLogSinkSinkType - The type of the log sink.
type LogioLogSinkSinkType string

const (
	LogioLogSinkSinkTypeLogzio LogioLogSinkSinkType = "logzio"
)

func (e LogioLogSinkSinkType) ToPointer() *LogioLogSinkSinkType {
	return &e
}

func (e *LogioLogSinkSinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "logzio":
		*e = LogioLogSinkSinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogioLogSinkSinkType: %v", v)
	}
}

// LogioLogSink - Create a log sink using Logz.io
type LogioLogSink struct {
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
	// Details about the Logz.io log sink.
	SinkData LogioLogSinkSinkData `json:"sinkData"`
	// The type of the log sink.
	SinkType LogioLogSinkSinkType `json:"sinkType"`
	// If `true`, we will do additional parsing on your JSON formatted log lines and your extract custom labels
	UseCustomLabels *bool `json:"useCustomLabels,omitempty"`
}

func (o *LogioLogSink) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LogioLogSink) GetForwardAccessLogs() *bool {
	if o == nil {
		return nil
	}
	return o.ForwardAccessLogs
}

func (o *LogioLogSink) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *LogioLogSink) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *LogioLogSink) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *LogioLogSink) GetSinkData() LogioLogSinkSinkData {
	if o == nil {
		return LogioLogSinkSinkData{}
	}
	return o.SinkData
}

func (o *LogioLogSink) GetSinkType() LogioLogSinkSinkType {
	if o == nil {
		return LogioLogSinkSinkType("")
	}
	return o.SinkType
}

func (o *LogioLogSink) GetUseCustomLabels() *bool {
	if o == nil {
		return nil
	}
	return o.UseCustomLabels
}
