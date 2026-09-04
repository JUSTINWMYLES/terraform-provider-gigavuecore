package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGigaPortsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadGigaPortsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadGigaPortsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadGigaPortsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadGigaPortsDataSourceMetadata(t *testing.T) {
	d := NewLoadGigaPortsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_giga_ports" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_giga_ports")
	}
}
