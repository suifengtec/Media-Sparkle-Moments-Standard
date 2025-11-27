package msp

import (
	"encoding/json"
	"errors"
	"os"
)

type Seg struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Slug    string `json:"slug"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

type MediaSparkleMoments struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Path  string `json:"path"`
	Segs  []Seg  `json:"list"`
}

func GetObjByJsonPath(mspJsonPath string) (MediaSparkleMoments, error) {
	var ret MediaSparkleMoments
	bs, err := os.ReadFile(mspJsonPath)
	if err != nil {
		return ret, err
	}
	var argInJson MediaSparkleMoments
	err = json.Unmarshal(bs, &argInJson)
	if err != nil {
		return ret, err
	}
	if argInJson.Path == "" {
		return ret, errors.New("media path is empty")
	}
	if !IsFile(argInJson.Path) {
		return ret, errors.New("media path is not a file")
	}
	return argInJson, nil
}

func GetMediaPathByJsonPath(mspJsonPath string) (string, error) {
	bs, err := os.ReadFile(mspJsonPath)
	if err != nil {
		return "", err
	}
	var argInJson MediaSparkleMoments
	err = json.Unmarshal(bs, &argInJson)
	if err != nil {
		return "", err
	}
	if argInJson.Path == "" {
		return "", errors.New("media path is empty")
	}
	if !IsFile(argInJson.Path) {
		return "", errors.New("media path is not a file")
	}

	return argInJson.Path, nil
}

func GetSegListByJsonPath(mspJsonPath string) ([]Seg, error) {

	bs, err := os.ReadFile(mspJsonPath)
	if err != nil {
		return nil, err
	}
	var argInJson MediaSparkleMoments
	err = json.Unmarshal(bs, &argInJson)
	if err != nil {
		return nil, err
	}
	if argInJson.Path == "" {
		return nil, errors.New("media path is empty")
	}
	if !IsFile(argInJson.Path) {
		return nil, errors.New("media path is not a file")
	}
	return argInJson.Segs, nil
}
