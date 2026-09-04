package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTunnelApplicationDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTunnelApplicationDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTunnelApplicationDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTunnelApplicationDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTunnelApplicationDataSourceMetadata(t *testing.T) {
	d := NewGetAllTunnelApplicationDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_tunnel_application" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_tunnel_application")
	}
}
