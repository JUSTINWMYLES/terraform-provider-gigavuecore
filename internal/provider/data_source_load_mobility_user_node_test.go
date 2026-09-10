package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityUserNodeDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMobilityUserNodeDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMobilityUserNodeDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMobilityUserNodeDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMobilityUserNodeDataSourceMetadata(t *testing.T) {
	d := NewLoadMobilityUserNodeDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_mobility_user_node" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_mobility_user_node")
	}
}
