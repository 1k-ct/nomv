package user

import (
	"net/http"

	"github.com/1k-ct/clonefile/nomv/src/needmov/db"
	youtubeapi "github.com/1k-ct/nomv/src/needmov/youtubeAPI"
	"github.com/gin-gonic/gin"
)

// CreateVideoInfo "/regVideo" PrintVideoInfoをdbに登録
func (pc Controller) CreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("videoURL")

	videoID, videoName, videoDescription, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL)
	db.InsertVideoInfo(videoID, videoName, videoDescription, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)
	c.Redirect(http.StatusFound, "/") // http.StatusFound = 302
}

// CreateChannelInfo "/regchannel" PrintChannelInfoをdbに登録
func (pc Controller) CreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("channelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount := youtubeapi.PrintChannelInfo(channelURL)
	db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	c.Redirect(http.StatusFound, "/") // http.StatusFound = 302
}
