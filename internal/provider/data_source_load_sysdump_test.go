package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSysdumpDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSysdumpDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSysdumpDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSysdumpDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSysdumpDataSourceMetadata(t *testing.T) {
	d := NewLoadSysdumpDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_sysdump" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_sysdump")
	}
}
