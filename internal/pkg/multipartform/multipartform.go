package multipartform

const (
	MultipartFormField = iota
	MultipartFormFile
)

type MultipartField struct {
	Type int
	Name string
}

type MultipartForm map[MultipartField]string

func New() MultipartForm {
	return make(MultipartForm)
}

func (m MultipartForm) AddField(name string, value string) {
	m[MultipartField{Type: MultipartFormField, Name: name}] = value
}

func (m MultipartForm) AddFile(name string, file string) {
	m[MultipartField{Type: MultipartFormFile, Name: name}] = file
}
