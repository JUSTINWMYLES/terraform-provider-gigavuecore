package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilitySamNodeConfigsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMobilitySamNodeConfigsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMobilitySamNodeConfigsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMobilitySamNodeConfigsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMobilitySamNodeConfigsDataSourceMetadata(t *testing.T) {
	d := NewLoadMobilitySamNodeConfigsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_mobility_sam_node_configs" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_mobility_sam_node_configs")
	}
}
