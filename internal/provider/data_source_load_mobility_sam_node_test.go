package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMobilitySamNodeDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMobilitySamNodeDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMobilitySamNodeDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMobilitySamNodeDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMobilitySamNodeDataSourceMetadata(t *testing.T) {
	d := NewLoadMobilitySamNodeDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_mobility_sam_node" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_mobility_sam_node")
	}
}
