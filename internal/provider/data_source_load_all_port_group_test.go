package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPortGroupDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllPortGroupDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllPortGroupDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllPortGroupDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllPortGroupDataSourceMetadata(t *testing.T) {
	d := NewLoadAllPortGroupDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_port_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_port_group")
	}
}
