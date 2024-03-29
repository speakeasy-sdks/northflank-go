// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	"net/http"
)

type PauseLogSinkRequest struct {
	LogSinkID string `pathParam:"style=simple,explode=false,name=logSinkId"`
}

func (o *PauseLogSinkRequest) GetLogSinkID() string {
	if o == nil {
		return ""
	}
	return o.LogSinkID
}

type PauseLogSinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The operation was performed successfully.
	SuccessResult *shared.SuccessResult
}

func (o *PauseLogSinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PauseLogSinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PauseLogSinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PauseLogSinkResponse) GetSuccessResult() *shared.SuccessResult {
	if o == nil {
		return nil
	}
	return o.SuccessResult
}
