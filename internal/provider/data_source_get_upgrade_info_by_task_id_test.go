package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeInfoByTaskIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUpgradeInfoByTaskIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUpgradeInfoByTaskIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUpgradeInfoByTaskIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUpgradeInfoByTaskIdDataSourceMetadata(t *testing.T) {
	d := NewGetUpgradeInfoByTaskIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_upgrade_info_by_task_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_upgrade_info_by_task_id")
	}
}
