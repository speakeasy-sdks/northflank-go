// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/utils"
	"time"
)

// Integrations - An integration object.
type Integrations struct {
	// The time the integration was created.
	CreatedAt time.Time `json:"createdAt"`
	// A short description of the integration.
	Description *string `json:"description,omitempty"`
	// Identifier for the integration.
	ID string `json:"id"`
	// The name of the integration.
	Name string `json:"name"`
	// The cloud provider to which this integration belongs to.
	Provider *string `json:"provider,omitempty"`
}

func (i Integrations) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Integrations) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Integrations) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Integrations) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Integrations) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Integrations) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Integrations) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

// ListIntegrationsResultData - Result data.
type ListIntegrationsResultData struct {
	// An array of integrations.
	Integrations []Integrations `json:"integrations"`
}

func (o *ListIntegrationsResultData) GetIntegrations() []Integrations {
	if o == nil {
		return []Integrations{}
	}
	return o.Integrations
}

// ListIntegrationsResultPagination - Data about the endpoint pagination.
type ListIntegrationsResultPagination struct {
	// The number of results returned by this request.
	Count float32 `json:"count"`
	// The cursor to access the next page of results.
	Cursor *string `json:"cursor,omitempty"`
	// Is there another page of results available?
	HasNextPage bool `json:"hasNextPage"`
}

func (o *ListIntegrationsResultPagination) GetCount() float32 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *ListIntegrationsResultPagination) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListIntegrationsResultPagination) GetHasNextPage() bool {
	if o == nil {
		return false
	}
	return o.HasNextPage
}

// ListIntegrationsResult - Response object.
type ListIntegrationsResult struct {
	// Result data.
	Data ListIntegrationsResultData `json:"data"`
	// Data about the endpoint pagination.
	Pagination ListIntegrationsResultPagination `json:"pagination"`
}

func (o *ListIntegrationsResult) GetData() ListIntegrationsResultData {
	if o == nil {
		return ListIntegrationsResultData{}
	}
	return o.Data
}

func (o *ListIntegrationsResult) GetPagination() ListIntegrationsResultPagination {
	if o == nil {
		return ListIntegrationsResultPagination{}
	}
	return o.Pagination
}
