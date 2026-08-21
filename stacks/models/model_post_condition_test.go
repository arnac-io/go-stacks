package models

import (
	"encoding/json"
	"testing"
)

func setVariants(pc PostCondition) []string {
	var set []string
	if pc.PostConditionFungible != nil {
		set = append(set, "fungible")
	}
	if pc.PostConditionNonFungible != nil {
		set = append(set, "non_fungible")
	}
	if pc.PostConditionPox != nil {
		set = append(set, "pox")
	}
	if pc.PostConditionStaking != nil {
		set = append(set, "staking")
	}
	if pc.PostConditionStx != nil {
		set = append(set, "stx")
	}
	return set
}

func TestPostConditionUnmarshalDispatchesOnType(t *testing.T) {
	principal := `"principal": {"type_id": "principal_standard", "address": "SP10B4KJ16AFGCYH7BQKEQ39G1SRYCHVP3XBRPK65"}`
	asset := `"asset": {"asset_name": "token", "contract_address": "SP10B4KJ16AFGCYH7BQKEQ39G1SRYCHVP3XBRPK65", "contract_name": "token-contract"}`

	tests := []struct {
		name    string
		data    string
		want    string
		wantErr bool
	}{
		{
			name: "stx",
			data: `{` + principal + `, "condition_code": "sent_equal_to", "amount": "1000", "type": "stx"}`,
			want: "stx",
		},
		{
			name: "fungible",
			data: `{` + principal + `, ` + asset + `, "condition_code": "sent_equal_to", "amount": "1000", "type": "fungible"}`,
			want: "fungible",
		},
		{
			name: "non fungible",
			data: `{` + principal + `, ` + asset + `, "asset_value": {"hex": "0x01", "repr": "u1"}, "condition_code": "sent", "type": "non_fungible"}`,
			want: "non_fungible",
		},
		{
			name: "pox",
			data: `{` + principal + `, "condition_code": "performed", "type": "pox"}`,
			want: "pox",
		},
		{
			name: "staking is not read as stx despite identical fields",
			data: `{"principal": {"type_id": "principal_standard", "address": "SP39PWER4NT0E7EK73354ZJKP9HCKBYGK72609M9G"}, "condition_code": "sent_less_than_or_equal_to", "amount": "346002000000", "type": "staking"}`,
			want: "staking",
		},
		{
			name:    "unrecognised type",
			data:    `{` + principal + `, "condition_code": "sent_equal_to", "amount": "1000", "type": "something_new"}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var pc PostCondition
			err := json.Unmarshal([]byte(test.data), &pc)

			if test.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got variants %v", setVariants(pc))
				}
				return
			}
			if err != nil {
				t.Fatalf("unmarshal: %v", err)
			}

			got := setVariants(pc)
			if len(got) != 1 || got[0] != test.want {
				t.Fatalf("expected only %q to be set, got %v", test.want, got)
			}
		})
	}
}
