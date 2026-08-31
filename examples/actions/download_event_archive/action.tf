action "gigavuecore_download_event_archive" "example" {
  config {
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
