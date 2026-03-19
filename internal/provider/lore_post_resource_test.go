package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLorePostResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and verify
			{
				Config: providerConfig + `
resource "tcg-sandbox_game" "test" {
  name                      = "Terraform Acc LorePost Test"
  description               = "Game for lore_post acceptance test"
  banner_image_path         = "testdata/test_banner.png"
  banner_vertical_alignment = 50
  attributes = {
    "power" = "number"
  }
}

resource "tcg-sandbox_lore_post" "test" {
  game_id = tcg-sandbox_game.test.id
  title   = "The Beginning"
  content = "# Chapter 1\n\nOnce upon a time..."
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("tcg-sandbox_lore_post.test", "id"),
					resource.TestCheckResourceAttrSet("tcg-sandbox_lore_post.test", "game_id"),
					resource.TestCheckResourceAttr("tcg-sandbox_lore_post.test", "title", "The Beginning"),
					resource.TestCheckResourceAttr("tcg-sandbox_lore_post.test", "content", "# Chapter 1\n\nOnce upon a time..."),
				),
			},
			// Update title and content
			{
				Config: providerConfig + `
resource "tcg-sandbox_game" "test" {
  name                      = "Terraform Acc LorePost Test"
  description               = "Game for lore_post acceptance test"
  banner_image_path         = "testdata/test_banner.png"
  banner_vertical_alignment = 50
  attributes = {
    "power" = "number"
  }
}

resource "tcg-sandbox_lore_post" "test" {
  game_id = tcg-sandbox_game.test.id
  title   = "The Updated Beginning"
  content = "# Chapter 1 (Revised)\n\nIt was a dark and stormy night..."
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tcg-sandbox_lore_post.test", "title", "The Updated Beginning"),
					resource.TestCheckResourceAttr("tcg-sandbox_lore_post.test", "content", "# Chapter 1 (Revised)\n\nIt was a dark and stormy night..."),
				),
			},
		},
	})
}
