package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeJobsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUpgradeJobsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUpgradeJobsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUpgradeJobsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUpgradeJobsDataSourceMetadata(t *testing.T) {
	d := NewGetUpgradeJobsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_upgrade_jobs" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_upgrade_jobs")
	}
}
