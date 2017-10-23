package app

const (
	htmlIndexHead = `
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0">
	<link rel="stylesheet" href="styles.css">
	<script src="phaser.min.js"></script>
	<script src="main.js"></script>
</head>`

	htmlIndexBody = `
<body>
	<div id="myCanvas"></div>
</body>`
)

func htmlIndex() []byte {
	return []byte(`
<html>` +
		htmlIndexHead +
		htmlIndexBody + `
</html>`)
}
