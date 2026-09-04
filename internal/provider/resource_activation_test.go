package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestActivationResourceSchemaValidation verifies that the generated resource schema is valid.
func TestActivationResourceSchemaValidation(t *testing.T) {
	r := &ActivationResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestActivationResourceMetadata verifies that the generated resource reports the expected type name.
func TestActivationResourceMetadata(t *testing.T) {
	r := &ActivationResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_activation" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_activation")
	}
}
