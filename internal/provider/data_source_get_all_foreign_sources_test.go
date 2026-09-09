package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllForeignSourcesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllForeignSourcesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllForeignSourcesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllForeignSourcesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllForeignSourcesDataSourceMetadata(t *testing.T) {
	d := NewGetAllForeignSourcesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_foreign_sources" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_foreign_sources")
	}
}
