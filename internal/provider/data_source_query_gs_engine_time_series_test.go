package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryGsEngineTimeSeriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryGsEngineTimeSeriesDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryGsEngineTimeSeriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryGsEngineTimeSeriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryGsEngineTimeSeriesDataSourceMetadata(t *testing.T) {
	d := NewQueryGsEngineTimeSeriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_gs_engine_time_series" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_gs_engine_time_series")
	}
}
