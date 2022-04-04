package user

import (
	"github.com/fahadkhan-fk/go-gin-backend/pkg/postgres"
	"github.com/fahadkhan-fk/go-gin-backend/pkg/schema"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreateUser func (create a new record in the database).
func Create(c *gin.Context) {
	var (
		user   *schema.User
		err    *gorm.DB
		client = postgres.GetClient()
	)
	c.BindJSON(&user)

	user, err = schema.AddNewUser(client, user)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, user)
	}
}

//GetUsers func (get all the records from the database).
func GetAll(c *gin.Context) {
	var (
		users  *[]schema.User
		err    *gorm.DB
		client = postgres.GetClient()
	)

	users, err = schema.FetchAllUsers(client, users)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, users)
	}
}

//GetByID func (get one specific record from the database against the id provided).
func GetByID(c *gin.Context) {
	var (
		user   *schema.User
		err    *gorm.DB
		client = postgres.GetClient()
	)

	userID := c.Params.ByName("id")
	user, err = schema.GetUserByID(client, user, userID)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

//UpdateUser func (updates the record in the database against the id provided).
func Update(c *gin.Context) {
	var (
		user   *schema.User
		err    *gorm.DB
		client = postgres.GetClient()
	)

	userID := c.Params.ByName("id")
	user, err = schema.GetUserByID(client, user, userID)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	//binding the data in JSON and saving in DB.
	c.BindJSON(&user)
	schema.UpdateUser(client, user)

	c.JSON(200, user)
}

//DeleteUser func (delete the record in the database against the id provided).
func Delete(c *gin.Context) {
	var (
		user   *schema.User
		err    *gorm.DB
		client = postgres.GetClient()
	)

	userID := c.Params.ByName("id")
	user, err = schema.GetUserByID(client, user, userID)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	schema.DeleteUser(client, user)
	c.JSON(200, "User is deleted")
}
