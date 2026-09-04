package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHbProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestHbProfileResourceSchemaValidation(t *testing.T) {
	r := &HbProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestHbProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestHbProfileResourceMetadata(t *testing.T) {
	r := &HbProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_hb_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_hb_profile")
	}
}
