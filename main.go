package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Q struct {
	Text    string
	Choices []string
	Answer  int
}

var quiz = []Q{
	{"rune 型の正体は？", []string{"uint8", "int32", "string", "byte"}, 1},
	{"len(\"あ\") の結果は？", []string{"1", "2", "3", "4"}, 2},
}

const html = `<!DOCTYPE html>
<html lang="ja"><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css">
<script src="https://unpkg.com/htmx.org"></script>
<style>main{animation:f .4s}@keyframes f{from{opacity:0;transform:translateY(8px)}}</style>
</head><body>
<main class="container">
{{if .Done}}
  <h1>{{.Score}} / {{.Total}} 正解！🎉</h1>
  <a href="/" role="button">もう一度</a>
{{else}}
  <h1>🦫 Go クイズ</h1>
  <form hx-post="/submit" hx-target="main" hx-swap="outerHTML">
  {{range $i, $q := .Quiz}}
    <fieldset>
      <legend><strong>{{$q.Text}}</strong></legend>
      {{range $j, $c := $q.Choices}}
      <label><input type="radio" name="q{{$i}}" value="{{$j}}"> {{$c}}</label>
      {{end}}
    </fieldset>
  {{end}}
  <button type="submit">採点</button>
  </form>
{{end}}
</main>
</body></html>`

var tmpl = template.Must(template.New("").Parse(html))

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, map[string]any{"Quiz": quiz})
	})

	mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		score := 0
		for i, q := range quiz {
			if v, _ := strconv.Atoi(r.FormValue(fmt.Sprintf("q%d", i))); v == q.Answer {
				score++
			}
		}
		tmpl.Execute(w, map[string]any{
			"Done": true, "Score": score, "Total": len(quiz),
		})
	})

	http.ListenAndServe(":8080", mux)
}
