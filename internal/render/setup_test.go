package render

import (
	"WebServer/internal/config"
	"WebServer/internal/models"
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

type myWriter struct{}

func (tw *myWriter) Header() http.Header{
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int){

}

func (tw *myWriter) Write(b []byte) (int, error){
	length := len(b)
	return length, nil
}

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())

}
