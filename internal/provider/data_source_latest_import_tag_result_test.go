package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLatestImportTagResultDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLatestImportTagResultDataSourceSchemaValidation(t *testing.T) {
	d := NewLatestImportTagResultDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLatestImportTagResultDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLatestImportTagResultDataSourceMetadata(t *testing.T) {
	d := NewLatestImportTagResultDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_latest_import_tag_result" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_latest_import_tag_result")
	}
}
