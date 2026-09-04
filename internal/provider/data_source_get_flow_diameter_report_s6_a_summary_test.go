package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowDiameterReportS6ASummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFlowDiameterReportS6ASummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFlowDiameterReportS6ASummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFlowDiameterReportS6ASummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFlowDiameterReportS6ASummaryDataSourceMetadata(t *testing.T) {
	d := NewGetFlowDiameterReportS6ASummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_flow_diameter_report_s6_a_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_flow_diameter_report_s6_a_summary")
	}
}
