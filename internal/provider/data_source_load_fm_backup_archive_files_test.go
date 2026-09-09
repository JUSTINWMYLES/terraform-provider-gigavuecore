package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmBackupArchiveFilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadFmBackupArchiveFilesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadFmBackupArchiveFilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadFmBackupArchiveFilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadFmBackupArchiveFilesDataSourceMetadata(t *testing.T) {
	d := NewLoadFmBackupArchiveFilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_fm_backup_archive_files" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_fm_backup_archive_files")
	}
}
