package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyVizEndPointsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadTopologyVizEndPointsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadTopologyVizEndPointsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadTopologyVizEndPointsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadTopologyVizEndPointsDataSourceMetadata(t *testing.T) {
	d := NewLoadTopologyVizEndPointsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_topology_viz_end_points" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_topology_viz_end_points")
	}
}
