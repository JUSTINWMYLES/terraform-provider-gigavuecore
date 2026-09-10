package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHeaderStripResourceSchemaValidation verifies that the generated resource schema is valid.
func TestHeaderStripResourceSchemaValidation(t *testing.T) {
	r := &HeaderStripResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestHeaderStripResourceMetadata verifies that the generated resource reports the expected type name.
func TestHeaderStripResourceMetadata(t *testing.T) {
	r := &HeaderStripResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_header_strip" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_header_strip")
	}
}
