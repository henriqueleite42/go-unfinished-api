package files

import (
	"io"
)

type Image = io.Reader

type ImageUsecase interface {
	CreateFromURL(URL string) (string, error)
}

type ImageRepository interface {
	CreateFromURL(URL string) (string, error)
}
