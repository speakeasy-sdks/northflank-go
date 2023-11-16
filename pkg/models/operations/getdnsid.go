// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/models/shared"
	"net/http"
)

type GetDNSIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Data about the DNS ID.
	DNSIDResult *shared.DNSIDResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDNSIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDNSIDResponse) GetDNSIDResult() *shared.DNSIDResult {
	if o == nil {
		return nil
	}
	return o.DNSIDResult
}

func (o *GetDNSIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDNSIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
