package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"iris/application/auth"
	"iris/application/fieldwork"
	"iris/application/social"
	authTypes "iris/domain/types/auth"
	"iris/domain/types/base"
	social2 "iris/domain/types/social"
	authImplementations "iris/infra/cache/auth"
	fieldworkImplementations "iris/infra/cache/fieldwork"
	socialImplementations "iris/infra/cache/social"
	"iris/infra/random"
	authRoutes "iris/routes/auth"
	fieldworkRoutes "iris/routes/fieldwork"
	socialRoutes "iris/routes/social"
	"time"
)

func main() {
	// SERVICES

	authSvc := auth.NewService(
		random.TokenGenerator{},
		authImplementations.NewSessionManager(),
		random.Hasher{},
		authImplementations.NewUserRepository(),
	)
	fieldSvc := fieldwork.NewService(
		fieldworkImplementations.NewDebriefRepo(),
		fieldworkImplementations.NewGOIRepo(),
		fieldworkImplementations.NewMissionRepo(),
	)
	socialSvc := social.NewService(
		socialImplementations.NewPostRepo(),
		socialImplementations.NewCommentRepo(),
	)

	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	rg := g.Group("/api")

	// ROUTES

	authRoutes.Configure(rg, authSvc)
	fieldworkRoutes.Configure(rg, fieldSvc, authSvc)
	socialRoutes.Configure(rg, socialSvc, authSvc)

	// ADMIN CREATION
	user, _ := authSvc.IssueNewUser(*authTypes.NewAdminContext(), "sophia_t", "1234", authTypes.ClearanceTopSecret)
	otherUser, _ := authSvc.IssueNewUser(*authTypes.NewAdminContext(), "test", "1234", authTypes.ClearanceConfidential)

	fmt.Println(user.ID)
	fmt.Println(otherUser.ID)

	// SAMPLE POST
	newCtx := authTypes.NewContext(user, "")
	_ = socialSvc.MakePost(*newCtx, &social2.Post{
		Entity:  base.Entity{},
		Title:   "Test post!",
		UserID:  1,
		Message: "Hey everyone :)\nTest with a newline",
	})

	_ = socialSvc.MakePost(*newCtx, &social2.Post{
		Entity: base.Entity{
			CreatedAt: time.Now(),
		},
		Title:   "Test post 2!",
		UserID:  1,
		Message: "Hey everyone :)\nTest with a newline",
	})

	_ = socialSvc.MakeComment(*newCtx, &social2.Comment{
		Entity:     base.Entity{},
		Message:    "Hey everyone :)\nTest with a newline",
		ParentType: "post",
		ParentID:   1,
		UserID:     1,
	})
	_ = socialSvc.MakeComment(*newCtx, &social2.Comment{
		Entity:     base.Entity{},
		Message:    "Hey everyone :)\nTest with a newline, and some more stuff, and some more...",
		ParentType: "post",
		ParentID:   1,
		UserID:     1,
	})

	// RUN

	g.Run(":8080")
}
