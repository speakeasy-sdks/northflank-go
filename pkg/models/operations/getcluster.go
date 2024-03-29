// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/northflank-go/v3/pkg/models/shared"
	"net/http"
)

type GetClusterRequest struct {
	ClusterID string `pathParam:"style=simple,explode=false,name=clusterId"`
}

func (o *GetClusterRequest) GetClusterID() string {
	if o == nil {
		return ""
	}
	return o.ClusterID
}

type GetClusterResponse struct {
	// Details about the given cluster.
	ClusterDetailsResult *shared.ClusterDetailsResult
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetClusterResponse) GetClusterDetailsResult() *shared.ClusterDetailsResult {
	if o == nil {
		return nil
	}
	return o.ClusterDetailsResult
}

func (o *GetClusterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetClusterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetClusterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
