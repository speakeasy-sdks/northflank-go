// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
	"net/http"
)

type AddSubDomainRequest struct {
	// Request body
	AddSubDomainRequest shared.AddSubDomainRequest `request:"mediaType=application/json"`
	Domain              string                     `pathParam:"style=simple,explode=false,name=domain"`
}

func (o *AddSubDomainRequest) GetAddSubDomainRequest() shared.AddSubDomainRequest {
	if o == nil {
		return shared.AddSubDomainRequest{}
	}
	return o.AddSubDomainRequest
}

func (o *AddSubDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

type AddSubDomainResponse struct {
	// Details about the newly added subdomain.
	AddSubDomainResult *shared.AddSubDomainResult
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddSubDomainResponse) GetAddSubDomainResult() *shared.AddSubDomainResult {
	if o == nil {
		return nil
	}
	return o.AddSubDomainResult
}

func (o *AddSubDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddSubDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddSubDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
