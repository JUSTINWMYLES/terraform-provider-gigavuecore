package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPpsSourceDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadPpsSourceDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadPpsSourceDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadPpsSourceDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadPpsSourceDataSourceMetadata(t *testing.T) {
	d := NewLoadPpsSourceDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_pps_source" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_pps_source")
	}
}
