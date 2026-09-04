package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetL2CircuitResourceConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetL2CircuitResourceConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewGetL2CircuitResourceConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetL2CircuitResourceConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetL2CircuitResourceConfigDataSourceMetadata(t *testing.T) {
	d := NewGetL2CircuitResourceConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_l2_circuit_resource_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_l2_circuit_resource_config")
	}
}
