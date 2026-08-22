---
page_title: "gigavuecore_get_hsm_group_status_report Data Source - gigavuecore"
subcategory: ""
description: |-
  Load HSM Group status report
---

# gigavuecore_get_hsm_group_status_report Data Source

Load HSM Group status report

## Example Usage

```terraform
data "gigavuecore_get_hsm_group_status_report" "example" {
  alias = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - alias of target HSM Group
* `type` (String, optional) - report type, valid only if alias is specified

### Attributes

In addition to all arguments above, the following attributes are exported:

* `report` (String, computed) - status report
* `session_stats` (List(Object({engine_id, hsm_avg_rtt, hsm_late, hsm_req_sent, hsm_resp_rcvd, hsmc_avg_pkcs11_rtt, hsmc_avg_rtt, hsmc_avgq_delay, hsmc_err_pkcs11, hsmc_err_send_resp, hsmc_errq_full, hsmc_max_pkcs11_rtt, hsmc_max_rtt, hsmc_maxq_delay, hsmc_min_pkcs11_rtt, hsmc_min_rtt, hsmc_minq_delay, hsmc_req_rcvd, hsmc_resp_sent})), computed)

