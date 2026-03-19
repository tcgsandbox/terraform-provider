package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGameSetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and verify
			{
				Config: providerConfig + `
resource "tcg-sandbox_game" "test" {
  name                      = "Terraform Acc GameSet Test"
  description               = "Game for game_set acceptance test"
  banner_image_path         = "testdata/test_banner.png"
  banner_vertical_alignment = 50
  attributes = {
    "power" = "number"
  }
}

resource "tcg-sandbox_game_set" "test" {
  id      = "test-set"
  game_id = tcg-sandbox_game.test.id
  name    = "Test Set"
  description = "A test set"
  attributes = {
    "power" = "number"
  }
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tcg-sandbox_game_set.test", "id", "test-set"),
					resource.TestCheckResourceAttrSet("tcg-sandbox_game_set.test", "game_id"),
					resource.TestCheckResourceAttr("tcg-sandbox_game_set.test", "name", "Test Set"),
					resource.TestCheckResourceAttr("tcg-sandbox_game_set.test", "description", "A test set"),
					resource.TestCheckResourceAttr("tcg-sandbox_game_set.test", "attributes.power", "number"),
				),
			},
		},
	})
}
