// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/northflank-go/pkg/models/shared"
	"github.com/speakeasy-sdks/northflank-go/pkg/utils"
	"net/http"
)

type AddRegistryRequestBodyType string

const (
	AddRegistryRequestBodyTypeAddRegistryType1 AddRegistryRequestBodyType = "AddRegistryType1"
	AddRegistryRequestBodyTypeAddRegistryType2 AddRegistryRequestBodyType = "AddRegistryType2"
	AddRegistryRequestBodyTypeAddRegistryType3 AddRegistryRequestBodyType = "AddRegistryType3"
)

type AddRegistryRequestBody struct {
	AddRegistryType1 *shared.AddRegistryType1
	AddRegistryType2 *shared.AddRegistryType2
	AddRegistryType3 *shared.AddRegistryType3

	Type AddRegistryRequestBodyType
}

func CreateAddRegistryRequestBodyAddRegistryType1(addRegistryType1 shared.AddRegistryType1) AddRegistryRequestBody {
	typ := AddRegistryRequestBodyTypeAddRegistryType1

	return AddRegistryRequestBody{
		AddRegistryType1: &addRegistryType1,
		Type:             typ,
	}
}

func CreateAddRegistryRequestBodyAddRegistryType2(addRegistryType2 shared.AddRegistryType2) AddRegistryRequestBody {
	typ := AddRegistryRequestBodyTypeAddRegistryType2

	return AddRegistryRequestBody{
		AddRegistryType2: &addRegistryType2,
		Type:             typ,
	}
}

func CreateAddRegistryRequestBodyAddRegistryType3(addRegistryType3 shared.AddRegistryType3) AddRegistryRequestBody {
	typ := AddRegistryRequestBodyTypeAddRegistryType3

	return AddRegistryRequestBody{
		AddRegistryType3: &addRegistryType3,
		Type:             typ,
	}
}

func (u *AddRegistryRequestBody) UnmarshalJSON(data []byte) error {

	addRegistryType3 := new(shared.AddRegistryType3)
	if err := utils.UnmarshalJSON(data, &addRegistryType3, "", true, true); err == nil {
		u.AddRegistryType3 = addRegistryType3
		u.Type = AddRegistryRequestBodyTypeAddRegistryType3
		return nil
	}

	addRegistryType2 := new(shared.AddRegistryType2)
	if err := utils.UnmarshalJSON(data, &addRegistryType2, "", true, true); err == nil {
		u.AddRegistryType2 = addRegistryType2
		u.Type = AddRegistryRequestBodyTypeAddRegistryType2
		return nil
	}

	addRegistryType1 := new(shared.AddRegistryType1)
	if err := utils.UnmarshalJSON(data, &addRegistryType1, "", true, true); err == nil {
		u.AddRegistryType1 = addRegistryType1
		u.Type = AddRegistryRequestBodyTypeAddRegistryType1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AddRegistryRequestBody) MarshalJSON() ([]byte, error) {
	if u.AddRegistryType1 != nil {
		return utils.MarshalJSON(u.AddRegistryType1, "", true)
	}

	if u.AddRegistryType2 != nil {
		return utils.MarshalJSON(u.AddRegistryType2, "", true)
	}

	if u.AddRegistryType3 != nil {
		return utils.MarshalJSON(u.AddRegistryType3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type AddRegistryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Data about the newly created credentials.
	RegistriesResult *shared.RegistriesResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddRegistryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddRegistryResponse) GetRegistriesResult() *shared.RegistriesResult {
	if o == nil {
		return nil
	}
	return o.RegistriesResult
}

func (o *AddRegistryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddRegistryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
