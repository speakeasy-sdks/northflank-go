// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Cloudfront struct {
	Enabled bool `json:"enabled"`
}

func (o *Cloudfront) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

// Cdn - Optional CDN configuration. Currently only available for select users.
type Cdn struct {
	Cloudfront *Cloudfront `json:"cloudfront,omitempty"`
}

func (o *Cdn) GetCloudfront() *Cloudfront {
	if o == nil {
		return nil
	}
	return o.Cloudfront
}

type AddSubDomainRequest struct {
	// Optional CDN configuration. Currently only available for select users.
	Cdn *Cdn `json:"cdn,omitempty"`
	// A subdomain to be added.
	Subdomain string `json:"subdomain"`
}

func (o *AddSubDomainRequest) GetCdn() *Cdn {
	if o == nil {
		return nil
	}
	return o.Cdn
}

func (o *AddSubDomainRequest) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
