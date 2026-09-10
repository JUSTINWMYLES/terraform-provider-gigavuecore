package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNodesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNodesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNodesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNodesDataSourceMetadata(t *testing.T) {
	d := NewGetNodesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_nodes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_nodes")
	}
}
