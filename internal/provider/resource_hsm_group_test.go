package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHsmGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestHsmGroupResourceSchemaValidation(t *testing.T) {
	r := &HsmGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestHsmGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestHsmGroupResourceMetadata(t *testing.T) {
	r := &HsmGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_hsm_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_hsm_group")
	}
}
