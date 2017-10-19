package app

const (
	jsMainPreload = `
function preload() {
}`

	jsMainCreate = `
function create() {
	game.load.onFileComplete.add(function(__, key) {
		console.log('key='+key)
	})
}`

	jsMainTween = `
function tween(obj, to, ms, fn) {
	var t = game.add.tween(obj);
	t.to(to, ms);
	t.frameBased = true;
	t.onComplete.add(fn);
	t.start();
}`
)

func jsMain() []byte {
	return []byte(`
var game = new Phaser.Game(800, 600, null, null, { preload: preload, create: create });` +
		jsMainPreload +
		jsMainCreate +
		jsMainTween)
}
