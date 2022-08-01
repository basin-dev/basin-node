# CapabilitySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCapabilitySchema

`func NewCapabilitySchema() *CapabilitySchema`

NewCapabilitySchema instantiates a new CapabilitySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySchemaWithDefaults

`func NewCapabilitySchemaWithDefaults() *CapabilitySchema`

NewCapabilitySchemaWithDefaults instantiates a new CapabilitySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CapabilitySchema) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CapabilitySchema) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CapabilitySchema) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CapabilitySchema) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpiration

`func (o *CapabilitySchema) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CapabilitySchema) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CapabilitySchema) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CapabilitySchema) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


