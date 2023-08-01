// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CDNRequest - Request body
type CDNRequest struct {
	// Provider for which the CDN on the subdomain should be disabled.
	Provider string `json:"provider"`
}

func (o *CDNRequest) GetProvider() string {
	if o == nil {
		return ""
	}
	return o.Provider
}
