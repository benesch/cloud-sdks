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
	Size *SizeEnum `json:"size,omitempty"`
	StorageMb *int32 `json:"storageMb,omitempty"`
	MaterializedExtraArgs *[]string `json:"materializedExtraArgs,omitempty"`
	MzVersion *string `json:"mzVersion,omitempty"`
}

// NewDeploymentRequest instantiates a new DeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentRequest() *DeploymentRequest {
	this := DeploymentRequest{}
	var storageMb int32 = 100
	this.StorageMb = &storageMb
	return &this
}

// NewDeploymentRequestWithDefaults instantiates a new DeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentRequestWithDefaults() *DeploymentRequest {
	this := DeploymentRequest{}
	var storageMb int32 = 100
	this.StorageMb = &storageMb
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DeploymentRequest) GetSize() SizeEnum {
	if o == nil || o.Size == nil {
		var ret SizeEnum
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRequest) GetSizeOk() (*SizeEnum, bool) {
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

// SetSize gets a reference to the given SizeEnum and assigns it to the Size field.
func (o *DeploymentRequest) SetSize(v SizeEnum) {
	o.Size = &v
}

// GetStorageMb returns the StorageMb field value if set, zero value otherwise.
func (o *DeploymentRequest) GetStorageMb() int32 {
	if o == nil || o.StorageMb == nil {
		var ret int32
		return ret
	}
	return *o.StorageMb
}

// GetStorageMbOk returns a tuple with the StorageMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRequest) GetStorageMbOk() (*int32, bool) {
	if o == nil || o.StorageMb == nil {
		return nil, false
	}
	return o.StorageMb, true
}

// HasStorageMb returns a boolean if a field has been set.
func (o *DeploymentRequest) HasStorageMb() bool {
	if o != nil && o.StorageMb != nil {
		return true
	}

	return false
}

// SetStorageMb gets a reference to the given int32 and assigns it to the StorageMb field.
func (o *DeploymentRequest) SetStorageMb(v int32) {
	o.StorageMb = &v
}

// GetMaterializedExtraArgs returns the MaterializedExtraArgs field value if set, zero value otherwise.
func (o *DeploymentRequest) GetMaterializedExtraArgs() []string {
	if o == nil || o.MaterializedExtraArgs == nil {
		var ret []string
		return ret
	}
	return *o.MaterializedExtraArgs
}

// GetMaterializedExtraArgsOk returns a tuple with the MaterializedExtraArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRequest) GetMaterializedExtraArgsOk() (*[]string, bool) {
	if o == nil || o.MaterializedExtraArgs == nil {
		return nil, false
	}
	return o.MaterializedExtraArgs, true
}

// HasMaterializedExtraArgs returns a boolean if a field has been set.
func (o *DeploymentRequest) HasMaterializedExtraArgs() bool {
	if o != nil && o.MaterializedExtraArgs != nil {
		return true
	}

	return false
}

// SetMaterializedExtraArgs gets a reference to the given []string and assigns it to the MaterializedExtraArgs field.
func (o *DeploymentRequest) SetMaterializedExtraArgs(v []string) {
	o.MaterializedExtraArgs = &v
}

// GetMzVersion returns the MzVersion field value if set, zero value otherwise.
func (o *DeploymentRequest) GetMzVersion() string {
	if o == nil || o.MzVersion == nil {
		var ret string
		return ret
	}
	return *o.MzVersion
}

// GetMzVersionOk returns a tuple with the MzVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRequest) GetMzVersionOk() (*string, bool) {
	if o == nil || o.MzVersion == nil {
		return nil, false
	}
	return o.MzVersion, true
}

// HasMzVersion returns a boolean if a field has been set.
func (o *DeploymentRequest) HasMzVersion() bool {
	if o != nil && o.MzVersion != nil {
		return true
	}

	return false
}

// SetMzVersion gets a reference to the given string and assigns it to the MzVersion field.
func (o *DeploymentRequest) SetMzVersion(v string) {
	o.MzVersion = &v
}

func (o DeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.StorageMb != nil {
		toSerialize["storageMb"] = o.StorageMb
	}
	if o.MaterializedExtraArgs != nil {
		toSerialize["materializedExtraArgs"] = o.MaterializedExtraArgs
	}
	if o.MzVersion != nil {
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


