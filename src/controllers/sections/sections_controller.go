package sections

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/domain/sections"
	"github.com/dzikrisyafi/kursusvirtual_topics-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var section sections.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.SectionsService.CreateSection(section)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusCreated("success created new section", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	sectionID, idErr := controller_utils.GetIDInt(c.Param("section_id"), "section id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	section, getErr := services.SectionsService.GetSection(sectionID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get section data", section.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAll(c *gin.Context) {
	courseID, idErr := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	sections, getErr := services.SectionsService.GetAllSectionByCourseID(courseID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get section data", sections.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	sectionID, idErr := controller_utils.GetIDInt(c.Param("section_id"), "section id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var section sections.Section
	if err := c.ShouldBindJSON(&section); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	section.ID = sectionID
	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := services.SectionsService.UpdateSection(isPartial, section)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusOK("success updated section data", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	sectionID, idErr := controller_utils.GetIDInt(c.Param("section_id"), "section id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.SectionsService.DeleteSection(sectionID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted section", "status": http.StatusOK})
}

func DeleteAll(c *gin.Context) {
	courseID, idErr := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.SectionsService.DeleteSectionByCourseID(courseID, c.Query("access_token")); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}
