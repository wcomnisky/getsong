package getsong

import (
	"fmt"
	"strings"
	"testing"

	"github.com/iawia002/lux/extractors"
	"github.com/iawia002/lux/extractors/youtube"
	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func TestUseAnnie(t *testing.T) {
	datas, err := youtube.New().Extract("https://www.youtube.com/watch?v=qxiOMm_x3Xg", extractors.Options{})
	assert.Nil(t, err)
	biggestSize := int64(0)
	bestURL := ""
	exten := ""
	for _, data := range datas {
		fmt.Printf("%+v\n", data)
		for _, stream := range data.Streams {
			fmt.Printf("%+v\n", stream)
			if strings.Contains(stream.Quality, "audio/mp4") {
				fmt.Printf("%+v\n", stream.Parts[0])
				if stream.Parts[0].Size > biggestSize {
					bestURL = stream.Parts[0].URL
					biggestSize = stream.Parts[0].Size
					exten = stream.Parts[0].Ext
				}
			}
		}
	}
	fmt.Println(bestURL)
	fmt.Println(exten)
}

func TestGetSongAPI(t *testing.T) {
	_, err := GetSong("Old Records", "Allen Toussaint", Options{
		ShowProgress: true,
		Debug:        true,
	})
	assert.Nil(t, err)
	_, err = GetSong("Eva", "Haerts", Options{
		ShowProgress: true,
		Debug:        true,
	})
	assert.Nil(t, err)
}

func TestGetPage(t *testing.T) {

	log.SetLevel("debug")
	html, err := getPage("https://www.youtube.com/watch?v=qxiOMm_x3Xg")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(html, "<html"))
}
func TestGetFfmpeg(t *testing.T) {

	OptionShowProgressBar = true

	locationToBinary, err := getFfmpegBinary()
	fmt.Println(locationToBinary)
	assert.NotEqual(t, "", locationToBinary)
	assert.Nil(t, err)
}

func TestGetYouTubeInfo(t *testing.T) {
	log.SetLevel("debug")
	info, err := getYoutubeVideoInfo("qxiOMm_x3Xg")
	assert.Nil(t, err)
	fmt.Printf("%+v\n", info)
}

func TestOne(t *testing.T) {
	_, err := GetSong("Old Records", "Allen Toussaint", Options{
		ShowProgress: true,
		Debug:        true,
	})
	assert.Nil(t, err)
}

func TestGetMusicVideoID(t *testing.T) {
	log.SetLevel("trace")

	// this one is tricky because the band name is spelled weird and requires
	// clicking through to force youtube to search the wrong spelling
	id, err := GetMusicVideoID("eva", "haerts")
	log.Infof("eva: %s", id)
	assert.Nil(t, err)
	assert.Equal(t, "qxiOMm_x3Xg", id)

	id, err = GetMusicVideoID("movies", "Weyes Blood")
	log.Infof("movies: %s", id)
	assert.Nil(t, err)
	assert.True(t, id == "RFtRq6t3jOo" || id == "xniRJsus8pk")

	// this one is trick because its the second result
	id, err = GetMusicVideoID("old records", "allen toussaint")
	log.Infof("old records: %s", id)
	assert.Nil(t, err)
	assert.True(t, id == "oa6KzRfvtAs" || id == "obtJEJ4VPmk")

	// try one with puncuation
	id, err = GetMusicVideoID("hey, ma", "bon iver")
	log.Infof("hey, ma: %s", id)
	assert.Nil(t, err)
	assert.True(t, id == "HDAKS18Gv1U")

	// skip the most popular result to get the provided to youtube version
	id, err = GetMusicVideoID("true", "spandau ballet")
	log.Infof("true: %s", id)
	assert.Nil(t, err)
	assert.True(t, id == "ITX-SEsyGRg" || id == "2H1N6KdU-L0" || id == "sWBueqYA2Es" || id == "TVeSwMMvkP4")

	// pick one that is not the first
	id, err = GetMusicVideoID("i know what love is", "don white")
	log.Infof("i know what: %s", id)
	assert.Nil(t, err)
	assert.Equal(t, "3LRu9mjiyKo", id)
}

func TestParseDurationString(t *testing.T) {
	assert.Equal(t, int64(470001), ParseDurationString("00:07:50.01,"))
}
