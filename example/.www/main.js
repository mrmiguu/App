
	var game = new Phaser.Game(800, 600, Phaser.CANVAS, '', { preload: preload, create: create });
	function preload() {
	game.load.image('einstein', 'd41d8cd98f00b204e9800998ecf8427e.jpg');
	}
	function create() {
	var s = game.add.sprite(80, 0, 'einstein');
	s.rotation = 0.14;
	}