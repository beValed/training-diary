package handlers

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"training_diary/configs/entities"
	"training_diary/internal/buttons"
	"training_diary/internal/messages"
)

var (
	countForCallback int		// crutch for callback
	parametersStr    []string	// collects str parameters
	parametersInt    []int		// collects int parameters
)

// starts bot

func StartBot(app *entities.Application) func(m *tbot.Message) {
	return func(m *tbot.Message) {
		_, err := app.Client.SendMessage(m.Chat.ID, messages.MessageForStart)
		if err != nil {
			return
		}
	}
}

// gives options in '/inform'

func GetOptions(app *entities.Application) func(m *tbot.Message) {
	return func(m *tbot.Message) {
		btnHeight := buttons.MakeButtonsForHeight()
		btnWeight := buttons.MakeButtonsForWeight()
		btnGender := buttons.MakeButtonsForGender()
		_, errGender := app.Client.SendMessage(m.Chat.ID, "Pick your gender:", tbot.OptInlineKeyboardMarkup(btnGender))
		if errGender != nil {
			return
		}
		_, errHeight := app.Client.SendMessage(m.Chat.ID, "Pick your height (in cm):", tbot.OptInlineKeyboardMarkup(btnHeight))
		if errHeight != nil {
			return
		}
		_, errWeight := app.Client.SendMessage(m.Chat.ID, "Pick your weight (in kg):", tbot.OptInlineKeyboardMarkup(btnWeight))
		if errWeight != nil {
			return
		}
	}
}

// saves data in slices and sends callback

func CallbackHandler(app *entities.Application) func(cq *tbot.CallbackQuery) {
	return func(cq *tbot.CallbackQuery) {
		var humanChoice int
		humanChoice = convertStrToInt(cq.Data)
		errDelete := app.Client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
		if errDelete != nil {
			return
		} else {
			countForCallback++
		}

		switch countForCallback {
		case 1:
			if humanChoice != 0 {
				parametersInt = append(parametersInt, humanChoice)
			}
		case 2:
			if humanChoice != 0 {
				parametersInt = append(parametersInt, humanChoice)
			}
		default:
			if humanChoice != 0 {
				parametersInt = append(parametersInt, humanChoice)
			}
		}

		if countForCallback%3 == 0 {
			_, errSend := app.Client.SendMessage(cq.Message.Chat.ID, messages.MessageForCallback)
			if errSend != nil {
				return
			}
		}
		fmt.Println("String parameters - ", parametersStr)
		fmt.Println("Numeric parameters - ", parametersInt)
	}
}

// converts str to int and saves str parameters in str slice

func convertStrToInt(someString string) int {
	var humanChoice int
	if someString == "< 150" || someString == "150 - 170" || someString == "170 - 180" ||
		someString == "180 - 190" || someString == "190 +" || someString == "< 50" || someString == "50 - 65" ||
		someString == "65 - 80" || someString == "80 - 90" || someString == "90 - 100" || someString == "100 +" {

		switch someString {
		case "< 50":
			humanChoice = 45
		case "50 - 65":
			humanChoice = 60
		case "65 - 80":
			humanChoice = 75
		case "80 - 90":
			humanChoice = 85
		case "90 - 100":
			humanChoice = 95
		case "100 +":
			humanChoice = 110
		case "< 150":
			humanChoice = 145
		case "150 - 170":
			humanChoice = 160
		case "170 - 180":
			humanChoice = 175
		case "180 - 190":
			humanChoice = 185
		case "190 +":
			humanChoice = 200
		}

	} else {
		parametersStr = append(parametersStr, someString)
	}
	return humanChoice
}

func GetTraining(app *entities.Application) func(m *tbot.Message) {
	return func(m *tbot.Message) {

	}

}

func GetNutrition(app *entities.Application) func(m *tbot.Message) {
	return func(m *tbot.Message) {

	}

}

func GetTips(app *entities.Application) func(m *tbot.Message) {
	return func(m *tbot.Message) {

	}

}

func MakeDonation(app *entities.Application) func(m *tbot.Message) {
	return func(m *tbot.Message) {

	}

}
