resource "gigavuecore_pcap_profile" "example" {
  alias_list = "example"
  cluster_id = "example"
  pcap_configs = [{
    alias        = "example"
    channel_port = "example"
    direction    = "example"
    packet_limit = 0
    pcap_rules_list = [{
      dscp = "example"
      dst_mac = {
        address = "example"
        mask    = "example"
      }
      dstipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      ether_type       = "example"
      inner_vlan       = 0
      ip4_frag         = "example"
      ip4_ttl          = 0
      ip_ver           = "example"
      packet_hit_count = 0
      portdst          = 0
      portsrc          = 0
      protocol         = 0
      rule_id          = 0
      src_mac = {
        address = "example"
        mask    = "example"
      }
      srcipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      tcpctl = 0
      vlan   = 0
    }]
    port = "example"
  }]
  port_ids = "example"
}
