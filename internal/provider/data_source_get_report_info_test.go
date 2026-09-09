package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetReportInfoDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetReportInfoDataSourceSchemaValidation(t *testing.T) {
	d := NewGetReportInfoDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetReportInfoDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetReportInfoDataSourceMetadata(t *testing.T) {
	d := NewGetReportInfoDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_report_info" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_report_info")
	}
}
