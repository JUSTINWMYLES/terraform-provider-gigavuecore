package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestV3UserResourceSchemaValidation verifies that the generated resource schema is valid.
func TestV3UserResourceSchemaValidation(t *testing.T) {
	r := &V3UserResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestV3UserResourceMetadata verifies that the generated resource reports the expected type name.
func TestV3UserResourceMetadata(t *testing.T) {
	r := &V3UserResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_v3_user" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_v3_user")
	}
}
