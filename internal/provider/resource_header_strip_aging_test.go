package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHeaderStripAgingResourceSchemaValidation verifies that the generated resource schema is valid.
func TestHeaderStripAgingResourceSchemaValidation(t *testing.T) {
	r := &HeaderStripAgingResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestHeaderStripAgingResourceMetadata verifies that the generated resource reports the expected type name.
func TestHeaderStripAgingResourceMetadata(t *testing.T) {
	r := &HeaderStripAgingResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_header_strip_aging" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_header_strip_aging")
	}
}
