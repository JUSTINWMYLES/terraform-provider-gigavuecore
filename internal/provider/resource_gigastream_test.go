package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGigastreamResourceSchemaValidation verifies that the generated resource schema is valid.
func TestGigastreamResourceSchemaValidation(t *testing.T) {
	r := &GigastreamResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGigastreamResourceMetadata verifies that the generated resource reports the expected type name.
func TestGigastreamResourceMetadata(t *testing.T) {
	r := &GigastreamResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_gigastream" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_gigastream")
	}
}
