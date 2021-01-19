package provider

import (
	"context"
	"log"

	"github.com/Mongey/terraform-provider-middesk/internal/middesk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MIDDESK_API_KEY", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"middesk_webhook": webhookResource(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Printf("[INFO] Initializing Middesk client")
	apiKey := d.Get("api_key").(string)
	cfg := &middesk.Config{APIKey: apiKey}
	c := middesk.NewClient(cfg)

	return c, nil
}
