package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPolicyResourceSchemaValidation verifies that the generated resource schema is valid.
func TestPolicyResourceSchemaValidation(t *testing.T) {
	r := &PolicyResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestPolicyResourceMetadata verifies that the generated resource reports the expected type name.
func TestPolicyResourceMetadata(t *testing.T) {
	r := &PolicyResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_policy" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_policy")
	}
}
