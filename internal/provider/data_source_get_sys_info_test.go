package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSysInfoDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSysInfoDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSysInfoDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSysInfoDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSysInfoDataSourceMetadata(t *testing.T) {
	d := NewGetSysInfoDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_sys_info" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_sys_info")
	}
}
