package main

import (
	"fmt"
	"log"
	"net/http"

	wf "github.com/writeas/go-webfinger"

	"github.com/mrbotchi-team/mrbotchi/app"
	"github.com/mrbotchi-team/mrbotchi/handlers"
	"github.com/mrbotchi-team/mrbotchi/webfinger"
)

var (
	version  string
	revision string
)

func PrintWakeupMessage() {
	fmt.Println("==========================================================================================")
	fmt.Print("\n")
	fmt.Println("Developed by silverscat_3")
	fmt.Println(" _ __ ___  _ __")
	fmt.Println("| '_ ` _ \\| '__|")
	fmt.Println("| | | | | | |")
	fmt.Println("|_| |_| |_|_|          _       _     _")
	fmt.Println("          | |         | |     | |   (_)")
	fmt.Println("          | |__   ___ | |_ ___| |__  _")
	fmt.Println("          | '_ \\ / _ \\| __/ __| '_ \\| |")
	fmt.Println("          | |_) | (_) | || (__| | | | |")
	fmt.Println("          |_.__/ \\___/ \\__\\___|_| |_|_|")
	fmt.Println("Version:", version, ",", "Revision:", revision)
	fmt.Print("\n")
	fmt.Println("==========================================================================================")
}

func main() {
	PrintWakeupMessage()

	app := app.NewApp()

	hostMeta := webfinger.HostMetaHandler{Host: app.Config.Host}
	app.Router.Get("/.well-known/host-meta", handlers.HandlerFunc(hostMeta.Get).ServeHTTP)

	webfinger := wf.Default(webfinger.WebfingerResolver{UserName: app.Config.User.Name, Host: app.Config.Host})
	webfinger.NoTLSHandler = nil
	app.Router.Get(wf.WebFingerPath, http.HandlerFunc(webfinger.Webfinger))

	hs := handlers.HandlerFactory(app)
	for endpoint, h := range hs {
		app.Router.Get(endpoint, handlers.HandlerFunc(h.Get).ServeHTTP)
		app.Router.Post(endpoint, handlers.HandlerFunc(h.Post).ServeHTTP)
		app.Router.Put(endpoint, handlers.HandlerFunc(h.Put).ServeHTTP)
		app.Router.Delete(endpoint, handlers.HandlerFunc(h.Delete).ServeHTTP)
	}

	fmt.Println("I'm HTTP listen on :3000. Have a nice day!")
	log.Fatalln(http.ListenAndServe(":3000", app.Router))
}
