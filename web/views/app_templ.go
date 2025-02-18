// Code generated by templ@v0.2.334 DO NOT EDIT.

package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func App(body templ.Component) templ.Component {
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
		_, err = templBuffer.WriteString("<head><meta charset=\"utf-8\"><title>")
		if err != nil {
			return err
		}
		var_2 := `Conduit`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link href=\"//code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css\" rel=\"stylesheet\" type=\"text/css\"><link href=\"//fonts.googleapis.com/css?family=Titillium+Web:700|Source+Serif+Pro:400,700|Merriweather+Sans:400,700|Source+Sans+Pro:400,300,600,700,300italic,400italic,600italic,700italic\" rel=\"stylesheet\" type=\"text/css\"><link rel=\"stylesheet\" href=\"//demo.productionready.io/main.css\"><script src=\"https://unpkg.com/htmx.org@1.9.5\">")
		if err != nil {
			return err
		}
		var_3 := ``
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body><nav class=\"navbar navbar-light\"><div class=\"container\"><a class=\"navbar-brand\" href=\"/\">")
		if err != nil {
			return err
		}
		var_4 := `conduit`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><ul class=\"nav navbar-nav pull-xs-right\"><li class=\"nav-item\"><a class=\"nav-link active\" href=\"/\">")
		if err != nil {
			return err
		}
		var_5 := `Home`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/login\">")
		if err != nil {
			return err
		}
		var_6 := `Sign in`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/register\">")
		if err != nil {
			return err
		}
		var_7 := `Sign up`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul></div></nav>")
		if err != nil {
			return err
		}
		err = body.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<footer><div class=\"container\"><a href=\"/\" class=\"logo-font\">")
		if err != nil {
			return err
		}
		var_8 := `conduit`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><span class=\"attribution\">")
		if err != nil {
			return err
		}
		var_9 := `An interactive learning project from `
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<a href=\"https://thinkster.io\">")
		if err != nil {
			return err
		}
		var_10 := `Thinkster`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a>")
		if err != nil {
			return err
		}
		var_11 := `. Code &amp;`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_12 := `design licensed under MIT.`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></div></footer></body>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
