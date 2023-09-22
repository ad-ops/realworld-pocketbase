// Code generated by templ@v0.2.334 DO NOT EDIT.

package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"home-page\"><div class=\"banner\"><div class=\"container\"><h1 class=\"logo-font\">")
		if err != nil {
			return err
		}
		var_2 := `conduit`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p>")
		if err != nil {
			return err
		}
		var_3 := `A place to share your knowledge.`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div></div><div class=\"container page\"><div class=\"row\"><div class=\"col-md-9\"><div class=\"feed-toggle\"><ul class=\"nav nav-pills outline-active\"><li class=\"nav-item\"><a class=\"nav-link\" href=\"\">")
		if err != nil {
			return err
		}
		var_4 := `Your Feed`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li><li class=\"nav-item\"><a class=\"nav-link active\" href=\"\">")
		if err != nil {
			return err
		}
		var_5 := `Global Feed`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul></div><div hx-get=\"/hx/articles\" hx-trigger=\"load\" hx-swap=\"innerHTML\">")
		if err != nil {
			return err
		}
		var_6 := `loading articles...`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><ul class=\"pagination\"><li class=\"page-item active\"><a class=\"page-link\" href=\"\">")
		if err != nil {
			return err
		}
		var_7 := `1`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li><li class=\"page-item\"><a class=\"page-link\" href=\"\">")
		if err != nil {
			return err
		}
		var_8 := `2`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul></div><div class=\"col-md-3\"><div class=\"sidebar\"><p>")
		if err != nil {
			return err
		}
		var_9 := `Popular Tags`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><div hx-get=\"/hx/tags\" hx-trigger=\"load\" hx-swap=\"innerHTML\" class=\"tag-list\">")
		if err != nil {
			return err
		}
		var_10 := `loading tags...`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div></div></div></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
