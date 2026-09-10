package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGsDumpDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadGsDumpDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadGsDumpDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadGsDumpDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadGsDumpDataSourceMetadata(t *testing.T) {
	d := NewLoadGsDumpDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_gs_dump" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_gs_dump")
	}
}
