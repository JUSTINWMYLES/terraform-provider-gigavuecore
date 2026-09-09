package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityControlNodeConfigsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMobilityControlNodeConfigsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMobilityControlNodeConfigsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMobilityControlNodeConfigsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMobilityControlNodeConfigsDataSourceMetadata(t *testing.T) {
	d := NewLoadMobilityControlNodeConfigsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_mobility_control_node_configs" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_mobility_control_node_configs")
	}
}
