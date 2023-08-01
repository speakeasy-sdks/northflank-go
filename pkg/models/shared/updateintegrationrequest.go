// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateIntegrationRequestCredentials - Cloud provider credential input, required fields dependent on which provider is chosen.
type UpdateIntegrationRequestCredentials struct {
	// AWS access key.
	AccessKey *string `json:"accessKey,omitempty"`
	// DO API key.
	APIKey *string `json:"apiKey,omitempty"`
	// Contents of a GCP key file.
	KeyfileJSON *string `json:"keyfileJson,omitempty"`
	// AWS secret key.
	SecretKey *string `json:"secretKey,omitempty"`
}

func (o *UpdateIntegrationRequestCredentials) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *UpdateIntegrationRequestCredentials) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *UpdateIntegrationRequestCredentials) GetKeyfileJSON() *string {
	if o == nil {
		return nil
	}
	return o.KeyfileJSON
}

func (o *UpdateIntegrationRequestCredentials) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

// UpdateIntegrationRequest - Request body
type UpdateIntegrationRequest struct {
	// Cloud provider credential input, required fields dependent on which provider is chosen.
	Credentials UpdateIntegrationRequestCredentials `json:"credentials"`
	// The description of the integration.
	Description *string `json:"description,omitempty"`
}

func (o *UpdateIntegrationRequest) GetCredentials() UpdateIntegrationRequestCredentials {
	if o == nil {
		return UpdateIntegrationRequestCredentials{}
	}
	return o.Credentials
}

func (o *UpdateIntegrationRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
