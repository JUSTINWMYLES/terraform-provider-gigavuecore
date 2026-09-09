package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryPortGroupLbTimeSeriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryPortGroupLbTimeSeriesDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryPortGroupLbTimeSeriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryPortGroupLbTimeSeriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryPortGroupLbTimeSeriesDataSourceMetadata(t *testing.T) {
	d := NewQueryPortGroupLbTimeSeriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_port_group_lb_time_series" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_port_group_lb_time_series")
	}
}
