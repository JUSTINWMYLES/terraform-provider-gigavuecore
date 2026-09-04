package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestKeyMapResourceSchemaValidation verifies that the generated resource schema is valid.
func TestKeyMapResourceSchemaValidation(t *testing.T) {
	r := &KeyMapResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestKeyMapResourceMetadata verifies that the generated resource reports the expected type name.
func TestKeyMapResourceMetadata(t *testing.T) {
	r := &KeyMapResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_key_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_key_map")
	}
}
