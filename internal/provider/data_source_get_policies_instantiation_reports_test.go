package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPoliciesInstantiationReportsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPoliciesInstantiationReportsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPoliciesInstantiationReportsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPoliciesInstantiationReportsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPoliciesInstantiationReportsDataSourceMetadata(t *testing.T) {
	d := NewGetPoliciesInstantiationReportsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_policies_instantiation_reports" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_policies_instantiation_reports")
	}
}
