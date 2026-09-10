package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetForeignMastersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetForeignMastersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetForeignMastersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetForeignMastersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetForeignMastersDataSourceMetadata(t *testing.T) {
	d := NewGetForeignMastersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_foreign_masters" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_foreign_masters")
	}
}
