package models

import (
	"encoding/json"
	"testing"
)

func TestPostConditionUnmarshalPox(t *testing.T) {
	data := []byte(`{
		"principal": {
			"type_id": "principal_standard",
			"address": "SP10B4KJ16AFGCYH7BQKEQ39G1SRYCHVP3XBRPK65"
		},
		"condition_code": "performed",
		"type": "pox"
	}`)

	var pc PostCondition
	if err := json.Unmarshal(data, &pc); err != nil {
		t.Fatalf("unmarshal pox post-condition: %v", err)
	}
	if pc.PostConditionPox == nil {
		t.Fatal("expected PostConditionPox to be set")
	}
	if pc.PostConditionPox.Type != "pox" {
		t.Fatalf("expected type pox, got %q", pc.PostConditionPox.Type)
	}
	if pc.PostConditionStx != nil || pc.PostConditionFungible != nil || pc.PostConditionNonFungible != nil {
		t.Fatal("expected only the pox variant to be set")
	}
}

func TestPostConditionUnmarshalStxStillMatchesStx(t *testing.T) {
	data := []byte(`{
		"principal": {
			"type_id": "principal_standard",
			"address": "SP10B4KJ16AFGCYH7BQKEQ39G1SRYCHVP3XBRPK65"
		},
		"condition_code": "sent_equal_to",
		"amount": "1000",
		"type": "stx"
	}`)

	var pc PostCondition
	if err := json.Unmarshal(data, &pc); err != nil {
		t.Fatalf("unmarshal stx post-condition: %v", err)
	}
	if pc.PostConditionStx == nil {
		t.Fatal("expected PostConditionStx to be set")
	}
	if pc.PostConditionPox != nil {
		t.Fatal("stx post-condition must not match the pox variant")
	}
}
