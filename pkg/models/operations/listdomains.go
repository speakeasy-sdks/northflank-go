// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"net/http"
)

type ListDomainsRequest struct {
	Cursor  *string `queryParam:"style=form,explode=true,name=cursor"`
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListDomainsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListDomainsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListDomainsRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListDomainsResponse struct {
	ContentType string
	// A list of domains registered to this account.
	ListDomainsResult *shared.ListDomainsResult
	StatusCode        int
	RawResponse       *http.Response
}

func (o *ListDomainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDomainsResponse) GetListDomainsResult() *shared.ListDomainsResult {
	if o == nil {
		return nil
	}
	return o.ListDomainsResult
}

func (o *ListDomainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDomainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
