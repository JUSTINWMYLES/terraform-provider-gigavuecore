package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllMapMigrationResultsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllMapMigrationResultsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllMapMigrationResultsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllMapMigrationResultsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllMapMigrationResultsDataSourceMetadata(t *testing.T) {
	d := NewGetAllMapMigrationResultsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_map_migration_results" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_map_migration_results")
	}
}
