package gaemvc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"appengine"
	"appengine/urlfetch"
	"github.com/google/go-github/github"
)

type IAction interface {
	Apply()
}

type Handler struct {
	Action IAction
}

func (this *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vi := newFromInstance(this.Action)
	vi.Elem().FieldByName("Action").Set(reflect.ValueOf(Action{w, r}))
	action, _ := vi.Interface().(IAction)

	action.Apply() //vi.MethodByName("Apply").Call(nil)
}

type Action struct {
	Writer http.ResponseWriter
	Req    *http.Request
}

func (this *Action) Apply() {
	log.Panicln("Should Implement this Interface!")
}

func (this *Action) Render() *Render {
	return NewRender(this.Writer, this.Req)
}

func (this *Action) SendError(err string, status int) {
	http.Error(this.Writer, err, status)
}

func (this *Action) Send(content string) {
	fmt.Fprintf(this.Writer, content)
}

type GaeAction struct {
	Action
}

func (this *GaeAction) GaeContext() appengine.Context {
	return appengine.NewContext(this.Req)
}

func (this *GaeAction) HttpClient() *http.Client {
	return urlfetch.Client(this.GaeContext())
}

type GaeGithubAction struct {
	GaeAction
}

func (this *GaeGithubAction) GithubClient() *github.Client {
	return GithubClient(this.Req)
}

// --------------------------------------------------------

func GithubClient(r *http.Request) *github.Client {
	httpClient := urlfetch.Client(appengine.NewContext(r))
	return github.NewClient(httpClient)
}

func newFromInstance(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	t := reflect.Indirect(v).Type()
	return reflect.New(t)
}
