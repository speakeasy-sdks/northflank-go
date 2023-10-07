// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
)

type CreateDomainRequest struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The domain name to register.
	Domain string `json:"domain"`
}

func (c CreateDomainRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateDomainRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateDomainRequest) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *CreateDomainRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}
