package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGsDumpFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadGsDumpFileDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadGsDumpFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadGsDumpFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadGsDumpFileDataSourceMetadata(t *testing.T) {
	d := NewLoadGsDumpFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_gs_dump_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_gs_dump_file")
	}
}
