package controller

import (
	"fmt"

	"github.com/Strike-official/global-getting-started/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/strike-official/go-sdk/strike"
)

func Getting_started(ctx *gin.Context) {
	var request model.Strike_Meta_Request_Structure
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
	answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{"https://raw.githubusercontent.com/shashank404error/global-getting-started/main/media/namaste.png"}, []string{}).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Congractulations! you just creared your new bot", "#41a800", true)

	ctx.JSON(200, strikeObject)
}

// func SettingHelper(name, id string) *strike.Response_structure {
// 	strikeObject := strike.Create("setting", "http://8512-2405-201-a407-9cb6-2016-3fe-e295-bed2.ngrok.io/ytbot/setting/set")

// 	question_object := strikeObject.Question("preference").
// 		QuestionCard().SetHeaderToQuestion(1, strike.FULL_WIDTH).AddTextRowToQuestion(strike.H4, "Hi "+name+", please select your favorite topics", "#0088e3", false)

// 	// Add answer
// 	answer_object := question_object.Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION)

// 	topicArr := database.GetAllTags()

// 	//get isSelected Array
// 	//
// 	user_tag_name, _ := database.GetTagsForAUser(id)
// 	isSelectedArr := database.GetBooleanArray(topicArr, user_tag_name)

// 	for i := 0; i < len(topicArr); i++ {
// 		if isSelectedArr[i] {
// 			answer_object = answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).
// 				AddTextRowToAnswer(strike.H4, topicArr[i], "#41a800", true)
// 		} else {
// 			answer_object = answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).
// 				AddTextRowToAnswer(strike.H4, topicArr[i], "#0088e3", true)
// 		}
// 	}
// 	return strikeObject
// }

// func Set_Setting(ctx *gin.Context) {
// 	var request model.Strike_Meta_Request_Structure
// 	if err := ctx.BindJSON(&request); err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	preference_array := request.User_session_variables.Preference
// 	fmt.Println(preference_array)

// 	id := database.GetUserRDS(request)
// 	//Adjust preference according to the user
// 	added_tags := database.UpdateUserPreference(preference_array, id)
// 	var added_tags_string string
// 	for _, t := range added_tags {
// 		added_tags_string = added_tags_string + " " + t + ","
// 	}
// 	//show in app

// 	strike_object := strike.Create("setting", "http://8512-2405-201-a407-9cb6-2016-3fe-e295-bed2.ngrok.io/ytbot/setting")
// 	question_object := strike_object.Question("").
// 		QuestionCard().SetHeaderToQuestion(1, strike.FULL_WIDTH)

// 	if len(added_tags) > 0 {
// 		question_object = question_object.AddTextRowToQuestion(strike.H4, added_tags_string+" added to your preference", "#41a800", false)
// 	} else {
// 		question_object = question_object.AddTextRowToQuestion(strike.H4, "No new tags added to your preference", "#0088e3", false)
// 	}

// 	question_object.Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION)

// 	ctx.JSON(200, strike_object)

// }

// func Youtubefeed(ctx *gin.Context) {
// 	var request model.Strike_Meta_Request_Structure
// 	if err := ctx.BindJSON(&request); err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	name := request.Bybrisk_session_variables.Username

// 	//get isSelected Array
// 	id := database.GetUserRDS(request)
// 	selectedTagNames, selectedTags := database.GetTagsForAUser(id)
// 	fmt.Println("selectedTags -> ", selectedTagNames)

// 	// ----------------------------
// 	// Create strike object
// 	// ----------------------------

// 	if len(selectedTags) == 0 {
// 		fmt.Println("Going to Setting handler")
// 		strikeObj := SettingHelper(name, id)
// 		ctx.JSON(200, strikeObj)
// 		return
// 	}

// 	// ----------------------------
// 	// Get YT Videos
// 	// ----------------------------
// 	ytvideos := getYTVideos2(id, selectedTags, selectedTagNames)

// 	strikeObj := strike.Create("setting", "http://8512-2405-201-a407-9cb6-2016-3fe-e295-bed2.ngrok.io/ytbot/feed")
// 	ques1 := strikeObj.Question("Hi "+name+", here you go!").
// 		QuestionCard().
// 		SetHeaderToQuestion(1, strike.FULL_WIDTH).
// 		AddTextRowToQuestion(strike.H4, "Here you go!", "#0088e3", false)

// 	ans1 := ques1.Answer(false).
// 		AnswerCardArray(strike.VERTICAL_ORIENTATION)

// 	for _, v := range ytvideos {
// 		id := strings.Split(v.Url, "?v=")
// 		thumb := "https://img.youtube.com/vi/" + id[1] + "/0.jpg"
// 		fmt.Println(v, thumb)
// 		ans1 = ans1.AnswerCard().
// 			SetHeaderToAnswer(1, strike.FULL_WIDTH).
// 			AddGraphicRowToAnswer("video_row", []string{v.Url}, []string{thumb}).
// 			AddTextRowToAnswer("h4", v.Title, "#320D34", true).
// 			AddTextRowToAnswer("h5", v.Tag, "#363636", false)
// 	}

// 	ans1.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).
// 		AddTextRowToAnswer(strike.H3, "Show More 👇", "#1064eb", true)

// 	ctx.JSON(200, strikeObj)

// }

// func getYTVideos(id string, selectedTags []string) []model.VideoInfo {
// 	// ytvideos := database.GetVideosForUser(id, selectedTags)
// 	// fmt.Println(ytvideos)

// 	var ytvideos []model.VideoInfo
// 	maxLength := len(model.VideoDB) - 1
// 	rand.Seed(time.Now().UnixNano())
// 	maxVideo := 10
// 	for {
// 		r := rand.Intn(maxLength)
// 		if stringInSlice(model.VideoDB[r].Tag, selectedTags) {
// 			ytvideos = append(ytvideos, model.VideoDB[r])
// 			maxVideo--
// 		}
// 		if maxVideo == 0 {
// 			break
// 		}
// 	}
// 	return ytvideos
// }

// func getYTVideos2(id string, selectedTags, selectedTagsNames []string) []model.VideoInfo {
// 	// ytvideos := database.GetVideosForUser(id, selectedTags)
// 	// fmt.Println(ytvideos)

// 	// var ytvideos []model.VideoInfo
// 	// maxLength := len(model.VideoDB) - 1
// 	// rand.Seed(time.Now().UnixNano())
// 	videoToBeShown := 10
// 	interestedTopicsCount := len(selectedTags)
// 	if interestedTopicsCount > videoToBeShown {
// 		selectedTags = selectedTags[:10]
// 	}

// 	ytvideos := getYTVideosHelper2(selectedTags, selectedTagsNames, videoToBeShown)
// 	return ytvideos
// }

// func getYTVideosHelper2(selectedTags, selectedTagsNames []string, videosTobeShownCount int) []model.VideoInfo {
// 	var listOfVideos []model.VideoInfo
// 	for i, tag := range selectedTags {
// 		noOfChannels := math.Round(math.Ceil(float64(videosTobeShownCount) / float64(len(selectedTags))))
// 		randomChanForTag := fetchRandomChan(tag, int(noOfChannels))
// 		fmt.Println("randomChannelsFor Tag ->", randomChanForTag)
// 		for _, ytchan := range randomChanForTag {
// 			noOfVideo := 1
// 			randomVideoForChan := fetchRandomVideo(ytchan, selectedTagsNames[i], noOfVideo)
// 			listOfVideos = append(listOfVideos, randomVideoForChan...)
// 		}
// 	}
// 	return listOfVideos
// }

// func fetchRandomChan(tag string, noOfChannels int) []string {
// 	fmt.Println("Tag is ", tag)
// 	t := model.Tag2ChanMap[tag]
// 	fmt.Println("TAG COUNT", len(t))
// 	var c []string
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < noOfChannels; i++ {
// 		idx := rand.Intn(len(t))
// 		c = append(c, t[idx])
// 		removeFromSlice(t, idx)
// 	}
// 	fmt.Println("Channel is ", c)
// 	return c
// }

// func fetchRandomVideo(ytchan, tag string, noOfVideo int) []model.VideoInfo {
// 	fmt.Println("ytchan", ytchan)
// 	t := model.Chan2Video[ytchan]
// 	var v []model.VideoInfo
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < noOfVideo; i++ {
// 		fmt.Println("fetchRandomVideo", ytchan, len(t))
// 		idx := rand.Intn(len(t))
// 		vtmp := t[idx]
// 		vtmp.Tag = tag
// 		v = append(v, vtmp)
// 		// removeFromSliceVideo(v, idx)
// 	}
// 	return v
// }

// func stringInSlice(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }

// func removeFromSlice(s []string, i int) []string {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

// func removeFromSliceVideo(s []model.VideoInfo, i int) []model.VideoInfo {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }
