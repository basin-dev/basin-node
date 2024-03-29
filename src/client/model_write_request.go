/*
Basin RPC API

Basin RPC API

API version: 1.0.11
Contact: sestinj@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// WriteRequest struct for WriteRequest
type WriteRequest struct {
	Url string `json:"url"`
	Value string `json:"value"`
}

// NewWriteRequest instantiates a new WriteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteRequest(url string, value string) *WriteRequest {
	this := WriteRequest{}
	this.Url = url
	this.Value = value
	return &this
}

// NewWriteRequestWithDefaults instantiates a new WriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRequestWithDefaults() *WriteRequest {
	this := WriteRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *WriteRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WriteRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WriteRequest) SetUrl(v string) {
	o.Url = v
}

// GetValue returns the Value field value
func (o *WriteRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WriteRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *WriteRequest) SetValue(v string) {
	o.Value = v
}

func (o WriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWriteRequest struct {
	value *WriteRequest
	isSet bool
}

func (v NullableWriteRequest) Get() *WriteRequest {
	return v.value
}

func (v *NullableWriteRequest) Set(val *WriteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteRequest(val *WriteRequest) *NullableWriteRequest {
	return &NullableWriteRequest{value: val, isSet: true}
}

func (v NullableWriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


