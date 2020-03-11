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

// WidgetLayout the layout for a widget on a free dashboard
type WidgetLayout struct {
	// The height of the widget
	Height int64 `json:"height"`
	// The width of the widget
	Width int64 `json:"width"`
	// The position of the widget on the X axis
	X int64 `json:"x"`
	// The position of the widget on the Y axis
	Y int64 `json:"y"`
}

// NewWidgetLayout instantiates a new WidgetLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetLayout(height int64, width int64, x int64, y int64) *WidgetLayout {
	this := WidgetLayout{}
	this.Height = height
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewWidgetLayoutWithDefaults instantiates a new WidgetLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetLayoutWithDefaults() *WidgetLayout {
	this := WidgetLayout{}
	return &this
}

// GetHeight returns the Height field value
func (o *WidgetLayout) GetHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Height
}

// SetHeight sets field value
func (o *WidgetLayout) SetHeight(v int64) {
	o.Height = v
}

// GetWidth returns the Width field value
func (o *WidgetLayout) GetWidth() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Width
}

// SetWidth sets field value
func (o *WidgetLayout) SetWidth(v int64) {
	o.Width = v
}

// GetX returns the X field value
func (o *WidgetLayout) GetX() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.X
}

// SetX sets field value
func (o *WidgetLayout) SetX(v int64) {
	o.X = v
}

// GetY returns the Y field value
func (o *WidgetLayout) GetY() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Y
}

// SetY sets field value
func (o *WidgetLayout) SetY(v int64) {
	o.Y = v
}

func (o WidgetLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["width"] = o.Width
	}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetLayout struct {
	value *WidgetLayout
	isSet bool
}

func (v NullableWidgetLayout) Get() *WidgetLayout {
	return v.value
}

func (v NullableWidgetLayout) Set(val *WidgetLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetLayout) IsSet() bool {
	return v.isSet
}

func (v NullableWidgetLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetLayout(val *WidgetLayout) *NullableWidgetLayout {
	return &NullableWidgetLayout{value: val, isSet: true}
}

func (v NullableWidgetLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}