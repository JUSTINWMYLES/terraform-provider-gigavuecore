package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllMapMigrationResultsByAliasDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllMapMigrationResultsByAliasDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllMapMigrationResultsByAliasDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllMapMigrationResultsByAliasDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllMapMigrationResultsByAliasDataSourceMetadata(t *testing.T) {
	d := NewGetAllMapMigrationResultsByAliasDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_map_migration_results_by_alias" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_map_migration_results_by_alias")
	}
}
