// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetDomainResultDataStatus - The status of the domain verification.
type GetDomainResultDataStatus string

const (
	GetDomainResultDataStatusPending  GetDomainResultDataStatus = "pending"
	GetDomainResultDataStatusVerified GetDomainResultDataStatus = "verified"
)

func (e GetDomainResultDataStatus) ToPointer() *GetDomainResultDataStatus {
	return &e
}

func (e *GetDomainResultDataStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "verified":
		*e = GetDomainResultDataStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDomainResultDataStatus: %v", v)
	}
}

// GetDomainResultDataSubdomains - Details about a subdomain.
type GetDomainResultDataSubdomains struct {
	// The full domain including the subdomain.
	FullName string `json:"fullName"`
	// The subdomain added, or -default for the empty subdomain.
	Name string `json:"name"`
}

func (o *GetDomainResultDataSubdomains) GetFullName() string {
	if o == nil {
		return ""
	}
	return o.FullName
}

func (o *GetDomainResultDataSubdomains) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// GetDomainResultData - Result data.
type GetDomainResultData struct {
	// The hostname to add to your domain's DNS records as a TXT record to verify the domain.
	Hostname string `json:"hostname"`
	// The domain name.
	Name string `json:"name"`
	// The status of the domain verification.
	Status GetDomainResultDataStatus `json:"status"`
	// A list of subdomains added to this domain.
	Subdomains []GetDomainResultDataSubdomains `json:"subdomains"`
	// The token to add as the content of the TXT record to verify the domain.
	Token string `json:"token"`
}

func (o *GetDomainResultData) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *GetDomainResultData) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetDomainResultData) GetStatus() GetDomainResultDataStatus {
	if o == nil {
		return GetDomainResultDataStatus("")
	}
	return o.Status
}

func (o *GetDomainResultData) GetSubdomains() []GetDomainResultDataSubdomains {
	if o == nil {
		return []GetDomainResultDataSubdomains{}
	}
	return o.Subdomains
}

func (o *GetDomainResultData) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// GetDomainResult - Response object.
type GetDomainResult struct {
	// Result data.
	Data GetDomainResultData `json:"data"`
}

func (o *GetDomainResult) GetData() GetDomainResultData {
	if o == nil {
		return GetDomainResultData{}
	}
	return o.Data
}
