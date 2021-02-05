package controllers

import (
	models "crud_gin/crudproject/connection"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

const database, collection = "go-development", "people"

type user struct {
	Name string `json:"name",bson:"name"`
	Age  int    `json:"age",bson:"age"`
	City string `json:"city",bson:"city"`
}

func Index(c *gin.Context) {

	var users primitive.A
	db := models.SetupDB()
	defer db.Close()
	err := db.DB(database).C(collection).Find(bson.M{}).All(&users)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user:", users)

	c.HTML(http.StatusOK, "Index", users)

}
func New(c *gin.Context) {

	c.HTML(200, "New", gin.H{})
}

func Insert(c *gin.Context) {
	db := models.SetupDB()
	var person user

	person.Name = c.Request.PostFormValue("name")
	person.City = c.Request.PostFormValue("city")
	ae := c.Request.PostFormValue("age")
	person.Age, _ = strconv.Atoi(ae)

	fmt.Println(person)
	err := db.DB(database).C(collection).Insert(person)
	if err != nil {
		log.Fatal(err)
	}

	c.Redirect(301, "/")

}

// Get Profile of a particular User by Name

func Show(c *gin.Context) {
	db := models.SetupDB()
	Name := c.Query("id")

	fmt.Println("name id is", Name)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	db.DB(database).C(collection).Find(bson.M{"name": Name}).One(&result)

	c.HTML(200, "Show", result)
}

//Update Profile of User
func Edit(c *gin.Context) {

	db := models.SetupDB()
	Name := c.Request.URL.Query().Get("id")
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	db.DB(database).C(collection).Find(bson.M{"name": Name}).One(&result)
	c.HTML(200, "Edit", result)

}

func UpdateProfile(c *gin.Context) {

	db := models.SetupDB()
	var user user

	uid := c.Request.PostFormValue("uid")
	user.Name = c.Request.PostFormValue("name")
	user.City = c.Request.PostFormValue("city")
	ae := c.Request.PostFormValue("age")
	user.Age, _ = strconv.Atoi(ae)

	update := bson.D{{"$set", bson.D{{"city", user.City}, {"age", user.Age}, {"name", user.Name}}}}

	err := db.DB(database).C(collection).Update(bson.M{"name": uid}, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("update is", user)
	c.Redirect(301, "/")
}

// Delete Profile of User

func DeleteProfile(c *gin.Context) {
	db := models.SetupDB()
	defer db.Close()
	Name := c.Request.URL.Query().Get("id")
	err := db.DB(database).C(collection).Remove(bson.M{"name": Name})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deleted successfully")
	c.Redirect(301, "/")
}
