package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGigastreamDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllGigastreamDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllGigastreamDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllGigastreamDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllGigastreamDataSourceMetadata(t *testing.T) {
	d := NewLoadAllGigastreamDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_gigastream" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_gigastream")
	}
}
