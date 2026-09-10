package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilityUserNodeConfigsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMobilityUserNodeConfigsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMobilityUserNodeConfigsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMobilityUserNodeConfigsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMobilityUserNodeConfigsDataSourceMetadata(t *testing.T) {
	d := NewLoadMobilityUserNodeConfigsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_mobility_user_node_configs" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_mobility_user_node_configs")
	}
}
