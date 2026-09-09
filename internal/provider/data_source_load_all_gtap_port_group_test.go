package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGtapPortGroupDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllGtapPortGroupDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllGtapPortGroupDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllGtapPortGroupDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllGtapPortGroupDataSourceMetadata(t *testing.T) {
	d := NewLoadAllGtapPortGroupDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_gtap_port_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_gtap_port_group")
	}
}
