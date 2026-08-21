package models

import (
	"encoding/json"
	"fmt"
)

// PostCondition Post-conditionscan limit the damage done to a user's assets
type PostCondition struct {
	PostConditionFungible    *PostConditionFungible
	PostConditionNonFungible *PostConditionNonFungible
	PostConditionPox         *PostConditionPox
	PostConditionStaking     *PostConditionStaking
	PostConditionStx         *PostConditionStx
}

// Unmarshal JSON data into the pointer matching the "type" discriminator.
// Staking and stx carry identical fields, so trying each variant in turn
// silently picks the wrong one.
func (dst *PostCondition) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	dst.PostConditionFungible = nil
	dst.PostConditionNonFungible = nil
	dst.PostConditionPox = nil
	dst.PostConditionStaking = nil
	dst.PostConditionStx = nil

	switch discriminator.Type {
	case "fungible":
		dst.PostConditionFungible = &PostConditionFungible{}
		return json.Unmarshal(data, dst.PostConditionFungible)
	case "non_fungible":
		dst.PostConditionNonFungible = &PostConditionNonFungible{}
		return json.Unmarshal(data, dst.PostConditionNonFungible)
	case "pox":
		dst.PostConditionPox = &PostConditionPox{}
		return json.Unmarshal(data, dst.PostConditionPox)
	case "staking":
		dst.PostConditionStaking = &PostConditionStaking{}
		return json.Unmarshal(data, dst.PostConditionStaking)
	case "stx":
		dst.PostConditionStx = &PostConditionStx{}
		return json.Unmarshal(data, dst.PostConditionStx)
	default:
		return fmt.Errorf("unknown post condition type %q in anyOf(PostCondition)", discriminator.Type)
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PostCondition) MarshalJSON() ([]byte, error) {
	if src.PostConditionFungible != nil {
		return json.Marshal(&src.PostConditionFungible)
	}

	if src.PostConditionNonFungible != nil {
		return json.Marshal(&src.PostConditionNonFungible)
	}

	if src.PostConditionPox != nil {
		return json.Marshal(&src.PostConditionPox)
	}

	if src.PostConditionStaking != nil {
		return json.Marshal(&src.PostConditionStaking)
	}

	if src.PostConditionStx != nil {
		return json.Marshal(&src.PostConditionStx)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePostCondition struct {
	value *PostCondition
	isSet bool
}

func (v NullablePostCondition) Get() *PostCondition {
	return v.value
}

func (v *NullablePostCondition) Set(val *PostCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCondition(val *PostCondition) *NullablePostCondition {
	return &NullablePostCondition{value: val, isSet: true}
}

func (v NullablePostCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
