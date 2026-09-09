package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGtpBackupFilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetGtpBackupFilesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetGtpBackupFilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetGtpBackupFilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetGtpBackupFilesDataSourceMetadata(t *testing.T) {
	d := NewGetGtpBackupFilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_gtp_backup_files" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_gtp_backup_files")
	}
}
