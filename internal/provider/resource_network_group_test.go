package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetworkGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNetworkGroupResourceSchemaValidation(t *testing.T) {
	r := &NetworkGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNetworkGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestNetworkGroupResourceMetadata(t *testing.T) {
	r := &NetworkGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_network_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_network_group")
	}
}
