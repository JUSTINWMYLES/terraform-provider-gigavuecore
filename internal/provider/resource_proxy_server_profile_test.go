package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestProxyServerProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestProxyServerProfileResourceSchemaValidation(t *testing.T) {
	r := &ProxyServerProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestProxyServerProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestProxyServerProfileResourceMetadata(t *testing.T) {
	r := &ProxyServerProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_proxy_server_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_proxy_server_profile")
	}
}
