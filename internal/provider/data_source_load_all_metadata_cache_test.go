package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMetadataCacheDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllMetadataCacheDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllMetadataCacheDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllMetadataCacheDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllMetadataCacheDataSourceMetadata(t *testing.T) {
	d := NewLoadAllMetadataCacheDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_metadata_cache" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_metadata_cache")
	}
}
