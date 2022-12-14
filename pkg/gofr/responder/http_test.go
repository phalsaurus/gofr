package responder

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr/template"
	"developer.zopsmart.com/go/gofr/pkg/log"
	"developer.zopsmart.com/go/gofr/pkg/middleware"
)

func TestNewContextualResponder(t *testing.T) {
	var (
		w             = httptest.NewRecorder()
		correlationID = "50deb4921d7bc1b441bb992c4874a147"
	)

	path := "/dummy"
	testCases := []struct {
		contentType         string
		correlationIDHeader string
		want                Responder
	}{
		{"", "X-B3-TraceID", &HTTP{w: w, resType: JSON, method: "GET", path: path, correlationID: correlationID}},
		{"text/xml", "X-B3-TraceID", &HTTP{w: w, resType: XML, method: "GET", path: path, correlationID: correlationID}},
		{"application/xml", "X-B3-TraceID", &HTTP{w: w, resType: XML, method: "GET", path: path, correlationID: correlationID}},
		{"text/json", "X-Correlation-ID", &HTTP{w: w, resType: JSON, method: "GET", path: path, correlationID: correlationID}},
		{"application/json", "X-Correlation-ID", &HTTP{w: w, resType: JSON, method: "GET", path: path, correlationID: correlationID}},
		{"text/plain", "X-Correlation-ID", &HTTP{w: w, resType: TEXT, method: "GET", path: path, correlationID: correlationID}},
	}

	for _, tc := range testCases {
		r := httptest.NewRequest("GET", "/dummy", nil)

		// handler to set the routeKey in request context
		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			r = req
		})

		muxRouter := mux.NewRouter()
		muxRouter.NewRoute().Path(r.URL.Path).Methods("GET").Handler(handler)
		muxRouter.ServeHTTP(w, r)

		*r = *r.Clone(context.WithValue(r.Context(), middleware.CorrelationIDKey, correlationID))

		r.Header.Set("Content-Type", tc.contentType)
		r.Header.Set(tc.correlationIDHeader, correlationID)

		if got := NewContextualResponder(w, r); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("NewContextualResponder() = %v, want %v", got, tc.want)
		}
	}
}

func TestHTTP_Respond(t *testing.T) {
	createDefaultTemplate()

	defer deleteDefaultTemplate()

	data := struct {
		Title string
		Items []string
	}{
		Title: "Default Gofr Template",
		Items: []string{
			"Welcome to Gofr",
		},
	}

	w := httptest.NewRecorder()

	type args struct {
		statusCode int
		data       interface{}
	}

	testCases := []struct {
		resType responseType
		args    args
		want    string
	}{
		{resType: 999, args: args{statusCode: 500, data: template.Template{Directory: "./", File: "default.html",
			Data: "test data", Type: template.HTML}}, want: ""},
		{resType: 999, args: args{statusCode: 200, data: template.Template{Directory: "./", File: "default.html",
			Data: data, Type: template.HTML}}, want: "text/html"},
		{JSON, args{200, `{"name": "gofr"}`}, "application/json"},
		{XML, args{200, `<name>gofr</name>`}, "application/xml"},
		{TEXT, args{200, `name: gofr`}, "text/plain"},
		{TEXT, args{200, template.File{Content: []byte(`<html></html>`), ContentType: "text/html"}}, "text/html"},
	}

	for _, tc := range testCases {
		h := HTTP{
			w:       w,
			resType: tc.resType,
		}
		h.Respond(tc.args.data, nil)

		if got := h.w.Header().Get("Content-Type"); got != tc.want {
			t.Errorf("got %v, want: %v", got, tc.want)
		}
	}
}

func createDefaultTemplate() {
	rootDir, _ := os.Getwd()
	logger := log.NewLogger()
	f, err := os.Create(rootDir + "/default.html")

	if err != nil {
		logger.Error(err)
	}

	_, err = f.WriteString(`<!DOCTYPE html>
	<html>
	<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	</head>
	<body>
	{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
	</html>`)

	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("Template created!")
	}

	err = f.Close()
	if err != nil {
		logger.Error(err)
	}
}

func deleteDefaultTemplate() {
	rootDir, _ := os.Getwd()
	logger := log.NewLogger()
	err := os.Remove(rootDir + "/default.html")

	if err != nil {
		logger.Error(err)
	}
}

func TestHTTP_Respond_PartialError(t *testing.T) {
	w := httptest.NewRecorder()

	type args struct {
		statusCode int
		data       interface{}
		err        error
	}

	testCases := []struct {
		resType responseType
		args    args
		want    string
	}{
		{JSON, args{206, map[string]interface{}{"name": "Alice"}, errors.EntityNotFound{
			Entity: "store",
			ID:     "1",
		}}, "application/json"},
	}

	for _, tc := range testCases {
		h := HTTP{
			w:       w,
			resType: tc.resType,
		}
		h.Respond(tc.args.data, tc.args.err)

		if got := h.w.Header().Get("Content-Type"); got != tc.want {
			t.Errorf("got %v, want: %v", got, tc.want)
		}
	}
}
