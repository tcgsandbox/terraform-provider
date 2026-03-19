package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// --- mapLorePostToState ---

func TestMapLorePostToState_AllFields(t *testing.T) {
	t.Parallel()
	post := &LorePost{
		Id:     "post-123",
		GameId: "game-456",
		Title:  "The Origin Story",
	}

	var id, gameID, title types.String
	mapLorePostToState(post, &id, &gameID, &title)

	if id.ValueString() != "post-123" {
		t.Errorf("expected id=post-123, got %s", id.ValueString())
	}
	if gameID.ValueString() != "game-456" {
		t.Errorf("expected gameID=game-456, got %s", gameID.ValueString())
	}
	if title.ValueString() != "The Origin Story" {
		t.Errorf("expected title=The Origin Story, got %s", title.ValueString())
	}
}

func TestMapLorePostToState_OverwritesPriorValues(t *testing.T) {
	t.Parallel()
	post := &LorePost{
		Id:     "new-post",
		GameId: "new-game",
		Title:  "New Title",
	}

	id := types.StringValue("old-post")
	gameID := types.StringValue("old-game")
	title := types.StringValue("Old Title")

	mapLorePostToState(post, &id, &gameID, &title)

	if id.ValueString() != "new-post" {
		t.Errorf("expected id=new-post, got %s", id.ValueString())
	}
	if gameID.ValueString() != "new-game" {
		t.Errorf("expected gameID=new-game, got %s", gameID.ValueString())
	}
	if title.ValueString() != "New Title" {
		t.Errorf("expected title=New Title, got %s", title.ValueString())
	}
}

func TestMapLorePostToState_EmptyStrings(t *testing.T) {
	t.Parallel()
	post := &LorePost{
		Id:     "",
		GameId: "",
		Title:  "",
	}

	var id, gameID, title types.String
	mapLorePostToState(post, &id, &gameID, &title)

	if id.ValueString() != "" {
		t.Errorf("expected empty id, got %s", id.ValueString())
	}
	if gameID.ValueString() != "" {
		t.Errorf("expected empty gameID, got %s", gameID.ValueString())
	}
	if title.ValueString() != "" {
		t.Errorf("expected empty title, got %s", title.ValueString())
	}
}
