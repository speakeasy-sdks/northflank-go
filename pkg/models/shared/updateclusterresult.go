// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/northflank-go/v2/pkg/utils"
	"time"
)

// UpdateClusterResultAutoscaling - Auto scaling settings to use for the node pool. Requires that the cloud provider supports this feature.
type UpdateClusterResultAutoscaling struct {
	Enabled *bool  `default:"false" json:"enabled"`
	Max     *int64 `json:"max,omitempty"`
	Min     *int64 `json:"min,omitempty"`
}

func (u UpdateClusterResultAutoscaling) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateClusterResultAutoscaling) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateClusterResultAutoscaling) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpdateClusterResultAutoscaling) GetMax() *int64 {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *UpdateClusterResultAutoscaling) GetMin() *int64 {
	if o == nil {
		return nil
	}
	return o.Min
}

// UpdateClusterResultLabels - Set of label keys and values that can be used to determine scheduling via resource tags.
type UpdateClusterResultLabels struct {
}

type UpdateClusterResultNodePools struct {
	// Auto scaling settings to use for the node pool. Requires that the cloud provider supports this feature.
	Autoscaling *UpdateClusterResultAutoscaling `json:"autoscaling,omitempty"`
	// Zones in which the node pool should be provisioned.
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// Disk size in GB
	DiskSize int64 `json:"diskSize"`
	// The disk type to use.
	DiskType *string `json:"diskType,omitempty"`
	// ID of existing node pool. Must be passed when modifying existing node pools. Not relevant for new node pools
	ID *string `json:"id,omitempty"`
	// Set of label keys and values that can be used to determine scheduling via resource tags.
	Labels *UpdateClusterResultLabels `json:"labels,omitempty"`
	// Number of nodes to the node pool should be provisioned with.
	NodeCount int64 `json:"nodeCount"`
	// Machine type to be used by the node pool.
	NodeType string `json:"nodeType"`
	// Configures node pool with preemptible / spot instances if enabled.
	Preemptible *bool `default:"false" json:"preemptible"`
	// When 'provider' is 'azure', at least one system node pool is required per cluster.
	SystemPool *bool `json:"systemPool,omitempty"`
}

func (u UpdateClusterResultNodePools) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateClusterResultNodePools) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateClusterResultNodePools) GetAutoscaling() *UpdateClusterResultAutoscaling {
	if o == nil {
		return nil
	}
	return o.Autoscaling
}

func (o *UpdateClusterResultNodePools) GetAvailabilityZones() []string {
	if o == nil {
		return nil
	}
	return o.AvailabilityZones
}

func (o *UpdateClusterResultNodePools) GetDiskSize() int64 {
	if o == nil {
		return 0
	}
	return o.DiskSize
}

func (o *UpdateClusterResultNodePools) GetDiskType() *string {
	if o == nil {
		return nil
	}
	return o.DiskType
}

func (o *UpdateClusterResultNodePools) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateClusterResultNodePools) GetLabels() *UpdateClusterResultLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *UpdateClusterResultNodePools) GetNodeCount() int64 {
	if o == nil {
		return 0
	}
	return o.NodeCount
}

func (o *UpdateClusterResultNodePools) GetNodeType() string {
	if o == nil {
		return ""
	}
	return o.NodeType
}

func (o *UpdateClusterResultNodePools) GetPreemptible() *bool {
	if o == nil {
		return nil
	}
	return o.Preemptible
}

func (o *UpdateClusterResultNodePools) GetSystemPool() *bool {
	if o == nil {
		return nil
	}
	return o.SystemPool
}

type UpdateClusterResultState struct {
	State          *string    `json:"state,omitempty"`
	TransitionTime *time.Time `json:"transitionTime,omitempty"`
}

func (u UpdateClusterResultState) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateClusterResultState) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateClusterResultState) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *UpdateClusterResultState) GetTransitionTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.TransitionTime
}

type UpdateClusterResultStatus struct {
	NextUpdateAfter *time.Time                `json:"nextUpdateAfter,omitempty"`
	State           *UpdateClusterResultState `json:"state,omitempty"`
}

func (u UpdateClusterResultStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateClusterResultStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateClusterResultStatus) GetNextUpdateAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.NextUpdateAfter
}

func (o *UpdateClusterResultStatus) GetState() *UpdateClusterResultState {
	if o == nil {
		return nil
	}
	return o.State
}

// UpdateClusterResultData - Result data.
type UpdateClusterResultData struct {
	// The time the cluster was created.
	CreatedAt time.Time `json:"createdAt"`
	// Indicates if provider resource deletion has been requested by the user.
	DeletionRequested bool `json:"deletionRequested"`
	// A short description of the cluster.
	Description *string `json:"description,omitempty"`
	// Identifier for the cluster.
	ID string `json:"id"`
	// ID of the provider integration used by this cluster.
	IntegrationID string `json:"integrationId"`
	// The name of the cluster.
	Name      string                       `json:"name"`
	NodePools UpdateClusterResultNodePools `json:"nodePools"`
	// The cloud provider to which this cluster belongs to.
	Provider *string                    `json:"provider,omitempty"`
	Status   *UpdateClusterResultStatus `json:"status,omitempty"`
}

func (u UpdateClusterResultData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateClusterResultData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateClusterResultData) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *UpdateClusterResultData) GetDeletionRequested() bool {
	if o == nil {
		return false
	}
	return o.DeletionRequested
}

func (o *UpdateClusterResultData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateClusterResultData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateClusterResultData) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *UpdateClusterResultData) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateClusterResultData) GetNodePools() UpdateClusterResultNodePools {
	if o == nil {
		return UpdateClusterResultNodePools{}
	}
	return o.NodePools
}

func (o *UpdateClusterResultData) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *UpdateClusterResultData) GetStatus() *UpdateClusterResultStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// UpdateClusterResult - Response object.
type UpdateClusterResult struct {
	// Result data.
	Data UpdateClusterResultData `json:"data"`
}

func (o *UpdateClusterResult) GetData() UpdateClusterResultData {
	if o == nil {
		return UpdateClusterResultData{}
	}
	return o.Data
}
