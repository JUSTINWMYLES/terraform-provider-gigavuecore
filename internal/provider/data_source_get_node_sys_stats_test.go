package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSysStatsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNodeSysStatsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNodeSysStatsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNodeSysStatsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNodeSysStatsDataSourceMetadata(t *testing.T) {
	d := NewGetNodeSysStatsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_node_sys_stats" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_node_sys_stats")
	}
}
