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

// LogsArchiveDefinition The definition of an archive.
type LogsArchiveDefinition struct {
	Attributes *LogsArchiveAttributes `json:"attributes,omitempty"`
	// The archive ID.
	Id *string `json:"id,omitempty"`
	// The type of the resource. The value should always be archives.
	Type string `json:"type"`
}

// NewLogsArchiveDefinition instantiates a new LogsArchiveDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsArchiveDefinition(type_ string) *LogsArchiveDefinition {
	this := LogsArchiveDefinition{}
	this.Type = type_
	return &this
}

// NewLogsArchiveDefinitionWithDefaults instantiates a new LogsArchiveDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsArchiveDefinitionWithDefaults() *LogsArchiveDefinition {
	this := LogsArchiveDefinition{}
	var type_ string = "archives"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *LogsArchiveDefinition) GetAttributes() LogsArchiveAttributes {
	if o == nil || o.Attributes == nil {
		var ret LogsArchiveAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveDefinition) GetAttributesOk() (*LogsArchiveAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *LogsArchiveDefinition) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given LogsArchiveAttributes and assigns it to the Attributes field.
func (o *LogsArchiveDefinition) SetAttributes(v LogsArchiveAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogsArchiveDefinition) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveDefinition) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogsArchiveDefinition) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogsArchiveDefinition) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *LogsArchiveDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsArchiveDefinition) SetType(v string) {
	o.Type = v
}

func (o LogsArchiveDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLogsArchiveDefinition struct {
	value *LogsArchiveDefinition
	isSet bool
}

func (v NullableLogsArchiveDefinition) Get() *LogsArchiveDefinition {
	return v.value
}

func (v *NullableLogsArchiveDefinition) Set(val *LogsArchiveDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveDefinition(val *LogsArchiveDefinition) *NullableLogsArchiveDefinition {
	return &NullableLogsArchiveDefinition{value: val, isSet: true}
}

func (v NullableLogsArchiveDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
