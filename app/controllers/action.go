package controllers

import (
	"log"
	"net/http"
	"reflect"

	"appengine"
	"appengine/urlfetch"
	"github.com/google/go-github/github"

	"app/views"
)

type IAction interface {
	Apply()
}

type Handler struct {
	Action IAction
}

type Action struct {
	Writer http.ResponseWriter
	Req    *http.Request
}

func (this *Action) Apply() {
	log.Panicln("Should Implement this Interface!")
}

func (this *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := reflect.ValueOf(this.Action)
	t := reflect.Indirect(v).Type()
	vi := reflect.New(t)

	vi.Elem().FieldByName("Action").Set(reflect.ValueOf(Action{w, r}))
	vi.MethodByName("Apply").Call(nil)
}

func (this *Action) Render() *views.Render {
	return views.New(this.Writer, this.Req)
}

func (this *Action) GithubClient() *github.Client {
	return GithubClient(this.Req)
}

func GithubClient(r *http.Request) *github.Client {
	httpClient := urlfetch.Client(appengine.NewContext(r))
	return github.NewClient(httpClient)
}
