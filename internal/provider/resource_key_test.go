package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestKeyResourceSchemaValidation verifies that the generated resource schema is valid.
func TestKeyResourceSchemaValidation(t *testing.T) {
	r := &KeyResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestKeyResourceMetadata verifies that the generated resource reports the expected type name.
func TestKeyResourceMetadata(t *testing.T) {
	r := &KeyResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_key" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_key")
	}
}
