package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllSffpProfileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllSffpProfileDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllSffpProfileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllSffpProfileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllSffpProfileDataSourceMetadata(t *testing.T) {
	d := NewGetAllSffpProfileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_sffp_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_sffp_profile")
	}
}
