// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Flags - An object with feature flags to indicate (un)supported features
type Flags struct {
	// Node pool autoscaling support
	Autoscaling *bool `json:"autoscaling,omitempty"`
	// Preemptible/Spot node pool support
	Preemptible *bool `json:"preemptible,omitempty"`
}

func (o *Flags) GetAutoscaling() *bool {
	if o == nil {
		return nil
	}
	return o.Autoscaling
}

func (o *Flags) GetPreemptible() *bool {
	if o == nil {
		return nil
	}
	return o.Preemptible
}

// Providers - A provider object
type Providers struct {
	// An array of supported node disk sizes
	DiskSizes []float32 `json:"diskSizes"`
	// The kubernetes engine used.
	Engine string `json:"engine"`
	// An object with feature flags to indicate (un)supported features
	Flags Flags `json:"flags"`
	// The ID of the provider.
	ID string `json:"id"`
	// An array of available kubernetes versions
	KubernetesVersions []string `json:"kubernetesVersions"`
	// The name of the provider.
	Name string `json:"name"`
	// An array of supported node types
	NodeTypes []string `json:"nodeTypes"`
	// An array of available regions
	Regions []string `json:"regions"`
}

func (o *Providers) GetDiskSizes() []float32 {
	if o == nil {
		return []float32{}
	}
	return o.DiskSizes
}

func (o *Providers) GetEngine() string {
	if o == nil {
		return ""
	}
	return o.Engine
}

func (o *Providers) GetFlags() Flags {
	if o == nil {
		return Flags{}
	}
	return o.Flags
}

func (o *Providers) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Providers) GetKubernetesVersions() []string {
	if o == nil {
		return []string{}
	}
	return o.KubernetesVersions
}

func (o *Providers) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Providers) GetNodeTypes() []string {
	if o == nil {
		return []string{}
	}
	return o.NodeTypes
}

func (o *Providers) GetRegions() []string {
	if o == nil {
		return []string{}
	}
	return o.Regions
}

// CloudProvidersResultData - Result data.
type CloudProvidersResultData struct {
	// An array of supported cloud providers
	Providers []Providers `json:"providers"`
}

func (o *CloudProvidersResultData) GetProviders() []Providers {
	if o == nil {
		return []Providers{}
	}
	return o.Providers
}

// CloudProvidersResult - Response object.
type CloudProvidersResult struct {
	// Result data.
	Data CloudProvidersResultData `json:"data"`
}

func (o *CloudProvidersResult) GetData() CloudProvidersResultData {
	if o == nil {
		return CloudProvidersResultData{}
	}
	return o.Data
}
