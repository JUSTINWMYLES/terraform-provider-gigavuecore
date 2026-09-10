package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIcapProfilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllIcapProfilesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllIcapProfilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllIcapProfilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllIcapProfilesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllIcapProfilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_icap_profiles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_icap_profiles")
	}
}
