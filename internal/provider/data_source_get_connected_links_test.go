package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetConnectedLinksDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetConnectedLinksDataSourceSchemaValidation(t *testing.T) {
	d := NewGetConnectedLinksDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetConnectedLinksDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetConnectedLinksDataSourceMetadata(t *testing.T) {
	d := NewGetConnectedLinksDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_connected_links" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_connected_links")
	}
}
