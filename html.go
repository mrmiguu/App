package app

const (
	defaultHTML = `
	<html>

	<body style="margin:0;">
		<script src="phaser.min.js"></script>
		<script src="main.js"></script>
	</body>

	</html>`

	defaultJSPreloadBegin = `
	function preload() {`

	defaultJSPreload = `
	game.load.image('einstein', 'd41d8cd98f00b204e9800998ecf8427e.jpg');`

	defaultJSPreloadEnd = `
	}`

	defaultJSCreateBegin = `
	function create() {`

	defaultJSCreate = `
	var s = game.add.sprite(80, 0, 'einstein');
	s.rotation = 0.14;`

	defaultJSCreateEnd = `
	}`

	defaultJS = `
	var game = new Phaser.Game(800, 600, Phaser.CANVAS, '', { preload: preload, create: create });` +

		defaultJSPreloadBegin +
		defaultJSPreload +
		defaultJSPreloadEnd +

		defaultJSCreateBegin +
		defaultJSCreate +
		defaultJSCreateEnd
)
