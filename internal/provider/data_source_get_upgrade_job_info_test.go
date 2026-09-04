package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeJobInfoDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUpgradeJobInfoDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUpgradeJobInfoDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUpgradeJobInfoDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUpgradeJobInfoDataSourceMetadata(t *testing.T) {
	d := NewGetUpgradeJobInfoDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_upgrade_job_info" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_upgrade_job_info")
	}
}
