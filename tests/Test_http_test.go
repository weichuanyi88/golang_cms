package tests

import (
    "reflect"
    "strings"
    "net/http"
    "testing"
    "fmt"
)

func Test_http(t *testing.T)  {
    req, _ := http.NewRequest("POST", "http://www.google.com/search",
        strings.NewReader("z=post&both=y&prio=2&empty=&sdljfsdj=112"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    fmt.Println(req.PostFormValue("sdljfsdj"))
    if q := req.FormValue("q"); q != "foo" {
        t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
    }
    if z := req.FormValue("z"); z != "post" {
        t.Errorf(`req.FormValue("z") = %q, want "post"`, z)
    }
    if bq, found := req.PostForm["q"]; found {
        t.Errorf(`req.PostForm["q"] = %q, want no entry in map`, bq)
    }
    if bz := req.PostFormValue("z"); bz != "post" {
        t.Errorf(`req.PostFormValue("z") = %q, want "post"`, bz)
    }
    if qs := req.Form["q"]; !reflect.DeepEqual(qs, []string{"foo", "bar"}) {
        t.Errorf(`req.Form["q"] = %q, want ["foo", "bar"]`, qs)
    }
    if both := req.Form["both"]; !reflect.DeepEqual(both, []string{"y", "x"}) {
        t.Errorf(`req.Form["both"] = %q, want ["y", "x"]`, both)
    }
    if prio := req.FormValue("prio"); prio != "2" {
        t.Errorf(`req.FormValue("prio") = %q, want "2" (from body)`, prio)
    }
    if empty := req.FormValue("empty"); empty != "" {
        t.Errorf(`req.FormValue("empty") = %q, want "" (from body)`, empty)
    }
}
