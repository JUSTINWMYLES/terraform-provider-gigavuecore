package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListConnectionsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListConnectionsDataSourceSchemaValidation(t *testing.T) {
	d := NewListConnectionsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListConnectionsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListConnectionsDataSourceMetadata(t *testing.T) {
	d := NewListConnectionsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_connections" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_connections")
	}
}
