package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeExecutionTimelineDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUpgradeExecutionTimelineDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUpgradeExecutionTimelineDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUpgradeExecutionTimelineDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUpgradeExecutionTimelineDataSourceMetadata(t *testing.T) {
	d := NewGetUpgradeExecutionTimelineDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_upgrade_execution_timeline" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_upgrade_execution_timeline")
	}
}
