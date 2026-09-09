package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetworkProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNetworkProfileResourceSchemaValidation(t *testing.T) {
	r := &NetworkProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNetworkProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestNetworkProfileResourceMetadata(t *testing.T) {
	r := &NetworkProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_network_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_network_profile")
	}
}
