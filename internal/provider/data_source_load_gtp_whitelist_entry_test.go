package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadGtpWhitelistEntryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadGtpWhitelistEntryDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadGtpWhitelistEntryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadGtpWhitelistEntryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadGtpWhitelistEntryDataSourceMetadata(t *testing.T) {
	d := NewLoadGtpWhitelistEntryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_gtp_whitelist_entry" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_gtp_whitelist_entry")
	}
}
