package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLastNPeriodsSummaryVolumesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetLastNPeriodsSummaryVolumesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetLastNPeriodsSummaryVolumesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetLastNPeriodsSummaryVolumesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetLastNPeriodsSummaryVolumesDataSourceMetadata(t *testing.T) {
	d := NewGetLastNPeriodsSummaryVolumesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_last_n_periods_summary_volumes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_last_n_periods_summary_volumes")
	}
}
