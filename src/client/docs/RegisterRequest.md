# RegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Permissions** | [**[]PermissionJson**](PermissionJson.md) |  | 
**Adapter** | [**AdapterJson**](AdapterJson.md) |  | 
**Schema** | **map[string]interface{}** |  | 

## Methods

### NewRegisterRequest

`func NewRegisterRequest(url string, permissions []PermissionJson, adapter AdapterJson, schema map[string]interface{}, ) *RegisterRequest`

NewRegisterRequest instantiates a new RegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRequestWithDefaults

`func NewRegisterRequestWithDefaults() *RegisterRequest`

NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RegisterRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RegisterRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RegisterRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPermissions

`func (o *RegisterRequest) GetPermissions() []PermissionJson`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RegisterRequest) GetPermissionsOk() (*[]PermissionJson, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RegisterRequest) SetPermissions(v []PermissionJson)`

SetPermissions sets Permissions field to given value.


### GetAdapter

`func (o *RegisterRequest) GetAdapter() AdapterJson`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *RegisterRequest) GetAdapterOk() (*AdapterJson, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *RegisterRequest) SetAdapter(v AdapterJson)`

SetAdapter sets Adapter field to given value.


### GetSchema

`func (o *RegisterRequest) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RegisterRequest) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RegisterRequest) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


