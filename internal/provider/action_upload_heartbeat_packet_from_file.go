package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadHeartbeatPacketFromFileAction)(nil)

// UploadHeartbeatPacketFromFileAction is the generated Terraform action implementation.
type UploadHeartbeatPacketFromFileAction struct {
}

// UploadHeartbeatPacketFromFileActionModel describes the action configuration shape.
type UploadHeartbeatPacketFromFileActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	HbPacket  types.String `tfsdk:"hb_packet" json:"hbPacket"`
}

// NewUploadHeartbeatPacketFromFileAction returns a new instance of the generated action.
func NewUploadHeartbeatPacketFromFileAction() action.Action {
	return &UploadHeartbeatPacketFromFileAction{}
}

// Metadata returns the action type name.
func (r *UploadHeartbeatPacketFromFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_heartbeat_packet_from_file"
}

// Schema returns the action schema.
func (r *UploadHeartbeatPacketFromFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload a Heartbeat Packet from local file", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target Heartbeat Profile", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "hb_packet": schema.StringAttribute{MarkdownDescription: "Attached Heartbeat Packet. In native binary format", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadHeartbeatPacketFromFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadHeartbeatPacketFromFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
