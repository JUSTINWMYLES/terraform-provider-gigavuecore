package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetConnectedNodesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetConnectedNodesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetConnectedNodesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetConnectedNodesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetConnectedNodesDataSourceMetadata(t *testing.T) {
	d := NewGetConnectedNodesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_connected_nodes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_connected_nodes")
	}
}
