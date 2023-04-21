package book

import "errors"

func ShowBookInfo(title, author string) (string, error) {
	if title == "" {
		return "", errors.New("title is empty")
	}
	if author == "" {
		return "", errors.New("author is empty")
	}
	return "《 " + title + " 》" + ", " + author, nil
}
