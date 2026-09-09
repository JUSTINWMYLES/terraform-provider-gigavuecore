package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemDiagDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemDiagDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemDiagDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemDiagDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemDiagDataSourceMetadata(t *testing.T) {
	d := NewGetSystemDiagDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_system_diag" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_system_diag")
	}
}
