package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTopologyVizConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadTopologyVizConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadTopologyVizConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadTopologyVizConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadTopologyVizConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadTopologyVizConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_topology_viz_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_topology_viz_config")
	}
}
