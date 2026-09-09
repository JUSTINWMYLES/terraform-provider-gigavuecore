package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlowSamplingSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFlowSamplingSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFlowSamplingSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFlowSamplingSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFlowSamplingSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetAllFlowSamplingSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_flow_sampling_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_flow_sampling_summary")
	}
}
