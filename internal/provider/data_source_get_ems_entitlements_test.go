package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmsEntitlementsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEmsEntitlementsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEmsEntitlementsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEmsEntitlementsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEmsEntitlementsDataSourceMetadata(t *testing.T) {
	d := NewGetEmsEntitlementsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ems_entitlements" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ems_entitlements")
	}
}
