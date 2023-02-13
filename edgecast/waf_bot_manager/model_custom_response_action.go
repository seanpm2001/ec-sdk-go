// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
WAF API

The WAF API is a RESTful server application for managing customer configuration settings.

API version: 1.0
*/

package waf_bot_manager

import (
	"encoding/json"
)

// CustomResponseAction struct for CustomResponseAction
type CustomResponseAction struct {
	Id                   *string          `json:"id,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	EnfType              *string          `json:"enf_type,omitempty"`
	ResponseBodyBase64   *string          `json:"response_body_base64,omitempty"`
	Status               *int32           `json:"status,omitempty"`
	ResponseHeaders      []ResponseHeader `json:"response_headers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomResponseAction CustomResponseAction

// NewCustomResponseAction instantiates a new CustomResponseAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomResponseAction() *CustomResponseAction {
	this := CustomResponseAction{}
	return &this
}

// NewCustomResponseActionWithDefaults instantiates a new CustomResponseAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomResponseActionWithDefaults() *CustomResponseAction {
	this := CustomResponseAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomResponseAction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomResponseAction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomResponseAction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomResponseAction) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomResponseAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomResponseAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomResponseAction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomResponseAction) SetName(v string) {
	o.Name = &v
}

// GetEnfType returns the EnfType field value if set, zero value otherwise.
func (o *CustomResponseAction) GetEnfType() string {
	if o == nil || o.EnfType == nil {
		var ret string
		return ret
	}
	return *o.EnfType
}

// GetEnfTypeOk returns a tuple with the EnfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomResponseAction) GetEnfTypeOk() (*string, bool) {
	if o == nil || o.EnfType == nil {
		return nil, false
	}
	return o.EnfType, true
}

// HasEnfType returns a boolean if a field has been set.
func (o *CustomResponseAction) HasEnfType() bool {
	if o != nil && o.EnfType != nil {
		return true
	}

	return false
}

// SetEnfType gets a reference to the given string and assigns it to the EnfType field.
func (o *CustomResponseAction) SetEnfType(v string) {
	o.EnfType = &v
}

// GetResponseBodyBase64 returns the ResponseBodyBase64 field value if set, zero value otherwise.
func (o *CustomResponseAction) GetResponseBodyBase64() string {
	if o == nil || o.ResponseBodyBase64 == nil {
		var ret string
		return ret
	}
	return *o.ResponseBodyBase64
}

// GetResponseBodyBase64Ok returns a tuple with the ResponseBodyBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomResponseAction) GetResponseBodyBase64Ok() (*string, bool) {
	if o == nil || o.ResponseBodyBase64 == nil {
		return nil, false
	}
	return o.ResponseBodyBase64, true
}

// HasResponseBodyBase64 returns a boolean if a field has been set.
func (o *CustomResponseAction) HasResponseBodyBase64() bool {
	if o != nil && o.ResponseBodyBase64 != nil {
		return true
	}

	return false
}

// SetResponseBodyBase64 gets a reference to the given string and assigns it to the ResponseBodyBase64 field.
func (o *CustomResponseAction) SetResponseBodyBase64(v string) {
	o.ResponseBodyBase64 = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomResponseAction) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomResponseAction) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomResponseAction) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *CustomResponseAction) SetStatus(v int32) {
	o.Status = &v
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise.
func (o *CustomResponseAction) GetResponseHeaders() []ResponseHeader {
	if o == nil || o.ResponseHeaders == nil {
		var ret []ResponseHeader
		return ret
	}
	return o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomResponseAction) GetResponseHeadersOk() ([]ResponseHeader, bool) {
	if o == nil || o.ResponseHeaders == nil {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *CustomResponseAction) HasResponseHeaders() bool {
	if o != nil && o.ResponseHeaders != nil {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given []ResponseHeader and assigns it to the ResponseHeaders field.
func (o *CustomResponseAction) SetResponseHeaders(v []ResponseHeader) {
	o.ResponseHeaders = v
}

func (o CustomResponseAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.EnfType != nil {
		toSerialize["enf_type"] = o.EnfType
	}
	if o.ResponseBodyBase64 != nil {
		toSerialize["response_body_base64"] = o.ResponseBodyBase64
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ResponseHeaders != nil {
		toSerialize["response_headers"] = o.ResponseHeaders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomResponseAction) UnmarshalJSON(bytes []byte) (err error) {
	varCustomResponseAction := _CustomResponseAction{}

	if err = json.Unmarshal(bytes, &varCustomResponseAction); err == nil {
		*o = CustomResponseAction(varCustomResponseAction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enf_type")
		delete(additionalProperties, "response_body_base64")
		delete(additionalProperties, "status")
		delete(additionalProperties, "response_headers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomResponseAction struct {
	value *CustomResponseAction
	isSet bool
}

func (v NullableCustomResponseAction) Get() *CustomResponseAction {
	return v.value
}

func (v *NullableCustomResponseAction) Set(val *CustomResponseAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomResponseAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomResponseAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomResponseAction(val CustomResponseAction) NullableCustomResponseAction {
	return NullableCustomResponseAction{value: &val, isSet: true}
}

func (v NullableCustomResponseAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomResponseAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
