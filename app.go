package app

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mrmiguu/app/js"
)

var (
	Addr = "localhost:80"

	temp = " "
	osfl sync.Mutex
)

func init() {
	Must(os.RemoveAll(temp))
	Must(os.Mkdir(temp, os.ModePerm))
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func findRand() string {
	var name string
	for {
		name = fmt.Sprintf("%x", md5.Sum([]byte(name)))
		if _, err := os.Stat(temp + "/" + name); os.IsNotExist(err) {
			break
		}
	}
	return name
}

type Image struct {
	key string
	v   js.Var
}

func AddImage(url string) (Image, error) {
	var img Image

	b, err := ioutil.ReadFile(url)
	if err != nil {
		println("failed to read `" + url + "`")
		resp, err := http.Get(url)
		if err != nil {
			println("failed to get `" + url + "`")
			return img, err
		}
		defer resp.Body.Close()
		b, err = ioutil.ReadAll(resp.Body)
	}

	osfl.Lock()
	key := findRand()
	err = ioutil.WriteFile(temp+"/"+key, b, os.ModePerm)
	osfl.Unlock()
	if err != nil {
		println("failed to write `" + key + "`")
		return img, err
	}

	img.key = key
	img.v = js.AddImage(key)
	return img, nil
}

func (i *Image) Pos() (int, int) {
	return -1, -1
}

func (i *Image) Size() (int, int) {
	return -1, -1
}

func (i *Image) Show(b bool, d ...time.Duration) {
	ms := toMS(d...)
	a := 1
	if !b {
		a = 0
	}
	js.Tween(i.v, js.O{"alpha": a}, ms)
}

func (i *Image) Resize(width, height int, d ...time.Duration) {
}

func toMS(d ...time.Duration) int {
	if len(d) >= 2 {
		panic("too many arguments")
	}
	dur := 1
	if len(d) == 1 {
		dur = int(d[0].Seconds() * 1000)
	}
	return dur
}

func Serve() {
	phaserminjs, err := ioutil.ReadFile("phaser.min.js")
	Must(err)
	Must(ioutil.WriteFile(temp+"/phaser.min.js", phaserminjs, os.ModePerm))
	Must(ioutil.WriteFile(temp+"/index.html", htmlIndex(), os.ModePerm))
	Must(ioutil.WriteFile(temp+"/styles.css", cssStyles(), os.ModePerm))
	Must(ioutil.WriteFile(temp+"/main.js", js.Compile(), os.ModePerm))
	http.Handle("/", http.FileServer(http.Dir(temp)))

	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/"+temp+"/_", func(w http.ResponseWriter, r *http.Request) {
		conn, err := up.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for {
			t, b, err := conn.ReadMessage()
			if err != nil || t != websocket.BinaryMessage {
				return
			}
			println(string(b))
		}
	})

	log.Fatal(http.ListenAndServe(Addr, nil))
}
