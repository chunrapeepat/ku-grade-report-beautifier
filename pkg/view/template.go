package view

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	tpIndex  = parseTemplate("layout/root.tmpl", "index.tmpl")
	tpReport = parseTemplate("layout/root.tmpl", "report.tmpl")
)

const templateDir = "template"

func joinTemplateDir(files ...string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	_, err := t.ParseFiles(joinTemplateDir(files...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}
