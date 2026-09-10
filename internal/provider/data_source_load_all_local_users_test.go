package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllLocalUsersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllLocalUsersDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllLocalUsersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllLocalUsersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllLocalUsersDataSourceMetadata(t *testing.T) {
	d := NewLoadAllLocalUsersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_local_users" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_local_users")
	}
}
