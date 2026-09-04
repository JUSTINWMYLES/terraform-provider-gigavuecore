package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllUsersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllUsersDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllUsersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllUsersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllUsersDataSourceMetadata(t *testing.T) {
	d := NewLoadAllUsersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_users" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_users")
	}
}
