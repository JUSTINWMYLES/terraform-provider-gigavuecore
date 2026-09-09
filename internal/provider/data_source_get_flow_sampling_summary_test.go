package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowSamplingSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFlowSamplingSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFlowSamplingSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFlowSamplingSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFlowSamplingSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetFlowSamplingSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_flow_sampling_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_flow_sampling_summary")
	}
}
