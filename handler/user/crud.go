package user

import (
	"fmt"

	"github.com/fahadkhan-fk/go-gin-backend/pkg/postgres"
	"github.com/fahadkhan-fk/go-gin-backend/pkg/schema"
	"github.com/gin-gonic/gin"
)

//CreateUser func (create a new record in the database).
func Create(c *gin.Context) {
	var (
		user   *schema.User
		err    error
		client = postgres.GetClient()
	)
	c.BindJSON(&user)

	user, err = schema.AddNewUser(client, user)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, user)
	}
}

//GetUsers func (get all the records from the database).
func GetAll(c *gin.Context) {
	client := postgres.GetClient()
	users, err := schema.FetchAllUsers(client, &[]schema.User{})
	if err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, users)
	}
}

//GetByID func (get one specific record from the database against the id provided).
func GetByID(c *gin.Context) {
	client := postgres.GetClient()
	userID := c.Params.ByName("id")
	user, err := schema.GetUserByID(client, &schema.User{}, userID)
	if err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}

//UpdateUser func (updates the record in the database against the id provided).
func Update(c *gin.Context) {
	client := postgres.GetClient()
	userID := c.Params.ByName("id")
	user, err := schema.GetUserByID(client, &schema.User{}, userID)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}

	//binding the data in JSON and saving in DB.
	c.BindJSON(&user)
	schema.UpdateUser(client, user)
	c.JSON(200, user)
}

//DeleteUser func (delete the record in the database against the id provided).
func Delete(c *gin.Context) {
	client := postgres.GetClient()
	userID := c.Params.ByName("id")
	user, err := schema.GetUserByID(client, &schema.User{}, userID)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}

	schema.DeleteUser(client, user)
	c.JSON(200, fmt.Sprintf("User %s is deleted", userID))
}
