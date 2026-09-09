package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGsopInfoDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetGsopInfoDataSourceSchemaValidation(t *testing.T) {
	d := NewGetGsopInfoDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetGsopInfoDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetGsopInfoDataSourceMetadata(t *testing.T) {
	d := NewGetGsopInfoDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_gsop_info" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_gsop_info")
	}
}
