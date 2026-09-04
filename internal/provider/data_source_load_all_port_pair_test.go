package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortPairDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllPortPairDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllPortPairDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllPortPairDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllPortPairDataSourceMetadata(t *testing.T) {
	d := NewLoadAllPortPairDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_port_pair" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_port_pair")
	}
}
