---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coralogix_tco_policy_override Resource - terraform-provider-coralogix"
subcategory: ""
description: |-
  Coralogix TCO-Policy-Override. For more information - https://coralogix.com/docs/tco-optimizer-api/#policy-overrides .
---

# coralogix_tco_policy_override (Resource)

Coralogix TCO-Policy-Override. For more information - https://coralogix.com/docs/tco-optimizer-api/#policy-overrides .

## Example Usage

```hcl
resource "coralogix_tco_policy_override" "tco_policy" {
  priority         = "medium"
  severity         = "debug"
  application_name = "prod"
  subsystem_name   = "mobile"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_name` (String) The application to apply the policy on. Applies the policy on all the applications by default.
- `priority` (String) The policy-override priority. Can be one of ["high" "medium" "low" "block"].
- `severity` (String) The severity to apply the policy on. Can be one of ["debug" "verbose" "info" "warning" "error" "critical"].
- `subsystem_name` (String) The subsystem to apply the policy on. Applies the policy on all the subsystems by default.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

