package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCustomerNameDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCustomerNameDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCustomerNameDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCustomerNameDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCustomerNameDataSourceMetadata(t *testing.T) {
	d := NewGetCustomerNameDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_customer_name" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_customer_name")
	}
}
