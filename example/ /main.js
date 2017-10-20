
var game = new Phaser.Game(800, 600, null, null, { preload: preload, create: create });
function preload(){
}
function create(){
game.load.onFileComplete.add(function(__, key) {
	console.log('key='+key)
})
game.load.image('d41d8cd98f00b204e9800998ecf8427e.png', 'd41d8cd98f00b204e9800998ecf8427e.png');
game.load.start();
game.load.image('74be16979710d4c4e7c6647856088456.png', '74be16979710d4c4e7c6647856088456.png');
game.load.start();
game.load.image('acf7ef943fdeb3cbfed8dd0d8f584731.png', 'acf7ef943fdeb3cbfed8dd0d8f584731.png');
game.load.start();
game.load.image('5a8dccb220de5c6775c873ead6ff2e43.png', '5a8dccb220de5c6775c873ead6ff2e43.png');
game.load.start();
game.load.image('76682f743ae018364a082b2e87f2d2f5.png', '76682f743ae018364a082b2e87f2d2f5.png');
game.load.start();
game.load.image('d41d8cd98f00b204e9800998ecf8427e.jpg', 'd41d8cd98f00b204e9800998ecf8427e.jpg');
game.load.start();
game.load.image('74be16979710d4c4e7c6647856088456.jpg', '74be16979710d4c4e7c6647856088456.jpg');
game.load.start();}
function tween(obj, to, ms, fn) {
	var t = game.add.tween(obj);
	t.to(to, ms);
	t.frameBased = true;
	t.onComplete.add(fn);
	t.start();
}