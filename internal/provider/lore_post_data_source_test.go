package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLorePostDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "tcg-sandbox_game" "test" {
  name                      = "Terraform Acc LorePost DS Test"
  description               = "Game for lore_post data source test"
  banner_image_path         = "testdata/test_banner.png"
  banner_vertical_alignment = 50
  attributes = {
    "power" = "number"
  }
}

resource "tcg-sandbox_lore_post" "test" {
  game_id = tcg-sandbox_game.test.id
  title   = "DS Test Post"
  content = "Some lore content"
}

data "tcg-sandbox_lore_post" "test" {
  id      = tcg-sandbox_lore_post.test.id
  game_id = tcg-sandbox_game.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.tcg-sandbox_lore_post.test", "id",
						"tcg-sandbox_lore_post.test", "id",
					),
					resource.TestCheckResourceAttrPair(
						"data.tcg-sandbox_lore_post.test", "title",
						"tcg-sandbox_lore_post.test", "title",
					),
					resource.TestCheckResourceAttrPair(
						"data.tcg-sandbox_lore_post.test", "game_id",
						"tcg-sandbox_lore_post.test", "game_id",
					),
				),
			},
		},
	})
}
