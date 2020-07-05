package app

import (
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/controllers/activity"
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/controllers/sections"
)

func mapUrls() {
	// sections end point
	sectionsGroup := router.Group("/sections")
	sectionsGroup.Use(middleware.Auth())
	{
		sectionsGroup.POST("", sections.Create)
		sectionsGroup.GET("/:section_id", sections.Get)
		sectionsGroup.PUT("/:section_id", sections.Update)
		sectionsGroup.PATCH("/:section_id", sections.Update)
		sectionsGroup.DELETE("/:section_id", sections.Delete)
	}

	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.GET("/sections/:course_id", sections.GetAll)
		internalGroup.DELETE("/sections/:course_id", sections.DeleteAll)
	}

	activityGroup := router.Group("/activity")
	activityGroup.Use(middleware.Auth())
	{
		activityGroup.POST("/", activity.Create)
		activityGroup.GET("/:activity_id", activity.Get)
		activityGroup.PUT("/:activity_id", activity.Update)
		activityGroup.PATCH("/:activity_id", activity.Update)
		activityGroup.DELETE("/:activity_id", activity.Delete)
	}
}
