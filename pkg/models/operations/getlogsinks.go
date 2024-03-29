// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/utils"
	"net/http"
)

type GetLogSinksRequest struct {
	Cursor  *string `queryParam:"style=form,explode=true,name=cursor"`
	PerPage *int64  `default:"50" queryParam:"style=form,explode=true,name=per_page"`
}

func (g GetLogSinksRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLogSinksRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLogSinksRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetLogSinksRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetLogSinksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The list of log sinks.
	GetLogSinksResult *shared.GetLogSinksResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetLogSinksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLogSinksResponse) GetGetLogSinksResult() *shared.GetLogSinksResult {
	if o == nil {
		return nil
	}
	return o.GetLogSinksResult
}

func (o *GetLogSinksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLogSinksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
