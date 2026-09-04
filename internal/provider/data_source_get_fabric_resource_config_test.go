package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFabricResourceConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFabricResourceConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFabricResourceConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFabricResourceConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFabricResourceConfigDataSourceMetadata(t *testing.T) {
	d := NewGetFabricResourceConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_fabric_resource_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_fabric_resource_config")
	}
}
