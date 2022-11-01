package gurl

import (
	"fmt"
	"net/http"
	"strings"
)

type wrappedBuilder struct {
	prefix string
	strings.Builder
}

func (w *wrappedBuilder) WriteHeaders(headers http.Header) {
	for k, vs := range headers {
		for _, v := range vs {
			w.Printf("%v: %v", k, v)
		}
	}
}

func (w *wrappedBuilder) Println() {
	w.WriteString("\n")
}

func (w *wrappedBuilder) Printf(s string, a ...any) {
	w.WriteString(fmt.Sprintf("%v %v\n", w.prefix, fmt.Sprintf(s, a...)))
}