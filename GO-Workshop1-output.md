# 📁 PROJECT EXPORT FOR LLMs

## 📊 Project Information

- **Project Name**: `GO-Workshop1`
- **Generated On**: 2026-01-31 05:13:36 (Asia/Bangkok / GMT+07:00)
- **Total Files Processed**: 9
- **Export Tool**: Easy Whole Project to Single Text File for LLMs v1.1.0
- **Tool Author**: Jota / José Guilherme Pandolfi

### ⚙️ Export Configuration

| Setting | Value |
|---------|-------|
| Language | `en` |
| Max File Size | `1 MB` |
| Include Hidden Files | `false` |
| Output Format | `both` |

## 🌳 Project Structure

```
├── 📁 common/
│   └── 📄 paging.go (251 B)
├── 📁 handler/
│   └── 📄 usershandler.go (4.5 KB)
├── 📁 models/
│   └── 📄 user.go (1.67 KB)
├── 📄 gitignore (558 B)
├── 📄 go.mod (2.04 KB)
├── 📄 go.sum (9.06 KB)
├── 📄 main.go (1.76 KB)
├── 📄 Makefile (299 B)
└── 📄 README.md (2.01 KB)
```

## 📑 Table of Contents

**Project Files:**

- [📄 common/paging.go](#📄-common-paging-go)
- [📄 handler/usershandler.go](#📄-handler-usershandler-go)
- [📄 models/user.go](#📄-models-user-go)
- [📄 main.go](#📄-main-go)
- [📄 README.md](#📄-readme-md)

---

## 📈 Project Statistics

| Metric | Count |
|--------|-------|
| Total Files | 9 |
| Total Directories | 3 |
| Text Files | 5 |
| Binary Files | 4 |
| Total Size | 22.12 KB |

### 📄 File Types Distribution

| Extension | Count |
|-----------|-------|
| `.go` | 4 |
| `no extension` | 2 |
| `.mod` | 1 |
| `.sum` | 1 |
| `.md` | 1 |

## 💻 File Code Contents

### <a id="📄-common-paging-go"></a>📄 `common/paging.go`

**File Info:**
- **Size**: 251 B
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `common/paging.go`
- **Relative Path**: `common`
- **Created**: 2026-01-31 04:26:46 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-31 04:54:44 (Asia/Bangkok / GMT+07:00)
- **MD5**: `070ac19c30684f6eb4cb10a893567f13`
- **SHA256**: `b7d7af4d6f3bc2a32750cfb3080750837b327cf0a68fc92e33674bfe5ee2ccbc`
- **Encoding**: ASCII

**File code content:**

```go
package common

type Paging struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (p *Paging) Preset() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 2 {
		p.Limit = 2
	}
	if p.Limit > 50 {
		p.Limit = 50
	}
}

```

---

### <a id="📄-handler-usershandler-go"></a>📄 `handler/usershandler.go`

**File Info:**
- **Size**: 4.5 KB
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `handler/usershandler.go`
- **Relative Path**: `handler`
- **Created**: 2026-01-31 05:08:52 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-31 05:13:35 (Asia/Bangkok / GMT+07:00)
- **MD5**: `f9e0e933671f5127fbb3f250e0247225`
- **SHA256**: `ebd10b2e99eeb6a2f326436416e6ca1b74e3978d289fc9c7235b0b661048f979`
- **Encoding**: ASCII

**File code content:**

```go
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

```

---

### <a id="📄-models-user-go"></a>📄 `models/user.go`

**File Info:**
- **Size**: 1.67 KB
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `models/user.go`
- **Relative Path**: `models`
- **Created**: 2026-01-30 04:46:53 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-31 03:58:36 (Asia/Bangkok / GMT+07:00)
- **MD5**: `37134baaabef05f5553fcb0f29f0f1e1`
- **SHA256**: `361581d9fb9d6290e703896ea7fa9d7d41f2bf1c431e62a8521753b1cf8d5ca1`
- **Encoding**: UTF-8

**File code content:**

```go
package models

import (
	"errors"
	"strings"
	"time"
)

var (
	Err_UsernameCannotBeEmpty = errors.New("Username cannot be empty")
	Err_PasswordRange         = errors.New("Password length must be great than 0 and less than 5 characters")
)

type User struct {
	Id       int    `json:"id,omitempty" gorm:"column:id"`
	UserName string `json:"userName" gorm:"column:user_name"`
	Email    string `json:"email" gorm:"column:email"`
	//FirstName string `json:"firstName" gorm: "column:firstName"`
	//LastName  string `json:"lastName" gorm: "column:lastName"`
	EncryptedPassword string    `json:"encryptedPassword" gorm:"column:encrypted_password"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy         string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy         string    `json:"updated_by" gorm:"column:updated_by"`
	DeleteFlag        bool      `json:"deleted_flag" gorm:"column:deleted_flag"`
}

// vì ng dùng ko dùng hết cái user đó nên có 1 struct DTO để lấy 1 số trường user nhập
// Model này dành cho user nhập liệu (DTO)
type UserCreation struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserUpdateParams struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func (u *UserCreation) Validate() error {
	u.UserName = strings.TrimSpace(u.UserName)
	if u.UserName == "" {
		return Err_UsernameCannotBeEmpty
	}

	u.Password = strings.TrimSpace(u.Password)

	if len(u.Password) <= 0 || len(u.Password) > 5 {
		return Err_PasswordRange
	}

	return nil
}

```

---

### <a id="📄-main-go"></a>📄 `main.go`

**File Info:**
- **Size**: 1.76 KB
- **Extension**: `.go`
- **Language**: `go`
- **Location**: `main.go`
- **Relative Path**: `root`
- **Created**: 2026-01-30 02:26:48 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-31 05:11:46 (Asia/Bangkok / GMT+07:00)
- **MD5**: `3b422484cf7e4f5f974e88f957055b48`
- **SHA256**: `89730ca42fa793d0fd53b1c8e407d806f1f8949e141ea690dc93fa658d88b087`
- **Encoding**: ASCII

**File code content:**

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	user_handler "go.learning.com/go2025/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//connect database
	dsn := "host=localhost user=root password=123456 dbname=daithuvien port=5450 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Connectdb successful", db)

	//create gin http server
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	v1Route := r.Group("/api/v1")
	{
		// Define a simple GET endpoint
		v1Route.GET("/healthcheck", func(c *gin.Context) {
			// Return JSON response
			c.JSON(http.StatusOK, gin.H{
				"message": "Server is running now",
			})
		})

		//User
		usersRoute := v1Route.Group("/users")
		{
			usersRoute.GET("/", user_handler.GetAllUsersHandler(db))      //Get All users
			usersRoute.GET("/:id", user_handler.GetUserByIdHandler(db))   //Get User by id
			usersRoute.PATCH("/:id", user_handler.UpdateUserHandler(db))  //Update user by id
			usersRoute.DELETE("/:id", user_handler.DeleteUserHandler(db)) //Delete user by id
			usersRoute.POST("", user_handler.CreatedUserHandler(db))      //Create new user
		}

		storiesRoute := v1Route.Group("/stories")
		{
			storiesRoute.GET("/")       //Get All stories
			storiesRoute.GET("/:id")    //Get stories by id
			storiesRoute.PATCH("/:id")  //Update stories by id
			storiesRoute.DELETE("/:id") //Delete stories by id
			storiesRoute.POST("")       //Create new stories
		}

	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

```

---

### <a id="📄-readme-md"></a>📄 `README.md`

**File Info:**
- **Size**: 2.01 KB
- **Extension**: `.md`
- **Language**: `text`
- **Location**: `README.md`
- **Relative Path**: `root`
- **Created**: 2026-01-30 02:27:41 (Asia/Bangkok / GMT+07:00)
- **Modified**: 2026-01-30 08:24:01 (Asia/Bangkok / GMT+07:00)
- **MD5**: `1469ba7e6bdeeb57070d1eca6d2909fe`
- **SHA256**: `968082cdf481855e63bbe1117d803ee43326d32fa019e6a7dfaa0b87273f91a0`
- **Encoding**: UTF-8

**File code content:**

````markdown
docker run --name postgre16 -p 5450:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=daithuvien -d postgres:16-alpine


docker run: Tạo mới & khởi động 
--name postgre16: Đặt tên cho container này là postgre16

-p 5432:5432: Port mapping -p [Cổng máy thật(host)]:[Cổng trong container] => Nối cổng 5432 trên lap của mình với 5432 của cái container đó => Giúp DBeaver, code Go/python có thể kết nối vô db thông qua địa chỉ localhost:5432

-e POSTGRES_USER=root: Biến môi trường để tạo tài khoản user mặc định 

-e POSTGRES_PASSWORD=password: Mật khẩu cho user trên 

-d (Detached mode): Chạy ngầm, sau khi chạy lên xong thì terminal sẽ trả lại quyền điều khiển cho mình, chứ nếu ko có cờ này thì terminal sẽ bị kẹt ở màn log của DB, tắt terminal là tắt DB

postgres:16-alpine (Image): Là bản thiết kế để tạo container, gồm tên image -> postgres, phiên bản postgreSQL: 16, alpine là bản siêu nhẹ 

----------------------------------------
Cấu hình để kết nối vô db
Host: localhost
Port: 5432
User: root
Password: password
Database name: root

----------------------------------------

docker ps: Xem danh sách các container đang hoạt động

Truy cập vào chế độ tương tác với db: docker exec -it postgre16 /bin/bash

Truy cập psql: psql

root=# show dbs
root=# \l

root=# create database testdb
root=# drop database testdb;
root=# exit


Lệnh tạo db từ ngoài: docker exec -it postgre16 createdb --username=root --owner=root testdb

Lệnh xóa db từ ngoài: docker exec -it postgre16 dropdb testdb

---------------------------------


create table users
(
    id                serial primary key,
    user_name          character(50),
    email           character(50),
    encrypted_password character(50),
    create_at timestamptz,
    updated_at timestamptz,
    create_by character(50),
    updated_by character(50),
    delete_flag bool 
);

````

---

## 🚫 Binary/Excluded Files

The following files were not included in the text content:

- `gitignore`
- `go.mod`
- `go.sum`
- `Makefile`

