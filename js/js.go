package js

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var (
	images = make(chan map[string]bool, 1)

	preloadBody = make(chan string, 1)
	preloadTail = make(chan string, 1)
	createBody  = make(chan string, 1)
	createTail  = make(chan string, 1)
)

type Var string
type O map[string]interface{}

func (obj O) String() string {
	var fields []string
	for k, v := range obj {
		fields = append(fields, k+`:`+fmt.Sprint(v))
	}
	return `{` + strings.Join(fields, `,`) + `}`
}

func init() {
	images <- map[string]bool{}

	preloadBody <- ``
	preloadTail <- ``
	createBody <- ``
	createTail <- ``
}

func Tween(v Var, obj O, ms int) {
	createBody <- <-createBody + `
tween(` + string(v) + `, ` + obj.String() + `, ` + strconv.Itoa(ms) + `);`
}

func AddImage(key string) Var {
	imgs := <-images
	if _, found := imgs[key]; !found {
		imgs[key] = true

		preloadBody <- <-preloadBody + `
loadImage('` + key + `');`

	}
	images <- imgs
	v := `_` + key

	createBody <- <-createBody + `
var ` + v + ` = addImage('` + key + `')`

	return Var(v)
}

func compile() []byte {
	return []byte(`

var game = new Phaser.Game(800, 600, null, null, { preload: preload, create: create });

`)
}

func compilePreload() []byte {
	return []byte(`

function preload() {
	` + <-preloadBody + `
	` + <-preloadTail + `
}

`)
}

func compileCreate() []byte {
	return []byte(`

function create() {
	` + <-createBody + `
	` + <-createTail + `
}

`)
}

func compileFunctions() []byte {
	return []byte(`

function loadImage(key) {
	game.load.image(key, key);
}

function addImage(key) {
	var img = game.add.image(game.world.centerX, game.world.centerY, key);
	img.alpha = 0;
	img.anchor.setTo(0.5, 0.5);
	return img
}

function tween(obj, to, ms, fn) {
	var t = game.add.tween(obj);
	t.to(to, ms);
	t.frameBased = true;
	// t.onComplete.add(fn);
	t.start();
}

`)
}

func Compile() []byte {
	return bytes.Join([][]byte{
		compile(),
		compilePreload(),
		compileCreate(),
		compileFunctions(),
	}, nil)
}

func gid() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	gid, _ := strconv.ParseUint(string(b), 10, 64)
	return gid
}
