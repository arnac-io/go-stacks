package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/arnac-io/go-stacks/stacks/utils"
)

var _ utils.MappedNullable = &PostConditionPox{}

// PostConditionPox struct for PostConditionPox
type PostConditionPox struct {
	Principal     PostConditionPrincipal `json:"principal"`
	ConditionCode string                 `json:"condition_code"`
	Type          string                 `json:"type"`
}

type _PostConditionPox PostConditionPox

func NewPostConditionPox(principal PostConditionPrincipal, conditionCode string, type_ string) *PostConditionPox {
	this := PostConditionPox{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Type = type_
	return &this
}

func NewPostConditionPoxWithDefaults() *PostConditionPox {
	this := PostConditionPox{}
	return &this
}

func (o *PostConditionPox) GetPrincipal() PostConditionPrincipal {
	if o == nil {
		var ret PostConditionPrincipal
		return ret
	}

	return o.Principal
}

func (o *PostConditionPox) GetPrincipalOk() (*PostConditionPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

func (o *PostConditionPox) SetPrincipal(v PostConditionPrincipal) {
	o.Principal = v
}

func (o *PostConditionPox) GetConditionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionCode
}

func (o *PostConditionPox) GetConditionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

func (o *PostConditionPox) SetConditionCode(v string) {
	o.ConditionCode = v
}

func (o *PostConditionPox) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

func (o *PostConditionPox) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

func (o *PostConditionPox) SetType(v string) {
	o.Type = v
}

func (o PostConditionPox) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionPox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostConditionPox) UnmarshalJSON(data []byte) (err error) {
	requiredProperties := []string{
		"principal",
		"condition_code",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostConditionPox := _PostConditionPox{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionPox)

	if err != nil {
		return err
	}

	*o = PostConditionPox(varPostConditionPox)

	return err
}

type NullablePostConditionPox struct {
	value *PostConditionPox
	isSet bool
}

func (v NullablePostConditionPox) Get() *PostConditionPox {
	return v.value
}

func (v *NullablePostConditionPox) Set(val *PostConditionPox) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionPox) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionPox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionPox(val *PostConditionPox) *NullablePostConditionPox {
	return &NullablePostConditionPox{value: val, isSet: true}
}

func (v NullablePostConditionPox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionPox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
