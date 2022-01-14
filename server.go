package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/strike-official/go-sdk/strike"
)

type Strike_Meta_Request_Structure struct {

	// Bybrisk variable from strike bot
	//
	Bybrisk_session_variables Bybrisk_session_variables_struct `json: "bybrisk_session_variables"`

	// Our own variable from previous API
	//
	User_session_variables User_session_variables_struct `json: "user_session_variables"`
}

type Bybrisk_session_variables_struct struct {

	// User ID on Bybrisk
	//
	UserId string `json:"userId"`

	// Our own business Id in Bybrisk
	//
	BusinessId string `json:"businessId"`

	// Handler Name for the API chain
	//
	Handler string `json:"handler"`

	// Current location of the user
	//
	Location GeoLocation_struct `json:"location"`

	// Username of the user
	//
	Username string `json:"username"`

	// Address of the user
	//
	Address string `json:"address"`

	// Phone number of the user
	//
	Phone string `json:"phone"`
}

type GeoLocation_struct struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type User_session_variables_struct struct {
	TextInput     string             `json:"textInput"`
	LocationInput GeoLocation_struct `json:"locationInput"`
	NumberInput   string             `json:"numberInput"`
	DateInput     []string           `json:"dateInput"`
	Card          []string           `json:"card"`
}

type AppConfig struct {
	Port  string `json:"port"`
	APIEp string `json:"apiep"`
}

var conf *AppConfig

func main() {
	conf = &AppConfig{Port: ":7001", APIEp: "http://ec2-18-218-96-97.us-east-2.compute.amazonaws.com"}
	// Init Routes
	router := gin.Default()
	bot := router.Group("/global-getting-started/")
	{
		bot.POST("/getting-started", Getting_started)
	}

	// Start serving the application
	err := router.Run(conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}
}

func Getting_started(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(request)
	name := request.Bybrisk_session_variables.Username
	strikeObject := strike.Create("getting_started", "")

	question_object := strikeObject.Question("").
		QuestionText().
		SetTextToQuestion("Hi "+name+"! Welcome to strike developers community", "Text Description, getting used for testing purpose.")

	answer_object := question_object.Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION)
	answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{"https://raw.githubusercontent.com/Strike-official/global-getting-started/main/media/Pebble%20People%20-%20Working%20From%20Home.png"}, []string{}).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Congractulations! you just creared your new bot", "#41a800", true)

	ctx.JSON(200, strikeObject)
}
