package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllV3UsersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllV3UsersDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllV3UsersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllV3UsersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllV3UsersDataSourceMetadata(t *testing.T) {
	d := NewLoadAllV3UsersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_v3_users" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_v3_users")
	}
}
