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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) - System Diag PSU Information (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `ac_or_dc` (String) - AC or DC. NA stands for Not Available
* `busy_fault` (String) - PSU Busy Fault
* `capacity` (String) - Capacity (in Watts)
* `communication_fault` (String) - Communication, Logic and Memory Fault
* `hardware_revision` (String) - Hardware Revision
* `hardware_type` (String) - Hardware Type
* `input_current` (String) - Input Current (in Ampere)
* `input_over_voltage_fault` (String) - Input Over Voltage Fault
* `input_power` (String) - Input Power (in Watts)
* `input_under_voltage_fault` (String) - Input Under Voltage Fault
* `input_voltage` (String) - Input Voltage (in Volts)
* `max_permit_input_voltage` (String) - Maximum Input Voltage (in Volts)
* `min_permit_input_voltage` (String) - Minimum Input Voltage (in Volts)
* `off_fault` (String) - PSU Off Fault
* `output_current` (String) - Output Current (in Ampere)
* `output_over_current_fault` (String) - Output Over Current Fault
* `output_over_voltage_fault` (String) - Output Over Voltage Fault
* `output_power` (String) - Output Power (in Watts)
* `output_voltage` (String) - Output Voltage (in Volts)
* `product_code` (String) - Product Code
* `psu_fan_speed` (String) - PSU Fan Speed (RPM)
* `serial_number` (String) - Serial Number
* `slot_id` (String) - Power Module Slot ID
* `status` (String) - Current PSU Status
* `temperature1` (String) - Inlet Temperature (in Celsius)
* `temperature2` (String) - Heat Sink Temperature (in Celsius)
* `temperature3` (String) - Temperature (in Celsius)
* `temperature_fault` (String) - Temperature Fault

