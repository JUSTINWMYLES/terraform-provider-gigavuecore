---
page_title: "gigavuecore_get_system_diag_psu_per_slot Data Source - gigavuecore"
subcategory: ""
description: |-
  get system diagnostics PSU information per Slot Id
---

# gigavuecore_get_system_diag_psu_per_slot Data Source

get system diagnostics PSU information per Slot Id

## Example Usage

```terraform
data "gigavuecore_get_system_diag_psu_per_slot" "example" {
  slot_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `slot_id` (String, required) - Power Module Slot ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `ac_or_dc` (String, computed) - AC or DC. NA stands for Not Available
* `busy_fault` (String, computed) - PSU Busy Fault
* `capacity` (String, computed) - Capacity (in Watts)
* `communication_fault` (String, computed) - Communication, Logic and Memory Fault
* `hardware_revision` (String, computed) - Hardware Revision
* `hardware_type` (String, computed) - Hardware Type
* `input_current` (String, computed) - Input Current (in Ampere)
* `input_over_voltage_fault` (String, computed) - Input Over Voltage Fault
* `input_power` (String, computed) - Input Power (in Watts)
* `input_under_voltage_fault` (String, computed) - Input Under Voltage Fault
* `input_voltage` (String, computed) - Input Voltage (in Volts)
* `max_permit_input_voltage` (String, computed) - Maximum Input Voltage (in Volts)
* `min_permit_input_voltage` (String, computed) - Minimum Input Voltage (in Volts)
* `off_fault` (String, computed) - PSU Off Fault
* `output_current` (String, computed) - Output Current (in Ampere)
* `output_over_current_fault` (String, computed) - Output Over Current Fault
* `output_over_voltage_fault` (String, computed) - Output Over Voltage Fault
* `output_power` (String, computed) - Output Power (in Watts)
* `output_voltage` (String, computed) - Output Voltage (in Volts)
* `product_code` (String, computed) - Product Code
* `psu_fan_speed` (String, computed) - PSU Fan Speed (RPM)
* `serial_number` (String, computed) - Serial Number
* `status` (String, computed) - Current PSU Status
* `temperature1` (String, computed) - Inlet Temperature (in Celsius)
* `temperature2` (String, computed) - Heat Sink Temperature (in Celsius)
* `temperature3` (String, computed) - Temperature (in Celsius)
* `temperature_fault` (String, computed) - Temperature Fault


