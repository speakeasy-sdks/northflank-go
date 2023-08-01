// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"net/http"
)

type GetIntegrationRequest struct {
	IntegrationID string `pathParam:"style=simple,explode=false,name=integrationId"`
}

func (o *GetIntegrationRequest) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

type GetIntegrationResponse struct {
	ContentType string
	// Details about the given integration.
	GetIntegrationResult *shared.GetIntegrationResult
	StatusCode           int
	RawResponse          *http.Response
}

func (o *GetIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIntegrationResponse) GetGetIntegrationResult() *shared.GetIntegrationResult {
	if o == nil {
		return nil
	}
	return o.GetIntegrationResult
}

func (o *GetIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
