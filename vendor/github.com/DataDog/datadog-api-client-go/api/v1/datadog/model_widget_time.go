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

// WidgetTime Time setting for the widget.
type WidgetTime struct {
	LiveSpan *WidgetLiveSpan `json:"live_span,omitempty"`
}

// NewWidgetTime instantiates a new WidgetTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetTime() *WidgetTime {
	this := WidgetTime{}
	return &this
}

// NewWidgetTimeWithDefaults instantiates a new WidgetTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetTimeWithDefaults() *WidgetTime {
	this := WidgetTime{}
	return &this
}

// GetLiveSpan returns the LiveSpan field value if set, zero value otherwise.
func (o *WidgetTime) GetLiveSpan() WidgetLiveSpan {
	if o == nil || o.LiveSpan == nil {
		var ret WidgetLiveSpan
		return ret
	}
	return *o.LiveSpan
}

// GetLiveSpanOk returns a tuple with the LiveSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetTime) GetLiveSpanOk() (*WidgetLiveSpan, bool) {
	if o == nil || o.LiveSpan == nil {
		return nil, false
	}
	return o.LiveSpan, true
}

// HasLiveSpan returns a boolean if a field has been set.
func (o *WidgetTime) HasLiveSpan() bool {
	if o != nil && o.LiveSpan != nil {
		return true
	}

	return false
}

// SetLiveSpan gets a reference to the given WidgetLiveSpan and assigns it to the LiveSpan field.
func (o *WidgetTime) SetLiveSpan(v WidgetLiveSpan) {
	o.LiveSpan = &v
}

func (o WidgetTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LiveSpan != nil {
		toSerialize["live_span"] = o.LiveSpan
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetTime struct {
	value *WidgetTime
	isSet bool
}

func (v NullableWidgetTime) Get() *WidgetTime {
	return v.value
}

func (v *NullableWidgetTime) Set(val *WidgetTime) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetTime) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetTime(val *WidgetTime) *NullableWidgetTime {
	return &NullableWidgetTime{value: val, isSet: true}
}

func (v NullableWidgetTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
