package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetProcessedVolumesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetProcessedVolumesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetProcessedVolumesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetProcessedVolumesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetProcessedVolumesDataSourceMetadata(t *testing.T) {
	d := NewGetProcessedVolumesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_processed_volumes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_processed_volumes")
	}
}
