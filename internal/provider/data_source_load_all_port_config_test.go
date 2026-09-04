package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllPortConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllPortConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllPortConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllPortConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadAllPortConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_port_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_port_config")
	}
}
