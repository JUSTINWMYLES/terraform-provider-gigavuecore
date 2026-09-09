package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPeriodVolumeDetailedByDateDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPeriodVolumeDetailedByDateDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPeriodVolumeDetailedByDateDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPeriodVolumeDetailedByDateDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPeriodVolumeDetailedByDateDataSourceMetadata(t *testing.T) {
	d := NewGetPeriodVolumeDetailedByDateDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_period_volume_detailed_by_date" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_period_volume_detailed_by_date")
	}
}
