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

func testAccLocalUserResourceConfig(serverURL string, name string) string {
	return fmt.Sprintf("provider \"gigavuecore\" {\n  endpoint = \"%s\"\n  username = \"example\"\n  password = \"example\"\n}\nresource \"gigavuecore_local_user\" \"example\" {\n  account_status = \"%s\"\n  capability = \"admin\"\n  cluster_id = \"example\"\n  current_password = \"example\"\n  enabled = true\n  full_name = \"example\"\n  roles = [\"example\"]\n  user_pwd = \"example\"\n  username = \"example\"\n}\n", serverURL, name)
}

// newLocalUserResourceMockServer returns an httptest server that stubs the LocalUserResource CRUD endpoints.
// The server echoes request bodies so that create/update responses reflect the values sent by the test.
func newLocalUserResourceMockServer() *httptest.Server {
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
		id := strings.Trim(strings.TrimPrefix(r.URL.Path, "/system/localUsers"), "/")
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
			if _, ok := body["username"]; !ok {
				body["username"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["username"])
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
			if _, ok := body["username"]; !ok {
				body["username"] = "example-id"
			}
			id = fmt.Sprintf("%v", body["username"])
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
	mux.HandleFunc("/system/localUsers", handler0)
	mux.HandleFunc("/system/localUsers/", handler0)
	return httptest.NewServer(mux)
}

// TestAccLocalUserResourceLifecycle verifies create, update, delete, and import flows against a mock API.
func TestAccLocalUserResourceLifecycle(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	server := newLocalUserResourceMockServer()
	defer server.Close()
	resource.Test(t, resource.TestCase{ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){"gigavuecore": providerserver.NewProtocol6WithError(New())}, Steps: []resource.TestStep{resource.TestStep{Config: testAccLocalUserResourceConfig(server.URL, "accountDisabled"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_local_user.example", "username"), resource.TestCheckResourceAttr("gigavuecore_local_user.example", "account_status", "accountDisabled"))}, resource.TestStep{Config: testAccLocalUserResourceConfig(server.URL, "passwordDisabled"), Check: resource.ComposeAggregateTestCheckFunc(resource.TestCheckResourceAttrSet("gigavuecore_local_user.example", "username"), resource.TestCheckResourceAttr("gigavuecore_local_user.example", "account_status", "passwordDisabled"))}, resource.TestStep{ResourceName: "gigavuecore_local_user.example", ImportState: true, ImportStateId: "imported-username/imported-cluster_id"}}})
}
