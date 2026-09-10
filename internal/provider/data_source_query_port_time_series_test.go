package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryPortTimeSeriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryPortTimeSeriesDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryPortTimeSeriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryPortTimeSeriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryPortTimeSeriesDataSourceMetadata(t *testing.T) {
	d := NewQueryPortTimeSeriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_port_time_series" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_port_time_series")
	}
}
