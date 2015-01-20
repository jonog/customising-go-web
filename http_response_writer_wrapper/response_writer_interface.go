import "net/http"

type ResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(int)
}