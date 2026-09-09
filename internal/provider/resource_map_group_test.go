package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMapGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestMapGroupResourceSchemaValidation(t *testing.T) {
	r := &MapGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestMapGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestMapGroupResourceMetadata(t *testing.T) {
	r := &MapGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_map_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_map_group")
	}
}
