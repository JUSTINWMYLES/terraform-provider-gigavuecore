package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestLocalUserResourceSchemaValidation verifies that the generated resource schema is valid.
func TestLocalUserResourceSchemaValidation(t *testing.T) {
	r := &LocalUserResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLocalUserResourceMetadata verifies that the generated resource reports the expected type name.
func TestLocalUserResourceMetadata(t *testing.T) {
	r := &LocalUserResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_local_user" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_local_user")
	}
}
