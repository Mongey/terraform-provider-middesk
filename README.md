# `terraform-provider-middesk`

A [Terraform][1] provider for managing [Middesk][2] resources.

## Installation

Download and extract the [latest release](https://github.com/Mongey/terraform-provider-middesk/releases/latest) to
your [terraform plugin directory][third-party-plugins] (typically `~/.terraform.d/plugins/`) or define the plugin in the required_providers block.

```hcl
terraform {
  required_providers {
    middesk = {
      source = "Mongey/middesk"
    }
  }
}
```

## Example

Configure the provider directly, or set the ENV variables `MIDDESK_API_KEY`

```hcl
terraform {
  required_providers {
    middesk = {
      source = "Mongey/middesk"
    }
  }
}

provider "middesk" {
  api_key = "abc_123"
}

resource "middesk_webhook" "example" {
  name   = "My Application"
  url    = "https://example.com/middesk/callback"
  secret = "hunter2"
}
```

[1]: https://www.terraform.io
[2]: https://www.middesk.com
[third-party-plugins]: https://www.terraform.io/docs/configuration/providers.html#third-party-plugins
