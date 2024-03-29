// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	"net/http"
)

type GetInvoiceDetailsRequest struct {
	AddonID   *string `queryParam:"style=form,explode=true,name=addonId"`
	JobID     *string `queryParam:"style=form,explode=true,name=jobId"`
	ProjectID *string `queryParam:"style=form,explode=true,name=projectId"`
	ServiceID *string `queryParam:"style=form,explode=true,name=serviceId"`
	Timestamp *int64  `queryParam:"style=form,explode=true,name=timestamp"`
}

func (o *GetInvoiceDetailsRequest) GetAddonID() *string {
	if o == nil {
		return nil
	}
	return o.AddonID
}

func (o *GetInvoiceDetailsRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *GetInvoiceDetailsRequest) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *GetInvoiceDetailsRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *GetInvoiceDetailsRequest) GetTimestamp() *int64 {
	if o == nil {
		return nil
	}
	return o.Timestamp
}

type GetInvoiceDetailsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Details about an invoice.
	InvoiceDetailsResult *shared.InvoiceDetailsResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetInvoiceDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInvoiceDetailsResponse) GetInvoiceDetailsResult() *shared.InvoiceDetailsResult {
	if o == nil {
		return nil
	}
	return o.InvoiceDetailsResult
}

func (o *GetInvoiceDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInvoiceDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
