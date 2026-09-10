package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestHsmResourceSchemaValidation verifies that the generated resource schema is valid.
func TestHsmResourceSchemaValidation(t *testing.T) {
	r := &HsmResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestHsmResourceMetadata verifies that the generated resource reports the expected type name.
func TestHsmResourceMetadata(t *testing.T) {
	r := &HsmResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_hsm" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_hsm")
	}
}
