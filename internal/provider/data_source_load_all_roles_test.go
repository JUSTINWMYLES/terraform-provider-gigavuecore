package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllRolesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllRolesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllRolesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllRolesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllRolesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllRolesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_roles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_roles")
	}
}
