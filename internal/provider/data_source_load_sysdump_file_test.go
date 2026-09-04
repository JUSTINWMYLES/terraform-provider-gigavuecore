package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSysdumpFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSysdumpFileDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSysdumpFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSysdumpFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSysdumpFileDataSourceMetadata(t *testing.T) {
	d := NewLoadSysdumpFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_sysdump_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_sysdump_file")
	}
}
