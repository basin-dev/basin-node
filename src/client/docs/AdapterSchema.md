# AdapterSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdapterName** | **string** |  | 
**Config** | **map[string]interface{}** |  | 

## Methods

### NewAdapterSchema

`func NewAdapterSchema(adapterName string, config map[string]interface{}, ) *AdapterSchema`

NewAdapterSchema instantiates a new AdapterSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterSchemaWithDefaults

`func NewAdapterSchemaWithDefaults() *AdapterSchema`

NewAdapterSchemaWithDefaults instantiates a new AdapterSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapterName

`func (o *AdapterSchema) GetAdapterName() string`

GetAdapterName returns the AdapterName field if non-nil, zero value otherwise.

### GetAdapterNameOk

`func (o *AdapterSchema) GetAdapterNameOk() (*string, bool)`

GetAdapterNameOk returns a tuple with the AdapterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterName

`func (o *AdapterSchema) SetAdapterName(v string)`

SetAdapterName sets AdapterName field to given value.


### GetConfig

`func (o *AdapterSchema) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AdapterSchema) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AdapterSchema) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


