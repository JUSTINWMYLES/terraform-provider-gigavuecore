package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllExternalExportTargetServerDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllExternalExportTargetServerDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllExternalExportTargetServerDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllExternalExportTargetServerDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllExternalExportTargetServerDataSourceMetadata(t *testing.T) {
	d := NewGetAllExternalExportTargetServerDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_external_export_target_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_external_export_target_server")
	}
}
