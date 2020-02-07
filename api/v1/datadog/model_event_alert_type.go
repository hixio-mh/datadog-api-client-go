/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// EventAlertType If it is an alert event, set its type between: error, warning, info, and success.
type EventAlertType string

// List of EventAlertType
const (
	EVENTALERTTYPE_ERROR   EventAlertType = "error"
	EVENTALERTTYPE_WARNING EventAlertType = "warning"
	EVENTALERTTYPE_INFO    EventAlertType = "info"
	EVENTALERTTYPE_SUCCESS EventAlertType = "success"
)

// Ptr returns reference to EventAlertType value
func (v EventAlertType) Ptr() *EventAlertType {
	return &v
}

type NullableEventAlertType struct {
	Value        EventAlertType
	ExplicitNull bool
}

func (v NullableEventAlertType) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEventAlertType) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}