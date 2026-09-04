package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineSslProfilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllInlineSslProfilesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllInlineSslProfilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllInlineSslProfilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllInlineSslProfilesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllInlineSslProfilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_inline_ssl_profiles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_inline_ssl_profiles")
	}
}
