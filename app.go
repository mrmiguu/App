package app

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	Addr = "localhost:80"

	temp = " "
	osfl sync.Mutex
)

func init() {
	Must(os.RemoveAll(temp))
	Must(os.MkdirAll(temp, os.ModePerm))
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func findRand(ext string) string {
	var name string
	for {
		name = fmt.Sprintf("%x", md5.Sum([]byte(name)))
		if _, err := os.Stat(temp + "/" + name + ext); os.IsNotExist(err) {
			break
		}
	}
	return name + ext
}

func Serve() {
	phaserminjs, err := ioutil.ReadFile("phaser.min.js")
	Must(err)
	Must(ioutil.WriteFile(temp+"/phaser.min.js", phaserminjs, os.ModePerm))
	Must(ioutil.WriteFile(temp+"/index.html", htmlIndex(), os.ModePerm))
	Must(ioutil.WriteFile(temp+"/styles.css", cssStyles(), os.ModePerm))
	Must(ioutil.WriteFile(temp+"/main.js", jsMain(), os.ModePerm))
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

type Image struct {
	key string
}

func LoadImage(url string) <-chan *Image {
	c := make(chan *Image)

	go func() {
		b, err := ioutil.ReadFile(url)
		if err != nil {
			println("failed to read `" + url + "`")
			resp, err := http.Get(url)
			if err != nil {
				println("failed to get `" + url + "`")
				c <- nil
				return
			}
			defer resp.Body.Close()
			b, err = ioutil.ReadAll(resp.Body)
		}

		var ext string
		idx := strings.LastIndex(url, ".")
		if idx != -1 {
			ext = url[idx:]
		}

		osfl.Lock()
		key := findRand(ext)
		err = ioutil.WriteFile(temp+"/"+key, b, os.ModePerm)
		osfl.Unlock()
		if err != nil {
			println("failed to write `" + key + "`")
			c <- nil
			return
		}

		c <- &Image{key}
	}()

	return c
}

func (i *Image) Pos() (int, int) {
	return -1, -1
}

func (i *Image) Size() (int, int) {
	return -1, -1
}

func (i *Image) Show(b bool, d ...time.Duration) {
}

func (i *Image) Resize(width, height int, d ...time.Duration) {
}

func duration(d ...time.Duration) int {
	if len(d) >= 2 {
		panic("too many arguments")
	}
	dur := 1
	if len(d) == 1 {
		dur = int(d[0].Seconds() * 1000)
	}
	return dur
}
