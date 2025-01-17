/*
 * peridot/proto/v1/batch.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peridotopenapi

import (
	"encoding/json"
)

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	PackageName *string `json:"packageName,omitempty"`
	PackageId *string `json:"packageId,omitempty"`
	ScmHash *string `json:"scmHash,omitempty"`
	DisableChecks *bool `json:"disableChecks,omitempty"`
	Branches *[]string `json:"branches,omitempty"`
	ModuleVariant *bool `json:"moduleVariant,omitempty"`
	SideNvrs *[]string `json:"sideNvrs,omitempty"`
	SetInactive *bool `json:"setInactive,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *InlineObject2) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *InlineObject2) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *InlineObject2) SetPackageName(v string) {
	o.PackageName = &v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *InlineObject2) GetPackageId() string {
	if o == nil || o.PackageId == nil {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetPackageIdOk() (*string, bool) {
	if o == nil || o.PackageId == nil {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *InlineObject2) HasPackageId() bool {
	if o != nil && o.PackageId != nil {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *InlineObject2) SetPackageId(v string) {
	o.PackageId = &v
}

// GetScmHash returns the ScmHash field value if set, zero value otherwise.
func (o *InlineObject2) GetScmHash() string {
	if o == nil || o.ScmHash == nil {
		var ret string
		return ret
	}
	return *o.ScmHash
}

// GetScmHashOk returns a tuple with the ScmHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetScmHashOk() (*string, bool) {
	if o == nil || o.ScmHash == nil {
		return nil, false
	}
	return o.ScmHash, true
}

// HasScmHash returns a boolean if a field has been set.
func (o *InlineObject2) HasScmHash() bool {
	if o != nil && o.ScmHash != nil {
		return true
	}

	return false
}

// SetScmHash gets a reference to the given string and assigns it to the ScmHash field.
func (o *InlineObject2) SetScmHash(v string) {
	o.ScmHash = &v
}

// GetDisableChecks returns the DisableChecks field value if set, zero value otherwise.
func (o *InlineObject2) GetDisableChecks() bool {
	if o == nil || o.DisableChecks == nil {
		var ret bool
		return ret
	}
	return *o.DisableChecks
}

// GetDisableChecksOk returns a tuple with the DisableChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetDisableChecksOk() (*bool, bool) {
	if o == nil || o.DisableChecks == nil {
		return nil, false
	}
	return o.DisableChecks, true
}

// HasDisableChecks returns a boolean if a field has been set.
func (o *InlineObject2) HasDisableChecks() bool {
	if o != nil && o.DisableChecks != nil {
		return true
	}

	return false
}

// SetDisableChecks gets a reference to the given bool and assigns it to the DisableChecks field.
func (o *InlineObject2) SetDisableChecks(v bool) {
	o.DisableChecks = &v
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *InlineObject2) GetBranches() []string {
	if o == nil || o.Branches == nil {
		var ret []string
		return ret
	}
	return *o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBranchesOk() (*[]string, bool) {
	if o == nil || o.Branches == nil {
		return nil, false
	}
	return o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *InlineObject2) HasBranches() bool {
	if o != nil && o.Branches != nil {
		return true
	}

	return false
}

// SetBranches gets a reference to the given []string and assigns it to the Branches field.
func (o *InlineObject2) SetBranches(v []string) {
	o.Branches = &v
}

// GetModuleVariant returns the ModuleVariant field value if set, zero value otherwise.
func (o *InlineObject2) GetModuleVariant() bool {
	if o == nil || o.ModuleVariant == nil {
		var ret bool
		return ret
	}
	return *o.ModuleVariant
}

// GetModuleVariantOk returns a tuple with the ModuleVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetModuleVariantOk() (*bool, bool) {
	if o == nil || o.ModuleVariant == nil {
		return nil, false
	}
	return o.ModuleVariant, true
}

// HasModuleVariant returns a boolean if a field has been set.
func (o *InlineObject2) HasModuleVariant() bool {
	if o != nil && o.ModuleVariant != nil {
		return true
	}

	return false
}

// SetModuleVariant gets a reference to the given bool and assigns it to the ModuleVariant field.
func (o *InlineObject2) SetModuleVariant(v bool) {
	o.ModuleVariant = &v
}

// GetSideNvrs returns the SideNvrs field value if set, zero value otherwise.
func (o *InlineObject2) GetSideNvrs() []string {
	if o == nil || o.SideNvrs == nil {
		var ret []string
		return ret
	}
	return *o.SideNvrs
}

// GetSideNvrsOk returns a tuple with the SideNvrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetSideNvrsOk() (*[]string, bool) {
	if o == nil || o.SideNvrs == nil {
		return nil, false
	}
	return o.SideNvrs, true
}

// HasSideNvrs returns a boolean if a field has been set.
func (o *InlineObject2) HasSideNvrs() bool {
	if o != nil && o.SideNvrs != nil {
		return true
	}

	return false
}

// SetSideNvrs gets a reference to the given []string and assigns it to the SideNvrs field.
func (o *InlineObject2) SetSideNvrs(v []string) {
	o.SideNvrs = &v
}

// GetSetInactive returns the SetInactive field value if set, zero value otherwise.
func (o *InlineObject2) GetSetInactive() bool {
	if o == nil || o.SetInactive == nil {
		var ret bool
		return ret
	}
	return *o.SetInactive
}

// GetSetInactiveOk returns a tuple with the SetInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetSetInactiveOk() (*bool, bool) {
	if o == nil || o.SetInactive == nil {
		return nil, false
	}
	return o.SetInactive, true
}

// HasSetInactive returns a boolean if a field has been set.
func (o *InlineObject2) HasSetInactive() bool {
	if o != nil && o.SetInactive != nil {
		return true
	}

	return false
}

// SetSetInactive gets a reference to the given bool and assigns it to the SetInactive field.
func (o *InlineObject2) SetSetInactive(v bool) {
	o.SetInactive = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PackageName != nil {
		toSerialize["packageName"] = o.PackageName
	}
	if o.PackageId != nil {
		toSerialize["packageId"] = o.PackageId
	}
	if o.ScmHash != nil {
		toSerialize["scmHash"] = o.ScmHash
	}
	if o.DisableChecks != nil {
		toSerialize["disableChecks"] = o.DisableChecks
	}
	if o.Branches != nil {
		toSerialize["branches"] = o.Branches
	}
	if o.ModuleVariant != nil {
		toSerialize["moduleVariant"] = o.ModuleVariant
	}
	if o.SideNvrs != nil {
		toSerialize["sideNvrs"] = o.SideNvrs
	}
	if o.SetInactive != nil {
		toSerialize["setInactive"] = o.SetInactive
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


