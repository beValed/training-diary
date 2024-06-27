package buttons

import "github.com/yanzay/tbot/v2"

// makes buttons for height

func MakeButtonsForHeight() *tbot.InlineKeyboardMarkup {
	btnLowHeight := tbot.InlineKeyboardButton{
		Text:         "< 150",
		CallbackData: "< 150",
	}
	btnUpperLowHeight := tbot.InlineKeyboardButton{
		Text:         "150 - 170",
		CallbackData: "150 - 170",
	}
	btnMiddleHeight := tbot.InlineKeyboardButton{
		Text:         "170 - 180",
		CallbackData: "170 - 180",
	}
	btnUpperMiddleHeight := tbot.InlineKeyboardButton{
		Text:         "180 - 190",
		CallbackData: "180 - 190",
	}
	btnHighHeight := tbot.InlineKeyboardButton{
		Text:         "190 +",
		CallbackData: "190 +",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnLowHeight, btnUpperLowHeight, btnMiddleHeight,
				btnUpperMiddleHeight, btnHighHeight},
		},
	}
}

// makes buttons for weight

func MakeButtonsForWeight() *tbot.InlineKeyboardMarkup {
	btnLowWeight := tbot.InlineKeyboardButton{
		Text:         "< 50",
		CallbackData: "< 50",
	}
	btnUpperLowWeight := tbot.InlineKeyboardButton{
		Text:         "50 - 65",
		CallbackData: "50 - 65",
	}
	btnMiddleWeight := tbot.InlineKeyboardButton{
		Text:         "65 - 80",
		CallbackData: "65 - 80",
	}
	btnUpperMiddleWeight := tbot.InlineKeyboardButton{
		Text:         "80 - 90",
		CallbackData: "80 - 90",
	}
	btnHighWeight := tbot.InlineKeyboardButton{
		Text:         "90 - 100",
		CallbackData: "90 - 100",
	}
	btnUpperHighWeight := tbot.InlineKeyboardButton{
		Text:         "100 +",
		CallbackData: "100 +",
	}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnLowWeight, btnUpperLowWeight, btnMiddleWeight,
				btnUpperMiddleWeight, btnHighWeight, btnUpperHighWeight},
		},
	}
}

// makes buttons for gender

func MakeButtonsForGender() *tbot.InlineKeyboardMarkup {
	btnMale := tbot.InlineKeyboardButton{
		Text:         "Male",
		CallbackData: "male",
	}
	btnFemale := tbot.InlineKeyboardButton{
		Text:         "Female",
		CallbackData: "female",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnMale, btnFemale},
		},
	}
}
