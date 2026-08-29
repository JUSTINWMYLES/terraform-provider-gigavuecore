---
page_title: "gigavuecore_get_sys_info Data Source - gigavuecore"
subcategory: ""
description: |-
  Cms build info for FmHa
---

# gigavuecore_get_sys_info Data Source

Cms build info for FmHa

## Example Usage

```terraform
data "gigavuecore_get_sys_info" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `additional_dns` (String, computed) - Additional Dns of build
* `additional_domain_name` (String, computed) - Additional Domain Name of build
* `boot_image` (String, computed) - Boot Image of build
* `build_date` (String, computed) - Build Date of build
* `build_id` (String, computed) - Build Id of build
* `default_gateway` (String, computed) - Default Gateway of build
* `default_ipv6_gateway` (String, computed) - Default Ipv6 Gateway of build
* `dns_name` (String, computed) - DNS Name of build
* `domain_name` (String, computed) - Domain Name of build
* `fm_time` (String, computed) - FM time of build
* `fm_uptime` (String, computed) - FM Up time of build
* `host_id` (String, computed) - Host Id of build
* `hostname` (String, computed) - DNS name or IP Address of build
* `mac` (String, computed) - Mac of build
* `ntp_ip` (String, computed) - ntp Ip of build
* `platform` (String, computed) - Platform of build
* `primary_dns` (String, computed) - Primary Dns of build
* `primary_ip` (String, computed) - Primary Ip of build
* `primary_ip_mask` (String, computed) - Primary Ip Mask of build
* `primary_ip_mask_len` (String, computed) - Primary Ip Mask Len of build
* `primary_ipv6` (String, computed) - Primary Ipv6 of build
* `primary_ipv6_mask` (String, computed) - Primary Ipv6 Mask of build
* `product_name` (String, computed) - Product Name of build
* `system_timestamp` (Number, computed) - System time stamp of build
* `system_timestamp_utc` (String, computed) - System time stamp Utc of build
* `version` (String, computed) - Version of build


