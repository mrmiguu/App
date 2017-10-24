package my

type Space struct{}
type Image struct{}

func New(key string) Space {
	return Space{}
}

func (s *Space) AddImage(url string) Image {
	return Image{}
}
