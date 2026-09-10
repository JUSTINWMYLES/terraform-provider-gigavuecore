package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetImageServerRepoByVersionDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetImageServerRepoByVersionDataSourceSchemaValidation(t *testing.T) {
	d := NewGetImageServerRepoByVersionDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetImageServerRepoByVersionDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetImageServerRepoByVersionDataSourceMetadata(t *testing.T) {
	d := NewGetImageServerRepoByVersionDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_image_server_repo_by_version" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_image_server_repo_by_version")
	}
}
