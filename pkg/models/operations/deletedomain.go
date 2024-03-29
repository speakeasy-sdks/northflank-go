// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	"net/http"
)

type DeleteDomainRequest struct {
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
}

func (o *DeleteDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

type DeleteDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The operation was performed successfully.
	SuccessResult *shared.SuccessResult
}

func (o *DeleteDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteDomainResponse) GetSuccessResult() *shared.SuccessResult {
	if o == nil {
		return nil
	}
	return o.SuccessResult
}
