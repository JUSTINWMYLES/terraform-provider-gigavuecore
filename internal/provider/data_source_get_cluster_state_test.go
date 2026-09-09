package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterStateDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetClusterStateDataSourceSchemaValidation(t *testing.T) {
	d := NewGetClusterStateDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetClusterStateDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetClusterStateDataSourceMetadata(t *testing.T) {
	d := NewGetClusterStateDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_cluster_state" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_cluster_state")
	}
}
