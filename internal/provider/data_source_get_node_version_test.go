package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeVersionDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNodeVersionDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNodeVersionDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNodeVersionDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNodeVersionDataSourceMetadata(t *testing.T) {
	d := NewGetNodeVersionDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_node_version" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_node_version")
	}
}
