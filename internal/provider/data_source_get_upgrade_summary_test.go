package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUpgradeSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUpgradeSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUpgradeSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUpgradeSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUpgradeSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetUpgradeSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_upgrade_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_upgrade_summary")
	}
}
