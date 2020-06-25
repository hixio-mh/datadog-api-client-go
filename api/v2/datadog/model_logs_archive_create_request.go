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

// LogsArchiveCreateRequest The logs archive.
type LogsArchiveCreateRequest struct {
	Data *LogsArchiveCreateRequestDefinition `json:"data,omitempty"`
}

// NewLogsArchiveCreateRequest instantiates a new LogsArchiveCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsArchiveCreateRequest() *LogsArchiveCreateRequest {
	this := LogsArchiveCreateRequest{}
	return &this
}

// NewLogsArchiveCreateRequestWithDefaults instantiates a new LogsArchiveCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsArchiveCreateRequestWithDefaults() *LogsArchiveCreateRequest {
	this := LogsArchiveCreateRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequest) GetData() LogsArchiveCreateRequestDefinition {
	if o == nil || o.Data == nil {
		var ret LogsArchiveCreateRequestDefinition
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequest) GetDataOk() (*LogsArchiveCreateRequestDefinition, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LogsArchiveCreateRequestDefinition and assigns it to the Data field.
func (o *LogsArchiveCreateRequest) SetData(v LogsArchiveCreateRequestDefinition) {
	o.Data = &v
}

func (o LogsArchiveCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLogsArchiveCreateRequest struct {
	value *LogsArchiveCreateRequest
	isSet bool
}

func (v NullableLogsArchiveCreateRequest) Get() *LogsArchiveCreateRequest {
	return v.value
}

func (v *NullableLogsArchiveCreateRequest) Set(val *LogsArchiveCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveCreateRequest(val *LogsArchiveCreateRequest) *NullableLogsArchiveCreateRequest {
	return &NullableLogsArchiveCreateRequest{value: val, isSet: true}
}

func (v NullableLogsArchiveCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
