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

// LogsArithmeticProcessorType Type of logs arithmetic processor.
type LogsArithmeticProcessorType string

// List of LogsArithmeticProcessorType
const (
	LOGSARITHMETICPROCESSORTYPE_ARITHMETIC_PROCESSOR LogsArithmeticProcessorType = "arithmetic-processor"
)

func (v *LogsArithmeticProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsArithmeticProcessorType(value)
	for _, existing := range []LogsArithmeticProcessorType{"arithmetic-processor"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsArithmeticProcessorType", *v)
}

// Ptr returns reference to LogsArithmeticProcessorType value
func (v LogsArithmeticProcessorType) Ptr() *LogsArithmeticProcessorType {
	return &v
}

type NullableLogsArithmeticProcessorType struct {
	value *LogsArithmeticProcessorType
	isSet bool
}

func (v NullableLogsArithmeticProcessorType) Get() *LogsArithmeticProcessorType {
	return v.value
}

func (v *NullableLogsArithmeticProcessorType) Set(val *LogsArithmeticProcessorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArithmeticProcessorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArithmeticProcessorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArithmeticProcessorType(val *LogsArithmeticProcessorType) *NullableLogsArithmeticProcessorType {
	return &NullableLogsArithmeticProcessorType{value: val, isSet: true}
}

func (v NullableLogsArithmeticProcessorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArithmeticProcessorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
