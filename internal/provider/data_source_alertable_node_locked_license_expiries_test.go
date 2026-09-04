package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestAlertableNodeLockedLicenseExpiriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestAlertableNodeLockedLicenseExpiriesDataSourceSchemaValidation(t *testing.T) {
	d := NewAlertableNodeLockedLicenseExpiriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestAlertableNodeLockedLicenseExpiriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestAlertableNodeLockedLicenseExpiriesDataSourceMetadata(t *testing.T) {
	d := NewAlertableNodeLockedLicenseExpiriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_alertable_node_locked_license_expiries" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_alertable_node_locked_license_expiries")
	}
}
