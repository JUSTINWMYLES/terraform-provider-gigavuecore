package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryTunneledPortTimeSeriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryTunneledPortTimeSeriesDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryTunneledPortTimeSeriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryTunneledPortTimeSeriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryTunneledPortTimeSeriesDataSourceMetadata(t *testing.T) {
	d := NewQueryTunneledPortTimeSeriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_tunneled_port_time_series" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_tunneled_port_time_series")
	}
}
