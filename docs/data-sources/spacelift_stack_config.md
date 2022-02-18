---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "utils_spacelift_stack_config Data Source - terraform-provider-utils"
subcategory: ""
description: |-
  The spacelift_stack_config data source accepts a list of stack config file names and returns a map of Spacelift stack configurations.
---

# utils_spacelift_stack_config (Data Source)

The `spacelift_stack_config` data source accepts a list of stack config file names and returns a map of Spacelift stack configurations.

## Example Usage

```terraform
locals {
  result = yamldecode(data.utils_spacelift_stack_config.example.output)
}

data "utils_spacelift_stack_config" "example" {
  process_component_deps     = true
  process_stack_deps         = true
  process_imports            = true
  stack_config_path_template = "stacks/%s.yaml"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **stack_config_path_template** (String) Stack config path template.

### Optional

- **base_path** (String) Stack config base path.
- **id** (String) The ID of this resource.
- **input** (List of String) A list of stack config file names.
- **process_component_deps** (Boolean) A boolean flag to enable/disable processing config dependencies for the components.
- **process_imports** (Boolean) A boolean flag to enable/disable processing stack imports.
- **process_stack_deps** (Boolean) A boolean flag to enable/disable processing all stack dependencies for the components.

### Read-Only

- **output** (String) A map of Spacelift stack configurations.

