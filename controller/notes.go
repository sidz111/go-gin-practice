package controller

import "github.com/gin-gonic/gin"

type NoteController struct {
}

func (n *NoteController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.PostNotes())
}
func (n *NoteController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Msg": "Get Method From Notes",
		})
	}
}

func (n *NoteController) PostNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Msg": "Post Method From Notes",
		})
	}
}
