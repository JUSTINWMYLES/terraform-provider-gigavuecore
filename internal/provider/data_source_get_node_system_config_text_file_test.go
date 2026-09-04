package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeSystemConfigTextFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNodeSystemConfigTextFileDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNodeSystemConfigTextFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNodeSystemConfigTextFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNodeSystemConfigTextFileDataSourceMetadata(t *testing.T) {
	d := NewGetNodeSystemConfigTextFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_node_system_config_text_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_node_system_config_text_file")
	}
}
