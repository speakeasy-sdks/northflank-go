// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
	"net/http"
)

type CreateIntegrationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Details about the created integration.
	CreateIntegrationResult *shared.CreateIntegrationResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIntegrationResponse) GetCreateIntegrationResult() *shared.CreateIntegrationResult {
	if o == nil {
		return nil
	}
	return o.CreateIntegrationResult
}

func (o *CreateIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
