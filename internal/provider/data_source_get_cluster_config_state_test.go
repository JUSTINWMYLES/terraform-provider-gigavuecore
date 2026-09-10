package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterConfigStateDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetClusterConfigStateDataSourceSchemaValidation(t *testing.T) {
	d := NewGetClusterConfigStateDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetClusterConfigStateDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetClusterConfigStateDataSourceMetadata(t *testing.T) {
	d := NewGetClusterConfigStateDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_cluster_config_state" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_cluster_config_state")
	}
}
