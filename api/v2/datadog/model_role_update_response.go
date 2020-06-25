/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// RoleUpdateResponse Response containing information about an updated role.
type RoleUpdateResponse struct {
	Data *RoleUpdateResponseData `json:"data,omitempty"`
}

// NewRoleUpdateResponse instantiates a new RoleUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdateResponse() *RoleUpdateResponse {
	this := RoleUpdateResponse{}
	return &this
}

// NewRoleUpdateResponseWithDefaults instantiates a new RoleUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateResponseWithDefaults() *RoleUpdateResponse {
	this := RoleUpdateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RoleUpdateResponse) GetData() RoleUpdateResponseData {
	if o == nil || o.Data == nil {
		var ret RoleUpdateResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdateResponse) GetDataOk() (*RoleUpdateResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RoleUpdateResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RoleUpdateResponseData and assigns it to the Data field.
func (o *RoleUpdateResponse) SetData(v RoleUpdateResponseData) {
	o.Data = &v
}

func (o RoleUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRoleUpdateResponse struct {
	value *RoleUpdateResponse
	isSet bool
}

func (v NullableRoleUpdateResponse) Get() *RoleUpdateResponse {
	return v.value
}

func (v *NullableRoleUpdateResponse) Set(val *RoleUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdateResponse(val *RoleUpdateResponse) *NullableRoleUpdateResponse {
	return &NullableRoleUpdateResponse{value: val, isSet: true}
}

func (v NullableRoleUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
