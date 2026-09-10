package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineNetworkLagsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllInlineNetworkLagsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllInlineNetworkLagsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllInlineNetworkLagsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllInlineNetworkLagsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllInlineNetworkLagsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_inline_network_lags" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_inline_network_lags")
	}
}
