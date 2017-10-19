package app

const (
	cssStylesBody = `
body {
	margin: 0;
}`
)

func cssStyles() []byte {
	return []byte(`` +
		cssStylesBody)
}
