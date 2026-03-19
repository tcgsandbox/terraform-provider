package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGameSetDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "tcg-sandbox_game" "test" {
  name                      = "Terraform Acc GameSet DS Test"
  description               = "Game for game_set data source test"
  banner_image_path         = "testdata/test_banner.png"
  banner_vertical_alignment = 50
  attributes = {
    "speed" = "number"
  }
}

resource "tcg-sandbox_game_set" "test" {
  id      = "ds-test-set"
  game_id = tcg-sandbox_game.test.id
  name    = "DS Test Set"
  attributes = {
    "speed" = "number"
  }
}

data "tcg-sandbox_game_set" "test" {
  id      = tcg-sandbox_game_set.test.id
  game_id = tcg-sandbox_game.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.tcg-sandbox_game_set.test", "id",
						"tcg-sandbox_game_set.test", "id",
					),
					resource.TestCheckResourceAttrPair(
						"data.tcg-sandbox_game_set.test", "name",
						"tcg-sandbox_game_set.test", "name",
					),
					resource.TestCheckResourceAttrPair(
						"data.tcg-sandbox_game_set.test", "game_id",
						"tcg-sandbox_game_set.test", "game_id",
					),
				),
			},
		},
	})
}
