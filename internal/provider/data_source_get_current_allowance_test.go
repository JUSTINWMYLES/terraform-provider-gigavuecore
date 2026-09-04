package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentAllowanceDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCurrentAllowanceDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCurrentAllowanceDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCurrentAllowanceDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCurrentAllowanceDataSourceMetadata(t *testing.T) {
	d := NewGetCurrentAllowanceDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_current_allowance" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_current_allowance")
	}
}
