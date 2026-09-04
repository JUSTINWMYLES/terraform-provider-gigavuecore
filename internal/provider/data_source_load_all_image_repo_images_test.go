package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllImageRepoImagesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllImageRepoImagesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllImageRepoImagesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllImageRepoImagesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllImageRepoImagesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllImageRepoImagesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_image_repo_images" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_image_repo_images")
	}
}
