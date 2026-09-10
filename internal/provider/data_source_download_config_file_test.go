package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestDownloadConfigFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestDownloadConfigFileDataSourceSchemaValidation(t *testing.T) {
	d := NewDownloadConfigFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestDownloadConfigFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestDownloadConfigFileDataSourceMetadata(t *testing.T) {
	d := NewDownloadConfigFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_download_config_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_download_config_file")
	}
}
