package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Zerohated/heart-beat-checker/internal/model"
	"github.com/gin-gonic/gin"
)

// EchoHandler return 200 and empty response with internal code 0
func (controller *Controller) EchoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, newRespOK(CodeOK, nil))
	return
}

type RegisterReq struct {
	UID      int    `json:"uid"`
	UserName string `json:"username"`
	Message  string `json:"message"`
	Path     string `json:"path"`
}

func (controller *Controller) RegisterUser(c *gin.Context) {
	var err error
	req := RegisterReq{}
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, newRespErr(fmt.Sprintf("bind json failed: %s", err.Error())))
		return
	}
	now := time.Now()
	user := model.User{
		UID:      req.UID,
		Username: req.UserName,
		Message:  req.Message,
		Path:     req.Path,
		LastSeen: &now,
	}
	err = model.CreateOrUpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newRespErr(fmt.Sprintf("db error: %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, newRespOK(CodeOK, "ok"))
	return
}

func (controller *Controller) GetUserList(c *gin.Context) {
	var err error
	users := []*model.User{}
	users, err = model.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, newRespErr(fmt.Sprintf("db error: %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, newRespOK(CodeOK, users))
	return
}
