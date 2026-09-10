package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFlexInlineResourceSchemaValidation verifies that the generated resource schema is valid.
func TestFlexInlineResourceSchemaValidation(t *testing.T) {
	r := &FlexInlineResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFlexInlineResourceMetadata verifies that the generated resource reports the expected type name.
func TestFlexInlineResourceMetadata(t *testing.T) {
	r := &FlexInlineResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_flex_inline" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_flex_inline")
	}
}
