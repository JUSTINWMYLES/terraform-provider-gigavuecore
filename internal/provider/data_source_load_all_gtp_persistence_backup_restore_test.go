package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGtpPersistenceBackupRestoreDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllGtpPersistenceBackupRestoreDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllGtpPersistenceBackupRestoreDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllGtpPersistenceBackupRestoreDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllGtpPersistenceBackupRestoreDataSourceMetadata(t *testing.T) {
	d := NewLoadAllGtpPersistenceBackupRestoreDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_gtp_persistence_backup_restore" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_gtp_persistence_backup_restore")
	}
}
