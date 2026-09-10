package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityControlNodeDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMobilityControlNodeDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMobilityControlNodeDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMobilityControlNodeDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMobilityControlNodeDataSourceMetadata(t *testing.T) {
	d := NewLoadMobilityControlNodeDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_mobility_control_node" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_mobility_control_node")
	}
}
