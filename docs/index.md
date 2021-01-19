---
layout: ""
page_title: "Provider: Middesk"
description: |-
  The Middesk provider provides resources to interact with a Middesk API.
---

# Middesk Provider

The Middesk provider provides resources to interact with the Middesk API.

## Example Usage

```terraform
provider "middesk" {
  api_key = "abc_123" # optionally use MIDDESK_API_KEY env var
}
```

## Schema

### Optional

- **api_key** (String, Sensitive)
