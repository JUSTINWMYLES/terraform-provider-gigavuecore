package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllicapClientDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllicapClientDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllicapClientDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllicapClientDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllicapClientDataSourceMetadata(t *testing.T) {
	d := NewGetAllicapClientDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_allicap_client" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_allicap_client")
	}
}
