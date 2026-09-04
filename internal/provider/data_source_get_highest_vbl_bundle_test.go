package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetHighestVblBundleDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetHighestVblBundleDataSourceSchemaValidation(t *testing.T) {
	d := NewGetHighestVblBundleDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetHighestVblBundleDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetHighestVblBundleDataSourceMetadata(t *testing.T) {
	d := NewGetHighestVblBundleDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_highest_vbl_bundle" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_highest_vbl_bundle")
	}
}
