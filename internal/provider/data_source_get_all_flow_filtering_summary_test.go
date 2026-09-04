package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlowFilteringSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFlowFilteringSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFlowFilteringSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFlowFilteringSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFlowFilteringSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetAllFlowFilteringSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_flow_filtering_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_flow_filtering_summary")
	}
}
