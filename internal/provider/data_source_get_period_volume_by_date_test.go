package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPeriodVolumeByDateDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPeriodVolumeByDateDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPeriodVolumeByDateDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPeriodVolumeByDateDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPeriodVolumeByDateDataSourceMetadata(t *testing.T) {
	d := NewGetPeriodVolumeByDateDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_period_volume_by_date" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_period_volume_by_date")
	}
}
