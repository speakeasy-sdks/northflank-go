// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"net/http"
)

type GetSubDomainRequest struct {
	Domain    string `pathParam:"style=simple,explode=false,name=domain"`
	Subdomain string `pathParam:"style=simple,explode=false,name=subdomain"`
}

func (o *GetSubDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *GetSubDomainRequest) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}

type GetSubDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Details about the subdomain.
	GetSubDomainResult *shared.GetSubDomainResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSubDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSubDomainResponse) GetGetSubDomainResult() *shared.GetSubDomainResult {
	if o == nil {
		return nil
	}
	return o.GetSubDomainResult
}

func (o *GetSubDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSubDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
