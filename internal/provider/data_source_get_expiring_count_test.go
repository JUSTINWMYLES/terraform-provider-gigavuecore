package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetExpiringCountDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetExpiringCountDataSourceSchemaValidation(t *testing.T) {
	d := NewGetExpiringCountDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetExpiringCountDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetExpiringCountDataSourceMetadata(t *testing.T) {
	d := NewGetExpiringCountDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_expiring_count" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_expiring_count")
	}
}
