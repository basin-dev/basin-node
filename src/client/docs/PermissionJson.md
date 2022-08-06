# PermissionJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** |  | [optional] 
**Capabilities** | Pointer to [**[]CapabilitySchema**](CapabilitySchema.md) |  | [optional] 
**Entities** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPermissionJson

`func NewPermissionJson() *PermissionJson`

NewPermissionJson instantiates a new PermissionJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionJsonWithDefaults

`func NewPermissionJsonWithDefaults() *PermissionJson`

NewPermissionJsonWithDefaults instantiates a new PermissionJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PermissionJson) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PermissionJson) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PermissionJson) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *PermissionJson) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCapabilities

`func (o *PermissionJson) GetCapabilities() []CapabilitySchema`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PermissionJson) GetCapabilitiesOk() (*[]CapabilitySchema, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PermissionJson) SetCapabilities(v []CapabilitySchema)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PermissionJson) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetEntities

`func (o *PermissionJson) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PermissionJson) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PermissionJson) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *PermissionJson) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


