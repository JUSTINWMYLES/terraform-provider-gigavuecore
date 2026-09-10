package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllForeignMastersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllForeignMastersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllForeignMastersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllForeignMastersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllForeignMastersDataSourceMetadata(t *testing.T) {
	d := NewGetAllForeignMastersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_foreign_masters" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_foreign_masters")
	}
}
