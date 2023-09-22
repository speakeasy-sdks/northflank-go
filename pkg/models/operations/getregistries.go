// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
	"net/http"
)

type GetRegistriesRequest struct {
	Cursor  *string `queryParam:"style=form,explode=true,name=cursor"`
	PerPage *int64  `default:"50" queryParam:"style=form,explode=true,name=per_page"`
}

func (g GetRegistriesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRegistriesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRegistriesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetRegistriesRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetRegistriesResponse struct {
	ContentType string
	// A list of registry credentials saved to this account.
	RegistriesResult *shared.RegistriesResult
	StatusCode       int
	RawResponse      *http.Response
}

func (o *GetRegistriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRegistriesResponse) GetRegistriesResult() *shared.RegistriesResult {
	if o == nil {
		return nil
	}
	return o.RegistriesResult
}

func (o *GetRegistriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRegistriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
