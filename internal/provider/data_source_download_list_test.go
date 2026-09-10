package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestDownloadListDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestDownloadListDataSourceSchemaValidation(t *testing.T) {
	d := NewDownloadListDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestDownloadListDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestDownloadListDataSourceMetadata(t *testing.T) {
	d := NewDownloadListDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_download_list" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_download_list")
	}
}
