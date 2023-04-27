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

// KnownBotObj struct for KnownBotObj
type KnownBotObj struct {
	ActionType string `json:"action_type"`
	// must be one of available bot tokens
	BotToken             string `json:"bot_token"`
	AdditionalProperties map[string]interface{}
}

type _KnownBotObj KnownBotObj

// NewKnownBotObj instantiates a new KnownBotObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnownBotObj(actionType string, botToken string) *KnownBotObj {
	this := KnownBotObj{}
	this.ActionType = actionType
	this.BotToken = botToken
	return &this
}

// NewKnownBotObjWithDefaults instantiates a new KnownBotObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnownBotObjWithDefaults() *KnownBotObj {
	this := KnownBotObj{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *KnownBotObj) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *KnownBotObj) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *KnownBotObj) SetActionType(v string) {
	o.ActionType = v
}

// GetBotToken returns the BotToken field value
func (o *KnownBotObj) GetBotToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BotToken
}

// GetBotTokenOk returns a tuple with the BotToken field value
// and a boolean to check if the value has been set.
func (o *KnownBotObj) GetBotTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BotToken, true
}

// SetBotToken sets field value
func (o *KnownBotObj) SetBotToken(v string) {
	o.BotToken = v
}

func (o KnownBotObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action_type"] = o.ActionType
	}
	if true {
		toSerialize["bot_token"] = o.BotToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KnownBotObj) UnmarshalJSON(bytes []byte) (err error) {
	varKnownBotObj := _KnownBotObj{}

	if err = json.Unmarshal(bytes, &varKnownBotObj); err == nil {
		*o = KnownBotObj(varKnownBotObj)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action_type")
		delete(additionalProperties, "bot_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKnownBotObj struct {
	value *KnownBotObj
	isSet bool
}

func (v NullableKnownBotObj) Get() *KnownBotObj {
	return v.value
}

func (v *NullableKnownBotObj) Set(val *KnownBotObj) {
	v.value = val
	v.isSet = true
}

func (v NullableKnownBotObj) IsSet() bool {
	return v.isSet
}

func (v *NullableKnownBotObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnownBotObj(val KnownBotObj) NullableKnownBotObj {
	return NullableKnownBotObj{value: &val, isSet: true}
}

func (v NullableKnownBotObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnownBotObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}