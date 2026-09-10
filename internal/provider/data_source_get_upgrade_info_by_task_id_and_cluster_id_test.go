package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUpgradeInfoByTaskIdAndClusterIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUpgradeInfoByTaskIdAndClusterIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUpgradeInfoByTaskIdAndClusterIdDataSourceMetadata(t *testing.T) {
	d := NewGetUpgradeInfoByTaskIdAndClusterIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_upgrade_info_by_task_id_and_cluster_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_upgrade_info_by_task_id_and_cluster_id")
	}
}
