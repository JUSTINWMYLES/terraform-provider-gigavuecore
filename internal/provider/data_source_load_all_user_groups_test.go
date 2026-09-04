package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllUserGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllUserGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllUserGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllUserGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllUserGroupsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllUserGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_user_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_user_groups")
	}
}
