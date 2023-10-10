// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
)

// AddRegistryType3Auths - The `auths` data extracted from your Docker config file.
type AddRegistryType3Auths struct {
}

// AddRegistryType3Provider - The registry provider associated with this set of credentials.
type AddRegistryType3Provider string

const (
	AddRegistryType3ProviderDockerhub AddRegistryType3Provider = "dockerhub"
	AddRegistryType3ProviderGcr       AddRegistryType3Provider = "gcr"
	AddRegistryType3ProviderGcrEu     AddRegistryType3Provider = "gcr-eu"
	AddRegistryType3ProviderGcrUs     AddRegistryType3Provider = "gcr-us"
	AddRegistryType3ProviderGitlab    AddRegistryType3Provider = "gitlab"
	AddRegistryType3ProviderGithub    AddRegistryType3Provider = "github"
	AddRegistryType3ProviderCustom    AddRegistryType3Provider = "custom"
)

func (e AddRegistryType3Provider) ToPointer() *AddRegistryType3Provider {
	return &e
}

func (e *AddRegistryType3Provider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dockerhub":
		fallthrough
	case "gcr":
		fallthrough
	case "gcr-eu":
		fallthrough
	case "gcr-us":
		fallthrough
	case "gitlab":
		fallthrough
	case "github":
		fallthrough
	case "custom":
		*e = AddRegistryType3Provider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddRegistryType3Provider: %v", v)
	}
}

// AddRegistryType3Restrictions - Data about whether the credentials are restricted to certain projects.
type AddRegistryType3Restrictions struct {
	// An array of projects the credentials are restricted to, if applicable.
	Projects []string `json:"projects,omitempty"`
	// Whether the credentials are restricted to specific projects.
	Restricted *bool `default:"false" json:"restricted"`
}

func (a AddRegistryType3Restrictions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AddRegistryType3Restrictions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AddRegistryType3Restrictions) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *AddRegistryType3Restrictions) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

// AddRegistryType3 - Validate with a docker config file.
type AddRegistryType3 struct {
	// The `auths` data extracted from your Docker config file.
	Auths AddRegistryType3Auths `json:"auths"`
	// Description of the credentials.
	Description string `json:"description"`
	// Name of the credentials.
	Name string `json:"name"`
	// The registry provider associated with this set of credentials.
	Provider AddRegistryType3Provider `json:"provider"`
	// Data about whether the credentials are restricted to certain projects.
	Restrictions *AddRegistryType3Restrictions `json:"restrictions,omitempty"`
}

func (o *AddRegistryType3) GetAuths() AddRegistryType3Auths {
	if o == nil {
		return AddRegistryType3Auths{}
	}
	return o.Auths
}

func (o *AddRegistryType3) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *AddRegistryType3) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddRegistryType3) GetProvider() AddRegistryType3Provider {
	if o == nil {
		return AddRegistryType3Provider("")
	}
	return o.Provider
}

func (o *AddRegistryType3) GetRestrictions() *AddRegistryType3Restrictions {
	if o == nil {
		return nil
	}
	return o.Restrictions
}
