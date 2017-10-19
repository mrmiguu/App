package app

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Addr = "localhost:80"
	Temp = "www"

	run sync.Once
	api string
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func findRand(ext string) string {
	var name string
	for {
		name = fmt.Sprintf("%x", md5.Sum([]byte(name)))
		if _, err := os.Stat(Temp + "/" + name + ext); os.IsNotExist(err) {
			break
		}
	}
	return name + ext
}

func new() {
	if end := len(Temp) - 1; Temp[end] == '/' {
		Temp = Temp[:end]
	}
	Must(os.RemoveAll(Temp))
	Must(os.MkdirAll(Temp, os.ModePerm))
	Must(ioutil.WriteFile("index.html", []byte(htmlDefaultIndex), os.ModePerm))
	http.Handle("/", http.FileServer(http.Dir(Temp)))

	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/"+Temp+"/.", func(w http.ResponseWriter, r *http.Request) {
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

	go http.ListenAndServe(Addr, nil)
}

type Image struct {
	key string
}

func LoadImage(url string) <-chan *Image {
	run.Do(new)

	c := make(chan *Image)
	go func() {
		b, err := ioutil.ReadFile(url)
		if err != nil {
			resp, err := http.Get(url)
			if err != nil {
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
		key := findRand(ext)

		err = ioutil.WriteFile(Temp+"/"+key, b, os.ModePerm)
		if err != nil {
			c <- nil
			return
		}

		c <- &Image{key}
	}()
	return c
}
