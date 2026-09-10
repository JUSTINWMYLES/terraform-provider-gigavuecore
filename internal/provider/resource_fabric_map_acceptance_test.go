package provider

import (
	json "encoding/json"
	"fmt"
	"io"
	http "net/http"
	httptest "net/http/httptest"
	"strings"
	"sync"
	"testing"
)
import (
	providerserver "github.com/hashicorp/terraform-plugin-framework/providerserver"
	tfprotov6 "github.com/hashicorp/terraform-plugin-go/tfprotov6"
	resource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccFabricMapResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_fabric_map\" \"example\" {\n  afm_map = {\n    ap_rules = {\n      drop_rules = [{\n        application_profile = \"example\"\n        rule_id = 1\n      }]\n      pass_rules = [{\n        application_profile = \"example\"\n        rule_id = 1\n      }]\n    }\n    cluster_id = \"example\"\n    comment = \"example\"\n    dst_ports = [\"example\"]\n    egress_gigastream = [\"example\"]\n    enable = true\n    encap_tunnel = \"example\"\n    flex_inline = {\n      a_to_b = {\n        ib_pathway = \"example\"\n        tools = [\"example\"]\n        type = \"bypass\"\n      }\n      b_to_a = {\n        ib_pathway = \"example\"\n        tools = [\"example\"]\n        type = \"bypass\"\n      }\n      oob_copy = [{\n        direction = \"aToB\"\n        dst_ports = [\"example\"]\n        src_ports = [\"example\"]\n        tag = {\n          type = \"none\"\n        }\n      }]\n      svt_mode = true\n      svt_tag = 0\n      tag = {\n        tag_protocol_id = \"0x8100\"\n        type = \"auto\"\n        vlan_id = 0\n      }\n    }\n    flex_inline_failover = \"bypass\"\n    flex_inline_vlan_id = 1\n    flow_rules = {\n      drop_rules = [{\n        gtp = {\n          imei = \"*\"\n          imsi = \"*\"\n          interface = \"Gn\"\n          msisdn = \"*\"\n          version = \"any\"\n        }\n        rule_id = 1\n      }]\n      pass_rules = [{\n        gtp = {\n          imei = \"*\"\n          imsi = \"*\"\n          interface = \"Gn\"\n          msisdn = \"*\"\n          version = \"any\"\n        }\n        rule_id = 1\n      }]\n    }\n    flow_sample5_g_overlap_rules = {\n      pass_rules = [{\n        comment = \"example\"\n        flow5_g = {\n          dnn = \"example\"\n          gpsi = \"*\"\n          nas_5_qi = \"0\"\n          nci = \"*\"\n          nsiid = \"0\"\n          pei = \"*\"\n          plmn_id = \"*\"\n          supi = \"*\"\n          tac = \"*\"\n        }\n        percentage = 0\n        rule_id = 1\n      }]\n    }\n    flow_sample5_g_rules = {\n      pass_rules = [{\n        comment = \"example\"\n        flow5_g = {\n          dnn = \"example\"\n          gpsi = \"*\"\n          nci = \"a1b2c3d4e\"\n          nsiid = \"0\"\n          pei = \"*\"\n          plmn_id = \"123.45\"\n          supi = \"*\"\n          tac = \"*\"\n        }\n        percentage = 0\n        priority = 1\n        rule_id = 1\n      }]\n    }\n    flow_sample_diameter_rules = {\n      pass_rules = [{\n        diameter = {\n          user_name = \"*\"\n        }\n        interface = \"s6a\"\n        percentage = 0\n        rule_id = 1\n      }]\n    }\n    flow_sample_overlap_rules = {\n      pass_rules = [{\n        comment = \"example\"\n        gtp = {\n          apn = \"example\"\n          eci = \"a1b2c3d4\"\n          imei = \"*\"\n          imsi = \"*\"\n          interface = \"Gn\"\n          msisdn = \"*\"\n          nas_5_qi = \"0\"\n          nci = \"*\"\n          plmn_id = \"123.45\"\n          qci = 0\n          snssai = \"0\"\n          tac = \"abc1\"\n          tac_5_g = \"*\"\n          version = \"any\"\n        }\n        percentage = 0\n        periodic_recalc = true\n        priority = 1\n        rule_id = 1\n      }]\n    }\n    flow_sample_rules = {\n      pass_rules = [{\n        comment = \"example\"\n        gtp = {\n          apn = \"example\"\n          eci = \"a1b2c3d4\"\n          imei = \"*\"\n          imsi = \"*\"\n          interface = \"Gn\"\n          msisdn = \"*\"\n          nas_5_qi = \"0\"\n          nci = \"*\"\n          plmn_id = \"123.45\"\n          qci = 0\n          snssai = \"0\"\n          tac = \"abc1\"\n          tac_5_g = \"*\"\n          version = \"any\"\n        }\n        percentage = 0\n        periodic_recalc = true\n        priority = 1\n        rule_id = 1\n      }]\n    }\n    flow_sample_sip_rules = {\n      pass_rules = [{\n        percentage = 0\n        rule_id = 1\n        sip = {\n          callee_id = \"example\"\n          callee_id_range = {\n            max_value = \"example\"\n            value = \"example\"\n          }\n          caller_id = \"example\"\n          caller_id_range = {\n            max_value = \"example\"\n            value = \"example\"\n          }\n          id_range = {\n            max_value = \"example\"\n            value = \"example\"\n          }\n        }\n      }]\n    }\n    flow_whitelist5_g_overlap_rules = {\n      dnn = \"example\"\n      type = \"example\"\n    }\n    flow_whitelist5_g_rules = {\n      dnn = \"example\"\n      type = \"example\"\n      whitelist_databases = [\"example\"]\n    }\n    flow_whitelist_overlap_rules = {\n      pass_rules = [{\n        flow5_g = {\n          dnn = \"example\"\n          type = \"example\"\n          whitelist_databases = [\"example\"]\n        }\n        gtp = {\n          apn = \"example\"\n          interface = \"Gn\"\n          type = \"example\"\n          version = \"v1\"\n          whitelist_databases = [\"example\"]\n        }\n        rule_id = 1\n        sip = {\n          type = \"all\"\n        }\n      }]\n    }\n    flow_whitelist_rules = {\n      pass_rules = [{\n        flow5_g = {\n          dnn = \"example\"\n          type = \"example\"\n          whitelist_databases = [\"example\"]\n        }\n        gtp = {\n          apn = \"example\"\n          interface = \"Gn\"\n          type = \"example\"\n          version = \"v1\"\n          whitelist_databases = [\"example\"]\n        }\n        rule_id = 1\n        sip = {\n          type = \"all\"\n        }\n      }]\n    }\n    fstype = {\n      offset = 1\n      timer = 15\n      type = \"_default\"\n    }\n    gs_rules = {\n      drop_rules = null\n      pass_rules = null\n    }\n    gsop = \"example\"\n    inline_traffic_path = \"normal\"\n    inline_traffic_type = \"symmetric\"\n    ip_rewrite = {\n      dst_ip = \"example\"\n      src_ip = \"example\"\n    }\n    mod_time = 0\n    null_dst_port = true\n    order = 0\n    rewrite = {\n      dst_mac = \"example\"\n      src_mac = \"example\"\n    }\n    roles = {\n      editors = [\"example\"]\n      listeners = [\"example\"]\n      owners = [\"example\"]\n      viewers = [\"example\"]\n    }\n    rule_matching = \"normal\"\n    rules = {\n      drop_rules = null\n      pass_rules = null\n    }\n    rx_cluster_ports = [\"example\"]\n    src_ports = [\"example\"]\n    sub_type = \"byRule\"\n    traffic_type = \"control\"\n    tx_cluster_ports = [\"example\"]\n    type = \"regular\"\n    vlan_tag = {\n      tag_protocol_id = \"0x8100\"\n      vlan_action = \"add\"\n      vlan_id = 0\n    }\n  }\n  alias = \"example\"\n  ap_rules = {\n    drop_rules = [{\n      application_profile = \"example\"\n      rule_id = 1\n    }]\n    pass_rules = [{\n      application_profile = \"example\"\n      rule_id = 1\n    }]\n  }\n  child_map_aliases = [\"example\"]\n  cluster_id = \"%s\"\n  comment = \"example\"\n  config_status = \"example\"\n  decap_aliases = [\"example\"]\n  dst_ports = [\"example\"]\n  egress_gigastream = [\"example\"]\n  enable = true\n  encap_aliases = [\"example\"]\n  encap_tunnel = \"example\"\n  error_message = \"example\"\n  flex_inline = {\n    a_to_b = {\n      ib_pathway = \"example\"\n      tools = [\"example\"]\n      type = \"bypass\"\n    }\n    b_to_a = {\n      ib_pathway = \"example\"\n      tools = [\"example\"]\n      type = \"bypass\"\n    }\n    oob_copy = [{\n      direction = \"aToB\"\n      dst_ports = [\"example\"]\n      src_ports = [\"example\"]\n      tag = {\n        type = \"none\"\n      }\n    }]\n    svt_mode = true\n    svt_tag = 0\n    tag = {\n      tag_protocol_id = \"0x8100\"\n      type = \"auto\"\n      vlan_id = 0\n    }\n  }\n  flex_inline_failover = \"bypass\"\n  flex_inline_vlan_id = 1\n  flow_rules = {\n    drop_rules = [{\n      gtp = {\n        imei = \"*\"\n        imsi = \"*\"\n        interface = \"Gn\"\n        msisdn = \"*\"\n        version = \"any\"\n      }\n      rule_id = 1\n    }]\n    pass_rules = [{\n      gtp = {\n        imei = \"*\"\n        imsi = \"*\"\n        interface = \"Gn\"\n        msisdn = \"*\"\n        version = \"any\"\n      }\n      rule_id = 1\n    }]\n  }\n  flow_sample5_g_overlap_rules = {\n    pass_rules = [{\n      comment = \"example\"\n      flow5_g = {\n        dnn = \"example\"\n        gpsi = \"*\"\n        nas_5_qi = \"0\"\n        nci = \"*\"\n        nsiid = \"0\"\n        pei = \"*\"\n        plmn_id = \"*\"\n        supi = \"*\"\n        tac = \"*\"\n      }\n      percentage = 0\n      rule_id = 1\n    }]\n  }\n  flow_sample5_g_rules = {\n    pass_rules = [{\n      comment = \"example\"\n      flow5_g = {\n        dnn = \"example\"\n        gpsi = \"*\"\n        nci = \"a1b2c3d4e\"\n        nsiid = \"0\"\n        pei = \"*\"\n        plmn_id = \"123.45\"\n        supi = \"*\"\n        tac = \"*\"\n      }\n      percentage = 0\n      priority = 1\n      rule_id = 1\n    }]\n  }\n  flow_sample_diameter_rules = {\n    pass_rules = [{\n      diameter = {\n        user_name = \"*\"\n      }\n      interface = \"s6a\"\n      percentage = 0\n      rule_id = 1\n    }]\n  }\n  flow_sample_overlap_rules = {\n    pass_rules = [{\n      comment = \"example\"\n      gtp = {\n        apn = \"example\"\n        eci = \"a1b2c3d4\"\n        imei = \"*\"\n        imsi = \"*\"\n        interface = \"Gn\"\n        msisdn = \"*\"\n        nas_5_qi = \"0\"\n        nci = \"*\"\n        plmn_id = \"123.45\"\n        qci = 0\n        snssai = \"0\"\n        tac = \"abc1\"\n        tac_5_g = \"*\"\n        version = \"any\"\n      }\n      percentage = 0\n      periodic_recalc = true\n      priority = 1\n      rule_id = 1\n    }]\n  }\n  flow_sample_rules = {\n    pass_rules = [{\n      comment = \"example\"\n      gtp = {\n        apn = \"example\"\n        eci = \"a1b2c3d4\"\n        imei = \"*\"\n        imsi = \"*\"\n        interface = \"Gn\"\n        msisdn = \"*\"\n        nas_5_qi = \"0\"\n        nci = \"*\"\n        plmn_id = \"123.45\"\n        qci = 0\n        snssai = \"0\"\n        tac = \"abc1\"\n        tac_5_g = \"*\"\n        version = \"any\"\n      }\n      percentage = 0\n      periodic_recalc = true\n      priority = 1\n      rule_id = 1\n    }]\n  }\n  flow_sample_sip_rules = {\n    pass_rules = [{\n      percentage = 0\n      rule_id = 1\n      sip = {\n        callee_id = \"example\"\n        callee_id_range = {\n          max_value = \"example\"\n          value = \"example\"\n        }\n        caller_id = \"example\"\n        caller_id_range = {\n          max_value = \"example\"\n          value = \"example\"\n        }\n        id_range = {\n          max_value = \"example\"\n          value = \"example\"\n        }\n      }\n    }]\n  }\n  flow_whitelist5_g_overlap_rules = {\n    dnn = \"example\"\n    type = \"example\"\n  }\n  flow_whitelist5_g_rules = {\n    dnn = \"example\"\n    type = \"example\"\n    whitelist_databases = [\"example\"]\n  }\n  flow_whitelist_overlap_rules = {\n    pass_rules = [{\n      flow5_g = {\n        dnn = \"example\"\n        type = \"example\"\n        whitelist_databases = [\"example\"]\n      }\n      gtp = {\n        apn = \"example\"\n        interface = \"Gn\"\n        type = \"example\"\n        version = \"v1\"\n        whitelist_databases = [\"example\"]\n      }\n      rule_id = 1\n      sip = {\n        type = \"all\"\n      }\n    }]\n  }\n  flow_whitelist_rules = {\n    pass_rules = [{\n      flow5_g = {\n        dnn = \"example\"\n        type = \"example\"\n        whitelist_databases = [\"example\"]\n      }\n      gtp = {\n        apn = \"example\"\n        interface = \"Gn\"\n        type = \"example\"\n        version = \"v1\"\n        whitelist_databases = [\"example\"]\n      }\n      rule_id = 1\n      sip = {\n        type = \"all\"\n      }\n    }]\n  }\n  fstype = {\n    offset = 1\n    timer = 15\n    type = \"_default\"\n  }\n  gs_rules = {\n    drop_rules = null\n    pass_rules = null\n  }\n  gsop = \"example\"\n  health_state = \"green\"\n  health_state_reasons = [{\n    message = \"example\"\n    severity = \"green\"\n    traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n  }]\n  inline_traffic_path = \"normal\"\n  inline_traffic_type = \"symmetric\"\n  ip_rewrite = {\n    dst_ip = \"example\"\n    src_ip = \"example\"\n  }\n  null_dst_port = true\n  order = 0\n  parent_map_aliases = [\"example\"]\n  rewrite = {\n    dst_mac = \"example\"\n    src_mac = \"example\"\n  }\n  roles = {\n    editors = [\"example\"]\n    listeners = [\"example\"]\n    owners = [\"example\"]\n    viewers = [\"example\"]\n  }\n  rule_matching = \"normal\"\n  rules = {\n    drop_rules = null\n    pass_rules = null\n  }\n  src_ports = [\"example\"]\n  sub_type = \"byRule\"\n  traffic_type = \"control\"\n  type = \"regular\"\n  updated_time = 1.0\n  vlan_tag = {\n    tag_protocol_id = \"0x8100\"\n    vlan_action = \"add\"\n    vlan_id = 0\n  }\n}\n", serverURL, name)
}

// newFabricMapResourceMockServer returns an httptest server that stubs the FabricMapResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newFabricMapResourceMockServer() *httptest.Server {
	mux := http.NewServeMux()
	state0 := make(map[string]map[string]interface{})
	var mu0 sync.Mutex
	lastKey0 := ""
	handler0 := func(w http.ResponseWriter, r *http.Request) {
		if bu0, bp0, bok0 := r.BasicAuth(); !bok0 || bu0 != "example" || bp0 != "example" {
			http.Error(w, "missing basic auth credential", http.StatusUnauthorized)
			return
		}
		mu0.Lock()
		defer mu0.Unlock()
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/fabricMaps"), "/")
		if id == "" {
			id = "example-id"
		}
		switch r.Method {
		case http.MethodPost:
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["alias"]; !ok {
				body["alias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["alias"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			_ = json.NewEncoder(w).Encode(body)
			lastKey0 = id
			return
		case http.MethodGet:
			body, ok := state0[id]
			if !ok && lastKey0 != "" {
				body = state0[lastKey0]
			}
			if body == nil {
				http.NotFound(w, r)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			return
		case http.MethodPut, http.MethodPatch:
			body := make(map[string]interface{})
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if _, ok := body["alias"]; !ok {
				body["alias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["alias"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			_ = json.NewEncoder(w).Encode(body)
			lastKey0 = id
			return
		case http.MethodDelete:
			delete(state0, id)
			w.WriteHeader(204)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/fabricMaps", handler0)
	mux.HandleFunc("/fabricMaps/", handler0)
	return httptest.NewServer(mux)
}

// TestAccFabricMapResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccFabricMapResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newFabricMapResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccFabricMapResourceConfig(server.URL, "example"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_fabric_map.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_fabric_map.example", "cluster_id", "example"))}, resource.TestStep{Config: testAccFabricMapResourceConfig(server.URL, "updated"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_fabric_map.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_fabric_map.example", "cluster_id", "updated"))}, resource.TestStep{ResourceName: "gigavuecore_fabric_map.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
