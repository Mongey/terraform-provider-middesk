package provider

import (
	"context"
	"log"

	"github.com/Mongey/terraform-provider-middesk/internal/middesk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func webhookResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: webhookCreate,
		ReadContext:   webhookRead,
		UpdateContext: webhookUpdate,
		DeleteContext: webhookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "The URL of the webhook endpoint.",
			},
			"secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The endpoint's secret, used to generate webhook signatures. If missing, a signature will not be provided in webhook request headers.",
			},
		},
	}
}

func webhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*middesk.Client)

	log.Printf("[INFO] Creating Middesk webhook")
	secret := d.Get("secret").(string)

	req := &middesk.WebhookRequest{
		URL:    d.Get("url").(string),
		Secret: &secret,
	}

	log.Printf("[INFO] Really Creating Middesk webhook")
	webhook, err := c.CreateWebhook(req)
	log.Printf("[INFO] Attempted to create Middesk webhook: %v\n %v", err, webhook)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(webhook.ID)

	return nil
}

func webhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*middesk.Client)
	id := d.Id()

	log.Printf("[INFO] Fetching Middesk webhook %s", id)
	webhook, err := c.GetWebhook(id)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Got Middesk webhook %v", webhook)
	err = d.Set("url", webhook.URL)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func webhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*middesk.Client)
	id := d.Id()

	secret := d.Get("secret").(string)
	req := &middesk.WebhookRequest{
		URL:    d.Get("url").(string),
		Secret: &secret,
	}
	_, err := c.UpdateWebhook(id, req)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func webhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*middesk.Client)
	err := c.DeleteWebhook(d.Id())

	return diag.FromErr(err)
}
