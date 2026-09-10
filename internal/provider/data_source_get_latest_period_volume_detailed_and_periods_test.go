package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetLatestPeriodVolumeDetailedAndPeriodsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetLatestPeriodVolumeDetailedAndPeriodsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetLatestPeriodVolumeDetailedAndPeriodsDataSourceMetadata(t *testing.T) {
	d := NewGetLatestPeriodVolumeDetailedAndPeriodsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_latest_period_volume_detailed_and_periods" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_latest_period_volume_detailed_and_periods")
	}
}
