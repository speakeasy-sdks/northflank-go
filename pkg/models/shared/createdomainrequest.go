// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateDomainRequest struct {
	// The domain name to register.
	Domain string `json:"domain"`
}

func (o *CreateDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}
