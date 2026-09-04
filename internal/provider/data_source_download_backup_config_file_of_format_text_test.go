package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestDownloadBackupConfigFileOfFormatTextDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestDownloadBackupConfigFileOfFormatTextDataSourceSchemaValidation(t *testing.T) {
	d := NewDownloadBackupConfigFileOfFormatTextDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestDownloadBackupConfigFileOfFormatTextDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestDownloadBackupConfigFileOfFormatTextDataSourceMetadata(t *testing.T) {
	d := NewDownloadBackupConfigFileOfFormatTextDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_download_backup_config_file_of_format_text" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_download_backup_config_file_of_format_text")
	}
}
