package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUserTokensForPrivilegeUsersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUserTokensForPrivilegeUsersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUserTokensForPrivilegeUsersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUserTokensForPrivilegeUsersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUserTokensForPrivilegeUsersDataSourceMetadata(t *testing.T) {
	d := NewGetUserTokensForPrivilegeUsersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_user_tokens_for_privilege_users" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_user_tokens_for_privilege_users")
	}
}
