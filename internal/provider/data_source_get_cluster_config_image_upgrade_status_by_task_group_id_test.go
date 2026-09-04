package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSourceMetadata(t *testing.T) {
	d := NewGetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id")
	}
}
