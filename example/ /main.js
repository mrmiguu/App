
var files = [];
var game = new Phaser.Game(800, 600, null, null, { create: create });
function create(){
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
});
(function() {
	var i = files.length;
	files.push({ key: 'd41d8cd98f00b204e9800998ecf8427e' });
	files[i].promise = new Promise(function(ok) {
		files[i].ok = ok;
		game.load.image('d41d8cd98f00b204e9800998ecf8427e', 'd41d8cd98f00b204e9800998ecf8427e');
		game.load.start();
	});
	files[i].promise.then(function() {
		var obj = game.add.sprite(game.world.centerX, game.world.centerY, 'd41d8cd98f00b204e9800998ecf8427e');
		obj.anchor.setTo(0.5, 0.5);
		obj.bringToTop();
		console.log('file #'+i+' loaded');
	});
})();
(function() {
	var i = files.length;
	files.push({ key: '74be16979710d4c4e7c6647856088456' });
	files[i].promise = new Promise(function(ok) {
		files[i].ok = ok;
		game.load.image('74be16979710d4c4e7c6647856088456', '74be16979710d4c4e7c6647856088456');
		game.load.start();
	});
	files[i].promise.then(function() {
		var obj = game.add.sprite(game.world.centerX, game.world.centerY, '74be16979710d4c4e7c6647856088456');
		obj.anchor.setTo(0.5, 0.5);
		obj.bringToTop();
		console.log('file #'+i+' loaded');
	});
})();
(function() {
	var i = files.length;
	files.push({ key: 'acf7ef943fdeb3cbfed8dd0d8f584731' });
	files[i].promise = new Promise(function(ok) {
		files[i].ok = ok;
		game.load.image('acf7ef943fdeb3cbfed8dd0d8f584731', 'acf7ef943fdeb3cbfed8dd0d8f584731');
		game.load.start();
	});
	files[i].promise.then(function() {
		var obj = game.add.sprite(game.world.centerX, game.world.centerY, 'acf7ef943fdeb3cbfed8dd0d8f584731');
		obj.anchor.setTo(0.5, 0.5);
		obj.bringToTop();
		console.log('file #'+i+' loaded');
	});
})();
(function() {
	var i = files.length;
	files.push({ key: '5a8dccb220de5c6775c873ead6ff2e43' });
	files[i].promise = new Promise(function(ok) {
		files[i].ok = ok;
		game.load.image('5a8dccb220de5c6775c873ead6ff2e43', '5a8dccb220de5c6775c873ead6ff2e43');
		game.load.start();
	});
	files[i].promise.then(function() {
		var obj = game.add.sprite(game.world.centerX, game.world.centerY, '5a8dccb220de5c6775c873ead6ff2e43');
		obj.anchor.setTo(0.5, 0.5);
		obj.bringToTop();
		console.log('file #'+i+' loaded');
	});
})();}
function tween(obj, to, ms, fn) {
	var t = game.add.tween(obj);
	t.to(to, ms);
	t.frameBased = true;
	t.onComplete.add(fn);
	t.start();
}