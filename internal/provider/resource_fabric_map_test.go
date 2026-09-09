package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFabricMapResourceSchemaValidation verifies that the generated resource schema is valid.
func TestFabricMapResourceSchemaValidation(t *testing.T) {
	r := &FabricMapResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFabricMapResourceMetadata verifies that the generated resource reports the expected type name.
func TestFabricMapResourceMetadata(t *testing.T) {
	r := &FabricMapResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_fabric_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_fabric_map")
	}
}
