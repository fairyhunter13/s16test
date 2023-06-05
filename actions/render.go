package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"io"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}

type jsonBytesRenderer struct {
	value []byte
}

func (s jsonBytesRenderer) ContentType() string {
	return "application/json; charset=utf-8"
}

func (s jsonBytesRenderer) Render(w io.Writer, _ render.Data) error {
	_, err := w.Write(s.value)
	return err
}

// JSONBytes renders the []byte value using the "application/json"
// content type.
func JSONBytes(v []byte) render.Renderer {
	return jsonBytesRenderer{value: v}
}
