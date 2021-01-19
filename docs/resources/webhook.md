---
page_title: "middesk_webhook Resource - terraform-provider-middesk"
subcategory: ""
description: |-
  
---

# Resource `middesk_webhook`





## Schema

### Required

- **url** (String) The URL of the webhook endpoint.

### Optional

- **id** (String) The ID of this resource.
- **secret** (String, Sensitive) The endpoint's secret, used to generate webhook signatures. If missing, a signature will not be provided in webhook request headers.


