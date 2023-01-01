package chooseadventure

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

const defaultCYOATemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>
`

var defaultTpl *template.Template

// Called whenever module is imported
func init() {
	defaultTpl = template.Must(template.New("").Parse(defaultCYOATemplate))
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

type handler struct {
	story  Story
	tpl    *template.Template
	pathFn func(r *http.Request) string
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.tpl = t
	}
}

func WithPathFn(pathFn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = pathFn
	}
}

func NewHandler(story Story, opts ...HandlerOption) http.Handler {
	h := handler{story, defaultTpl, defaultPathFn}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.story[path]; ok {
		err := h.tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Chapter not found.", http.StatusNotFound)
	}
}
