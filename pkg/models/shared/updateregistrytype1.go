// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateRegistryType1Restrictions - Data about whether the credentials are restricted to certain projects.
type UpdateRegistryType1Restrictions struct {
	// An array of projects the credentials are restricted to, if applicable.
	Projects []string `json:"projects,omitempty"`
	// Whether the credentials are restricted to specific projects.
	Restricted bool `json:"restricted"`
}

func (o *UpdateRegistryType1Restrictions) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *UpdateRegistryType1Restrictions) GetRestricted() bool {
	if o == nil {
		return false
	}
	return o.Restricted
}

// UpdateRegistryType1 - Don't update the credentials.
type UpdateRegistryType1 struct {
	// Description of the credentials.
	Description *string `json:"description,omitempty"`
	// Data about whether the credentials are restricted to certain projects.
	Restrictions *UpdateRegistryType1Restrictions `json:"restrictions,omitempty"`
}

func (o *UpdateRegistryType1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateRegistryType1) GetRestrictions() *UpdateRegistryType1Restrictions {
	if o == nil {
		return nil
	}
	return o.Restrictions
}
