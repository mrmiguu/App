package app

const (
	htmlIndexHead = `
<head>
	<link rel="stylesheet" href="styles.css">
	<script src="phaser.min.js"></script>
	<script src="main.js"></script>
</head>`

	htmlIndexBody = `
<body>
</body>`
)

func htmlIndex() []byte {
	return []byte(`
<html>` +
		htmlIndexHead +
		htmlIndexBody + `
</html>`)
}
