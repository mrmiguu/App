package app

import (
	"crypto/md5"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mrmiguu/app/js"
)

var (
	Addr = "localhost:80"

	Width = 800

	Height = 600

	temp = " "
	osfl sync.Mutex

	once  sync.Once
	serve sync.Mutex

	upgr = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	fatal = make(chan error)
)

func init() {
	println("[removing old...]")
	must(os.RemoveAll(temp))
	println("[removing old  !]")

	println("[creating new...]")
	must(os.Mkdir(temp, os.ModePerm))
	println("[creating new  !]")

	println("[serving...]")
	serve.Lock()
}

func bind() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("[connection...]")
		serve.Lock()
		serve.Unlock()
		println("[connection  !]")

		http.ServeFile(w, r, temp)
	})

	http.HandleFunc("/"+temp+"/_", onConnection)

	go func() {
		err := http.ListenAndServe(Addr, nil)
		if err != nil {
			fatal <- err
		}
	}()
}

func onConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgr.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for {
		t, b, err := conn.ReadMessage()
		if err != nil || t != websocket.BinaryMessage {
			return
		}
		println("packet:", string(b))
	}
}

func NewGroup(nameA, nameB string, names ...string) {
	strings.Join(append([]string{nameA, nameB}, names...), "▼")
}

type Group struct {
	// Live
}

type Image struct {
	lock          *sync.RWMutex
	data          image.Image
	x, y          int
	width, height int
	key           string
	v             js.Var
}

func AddImage(url string) (Image, error) {
	once.Do(bind)

	var i Image
	var r io.Reader

	f, err := os.Open(url)
	if err != nil {
		println("failed to read `" + url + "`")
		resp, err := http.Get(url)
		if err != nil {
			println("failed to get `" + url + "`")
			return i, err
		}
		defer resp.Body.Close()
		r = resp.Body
	} else {
		defer f.Close()
		r = f
	}

	img, _, err := image.Decode(r)
	if err != nil {
		return i, err
	}

	osfl.Lock()
	key := findRand()
	f2, err := os.Create(temp + "/" + key)
	osfl.Unlock()
	if err != nil {
		println("failed to write `" + key + "`")
		return i, err
	}
	defer f2.Close()

	ext := strings.LastIndex(url, ".")
	if ext == -1 {
		return i, errors.New("unknown image type")
	}

	switch strings.ToLower(url[ext:]) {
	case ".jpeg":
		fallthrough
	case ".jpg":
		err = jpeg.Encode(f2, img, nil)
	case ".png":
		err = png.Encode(f2, img)
	default:
		return i, errors.New("unsupported image type")
	}
	if err != nil {
		println("failed to encode `" + key + "`")
		return i, err
	}

	i.lock = &sync.RWMutex{}
	i.data = img

	i.x = Width / 2
	i.y = Height / 2

	size := img.Bounds().Size()
	i.width = size.X
	i.height = size.Y

	i.key = key
	i.v = js.AddImage(key)

	return i, nil
}

func (i *Image) Pos() (int, int) {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.x, i.y
}

func (i *Image) Size() (int, int) {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.width, i.height
}

func (i *Image) Show(b bool, d ...time.Duration) {
	a := 1
	if !b {
		a = 0
	}
	js.Tween(i.v, js.O{"alpha": a}, toMS(d...))
}

func (i *Image) Move(x, y int, d ...time.Duration) {
	i.lock.Lock()
	defer i.lock.Unlock()

	js.Tween(i.v, js.O{"x": x, "y": y}, toMS(d...))

	i.x = x
	i.y = y
}

func (i *Image) Resize(width, height int, d ...time.Duration) {
	i.lock.Lock()
	defer i.lock.Unlock()

	js.Tween(i.v, js.O{"width": width, "height": height}, toMS(d...))

	i.width = width
	i.height = height
}

func must(err error) {
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
	once.Do(bind)

	phaserminjs, err := ioutil.ReadFile("phaser.min.js")
	must(err)

	osfl.Lock()
	must(ioutil.WriteFile(temp+"/phaser.min.js", phaserminjs, os.ModePerm))
	must(ioutil.WriteFile(temp+"/index.html", htmlIndex(), os.ModePerm))
	must(ioutil.WriteFile(temp+"/styles.css", cssStyles(), os.ModePerm))

	js.Width = strconv.Itoa(Width)
	js.Height = strconv.Itoa(Height)
	b := js.Compile()

	must(ioutil.WriteFile(temp+"/main.js", b, os.ModePerm))
	osfl.Unlock()

	serve.Unlock()
	println("[serving  !]")

	if err := <-fatal; err != nil {
		panic(err)
	}
}
