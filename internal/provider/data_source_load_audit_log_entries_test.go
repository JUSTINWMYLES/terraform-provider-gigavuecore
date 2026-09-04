package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAuditLogEntriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAuditLogEntriesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAuditLogEntriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAuditLogEntriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAuditLogEntriesDataSourceMetadata(t *testing.T) {
	d := NewLoadAuditLogEntriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_audit_log_entries" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_audit_log_entries")
	}
}
