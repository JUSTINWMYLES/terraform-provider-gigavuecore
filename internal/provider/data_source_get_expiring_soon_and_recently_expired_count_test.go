package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetExpiringSoonAndRecentlyExpiredCountDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSourceSchemaValidation(t *testing.T) {
	d := NewGetExpiringSoonAndRecentlyExpiredCountDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetExpiringSoonAndRecentlyExpiredCountDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetExpiringSoonAndRecentlyExpiredCountDataSourceMetadata(t *testing.T) {
	d := NewGetExpiringSoonAndRecentlyExpiredCountDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_expiring_soon_and_recently_expired_count" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_expiring_soon_and_recently_expired_count")
	}
}
