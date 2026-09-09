package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAcmeServerDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAcmeServerDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAcmeServerDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAcmeServerDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAcmeServerDetailsDataSourceMetadata(t *testing.T) {
	d := NewGetAcmeServerDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_acme_server_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_acme_server_details")
	}
}
