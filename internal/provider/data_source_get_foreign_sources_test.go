package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetForeignSourcesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetForeignSourcesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetForeignSourcesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetForeignSourcesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetForeignSourcesDataSourceMetadata(t *testing.T) {
	d := NewGetForeignSourcesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_foreign_sources" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_foreign_sources")
	}
}
