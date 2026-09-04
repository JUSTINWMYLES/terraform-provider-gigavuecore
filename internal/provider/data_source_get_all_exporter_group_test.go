package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllExporterGroupDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllExporterGroupDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllExporterGroupDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllExporterGroupDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllExporterGroupDataSourceMetadata(t *testing.T) {
	d := NewGetAllExporterGroupDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_exporter_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_exporter_group")
	}
}
