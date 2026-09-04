package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetworkLagResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNetworkLagResourceSchemaValidation(t *testing.T) {
	r := &NetworkLagResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNetworkLagResourceMetadata verifies that the generated resource reports the expected type name.
func TestNetworkLagResourceMetadata(t *testing.T) {
	r := &NetworkLagResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_network_lag" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_network_lag")
	}
}
