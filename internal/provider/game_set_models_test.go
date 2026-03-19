package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// --- mapGameSetFromAPI ---

func TestMapGameSetFromAPI_AllFields(t *testing.T) {
	t.Parallel()
	desc := "A test set"
	set := &GameSet{
		Id:          "set-abc",
		GameId:      "game-xyz",
		Name:        "Base Set",
		Description: &desc,
		Attributes: map[string]GameAttributeType{
			"power": "number",
			"name":  "string",
		},
	}

	var state gameSetModel
	mapGameSetFromAPI(set, &state)

	if state.ID.ValueString() != "set-abc" {
		t.Errorf("expected ID=set-abc, got %s", state.ID.ValueString())
	}
	if state.GameID.ValueString() != "game-xyz" {
		t.Errorf("expected GameID=game-xyz, got %s", state.GameID.ValueString())
	}
	if state.Name.ValueString() != "Base Set" {
		t.Errorf("expected Name=Base Set, got %s", state.Name.ValueString())
	}
	if state.Description.ValueString() != "A test set" {
		t.Errorf("expected Description=A test set, got %s", state.Description.ValueString())
	}
	if len(state.Attributes) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(state.Attributes))
	}
	if state.Attributes["power"].ValueString() != "number" {
		t.Errorf("expected power=number, got %s", state.Attributes["power"].ValueString())
	}
	if state.Attributes["name"].ValueString() != "string" {
		t.Errorf("expected name=string, got %s", state.Attributes["name"].ValueString())
	}
}

func TestMapGameSetFromAPI_NilDescription(t *testing.T) {
	t.Parallel()
	set := &GameSet{
		Id:         "set-1",
		GameId:     "game-1",
		Name:       "Minimal Set",
		Attributes: map[string]GameAttributeType{},
	}

	var state gameSetModel
	mapGameSetFromAPI(set, &state)

	if !state.Description.IsNull() {
		t.Errorf("expected null description, got %s", state.Description.ValueString())
	}
	if len(state.Attributes) != 0 {
		t.Errorf("expected 0 attributes, got %d", len(state.Attributes))
	}
}

func TestMapGameSetFromAPI_PreservesExistingStateFields(t *testing.T) {
	t.Parallel()
	set := &GameSet{
		Id:         "new-id",
		GameId:     "new-game",
		Name:       "New Name",
		Attributes: map[string]GameAttributeType{"hp": "number"},
	}

	state := gameSetModel{
		ID:     types.StringValue("old-id"),
		GameID: types.StringValue("old-game"),
	}
	mapGameSetFromAPI(set, &state)

	// Should overwrite all fields from API
	if state.ID.ValueString() != "new-id" {
		t.Errorf("expected ID=new-id, got %s", state.ID.ValueString())
	}
	if state.GameID.ValueString() != "new-game" {
		t.Errorf("expected GameID=new-game, got %s", state.GameID.ValueString())
	}
}
