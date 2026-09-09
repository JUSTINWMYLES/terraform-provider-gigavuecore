package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpCountersByAliasDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPtpCountersByAliasDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPtpCountersByAliasDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPtpCountersByAliasDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPtpCountersByAliasDataSourceMetadata(t *testing.T) {
	d := NewGetPtpCountersByAliasDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ptp_counters_by_alias" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ptp_counters_by_alias")
	}
}
