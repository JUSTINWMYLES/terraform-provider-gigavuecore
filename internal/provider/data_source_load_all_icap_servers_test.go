package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIcapServersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllIcapServersDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllIcapServersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllIcapServersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllIcapServersDataSourceMetadata(t *testing.T) {
	d := NewLoadAllIcapServersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_icap_servers" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_icap_servers")
	}
}
