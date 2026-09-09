package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTunnelApplicationResourceSchemaValidation verifies that the generated resource schema is valid.
func TestTunnelApplicationResourceSchemaValidation(t *testing.T) {
	r := &TunnelApplicationResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestTunnelApplicationResourceMetadata verifies that the generated resource reports the expected type name.
func TestTunnelApplicationResourceMetadata(t *testing.T) {
	r := &TunnelApplicationResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_tunnel_application" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_tunnel_application")
	}
}
