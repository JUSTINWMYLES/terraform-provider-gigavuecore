package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAaaAuthConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAaaAuthConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAaaAuthConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAaaAuthConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAaaAuthConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadAaaAuthConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_aaa_auth_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_aaa_auth_config")
	}
}
