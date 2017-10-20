package app

const (
	jsMainCreateHeader = `
function create()`
	jsMainCreateBody = `
game.load.onFileComplete.add(function(__, key) {
	console.log('key='+key);
	for (var i in files) {
		if (files[i].key != key) {
			continue;
		}
		files[i].ld()
	}
})`

	jsMainTween = `
function tween(obj, to, ms, fn) {
	var t = game.add.tween(obj);
	t.to(to, ms);
	t.frameBased = true;
	t.onComplete.add(fn);
	t.start();
}`
)

func jsLoadImage(key string) string {
	return `
var i = files.length;
files.push({
	key: '` + key + `',
	ld: function() {
		if (i > 0) {
			files[i-1].ld.then(ok);
		} else {
		}
	}
});
game.load.image('` + key + `', '` + key + `');
game.load.start();`
}

func jsMain(jsMainCreate string) []byte {
	return []byte(`
var files = [];
var game = new Phaser.Game(800, 600, null, null, { create: create });` +
		jsMainCreateHeader + `{` +
		jsMainCreateBody +
		jsMainCreate + `}` +

		jsMainTween)
}
