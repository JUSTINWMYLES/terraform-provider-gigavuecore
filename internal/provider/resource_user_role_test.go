package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestUserRoleResourceSchemaValidation verifies that the generated resource schema is valid.
func TestUserRoleResourceSchemaValidation(t *testing.T) {
	r := &UserRoleResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestUserRoleResourceMetadata verifies that the generated resource reports the expected type name.
func TestUserRoleResourceMetadata(t *testing.T) {
	r := &UserRoleResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_user_role" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_user_role")
	}
}
