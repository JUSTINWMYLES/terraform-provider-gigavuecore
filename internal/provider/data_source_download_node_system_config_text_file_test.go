package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestDownloadNodeSystemConfigTextFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestDownloadNodeSystemConfigTextFileDataSourceSchemaValidation(t *testing.T) {
	d := NewDownloadNodeSystemConfigTextFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestDownloadNodeSystemConfigTextFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestDownloadNodeSystemConfigTextFileDataSourceMetadata(t *testing.T) {
	d := NewDownloadNodeSystemConfigTextFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_download_node_system_config_text_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_download_node_system_config_text_file")
	}
}
