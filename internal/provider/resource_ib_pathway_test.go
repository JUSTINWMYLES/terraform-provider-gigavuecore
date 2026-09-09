package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestIbPathwayResourceSchemaValidation verifies that the generated resource schema is valid.
func TestIbPathwayResourceSchemaValidation(t *testing.T) {
	r := &IbPathwayResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestIbPathwayResourceMetadata verifies that the generated resource reports the expected type name.
func TestIbPathwayResourceMetadata(t *testing.T) {
	r := &IbPathwayResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_ib_pathway" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_ib_pathway")
	}
}
