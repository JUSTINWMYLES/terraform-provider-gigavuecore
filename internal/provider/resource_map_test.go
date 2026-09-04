package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMapResourceSchemaValidation verifies that the generated resource schema is valid.
func TestMapResourceSchemaValidation(t *testing.T) {
	r := &MapResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestMapResourceMetadata verifies that the generated resource reports the expected type name.
func TestMapResourceMetadata(t *testing.T) {
	r := &MapResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_map")
	}
}
