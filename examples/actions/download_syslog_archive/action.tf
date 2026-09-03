action "gigavuecore_download_syslog_archive" "example" {
  config {
    cluster_id = "example"
    destination = {
      sftp = {
        file_path    = "example"
        host_address = "example"
        password     = "example"
        username     = "example"
      }
    }
    purge = true
    scope = {
      start_date = "example"
    }
  }
}
