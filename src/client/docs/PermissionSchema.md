# PermissionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** |  | [optional] 
**Capabilities** | Pointer to [**[]CapabilitySchema**](CapabilitySchema.md) |  | [optional] 
**Entities** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPermissionSchema

`func NewPermissionSchema() *PermissionSchema`

NewPermissionSchema instantiates a new PermissionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSchemaWithDefaults

`func NewPermissionSchemaWithDefaults() *PermissionSchema`

NewPermissionSchemaWithDefaults instantiates a new PermissionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PermissionSchema) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PermissionSchema) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PermissionSchema) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *PermissionSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCapabilities

`func (o *PermissionSchema) GetCapabilities() []CapabilitySchema`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *PermissionSchema) GetCapabilitiesOk() (*[]CapabilitySchema, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *PermissionSchema) SetCapabilities(v []CapabilitySchema)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *PermissionSchema) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetEntities

`func (o *PermissionSchema) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PermissionSchema) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PermissionSchema) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *PermissionSchema) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


