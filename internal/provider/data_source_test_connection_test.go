package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestTestConnectionDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestTestConnectionDataSourceSchemaValidation(t *testing.T) {
	d := NewTestConnectionDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestTestConnectionDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestTestConnectionDataSourceMetadata(t *testing.T) {
	d := NewTestConnectionDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_test_connection" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_test_connection")
	}
}
