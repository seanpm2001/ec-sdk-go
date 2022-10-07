// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
Customer Origins API v3

List of API of config Customer Origin.

API version: 0.5.0
Contact: portals-streaming@edgecast.com
*/

package originv3

import (
	"encoding/json"
)

// CustomerOriginFailoverOrder struct for CustomerOriginFailoverOrder
type CustomerOriginFailoverOrder struct {
	Id                   *float32 `json:"id,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	Host                 *string  `json:"host,omitempty"`
	Port                 *int32   `json:"port,omitempty"`
	IsPrimary            *bool    `json:"is_primary,omitempty"`
	StorageTypeId        *int32   `json:"storage_type_id,omitempty"`
	ProtocolTypeId       *int32   `json:"protocol_type_id,omitempty"`
	FailoverOrder        *int32   `json:"failover_order,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerOriginFailoverOrder CustomerOriginFailoverOrder

// NewCustomerOriginFailoverOrder instantiates a new CustomerOriginFailoverOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerOriginFailoverOrder() *CustomerOriginFailoverOrder {
	this := CustomerOriginFailoverOrder{}
	return &this
}

// NewCustomerOriginFailoverOrderWithDefaults instantiates a new CustomerOriginFailoverOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerOriginFailoverOrderWithDefaults() *CustomerOriginFailoverOrder {
	this := CustomerOriginFailoverOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetId() float32 {
	if o == nil || o.Id == nil {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetIdOk() (*float32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *CustomerOriginFailoverOrder) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerOriginFailoverOrder) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CustomerOriginFailoverOrder) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *CustomerOriginFailoverOrder) SetPort(v int32) {
	o.Port = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *CustomerOriginFailoverOrder) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetStorageTypeId returns the StorageTypeId field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetStorageTypeId() int32 {
	if o == nil || o.StorageTypeId == nil {
		var ret int32
		return ret
	}
	return *o.StorageTypeId
}

// GetStorageTypeIdOk returns a tuple with the StorageTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetStorageTypeIdOk() (*int32, bool) {
	if o == nil || o.StorageTypeId == nil {
		return nil, false
	}
	return o.StorageTypeId, true
}

// HasStorageTypeId returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasStorageTypeId() bool {
	if o != nil && o.StorageTypeId != nil {
		return true
	}

	return false
}

// SetStorageTypeId gets a reference to the given int32 and assigns it to the StorageTypeId field.
func (o *CustomerOriginFailoverOrder) SetStorageTypeId(v int32) {
	o.StorageTypeId = &v
}

// GetProtocolTypeId returns the ProtocolTypeId field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetProtocolTypeId() int32 {
	if o == nil || o.ProtocolTypeId == nil {
		var ret int32
		return ret
	}
	return *o.ProtocolTypeId
}

// GetProtocolTypeIdOk returns a tuple with the ProtocolTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetProtocolTypeIdOk() (*int32, bool) {
	if o == nil || o.ProtocolTypeId == nil {
		return nil, false
	}
	return o.ProtocolTypeId, true
}

// HasProtocolTypeId returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasProtocolTypeId() bool {
	if o != nil && o.ProtocolTypeId != nil {
		return true
	}

	return false
}

// SetProtocolTypeId gets a reference to the given int32 and assigns it to the ProtocolTypeId field.
func (o *CustomerOriginFailoverOrder) SetProtocolTypeId(v int32) {
	o.ProtocolTypeId = &v
}

// GetFailoverOrder returns the FailoverOrder field value if set, zero value otherwise.
func (o *CustomerOriginFailoverOrder) GetFailoverOrder() int32 {
	if o == nil || o.FailoverOrder == nil {
		var ret int32
		return ret
	}
	return *o.FailoverOrder
}

// GetFailoverOrderOk returns a tuple with the FailoverOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerOriginFailoverOrder) GetFailoverOrderOk() (*int32, bool) {
	if o == nil || o.FailoverOrder == nil {
		return nil, false
	}
	return o.FailoverOrder, true
}

// HasFailoverOrder returns a boolean if a field has been set.
func (o *CustomerOriginFailoverOrder) HasFailoverOrder() bool {
	if o != nil && o.FailoverOrder != nil {
		return true
	}

	return false
}

// SetFailoverOrder gets a reference to the given int32 and assigns it to the FailoverOrder field.
func (o *CustomerOriginFailoverOrder) SetFailoverOrder(v int32) {
	o.FailoverOrder = &v
}

func (o CustomerOriginFailoverOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.IsPrimary != nil {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if o.StorageTypeId != nil {
		toSerialize["storage_type_id"] = o.StorageTypeId
	}
	if o.ProtocolTypeId != nil {
		toSerialize["protocol_type_id"] = o.ProtocolTypeId
	}
	if o.FailoverOrder != nil {
		toSerialize["failover_order"] = o.FailoverOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomerOriginFailoverOrder) UnmarshalJSON(bytes []byte) (err error) {
	varCustomerOriginFailoverOrder := _CustomerOriginFailoverOrder{}

	if err = json.Unmarshal(bytes, &varCustomerOriginFailoverOrder); err == nil {
		*o = CustomerOriginFailoverOrder(varCustomerOriginFailoverOrder)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "is_primary")
		delete(additionalProperties, "storage_type_id")
		delete(additionalProperties, "protocol_type_id")
		delete(additionalProperties, "failover_order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerOriginFailoverOrder struct {
	value *CustomerOriginFailoverOrder
	isSet bool
}

func (v NullableCustomerOriginFailoverOrder) Get() *CustomerOriginFailoverOrder {
	return v.value
}

func (v *NullableCustomerOriginFailoverOrder) Set(val *CustomerOriginFailoverOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerOriginFailoverOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerOriginFailoverOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerOriginFailoverOrder(val CustomerOriginFailoverOrder) NullableCustomerOriginFailoverOrder {
	return NullableCustomerOriginFailoverOrder{value: &val, isSet: true}
}

func (v NullableCustomerOriginFailoverOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerOriginFailoverOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}