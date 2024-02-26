package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Kavindulakmal/ecommerce-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string) string {

}
func VerifyPassword(userpassword string, givenpassword string) (bool, string) {

}
func SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		}
	}
}
func Login() gin.HandlerFunc {

}
func ProductViewerAdmin() gin.HandlerFunc {

}
func SearchProduct() gin.HandlerFunc {

}
func SearchProductByQuery() gin.HandlerFunc {

}
