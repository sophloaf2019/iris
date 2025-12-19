package social

import (
	"errors"
	"iris/application/auth"
	"iris/application/social"
	socialTypes "iris/domain/types/social"
	"iris/routes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Configure(
	g *gin.RouterGroup,
	socialSvc *social.Service,
	authSvc *auth.Service,

) {
	authMW := routes.RequireAuth(authSvc)

	// ---- Posts ----

	g.GET("posts", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		posts, err := socialSvc.GetPosts(ctx)
		c.JSON(routes.SmartResponse(posts, err))
	})

	g.GET("posts/:id", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		post, err := socialSvc.GetPost(ctx, id)
		c.JSON(routes.SmartResponse(post, err))
	})

	g.POST("posts", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		var post socialTypes.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		c.JSON(routes.SmartResponse(post, socialSvc.MakePost(ctx, &post)))
	})

	g.PUT("posts/:id", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		var post socialTypes.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		post.ID = id
		c.JSON(routes.SmartResponse(post, socialSvc.SavePost(ctx, &post)))
	})

	g.DELETE("posts/:id", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		c.JSON(routes.SmartResponse(nil, socialSvc.DeletePost(ctx, id)))
	})

	// ---- Comments ----

	g.GET("comments", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		// Check if "id" is provided
		if idStr := c.Query("id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
				return
			}

			comment, err := socialSvc.GetComment(ctx, id)
			c.JSON(routes.SmartResponse(comment, err))
			return
		}

		// Otherwise, require parent_id and content_type
		parentIDStr := c.Query("parent_id")
		contentType := c.Query("content_type")

		if parentIDStr == "" || contentType == "" {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(
				errors.New("either id or both parent_id and content_type are required"),
			))
			return
		}

		parentID, err := strconv.Atoi(parentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		comments, err := socialSvc.GetCommentsFor(ctx, parentID, contentType)
		c.JSON(routes.SmartResponse(comments, err))
	})

	g.POST("comments", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		var comment socialTypes.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		c.JSON(routes.SmartResponse(comment, socialSvc.MakeComment(ctx, &comment)))
	})

	g.PUT("comments/:id", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		var comment socialTypes.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		comment.ID = id
		c.JSON(routes.SmartResponse(comment, socialSvc.SaveComment(ctx, &comment)))
	})

	g.DELETE("comments/:id", authMW, func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}

		c.JSON(routes.SmartResponse(nil, socialSvc.DeleteComment(ctx, id)))
	})
}
