package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/senomas/xgoapi/httputil"
	"github.com/senomas/xgoapi/model"
)

// ShowUser godoc
// @Summary Show an user
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [get]
func (c *Controller) ShowUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := model.UserOne(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary List users
// @Description get users
// @Tags users
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q" Format(email)
// @Success 200 {array} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users [get]
func (c *Controller) ListUsers(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	users, err := model.UsersAll(q)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// AddUser godoc
// @Summary Add an user
// @Description add by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.AddUser true "Add user"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users [post]
func (c *Controller) AddUser(ctx *gin.Context) {
	var addUser model.AddUser
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addUser.Validation(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	user := model.User{
		Name: addUser.Name,
	}
	err := user.Insert()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update an user
// @Description Update by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID"
// @Param  user body model.UpdateUser true "Update user"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [patch]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateUser model.UpdateUser
	var err error;
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	user := model.User{
		ID:   id,
		Name: updateUser.Name,
	}
	err = user.Update()
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete an user
// @Description Delete by user ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID" Format(int64)
// @Success 204 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := model.DeleteUser(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// UploadUserImage godoc
// @Summary Upload user image
// @Description Upload file
// @Tags users
// @Accept  multipart/form-data
// @Produce  json
// @Param  id path int true "User ID"
// @Param file formData file true "user image"
// @Success 200 {object} controller.Message
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id}/images [post]
func (c *Controller) UploadUserImage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload complete userID=%d filename=%s", id, file.Filename)})
}
