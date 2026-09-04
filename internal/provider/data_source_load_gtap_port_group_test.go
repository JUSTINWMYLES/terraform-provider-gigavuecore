package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGtapPortGroupDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadGtapPortGroupDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadGtapPortGroupDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadGtapPortGroupDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadGtapPortGroupDataSourceMetadata(t *testing.T) {
	d := NewLoadGtapPortGroupDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_gtap_port_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_gtap_port_group")
	}
}
