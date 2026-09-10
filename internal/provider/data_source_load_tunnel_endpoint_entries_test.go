package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadTunnelEndpointEntriesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadTunnelEndpointEntriesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadTunnelEndpointEntriesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadTunnelEndpointEntriesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadTunnelEndpointEntriesDataSourceMetadata(t *testing.T) {
	d := NewLoadTunnelEndpointEntriesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_tunnel_endpoint_entries" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_tunnel_endpoint_entries")
	}
}
