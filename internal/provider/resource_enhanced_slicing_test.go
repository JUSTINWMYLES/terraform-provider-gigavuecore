package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestEnhancedSlicingResourceSchemaValidation verifies that the generated resource schema is valid.
func TestEnhancedSlicingResourceSchemaValidation(t *testing.T) {
	r := &EnhancedSlicingResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestEnhancedSlicingResourceMetadata verifies that the generated resource reports the expected type name.
func TestEnhancedSlicingResourceMetadata(t *testing.T) {
	r := &EnhancedSlicingResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_enhanced_slicing" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_enhanced_slicing")
	}
}
