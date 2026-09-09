package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetHsmGroupStatusReportDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetHsmGroupStatusReportDataSourceSchemaValidation(t *testing.T) {
	d := NewGetHsmGroupStatusReportDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetHsmGroupStatusReportDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetHsmGroupStatusReportDataSourceMetadata(t *testing.T) {
	d := NewGetHsmGroupStatusReportDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_hsm_group_status_report" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_hsm_group_status_report")
	}
}
