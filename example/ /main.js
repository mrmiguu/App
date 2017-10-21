

var game = new Phaser.Game(800, 600, null, null, { preload: preload, create: create });



function preload() {
	
loadImage('d41d8cd98f00b204e9800998ecf8427e');
loadImage('74be16979710d4c4e7c6647856088456');
loadImage('acf7ef943fdeb3cbfed8dd0d8f584731');
loadImage('5a8dccb220de5c6775c873ead6ff2e43');
	
}



function create() {
	
var _d41d8cd98f00b204e9800998ecf8427e = addImage('d41d8cd98f00b204e9800998ecf8427e')
var _74be16979710d4c4e7c6647856088456 = addImage('74be16979710d4c4e7c6647856088456')
var _acf7ef943fdeb3cbfed8dd0d8f584731 = addImage('acf7ef943fdeb3cbfed8dd0d8f584731')
var _5a8dccb220de5c6775c873ead6ff2e43 = addImage('5a8dccb220de5c6775c873ead6ff2e43')
var p = new Promise(function(ok) {
	tween(_d41d8cd98f00b204e9800998ecf8427e, {alpha:1}, 1000, ok);
});
var p = new Promise(function(ok) {
	tween(_74be16979710d4c4e7c6647856088456, {alpha:1}, 1000, ok);
});
var p = new Promise(function(ok) {
	tween(_acf7ef943fdeb3cbfed8dd0d8f584731, {alpha:1}, 1000, ok);
});
var p = new Promise(function(ok) {
	tween(_5a8dccb220de5c6775c873ead6ff2e43, {alpha:1}, 1000, ok);
});
	
}



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

