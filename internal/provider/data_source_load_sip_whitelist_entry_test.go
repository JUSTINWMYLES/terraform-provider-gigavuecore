package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSipWhitelistEntryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSipWhitelistEntryDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSipWhitelistEntryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSipWhitelistEntryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSipWhitelistEntryDataSourceMetadata(t *testing.T) {
	d := NewLoadSipWhitelistEntryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_sip_whitelist_entry" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_sip_whitelist_entry")
	}
}
