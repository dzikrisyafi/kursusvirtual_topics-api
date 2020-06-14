package app

import (
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/controllers/activity"
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/controllers/sections"
)

func mapUrls() {
	router.POST("/sections", sections.Create)
	router.GET("/sections/:section_id", sections.Get)
	router.PUT("/sections/:section_id", sections.Update)
	router.PATCH("/sections/:section_id", sections.Update)
	router.DELETE("/sections/:section_id", sections.Delete)
	router.GET("/internal/sections/:course_id", sections.GetAll)
	router.DELETE("/internal/sections/:course_id", sections.DeleteAll)

	router.POST("/activity", activity.Create)
	router.GET("/activity/:activity_id", activity.Get)
	router.PUT("/activity/:activity_id", activity.Update)
	router.PATCH("/activity/:activity_id", activity.Update)
	router.DELETE("/activity/:activity_id", activity.Delete)
}
