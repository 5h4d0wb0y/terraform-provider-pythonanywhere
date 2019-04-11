package pythonanywhere

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClientWith(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		t.Fatalf("err:")
	}

	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header.Get("Content-Type"))
	//fmt.Println(string(body))
	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
