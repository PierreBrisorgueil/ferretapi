package main

/**
 * Module dependencies
 */
import (
	"context"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"time"
	"os/exec"
	"io/ioutil"
	"bytes"

	"github.com/gobs/args"
	"github.com/raff/godet"
	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	httpDriver "github.com/MontFerret/ferret/pkg/drivers/http"
)

var remote *godet.RemoteDebugger

type text struct {
	Query string `json:"text"`
}

type errResult struct {
	Code int `json:"code"`
	Type string `json:"type"`
	Message string `json:"message"`
	Description string `json:"description"`
	Ip string `json:"ip"`
}

type successResult struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Data string `json:"data"`
	Ip string `json:"ip"`
}

/**
 * Main
 */

func main() {
	log.Println(chrome("/usr/bin/google-chrome"))
	//log.Println(chrome("/Applications/Google\\ Chrome.app/Contents/MacOS/Google\\ Chrome"))
	http.HandleFunc("/", reqHandler)
	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}

/**
 * @desc request handler
 */
func reqHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			// 1. Decode json body
			d := json.NewDecoder(r.Body)
			p := &text{}
			err := d.Decode(p)

			// 2. Get external ip for log
			ip, err := getIp()

			// 3. Magic Ferret stuff 
			data, err := ferret(p.Query)

			// 4. Answer
			if err != nil {
				var r *errResult = &errResult{
					Code: 422,
					Type: "error",
					Message: "Unprocessable Entity",
					Description:  string(err.Error()),
					Ip: ip,
				}
				j, _ := json.Marshal(r)
				w.WriteHeader(422)
				w.Write(j)
			} else {
				var r *successResult = &successResult{
					Type: "success",
					Message:  "ferret executed without errors ",
					Ip: ip,
					Data: string(data),
				}
				j, _ := json.Marshal(r)
				w.Write(j)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "I can't do that, please send me a Post request { text : 'your ferret things' } ")
	}
}

/**
 * @desc launch chrome and wait for it's availabality
 * @param {String} chrome path
 * @return {String} chrome status
 * some chrome options : 
 * path += " --single-process"
 * path += " --disable-dev-shm-usage"
 * path += " --ignore-certificate-errors"
 * path += " --enable-logging"
 * path += " --log-level=0"
 * path += " --disable-application-cache"
 * path += " --hide-scrollbars"
 */
 func chrome(path string) string {
	path += " --no-sandbox"
	path += " --headless"
	path += " --disable-gpu"
	path += " --remote-debugging-port=9222"
	path += " --disable-dev-shm-usage"

	if errRun := command(path, "start"); errRun != nil {
		// log.Println("cannot start browser", fmt.Sprint(errRun))
		return "cannot start browser" + fmt.Sprint(errRun)
	}
	var err error
	for i := 0; i < 20; i++ {
		if i > 0 {
			time.Sleep(500 * time.Millisecond)
		}
		remote, err = godet.Connect("localhost:9222", false)
		if err == nil {
			break
		}
		// log.Println("Error", err)
	}
	if err != nil {
		return "cannot connect to browser"
	}
	return "google ok"
}


/**
 * @desc ewecute command and arguments 
 * @input {String} command
 * @input {kind} kind of execution, 'run' - wait result, 'start' background
 * @return {error} error
 */
func command(input string, kind string) error {
	parts := args.GetArgs(input)
	cmd := exec.Command(parts[0], parts[1:]...)
	switch kind {
	case "run":
		return cmd.Run()
	default:
		return cmd.Start()
	}
}


/**
 * @desc launch ferret command and return return 
 * @param {String} term command for launch fql
 * @return {String} result
 */
func ferret(text string) ([]byte, error) {
	comp := compiler.New()
	program, err := comp.Compile(text)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	ctx = drivers.WithContext(ctx, cdp.NewDriver())
	ctx = drivers.WithContext(ctx, httpDriver.NewDriver(), drivers.AsDefault())
	out, err := program.Run(ctx)
	if err != nil {
		return nil, err
	}
	return out, nil
}

/**
 * @desc get external ip adress
 * @return {String, error}
 */
 func getIp() (string, error) {
	rsp, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(buf)), nil
}