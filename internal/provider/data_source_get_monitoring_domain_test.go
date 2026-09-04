package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMonitoringDomainDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetMonitoringDomainDataSourceSchemaValidation(t *testing.T) {
	d := NewGetMonitoringDomainDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetMonitoringDomainDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetMonitoringDomainDataSourceMetadata(t *testing.T) {
	d := NewGetMonitoringDomainDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_monitoring_domain" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_monitoring_domain")
	}
}
