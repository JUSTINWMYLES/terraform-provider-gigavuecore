package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAvisiPoliciesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAvisiPoliciesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAvisiPoliciesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAvisiPoliciesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAvisiPoliciesDataSourceMetadata(t *testing.T) {
	d := NewGetAvisiPoliciesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_avisi_policies" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_avisi_policies")
	}
}
