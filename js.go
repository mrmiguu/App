package app

const (
	jsMainCreateHeader = `
function create()`

	jsMainCreateBody = `
game.load.onFileComplete.add(function(__, key) {
	for (var i in files) {
		if (files[i].key !== key) {
			continue;
		}
		if (i > 0) {
			files[i-1].promise.then().then(files[i].ok);
		} else {
			files[i].ok();
		}
	}
});`

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
(function() {
	var i = files.length;
	files.push({ key: '` + key + `' });
	files[i].promise = new Promise(function(ok) {
		files[i].ok = ok;
		game.load.image('` + key + `', '` + key + `');
		game.load.start();
	});
	files[i].promise.then(function() {
		var obj = game.add.sprite(game.world.centerX, game.world.centerY, '` + key + `');
		obj.anchor.setTo(0.5, 0.5);
		obj.bringToTop();
		console.log('file #'+i+' loaded');
	});
})();`
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
