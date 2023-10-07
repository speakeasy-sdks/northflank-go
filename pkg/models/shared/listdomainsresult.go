// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
)

// ListDomainsResultDataDomainsStatus - The status of the domain verification.
type ListDomainsResultDataDomainsStatus string

const (
	ListDomainsResultDataDomainsStatusPending  ListDomainsResultDataDomainsStatus = "pending"
	ListDomainsResultDataDomainsStatusVerified ListDomainsResultDataDomainsStatus = "verified"
)

func (e ListDomainsResultDataDomainsStatus) ToPointer() *ListDomainsResultDataDomainsStatus {
	return &e
}

func (e *ListDomainsResultDataDomainsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "verified":
		*e = ListDomainsResultDataDomainsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListDomainsResultDataDomainsStatus: %v", v)
	}
}

// ListDomainsResultDataDomains - Details about a domain.
type ListDomainsResultDataDomains struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The hostname to add to your domain's DNS records as a TXT record to verify the domain.
	Hostname string `json:"hostname"`
	// The domain name.
	Name string `json:"name"`
	// The status of the domain verification.
	Status ListDomainsResultDataDomainsStatus `json:"status"`
	// The token to add as the content of the TXT record to verify the domain.
	Token string `json:"token"`
}

func (l ListDomainsResultDataDomains) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDomainsResultDataDomains) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDomainsResultDataDomains) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ListDomainsResultDataDomains) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *ListDomainsResultDataDomains) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListDomainsResultDataDomains) GetStatus() ListDomainsResultDataDomainsStatus {
	if o == nil {
		return ListDomainsResultDataDomainsStatus("")
	}
	return o.Status
}

func (o *ListDomainsResultDataDomains) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// ListDomainsResultData - Result data.
type ListDomainsResultData struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// A list of domains registered to this account.
	Domains []ListDomainsResultDataDomains `json:"domains"`
}

func (l ListDomainsResultData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDomainsResultData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDomainsResultData) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ListDomainsResultData) GetDomains() []ListDomainsResultDataDomains {
	if o == nil {
		return []ListDomainsResultDataDomains{}
	}
	return o.Domains
}

// ListDomainsResultPagination - Data about the endpoint pagination.
type ListDomainsResultPagination struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The number of results returned by this request.
	Count float32 `json:"count"`
	// The cursor to access the next page of results.
	Cursor *string `json:"cursor,omitempty"`
	// Is there another page of results available?
	HasNextPage bool `json:"hasNextPage"`
}

func (l ListDomainsResultPagination) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDomainsResultPagination) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDomainsResultPagination) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ListDomainsResultPagination) GetCount() float32 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *ListDomainsResultPagination) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListDomainsResultPagination) GetHasNextPage() bool {
	if o == nil {
		return false
	}
	return o.HasNextPage
}

// ListDomainsResult - Response object.
type ListDomainsResult struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Result data.
	Data ListDomainsResultData `json:"data"`
	// Data about the endpoint pagination.
	Pagination ListDomainsResultPagination `json:"pagination"`
}

func (l ListDomainsResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListDomainsResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListDomainsResult) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ListDomainsResult) GetData() ListDomainsResultData {
	if o == nil {
		return ListDomainsResultData{}
	}
	return o.Data
}

func (o *ListDomainsResult) GetPagination() ListDomainsResultPagination {
	if o == nil {
		return ListDomainsResultPagination{}
	}
	return o.Pagination
}
