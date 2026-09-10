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

func testAccFlexInlineResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_flex_inline\" \"example\" {\n  alias = \"example\"\n  cluster_configs = [{\n    cluster_id = \"example\"\n    export_criteria = {\n      lsb = 0\n    }\n    export_type = \"ipBased\"\n    ib_pathway = \"example\"\n    source = {\n      alias = \"example\"\n      type = \"IN\"\n    }\n  }]\n  health_state = \"%s\"\n  health_state_reasons = [{\n    message = \"example\"\n    severity = \"green\"\n    traffic_health_state_computation_type = \"PORT_LOW_UTIL\"\n  }]\n  keep_on_device = true\n  resilient_config = {\n    side_a = \"source\"\n    side_b = \"source\"\n  }\n  target_traffic_path = \"drop\"\n}\n", serverURL, name)
}

// newFlexInlineResourceMockServer returns an httptest server that stubs the FlexInlineResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newFlexInlineResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/flexInline"), "/")
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
	mux.HandleFunc("/flexInline", handler0)
	mux.HandleFunc("/flexInline/", handler0)
	return httptest.NewServer(mux)
}

// TestAccFlexInlineResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccFlexInlineResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newFlexInlineResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccFlexInlineResourceConfig(server.URL, "green"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_flex_inline.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_flex_inline.example", "health_state", "green"))}, resource.TestStep{Config: testAccFlexInlineResourceConfig(server.URL, "yellow"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_flex_inline.example", "alias"), resource.TestCheckResourceAttr("gigavuecore_flex_inline.example", "health_state", "yellow"))}, resource.TestStep{ResourceName: "gigavuecore_flex_inline.example", ImportState: true, ImportStateId: "imported-alias"}}})
}
