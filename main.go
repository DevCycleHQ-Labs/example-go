package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	devcycle "github.com/devcyclehq/go-server-sdk/v2"
)

func main() {
	godotenv.Load(".env")

	// Initialize a single instance of the DevCycle client
	initalizeDevCycle()

	// Log the current DevCycle variation to the console.
	go logVariation()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(devcycleMiddleware())
	router.Use(userMiddleware())
	router.GET("/", getGreeting)
	router.GET("/variables", getVariables)

	router.Run("localhost:8000")
}

// Add the DevCycle client to the request context
func devcycleMiddleware() gin.HandlerFunc {
	devcycleClient := getDevCycleClient()

	return func(c *gin.Context) {
		c.Set("devcycle", devcycleClient)
		c.Next()
	}
}

// Create a user object and add it to the request context
func userMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := devcycle.User{
			UserId: "example_user_id",
		}

		c.Set("user", user)
		c.Next()
	}
} 

// Fetch a greeting message baed on the "example-text" variable
func getGreeting(c *gin.Context) {
	user := c.Value("user").(devcycle.User)
	devcycleClient := c.Value("devcycle").(*devcycle.Client)

	variableValue, err := devcycleClient.VariableValue(user, "example-text", "default")

	if err != nil {
		log.Fatalf("Error getting variable value: %v", err)
	}

	var header string
	var body string
	switch step := variableValue.(string); step {
		case "step-1":
			header = "Welcome to DevCycle's example app."
			body = "If you got here through the onboarding flow, just follow the instructions to change and create new Variations and see how the app reacts to new Variable values."
		case "step-2":
			header = "Great! You've taken the first step in exploring DevCycle."
			body = "You've successfully toggled your very first Variation. You are now serving a different value to your users and you can see how the example app has reacted to this change. Next, go ahead and create a whole new Variation to see what else is possible in this app."
		case "step-3":
			header = "You're getting the hang of things."
			body = "By creating a new Variation with new Variable values and toggling it on for all users, you've already explored the fundamental concepts within DevCycle. There's still so much more to the platform, so go ahead and complete the onboarding flow and play around with the feature that controls this example in your dashboard."
		default:
			header = "Welcome to DevCycle's example app."
			body = "If you got to the example app on your own, follow our README guide to create the Feature and Variables you need to control this app in DevCycle."
	}

	content := []byte(fmt.Sprintf("<h2>%s</h2><p>%s</p><p><a href=\"/variables\">All Variables</a></p>", header, body))
    c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

// Fetch all DevCycle variables and return them as JSON
// This should just be used for debugging purposes
func getVariables(c *gin.Context) {
	user := c.Value("user").(devcycle.User)
	devcycleClient := c.Value("devcycle").(*devcycle.Client)

	variables, err := devcycleClient.AllVariables(user)

	if err != nil {
		log.Fatalf("Error getting variables: %v", err)
	}

	c.JSON(http.StatusOK, variables)
}
