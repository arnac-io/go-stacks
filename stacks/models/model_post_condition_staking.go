package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/arnac-io/go-stacks/stacks/utils"
)

var _ utils.MappedNullable = &PostConditionStaking{}

// PostConditionStaking struct for PostConditionStaking
type PostConditionStaking struct {
	Principal     PostConditionPrincipal `json:"principal"`
	ConditionCode string                 `json:"condition_code"`
	Amount        string                 `json:"amount"`
	Type          string                 `json:"type"`
}

type _PostConditionStaking PostConditionStaking

func NewPostConditionStaking(principal PostConditionPrincipal, conditionCode string, amount string, type_ string) *PostConditionStaking {
	this := PostConditionStaking{}
	this.Principal = principal
	this.ConditionCode = conditionCode
	this.Amount = amount
	this.Type = type_
	return &this
}

func NewPostConditionStakingWithDefaults() *PostConditionStaking {
	this := PostConditionStaking{}
	return &this
}

func (o *PostConditionStaking) GetPrincipal() PostConditionPrincipal {
	if o == nil {
		var ret PostConditionPrincipal
		return ret
	}

	return o.Principal
}

func (o *PostConditionStaking) GetPrincipalOk() (*PostConditionPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

func (o *PostConditionStaking) SetPrincipal(v PostConditionPrincipal) {
	o.Principal = v
}

func (o *PostConditionStaking) GetConditionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionCode
}

func (o *PostConditionStaking) GetConditionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionCode, true
}

func (o *PostConditionStaking) SetConditionCode(v string) {
	o.ConditionCode = v
}

func (o *PostConditionStaking) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

func (o *PostConditionStaking) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

func (o *PostConditionStaking) SetAmount(v string) {
	o.Amount = v
}

func (o *PostConditionStaking) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

func (o *PostConditionStaking) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

func (o *PostConditionStaking) SetType(v string) {
	o.Type = v
}

func (o PostConditionStaking) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostConditionStaking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	toSerialize["condition_code"] = o.ConditionCode
	toSerialize["amount"] = o.Amount
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostConditionStaking) UnmarshalJSON(data []byte) (err error) {
	requiredProperties := []string{
		"principal",
		"condition_code",
		"amount",
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

	varPostConditionStaking := _PostConditionStaking{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostConditionStaking)

	if err != nil {
		return err
	}

	*o = PostConditionStaking(varPostConditionStaking)

	return err
}

type NullablePostConditionStaking struct {
	value *PostConditionStaking
	isSet bool
}

func (v NullablePostConditionStaking) Get() *PostConditionStaking {
	return v.value
}

func (v *NullablePostConditionStaking) Set(val *PostConditionStaking) {
	v.value = val
	v.isSet = true
}

func (v NullablePostConditionStaking) IsSet() bool {
	return v.isSet
}

func (v *NullablePostConditionStaking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostConditionStaking(val *PostConditionStaking) *NullablePostConditionStaking {
	return &NullablePostConditionStaking{value: val, isSet: true}
}

func (v NullablePostConditionStaking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostConditionStaking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
