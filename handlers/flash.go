package handlers

var flash *Flash

func init() {
	flash = &Flash{}
}

type Flash struct {
	Type string
	Message string
}

func(f *Flash) setFlash(_type, msg string) {
	f.Type = _type
	f.Message = msg
}

func (f *Flash) delete() {
	f.Type = ""
	f.Message = ""
}