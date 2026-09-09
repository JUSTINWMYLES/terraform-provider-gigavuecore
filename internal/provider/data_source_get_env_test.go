package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEnvDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEnvDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEnvDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEnvDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEnvDataSourceMetadata(t *testing.T) {
	d := NewGetEnvDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_env" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_env")
	}
}
