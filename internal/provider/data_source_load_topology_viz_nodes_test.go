package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyVizNodesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadTopologyVizNodesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadTopologyVizNodesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadTopologyVizNodesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadTopologyVizNodesDataSourceMetadata(t *testing.T) {
	d := NewLoadTopologyVizNodesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_topology_viz_nodes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_topology_viz_nodes")
	}
}
