package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/arnac-io/go-stacks/stacks/utils"
)

// checks if the PoxCycleListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PoxCycleListResponse{}

// PoxCycleListResponse GET request that returns PoX cycles
type PoxCycleListResponse struct {
	// The number of cycles to return
	Limit int32 `json:"limit"`
	// The number to cycles to skip (starting at `0`)
	Offset int32 `json:"offset"`
	// The total number of cycles
	Total   int32      `json:"total"`
	Results []PoxCycle `json:"results"`
}

type _PoxCycleListResponse PoxCycleListResponse

// NewPoxCycleListResponse instantiates a new PoxCycleListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoxCycleListResponse(limit int32, offset int32, total int32, results []PoxCycle) *PoxCycleListResponse {
	this := PoxCycleListResponse{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Results = results
	return &this
}

// NewPoxCycleListResponseWithDefaults instantiates a new PoxCycleListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoxCycleListResponseWithDefaults() *PoxCycleListResponse {
	this := PoxCycleListResponse{}
	var offset int32 = 0
	this.Offset = offset
	return &this
}

// GetLimit returns the Limit field value
func (o *PoxCycleListResponse) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PoxCycleListResponse) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PoxCycleListResponse) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *PoxCycleListResponse) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PoxCycleListResponse) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PoxCycleListResponse) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *PoxCycleListResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PoxCycleListResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PoxCycleListResponse) SetTotal(v int32) {
	o.Total = v
}

// GetResults returns the Results field value
func (o *PoxCycleListResponse) GetResults() []PoxCycle {
	if o == nil {
		var ret []PoxCycle
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PoxCycleListResponse) GetResultsOk() ([]PoxCycle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PoxCycleListResponse) SetResults(v []PoxCycle) {
	o.Results = v
}

func (o PoxCycleListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoxCycleListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PoxCycleListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"offset",
		"total",
		"results",
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

	varPoxCycleListResponse := _PoxCycleListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoxCycleListResponse)

	if err != nil {
		return err
	}

	*o = PoxCycleListResponse(varPoxCycleListResponse)

	return err
}

type NullablePoxCycleListResponse struct {
	value *PoxCycleListResponse
	isSet bool
}

func (v NullablePoxCycleListResponse) Get() *PoxCycleListResponse {
	return v.value
}

func (v *NullablePoxCycleListResponse) Set(val *PoxCycleListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePoxCycleListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePoxCycleListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoxCycleListResponse(val *PoxCycleListResponse) *NullablePoxCycleListResponse {
	return &NullablePoxCycleListResponse{value: val, isSet: true}
}

func (v NullablePoxCycleListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoxCycleListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
