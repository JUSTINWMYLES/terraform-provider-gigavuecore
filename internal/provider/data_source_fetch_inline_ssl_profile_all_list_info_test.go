package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestFetchInlineSslProfileAllListInfoDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestFetchInlineSslProfileAllListInfoDataSourceSchemaValidation(t *testing.T) {
	d := NewFetchInlineSslProfileAllListInfoDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFetchInlineSslProfileAllListInfoDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestFetchInlineSslProfileAllListInfoDataSourceMetadata(t *testing.T) {
	d := NewFetchInlineSslProfileAllListInfoDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_fetch_inline_ssl_profile_all_list_info" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_fetch_inline_ssl_profile_all_list_info")
	}
}
