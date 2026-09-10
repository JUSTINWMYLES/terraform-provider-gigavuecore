package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListDeviceConfigFilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListDeviceConfigFilesDataSourceSchemaValidation(t *testing.T) {
	d := NewListDeviceConfigFilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListDeviceConfigFilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListDeviceConfigFilesDataSourceMetadata(t *testing.T) {
	d := NewListDeviceConfigFilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_device_config_files" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_device_config_files")
	}
}
