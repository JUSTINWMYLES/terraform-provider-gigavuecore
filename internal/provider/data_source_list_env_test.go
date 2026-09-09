package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListEnvDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListEnvDataSourceSchemaValidation(t *testing.T) {
	d := NewListEnvDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListEnvDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListEnvDataSourceMetadata(t *testing.T) {
	d := NewListEnvDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_env" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_env")
	}
}
