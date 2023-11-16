// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/utils"
)

// BranchesResultData - Result data.
type BranchesResultData struct {
	// The cursor returned from the previous page of results, used to request the next page.
	Cursor *string `json:"cursor,omitempty"`
	// The page number to access.
	Page *int64 `json:"page,omitempty"`
	// The number of results to display per request. Maximum of 100 results per page.
	PerPage *int64 `default:"50" json:"per_page"`
	// If provided, uses the given VCS link to access the repository's data.
	VcsLinkID *string `json:"vcs_link_id,omitempty"`
}

func (b BranchesResultData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BranchesResultData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BranchesResultData) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *BranchesResultData) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *BranchesResultData) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *BranchesResultData) GetVcsLinkID() *string {
	if o == nil {
		return nil
	}
	return o.VcsLinkID
}

// Pagination - Data about the endpoint pagination.
type Pagination struct {
	// The number of results returned by this request.
	Count float32 `json:"count"`
	// The cursor to access the next page of results.
	Cursor *string `json:"cursor,omitempty"`
	// Is there another page of results available?
	HasNextPage bool `json:"hasNextPage"`
}

func (o *Pagination) GetCount() float32 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *Pagination) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *Pagination) GetHasNextPage() bool {
	if o == nil {
		return false
	}
	return o.HasNextPage
}

// BranchesResult - Response object.
type BranchesResult struct {
	// Result data.
	Data BranchesResultData `json:"data"`
	// Data about the endpoint pagination.
	Pagination Pagination `json:"pagination"`
}

func (o *BranchesResult) GetData() BranchesResultData {
	if o == nil {
		return BranchesResultData{}
	}
	return o.Data
}

func (o *BranchesResult) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}
