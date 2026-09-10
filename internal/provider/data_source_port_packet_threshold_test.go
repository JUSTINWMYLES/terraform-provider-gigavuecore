package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestPortPacketThresholdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestPortPacketThresholdDataSourceSchemaValidation(t *testing.T) {
	d := NewPortPacketThresholdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestPortPacketThresholdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestPortPacketThresholdDataSourceMetadata(t *testing.T) {
	d := NewPortPacketThresholdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_port_packet_threshold" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_port_packet_threshold")
	}
}
