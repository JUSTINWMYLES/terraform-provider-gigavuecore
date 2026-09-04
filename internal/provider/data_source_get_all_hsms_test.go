package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllHsmsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllHsmsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllHsmsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllHsmsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllHsmsDataSourceMetadata(t *testing.T) {
	d := NewGetAllHsmsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_hsms" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_hsms")
	}
}
