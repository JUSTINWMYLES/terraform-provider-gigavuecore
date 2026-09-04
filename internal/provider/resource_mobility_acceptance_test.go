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

func testAccMobilityResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_mobility\" \"example\" {\n  deployed = true\n  health_state = \"%s\"\n  site_name = \"example\"\n  site_tag = \"example\"\n  sites = \"example\"\n  solution_alias = \"example\"\n  solution_type = \"NON_CUPS\"\n  tags = [{\n    tag_key = \"example\"\n    tag_values = [\"example\"]\n  }]\n  traffic_policies = {\n    for5_g = {\n      gtp_flow_timeout = 1\n      gtp_persistence = {\n        enabled = true\n        file_age_timeout = 10\n        interval = 10\n        restart_age_time = 10\n      }\n      load_balancing = {\n        app_type = \"flow5g\"\n        hashing_key = \"supi\"\n      }\n      overlap_mode = true\n      sampling = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          rules = [{\n            comment = \"example\"\n            control_plane_percentage = 0\n            dnn = \"example\"\n            gpsi = \"*\"\n            nas_5_qi = \"0\"\n            nci = \"*\"\n            nsiid = \"0\"\n            pei = \"*\"\n            plmn_id = \"*\"\n            supi = \"*\"\n            tac = \"*\"\n            user_plane_percentage = 0\n          }]\n          source_group_id = \"example\"\n          tags = [{\n            tag_key = \"example\"\n            tag_values = [\"example\"]\n          }]\n          tool = \"example\"\n        }]\n      }\n      whitelisting = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          rules = [{\n            dnn = \"example\"\n            type = \"example\"\n            whitelist_databases = [\"example\"]\n          }]\n          source_group_id = \"example\"\n          tags = [{\n            tag_key = \"example\"\n            tag_values = [\"example\"]\n          }]\n          tool = \"example\"\n        }]\n        multi_whitelists = [\"example\"]\n        white_list_alias = \"example\"\n      }\n    }\n    for_lte = {\n      gtp_flow_timeout = 1\n      gtp_persistence = {\n        enabled = true\n        file_age_timeout = 10\n        interval = 10\n        restart_age_time = 10\n      }\n      overlap_mode = true\n      sampling = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          rules = [{\n            apn = \"example\"\n            comment = \"example\"\n            control_plane_percentage = 0\n            eci = \"abc1\"\n            imei = \"*\"\n            imsi = \"*\"\n            interface = \"Gn\"\n            msisdn = \"*\"\n            nas_5_qi = \"0\"\n            nci = \"*\"\n            periodic_recalc = true\n            plmn_id = \"*\"\n            qci = 0\n            snssai = \"0\"\n            tac = \"abc1\"\n            tac_5_g = \"*\"\n            user_plane_percentage = 0\n            version = \"any\"\n          }]\n          source_group_id = \"example\"\n          tags = [{\n            tag_key = \"example\"\n            tag_values = [\"example\"]\n          }]\n          tool = \"example\"\n        }]\n      }\n      whitelisting = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          rules = [{\n            apn = \"example\"\n            interface = \"Gn\"\n            type = \"example\"\n            version = \"v1\"\n            whitelist_databases = [\"example\"]\n          }]\n          source_group_id = \"example\"\n          tool = \"example\"\n        }]\n        multi_whitelists = [\"example\"]\n        white_list_alias = \"example\"\n      }\n    }\n    for_non_cups_lte = {\n      flowfiltering = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          drop_rules = [{\n            imei = \"*\"\n            imsi = \"*\"\n            interface = \"Gn\"\n            msisdn = \"*\"\n            version = \"v1\"\n          }]\n          pass_rules = [{\n            imei = \"*\"\n            imsi = \"*\"\n            interface = \"Gn\"\n            msisdn = \"*\"\n            version = \"v1\"\n          }]\n          source_group_id = \"example\"\n          tags = [{\n            tag_key = \"example\"\n            tag_values = [\"example\"]\n          }]\n          tool = \"example\"\n        }]\n      }\n      gtp_flow_timeout = 1\n      gtp_persistence = {\n        enabled = true\n        file_age_timeout = 10\n        interval = 10\n        restart_age_time = 10\n      }\n      load_balancing = {\n        app_type = \"gtp\"\n        hashing_key = \"imsi\"\n      }\n      overlap_mode = true\n      sampling = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          rules = [{\n            apn = \"example\"\n            comment = \"example\"\n            eci = \"abc1\"\n            gtp_sample_percentage = 0\n            imei = \"*\"\n            imsi = \"*\"\n            interface = \"Gn\"\n            msisdn = \"*\"\n            periodic_recalc = true\n            plmn_id = \"*\"\n            qci = 0\n            tac = \"abc1\"\n            version = \"any\"\n          }]\n          source_group_id = \"example\"\n          tags = [{\n            tag_key = \"example\"\n            tag_values = [\"example\"]\n          }]\n          tool = \"example\"\n        }]\n      }\n      whitelisting = {\n        flow_maps = [{\n          alias = \"example\"\n          comment = \"example\"\n          rules = [{\n            apn = \"example\"\n            interface = \"Gn\"\n            type = \"example\"\n            version = \"v1\"\n            whitelist_databases = [\"example\"]\n          }]\n          source_group_id = \"example\"\n          tags = [{\n            tag_key = \"example\"\n            tag_values = [\"example\"]\n          }]\n          tool = \"example\"\n        }]\n        multi_whitelists = [\"example\"]\n        white_list_alias = \"example\"\n      }\n    }\n  }\n}\n", serverURL, name)
}

// newMobilityResourceMockServer returns an httptest server that stubs the MobilityResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newMobilityResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/intent/mobility"), "/")
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
			if _, ok := body["solutionAlias"]; !ok {
				body["solutionAlias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["solutionAlias"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(207)
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
			if _, ok := body["solutionAlias"]; !ok {
				body["solutionAlias"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["solutionAlias"])
			state0[id] = body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(207)
			_ = json.NewEncoder(w).Encode(body)
			lastKey0 = id
			return
		case http.MethodDelete:
			delete(state0, id)
			w.WriteHeader(207)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
	mux.HandleFunc("/intent/mobility", handler0)
	mux.HandleFunc("/intent/mobility/", handler0)
	return httptest.NewServer(mux)
}

// TestAccMobilityResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccMobilityResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newMobilityResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccMobilityResourceConfig(server.URL, "green"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_mobility.example", "solution_alias"), resource.TestCheckResourceAttr("gigavuecore_mobility.example", "health_state", "green"))}, resource.TestStep{Config: testAccMobilityResourceConfig(server.URL, "yellow"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_mobility.example", "solution_alias"), resource.TestCheckResourceAttr("gigavuecore_mobility.example", "health_state", "yellow"))}, resource.TestStep{ResourceName: "gigavuecore_mobility.example", ImportState: true, ImportStateId: "imported-solution_alias"}}})
}
