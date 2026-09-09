package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllCriticalNotificationsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllCriticalNotificationsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllCriticalNotificationsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllCriticalNotificationsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllCriticalNotificationsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllCriticalNotificationsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_critical_notifications" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_critical_notifications")
	}
}
