type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error) // HL
	WriteHeader(int)
}