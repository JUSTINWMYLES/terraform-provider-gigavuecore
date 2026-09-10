package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryNetflowExporterTimeSeriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryNetflowExporterTimeSeriesDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryNetflowExporterTimeSeriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryNetflowExporterTimeSeriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryNetflowExporterTimeSeriesDataSourceMetadata(t *testing.T) {
	d := NewQueryNetflowExporterTimeSeriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_netflow_exporter_time_series" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_netflow_exporter_time_series")
	}
}
