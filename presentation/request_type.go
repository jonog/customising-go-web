type Request struct {
      Method string
      URL *url.URL
      Header Header
      Body io.ReadCloser
      ContentLength int64
      Host string
      RemoteAddr string
      ...
}