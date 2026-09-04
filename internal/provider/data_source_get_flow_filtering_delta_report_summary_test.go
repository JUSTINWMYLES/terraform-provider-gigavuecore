package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowFilteringDeltaReportSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFlowFilteringDeltaReportSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFlowFilteringDeltaReportSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFlowFilteringDeltaReportSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFlowFilteringDeltaReportSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetFlowFilteringDeltaReportSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_flow_filtering_delta_report_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_flow_filtering_delta_report_summary")
	}
}
