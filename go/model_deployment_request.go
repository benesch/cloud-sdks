/*
 * Materialize Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mzcloud

import (
	"encoding/json"
)

// DeploymentRequest struct for DeploymentRequest
type DeploymentRequest struct {
	Size *string `json:"size,omitempty"`
	MzVersion string `json:"mzVersion"`
}

// NewDeploymentRequest instantiates a new DeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentRequest(mzVersion string) *DeploymentRequest {
	this := DeploymentRequest{}
	this.MzVersion = mzVersion
	return &this
}

// NewDeploymentRequestWithDefaults instantiates a new DeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentRequestWithDefaults() *DeploymentRequest {
	this := DeploymentRequest{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DeploymentRequest) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRequest) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DeploymentRequest) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *DeploymentRequest) SetSize(v string) {
	o.Size = &v
}

// GetMzVersion returns the MzVersion field value
func (o *DeploymentRequest) GetMzVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MzVersion
}

// GetMzVersionOk returns a tuple with the MzVersion field value
// and a boolean to check if the value has been set.
func (o *DeploymentRequest) GetMzVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MzVersion, true
}

// SetMzVersion sets field value
func (o *DeploymentRequest) SetMzVersion(v string) {
	o.MzVersion = v
}

func (o DeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["mzVersion"] = o.MzVersion
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentRequest struct {
	value *DeploymentRequest
	isSet bool
}

func (v NullableDeploymentRequest) Get() *DeploymentRequest {
	return v.value
}

func (v *NullableDeploymentRequest) Set(val *DeploymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentRequest(val *DeploymentRequest) *NullableDeploymentRequest {
	return &NullableDeploymentRequest{value: val, isSet: true}
}

func (v NullableDeploymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


