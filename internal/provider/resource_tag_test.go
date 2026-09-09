package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTagResourceSchemaValidation verifies that the generated resource schema is valid.
func TestTagResourceSchemaValidation(t *testing.T) {
	r := &TagResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestTagResourceMetadata verifies that the generated resource reports the expected type name.
func TestTagResourceMetadata(t *testing.T) {
	r := &TagResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_tag" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_tag")
	}
}
