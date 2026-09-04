resource "gigavuecore_pcap_profile" "example" {
  alias_list = "example"
  cluster_id = "example"
  pcap_configs = [{
    alias        = "example"
    channel_port = "example"
    direction    = "example"
    packet_limit = 1
    pcap_rules_list = [{
      dscp = "af11"
      dst_mac = {
        address = "example"
        mask    = "example"
      }
      dstipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      ether_type       = "example"
      inner_vlan       = 1
      ip4_frag         = "noFrag"
      ip4_ttl          = 0
      ip_ver           = "v4"
      packet_hit_count = 0
      portdst          = 0
      portsrc          = 0
      protocol         = 0
      rule_id          = 1
      src_mac = {
        address = "example"
        mask    = "example"
      }
      srcipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      tcpctl = 0
      vlan   = 1
    }]
    port = "example"
  }]
  port_ids = "example"
}
