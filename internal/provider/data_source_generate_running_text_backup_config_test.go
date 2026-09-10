package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGenerateRunningTextBackupConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGenerateRunningTextBackupConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewGenerateRunningTextBackupConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGenerateRunningTextBackupConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGenerateRunningTextBackupConfigDataSourceMetadata(t *testing.T) {
	d := NewGenerateRunningTextBackupConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_generate_running_text_backup_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_generate_running_text_backup_config")
	}
}
