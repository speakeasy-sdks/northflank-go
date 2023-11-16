// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
	"net/http"
)

type GetDomainRequest struct {
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
}

func (o *GetDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

type GetDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Details about the given domain.
	GetDomainResult *shared.GetDomainResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDomainResponse) GetGetDomainResult() *shared.GetDomainResult {
	if o == nil {
		return nil
	}
	return o.GetDomainResult
}

func (o *GetDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
