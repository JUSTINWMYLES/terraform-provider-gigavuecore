package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadIpv6NeighborsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadIpv6NeighborsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadIpv6NeighborsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadIpv6NeighborsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadIpv6NeighborsDataSourceMetadata(t *testing.T) {
	d := NewLoadIpv6NeighborsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_ipv6_neighbors" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_ipv6_neighbors")
	}
}
