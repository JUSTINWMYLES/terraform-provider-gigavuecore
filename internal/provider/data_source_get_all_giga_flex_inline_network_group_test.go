package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllGigaFlexInlineNetworkGroupDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllGigaFlexInlineNetworkGroupDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllGigaFlexInlineNetworkGroupDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllGigaFlexInlineNetworkGroupDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllGigaFlexInlineNetworkGroupDataSourceMetadata(t *testing.T) {
	d := NewGetAllGigaFlexInlineNetworkGroupDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_giga_flex_inline_network_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_giga_flex_inline_network_group")
	}
}
