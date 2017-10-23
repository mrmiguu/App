

var game = new Phaser.Game(600, 800, null, null, { preload: preload, create: create });



function preload() {
	game.canvas.oncontextmenu = function(e) { e.preventDefault(); };

	// var inW = window.innerWidth;
	// var inH = window.innerHeight;
	// if (800 > 600) {
	// 	var newW = (600 / 800) * inH;
	// 	game.scale.setMinMax(newW, inH, newW, inH);
	// } else {
	// 	var newH = (800 / 600) * inW;
	// 	game.scale.setMinMax(inW, newH, inW, newH);
	// }
	
	game.scale.scaleMode = Phaser.ScaleManager.SHOW_ALL;
	game.scale.pageAlignVertically = true;
	game.scale.pageAlignHorizontally = true;

	setTimeout(function() {
		document.body.style.visibility = 'visible';
	}, 200);

	
loadImage('d41d8cd98f00b204e9800998ecf8427e');
	
}



function create() {
	
var _d41d8cd98f00b204e9800998ecf8427e = addImage('d41d8cd98f00b204e9800998ecf8427e');
		var tween_1x0_d41d8cd98f00b204e9800998ecf8427e = new Promise(function(ok) {
			tween(_d41d8cd98f00b204e9800998ecf8427e, {width:600,height:600}, 1, ok);
			});
var tween_1x1_d41d8cd98f00b204e9800998ecf8427e = new Promise(function(ok) {
	tween_1x0_d41d8cd98f00b204e9800998ecf8427e.then(function() {
		tween(_d41d8cd98f00b204e9800998ecf8427e, {alpha:1}, 2500, ok);
	});
});
var tween_1x2_d41d8cd98f00b204e9800998ecf8427e = new Promise(function(ok) {
	tween_1x1_d41d8cd98f00b204e9800998ecf8427e.then(function() {
		tween(_d41d8cd98f00b204e9800998ecf8427e, {x:300,y:300}, 5000, ok);
	});
});
var tween_1x3_d41d8cd98f00b204e9800998ecf8427e = new Promise(function(ok) {
	tween_1x2_d41d8cd98f00b204e9800998ecf8427e.then(function() {
		tween(_d41d8cd98f00b204e9800998ecf8427e, {alpha:0}, 2500, ok);
	});
});
var tween_6x0_d41d8cd98f00b204e9800998ecf8427e = new Promise(function(ok) {
	tween_1x3_d41d8cd98f00b204e9800998ecf8427e.then(function() {
		tween(_d41d8cd98f00b204e9800998ecf8427e, {width:-600,height:-600}, 5000, ok);
	});
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

