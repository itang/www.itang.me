package gaemvc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"appengine"
	"appengine/urlfetch"
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func Handler(action IAction) http.Handler {
	return &actionWrapper{action}
}

type actionWrapper struct {
	Action IAction
}

func (this *actionWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vi := newFromInstance(this.Action)
	//vi.Elem().FieldByName("Action").Set(reflect.ValueOf(Action{w, r}))
	action, _ := vi.Interface().(IAction)

	action.Init(w, r)

	action.Apply() //vi.MethodByName("Apply").Call(nil)
}

type IAction interface {
	Init(w http.ResponseWriter, r *http.Request)
	Apply()
}

type Action struct {
	Resp http.ResponseWriter
	Req  *http.Request
}

func (this *Action) Init(w http.ResponseWriter, r *http.Request) {
	this.Resp = w
	this.Req = r
}

func (this *Action) Apply() {
	log.Panicln("Should Implement this Interface!")
}

func (this *Action) Render() *Render {
	return NewRender(this.Resp, this.Req)
}

func (this *Action) SendError(err string, status int) {
	http.Error(this.Resp, err, status)
}

func (this *Action) Send(content interface{}) {
	fmt.Fprint(this.Resp, content)
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
	t := oauth.Transport{
		Token:     &oauth.Token{AccessToken: "458a1257f03e0f5d65ebe83783349ce55007c89b"},
		Transport: &urlfetch.Transport{Context: appengine.NewContext(r)},
	}

	return github.NewClient(t.Client())
}

func newFromInstance(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	t := reflect.Indirect(v).Type()
	return reflect.New(t)
}
