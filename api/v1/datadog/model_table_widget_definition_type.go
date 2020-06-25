/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// TableWidgetDefinitionType Type of the table widget.
type TableWidgetDefinitionType string

// List of TableWidgetDefinitionType
const (
	TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE TableWidgetDefinitionType = "query_table"
)

func (v *TableWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TableWidgetDefinitionType(value)
	for _, existing := range []TableWidgetDefinitionType{"query_table"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TableWidgetDefinitionType", *v)
}

// Ptr returns reference to TableWidgetDefinitionType value
func (v TableWidgetDefinitionType) Ptr() *TableWidgetDefinitionType {
	return &v
}

type NullableTableWidgetDefinitionType struct {
	value *TableWidgetDefinitionType
	isSet bool
}

func (v NullableTableWidgetDefinitionType) Get() *TableWidgetDefinitionType {
	return v.value
}

func (v *NullableTableWidgetDefinitionType) Set(val *TableWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableTableWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableTableWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableWidgetDefinitionType(val *TableWidgetDefinitionType) *NullableTableWidgetDefinitionType {
	return &NullableTableWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableTableWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
