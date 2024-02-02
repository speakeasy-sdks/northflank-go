// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/utils"
	"time"
)

// GetIntegrationResultCredentials - Cloud provider credential input, required fields dependent on which provider is chosen.
type GetIntegrationResultCredentials struct {
	// AWS access key.
	AccessKey *string `json:"accessKey,omitempty"`
	// DO API key.
	APIKey *string `json:"apiKey,omitempty"`
	// Contents of a GCP key file.
	KeyfileJSON *string `json:"keyfileJson,omitempty"`
	// AWS secret key.
	SecretKey *string `json:"secretKey,omitempty"`
}

func (o *GetIntegrationResultCredentials) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *GetIntegrationResultCredentials) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *GetIntegrationResultCredentials) GetKeyfileJSON() *string {
	if o == nil {
		return nil
	}
	return o.KeyfileJSON
}

func (o *GetIntegrationResultCredentials) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

// GetIntegrationResultData - Result data.
type GetIntegrationResultData struct {
	// The time the integration was created.
	CreatedAt time.Time `json:"createdAt"`
	// Cloud provider credential input, required fields dependent on which provider is chosen.
	Credentials GetIntegrationResultCredentials `json:"credentials"`
	// A short description of the integration.
	Description *string `json:"description,omitempty"`
	// Identifier for the integration.
	ID string `json:"id"`
	// The name of the integration.
	Name string `json:"name"`
}

func (g GetIntegrationResultData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIntegrationResultData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetIntegrationResultData) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GetIntegrationResultData) GetCredentials() GetIntegrationResultCredentials {
	if o == nil {
		return GetIntegrationResultCredentials{}
	}
	return o.Credentials
}

func (o *GetIntegrationResultData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GetIntegrationResultData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetIntegrationResultData) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// GetIntegrationResult - Response object.
type GetIntegrationResult struct {
	// Result data.
	Data GetIntegrationResultData `json:"data"`
}

func (o *GetIntegrationResult) GetData() GetIntegrationResultData {
	if o == nil {
		return GetIntegrationResultData{}
	}
	return o.Data
}
