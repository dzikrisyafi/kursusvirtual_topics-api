package activity

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/domain/activity"
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var activity activity.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.ActivityService.CreateActivity(activity)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	activityID, idErr := controller_utils.GetIDInt(c.Param("activity_id"), "activity id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	activity, getErr := services.ActivityService.GetActivity(activityID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, activity)
}

func Update(c *gin.Context) {
	activityID, idErr := controller_utils.GetIDInt(c.Param("activity_id"), "activity id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var activity activity.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	activity.ID = activityID
	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := services.ActivityService.UpdateActivity(isPartial, activity)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	activityID, idErr := controller_utils.GetIDInt(c.Param("activity_id"), "acitivty id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.ActivityService.DeleteActivity(activityID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted activity", "status": http.StatusOK})
}
