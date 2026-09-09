package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryForApplicationsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryForApplicationsDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryForApplicationsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryForApplicationsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryForApplicationsDataSourceMetadata(t *testing.T) {
	d := NewQueryForApplicationsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_for_applications" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_for_applications")
	}
}
