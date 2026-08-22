---
page_title: "gigavuecore_get_system_diag_psu Data Source - gigavuecore"
subcategory: ""
description: |-
  get system diagnostics PSU information
---

# gigavuecore_get_system_diag_psu Data Source

get system diagnostics PSU information

## Example Usage

```terraform
data "gigavuecore_get_system_diag_psu" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({ac_or_dc, busy_fault, capacity, communication_fault, hardware_revision, hardware_type, input_current, input_over_voltage_fault, input_power, input_under_voltage_fault, input_voltage, max_permit_input_voltage, min_permit_input_voltage, off_fault, output_current, output_over_current_fault, output_over_voltage_fault, output_power, output_voltage, product_code, psu_fan_speed, serial_number, slot_id, status, temperature1, temperature2, temperature3, temperature_fault})), computed) - System Diag PSU Information

