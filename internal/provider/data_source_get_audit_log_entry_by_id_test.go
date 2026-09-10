package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAuditLogEntryByIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAuditLogEntryByIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAuditLogEntryByIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAuditLogEntryByIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAuditLogEntryByIdDataSourceMetadata(t *testing.T) {
	d := NewGetAuditLogEntryByIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_audit_log_entry_by_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_audit_log_entry_by_id")
	}
}
