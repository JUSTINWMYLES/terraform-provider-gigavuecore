package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetFlexInlineNetworkConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetFlexInlineNetworkConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewGetFlexInlineNetworkConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetFlexInlineNetworkConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetFlexInlineNetworkConfigDataSourceMetadata(t *testing.T) {
	d := NewGetFlexInlineNetworkConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_flex_inline_network_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_flex_inline_network_config")
	}
}
