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

// Log Object description of a log after being processed and stored by Datadog.
type Log struct {
	Attributes *LogAttributes `json:"attributes,omitempty"`
	// Unique ID of the Log.
	Id   *string  `json:"id,omitempty"`
	Type *LogType `json:"type,omitempty"`
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog() *Log {
	this := Log{}
	var type_ LogType = "log"
	this.Type = &type_
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	var type_ LogType = "log"
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Log) GetAttributes() LogAttributes {
	if o == nil || o.Attributes == nil {
		var ret LogAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetAttributesOk() (*LogAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Log) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given LogAttributes and assigns it to the Attributes field.
func (o *Log) SetAttributes(v LogAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Log) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Log) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Log) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Log) GetType() LogType {
	if o == nil || o.Type == nil {
		var ret LogType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetTypeOk() (*LogType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Log) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given LogType and assigns it to the Type field.
func (o *Log) SetType(v LogType) {
	o.Type = &v
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
