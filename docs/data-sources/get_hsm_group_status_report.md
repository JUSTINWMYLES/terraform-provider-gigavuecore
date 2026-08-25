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
  type  = null
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
* `session_stats` (Attributes List, computed) (see [below for nested schema](#nestedatt--session_stats))

<a id="nestedatt--session_stats"></a>
### Nested Schema for `session_stats`

Read-Only:

* `engine_id` (String)
* `hsm_avg_rtt` (Number)
* `hsm_late` (Number)
* `hsm_req_sent` (Number)
* `hsm_resp_rcvd` (Number)
* `hsmc_avg_pkcs11_rtt` (Number)
* `hsmc_avg_rtt` (Number)
* `hsmc_avgq_delay` (Number)
* `hsmc_err_pkcs11` (Number)
* `hsmc_err_send_resp` (Number)
* `hsmc_errq_full` (Number)
* `hsmc_max_pkcs11_rtt` (Number)
* `hsmc_max_rtt` (Number)
* `hsmc_maxq_delay` (Number)
* `hsmc_min_pkcs11_rtt` (Number)
* `hsmc_min_rtt` (Number)
* `hsmc_minq_delay` (Number)
* `hsmc_req_rcvd` (Number)
* `hsmc_resp_sent` (Number)

