package js

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var (
	Width = "800"

	Height = "600"

	images = make(chan map[string]bool, 1)

	preloadBody = make(chan string, 1)
	preloadTail = make(chan string, 1)
	createBody  = make(chan string, 1)
	createTail  = make(chan string, 1)

	lastGID   = make(chan string, 1)
	gidTweens = make(chan map[string]int, 1)
)

func init() {
	images <- map[string]bool{}

	preloadBody <- ``
	preloadTail <- ``
	createBody <- ``
	createTail <- ``

	gidTweens <- map[string]int{}
}

type Var string

type O map[string]interface{}

func (obj O) String() string {
	var fields []string
	for k, v := range obj {
		fields = append(fields, k+`:`+fmt.Sprint(v))
	}
	return `{` + strings.Join(fields, `,`) + `}`
}

func Tween(v Var, obj O, ms int) {
	body := <-createBody
	gids := <-gidTweens
	id := gid()
	n := gids[id]
	t := `tween_` + id + `x` + strconv.Itoa(n) + string(v)
	if len(gids) > 0 {
		lastv := <-lastGID

		createBody <- body + `
var ` + t + ` = new Promise(function(ok) {
	` + lastv + `.then(function() {
		tween(` + string(v) + `, ` + obj.String() + `, ` + strconv.Itoa(ms) + `, ok);
	});
});`

	} else {

		createBody <- body + `
		var ` + t + ` = new Promise(function(ok) {
			tween(` + string(v) + `, ` + obj.String() + `, ` + strconv.Itoa(ms) + `, ok);
			});`

	}
	gids[id] = n + 1
	lastGID <- t
	gidTweens <- gids
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
var ` + v + ` = addImage('` + key + `');`

	return Var(v)
}

func compile() []byte {
	return []byte(`

var game = new Phaser.Game(` + Width + `, ` + Height + `, null, null, { preload: preload, create: create });

`)
}

func compilePreload() []byte {
	return []byte(`

function preload() {
	game.canvas.oncontextmenu = function(e) { e.preventDefault(); };

	// var inW = window.innerWidth;
	// var inH = window.innerHeight;
	// if (` + Height + ` > ` + Width + `) {
	// 	var newW = (` + Width + ` / ` + Height + `) * inH;
	// 	game.scale.setMinMax(newW, inH, newW, inH);
	// } else {
	// 	var newH = (` + Height + ` / ` + Width + `) * inW;
	// 	game.scale.setMinMax(inW, newH, inW, newH);
	// }
	
	game.scale.scaleMode = Phaser.ScaleManager.SHOW_ALL;
	game.scale.pageAlignVertically = true;
	game.scale.pageAlignHorizontally = true;

	setTimeout(function() {
		document.body.style.visibility = 'visible';
	}, 200);

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
	t.onComplete.add(fn);
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

// func sameRoutine() bool {
// 	old := <-lastGID
// 	new := gid()
// 	b := new == old && old > 0
// 	lastGID <- new
// 	return b
// }

func gid() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}
