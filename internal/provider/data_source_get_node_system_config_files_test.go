package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSystemConfigFilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNodeSystemConfigFilesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNodeSystemConfigFilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNodeSystemConfigFilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNodeSystemConfigFilesDataSourceMetadata(t *testing.T) {
	d := NewGetNodeSystemConfigFilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_node_system_config_files" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_node_system_config_files")
	}
}
