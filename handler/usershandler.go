package user_handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.learning.com/go2025/common"
	"go.learning.com/go2025/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	bcryptCost = 12
)

func CreatedUserHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var userCreationModel models.UserCreation
		var userModel models.User

		//get data from request body
		if err := c.ShouldBindJSON(&userCreationModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//Validate input entity
		if err := userCreationModel.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//Make sure username is not existed
		//Make sure email is not existed

		encyptedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreationModel.Password), bcryptCost)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"encrypted_password_error": err.Error(),
			})
			return
		}

		//set usermodel data
		userModel.Email = userCreationModel.Email
		userModel.UserName = userCreationModel.UserName
		userModel.EncryptedPassword = string(encyptedPassword) //Must have function to Hash password
		userModel.CreatedAt = time.Now()
		userModel.UpdatedAt = time.Now()
		userModel.CreatedBy = "Admin"
		userModel.UpdatedBy = "Admin"
		userModel.DeleteFlag = false

		if err := db.Table("users").Create(&userModel).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Return JSON response
		c.IndentedJSON(http.StatusOK, gin.H{
			"success": userModel,
		})
	}
}

func GetAllUsersHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var result []models.User
		var paging common.Paging
		var total int64

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"get_paging_error": err.Error(),
			})
			return
		}

		paging.Preset()

		query := db.Table("users").Where("deleted_flag=?", false)

		if err := query.Select("id").Count(&total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"count_error": err.Error(),
			})
			return
		}

		if err := query.
			Offset((paging.Page - 1) * paging.Limit).
			Select("*").
			Limit(paging.Limit).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success": result,
			"paging":  paging,
			"total":   total,
		})

	}

}

func GetUserByIdHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var result models.User

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"convert_id_error": err.Error(),
			})
			return
		}

		if err := db.Table("users").Where("id=? AND deleted_flag = ?", id, false).First(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"get_user_error": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success": result,
		})

	}
}

func UpdateUserHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var userUpdateData models.UserUpdateParams

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"convert_id_error": err.Error(),
			})
			return
		}

		//Make sure user is existed --- enhance validation

		//get data from request body
		if err := c.ShouldBindJSON(&userUpdateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Table("users").Where("id=? AND deleted_flag = ?", id, false).Updates(&userUpdateData).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"update_user_error": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success": true,
		})

	}
}

func DeleteUserHandler(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"convert_id_error": err.Error(),
			})
			return
		}

		//Make sure user is existed --- enhance validation

		if err := db.Table("users").Where("id=? AND deleted_flag = ?", id, false).Updates(map[string]interface{}{
			"deleted_flag": true,
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"delete_user_error": err.Error(),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"success": true,
		})

	}
}
