// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	"net/http"
)

type AssignSubDomainRequest struct {
	// Request body
	AssignSubDomainRequest shared.AssignSubDomainRequest `request:"mediaType=application/json"`
	Domain                 string                        `pathParam:"style=simple,explode=false,name=domain"`
	Subdomain              string                        `pathParam:"style=simple,explode=false,name=subdomain"`
}

func (o *AssignSubDomainRequest) GetAssignSubDomainRequest() shared.AssignSubDomainRequest {
	if o == nil {
		return shared.AssignSubDomainRequest{}
	}
	return o.AssignSubDomainRequest
}

func (o *AssignSubDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *AssignSubDomainRequest) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}

type AssignSubDomainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The operation was performed successfully.
	SuccessResult *shared.SuccessResult
}

func (o *AssignSubDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AssignSubDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AssignSubDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AssignSubDomainResponse) GetSuccessResult() *shared.SuccessResult {
	if o == nil {
		return nil
	}
	return o.SuccessResult
}
