package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlowSipSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFlowSipSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFlowSipSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFlowSipSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFlowSipSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetFlowSipSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_flow_sip_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_flow_sip_summary")
	}
}
