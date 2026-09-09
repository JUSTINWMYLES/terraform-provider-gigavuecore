package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestAppVisibilityResourceSchemaValidation verifies that the generated resource schema is valid.
func TestAppVisibilityResourceSchemaValidation(t *testing.T) {
	r := &AppVisibilityResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestAppVisibilityResourceMetadata verifies that the generated resource reports the expected type name.
func TestAppVisibilityResourceMetadata(t *testing.T) {
	r := &AppVisibilityResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_app_visibility" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_app_visibility")
	}
}
