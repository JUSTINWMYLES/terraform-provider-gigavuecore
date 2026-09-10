package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestResourceSelectionDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestResourceSelectionDataSourceSchemaValidation(t *testing.T) {
	d := NewResourceSelectionDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestResourceSelectionDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestResourceSelectionDataSourceMetadata(t *testing.T) {
	d := NewResourceSelectionDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_resource_selection" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_resource_selection")
	}
}
