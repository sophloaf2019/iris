package fieldwork

import (
	"iris/application/auth"
	"iris/application/fieldwork"
	fieldworkTypes "iris/domain/types/fieldwork"
	"iris/routes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Configure(g *gin.RouterGroup, fieldSvc *fieldwork.Service, authSvc *auth.Service) {
	g.GET("goi", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		if idStr, ok := c.GetQuery("id"); ok {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
				return
			}
			goi, err := fieldSvc.GetGOI(ctx, id)
			c.JSON(routes.SmartResponse(goi, err))
			return
		}
		if slug, ok := c.GetQuery("slug"); ok {
			goi, err := fieldSvc.GetGOIBySlug(ctx, slug)
			c.JSON(routes.SmartResponse(goi, err))
			return
		}
		gois, err := fieldSvc.GetGOIs(ctx)
		c.JSON(routes.SmartResponse(gois, err))
	})
	g.POST("goi", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		var goi fieldworkTypes.GOI
		if err := c.ShouldBind(&goi); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(routes.SmartResponse(goi, fieldSvc.PostGOI(ctx, &goi)))
	})
	g.PUT("goi/:id", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		var goi fieldworkTypes.GOI
		if err := c.ShouldBindJSON(&goi); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		goi.ID = id
		c.JSON(routes.SmartResponse(goi, fieldSvc.UpdateGOI(ctx, &goi)))
	})
	g.DELETE("goi/:id", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		strID := c.Param("id")
		id, err := strconv.Atoi(strID)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(routes.SmartResponse(nil, fieldSvc.DeleteGOI(ctx, id)))
	})

	g.GET("mission", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		if idStr, ok := c.GetQuery("id"); ok {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
				return
			}
			mission, err := fieldSvc.GetMission(ctx, id)
			c.JSON(routes.SmartResponse(mission, err))
			return
		}
		if slug, ok := c.GetQuery("slug"); ok {
			mission, err := fieldSvc.GetMissionBySlug(ctx, slug)
			c.JSON(routes.SmartResponse(mission, err))
			return
		}
		missions, err := fieldSvc.GetMissions(ctx)
		c.JSON(routes.SmartResponse(missions, err))
	})
	g.POST("mission", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		var mission fieldworkTypes.Mission
		if err := c.ShouldBind(&mission); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(routes.SmartResponse(mission, fieldSvc.PostMission(ctx, &mission)))
	})
	g.PUT("mission/:id", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		var mission fieldworkTypes.Mission
		if err := c.ShouldBindJSON(&mission); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		mission.ID = id
		c.JSON(routes.SmartResponse(mission, fieldSvc.UpdateMission(ctx, &mission)))
	})
	g.DELETE("mission/:id", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		strID := c.Param("id")
		id, err := strconv.Atoi(strID)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(routes.SmartResponse(nil, fieldSvc.DeleteMission(ctx, id)))
	})

	g.GET("debrief", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		if idStr, ok := c.GetQuery("id"); ok {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
				return
			}
			debrief, err := fieldSvc.GetDebrief(ctx, id)
			c.JSON(routes.SmartResponse(debrief, err))
			return
		}
		if slug, ok := c.GetQuery("slug"); ok {
			debrief, err := fieldSvc.GetDebriefBySlug(ctx, slug)
			c.JSON(routes.SmartResponse(debrief, err))
			return
		}
		debriefs, err := fieldSvc.GetDebriefs(ctx)
		c.JSON(routes.SmartResponse(debriefs, err))
	})
	g.POST("debrief", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		var debrief fieldworkTypes.Debrief
		if err := c.ShouldBind(&debrief); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(routes.SmartResponse(debrief, fieldSvc.PostDebrief(ctx, &debrief)))
	})
	g.PUT("debrief/:id", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		var debrief fieldworkTypes.Debrief
		if err := c.ShouldBindJSON(&debrief); err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		debrief.ID = id
		c.JSON(routes.SmartResponse(debrief, fieldSvc.UpdateDebrief(ctx, &debrief)))
	})
	g.DELETE("debrief/:id", routes.RequireAuth(authSvc), func(c *gin.Context) {
		ctx, err := routes.AuthContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		strID := c.Param("id")
		id, err := strconv.Atoi(strID)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.ErrorResponse(err))
			return
		}
		c.JSON(routes.SmartResponse(nil, fieldSvc.DeleteDebrief(ctx, id)))
	})
}
