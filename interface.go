package codes

type ReadWriter interface {
	Read() string
	Write(string)
}

type Person struct {
}

func (Person) Read() string {
	return ""
}

func (Person) Write(string) {

}

//var rw ReadWriter = Person{}

var rw ReadWriter = Person{}
