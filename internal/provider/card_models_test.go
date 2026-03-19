package provider

import (
	"testing"
)

// --- mapCardToState ---

func TestMapCardToState_AllFields(t *testing.T) {
	t.Parallel()
	card := &Card{
		Id:                 "card-123",
		GameId:             "game-456",
		SetId:              "set-789",
		Name:               "Fire Dragon",
		Description:        strPtr("A powerful dragon card"),
		CardImagePublicUrl: strPtr("https://example.com/dragon.png"),
		Attributes: map[string]interface{}{
			"power":   "100",
			"element": "fire",
		},
	}

	id, gameID, setID, name, desc, imgURL, attrs := mapCardToState(card)

	if id.ValueString() != "card-123" {
		t.Errorf("expected id=card-123, got %s", id.ValueString())
	}
	if gameID.ValueString() != "game-456" {
		t.Errorf("expected gameID=game-456, got %s", gameID.ValueString())
	}
	if setID.ValueString() != "set-789" {
		t.Errorf("expected setID=set-789, got %s", setID.ValueString())
	}
	if name.ValueString() != "Fire Dragon" {
		t.Errorf("expected name=Fire Dragon, got %s", name.ValueString())
	}
	if desc.ValueString() != "A powerful dragon card" {
		t.Errorf("expected description, got %s", desc.ValueString())
	}
	if imgURL.ValueString() != "https://example.com/dragon.png" {
		t.Errorf("expected image URL, got %s", imgURL.ValueString())
	}
	if len(attrs) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(attrs))
	}
	if attrs["power"].ValueString() != "100" {
		t.Errorf("expected power=100, got %s", attrs["power"].ValueString())
	}
	if attrs["element"].ValueString() != "fire" {
		t.Errorf("expected element=fire, got %s", attrs["element"].ValueString())
	}
}

func TestMapCardToState_NilOptionalFields(t *testing.T) {
	t.Parallel()
	card := &Card{
		Id:         "card-1",
		GameId:     "game-1",
		SetId:      "set-1",
		Name:       "Basic Card",
		Attributes: map[string]interface{}{},
	}

	_, _, _, _, desc, imgURL, attrs := mapCardToState(card)

	if !desc.IsNull() {
		t.Errorf("expected null description, got %s", desc.ValueString())
	}
	if !imgURL.IsNull() {
		t.Errorf("expected null image URL, got %s", imgURL.ValueString())
	}
	if len(attrs) != 0 {
		t.Errorf("expected 0 attributes, got %d", len(attrs))
	}
}

func TestMapCardToState_NonStringAttributeValues(t *testing.T) {
	t.Parallel()
	card := &Card{
		Id:     "card-1",
		GameId: "game-1",
		SetId:  "set-1",
		Name:   "Test Card",
		Attributes: map[string]interface{}{
			"str_val": "hello",
			"num_val": float64(42),
			"bool_val": true,
		},
	}

	_, _, _, _, _, _, attrs := mapCardToState(card)

	if attrs["str_val"].ValueString() != "hello" {
		t.Errorf("expected str_val=hello, got %s", attrs["str_val"].ValueString())
	}
	if attrs["num_val"].ValueString() != "42" {
		t.Errorf("expected num_val=42, got %s", attrs["num_val"].ValueString())
	}
	if attrs["bool_val"].ValueString() != "true" {
		t.Errorf("expected bool_val=true, got %s", attrs["bool_val"].ValueString())
	}
}
