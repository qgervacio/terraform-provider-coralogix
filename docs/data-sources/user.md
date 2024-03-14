---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coralogix_user Data Source - terraform-provider-coralogix"
subcategory: ""
description: |-
  Coralogix User.
---

# coralogix_user (Data Source)

Coralogix User.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) User ID.

### Read-Only

- `active` (Boolean)
- `emails` (Attributes Set) (see [below for nested schema](#nestedatt--emails))
- `groups` (Set of String)
- `name` (Attributes) (see [below for nested schema](#nestedatt--name))
- `user_name` (String) User name.

<a id="nestedatt--emails"></a>
### Nested Schema for `emails`

Read-Only:

- `primary` (Boolean)
- `type` (String)
- `value` (String)


<a id="nestedatt--name"></a>
### Nested Schema for `name`

Read-Only:

- `family_name` (String)
- `given_name` (String)