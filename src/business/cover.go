package business

import (
	"bytes"
	"gold/src/configs"
	"gold/src/exception"
	"gold/src/logger"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
	"strconv"
)

var conf = configs.Get()
var log = logger.Get()

// Cover : Implements the Cover service methods.
var Cover = &coverService{}

type coverService struct{}

func (c *coverService) GetRandomCover(args map[string]interface{}) (*http.Response, error) {
	info, err := ioutil.ReadDir(conf.Path.Cover)
	if err != nil {
		log.Sugar().Errorf("Unexpected error while reading Cover dir: %s", err.Error())
		return nil, err
	}

	if len(info) == 0 {
		return nil, exception.CoverNotFound("")
	}

	randomFileName := info[rand.Intn(len(info))].Name()
	randomCover, err := ioutil.ReadFile(path.Join(conf.Path.Cover, randomFileName))
	if err != nil {
		log.Sugar().Errorf("Unexpected error while reading cover image file: %s", err.Error())
		return nil, err
	}

	height, exists := args["height"]
	if !exists {
		height = "0"
	}
	width, exists := args["width"]
	if !exists {
		width = "0"
	}

	intHeight, _ := strconv.ParseInt(height.(string), 10, 64)
	intWidth, _ := strconv.ParseInt(width.(string), 10, 64)

	compressed, err := imageService.Resize(randomCover, intHeight, intWidth)
	if err != nil {
		log.Sugar().Errorf("Unexpected error while resizing image: %s", err.Error())
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(compressed)),
	}, nil
}
