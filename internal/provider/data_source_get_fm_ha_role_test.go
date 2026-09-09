package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFmHaRoleDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFmHaRoleDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFmHaRoleDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFmHaRoleDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFmHaRoleDataSourceMetadata(t *testing.T) {
	d := NewGetFmHaRoleDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_fm_ha_role" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_fm_ha_role")
	}
}
