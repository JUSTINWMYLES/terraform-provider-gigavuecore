package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestValidateExternalExportServerDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestValidateExternalExportServerDataSourceSchemaValidation(t *testing.T) {
	d := NewValidateExternalExportServerDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestValidateExternalExportServerDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestValidateExternalExportServerDataSourceMetadata(t *testing.T) {
	d := NewValidateExternalExportServerDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_validate_external_export_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_validate_external_export_server")
	}
}
