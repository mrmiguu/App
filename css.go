package app

const (
	cssStylesBody = `
body {
	background: #000000;
	margin: 0;
	visibility: hidden;
}`
)

func cssStyles() []byte {
	return []byte(`` +
		cssStylesBody)
}
