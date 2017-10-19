
var game = new Phaser.Game(800, 600, null, null, { preload: preload, create: create });
function preload() {
}
function create() {
	game.load.onFileComplete.add(function(__, key) {
		console.log('key='+key)
	})
}