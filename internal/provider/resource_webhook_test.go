package provider

import (
	"log"
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func accProvider() map[string]*schema.Provider {
	log.Println("[INFO] Setting up override for a provider")
	provider := Provider()

	return map[string]*schema.Provider{
		"middesk": provider,
	}
}

func TestAcc_ACLCreateAndUpdate(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers:  accProvider(),
		IsUnitTest: false,
		Steps: []r.TestStep{
			{
				Config: testResourceWebhook_initialConfig,
				//Check:  testResourceACL_initialCheck,
			},
			{
				ResourceName:      "middesk_webhook.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testResourceWebhook_initialConfig = `
resource "middesk_webhook" "test" {
  url    = "https://example.org"
  secret = "abc123"
}
`
