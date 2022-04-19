package main

import (
	"sort"
	"strings"
)

const (
	Black     = "черный"
	Brown     = "коричневый"
	Darkblue  = "синий"
	Blue      = "голубой"
	Purple    = "фиолетовый"
	Red       = "красный"
	LightPink = "светло-розовый"
	DarkPink  = "темно-розовый"
	LightPinkWithoutDash = "светло розовый"
	DarkPinkWithoutDash  = "темно розовый"
	Orange    = "оранжевый"
	Yellow    = "желтый"
	Green     = "зеленый"
	Beige     = "бежевый"
	LightGray = "светло-серый"
	DarkGray  = "темно-серый"
	LightGrayWithoutDash = "светло серый"
	DarkGrayWithoutDash  = "темно серый"
	White     = "белый"
)

const (
	hintIDEasyLevel1 = 457239065
	pictureIDEasyLevel1 = 457239066
	hintIDEasyLevel2 = 457239148
	pictureIDEasyLevel2 = 457239149
	hintIDEasyLevel3 = 457239069
	pictureIDEasyLevel3 = 457239070
	hintIDEasyLevel4 = 457239071
	pictureIDEasyLevel4 = 457239072
	hintIDEasyLevel5 = 457239073
	pictureIDEasyLevel5 = 457239074
	hintIDEasyLevel6 = 457239116
	pictureIDEasyLevel6 = 457239117
	hintIDEasyLevel7 = 457239118
	pictureIDEasyLevel7 = 457239119
	hintIDEasyLevel8 = 457239120
	pictureIDEasyLevel8 = 457239121
	hintIDEasyLevel9 = 457239122
	pictureIDEasyLevel9 = 457239123
	hintIDEasyLevel10 = 457239124
	pictureIDEasyLevel10 = 457239125
)

const (
	hintIDMediumLevel1 = 457239095
	pictureIDMediumLevel1 = 457239096
	hintIDMediumLevel2 = 457239146
	pictureIDMediumLevel2 = 457239147
	hintIDMediumLevel3 = 457239099
	pictureIDMediumLevel3 = 457239100
	hintIDMediumLevel4 = 457239101
	pictureIDMediumLevel4 = 457239102
	hintIDMediumLevel5 = 457239103
	pictureIDMediumLevel5 = 457239104
	hintIDMediumLevel6 = 457239126
	pictureIDMediumLevel6 = 457239127
	hintIDMediumLevel7 = 457239128
	pictureIDMediumLevel7 = 457239129
	hintIDMediumLevel8 = 457239130
	pictureIDMediumLevel8 = 457239131
	hintIDMediumLevel9 = 457239132
	pictureIDMediumLevel9 = 457239133
	hintIDMediumLevel10 = 457239134
	pictureIDMediumLevel10 = 457239135
)

const (
	hintIDHardLevel1 = 457239105
	pictureIDHardLevel1 = 457239106
	hintIDHardLevel2 = 457239107
	pictureIDHardLevel2 = 457239108
	hintIDHardLevel3 = 457239109
	pictureIDHardLevel3 = 457239110
	hintIDHardLevel4 = 457239111
	pictureIDHardLevel4 = 457239112
	hintIDHardLevel5 = 457239113
	pictureIDHardLevel5 = 457239114
	hintIDHardLevel6 = 457239136
	pictureIDHardLevel6 = 457239137
	hintIDHardLevel7 = 457239138
	pictureIDHardLevel7 = 457239139
	hintIDHardLevel8 = 457239140
	pictureIDHardLevel8 = 457239141
	hintIDHardLevel9 = 457239142
	pictureIDHardLevel9 = 457239143
	hintIDHardLevel10 = 457239144
	pictureIDHardLevel10 = 457239145
)

var hintID, pictureID, numbOfColoringRules int
var breakTimeShort = "<break time=\"0.005s\"/>"
var breakTimeMiddle = "<break time=\"0.5s\"/>"
var breakTimeLong = "<break time=\"3s\"/>"
var colorsAndQuantity map[string]int
var endDrawingText, endDrawingSSML, resultText string
var pictureType string
var colorsSort []string
var pictureDraw bool

func easyLevel1(resp Response) Response {
	resp.Text = "Нарисуем яблоко. Для этого Вам понадобятся 4 цвета: " + Yellow + ", " + Red + ", " + Green + " и " + Brown + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем яблоко. Для этого Вам понадобятся четыре цв+ета\n"  + Yellow + "\n" + Red + "\n" + Green + " - и " + Brown + ". С какого цвета начнем?"

	resultText = "Поздравляю! Вы нарисовали яблоко."
	pictureType = "Образец для рисования яблока."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 2; colorsAndQuantity[Red] = 9; colorsAndQuantity[Green] = 3; colorsAndQuantity[Brown] = 1
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel1
	pictureID = pictureIDEasyLevel1
	pictureDraw = false
	return resp
}

func easyLevel2(resp Response) Response {
	resp.Text = "Давайте нарисуем мороженое. Для этого Вам необходимо 5 цветов: " + LightPink + ", " + Red + ", " + Blue + ", " + Darkblue + " и " + Brown + ". С какого цвета Вы хотите начать?"
	resp.TTS = "Давайте нарисуем мороженое. Для этого Вам необходимо пять цветов\n" + LightPink + "\n" + Red + "\n" + Blue + ", " + Darkblue + " - и " + Brown + ". С какого цвета Вы хотите начать?"

	resultText = "Поздравляю! Мороженое нарисовано."
	pictureType = "Образец для рисования мороженого."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[LightPink] = 3; colorsAndQuantity[Red] = 3; colorsAndQuantity[Blue] = 3; colorsAndQuantity[Darkblue] = 2; colorsAndQuantity[Brown] = 4
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel2
	pictureID = pictureIDEasyLevel2
	pictureDraw = false
	return resp
}

func easyLevel3(resp Response) Response {
	resp.Text = "Будем рисовать клубнику. Для этого Вам необходимо подготовить 3 цвета: " + Yellow + ", " + Red + " и " + Green + ". С какого цвета начнем рисовать?"
	resp.TTS = "Будем рисовать клубнику. Для этого Вам необходимо подготовить три цв+ета\n" + Yellow + "\n" + Red + " - и " + Green + ". С какого цвета начнем рисовать?"

	resultText = "Отлично! Вы нарисовали клубнику."
	pictureType = "Образец для рисования клубники."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 7; colorsAndQuantity[Red] = 9; colorsAndQuantity[Green] = 5
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel3
	pictureID = pictureIDEasyLevel3
	pictureDraw = false
	return resp
}

func easyLevel4(resp Response) Response {
	resp.Text = "Изобразим змею. Для этого Вам нужно 2 цвета: " + Green + " и " + Black + ". Какой цвет Вы хотите выбрать?"
	resp.TTS = "Изобразим змею. Для этого Вам нужно два цв+ета\n" + Green + " - и " + Black + ". Какой цвет Вы хотите выбрать?"

	resultText = "Поздравляю! Вы нарисовали змею."
	pictureType = "Образец для рисования змеи."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Green] = 13; colorsAndQuantity[Black] = 2
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel4
	pictureID = pictureIDEasyLevel4
	pictureDraw = false
	return resp
}

func easyLevel5(resp Response) Response {
	resp.Text = "Давайте изобразим осьминога. Для этого Вам пригодятся 3 цвета: " + Purple + ", " + Darkblue + " и " + Black + ". Какой цвет нарисуем первым?"
	resp.TTS = "Давайте изобразим осьминога. Для этого Вам пригодятся три цв+ета\n" + Purple + "\n" + Darkblue + " - и " + Black + ". Какой цвет нарисуем первым?"

	resultText = "Поздравляю! Осьминог нарисован."
	pictureType = "Образец для рисования осьминога."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Purple] = 7; colorsAndQuantity[Darkblue] = 10; colorsAndQuantity[Black] = 2
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel5
	pictureID = pictureIDEasyLevel5
	pictureDraw = false
	return resp
}

func easyLevel6(resp Response) Response {
	resp.Text = "Давайте нарисуем божью коровку. Для этого Вам пригодятся 2 цвета: " + Red + " и " + Black + ". Какой цвет нарисуем первым?"
	resp.TTS = "Давайте нарисуем божью коровку. Для этого Вам пригодятся два цв+ета\n" + Red + " - и " + Black + ". Какой цвет нарисуем первым?"

	resultText = "Поздравляю! Божья коровка нарисована."
	pictureType = "Образец для рисования божьей коровки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Red] = 9; colorsAndQuantity[Black] = 7
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel6
	pictureID = pictureIDEasyLevel6
	pictureDraw = false
	return resp
}

func easyLevel7(resp Response) Response {
	resp.Text = "Давайте нарисуем новогоднюю ёлку. Для этого Вам понадобятся 7 цветов: " + Yellow + ", " + LightPink + ", " + Red + ", " + Green + ", " + Darkblue + ", " + Brown + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Давайте нарисуем новогоднюю ёлку. Для этого Вам понадобятся семь цветов\n" + Yellow + "\n" + LightPink + "\n" + Red + "\n" + Green + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Отлично! Вы нарисовали новогоднюю ёлку."
	pictureType = "Образец для рисования новогодней ёлки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 5; colorsAndQuantity[LightPink] = 2; colorsAndQuantity[Red] = 1; colorsAndQuantity[Green] = 11; colorsAndQuantity[Darkblue] = 1; colorsAndQuantity[Brown] = 3; colorsAndQuantity[Black] = 4
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel7
	pictureID = pictureIDEasyLevel7
	pictureDraw = false
	return resp
}

func easyLevel8(resp Response) Response {
	resp.Text = "Давайте нарисуем краба. Для этого Вам понадобятся 3 цвета: " + Yellow + ", " + Orange + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Давайте нарисуем краба. Для этого Вам понадобятся три цв+ета\n" + Yellow + "\n" + Orange + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Отлично! Вы нарисовали краба."
	pictureType = "Образец для рисования краба."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 10; colorsAndQuantity[Orange] = 9; colorsAndQuantity[Black] = 2
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel8
	pictureID = pictureIDEasyLevel8
	pictureDraw = false
	return resp
}

func easyLevel9(resp Response) Response {
	resp.Text = "Будем рисовать персонажа «Молния Маккуин» из мульфильма «Тачки». Для этого Вам нужно 5 цветов: " + Yellow + ", " + Orange + ", " + Red + ", " + Darkblue + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Будем рисовать персонажа «Молния Маккуин» из мульфильма «Тачки». Для этого Вам нужно пять цветов\n" + Yellow + "\n" + Orange + "\n" + Red + "\n" + Darkblue + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Поздравляю! Вы нарисовали персонажа «Молния Маккуин»."
	pictureType = "Образец для рисования персонажа «Молния Маккуин»."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 5; colorsAndQuantity[Orange] = 3; colorsAndQuantity[Red] = 14; colorsAndQuantity[Darkblue] = 2; colorsAndQuantity[Black] = 7
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel9
	pictureID = pictureIDEasyLevel9
	pictureDraw = false
	return resp
}

func easyLevel10(resp Response) Response {
	resp.Text = "Нарисуем белку. Для этого Вам нужно 2 цвета: " + Orange + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Нарисуем белку. Для этого Вам нужно два цвета\n" + Orange + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Поздравляю! Вы нарисовали белку."
	pictureType = "Образец для рисования белки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Orange] = 13; colorsAndQuantity[Black] = 15
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDEasyLevel10
	pictureID = pictureIDEasyLevel10
	pictureDraw = false
	return resp
}

func mediumLevel1(resp Response) Response {
	resp.Text = "Нарисуем морскую свинку. Для этого Вам нужно подготовить 3 цвета: " + LightPink + ", " + Brown + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Нарисуем морскую свинку. Для этого Вам нужно подготовить три цв+ета\n" + LightPink + "\n" + Brown + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Поздравляю! Вы нарисовали морскую свинку."
	pictureType = "Образец для рисования морской свинки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[LightPink] = 2; colorsAndQuantity[Brown] = 24; colorsAndQuantity[Black] = 14
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel1
	pictureID = pictureIDMediumLevel1
	pictureDraw = false
	return resp
}

func mediumLevel2(resp Response) Response {
	resp.Text = "Нарисуем котенка. Для этого Вам понадобятся 5 цветов: " + Green + ", " + LightPink + ", " + Red + ", " + DarkGray + " и " + Black + ". Какой цвет хотите выбрать?"
	resp.TTS = "Нарисуем котенка. Для этого Вам понадобятся пять цвет+ов\n" + Green + "\n" + LightPink + "\n" + Red + "\n" + DarkGray + " - и " + Black + ". Какой цвет хотите выбрать?"

	resultText = "Поздравляю с нарисованным котенком!"
	pictureType = "Образец для рисования котенка."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Green] = 2; colorsAndQuantity[LightPink] = 5; colorsAndQuantity[Red] = 2; colorsAndQuantity[DarkGray] = 16; colorsAndQuantity[Black] = 20
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel2
	pictureID = pictureIDMediumLevel2
	pictureDraw = false
	return resp
}

func mediumLevel3(resp Response) Response {
	resp.Text = "Научимся рисовать синего кита. Для этого Вы должны подготовить 3 цвета: " + Blue + ", " + Darkblue + " и " + Black + ". Какой цвет Вы хотите выбрать?"
	resp.TTS = "Научимся рисовать синего кит+а. Для этого Вы должны подготовить три цв+ета\n" + Blue + "\n" + Darkblue + " - и " + Black + ". Какой цвет Вы хотите выбрать?"

	resultText = "Поздравляю! Вы нарисовали синего кита."
	pictureType = "Образец для рисования синего кита."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Blue] = 13; colorsAndQuantity[Darkblue] = 8; colorsAndQuantity[Black] = 12
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel3
	pictureID = pictureIDMediumLevel3
	pictureDraw = false
	return resp
}

func mediumLevel4(resp Response) Response {
	resp.Text = "Попробуем нарисовать сову. Для этого подготовьте 4 цвета: " + Beige + ", " + Yellow + ", " + Brown + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Попробуем нарисовать сову. Для этого подготовьте четыре цв+ета\n" + Beige + "\n" + Yellow + "\n" + Brown + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Поздравляю! Сова нарисована."
	pictureType = "Образец для рисования совы."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 15; colorsAndQuantity[Yellow] = 5; colorsAndQuantity[Brown] = 17; colorsAndQuantity[Black] = 5
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel4
	pictureID = pictureIDMediumLevel4
	pictureDraw = false
	return resp
}

func mediumLevel5(resp Response) Response {
	resp.Text = "Нарисуем персонажа «Миньон» из мультфильма «Гадкий я». Перед рисованием подготовьте 6 цветов: " + DarkGray + ", " + Yellow + ", " + Red + ", " + Darkblue + ", " + Brown + " и " + Black + ". Какой цвет выберем?"
	resp.TTS = "Нарисуем персонажа - Миньон - из мультфильма «Гадкий я». Перед рисованием подготовьте шесть цвет+ов\n" + DarkGray + "\n" + Yellow + "\n" + Red + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". Какой цвет выберем?"

	resultText = "Миньон нарисован. Поздравляю!"
	pictureType = "Образец для рисования миньона."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[DarkGray] = 8; colorsAndQuantity[Yellow] = 10; colorsAndQuantity[Red] = 2; colorsAndQuantity[Darkblue] = 6; colorsAndQuantity[Brown] = 1; colorsAndQuantity[Black] = 11
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel5
	pictureID = pictureIDMediumLevel5
	pictureDraw = false
	return resp
}

func mediumLevel6(resp Response) Response {
	resp.Text = "Нарисуем персонажа «Железный Человек» из вселенной «Marvel». Для рисования подготовьте 3 цвета: " + Yellow + ", " + Red + " и " + Black + ". Какой цвет выберем?"
	resp.TTS = "Нарисуем персонажа «Железный Человек» из вселенной «Marvel». Для рисования подготовьте три цвета\n" + Yellow + "\n" + Red + " - и " + Black + ". Какой цвет выберем?"

	resultText = "Вы нарисовали Железного Человека. Поздравляю!"
	pictureType = "Образец для рисования Железного Человека."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 13; colorsAndQuantity[Red] = 15; colorsAndQuantity[Black] = 17
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel6
	pictureID = pictureIDMediumLevel6
	pictureDraw = false
	return resp
}

func mediumLevel7(resp Response) Response {
	resp.Text = "Нарисуем коалу. Для этого подготовьте 6 цветов: " + LightGray + ", " + DarkGray + ", " + Orange + ", " + Green + ", " + Brown + " и " + Black + ". Какой цвет выберем первым?"
	resp.TTS = "Нарисуем коалу. Для этого подготовьте шесть цветов\n" + LightGray + "\n" + DarkGray + "\n" + Orange + "\n" + Green + "\n" + Brown + " - и " + Black + ". Какой цвет выберем первым?"

	resultText = "Коала нарисован. Поздравляю!"
	pictureType = "Образец для рисования коалы."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[LightGray] = 14; colorsAndQuantity[DarkGray] = 17; colorsAndQuantity[Orange] = 8; colorsAndQuantity[Green] = 5; colorsAndQuantity[Brown] = 9; colorsAndQuantity[Black] = 4
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel7
	pictureID = pictureIDMediumLevel7
	pictureDraw = false
	return resp
}

func mediumLevel8(resp Response) Response {
	resp.Text = "Нарисуем попугая. Для этого подготовьте 6 цветов: " + DarkGray + ", " + Yellow + ", " + Green + ", " + Darkblue + ", " + Brown + " и " + Black + ". Какой цвет выберем первым?"
	resp.TTS = "Нарисуем попугаем. Для этого подготовьте шесть цветов\n" + DarkGray + "\n" + Yellow + "\n" + Green + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". Какой цвет выберем первым?"

	resultText = "Попугай нарисован. Поздравляю!"
	pictureType = "Образец для рисования попугая."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[DarkGray] = 6; colorsAndQuantity[Yellow] = 13; colorsAndQuantity[Green] = 15; colorsAndQuantity[Darkblue] = 6; colorsAndQuantity[Brown] = 9; colorsAndQuantity[Black] = 13
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel8
	pictureID = pictureIDMediumLevel8
	pictureDraw = false
	return resp
}

func mediumLevel9(resp Response) Response {
	resp.Text = "Нарисуем персонажа «Стич» из мультфильма «Лило и Стич». Вам понадобятся 4 цвета: "  + LightPink + ", " + Blue + ", " + Darkblue + " и " + Black + ". Какой цвет выберем первым?"
	resp.TTS = "Нарисуем персонажа «Стич» из мультфильма «Лило и Стич». Вам понадобятся четыре цвета\n" + LightPink + "\n" + Blue + "\n" + Darkblue + " - и " + Black + ". Какой цвет выберем первым?"

	resultText = "Персонаж «Стич» нарисован. Поздравляю!"
	pictureType = "Образец для рисования персонажа «Стич»."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[LightPink] = 8; colorsAndQuantity[Blue] = 15; colorsAndQuantity[Darkblue] = 23; colorsAndQuantity[Black] = 23
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel9
	pictureID = pictureIDMediumLevel9
	pictureDraw = false
	return resp
}

func mediumLevel10(resp Response) Response {
	resp.Text = "Нарисуем пингвина. Вам понадобятся 3 цвета: " + Yellow + ", " + Orange + " и " + Black + ". Какой цвет выберем первым?"
	resp.TTS = "Нарисуем пингвина. Вам понадобятся три цвета\n" + Yellow + "\n" + Orange + " - и " + Black + ". Какой цвет выберем первым?"

	resultText = "Пингвин нарисован. Поздравляю!"
	pictureType = "Образец для рисования пингвина."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 5; colorsAndQuantity[Orange] = 9; colorsAndQuantity[Black] = 19
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDMediumLevel10
	pictureID = pictureIDMediumLevel10
	pictureDraw = false
	return resp
}

func hardLevel1(resp Response) Response {
	resp.Text = "Давайте нарисуем девушку Моана из мультфильма «Моана». Перед тем как приступить к рисованию Вам необходимо подготовить 8 цветов: " + Beige + ", " + Yellow + ", " + LightPink + ", " + Red + ", " + Blue + ", " + Darkblue + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Давайте нарисуем девушку - Моана - из мультфильма «Моана». Перед тем как приступить к рисованию Вам необходимо подготовить восемь цветов\n" + Beige + "\n" + Yellow + "\n" + LightPink + "\n" + Red + "\n" + Blue + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Моана нарисована. Поздравляю!"
	pictureType = "Образец для рисования девушки Моана."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 17; colorsAndQuantity[Yellow] = 3; colorsAndQuantity[LightPink] = 2; colorsAndQuantity[Red] = 8; colorsAndQuantity[Blue] = 1; colorsAndQuantity[Darkblue] = 2; colorsAndQuantity[Brown] = 9; colorsAndQuantity[Black] = 22
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel1
	pictureID = pictureIDHardLevel1
	pictureDraw = false
	return resp
}

func hardLevel2(resp Response) Response {
	resp.Text = "Нарисуем девушку Мерида из мультфильма «Храбрая сердцем». Для этого вам понадобятся 8 цветов: " + Beige + ", " + Yellow + ", " + Orange + ", " + LightPink + ", " + Blue + ", " + Darkblue + ", " + Brown + " и " + Black + ". Какой цвет выберем первым?"
	resp.TTS = "Нарисуем девушку - Мер+ида - из мультфильма «Храбрая сердцем». Для этого вам понадобятся восемь цветов\n" + Beige + "\n" + Yellow + "\n" + Orange + "\n" + LightPink + "\n" + Blue + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". Какой цвет выберем первым?"

	resultText = "Мерида нарисована. Поздравляю!"
	pictureType = "Образец для рисования девушки {Мерида}{Мер`ида}."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 12; colorsAndQuantity[Yellow] = 7; colorsAndQuantity[Orange] = 19; colorsAndQuantity[LightPink] = 2; colorsAndQuantity[Blue] = 1; colorsAndQuantity[Darkblue] = 16; colorsAndQuantity[Brown] = 10; colorsAndQuantity[Black] = 7
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel2
	pictureID = pictureIDHardLevel2
	pictureDraw = false
	return resp
}

func hardLevel3(resp Response) Response {
	resp.Text = "Изобразим лошадь. Для этого подготовьте 2 цвета: " + Brown + " и " + Black + ". Какой цвет нарисуем первым?"
	resp.TTS = "Изобразим лошадь. Для этого подготовьте два цвета\n" +  Brown + " - и " + Black + ". Какой цвет нарисуем первым?"

	resultText = "Лошадь нарисована."
	pictureType = "Образец для рисования лошади."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Brown] = 13; colorsAndQuantity[Black] = 18
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel3
	pictureID = pictureIDHardLevel3
	pictureDraw = false
	return resp
}

func hardLevel4(resp Response) Response {
	resp.Text = "Научимся рисовать собачку. Перед рисованием вам нужно подготовить 6 цветов: " + Beige + ", " + Orange + ", " + LightPink + ", " + Red + ", " + Brown + " и " + Black + ". Какой цвет будем рисовать первым?"
	resp.TTS = "Научимся рисовать собачку. Перед рисованием вам нужно подготовить шесть цветов\n" + Beige + "\n" + Orange + "\n" + LightPink + "\n" + Red + "\n" + Brown + " - и " + Black + ". Какой цвет будем рисовать первым?"

	resultText = "Собачка нарисована. Поздравляю!"
	pictureType = "Образец для рисования собачки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 8; colorsAndQuantity[Orange] = 13; colorsAndQuantity[LightPink] = 3; colorsAndQuantity[Red] = 1; colorsAndQuantity[Brown] = 19; colorsAndQuantity[Black] = 6
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel4
	pictureID = pictureIDHardLevel4
	pictureDraw = false
	return resp
}

func hardLevel5(resp Response) Response {
	resp.Text = "Попробуем изобразить слоненка. Для этого вам нужно 6 цветов: " + Yellow + ", " + LightPink + ", " + Blue + ", " + Purple + ", " + DarkGray + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Попробуем изобразить слоненка. Для этого вам нужно шесть цветов\n" + Yellow + "\n" + LightPink + "\n" + Blue + "\n" + Purple + "\n" + DarkGray + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Слоненок нарисован."
	pictureType = "Образец для рисования слоненка."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 7; colorsAndQuantity[LightPink] = 12; colorsAndQuantity[Blue] = 20; colorsAndQuantity[Purple] = 19; colorsAndQuantity[DarkGray] = 20; colorsAndQuantity[Black] = 26
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel5
	pictureID = pictureIDHardLevel5
	pictureDraw = false
	return resp
}

func hardLevel6(resp Response) Response {
	resp.Text = "Попробуем изобразить черепаху. Для этого вам нужно 4 цвета: " + Yellow + ", " + Green + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Попробуем изобразить черепаху. Для этого вам нужно четыре цвета\n" + Yellow + "\n" + Green + "\n" + Brown + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Черепаха нарисована."
	pictureType = "Образец для рисования черепахи."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 15; colorsAndQuantity[Green] = 25; colorsAndQuantity[Brown] = 17; colorsAndQuantity[Black] = 2
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel6
	pictureID = pictureIDHardLevel6
	pictureDraw = false
	return resp
}

func hardLevel7(resp Response) Response {
	resp.Text = "Нарисуем персонажа «Капитан Америка» из вселенной «Marvel». Для этого вам нужно 4 цвета: " + Beige + ", " + Red + ", " + Darkblue + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем персонажа «Капитан Америка» из вселенной «Marvel». Для этого вам нужно четыре цвета\n" + Beige + "\n" + Red + "\n" + Darkblue + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Персонаж «Капитан Америка» нарисован."
	pictureType = "Образец для рисования персонажа «Капитан Америка»."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 5; colorsAndQuantity[Red] = 17; colorsAndQuantity[Darkblue] = 19; colorsAndQuantity[Black] = 27
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel7
	pictureID = pictureIDHardLevel7
	pictureDraw = false
	return resp
}

func hardLevel8(resp Response) Response {
	resp.Text = "Научимся рисовать фламинго. Для этого вам нужно 4 цвета: " + Yellow + ", " + LightPink + ", " + DarkPink + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Научимся рисовать фламинго. Для этого вам нужно четыре цвета\n" + Yellow + "\n" + LightPink + "\n" + DarkPink + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Фламинго нарисован."
	pictureType = "Образец для рисования фламинго."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 3; colorsAndQuantity[LightPink] = 12; colorsAndQuantity[DarkPink] = 17; colorsAndQuantity[Black] = 10
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel8
	pictureID = pictureIDHardLevel8
	pictureDraw = false
	return resp
}

func hardLevel9(resp Response) Response {
	resp.Text = "Сейчас я научу Вас рисовать вертолет. Для этого вам понадобятся 4 цвета: " + Yellow + ", " + Orange + ", " + Blue + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Сейчас я научу Вас рисовать вертолет. Для этого вам понадобятся четыре цвета\n" + Yellow + "\n" + Orange + "\n" + Blue + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Вертолет нарисован."
	pictureType = "Образец для рисования вертолета."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 24; colorsAndQuantity[Orange] = 21; colorsAndQuantity[Blue] = 8; colorsAndQuantity[Black] = 12
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel9
	pictureID = pictureIDHardLevel9
	pictureDraw = false
	return resp
}

func hardLevel10(resp Response) Response {
	resp.Text = "Нарисуем дракончика. Для этого вам понадобятся 4 цвета: " + LightPink + ", " + Blue + ", " + Purple + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем дракончика. Для этого вам понадобятся четыре цвета\n" + LightPink + "\n" + Blue + "\n" + Purple + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Дракончик нарисован."
	pictureType = "Образец для рисования дракончика."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[LightPink] = 22; colorsAndQuantity[Blue] = 24; colorsAndQuantity[Purple] = 25; colorsAndQuantity[Black] = 3
	sortColorsAndQuantity()
	for color := range colorsAndQuantity {
		resp.AddButton(strings.Title(color), Payload{
			Text: color,
		})
	}

	hintID = hintIDHardLevel10
	pictureID = pictureIDHardLevel10
	pictureDraw = false
	return resp
}

func sortColorsAndQuantity() {
	colorsSort = nil
	for col := range colorsAndQuantity {
		colorsSort = append(colorsSort, col)
	}
	sort.Strings(colorsSort)
}

func drawingLevel(resp Response) Response {
	if color == LightPinkWithoutDash || color == DarkPinkWithoutDash || color == LightGrayWithoutDash || color == DarkGrayWithoutDash {
		color = strings.Replace(color, " ", "-", -1)
	} else if color == "светлый розовый" || color == "светлый серый" {
		color = strings.Replace(color, "светлый ", "светло-", -1)
	} else if color == "темный розовый" || color == "темный серый" {
		color = strings.Replace(color, "темный ", "темно-", -1)
	} else if color == "светлый-розовый" || color == "светлый-серый" {
		color = strings.Replace(color, "светлый", "светло", -1)
	} else if color == "темный-розовый" || color == "темный-серый" {
		color = strings.Replace(color, "темный", "темно", -1)
	}

	switch selected {
	case Easy:
		resp = easyLevels(resp)
	case Medium:
		resp = mediumLevels(resp)
	case Hard:
		resp = hardLevels(resp)
	case Mixed:
		resp = mixedLevels(resp)
	}

	numbOfColoringRules = 0

	for _, value := range colorsAndQuantity {
		numbOfColoringRules += value
	}

	if numbOfColoringRules == 0 {
		resp.Text = strings.TrimSuffix(resp.Text, endDrawingText) + resultText
		resp.SSML = strings.TrimSuffix(resp.SSML, endDrawingSSML+"</speak>") + "\n" + resultText + "</speak>"
		resp.Card = NewBigImage(
			pictureID,
		)
		resp.AddButton("Следующий рисунок", Payload{
			Text: "следующий рисунок",
		})
		pictureDraw = true
	} else if colorsAndQuantity[color] != 0 {
		resp.AddButton("Дальше", Payload{
			Text: "дальше",
		})
	} else if colorsAndQuantity[color] == 0 {
		resp = drawColorButtons(resp)
		resp = mainButtons(resp)
	}

	return resp
}

func drawColorButtons(resp Response) Response {
	if selected != None {
		for color, value := range colorsAndQuantity {
			if value != 0 {
				resp.AddButton(strings.Title(color), Payload{
					Text: color,
				})
			}
		}
	}
	return resp
}

func getRemainingColors() string {
	var colorsText = ""
	var lenColors = 1

	var lenColorsAndQuantityWithoutZero = 0
	for _,value := range colorsAndQuantity {
		if value != 0 {
			lenColorsAndQuantityWithoutZero++
		}
	}

	if lenColorsAndQuantityWithoutZero == 0  {
		return colorsText
	}

	for _, col := range colorsSort {
		if colorsAndQuantity[col] != 0 {
			if lenColors < lenColorsAndQuantityWithoutZero - 1 {
				colorsText += col + ", "
			} else if lenColors == lenColorsAndQuantityWithoutZero - 1 {
				colorsText += col + " или "
			} else {
				colorsText += col + " "
			}
			lenColors++
		}
	}

	return colorsText
}

func getRemainingColorsForSpeak() string {
	var colorsText = ""
	var lenColors = 1

	var lenColorsAndQuantityWithoutZero = 0
	for _,value := range colorsAndQuantity {
		if value != 0 {
			lenColorsAndQuantityWithoutZero++
		}
	}

	if lenColorsAndQuantityWithoutZero == 0  {
		return colorsText
	}

	for _, col := range colorsSort {
		if colorsAndQuantity[col] != 0 {
			if lenColors < lenColorsAndQuantityWithoutZero - 1 {
				colorsText += col + breakTimeShort
			} else if lenColors == lenColorsAndQuantityWithoutZero - 1 {
				colorsText += col + " или "
			} else {
				colorsText += col + " "
			}
			lenColors++
		}
	}

	return colorsText
}

func easyLevels(resp Response) Response {
	switch pictureID {
	case pictureIDEasyLevel1:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 2:
				resp.Text = "Е—8, Е—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемь, Е-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—8, Ё—9\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-восемь, Йо-девять" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Желтый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Желтый цвет уже ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 9:
				resp.Text = "От А—6 до А—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-шесть до А-девять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 8:
				resp.Text = "От Б—5 до Б—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-пять до Бэ-десять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "От В—4 до В—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Вэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "От Г—4 до Г—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-четыре до Гэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "От Д—5 до Д—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "От Е—4 до Е—7, Е—10 и Е—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-семь" + breakTimeLong +
						"Е-десять" + breakTimeLong +
						"И - Е-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "От Ё—4 до Ё—7, Ё—10 и Ё—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четыре до Йо-семь" + breakTimeLong +
						"Йо-десять" + breakTimeLong +
						"И - Йо-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "От Ж—5 до Ж—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-пять до Жэ-десять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—6 до З—9\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-шесть до Зэ-девять" + breakTimeLong +
						"Красный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Красный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Красный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Green:
			switch colorsAndQuantity[Green] {
			case 3:
				resp.Text = "Е—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-один" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "Ё—1, Ё—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-один, Йо-два" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—2\nЗеленый цвет закончился. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-два" + breakTimeLong +
						"Зеленый цвет ^закончился^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Зеленый цвет закончился. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Зеленый цвет ^закончился^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Д—2 до Д—4\nВыберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-два до Дэ-четыре" + breakTimeLong +
						"Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel2:
		switch color {
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 3:
				resp.Text = "От А—4 до Ж—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-четыре до Жэ-четыре" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "От А—5 до Ж—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-пять до Жэ-пять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—6 до Ж—6\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-шесть до Жэ-шесть" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"

			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 3:
				resp.Text = "От В—1 до Е—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-один до Е-один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "От Б—2 до Ё—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-два до Йо-два" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—3 до Ж—3\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-три до Жэ-три" + breakTimeLong +
						"Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Красный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Красный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"

			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 3:
				resp.Text = "От А—7 до Ж—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-семь до Жэ-семь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "От А—8 до Ж—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-восемь до Жэ-восемь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—9 до Ж—9\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-девять до Жэ-девять" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Голубой цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Голубой цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"

			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 2:
				resp.Text = "От А—10 до Ж—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-десять до Жэ-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—11 до Ж—11\nСиний цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-одиннадцать до Жэ-одиннадцать" + breakTimeLong +
						"Синий цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет закончен. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 4:
				resp.Text = "Г—12, Д—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-двенадцать, Дэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "Г—13, Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-тринадцать, Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "Г—14, Д—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четырнадцать, Дэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Г—15, Д—15\nКоричневый цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-пятнадцать, Дэ-пятнадцать" + breakTimeLong +
						"Коричневый цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет закончен. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel3:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 7:
				resp.Text = "Б—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четыре" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "В—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "Г—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-пять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "Д—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "Е—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-пять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "Ё—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—4\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четыре" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 9:
				resp.Text = "А—4, А—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-четыре, А-пять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 8:
				resp.Text = "Б—3 и от Б—5 до Б—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-три" + breakTimeLong +
						"И - от Бэ-пять до Бэ-семь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "От В—3 до В—6 и В—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-три до Вэ-шесть" + breakTimeLong +
						"И - Вэ-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "Г—4 и от Г—6 до Г—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четыре" + breakTimeLong +
						"И - от Гэ-шесть до Гэ-девять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "От Д—5 до Д—7, Д—9 и Д—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-семь" + breakTimeLong +
						"Дэ-девять" + breakTimeLong +
						"И - Дэ-десять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "Е—4 и от Е—6 до Е—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четыре" + breakTimeLong +
						"И - от Е-шесть до Е-девять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "От Ё—3 до Ё—6 и Ё—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-три до Йо-шесть" + breakTimeLong +
						"И - Йо-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "Ж—3 и от Ж—5 до Ж—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"И - от Жэ-пять до Жэ-семь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—4, З—5\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четыре, Зэ-пять" + breakTimeLong +
						"Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже закрасили клетки красным цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже закрасили клетки красным цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"

			}
		case Green:
			switch colorsAndQuantity[Green] {
			case 5:
				resp.Text = "В—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-два" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "Г—2, Г—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два, Гэ-три" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "От Д—1 до Д—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-один до Дэ-четыре" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "Е—2, Е—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-два, Е-три" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—2\nЗеленый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-два" + breakTimeLong +
						"Зеленый цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Зеленый цвет использован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Зеленый цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже использовали зеленый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже использовали зеленый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel4:
		switch color {
		case Green:
			switch colorsAndQuantity[Green] {
			case 13:
				resp.Text = "А—3, А—4 и от А—9 до А—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-три" + breakTimeLong +
						"А-четыре" + breakTimeLong +
						"И - от А-девять до А-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 12:
				resp.Text = "От Б—3 до Б—6 и от Б—8 до Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-три до Бэ-шесть" + breakTimeLong +
						"И - от Бэ-восемь до Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 11:
				resp.Text = "От В—4 до В—6, В—8, В—9, В—12 и В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Вэ-шесть" + breakTimeLong +
						"Вэ-восемь" + breakTimeLong +
						"Вэ-девять" + breakTimeLong +
						"Вэ-двенадцать" + breakTimeLong +
						"И - Вэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 10:
				resp.Text = "От Г—4 до Г—6, Г—8, Г—9, Г—12 и Г—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-четыре до Гэ-шесть" + breakTimeLong +
						"Гэ-восемь" + breakTimeLong +
						"Гэ-девять" + breakTimeLong +
						"Гэ-двенадцать" + breakTimeLong +
						"И - Гэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 9:
				resp.Text = "Д—5, Д—6, от Д—8 до Д—10, Д—12 и Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-пять" + breakTimeLong +
						"Дэ-шесть" + breakTimeLong +
						"От Дэ-восемь до Дэ-десять" + breakTimeLong +
						"Дэ-двенадцать" + breakTimeLong +
						"И - Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 8:
				resp.Text = "От Е—4 до Е—6, Е—9, Е—10, Е—12 и Е—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-шесть" + breakTimeLong +
						"Е-девять" + breakTimeLong +
						"Е-десять" + breakTimeLong +
						"Е-двенадцать" + breakTimeLong +
						"И - Е-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 7:
				resp.Text = "От Ё—3 до Ё—5, Ё—9, Ё—10, Ё—12 и Ё—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-три до Йо-пять" + breakTimeLong +
						"Йо-девять" + breakTimeLong +
						"Йо-десять" + breakTimeLong +
						"Йо-двенадцать" + breakTimeLong +
						"И - Йо-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 6:
				resp.Text = "От Ж—2 до Ж—4, Ж—9, Ж—10, Ж—12 и Ж—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-два до Жэ-четыре" + breakTimeLong +
						"Жэ-девять" + breakTimeLong +
						"Жэ-десять" + breakTimeLong +
						"Жэ-двенадцать" + breakTimeLong +
						"И - Жэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 5:
				resp.Text = "От З—1 до З—3 и от З—8 до З—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-один до Зэ-три" + breakTimeLong +
						"И - от Зэ-восемь до Зэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "И—1, И—2, от И—7 до И—9, И—11, И—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-один" + breakTimeLong +
						"И-два" + breakTimeLong +
						"От И-семь до И-девять" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"И-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "От Й—1 до Й—3, от Й—6 до Й—8, Й—11, Й—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - один до И-краткое - три" + breakTimeLong +
						"От И-краткое - шесть до И-краткое - восемь" + breakTimeLong +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И-краткое - двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "От К—2 до К—7, К—11 и К—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-два до Ка-семь" + breakTimeLong +
						"Ка-одиннадцать" + breakTimeLong +
						"И - Ка-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Л—3 до Л—6, Л—12 и Л—13\nЗеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-три до Эль-шесть" + breakTimeLong +
						"Эль-двенадцать" + breakTimeLong +
						"И - Эль-тринадцать" + breakTimeLong +
						"Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Зеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Зеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зеленый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 2:
				resp.Text = "А—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "В—3\nЧерный цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-три" + breakTimeLong +
						"Черный цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет использован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel5:
		switch color {
		case Purple:
			switch colorsAndQuantity[Purple] {
			case 7:
				resp.Text = "Б—11, Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-одиннадцать, Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 6:
				resp.Text = "От В—7 до В—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-семь до Вэ-девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 5:
				resp.Text = "Д—11, Д—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-одиннадцать, Дэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 4:
				resp.Text = "Е—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 3:
				resp.Text = "От Ё—7 до Ё—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-семь до Йо-девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 2:
				resp.Text = "И—7, И—10, И—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-семь" + breakTimeLong +
						"И-десять" + breakTimeLong +
						"И-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 1:
				colorsAndQuantity[Purple]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Й—8, Й—9\nВы закрасили все клетки фиолетовым цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - восемь, И-краткое - девять" + breakTimeLong +
						"Вы закрасили все клетки фиолетовым цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы закрасили все клетки фиолетовым цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы закрасили все клетки фиолетовым цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Фиолетовый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Фиолетовый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 10:
				resp.Text = "От В—10 до В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-десять до Вэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 9:
				resp.Text = "Г—3, Г—4 и от Г—6 до Г—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-три" + breakTimeLong +
						"Гэ-четыре" + breakTimeLong +
						"И - от Гэ-шесть до Гэ-девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 8:
				resp.Text = "От Д—2 до Д—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-два до Дэ-семь" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 7:
				resp.Text = "Е—1, Е—2, от Е—4 до Е—6 и от Е—11 до Е—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-один" + breakTimeLong +
						"Е-два" + breakTimeLong +
						"От Е-четыре до Е-шесть" + breakTimeLong +
						"И - от Е-одиннадцать до Е-тринадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 6:
				resp.Text = "От Ё—1 до Ё—6 и Ё—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-один до Йо-шесть" + breakTimeLong +
						"И - Йо-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "Ж—1, Ж—2 и от Ж—4 до Ж—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-один" + breakTimeLong +
						"Жэ-два" + breakTimeLong +
						"И - от Жэ-четыре до Жэ-девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "От З—2 до З—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-два до Зэ-шесть" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "И—3, И—4, И—6, И—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-три" + breakTimeLong +
						"И-четыре" + breakTimeLong +
						"И-шесть" + breakTimeLong +
						"И-двенадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "Й—7, Й—10, Й—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - семь" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"И-краткое - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "К—8, К—9\nСиний цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-восемь, Ка-девять" + breakTimeLong +
						"Синий цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет использован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 2:
				resp.Text = "Е—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—3\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel6:
			switch color {
			case Red:
				switch colorsAndQuantity[Red] {
				case 9:
					resp.Text = "От А—7 до А—11"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От А-семь до А-одиннадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 8:
					resp.Text = "От Б—6 до Б—8 и от Б—10 до Б—12"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Бэ-шесть до Бэ-восемь" + breakTimeLong +
							"И - от Бэ-десять до Бэ-двенадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 7:
					resp.Text = "В—5, В—6, от В—8 до В—10, В—12 и В—13"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Вэ-пять" + breakTimeLong +
							"Вэ-шесть" + breakTimeLong +
							"От Вэ-восемь до Вэ-десять" + breakTimeLong +
							"Вэ-двенадцать" + breakTimeLong +
							"И - Вэ-тринадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 6:
					resp.Text = "От Г—5 до Г—13"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Гэ-пять до Гэ-тринадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 5:
					resp.Text = "Д—13"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Дэ-тринадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 4:
					resp.Text = "От Е—5 до Е—13"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Е-пять до Е-тринадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 3:
					resp.Text = "Ё—5, Ё—6, от Ё—8 до Ё—10, Ё—12 и Ё—13"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Йо-пять" + breakTimeLong +
							"Йо-шесть" + breakTimeLong +
							"От Йо-восемь до Йо-десять" + breakTimeLong +
							"Йо-двенадцать" + breakTimeLong +
							"И - Йо-тринадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 2:
					resp.Text = "От Ж—6 до Ж—8 и от Ж—10 до Ж—12"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Жэ-шесть до Жэ-восемь" + breakTimeLong +
							"И - от Жэ-десять до Жэ-двенадцать" +
							"</speak>"
					colorsAndQuantity[Red]--
				case 1:
					colorsAndQuantity[Red]--
					var getRemainingColors = getRemainingColors()
					var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
					resp.Text = "От З—7 до З—11\nВы закрасили все клетки красным цветом. Выберите " + getRemainingColors + "цвет."
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Зэ-семь до Зэ-одиннадцать" + breakTimeLong +
							"Вы закрасили все клетки красным цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
							"</speak>"
					endDrawingText = "Вы закрасили все клетки красным цветом. Выберите " + getRemainingColors + "цвет."
					endDrawingSSML = "Вы закрасили все клетки красным цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
				case 0:
					var getRemainingColors = getRemainingColors()
					var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
					resp.Text = "Красный цвет использован. Выберите " + getRemainingColors + "цвет."
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Красный цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
							"</speak>"
				}
			case Black:
				switch colorsAndQuantity[Black] {
				case 7:
					resp.Text = "Б—9"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Бэ-девять" +
							"</speak>"
					colorsAndQuantity[Black]--
				case 6:
					resp.Text = "В—1, В—4, В—7 и В—11"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Вэ-один" + breakTimeLong +
							"Вэ-четыре" + breakTimeLong +
							"Вэ-семь" + breakTimeLong +
							"И - Вэ-одиннадцать" +
							"</speak>"
					colorsAndQuantity[Black]--
				case 5:
					resp.Text = "От Г—1 до Г—4"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Гэ-один до Гэ-четыре" +
							"</speak>"
					colorsAndQuantity[Black]--
				case 4:
					resp.Text = "От Д—3 до Д—12"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Дэ-три до Дэ-двенадцать" +
							"</speak>"
					colorsAndQuantity[Black]--
				case 3:
					resp.Text = "От Е—1 до Е—4"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"От Е-один до Е-четыре" +
							"</speak>"
					colorsAndQuantity[Black]--
				case 2:
					resp.Text = "Ё—1, Ё—4, Ё—7 и Ё—11"
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Йо-один" + breakTimeLong +
							"Йо-четыре" + breakTimeLong +
							"Йо-семь" + breakTimeLong +
							"И - Йо-одиннадцать" +
							"</speak>"
					colorsAndQuantity[Black]--
				case 1:
					colorsAndQuantity[Black]--
					var getRemainingColors = getRemainingColors()
					var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
					resp.Text = "Ж—9\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Жэ-девять" + breakTimeLong +
							"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
							"</speak>"
					endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
					endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
				case 0:
					var getRemainingColors = getRemainingColors()
					var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
					resp.Text = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
					resp.TTSType = "ssml"
					resp.SSML =
						"<speak>" +
							"Черный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
							"</speak>"
				}
			default:
				resp = processingDefaultColor(resp)
			}
	case pictureIDEasyLevel7:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 5:
				resp.Text = "Г—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "От Д—2 до Д—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-два до Дэ-четыре" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От Е—1 до Е—3 и Е—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-один до Е-три" + breakTimeLong +
						"И - Е-восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От Ё—2 до Ё—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-два до Йо-четыре" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—2\nВы закрасили все клетки желтым цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-два" + breakTimeLong +
						"Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Желтый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Желтый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 2:
				resp.Text = "В—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—11\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-одиннадцать" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—5, Е—12\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-пять, Е-двенадцать" + breakTimeLong +
						"Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Красный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Green:
			switch colorsAndQuantity[Green] {
			case 11:
				resp.Text = "А—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 10:
				resp.Text = "Б—9, Б—11 и Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-девять" + breakTimeLong +
						"Бэ-одиннадцать" + breakTimeLong +
						"И - Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 9:
				resp.Text = "В—7, В—9, В—10, В—12 и В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-семь" + breakTimeLong +
						"Вэ-девять" + breakTimeLong +
						"Вэ-десять" + breakTimeLong +
						"Вэ-двенадцать" + breakTimeLong +
						"И - Вэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 8:
				resp.Text = "Г—5, Г—7, Г—9, Г—12 и Г—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-пять" + breakTimeLong +
						"Гэ-семь" + breakTimeLong +
						"Гэ-девять" + breakTimeLong +
						"Гэ-двенадцать" + breakTimeLong +
						"И - Гэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 7:
				resp.Text = "От Д—5 до Д—9, Д—12 и Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-девять" + breakTimeLong +
						"Дэ-двенадцать" + breakTimeLong +
						"И - Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 6:
				resp.Text = "Е—4, Е—6, Е—7, от Е—9 до Е—11 и Е—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четыре" + breakTimeLong +
						"Е-шесть" + breakTimeLong +
						"Е-семь" + breakTimeLong +
						"От Е-девять до Е-одиннадцать" + breakTimeLong +
						"И - Е-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 5:
				resp.Text = "От Ё—5 до Ё—9, Ё—12 и Ё—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пять до Йо-девять" + breakTimeLong +
						"Йо-двенадцать" + breakTimeLong +
						"И - Йо-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "Ж—5, от Ж—7 до Ж—9, Ж—12 и Ж—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-пять" + breakTimeLong +
						"От Жэ-семь до Жэ-девять" + breakTimeLong +
						"Жэ-двенадцать" + breakTimeLong +
						"И - Жэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "З—7, З—9, З—10, З—12 и З—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-семь" + breakTimeLong +
						"Зэ-девять" + breakTimeLong +
						"Зэ-десять" + breakTimeLong +
						"Зэ-двенадцать" + breakTimeLong +
						"И - Зэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "И—9, И—11, И—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-девять" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"И-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Й—11\nВы закрасили все клетки зеленым цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - одиннадцать" + breakTimeLong +
						"Вы закрасили все клетки зеленым цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы закрасили все зеленым цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы закрасили все зеленым цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Зеленый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зеленый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Г—8\nСиний цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-восемь" + breakTimeLong +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 3:
				resp.Text = "Д—14, Д—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-четырнадцать, Дэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "Е—14, Е—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четырнадцать, Е-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—14, Ё—15\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-четырнадцать, Йо-пятнадцать" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 4:
				resp.Text = "Г—10, Г—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-десять, Гэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Д—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Ё—10, Ё—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-десять, Йо-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—11\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-одиннадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel8:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 10:
				resp.Text = "А—8, А—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-восемь, А-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 9:
				resp.Text = "Б—8, Б—10 и Б—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемь" + breakTimeLong +
						"Бэ-десять" + breakTimeLong +
						"И - Бэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "От В—2 до В—6, В—8 и В—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-два до Вэ-шесть" + breakTimeLong +
						"Вэ-восемь" + breakTimeLong +
						"И - Вэ-десять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "Г—12, Г—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-двенадцать, Гэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "Д—12, Д—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двенадцать, Дэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "Й—12, Й—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - двенадцать, И-краткое - четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "К—12, К—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-двенадцать, Ка-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От Л—2 до Л—6, Л—8 и Л—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-два до Эль-шесть" + breakTimeLong +
						"Эль-восемь" + breakTimeLong +
						"И - Эль-десять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "М—8, М—10 и М—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-восемь" + breakTimeLong +
						"Эм-десять" + breakTimeLong +
						"И - Эм-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Н—8, Н—9\nВы закрасили все клетки желтым цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-восемь, Эн-девять" + breakTimeLong +
						"Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Желтый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Желтый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 9:
				resp.Text = "Г—1, Г—3 и от Г—6 до Г—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-один" + breakTimeLong +
						"Гэ-три" + breakTimeLong +
						"И - от Гэ-шесть до Гэ-десять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 8:
				resp.Text = "От Д—5 до Д—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "От Е—5 до Е—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пять до Е-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "От Ё—5 до Ё—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пять до Йо-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "От Ж—5 до Ж—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-пять до Жэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "От З—5 до З—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-пять до Зэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От И—5 до И—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-пять до И-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "От Й—5 до Й—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - пять до И-краткое - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "К—1, К—3 и от К—6 до К—10\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-один" + breakTimeLong +
						"Ка-три" + breakTimeLong +
						"И - от Ка-шесть до Ка-десять" + breakTimeLong +
						"Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Оранжевый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 2:
				resp.Text = "Ё—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—4\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четыре" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel9:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 5:
				resp.Text = "А—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "Б—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четыре" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "В—4, В—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четыре, Вэ-пять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "Ж—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-пять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—5\nВы закрасили все клетки желтым цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-пять" + breakTimeLong +
						"Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы закрасили все клетки желтым цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Желтый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Желтый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 3:
				resp.Text = "А—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-два" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "Б—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-три" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "В—3\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-три" + breakTimeLong +
						"Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Оранжевый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 14:
				resp.Text = "А—1, А—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-один, А-четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 13:
				resp.Text = "Б—1, Б—2, Б—5 и Б—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-один" + breakTimeLong +
						"Бэ-два" + breakTimeLong +
						"Бэ-пять" + breakTimeLong +
						"И - Бэ-шесть" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 12:
				resp.Text = "В—2, В—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-два, Вэ-шесть" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 11:
				resp.Text = "От Г—2 до Г—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-два до Гэ-четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 10:
				resp.Text = "Д—1, Д—4, Д—6 и Д—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-один" + breakTimeLong +
						"Дэ-четыре" + breakTimeLong +
						"Дэ-шесть" + breakTimeLong +
						"И - Дэ-семь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 9:
				resp.Text = "Е—1, Е—4 и Е—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-один" + breakTimeLong +
						"Е-четыре" + breakTimeLong +
						"И - Е-пять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 8:
				resp.Text = "Ё—1 и от Ё—4 до Ё—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-один" + breakTimeLong +
						"И - от Йо-четыре до Йо-восемь" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "Ж—1, Ж—4 и от Ж—6 до Ж—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-один" + breakTimeLong +
						"Жэ-четыре" + breakTimeLong +
						"И - от Жэ-шесть до Жэ-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "З—2, З—4, З—5, З—7 и З—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-два" + breakTimeLong +
						"Зэ-четыре" + breakTimeLong +
						"Зэ-пять" + breakTimeLong +
						"Зэ-семь" + breakTimeLong +
						"И - Зэ-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "От И—3 до И—5, И—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-три до И-пять" + breakTimeLong +
						"И-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "Й—4, Й—5, Й—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - четыре" + breakTimeLong +
						"И-краткое - пять" + breakTimeLong +
						"И-краткое - восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "К—4, К—5 и К—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-четыре" + breakTimeLong +
						"Ка-пять" + breakTimeLong +
						"И - Ка-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "Л—4, Л—5 и Л—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-четыре" + breakTimeLong +
						"Эль-пять" + breakTimeLong +
						"И - Эль-семь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—6\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-шесть" + breakTimeLong +
						"Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Красный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 2:
				resp.Text = "Е—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—3\nСиний цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 7:
				resp.Text = "От А—5 до А—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-пять до А-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Б—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "От Г—5 до Г—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-пять до Гэ-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Д—5, Д—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-пять, Дэ-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От Е—6 до Е—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-шесть до Е-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Л—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—7, М—8\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-семь, Эм-восемь" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDEasyLevel10:
		switch color {
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 13:
				resp.Text = "Б—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 12:
				resp.Text = "От В—9 до В—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-девять до Вэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 11:
				resp.Text = "От Г—7 до Г—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-семь до Гэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 10:
				resp.Text = "Д—3, Д—4, Д—6, Д—7, Д—10 и Д—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-три" + breakTimeLong +
						"Дэ-четыре" + breakTimeLong +
						"Дэ-шесть" + breakTimeLong +
						"Дэ-семь" + breakTimeLong +
						"Дэ-десять" + breakTimeLong +
						"И - Дэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 9:
				resp.Text = "Е—4 и от Е—6 до Е—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четыре" + breakTimeLong +
						"И - от Е-шесть до Е-десять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 8:
				resp.Text = "От Ё—5 до Ё—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ё-пять до Ё-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "От Ж—5 до Ж—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-пять до Жэ-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "От З—5 до З—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-пять до Зэ-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "И—4 и от И—6 до И—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-четыре" + breakTimeLong +
						"И - от И-шесть до И-десять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "Й—3, Й—4, Й—6, Й—7, Й—10, Й—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - три" + breakTimeLong +
						"И-краткое - четыре" + breakTimeLong +
						"И-краткое - шесть" + breakTimeLong +
						"И-краткое - семь" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"И-краткое - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От К—7 до К—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-семь до Ка-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "От Л—9 до Л—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-девять до Эль-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—11\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-одиннадцать" + breakTimeLong +
						"Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Оранжевый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 15:
				resp.Text = "А—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Б—9, Б—10 и Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-девять" + breakTimeLong +
						"Бэ-десять" + breakTimeLong +
						"И - Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "От В—2 до В—8 и В—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-два до Вэ-восемь" + breakTimeLong +
						"И - Вэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Г—1, Г—6 и Г—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-один" + breakTimeLong +
						"Гэ-шесть" + breakTimeLong +
						"И - Гэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Д—2, Д—5, Д—8, Д—9 и Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-два" + breakTimeLong +
						"Дэ-пять" + breakTimeLong +
						"Дэ-восемь" + breakTimeLong +
						"Дэ-девять" + breakTimeLong +
						"И - Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Е—3, Е—5 и Е—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" + breakTimeLong +
						"Е-пять" + breakTimeLong +
						"И - Е-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "Ё—4, Ё—9, Ё—12 и Ё—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-четыре" + breakTimeLong +
						"Йо-девять" + breakTimeLong +
						"Йо-двенадцать" + breakTimeLong +
						"И - Йо-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "Ж—4, Ж—9, Ж—10, Ж—12 и Ж—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четыре" + breakTimeLong +
						"Жэ-девять" + breakTimeLong +
						"Жэ-десять" + breakTimeLong +
						"Жэ-двенадцать" + breakTimeLong +
						"И - Жэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "З—4, З—9, З—12 и З—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четыре" + breakTimeLong +
						"Зэ-девять" + breakTimeLong +
						"Зэ-двенадцать" + breakTimeLong +
						"И - Зэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "И—3, И—5, И—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-три" + breakTimeLong +
						"И-пять" + breakTimeLong +
						"И-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Й—2, Й—5, Й—8, Й—9, Й—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - два" + breakTimeLong +
						"И-краткое - пять" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"И-краткое - девять" + breakTimeLong +
						"И-краткое - тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "К—1, К—6 и К—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-один" + breakTimeLong +
						"Ка-шесть" + breakTimeLong +
						"И - Ка-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От Л—2 до Л—8 и Л—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-два до Эль-восемь" + breakTimeLong +
						"И - Эль-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "М—9, М—10 и М—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-девять" + breakTimeLong +
						"Эм-десять" + breakTimeLong +
						"И - Эм-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Н—11\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-одиннадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	}
	return resp
}

func mediumLevels(resp Response) Response {
	switch pictureID {
	case pictureIDMediumLevel1:
		switch color {
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 2:
				resp.Text = "Ё—11, З—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-одиннадцать, Зэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—12\nМожем переходить к другому цвету. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двенадцать" + breakTimeLong +
						"Можем переходить к другому цвету. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Можем переходить к другому цвету. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Можем переходить к другому цв+ету. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 24:
				resp.Text = "От А—11 до А—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-одиннадцать до А-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 23:
				resp.Text = "От Б—10 до Б—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-десять до Бэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 22:
				resp.Text = "От В—8 до В—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-восемь до Вэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 21:
				resp.Text = "От Г—4 до Г—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-четыре до Гэ-девять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 20:
				resp.Text = "От Д—2 до Д—4 и Д—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-два до Дэ-четыре" + breakTimeLong +
						"И - Дэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 19:
				resp.Text = "От Е—2 до Е—5, Е—21 и Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-два до Е-пять" + breakTimeLong +
						"Е-двадцать один" + breakTimeLong +
						"И - Е-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 18:
				resp.Text = "От Ё—1 до Ё—3 и Ё—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-один до Йо-три" + breakTimeLong +
						"И - Йо-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 17:
				resp.Text = "Ж—21, Ж—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать один, Жэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 16:
				resp.Text = "З—21, З—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двадцать один, Зэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 15:
				resp.Text = "И—1, И—2, И—22, И—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-один" + breakTimeLong +
						"И-два" + breakTimeLong +
						"И-двадцать два" + breakTimeLong +
						"И-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 14:
				resp.Text = "От Й—1 до Й—6, Й—9, Й—17, Й—18, Й—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - один до И-краткое - шесть" + breakTimeLong +
						"И-краткое - девять" + breakTimeLong +
						"И-краткое - семнадцать" + breakTimeLong +
						"И-краткое - восемнадцать" + breakTimeLong +
						"И-краткое - двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 13:
				resp.Text = "От К—1 до К—5, К—9, от К—16 до К—19, К—22 и К—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-один до Ка-пять" + breakTimeLong +
						"Ка-девять" + breakTimeLong +
						"От Ка-шестнадцать до Ка-девятнадцать" + breakTimeLong +
						"Ка-двадцать два" + breakTimeLong +
						"И - Ка-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 12:
				resp.Text = "От Л—2 до Л—11, от Л—15 до Л—19 и Л—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-два до Эль-одиннадцать" + breakTimeLong +
						"От Эль-пятнадцать до Эль-девятнадцать" + breakTimeLong +
						"И - Эль-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 11:
				resp.Text = "М—2, М—4 и от М—7 до М—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-два" + breakTimeLong +
						"Эм-четыре" + breakTimeLong +
						"И - от Эм-семь до Эм-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 10:
				resp.Text = "От Н—10 до Н—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-десять до Эн-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 9:
				resp.Text = "От О—11 до О—15 и О—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-одиннадцать до О-пятнадцать" + breakTimeLong +
						"И - О-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "П—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "Р—7, Р—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-семь, Эр-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "От С—6 до С—9, С—20 и С—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-шесть до Эс-девять" + breakTimeLong +
						"Эс-двадцать" + breakTimeLong +
						"И - Эс-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "От Т—7 до Т—12 и от Т—18 до Т—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-семь до Тэ-двенадцать" + breakTimeLong +
						"И - от Тэ-восемнадцать до Тэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "От У—8 до У—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От У-восемь до У-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От Ф—9 до Ф—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-девять до Эф-двадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От Х—10 до Х—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ха-десять до Ха-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ц—12 до Ц—18\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-двенадцать до Цэ-восемнадцать" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 14:
				resp.Text = "Б—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "В—2, В—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-два, Вэ-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Г—2, Г—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два, Гэ-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Д—5, Д—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-пять, Дэ-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Е—6, Е—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-шесть, Е-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "Ё—12, Ё—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двенадцать, Йо-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От Ж—14 до Ж—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Же-четырнадцать до Жэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "З—12, З—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двенадцать, Зэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Й—7, Й—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - семь, И-краткое - восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "К—6, К—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-шесть, Ка-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "М—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Н—2, Н—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-два, Эн-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От О—2 до О—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-два до О-четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "П—3\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-три" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel2:
		switch color {
		case Green:
			switch colorsAndQuantity[Green] {
			case 2:
				resp.Text = "Ж—7, З—7, Л—7 и М—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-семь" + breakTimeLong +
						"Зэ-семь" + breakTimeLong +
						"Эль-семь" + breakTimeLong +
						"И - Эм-семь" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—8, М—8\nВы закрасили все клетки зеленым цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-восемь, Эм-восемь" + breakTimeLong +
						"Вы закрасили все клетки зеленым цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы закрасили все клетки зеленым цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы закрасили все клетки зеленым цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Все клетки зеленым цветом уже раскрашены. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Все клетки зеленым цветом уже раскрашены. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 5:
				resp.Text = "Д—3, П—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-три, Пэ-три" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 4:
				resp.Text = "Д—4, Е—4, О—4 и П—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-четыре" + breakTimeLong +
						"Е-четыре" + breakTimeLong +
						"О-четыре" + breakTimeLong +
						"И - Пэ-четыре" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 3:
				resp.Text = "Д—5, П—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-пять, Пэ-пять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "От И—10 до К—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-десять до Ка-десять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Й—11\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - одиннадцать" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 2:
				resp.Text = "Ж—14, З—14, Л—14 и М—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четырнадцать" + breakTimeLong +
						"Зэ-четырнадцать" + breakTimeLong +
						"Эль-четырнадцать" + breakTimeLong +
						"И - Эм-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—15 до Л—15\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-пятнадцать до Эль-пятнадцать" + breakTimeLong +
						"Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Красный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Красный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case DarkGray:
			switch colorsAndQuantity[DarkGray] {
			case 16:
				resp.Text = "Б—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 15:
				resp.Text = "В—21, В—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать один, Вэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 14:
				resp.Text = "От Г—21 до Г—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двадцать один до Гэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 13:
				resp.Text = "От Д—7 до Д—9, Д—11 и Д—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-семь до Дэ-девять" + breakTimeLong +
						"Дэ-одиннадцать" + breakTimeLong +
						"И - Дэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 12:
				resp.Text = "От Е—7 до Е—9, Е—12 и Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-семь до Е-девять" + breakTimeLong +
						"Е-двенадцать" + breakTimeLong +
						"И - Е-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 11:
				resp.Text = "От Ё—6 до Ё—8, Ё—13 и от Ё—20 до Ё—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-шесть до Йо-восемь" + breakTimeLong +
						"Йо-тринадцать" + breakTimeLong +
						"И - от Йо-двадцать до Йо-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 10:
				resp.Text = "Ж—5, Ж—6, Ж—13 и от Ж—17 до Ж—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-пять" + breakTimeLong +
						"Жэ-шесть" + breakTimeLong +
						"Жэ-тринадцать" + breakTimeLong +
						"И - от Жэ-семнадцать до Жэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 9:
				resp.Text = "От З—4 до З—6 и от З—19 до З—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четыре до Зэ-шесть" + breakTimeLong +
						"И - от Зэ-девятнадцать до Зэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 8:
				resp.Text = "И—5 и от И—20 до И—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-пять" + breakTimeLong +
						"И - от И-двадцать до И-двадцать три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 7:
				resp.Text = "Й—4, Й—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - четыре, И-краткое - двадцать один" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 6:
				resp.Text = "К—5 и от К—20 до К—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-пять" + breakTimeLong +
						"И - от Ка-двадцать до Ка-двадцать три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 5:
				resp.Text = "От Л—4 до Л—6 и от Л—19 до Л—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-четыре до Эль-шесть" + breakTimeLong +
						"И - от Эль-девятнадцать до Эль-двадцать три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 4:
				resp.Text = "М—5, М—6, М—13 и от М—17 до М—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-пять" + breakTimeLong +
						"Эм-шесть" + breakTimeLong +
						"Эм-тринадцать" + breakTimeLong +
						"И - от Эм-семнадцать до Эм-девятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 3:
				resp.Text = "От Н—6 до Н—8, Н—13 и от Н—20 до Н—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-шесть до Эн-восемь" + breakTimeLong +
						"Эн-тринадцать" + breakTimeLong +
						"И - от Эн-двадцать до Эн-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 2:
				resp.Text = "От О—7 до О—9, О—12 и О—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-семь до О-девять" + breakTimeLong +
						"О-двенадцать" + breakTimeLong +
						"И - О-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 1:
				colorsAndQuantity[DarkGray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От П—7 до П—9 и П—11\nТемно-серый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-семь до Пэ-девять" + breakTimeLong +
						"И - Пэ-одиннадцать" + breakTimeLong +
						"Темно-серый цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Темно-серый цвет использован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Темно-серый цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже использовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже использовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 20:
				resp.Text = "А—22, А—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать два, А-двадцать три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 19:
				resp.Text = "Б—10, Б—21 и Б—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" + breakTimeLong +
						"Бэ-двадцать один" + breakTimeLong +
						"И - Бэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 18:
				resp.Text = "В—10, В—12, В—20 и В—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-десять" + breakTimeLong +
						"Вэ-двенадцать" + breakTimeLong +
						"Вэ-двадцать" + breakTimeLong +
						"И - Вэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "От Г—1 до Г—12, Г—20 и Г—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-один до Гэ-двенадцать" + breakTimeLong +
						"Гэ-двадцать" + breakTimeLong +
						"И - Гэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "Д—1, Д—2, Д—6, Д—10, Д—12, Д—20 и от Д—22 до Д—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-один" + breakTimeLong +
						"Дэ-два" + breakTimeLong +
						"Дэ-шесть" + breakTimeLong +
						"Дэ-десять" + breakTimeLong +
						"Дэ-двенадцать" + breakTimeLong +
						"Дэ-двадцать" + breakTimeLong +
						"И - От Дэ-двадцать два до Дэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "Е—2, Е—3, Е—5, Е—6, Е—10, Е—13, Е—20, Е—21 и Е—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-два" + breakTimeLong +
						"Е-три" + breakTimeLong +
						"Е-пять" + breakTimeLong +
						"Е-шесть" + breakTimeLong +
						"Е-десять" + breakTimeLong +
						"Е-тринадцать" + breakTimeLong +
						"Е-двадцать" + breakTimeLong +
						"Е-двадцать один" + breakTimeLong +
						"И - Е-двадцать четыре" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "От Ё—2 до Ё—5, Ё—10, Ё—14, от Ё—16 до Ё—19, Ё—24 и Ё—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-два до Йо-пять" + breakTimeLong +
						"Йо-десять" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"От Йо-шестнадцать до Йо-девятнадцать" + breakTimeLong +
						"Йо-двадцать четыре" + breakTimeLong +
						"И - Йо-двадцать пять" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "Ж—3, Ж—4, Ж—15, Ж—16, от Ж—20 до Ж—23 и Ж—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"Жэ-четыре" + breakTimeLong +
						"Жэ-пятнадцать" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"От Жэ-двадцать до Жэ-двадцать три" + breakTimeLong +
						"И - Жэ-двадцать пять" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "З—3, З—8 и З—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-три" + breakTimeLong +
						"Зэ-восемь" + breakTimeLong +
						"И - Зэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "И—3, И—4, И—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-три" + breakTimeLong +
						"И-четыре" + breakTimeLong +
						"И-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Й—3, Й—12 и от Й—22 до Й—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - три" + breakTimeLong +
						"И-краткое - двенадцать" + breakTimeLong +
						"И - от И-краткое - двадцать два до И-краткое - двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "К—3, К—4 и К—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-три" + breakTimeLong +
						"Ка-четыре" + breakTimeLong +
						"И - Ка-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "Л—3, Л—8 и Л—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-три" + breakTimeLong +
						"Эль-восемь" + breakTimeLong +
						"И - Эль-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "М—3, М—4, М—15, М—16, от М—20 до М—23 и М—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-три" + breakTimeLong +
						"Эм-четыре" + breakTimeLong +
						"Эм-пятнадцать" + breakTimeLong +
						"Эм-шестнадцать" + breakTimeLong +
						"От Эм-двадцать до Эм-двадцать три" + breakTimeLong +
						"И - Эм-двадцать пять" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "От Н—2 до Н—5, Н—10, Н—14, от Н—16 до Н—19, Н—24 и Н—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-два до Эн-пять" + breakTimeLong +
						"Эн-десять" + breakTimeLong +
						"Эн-четырнадцать" + breakTimeLong +
						"От Эн-шестнадцать до Эн-девятнадцать" + breakTimeLong +
						"Эн-двадцать четыре" + breakTimeLong +
						"И - Эн-двадцать пять" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "О—2, О—3, О—5, О—6, О—10, О—13, О—20, О—21 и О—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-два" + breakTimeLong +
						"О-три" + breakTimeLong +
						"О-пять" + breakTimeLong +
						"О-шесть" + breakTimeLong +
						"О-десять" + breakTimeLong +
						"О-тринадцать" + breakTimeLong +
						"О-двадцать" + breakTimeLong +
						"О-двадцать один" + breakTimeLong +
						"И - О-двадцать четыре" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "П—1, П—2, П—6, П—10, П—12 и от П—22 до П—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-один" + breakTimeLong +
						"Пэ-два" + breakTimeLong +
						"Пэ-шесть" + breakTimeLong +
						"Пэ-десять" + breakTimeLong +
						"Пэ-двенадцать" + breakTimeLong +
						"И - От Пэ-двадцать два до Пэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От Р—1 до Р—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эр-один до Эр-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "С—10, С—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-десять, Эс-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Т—10\nВы нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Тэ-десять" + breakTimeLong +
						"Вы нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Клетки уже закрашены черным цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Клетки уже закрашены черным цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel3:
		switch color {
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 13:
				resp.Text = "Д—1, Ё—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-один, Йо-один" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 12:
				resp.Text = "Г—2, Е—2 и Ж—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два" + breakTimeLong +
						"Е-два" + breakTimeLong +
						"И - Жэ-два" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 11:
				resp.Text = "Е—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 10:
				resp.Text = "От В—6 до З—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-шесть до Зэ-шесть" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 9:
				resp.Text = "Б—7, В—7, И—7, Й—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-семь" + breakTimeLong +
						"Вэ-семь" + breakTimeLong +
						"И-семь" + breakTimeLong +
						"И-краткое - семь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 8:
				resp.Text = "Б—8, К—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемь, Ка-восемь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 7:
				resp.Text = "Б—9, П—9 и С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-девять" + breakTimeLong +
						"Пэ-девять" + breakTimeLong +
						"И - Эс-девять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 6:
				resp.Text = "Б—10 и от П—10 до С—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" + breakTimeLong +
						"И - от Пэ-десять до Эс-десять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "Б—11 и от П—11 до С—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-одиннадцать" + breakTimeLong +
						"И - от Пэ-одиннадцать до Эс-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "Б—12, Р—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двенадцать, Эр-двенадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "П—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "З—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—15\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-пятнадцать" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Голубой цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Голубой цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 8:
				resp.Text = "От Г—7 до З—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-семь до Зэ-семь" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 7:
				resp.Text = "От В—8 до Й—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-восемь до И-краткое - восемь" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 6:
				resp.Text = "От В—9 до К—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-девять до Ка-девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "В—10, Г—10 и от Ё—10 до К—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-десять" + breakTimeLong +
						"Гэ-десять" + breakTimeLong +
						"И - от Йо-десять до Ка-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "В—11, Г—11 и от Ё—11 до К—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" + breakTimeLong +
						"Гэ-одиннадцать" + breakTimeLong +
						"И - от Йо-одиннадцать до Ка-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "От В—12 до Л—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-двенадцать до Эль-двенадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "От Ё—13 до М—13 и Р—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-тринадцать до Эм-тринадцать" + breakTimeLong +
						"И - Эр-тринадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—14 и от Л—14 до П—14\nСиний цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четырнадцать " + breakTimeLong +
						"И - от Эль-четырнадцать до Пэ-четырнадцать" + breakTimeLong +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 12:
				resp.Text = "От В—5 до З—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пять до Зэ-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Б—6, И—6, Й—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-шесть" + breakTimeLong +
						"И-шесть" + breakTimeLong +
						"И-краткое - шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "А—7, К—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-семь, Ка-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "А—8, Л—8, П—8 и С—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-восемь" + breakTimeLong +
						"Эль-восемь" + breakTimeLong +
						"Пэ-восемь" + breakTimeLong +
						"И - Эс-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "А—9, Л—9, О—9, Р—9 и Т—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-девять" + breakTimeLong +
						"Эль-девять" + breakTimeLong +
						"О-девять" + breakTimeLong +
						"Эр-девять" + breakTimeLong +
						"И - Тэ-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "А—10, Л—10, О—10 и Т—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-десять" + breakTimeLong +
						"Эль-десять" + breakTimeLong +
						"О-десять" + breakTimeLong +
						"И - Тэ-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "А—11, Д—11, Л—11, О—11 и Т—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-одиннадцать" + breakTimeLong +
						"Дэ-одиннадцать" + breakTimeLong +
						"Эль-одиннадцать" + breakTimeLong +
						"О-одиннадцать" + breakTimeLong +
						"И - Тэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "А—12, М—12, П—12 и С—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двенадцать" + breakTimeLong +
						"Эм-двенадцать" + breakTimeLong +
						"Пэ-двенадцать" + breakTimeLong +
						"И - Эс-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "А—13, Н—13, О—13 и С—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-тринадцать" + breakTimeLong +
						"Эн-тринадцать" + breakTimeLong +
						"О-тринадцать" + breakTimeLong +
						"И - Эс-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Б—14, Р—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четырнадцать, Эр-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От В—15 до П—15, кроме: З—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пятнадцать до Пэ-пятнадцать" + breakTimeMiddle +
						"Кроме: Зэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Д—16, Е—16, З—16, И—16\nРисование черного цвета завершено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-шестнадцать" + breakTimeLong +
						"Е-шестнадцать" + breakTimeLong +
						"Зэ-шестнадцать" + breakTimeLong +
						"И-шестнадцать" + breakTimeLong +
						"Рисование черного цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование черного цвета завершено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование черного цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel4:
		switch color {
		case Beige:
			switch colorsAndQuantity[Beige] {
			case 15:
				resp.Text = "От Б—7 до Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-семь до Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 14:
				resp.Text = "В—7, В—8, В—12 и от В—15 до В—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-семь" + breakTimeLong +
						"Вэ-восемь" + breakTimeLong +
						"Вэ-двенадцать" + breakTimeLong +
						"И - от Вэ-пятнадцать до Вэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 13:
				resp.Text = "Г—7, Г—13 и от Г—15 до Г—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-семь" + breakTimeLong +
						"Гэ-тринадцать" + breakTimeLong +
						"И - от Гэ-пятнадцать до Гэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 12:
				resp.Text = "Д—7, Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семь, Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 11:
				resp.Text = "Е—7, Е—13 и Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-семь" + breakTimeLong +
						"Е-тринадцать" + breakTimeLong +
						"И - Е-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 10:
				resp.Text = "Ё—8, Ё—12 и от Ё—17 до Ё—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-восемь" + breakTimeLong +
						"Йо-двенадцать" + breakTimeLong +
						"И - от Йо-семнадцать до Йо-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 9:
				resp.Text = "От Ж—9 до Ж—11 и от Ж—16 до Ж—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девять до Жэ-одиннадцать" + breakTimeLong +
						"И - от Жэ-шестнадцать до Жэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 8:
				resp.Text = "От З—16 до З—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-шестнадцать до Зэ-двадцать два" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 7:
				resp.Text = "От И—9 до И—11 и от И—16 до И—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-девять до И-одиннадцать" + breakTimeLong +
						"И - от И-шестнадцать до И-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 6:
				resp.Text = "Й—8, Й—12 и от Й—17 до Й—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - восемь" + breakTimeLong +
						"И-краткое - двенадцать" + breakTimeLong +
						"И - от И-краткое - семнадцать до И-краткое - двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 5:
				resp.Text = "К—7, К—13 и К—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-семь" + breakTimeLong +
						"Ка-тринадцать" + breakTimeLong +
						"И - Ка-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 4:
				resp.Text = "Л—7, Л—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-семь, Эль-тринадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 3:
				resp.Text = "М—7, М—13 и от М—15 до М—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-семь" + breakTimeLong +
						"Эм-тринадцать" + breakTimeLong +
						"И - от Эм-пятнадцать до Эм-двадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 2:
				resp.Text = "Н—7, Н—8, Н—12 и от Н—15 до Н—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-семь" + breakTimeLong +
						"Эн-восемь" + breakTimeLong +
						"Эн-двенадцать" + breakTimeLong +
						"И - от Эн-пятнадцать до Эн-двадцать один" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 1:
				colorsAndQuantity[Beige]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От О—7 до О—12\nБежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-семь до О-двенадцать" + breakTimeLong +
						"Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Бежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Бежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 5:
				resp.Text = "От Ж—12 до И—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двенадцать до И-двенадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "От Ж—13 до И—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-тринадцать до И-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "З—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "Ё—23, Й—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать три, И-краткое - двадцать три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Е—24 до Ж—24 и от И—24 до К—24\nВы дорисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двадцать четыре до Жэ-двадцать четыре" + breakTimeLong +
						"И - от И-двадцать четыре до Ка-двадцать четыре" + breakTimeLong +
						"Вы дорисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы дорисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы дорисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Желтый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Желтый цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 17:
				resp.Text = "От А—14 до А—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-четырнадцать до А-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 16:
				resp.Text = "От Б—1 до Б—4, Б—6 и от Б—13 до Б—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-один до Бэ-четыре" + breakTimeLong +
						"Бэ-шесть" + breakTimeLong +
						"И - от Бэ-тринадцать до Бэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 15:
				resp.Text = "От В—2 до В—6, В—13, В—14, В—22 и В—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-два до Вэ-шесть" + breakTimeLong +
						"Вэ-тринадцать" + breakTimeLong +
						"Вэ-четырнадцать" + breakTimeLong +
						"Вэ-двадцать два" + breakTimeLong +
						"И - Вэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 14:
				resp.Text = "От Г—3 до Г—6, Г—14 и от Г—21 до Г—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-три до Гэ-шесть" + breakTimeLong +
						"Гэ-четырнадцать" + breakTimeLong +
						"И - от Гэ-двадцать один до Гэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 13:
				resp.Text = "От Д—4 до Д—6 и от Д—14 до Д—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-четыре до Дэ-шесть" + breakTimeLong +
						"И - от Дэ-четырнадцать до Дэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 12:
				resp.Text = "От Е—4 до Е—6 и от Е—14 до Е—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-шесть" + breakTimeLong +
						"И - от Е-четырнадцать до Е-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 11:
				resp.Text = "От Ё—4 до Ё—7 и от Ё—13 до Ё—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четыре до Йо-семь" + breakTimeLong +
						"И - от Йо-тринадцать до Йо-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 10:
				resp.Text = "От Ж—4 до Ж—8, Ж—14 и Ж—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-четыре до Жэ-восемь" + breakTimeLong +
						"Жэ-четырнадцать" + breakTimeLong +
						"И - Жэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 9:
				resp.Text = "От З—4 до З—11 и З—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четыре до Зэ-одиннадцать" + breakTimeLong +
						"И - Зэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "От И—4 до И—8, И—14, И—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-четыре до И-восемь" + breakTimeLong +
						"И-четырнадцать" + breakTimeLong +
						"И-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "От Й—4 до Й—7 и от Й—13 до Й—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - четыре до И-краткое - семь" + breakTimeLong +
						"И - от И-краткое - тринадцать до И-краткое - шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "От К—4 до К—6 и от К—14 до К—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-четыре до Ка-шесть" + breakTimeLong +
						"И - от Ка-четырнадцать до Ка-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "От Л—4 до Л—6 и от Л—14 до Л—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-четыре до Эль-шесть" + breakTimeLong +
						"И - от Эль-четырнадцать до Эль-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "От М—3 до М—6, М—14 и от М—21 до М—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-три до Эм-шесть" + breakTimeLong +
						"Эм-четырнадцать" + breakTimeLong +
						"И - от Эм-двадцать один до Эм-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От Н—2 до Н—6, Н—13, Н—14, Н—22 и Н—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-два до Эн-шесть" + breakTimeLong +
						"Эн-тринадцать" + breakTimeLong +
						"Эн-четырнадцать" + breakTimeLong +
						"Эн-двадцать два" + breakTimeLong +
						"И - Эн-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От О—1 до О—4, О—6 и от О—13 до О—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-один до О-четыре" + breakTimeLong +
						"О-шесть" + breakTimeLong +
						"И - от О-тринадцать до О-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От П—14 до П—21\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-четырнадцать до Пэ-двадцать один" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 5:
				resp.Text = "От Г—8 до Е—8 и от К—8 до М—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-восемь до Е-восемь" + breakTimeLong +
						"И - от Ка-восемь до Эм-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "В—9, Ё—9, Й—9 и Н—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-девять" + breakTimeLong +
						"Йо-девять" + breakTimeLong +
						"И-краткое - девять" + breakTimeLong +
						"И - Эн-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "В—10, Д—10, Ё—10, Й—10, Л—10 и Н—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-десять" + breakTimeLong +
						"Дэ-десять" + breakTimeLong +
						"Йо-десять" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"Эль-десять" + breakTimeLong +
						"И - Эн-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "В—11, Ё—11, Й—11 и Н—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" + breakTimeLong +
						"Йо-одиннадцать" + breakTimeLong +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И - Эн-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Г—12 до Е—12 и от К—12 до М—12\nРисование черного цвета завершено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двенадцать до Е-двенадцать" + breakTimeLong +
						"И - от Ка-двенадцать до Эм-двенадцать" + breakTimeLong +
						"Рисование черного цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование черного цвета завершено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование черного цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel5:
		switch color {
		case DarkGray:
			switch colorsAndQuantity[DarkGray] {
			case 8:
				resp.Text = "Л—2, М—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-два, Эм-два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 7:
				resp.Text = "К—3, Н—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-три, Эн-три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 6:
				resp.Text = "В—4, Й—4 и О—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четыре" + breakTimeLong +
						"И-краткое - четыре" + breakTimeLong +
						"И - О-четыре" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 5:
				resp.Text = "От В—5 до Д—5, Й—5 и О—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пять до Дэ-пять" + breakTimeLong +
						"И-краткое - пять" + breakTimeLong +
						"И - О-пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 4:
				resp.Text = "От В—6 до Д—6, К—6 и Н—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-шесть до Дэ-шесть" + breakTimeLong +
						"Ка-шесть" + breakTimeLong +
						"И - Эн-шесть" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 3:
				resp.Text = "От В—7 до Д—7, Л—7 и М—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-семь до Дэ-семь" + breakTimeLong +
						"Эль-семь" + breakTimeLong +
						"И - Эм-семь" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 2:
				resp.Text = "В—8, Е—8, Ё—8, И—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-восемь" + breakTimeLong +
						"Е-восемь" + breakTimeLong +
						"Йо-восемь" + breakTimeLong +
						"И-восемь" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 1:
				colorsAndQuantity[DarkGray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—9\nВы нарисовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девять" + breakTimeLong +
						"Вы нарисовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Темно-серый цвет Вы уже нарисовали. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Темно-серый цвет Вы уже нарисовали. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 10:
				resp.Text = "От Л—1 до С—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-один до Эс-один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 9:
				resp.Text = "К—2 и от Н—2 до Т—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-два" + breakTimeLong +
						"И - от Эн-два до Тэ-два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "Й—3 и от О—3 до Т—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - три" + breakTimeLong +
						"И - от О-три до Тэ-три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "Й—6 и от О—6 до Т—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - шесть" + breakTimeLong +
						"И - от О-шесть до Тэ-шесть" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "Й—7, К—7 и от Н—7 до Т—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - семь" + breakTimeLong +
						"Ка-семь" + breakTimeLong +
						"И - от Эн-семь до Тэ-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "От Й—8 до Н—8 и от П—8 до Т—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - восемь до Эн-восемь" + breakTimeLong +
						"И - от Пэ-восемь до Тэ-восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "Й—9, К—9 и от О—9 до Р—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - девять" + breakTimeLong +
						"Ка-девять" + breakTimeLong +
						"И - от О-девять до Эр-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От К—10 до О—10, С—10 и Т—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-десять до О-десять" + breakTimeLong +
						"Эс-десять" + breakTimeLong +
						"И - Тэ-десять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "Й—11 и от П—11 до Т—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И - от Пэ-одиннадцать до Тэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—12 до Т—12\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-двенадцать до Тэ-двенадцать" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 2:
				resp.Text = "А—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шесть" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Б—5 до Б—7\nРисование красного цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-пять до Бэ-семь" + breakTimeLong +
						"Рисование красного цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование красного цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование красного цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали красный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали красный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 6:
				resp.Text = "С—9, Т—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-девять, Тэ-девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "Й—10, П—10 и Р—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - десять" + breakTimeLong +
						"Пэ-десять" + breakTimeLong +
						"И - Эр-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "От К—11 до О—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-одиннадцать до О-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "От Й—13 до Т—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - тринадцать до Тэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "От Й—14 до Т—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - четырнадцать до Тэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От К—15 до С—15\nСиний цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-пятнадцать до Эс-пятнадцать" + breakTimeLong +
						"Синий цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет использован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали синий цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали синий цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ё—9 до Ё—11\nКоричневый цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девять до Йо-одиннадцать" + breakTimeLong +
						"Коричневый цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет закончен. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 11:
				resp.Text = "Л—4, М—4 и от П—4 до Т—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-четыре" + breakTimeLong +
						"Эм-четыре" + breakTimeLong +
						"И - от Пэ-четыре до Тэ-четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "От Е—5 до И—5, Л—5, М—5 и от П—5 до Т—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пять до И-пять" + breakTimeLong +
						"Эль-пять" + breakTimeLong +
						"Эм-пять" + breakTimeLong +
						"И - от Пэ-пять до Тэ-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "От Е—6 до И—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-шесть до И-шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От Е—7 до И—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-семь до И-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "О—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "З—9, И—9 и от Л—9 до Н—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-девять" + breakTimeLong +
						"И-девять" + breakTimeLong +
						"И - от Эль-девять до Эн-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "И—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Ж—11, З—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-одиннадцать, Зэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Ж—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Л—16, М—16, П—16 и Р—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-шестнадцать" + breakTimeLong +
						"Эм-шестнадцать" + breakTimeLong +
						"Пэ-шестнадцать" + breakTimeLong +
						"И - Эр-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От К—17 до М—17 и от О—17 до Р—17\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-семнадцать до Эм-семнадцать" + breakTimeLong +
						"И - от О-семнадцать до Эр-семнадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel6:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 13:
				resp.Text = "В—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 12:
				resp.Text = "Г—9, Г—10, Г—17 и Г—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-девять" + breakTimeLong +
						"Гэ-десять" + breakTimeLong +
						"Гэ-семнадцать" + breakTimeLong +
						"И - Гэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 11:
				resp.Text = "От Д—4 до Д—7, от Д—10 до Д—12 и Д—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-четыре до Дэ-семь" + breakTimeLong +
						"От Дэ-десять до Дэ-двенадцать" + breakTimeLong +
						"И - Дэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 10:
				resp.Text = "От Е—4 до Е—7, от Е—10 до Е—13, Е—21 и Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-семь" + breakTimeLong +
						"От Е-десять до Е-тринадцать" + breakTimeLong +
						"Е-двадцать один" + breakTimeLong +
						"И - Е-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 9:
				resp.Text = "Ё—6, Ё—7, от Ё—10 до Ё—12, Ё—14 и Ё—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шесть" + breakTimeLong +
						"Йо-семь" + breakTimeLong +
						"От Йо-десять до Йо-двенадцать" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"И - Йо-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "Ж—7, от Ж—10 до Ж—12 и Ж—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-семь" + breakTimeLong +
						"От Жэ-десять до Жэ-двенадцать" + breakTimeLong +
						"И - Жэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "З—7, от З—10 до З—12 и З—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-семь" + breakTimeLong +
						"От Зэ-десять до Зэ-двенадцать" + breakTimeLong +
						"И - Зэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "И—7, от И—10 до И—12, И—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-семь" + breakTimeLong +
						"От И-десять до И-двенадцать" + breakTimeLong +
						"И-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "Й—6, Й—7, от Й—10 до Й—12, Й—14, Й—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - шесть" + breakTimeLong +
						"И-краткое - семь" + breakTimeLong +
						"От И-краткое - десять до И-краткое - двенадцать" + breakTimeLong +
						"И-краткое - четырнадцать" + breakTimeLong +
						"И-краткое - двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "От К—4 до К—7, от К—10 до К—13, К—21 и К—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-четыре до Ка-семь" + breakTimeLong +
						"От Ка-десять до Ка-тринадцать" + breakTimeLong +
						"Ка-двадцать один" + breakTimeLong +
						"И - Ка-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От Л—4 до Л—7, от Л—10 до Л—12 и Л—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-четыре до Эль-семь" + breakTimeLong +
						"От Эль-десять до Эль-двенадцать" + breakTimeLong +
						"И - Эль-двадцать один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "М—9, М—10, М—17 и М—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-девять" + breakTimeLong +
						"Эм-десять" + breakTimeLong +
						"Эм-семнадцать" + breakTimeLong +
						"И - Эм-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Н—17\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-семнадцать" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 15:
				resp.Text = "Б—18, Б—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемнадцать, Бэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 14:
				resp.Text = "От В—7 до В—10, В—18 и В—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-семь до Вэ-десять" + breakTimeLong +
						"Вэ-восемнадцать" + breakTimeLong +
						"И - Вэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 13:
				resp.Text = "От Г—4 до Г—6, Г—11, Г—12, Г—16 и Г—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-четыре до Гэ-шесть" + breakTimeLong +
						"Гэ-одиннадцать" + breakTimeLong +
						"Гэ-двенадцать" + breakTimeLong +
						"Гэ-шестнадцать" + breakTimeLong +
						"И - Гэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 12:
				resp.Text = "Д—3, Д—13, Д—15 и от Д—22 до Д—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-три" + breakTimeLong +
						"Дэ-тринадцать" + breakTimeLong +
						"Дэ-пятнадцать" + breakTimeLong +
						"И - от Дэ-двадцать два до Дэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 11:
				resp.Text = "Е—3, от Е—16 до Е—20, Е—23 и Е—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" + breakTimeLong +
						"От Е-шестнадцать до Е-двадцать" + breakTimeLong +
						"Е-двадцать три" + breakTimeLong +
						"И - Е-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 10:
				resp.Text = "От Ё—3 до Ё—5, Ё—15 и от Ё—18 до Ё—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-три до Йо-пять" + breakTimeLong +
						"Йо-пятнадцать" + breakTimeLong +
						"И - от Йо-восемнадцать до Йо-двадцать один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 9:
				resp.Text = "От Ж—2 до Ж—6, Ж—15 и от Ж—19 до Ж—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-два до Жэ-шесть" + breakTimeLong +
						"Жэ-пятнадцать" + breakTimeLong +
						"И - от Жэ-девятнадцать до Жэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 8:
				resp.Text = "От З—2 до З—6, З—15, З—20 и З—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-два до Зэ-шесть" + breakTimeLong +
						"Зэ-пятнадцать" + breakTimeLong +
						"Зэ-двадцать" + breakTimeLong +
						"И - Зэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "От И—2 до И—6, И—15 и от И—19 до И—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-два до И-шесть" + breakTimeLong +
						"И-пятнадцать" + breakTimeLong +
						"И - от И-девятнадцать до И-двадцать один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "От Й—3 до Й—5, Й—15 и от Й—18 до Й—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - три до И-краткое - пять" + breakTimeLong +
						"И-краткое - пятнадцать" + breakTimeLong +
						"И - от И-краткое - восемнадцать до И-краткое - двадцать один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "К—3, от К—16 до К—20, К—23 и К—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-три" + breakTimeLong +
						"От Ка-шестнадцать до Ка-двадцать" + breakTimeLong +
						"Ка-двадцать три" + breakTimeLong +
						"И - Ка-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "Л—3, Л—13, Л—15 и от Л—22 до Л—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-три" + breakTimeLong +
						"Эль-тринадцать" + breakTimeLong +
						"Эль-пятнадцать" + breakTimeLong +
						"И - от Эль-двадцать два до Эль-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "От М—4 до М—6, М—11, М—12, М—16 и М—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-четыре до Эм-шесть" + breakTimeLong +
						"Эм-одиннадцать" + breakTimeLong +
						"Эм-двенадцать" + breakTimeLong +
						"Эм-шестнадцать" + breakTimeLong +
						"И - Эм-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "От Н—7 до Н—10, Н—18 и Н—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-семь до Эн-десять" + breakTimeLong +
						"Эн-восемнадцать" + breakTimeLong +
						"И - Эн-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "О—18, О—19\nРисование красного цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О—восемнадцать, О-девятнадцать" + breakTimeLong +
						"Рисование красного цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование красного цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование красного цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали красный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали красный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 17:
				resp.Text = "А—18, А—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-восемнадцать, А-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "От Б—7 до Б—10, Б—17 и Б—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-семь до Бэ-десять" + breakTimeLong +
						"Бэ-семнадцать" + breakTimeLong +
						"И - Бэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "От В—4 до В—6, В—11, В—12, В—16, В—20 и от В—23 до В—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Вэ-шесть" + breakTimeLong +
						"Вэ-одиннадцать" + breakTimeLong +
						"Вэ-двенадцать" + breakTimeLong +
						"Вэ-шестнадцать" + breakTimeLong +
						"Вэ-двадцать" + breakTimeLong +
						"И - от Вэ-двадцать три до Вэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Г—3, Г—7, Г—8, Г—13, Г—15, Г—19, от Г—21 до Г—23 и Г—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-три" + breakTimeLong +
						"Гэ-семь" + breakTimeLong +
						"Гэ-восемь" + breakTimeLong +
						"Гэ-тринадцать" + breakTimeLong +
						"Гэ-пятнадцать" + breakTimeLong +
						"Гэ-девятнадцать" + breakTimeLong +
						"От Гэ-двадцать один до Гэ-двадцать три" + breakTimeLong +
						"И - Гэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "Д—2, Д—8, Д—14, от Д—16 до Д—20 и Д—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-два" + breakTimeLong +
						"Дэ-восемь" + breakTimeLong +
						"Дэ-четырнадцать" + breakTimeLong +
						"От Дэ-шестнадцать до Дэ-двадцать" + breakTimeLong +
						"И - Дэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Е—2, Е—8, Е—14, Е—15 и Е—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-два" + breakTimeLong +
						"Е-восемь" + breakTimeLong +
						"Е-четырнадцать" + breakTimeLong +
						"Е-пятнадцать" + breakTimeLong +
						"И - Е-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Ё—2, Ё—8, Ё—13, Ё—16, Ё—17 и от Ё—23 до Ё—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-два" + breakTimeLong +
						"Йо-восемь" + breakTimeLong +
						"Йо-тринадцать" + breakTimeLong +
						"Йо-шестнадцать" + breakTimeLong +
						"Йо-семнадцать" + breakTimeLong +
						"И - от Йо-двадцать три до Йо-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Ж—1, Ж—8, Ж—9, Ж—13, Ж—16, Ж—18 и Ж—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-один" + breakTimeLong +
						"Жэ-восемь" + breakTimeLong +
						"Жэ-девять" + breakTimeLong +
						"Жэ-тринадцать" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"Жэ-восемнадцать" + breakTimeLong +
						"И - Жэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "З—1, З—8, З—9, З—13, З—16, З—19 и З—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-один" + breakTimeLong +
						"Зэ-восемь" + breakTimeLong +
						"Зэ-девять" + breakTimeLong +
						"Зэ-тринадцать" + breakTimeLong +
						"Зэ-шестнадцать" + breakTimeLong +
						"Зэ-девятнадцать" + breakTimeLong +
						"И - Зэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "И—1, И—8, И—9, И—13, И—16, И—18, И—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-один" + breakTimeLong +
						"И-восемь" + breakTimeLong +
						"И-девять" + breakTimeLong +
						"И-тринадцать" + breakTimeLong +
						"И-шестнадцать" + breakTimeLong +
						"И-восемнадцать" + breakTimeLong +
						"И-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "Й—2, Й—8, Й—13, Й—16, Й—17 и от Й—23 до Й—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - два" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"И-краткое - тринадцать" + breakTimeLong +
						"И-краткое - шестнадцать" + breakTimeLong +
						"И-краткое - семнадцать" + breakTimeLong +
						"И - от И-краткое - двадцать три до И-краткое - двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "К—2, К—8, К—14, К—15 и К—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-два" + breakTimeLong +
						"Ка-восемь" + breakTimeLong +
						"Ка-четырнадцать" + breakTimeLong +
						"Ка-пятнадцать" + breakTimeLong +
						"И - Ка-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Л—2, Л—8, Л—14, от Л—16 до Л—20 и Л—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-два" + breakTimeLong +
						"Эль-восемь" + breakTimeLong +
						"Эль-четырнадцать" + breakTimeLong +
						"От Эль-шестнадцать до Эль-двадцать" + breakTimeLong +
						"И - Эль-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "М—3, М—7, М—8, М—13, М—15, М—19, от М—21 до М—23 и М—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-три" + breakTimeLong +
						"Эм-семь" + breakTimeLong +
						"Эм-восемь" + breakTimeLong +
						"Эм-тринадцать" + breakTimeLong +
						"Эм-пятнадцать" + breakTimeLong +
						"Эм-девятнадцать" + breakTimeLong +
						"От Эм-двадцать один до Эм-двадцать три" + breakTimeLong +
						"И - Эм-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От Н—4 до Н—6, Н—11, Н—12, Н—16, Н—20 и от Н—23 до Н—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-четыре до Эн-шесть" + breakTimeLong +
						"Эн-одиннадцать" + breakTimeLong +
						"Эн-двенадцать" + breakTimeLong +
						"Эн-шестнадцать" + breakTimeLong +
						"Эн-двадцать" + breakTimeLong +
						"И - от Эн-двадцать три до Эн-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От О—7 до О—10, О—17 и О—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-семь до О-десять" + breakTimeLong +
						"О-семнадцать" + breakTimeLong +
						"И - О-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "П—18, П—19\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-восемнадцать, Пэ-девятнадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel7:
		switch color {
		case LightGray:
			switch colorsAndQuantity[LightGray] {
			case 14:
				resp.Text = "От В—2 до Д—2, от Ё—2 до Л—2 и от Н—2 до П—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-два до Дэ-два" + breakTimeLong +
						"От Йо-два до Эль-два" + breakTimeLong +
						"И - от Эн-два до Пэ-два" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 13:
				resp.Text = "От Б—3 до Г—3, от Е—3 до М—3 и от О—3 до Р—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-три до Гэ-три" + breakTimeLong +
						"От Е-три до Эм-три" + breakTimeLong +
						"И - от О-три до Эр-три" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 12:
				resp.Text = "Б—4, В—4, от Д—4 до Н—4, П—4 и Р—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четыре" + breakTimeLong +
						"Вэ-четыре" + breakTimeLong +
						"От Дэ-четыре до Эн-четыре" + breakTimeLong +
						"Пэ-четыре" + breakTimeLong +
						"И - Эр-четыре" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 11:
				resp.Text = "Б—5, В—5, от Д—5 до З—5, от Й—5 до Н—5, П—5 и Р—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-пять" + breakTimeLong +
						"Вэ-пять" + breakTimeLong +
						"От Дэ-пять до Зэ-пять" + breakTimeLong +
						"От И-краткое - пять до Эн-пять" + breakTimeLong +
						"Пэ-пять" + breakTimeLong +
						"И - Эр-пять" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 10:
				resp.Text = "От В—6 до Д—6, Ё—6, Ж—6, К—6, Л—6, Н—6 и О—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-шесть до Дэ-шесть" + breakTimeLong +
						"Йо-шесть" + breakTimeLong +
						"Жэ-шесть" + breakTimeLong +
						"Ка-шесть" + breakTimeLong +
						"Эль-шесть" + breakTimeLong +
						"Эн-шесть" + breakTimeLong +
						"И - О-шесть" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 9:
				resp.Text = "Г—7, Д—7, Ё—7, Ж—7, К—7, Л—7, Н—7 и О—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-семь" + breakTimeLong +
						"Дэ-семь" + breakTimeLong +
						"Йо-семь" + breakTimeLong +
						"Жэ-семь" + breakTimeLong +
						"Ка-семь" + breakTimeLong +
						"Эль-семь" + breakTimeLong +
						"Эн-семь" + breakTimeLong +
						"И - О-семь" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 8:
				resp.Text = "От Д—8 до Ж—8 и от К—8 до О—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-восемь до Жэ-восемь" + breakTimeLong +
						"И - от Ка-восемь до О-восемь" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 7:
				resp.Text = "От Е—9 до О—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-девять до О-девять" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 6:
				resp.Text = "От Ё—10 до Н—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-десять до Эн-десять" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 5:
				resp.Text = "От Е—12 до К—12 и М—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двенадцать до Ка-двенадцать" + breakTimeLong +
						"И - Эм-двенадцать" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 4:
				resp.Text = "От Ё—13 до Й—13 и М—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-тринадцать до И-краткое - тринадцать" + breakTimeLong +
						"И - Эм-тринадцать" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 3:
				resp.Text = "Ё—14, К—14 и М—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-четырнадцать" + breakTimeLong +
						"Ка-четырнадцать" + breakTimeLong +
						"И - Эм-четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 2:
				resp.Text = "От З—15 до К—15 и М—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-пятнадцать до Ка-пятнадцать" + breakTimeLong +
						"И - Эм-пятнадцать" +
						"</speak>"
				colorsAndQuantity[LightGray]--
			case 1:
				colorsAndQuantity[LightGray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—16 до Л—16\nСветло-серый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-шестнадцать до Эль-шестнадцать" + breakTimeLong +
						"Светло-серый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-серый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-серый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали светло-серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали светло-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case DarkGray:
			switch colorsAndQuantity[DarkGray] {
			case 17:
				resp.Text = "От В—1 до Д—1, от Ё—1 до Л—1 и от Н—1 до П—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-один до Дэ-один" + breakTimeLong +
						"От Йо-один до Эль-один" + breakTimeLong +
						"И - от Эн-один до Пэ-один" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 16:
				resp.Text = "Б—2, Е—2, М—2 и Р—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-два" + breakTimeLong +
						"Е-два" + breakTimeLong +
						"Эм-два" + breakTimeLong +
						"И - Эр-два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 15:
				resp.Text = "А—3, Д—3, Н—3 и С—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-три" + breakTimeLong +
						"Дэ-три" + breakTimeLong +
						"Эн-три" + breakTimeLong +
						"И - Эс-три" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 14:
				resp.Text = "А—4, Г—4, О—4 и С—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-четыре" + breakTimeLong +
						"Гэ-четыре" + breakTimeLong +
						"О-четыре" + breakTimeLong +
						"И - Эс-четыре" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 13:
				resp.Text = "А—5, Г—5, О—5 и С—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пять" + breakTimeLong +
						"Гэ-пять" + breakTimeLong +
						"О-пять" + breakTimeLong +
						"И - Эс-пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 12:
				resp.Text = "П—6, Р—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-шесть, Эр-шесть" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 11:
				resp.Text = "П—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-семь" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 10:
				resp.Text = "П—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-восемь" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 9:
				resp.Text = "П—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-девять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 8:
				resp.Text = "О—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-десять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 7:
				resp.Text = "От Е—11 до Н—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-одиннадцать до Эн-одиннадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 6:
				resp.Text = "Л—12, Н—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-двенадцать, Эн-двенадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 5:
				resp.Text = "К—13, Л—13 и Н—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-тринадцать" + breakTimeLong +
						"Эль-тринадцать" + breakTimeLong +
						"И - Эн-тринадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 4:
				resp.Text = "От Ж—14 до Й—14, Л—14 и Н—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-четырнадцать до И-краткое - четырнадцать" + breakTimeLong +
						"Эль-четырнадцать" + breakTimeLong +
						"И - Эн-четырнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 3:
				resp.Text = "Ж—15, Л—15 и Н—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-пятнадцать" + breakTimeLong +
						"Эль-пятнадцать" + breakTimeLong +
						"И - Эн-пятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 2:
				resp.Text = "Ж—16, М—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-шестнадцать, Эм-шестнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 1:
				colorsAndQuantity[DarkGray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—17 до Л—17\nРисование темно-серого цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-семнадцать до Эль-семнадцать" + breakTimeLong +
						"Рисование темно-серого цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование темно-серого цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование темно-серого цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 8:
				resp.Text = "А—6, А—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шесть, А-семь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "От Б—7 до Б—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-семь до Бэ-девять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "В—7, В—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-семь, Вэ-девять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "Г—8 и от Г—10 до Г—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-восемь" + breakTimeLong +
						"И - от Гэ-десять до Гэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "От Д—12 до Д—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-двенадцать до Дэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От Е—13 до Е—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-тринадцать до Е-семнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "От Ё—16 до Ё—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-шестнадцать до Йо-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "И—18\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-восемнадцать" + breakTimeLong +
						"Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали оранжевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали оранжевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Green:
			switch colorsAndQuantity[Green] {
			case 5:
				resp.Text = "Б—14, Б—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четырнадцать, Бэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "В—14, В—15 и В—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четырнадцать" + breakTimeLong +
						"Вэ-пятнадцать" + breakTimeLong +
						"И - Вэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "Г—9, Г—16 и Г—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-девять" + breakTimeLong +
						"Гэ-шестнадцать" + breakTimeLong +
						"И - Гэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "Д—9, Д—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять, Дэ-десять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—10\nЗеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-десять" + breakTimeLong +
						"Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Зеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали зеленый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали зеленый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 9:
				resp.Text = "А—8, А—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-восемь, А-девять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "Б—6, Б—10 и Б—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-шесть" + breakTimeLong +
						"Бэ-десять" + breakTimeLong +
						"И - Бэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "В—8 и от В—10 до В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-восемь" + breakTimeLong +
						"И - от Вэ-десять до Вэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "Г—14, Г—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четырнадцать, Гэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "Д—11 и от Д—16 до Д—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-одиннадцать" + breakTimeLong +
						"И - от Дэ-шестнадцать до Дэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "Е—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е—восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "Ё—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "Ж—17, Ж—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-семнадцать, Жэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—18\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-восемнадцать" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 4:
				resp.Text = "И—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Е—6, от З—6 до Й—6 и М—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е—шесть" + breakTimeLong +
						"От Зэ-шесть до И-краткое - шесть" + breakTimeLong +
						"И - Эм-шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Е—7, от З—7 до Й—7 и М—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е—семь" + breakTimeLong +
						"От Зэ-семь до И-краткое - семь" + breakTimeLong +
						"И - Эм-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—8 до Й—8\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-восемь до И-краткое - восемь" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel8:
		switch color {
		case DarkGray:
			switch colorsAndQuantity[DarkGray] {
			case 6:
				resp.Text = "А—5, Б—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пять, Бэ-пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 5:
				resp.Text = "А—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шесть" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 4:
				resp.Text = "От Г—14 до Е—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-четырнадцать до Е-четырнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 3:
				resp.Text = "В—15, Г—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-пятнадцать, Гэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 2:
				resp.Text = "От Е—16 до З—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-шестнадцать до Зэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 1:
				colorsAndQuantity[DarkGray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Д—17, Е—17 и Ж—17\nРисование темно-серого цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семнадцать" + breakTimeLong +
						"Е-семнадцать" + breakTimeLong +
						"И - Жэ-семнадцать" + breakTimeLong +
						"Рисование темно-серого цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование темно-серого цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование темно-серого цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 13:
				resp.Text = "От Б—1 до Д—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-один до Дэ-один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 12:
				resp.Text = "А—2, Б—2, Д—2 и Е—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-два" + breakTimeLong +
						"Бэ-два" + breakTimeLong +
						"Дэ-два" + breakTimeLong +
						"И - Е-два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 11:
				resp.Text = "А—3, Б—3 и от Д—3 до Ё—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-три" + breakTimeLong +
						"Бэ-три" + breakTimeLong +
						"И - от Дэ-три до Йо-три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 10:
				resp.Text = "От В—4 до Ё—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Йо-четыре" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 9:
				resp.Text = "От В—5 до Д—5, Ё—5 и Ж—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пять до Дэ-пять" + breakTimeLong +
						"Йо-пять" + breakTimeLong +
						"И - Жэ-пять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "От Б—6 до Е—6 и З—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-шесть до Е-шесть" + breakTimeLong +
						"И - Зэ-шесть" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "От В—7 до Д—7, Ё—7 и Ж—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-семь до Дэ-семь" + breakTimeLong +
						"Йо-семь" + breakTimeLong +
						"И - Жэ-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "В—8, Г—8, Е—8 и от З—8 до Й—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-восемь" + breakTimeLong +
						"Гэ-восемь" + breakTimeLong +
						"Е-восемь" + breakTimeLong +
						"И - от Зэ-восемь до И-краткое - восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "Ё—9, Ж—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-девять, Жэ-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "Ё—10 и от З—10 до Й—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-десять" + breakTimeLong +
						"И - от Зэ-десять до И-краткое - десять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "Ж—11, К—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-одиннадцать, Ка-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От З—12 до Й—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-двенадцать до И-краткое - двенадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "К—13, Л—13\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-тринадцать, Эль-тринадцать" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Green:
			switch colorsAndQuantity[Green] {
			case 15:
				resp.Text = "Д—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-восемь" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 14:
				resp.Text = "От В—9 до Е—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-девять до Е-девять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 13:
				resp.Text = "От В—10 до Е—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-десять до Е-десять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 12:
				resp.Text = "От Г—11 до Ё—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-одиннадцать до Йо-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 11:
				resp.Text = "От Д—12 до Ж—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-двенадцать до Жэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 10:
				resp.Text = "От Д—13 до Й—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-тринадцать до И-краткое - тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 9:
				resp.Text = "От Ё—14 до Й—14, Л—14 и М—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четырнадцать до И-краткое - четырнадцать" + breakTimeLong +
						"Эль-четырнадцать" + breakTimeLong +
						"И - Эм-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 8:
				resp.Text = "Ж—15, З—15, Й—15, К—15, М—15 и Н—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-пятнадцать" + breakTimeLong +
						"Зэ-пятнадцать" + breakTimeLong +
						"И-краткое - пятнадцать" + breakTimeLong +
						"Ка-пятнадцать" + breakTimeLong +
						"Эм-пятнадцать" + breakTimeLong +
						"И - Эн-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 7:
				resp.Text = "К—16, Л—16 и О—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-шестнадцать" + breakTimeLong +
						"Эль-шестнадцать" + breakTimeLong +
						"И - О-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 6:
				resp.Text = "От Л—17 до Н—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-семнадцать до Эн-семнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 5:
				resp.Text = "Н—18, О—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-восемнадцать, О-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "О—19, П—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-девятнадцать, Пэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "П—20, Р—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-двадцать, Эр-двадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "Р—21, С—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-двадцать один, Эс-двадцать один" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "С—22\nЗеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-двадцать два" + breakTimeLong +
						"Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Зеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали зеленый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали зеленый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 6:
				resp.Text = "А—4, Б—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-четыре, Бэ-четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "П—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "Р—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "С—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-двадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "Т—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Тэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Т—22\nРисование синего цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Тэ-двадцать два" + breakTimeLong +
						"Рисование синего цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование синего цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование синего цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали синий цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали синий цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 9:
				resp.Text = "А—13, Б—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-тринадцать, Бэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "От А—14 до В—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А—четырнадцать до Вэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "Д—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "Г—16, Д—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-шестнадцать, Дэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "Ё—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-семнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "От Е—18 до Ж—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е—восемнадцать до Жэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От Ё—19 до И—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девятнадцать до И-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От Ж—20 до Й—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двадцать до И-краткое - двадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—21 до Й—21\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-двадцать один до И-краткое - двадцать один" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 13:
				resp.Text = "В—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Е—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Ё—6, Ж—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шесть, Жэ-шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Е—7, З—7, И—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е—семь" + breakTimeLong +
						"Зэ-семь" + breakTimeLong +
						"И-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "Ё—8, Ж—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-восемь, Жэ-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От З—9 до Й—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-девять до И-краткое - девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "Ж—10, К—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-десять, Ка-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "От З—11 до Й—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-одиннадцать до И-краткое - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "К—12, Л—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-двенадцать, Эль-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "К—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Л—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "М—16, Н—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-шестнадцать, Эн-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "О—17, П—17\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-семнадцать, Пэ-семнадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel9:
		switch color {
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 8:
				resp.Text = "А—2, А—3 и от А—5 до А—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-два" + breakTimeLong +
						"А-три" + breakTimeLong +
						"И - от А-пять до А-восемь" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 7:
				resp.Text = "От Б—2 до Б—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-два до Бэ-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 6:
				resp.Text = "От В—4 до В—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Вэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 5:
				resp.Text = "От Г—6 до Г—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-шесть до Гэ-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 4:
				resp.Text = "От Ф—6 до Ф—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-шесть до Эф-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 3:
				resp.Text = "От Х—4 до Х—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ха-четыре до Ха-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "От Ц—2 до Ц—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-два до Цэ-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ч—2 до Ч—5, Ч—7 и Ч—8\nРисование светло-розового цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Чэ-два до Чэ-пять" + breakTimeLong +
						"Чэ-семь" + breakTimeLong +
						"И - Чэ-восемь" + breakTimeLong +
						"Рисование светло-розового цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование светло-розового цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование светло-розового цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали светло-розовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали светло-розовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 15:
				resp.Text = "Ё—8, С—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-восемь, Эс-восемь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 14:
				resp.Text = "Е—9, Ж—9, Р—9 и Т—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девять" + breakTimeLong +
						"Жэ-девять" + breakTimeLong +
						"Эр-девять" + breakTimeLong +
						"И - Тэ-девять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 13:
				resp.Text = "Д—10, З—10, П—10 и У—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-десять" + breakTimeLong +
						"Зэ-десять" + breakTimeLong +
						"Пэ-десять" + breakTimeLong +
						"И - У-десять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 12:
				resp.Text = "Д—11, И—11, О—11 и У—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-одиннадцать" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"О-одиннадцать" + breakTimeLong +
						"И - У-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 11:
				resp.Text = "Д—12, И—12, О—12 и У—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двенадцать" + breakTimeLong +
						"И-двенадцать" + breakTimeLong +
						"О-двенадцать" + breakTimeLong +
						"И - У-двенадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 10:
				resp.Text = "Д—13, И—13, О—13 и У—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-тринадцать" + breakTimeLong +
						"И-тринадцать" + breakTimeLong +
						"О-тринадцать" + breakTimeLong +
						"И - У-тринадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 9:
				resp.Text = "Е—14, З—14, П—14 и Т—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четырнадцать" + breakTimeLong +
						"Зэ-четырнадцать" + breakTimeLong +
						"Пэ-четырнадцать" + breakTimeLong +
						"И - Тэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 8:
				resp.Text = "Ё—15, Ж—15, Р—15 и С—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-пятнадцать" + breakTimeLong +
						"Жэ-пятнадцать" + breakTimeLong +
						"Эр-пятнадцать" + breakTimeLong +
						"И - Эс-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 7:
				resp.Text = "От Ё—17 до С—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-семнадцать до Эс-семнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 6:
				resp.Text = "От К—18 до М—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-восемнадцать до Эм-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "Л—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "Л—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-двадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "Л—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-двадцать один" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "От К—22 до М—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать два до Эм-двадцать два" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От К—23 до М—23\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать три до Эм-двадцать три" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали голубой цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали голубой цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 23:
				resp.Text = "Б—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 22:
				resp.Text = "В—2, В—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-два, Вэ-три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 21:
				resp.Text = "Г—4, Г—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четыре, Гэ-пять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 20:
				resp.Text = "От Д—6 до Д—8 и от Д—14 до Д—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-шесть до Дэ-восемь" + breakTimeLong +
						"И - от Дэ-четырнадцать до Дэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 19:
				resp.Text = "От Е—15 до Е—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пятнадцать до Е-семнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 18:
				resp.Text = "Ё—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 17:
				resp.Text = "Ж—7, Ж—8, Ж—16, Ж—21 и Ж—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-семь" + breakTimeLong +
						"Жэ-восемь" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"Жэ-двадцать один" + breakTimeLong +
						"И - Жэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 16:
				resp.Text = "От З—7 до З—9, З—15, З—16, З—21 и З—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-семь до Зэ-девять" + breakTimeLong +
						"Зэ-пятнадцать" + breakTimeLong +
						"Зэ-шестнадцать" + breakTimeLong +
						"Зэ-двадцать один" + breakTimeLong +
						"И - Зэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 15:
				resp.Text = "От И—6 до И—10, от И—14 до И—16, И—19 и от И—21 до И—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-шесть до И-десять" + breakTimeLong +
						"От И-четырнадцать до И-шестнадцать" + breakTimeLong +
						"И-девятнадцать" + breakTimeLong +
						"И - от И-двадцать один до И-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 14:
				resp.Text = "От Й—6 до Й—13, Й—15, Й—16 и от Й—18 до Й—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - шесть до И-краткое - тринадцать" + breakTimeLong +
						"И-краткое - пятнадцать" + breakTimeLong +
						"И-краткое - шестнадцать" + breakTimeLong +
						"И - от И-краткое - восемнадцать до И-краткое - двадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 13:
				resp.Text = "От К—5 до К—12 и К—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-пять до Ка-двенадцать" + breakTimeLong +
						"И - Ка-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 12:
				resp.Text = "Л—4, от Л—6 до Л—12 и Л—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-четыре" + breakTimeLong +
						"От - Эль-шесть до Эль-двенадцать" + breakTimeLong +
						"И - Эль-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 11:
				resp.Text = "От М—5 до М—12 и М—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-пять до Эм-двенадцать" + breakTimeLong +
						"И - Эм-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 10:
				resp.Text = "Н—4, от Н—6 до Н—13, Н—15, Н—16 и от Н—18 до Н—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-четыре" + breakTimeLong +
						"От Эн-шесть до Эн-тринадцать" + breakTimeLong +
						"Эн-пятнадцать" + breakTimeLong +
						"Эн-шестнадцать" + breakTimeLong +
						"И - от Эн-восемнадцать до Эн-двадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 9:
				resp.Text = "От О—6 до О—10, от О—14 до О—16, О—19 и от О—21 до О—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-шесть до О-десять" + breakTimeLong +
						"От О-четырнадцать до О-шестнадцать" + breakTimeLong +
						"О-девятнадцать" + breakTimeLong +
						"И - от О-двадцать один до О-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 8:
				resp.Text = "От П—7 до П—9, П—15, П—16, П—21 и П—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-семь до Пэ-девять" + breakTimeLong +
						"Пэ-пятнадцать" + breakTimeLong +
						"Пэ-шестнадцать" + breakTimeLong +
						"Пэ-двадцать один" + breakTimeLong +
						"И - Пэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 7:
				resp.Text = "Р—7, Р—8, Р—16, Р—21 и Р—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-семь" + breakTimeLong +
						"Эр-восемь" + breakTimeLong +
						"Эр-шестнадцать" + breakTimeLong +
						"Эр-двадцать один" + breakTimeLong +
						"И - Эр-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 6:
				resp.Text = "С—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "От Т—15 до Т—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-пятнадцать до Тэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "От У—6 до У—8 и от У—14 до У—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От У-шесть до У-восемь" + breakTimeLong +
						"И - от У-четырнадцать до У-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "Ф—4, Ф—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эф-четыре, Эф-пять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "Х—2, Х—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ха-два, Ха-три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ц—1\nРисование синего цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Цэ-один" + breakTimeLong +
						"Рисование синего цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование синего цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование синего цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали синий цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали синий цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 23:
				resp.Text = "Б—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 22:
				resp.Text = "В—1, В—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-один, Вэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 21:
				resp.Text = "Г—2, Г—3 и от Г—10 до Г—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два" + breakTimeLong +
						"Гэ-три" + breakTimeLong +
						"И - от Гэ-десять до Гэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 20:
				resp.Text = "Д—4, Д—5 и Д—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-четыре" + breakTimeLong +
						"Дэ-пять" + breakTimeLong +
						"И - Дэ-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 19:
				resp.Text = "Е—8 и от Е—10 до Е—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемь" + breakTimeLong +
						"И - от Е—десять до Е—тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 18:
				resp.Text = "Ё—7, от Ё—9 до Ё—12, Ё—14 и от Ё—21 до Ё—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-семь" + breakTimeLong +
						"От Йо-девять до Йо-двенадцать" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"И - от Йо-двадцать один до Йо-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "Ж—6, Ж—10, от Ж—12 до Ж—14, Ж—20, Ж—22, Ж—23 и Ж—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-шесть" + breakTimeLong +
						"Жэ-десять" + breakTimeLong +
						"От Жэ-двенадцать до Жэ-четырнадцать" + breakTimeLong +
						"Жэ-двадцать" + breakTimeLong +
						"Жэ-двадцать два" + breakTimeLong +
						"Жэ-двадцать три" + breakTimeLong +
						"И - Жэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "З—6, от З—11 до З—13, З—19, З—20, З—22, З—23 и З—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-шесть" + breakTimeLong +
						"От Зэ-одиннадцать до Зэ-тринадцать" + breakTimeLong +
						"Зэ-девятнадцать" + breakTimeLong +
						"Зэ-двадцать" + breakTimeLong +
						"Зэ-двадцать два" + breakTimeLong +
						"Зэ-двадцать три" + breakTimeLong +
						"И - Зэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "И—5, И—18, И—20, И—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-пять" + breakTimeLong +
						"И-восемнадцать" + breakTimeLong +
						"И-двадцать" + breakTimeLong +
						"И-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Й—5, Й—14 и от Й—21 до Й—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - пять" + breakTimeLong +
						"И-краткое - четырнадцать" + breakTimeLong +
						"И - от И-краткое - двадцать один до И-краткое - двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "К—4, от К—13 до К—15, от К—19 до К—21 и К—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-четыре" + breakTimeLong +
						"От Ка-тринадцать до Ка-пятнадцать" + breakTimeLong +
						"От Ка-девятнадцать до Ка-двадцать один" + breakTimeLong +
						"И - Ка-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Л—3, Л—5, от Л—13 до Л—15 и Л—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-три" + breakTimeLong +
						"Эль-пять" + breakTimeLong +
						"От Эль-тринадцать до Эль-пятнадцать" + breakTimeLong +
						"И - Эль-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "М—4, от М—13 до М—15, от М—19 до М—21 и М—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-четыре" + breakTimeLong +
						"От Эм-тринадцать до Эм-пятнадцать" + breakTimeLong +
						"От Эм-девятнадцать до Эм-двадцать один" + breakTimeLong +
						"И - Эм-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Н—3, Н—5, Н—14 и от Н—21 до Н—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-три" + breakTimeLong +
						"Эн-пять" + breakTimeLong +
						"Эн-четырнадцать" + breakTimeLong +
						"И - от Эн-двадцать один до Эн-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "От О—3 до О—5, О—18, О—20 и О—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-три до О-пять" + breakTimeLong +
						"О-восемнадцать" + breakTimeLong +
						"О-двадцать" + breakTimeLong +
						"И - О-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "П—6, от П—11 до П—13, П—19, П—20, П—22, П—23 и П—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-шесть" + breakTimeLong +
						"От Пэ-одиннадцать до Пэ-тринадцать" + breakTimeLong +
						"Пэ-девятнадцать" + breakTimeLong +
						"Пэ-двадцать" + breakTimeLong +
						"Пэ-двадцать два" + breakTimeLong +
						"Пэ-двадцать три" + breakTimeLong +
						"И - Пэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "Р—6, Р—10, от Р—12 до Р—14, Р—20, Р—22, Р—23 и Р—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-шесть" + breakTimeLong +
						"Эр-десять" + breakTimeLong +
						"От Эр-двенадцать до Эр-четырнадцать" + breakTimeLong +
						"Эр-двадцать" + breakTimeLong +
						"Эр-двадцать два" + breakTimeLong +
						"Эр-двадцать три" + breakTimeLong +
						"И - Эр-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "С—7, от С—9 до С—12, С—14 и от С—21 до С—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-семь" + breakTimeLong +
						"От Эс-девять до Эс-двенадцать" + breakTimeLong +
						"Эс-четырнадцать" + breakTimeLong +
						"И - от Эс-двадцать один до Эс-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Т—8 и от Т—10 до Т—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Тэ-восемь" + breakTimeLong +
						"И - от Тэ—десять до Тэ—тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "У—4, У—5 и У—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"У-четыре" + breakTimeLong +
						"У-пять" + breakTimeLong +
						"И - У-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Ф—2, Ф—3 и от Ф—10 до Ф—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эф-два" + breakTimeLong +
						"Эф-три" + breakTimeLong +
						"И - от Эф-десять до Эф-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Х—1, Х—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ха-один, Ха-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ц—10\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Цэ-десять" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDMediumLevel10:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 5:
				resp.Text = "От З—12 до Й—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-двенадцать до И-краткое - двенадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "Г—21, Д—21, Н—21 и О—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-двадцать один" + breakTimeLong +
						"Дэ-двадцать один" + breakTimeLong +
						"Эн-двадцать один" + breakTimeLong +
						"И - О-двадцать один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От Г—22 до Е—22 и от М—22 до О—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двадцать два до Е-двадцать два" + breakTimeLong +
						"И - от Эм-двадцать два до О-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От В—23 до Е—23 и от М—23 до П—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-двадцать три до Е-двадцать три" + breakTimeLong +
						"И - от Эм-двадцать три до Пэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Г—24 до Е—24 и от М—24 до О—24\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двадцать четыре до Е-двадцать четыре" + breakTimeLong +
						"И - от Эм-двадцать четыре до О-двадцать четыре" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 9:
				resp.Text = "От З—11 до Й—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-одиннадцать до И-краткое - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 8:
				resp.Text = "Ж—12, К—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двенадцать, Ка-двенадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "От З—13 до Й—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-тринадцать до И-краткое - тринадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "От В—20 до Е—20 и от М—20 до П—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-двадцать до Е-двадцать" + breakTimeLong +
						"И - от Эм-двадцать до Пэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "В—21, Е—21, Ё—21, Л—21, М—21 и П—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать один" + breakTimeLong +
						"Е-двадцать один" + breakTimeLong +
						"Йо-двадцать один" + breakTimeLong +
						"Эль-двадцать один" + breakTimeLong +
						"Эм-двадцать один" + breakTimeLong +
						"И - Пэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "Б—22, В—22, Ё—22, Л—22, П—22 и Р—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать два" + breakTimeLong +
						"Вэ-двадцать два" + breakTimeLong +
						"Йо-двадцать два" + breakTimeLong +
						"Эль-двадцать два" + breakTimeLong +
						"Пэ-двадцать два" + breakTimeLong +
						"И - Эр-двадцать два" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "Б—23, Ё—23, Л—23 и Р—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать три" + breakTimeLong +
						"Йо-двадцать три" + breakTimeLong +
						"Эль-двадцать три" + breakTimeLong +
						"И - Эр-двадцать три" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "В—24, Ё—24, Л—24 и П—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать четыре" + breakTimeLong +
						"Йо-двадцать четыре" + breakTimeLong +
						"Эль-двадцать четыре" + breakTimeLong +
						"И - Пэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Г—25 до Е—25 и от М—25 до О—25\nРисование оранжевого цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двадцать пять до Е-двадцать пять" + breakTimeLong +
						"И - от Эм-двадцать пять до О-двадцать пять" + breakTimeLong +
						"Рисование оранжевого цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование оранжевого цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование оранжевого цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали оранжевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали оранжевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 19:
				resp.Text = "От А—19 до А—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-девятнадцать до А-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 18:
				resp.Text = "От Б—6 до Б—12 и от Б—18 до Б—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-шесть до Бэ-двенадцать" + breakTimeLong +
						"И - от Бэ-восемнадцать до Бэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "От В—5 до В—13 и от В—16 до В—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пять до Вэ-тринадцать" + breakTimeLong +
						"И - от Вэ-шестнадцать до Вэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "Г—4, Г—5 и от Г—12 до Г—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четыре" + breakTimeLong +
						"Гэ-пять" + breakTimeLong +
						"И - от Гэ-двенадцать до Гэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "Д—3, Д—4 и от Д—13 до Д—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-три" + breakTimeLong +
						"Дэ-четыре" + breakTimeLong +
						"И - от Дэ-тринадцать до Дэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Е—2, Е—3 и от Е—14 до Е—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-два" + breakTimeLong +
						"Е-три" + breakTimeLong +
						"И - от Е-четырнадцать до Е-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "От Ё—1 до Ё—3, от Ё—7 до Ё—9 и от Ё—14 до Ё—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-один до Йо-три" + breakTimeLong +
						"От Йо-семь до Йо-девять" + breakTimeLong +
						"И - от Йо-четырнадцать до Йо-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "От Ж—1 до Ж—3, от Ж—7 до Ж—9, Ж—15 и Ж—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-один до Жэ-три" + breakTimeLong +
						"От Жэ-семь до Жэ-девять" + breakTimeLong +
						"Жэ-пятнадцать" + breakTimeLong +
						"И - Жэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "От З—1 до З—4, З—15 и З—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-один до Зэ-четыре" + breakTimeLong +
						"Зэ-пятнадцать" + breakTimeLong +
						"И - Зэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "От И—1 до И—6, И—15, И—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-один до И-шесть" + breakTimeLong +
						"И-пятнадцать" + breakTimeLong +
						"И-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "От Й—1 до Й—4, Й—15, Й—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - один до И-краткое - четыре" + breakTimeLong +
						"И-краткое - пятнадцать" + breakTimeLong +
						"И-краткое - двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От К—1 до К—3, от К—7 до К—9, К—15 и К—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-один до Ка-три" + breakTimeLong +
						"От Ка-семь до Ка-девять" + breakTimeLong +
						"Ка-пятнадцать" + breakTimeLong +
						"И - Ка-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "От Л—1 до Л—3, от Л—7 до Л—9 и от Л—14 до Л—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-один до Эль-три" + breakTimeLong +
						"От Эль-семь до Эль-девять" + breakTimeLong +
						"И - от Эль-четырнадцать до Эль-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "М—2, М—3 и от М—14 до М—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-два" + breakTimeLong +
						"Эм-три" + breakTimeLong +
						"И - от Эм-четырнадцать до Эм-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Н—3, Н—4 и от Н—13 до Н—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-три" + breakTimeLong +
						"Эн-четыре" + breakTimeLong +
						"И - от Эн-тринадцать до ЭН-семнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "О—4, О—5 и от О—12 до О—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-четыре" + breakTimeLong +
						"О-пять" + breakTimeLong +
						"И - от О-двенадцать до О-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От П—5 до П—13 и от П—16 до П—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-пять до Пэ-тринадцать" + breakTimeLong +
						"И - от Пэ-шестнадцать до Пэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От Р—6 до Р—12 и от Р—18 до Р—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эр-шесть до Эр-двенадцать" + breakTimeLong +
						"И - от Эр-восемнадцать до Эр-двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От С—19 до С—22\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-девятнадцать до Эс-двадцать два" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	}
	return  resp
}

func hardLevels(resp Response) Response {
	switch pictureID {
	case pictureIDHardLevel1:
		switch color {
		case Beige:
			switch colorsAndQuantity[Beige] {
			case 17:
				resp.Text = "А—17, А—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-семнадцать, А-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 16:
				resp.Text = "От Б—16 до Б—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-шестнадцать до Бэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 15:
				resp.Text = "От В—15 до В—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пятнадцать до Вэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 14:
				resp.Text = "От Г—7 до Г—9, Г—15, Г—16 и Г—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-семь до Гэ-девять" + breakTimeLong +
						"Гэ-пятнадцать" + breakTimeLong +
						"Гэ-шестнадцать" + breakTimeLong +
						"И - Гэ-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 13:
				resp.Text = "Д—6, Д—7, Д—10, Д—13, Д—14 и от Д—24 до Д—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-шесть" + breakTimeLong +
						"Дэ-семь" + breakTimeLong +
						"Дэ-десять" + breakTimeLong +
						"Дэ-тринадцать" + breakTimeLong +
						"Дэ-четырнадцать" + breakTimeLong +
						"И - от Дэ-двадцать четыре до Дэ-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 12:
				resp.Text = "Е—6, Е—10, Е—11, Е—13, Е—14 и от Е—25 до Е—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-шесть" + breakTimeLong +
						"Е-десять" + breakTimeLong +
						"Е-одиннадцать" + breakTimeLong +
						"Е-тринадцать" + breakTimeLong +
						"Е-четырнадцать" + breakTimeLong +
						"И - от Е-двадцать пять до Е-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 11:
				resp.Text = "Ё—6, от Ё—9 до Ё—11, Ё—14, Ё—18 и от Ё—26 до Ё—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шесть" + breakTimeLong +
						"от Йо-девять до Йо-одиннадцать" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"Йо-восемнадцать" + breakTimeLong +
						"И - от Йо-двадцать шесть до Йо-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 10:
				resp.Text = "От Ж—6 до Ж—13 и Ж—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-шесть до Жэ-тринадцать" + breakTimeLong +
						"И - Жэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 9:
				resp.Text = "От З—6 до З—13 и З—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-шесть до Зэ-тринадцать" + breakTimeLong +
						"И - Зэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 8:
				resp.Text = "От И—5 до И—13, И—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-пять до И-тринадцать" + breakTimeLong +
						"И-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 7:
				resp.Text = "От Й—4 до Й—6, от Й—9 до Й—11, Й—18 и от Й—26 до Й—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - четыре до И-краткое - шесть" + breakTimeLong +
						"От И-краткое - девять до И-краткое - одиннадцать" + breakTimeLong +
						"И-краткое - восемнадцать" + breakTimeLong +
						"И - от И-краткое - двадцать шесть до И-краткое - двадцать девять" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 6:
				resp.Text = "К—5, К—6, К—10, К—11 и от К—25 до К—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-пять" + breakTimeLong +
						"Ка-шесть" + breakTimeLong +
						"Ка-десять" + breakTimeLong +
						"Ка-одиннадцать" +  breakTimeLong +
						"И - от Ка-двадцать пять до Ка-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 5:
				resp.Text = "Л—6, Л—7, Л—10, Л—15 и от Л—24 до Л—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-шесть" + breakTimeLong +
						"Эль-семь" + breakTimeLong +
						"Эль-десять" + breakTimeLong +
						"Эль-пятнадцать" +  breakTimeLong +
						"И - от Эль-двадцать четыре до Эль-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 4:
				resp.Text = "От М—6 до М—9, от М—14 до М—16, М—29 и М—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-шесть до Эм-девять" + breakTimeLong +
						"От Эм-четырнадцать до Эм-шестнадцать" + breakTimeLong +
						"Эм-двадцать девять" + breakTimeLong +
						"И - Эм-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 3:
				resp.Text = "От Н—15 до Н—17 и Н—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-пятнадцать до Эн-семнадцать" + breakTimeLong +
						"И - Эн-тридцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 2:
				resp.Text = "От О—16 до О—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-шестнадцать до О-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 1:
				colorsAndQuantity[Beige]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "П—17, П—18\nВы нарисовали бежевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-семнадцать, Пэ-восемнадцать" + breakTimeLong +
						"Вы нарисовали бежевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали бежевый цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали бежевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Бежевый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бежевый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 3:
				resp.Text = "К—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "И—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—16, Ж—16\nРисование желтого цвета завершено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-шестнадцать, Жэ-шестнадцать" + breakTimeLong +
						"Рисование желтого цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование желтого цвета завершено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование желтого цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 2:
				resp.Text = "Д—8, Л—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-восемь, Эль-восемь" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Д—9, Е—9, К—9 и Л—9\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять" + breakTimeLong +
						"Е-девять" + breakTimeLong +
						"Ка-девять" + breakTimeLong +
						"И - Эль-девять" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже был нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже был нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 8:
				resp.Text = "К—13, Л—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-тринадцать, Эль-тринадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "Г—14, Й—14 и Л—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четырнадцать" + breakTimeLong +
						"И-краткое - четырнадцать" + breakTimeLong +
						"И - Эль-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "От Д—15 до К—15, кроме: З—15, И—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пятнадцать до Ка-пятнадцать" + breakTimeMiddle +
						"Кроме: Зэ-пятнадцать, И-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "От Ё—16 до К—16, кроме: Ж—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо—шестнадцать до Ка-шестнадцать" + breakTimeMiddle +
						"Кроме: Жэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "От Ё—17 до Й—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-семнадцать до И-краткое - семнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "Ё—19, Й—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-девятнадцать, И-краткое - девятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "Ж—20, И—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать, И-двадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—21\nПоздравляю! Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двадцать один" + breakTimeLong +
						"Поздравляю! Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Поздравляю! Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Поздравляю! Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали красный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали красный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—14, З—15\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четырнадцать, Зэ-пятнадцать" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали голубой цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали голубой цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 2:
				resp.Text = "Ё—13, Й—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-тринадцать, И-краткое - тринадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—14, И—14\nСиний цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четырнадцать, И-четырнадцать" + breakTimeLong +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет уже был нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет уже был нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 9:
				resp.Text = "От Д—21 до Д—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-двадцать один до Дэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "От Е—20 до Е—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двадцать до Е-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "От Ё—20 до Ё—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать до Йо-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "Ж—19 и от Ж—21 до Ж—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-девятнадцать" + breakTimeLong +
						"И - от Жэ-двадцать один до Жэ-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "З—19, З—20 и от З—22 до З—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-девятнадцать" + breakTimeLong +
						"Зэ-двадцать" + breakTimeLong +
						"И - от Зэ-двадцать два до Зэ-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "И—19 и от И—21 до И—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-девятнадцать" + breakTimeLong +
						"И - от И-двадцать один до И-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От Й—20 до Й—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - двадцать до И-краткое - двадцать пять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От К—20 до К—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать до Ка-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Л—21 до Л—23\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-двадцать один до Эль-двадцать три" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 22:
				resp.Text = "От Е—1 до Й—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-один до И-краткое - один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 21:
				resp.Text = "От Г—2 до М—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-два до Эм-два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 20:
				resp.Text = "От В—3 до Н—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-три до Эн-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 19:
				resp.Text = "От Б—4 до О—4, кроме: Й—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-четыре до О-четыре" + breakTimeMiddle +
						"Кроме: И-краткое - четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 18:
				resp.Text = "От А—5 до З—5 и от Л—5 до О—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-пять до Зэ-пять" + breakTimeLong +
						"И - от Эль-пять до О-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "От А—6 до Г—6, Н—6 и О—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-шесть до Гэ-шесть" + breakTimeLong +
						"Эн-шесть" + breakTimeLong +
						"И - О-шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "От А—7 до В—7, Е—7, К—7, Н—7 и О—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-семь до Вэ-семь" + breakTimeLong +
						"Е-семь" + breakTimeLong +
						"Ка-семь" + breakTimeLong +
						"Эн-семь" + breakTimeLong +
						"И - О-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "От А—8 до В—8, Е—8, Ё—8, Й—8, К—8, Н—8 и О—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-восемь до Вэ-восемь" + breakTimeLong +
						"Е-восемь" + breakTimeLong +
						"Йо-восемь" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"Ка-восемь" + breakTimeLong +
						"Эн-восемь" + breakTimeLong +
						"И - О-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "От А—9 до В—9, Н—9 и О—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-девять до Вэ-девять" + breakTimeLong +
						"Эн-девять" + breakTimeLong +
						"И - О-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "От А—10 до Г—10 и от М—10 до П—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-десять до Гэ-десять" + breakTimeLong +
						"И - от Эм-десять до Пэ-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "От Б—11 до Д—11 и от Л—11 до Р—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-одиннадцать до Дэ-одиннадцать" + breakTimeLong +
						"И - от Эль-одиннадцать до Эр-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "От Б—12 до Ё—12, от Й—12 до С—12 и от Ц—12 до Ш—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-двенадцать до Йо-двенадцать" + breakTimeLong +
						"От И-краткое - двенадцать до Эс-двенадцать" + breakTimeLong +
						"И - от Цэ-двенадцать до Ша-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "От Б—13 до Г—13, от М—13 до Т—13, Х—13 и Ц—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-тринадцать до Гэ-тринадцать" + breakTimeLong +
						"От Эм-тринадцать до Тэ-тринадцать" + breakTimeLong +
						"Ха-тринадцать" + breakTimeLong +
						"И - Цэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "В—14 и от Н—14 до Ц—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четырнадцать" + breakTimeLong +
						"И - от Эн-четырнадцать до Цэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От О—15 до Ф—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-пятнадцать до Эф-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "Д—16, Л—16 и от П—16 до С—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-шестнадцать" + breakTimeLong +
						"Эль-шестнадцать" + breakTimeLong +
						"И - от Пэ-шестнадцать до Эс-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Д—17, Е—17, от К—17 до М—17, от Р—17 до Т—17, Ф—17 и Х—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семнадцать" + breakTimeLong +
						"Е-семнадцать" + breakTimeLong +
						"От Ка-семнадцать до Эм-семнадцать" + breakTimeLong +
						"От Эр-семнадцать до Тэ-семнадцать" + breakTimeLong +
						"Эф-семнадцать" + breakTimeLong +
						"И - Ха-семнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Е—18, от К—18 до Н—18, от Р—18 до Ф—18 и Ц—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемнадцать" + breakTimeLong +
						"От Ка-восемнадцать до Эн-восемнадцать" + breakTimeLong +
						"От Эр-восемнадцать до Эф-восемнадцать" + breakTimeLong +
						"И - Цэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "От К—19 до У—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-девятнадцать до У-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От Л—20 до С—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-двадцать до Эс-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От М—21 до Т—21, кроме: С—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-двадцать один до Тэ-двадцать один" + breakTimeMiddle +
						"Кроме: Эс-двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—22, О—22 и С—22\nРисование черного цвета закончено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-двадцать два" + breakTimeLong +
						"О-двадцать два" + breakTimeLong +
						"И - Эс-двадцать два" + breakTimeLong +
						"Рисование черного цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование черного цвета закончено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование черного цвета закончено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel2:
		switch color {
		case Beige:
			switch colorsAndQuantity[Beige] {
			case 12:
				resp.Text = "От Й—6 до Н—6, кроме: К—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - шесть до Эн-шесть" + breakTimeMiddle +
						"Кроме: Ка-шесть" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 11:
				resp.Text = "От И—7 до Р—7, кроме: Н—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-семь до Эр-семь" + breakTimeMiddle +
						"Кроме: Эн-семь" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 10:
				resp.Text = "З—8, И—8, от Л—8 до Н—8, Р—8 и С—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-восемь" + breakTimeLong +
						"И-восемь" + breakTimeLong +
						"От Эль-восемь до Эн-восемь" + breakTimeLong +
						"Эр-восемь" + breakTimeLong +
						"И - Эс-восемь" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 9:
				resp.Text = "Ж—9, З—9, от Л—9 до Н—9 и С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-девять" + breakTimeLong +
						"Зэ-девять" + breakTimeLong +
						"От Эль-девять до Эн-девять" + breakTimeLong +
						"И - Эс-девять" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 8:
				resp.Text = "З—10, от К—10 до О—10 и С—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-десять" + breakTimeLong +
						"От Ка-десять до О-десять" + breakTimeLong +
						"И - Эс-десять" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 7:
				resp.Text = "От И—11 до Р—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-одиннадцать до Эр-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 6:
				resp.Text = "От Й—12 до П—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - двенадцать до Пэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 5:
				resp.Text = "От Л—13 до Н—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-тринадцать до Эн-тринадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 4:
				resp.Text = "От И—14 до Р—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-четырнадцать до Эр-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 3:
				resp.Text = "От Й—15 до П—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - пятнадцать до Пэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 2:
				resp.Text = "От Ё—22 до И—22 и от Р—22 до У—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать два до И-двадцать два" + breakTimeLong +
						"И - от Эр-двадцать два до У-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 1:
				colorsAndQuantity[Beige]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ж—23, З—23, С—23 и Т—23\nБежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать три" + breakTimeLong +
						"Зэ-двадцать три" + breakTimeLong +
						"Эс-двадцать три" + breakTimeLong +
						"И - Тэ-двадцать три" + breakTimeLong +
						"Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Бежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Бежевый цвет Вы уже рисовали. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бежевый цвет Вы уже рисовали. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 7:
				resp.Text = "З—14, С—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четырнадцать, Эс-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "От Ж—15 до И—15 и от Р—15 до Т—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-пятнадцать до И-пятнадцать" + breakTimeLong +
						"И - от Эр-пятнадцать до Тэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "Ё—16, Ж—16, от Й—16 до П—16, Т—16 и У—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шестнадцать" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"От Йо-шестнадцать до Пэ-шестнадцать" + breakTimeLong +
						"Тэ-шестнадцать" + breakTimeLong +
						"И - У-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "М—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-семнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "Ё—19, Ж—19, Т—19 и У—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-девятнадцать" + breakTimeLong +
						"Жэ-девятнадцать" + breakTimeLong +
						"Тэ-девятнадцать" + breakTimeLong +
						"И - У-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От Е—20 до Ж—20 и от Т—20 до Ф—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двадцать до Жэ-двадцать" + breakTimeLong +
						"И - от Тэ-двадцать до Эф-двадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—21, Ф—21\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать один, Эф-двадцать один" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 19:
				resp.Text = "Й—1, М—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - один, Эм-один" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 18:
				resp.Text = "От З—2 до Н—2 и П—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-два до Эн-два" + breakTimeLong +
						"И - Пэ-два" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 17:
				resp.Text = "От Ж—3 до С—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-три до Эс-три" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 16:
				resp.Text = "От Ё—4 до Т—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четыре до Тэ-четыре" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 15:
				resp.Text = "От Е—5 до Т—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пять до Тэ-пять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 14:
				resp.Text = "От Е—6 до И—6, К—6 и от О—6 до У—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-шесть до И-шесть" + breakTimeLong +
						"Ка-шесть" + breakTimeLong +
						"И - от О-шесть до У-шесть" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 13:
				resp.Text = "От Е—7 до З—7, Н—7 и от С—7 до Ф—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-семь до Зэ-семь" + breakTimeLong +
						"Эн-семь" + breakTimeLong +
						"И - от Эс-семь до Эф-семь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 12:
				resp.Text = "Д—8, Ё—8, Ж—8 и от Т—8 до Ф—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-восемь" + breakTimeLong +
						"Йо-восемь" + breakTimeLong +
						"Жэ-восемь" + breakTimeLong +
						"И - от Тэ-восемь до Эф-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 11:
				resp.Text = "Д—9, Е—9 и от Т—9 до Х—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять" + breakTimeLong +
						"Е-девять" + breakTimeLong +
						"И - от Тэ-девять до Ха-девять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 10:
				resp.Text = "В—10, от Д—10 до Ё—10 и от Т—10 до Х—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-десять" + breakTimeLong +
						"От Дэ-десять до Йо-десять" + breakTimeLong +
						"И - от Тэ-десять до Ха-десять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 9:
				resp.Text = "В—11, от Д—11 до Ж—11 и от С—11 до Ф—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" + breakTimeLong +
						"От Дэ-одиннадцать до Жэ-одиннадцать" + breakTimeLong +
						"И - от Эс-одиннадцать до Эф-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 8:
				resp.Text = "В—12, Г—12, от Е—12 до З—12 и от Р—12 до Х—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двенадцать" + breakTimeLong +
						"Гэ-двенадцать" + breakTimeLong +
						"От Е-двенадцать до Зэ-двенадцать" + breakTimeLong +
						"И - от Эр-двенадцать до Ха-двенадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "От Б—13 до Г—13, от Е—13 до И—13, К—13 и от О—13 до Х—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-тринадцать до Гэ-тринадцать" + breakTimeLong +
						"От Е-тринадцать до И-тринадцать" + breakTimeLong +
						"Ка-тринадцать" + breakTimeLong +
						"И - от О-тринадцать до Ха-тринадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "От А—14 до Д—14, Ё—14, Ж—14 и от Т—14 до Ц—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-четырнадцать до Дэ-четырнадцать" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"Жэ-четырнадцать" + breakTimeLong +
						"И - от Тэ-четырнадцать до Цэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "От А—15 до Е—15 и от У—15 до Ч—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-пятнадцать до Е-пятнадцать" + breakTimeLong +
						"И - от У-пятнадцать до Чэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "От Б—16 до Е—16 и от Ф—16 до Ш—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-шестнадцать до Е-шестнадцать" + breakTimeLong +
						"И - от Эф-шестнадцать до Ша-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От Б—17 до Е—17 и от Ф—17 до Ш—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-семнадцать до Е-семнадцать" + breakTimeLong +
						"И - от Эф-семнадцать до Ша-семнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "В—18, Д—18, Р—18, Х—18 и Ч—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ—восемнадцать" + breakTimeLong +
						"Дэ-восемнадцать" + breakTimeLong +
						"Эр-восемнадцать" + breakTimeLong +
						"Ха-восемнадцать" + breakTimeLong +
						"И - Чэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—19, И—19, П—19 и С—19\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ—девятнадцать" + breakTimeLong +
						"И-девятнадцать" + breakTimeLong +
						"Пэ-девятнадцать" + breakTimeLong +
						"И - Эс-девятнадцать" + breakTimeLong +
						"Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали оранжевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали оранжевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 2:
				resp.Text = "И—9, Р—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-девять, Эр-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "И—10, Й—10, П—10 и Р—10\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-десять" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"Пэ-десять" + breakTimeLong +
						"И - Эр-десять" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Й—8, П—8\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - восемь, Пэ-восемь" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Голубой цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Голубой цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 16:
				resp.Text = "Й—9, К—9, О—9 и П—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - девять" + breakTimeLong +
						"Ка-девять" + breakTimeLong +
						"О-девять" + breakTimeLong +
						"И - Пэ-девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 15:
				resp.Text = "З—16, И—16, Р—16 и С—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-шестнадцать" + breakTimeLong +
						"И-шестнадцать" + breakTimeLong +
						"Эр-шестнадцать" + breakTimeLong +
						"И - Эс-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 14:
				resp.Text = "От Ё—17 до У—17, кроме: М—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-семнадцать до У-семнадцать" + breakTimeMiddle +
						"Кроме: Эм-семнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 13:
				resp.Text = "От Е—18 до Ф—18, кроме: И—18, Р—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-восемнадцать до Эф-восемнадцать" + breakTimeMiddle +
						"Кроме: И-восемнадцать, Эр-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 12:
				resp.Text = "Е—19, от К—19 до О—19 и Ф—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девятнадцать" + breakTimeLong +
						"От Ка-девятнадцать до О-девятнадцать" + breakTimeLong +
						"И - Эф-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 11:
				resp.Text = "От К—20 до О—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать до О-двадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 10:
				resp.Text = "От Ё—21 до З—21, от К—21 до Н—21 и от С—21 до У—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать один до Зэ-двадцать один" + breakTimeLong +
						"От Ка-двадцать один до Эн-двадцать один" + breakTimeLong +
						"И - от Эс-двадцать один до У-двадцать один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 9:
				resp.Text = "От Й—22 до М—22 и П—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - двадцать два до Эм-двадцать два" + breakTimeLong +
						"И - Пэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 8:
				resp.Text = "От И—23 до К—23 и от О—23 до Р—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-двадцать три до Ка-двадцать три" + breakTimeLong +
						"И - от О-двадцать три до Эр-двадцать три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 7:
				resp.Text = "От М—24 до С—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-двадцать четыре до Эс-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 6:
				resp.Text = "От Ж—25 до Т—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двадцать пять до Тэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "От Ё—26 до У—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать шесть до У-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "От Ё—27 до У—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать семь до У-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "От Ё—28 до У—28"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать восемь до У-двадцать восемь" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "От Е—29 до Ф—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двадцать девять до Эф-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Е—30 до Ф—30\nРисование синего цвета завершено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-тридцать до Эф-тридцать" + breakTimeLong +
						"Рисование синего цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование синего цвета завершено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование синего цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали синий цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали синий цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 10:
				resp.Text = "От В—6 до В—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-шесть до Вэ-восемь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 9:
				resp.Text = "От Г—9 до Г—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-девять до Гэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "Д—12, Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двенадцать, Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "Е—14, Ё—15, И—18, Й—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четырнадцать" + breakTimeLong +
						"Йо-пятнадцать" + breakTimeLong +
						"И-восемнадцать" + breakTimeLong +
						"И-краткое - девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "П—20, Р—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-двадцать, Эр-двадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "От О—21 до Р—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-двадцать один до Эр-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "Н—22, О—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-двадцать два, О-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "Е—23, Ё—23 и от Л—23 до Н—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать три" + breakTimeLong +
						"Йо-двадцать три" + breakTimeLong +
						"И - от Эль-двадцать три до Эн-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От Е—24 до Л—24, Т—24 и У—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двадцать четыре до Эль-двадцать четыре" + breakTimeLong +
						"Тэ-двадцать четыре" + breakTimeLong +
						"И - У-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—25 и от Ф—25 до Ч—25\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать пять" + breakTimeLong +
						"И - от Эф-двадцать пять до Чэ-двадцать пять" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Коричневый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Коричневый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 7:
				resp.Text = "В—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Г—6, Г—20 и Г—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-шесть" + breakTimeLong +
						"Гэ-двадцать" + breakTimeLong +
						"И - Гэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Д—7, Д—21 и Д—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семь" + breakTimeLong +
						"Дэ-двадцать один" + breakTimeLong +
						"И - Дэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Е—8, Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемь, Е-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Ё—9, Ж—10, З—11, И—12, Й—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-девять" + breakTimeLong +
						"Жэ-десять" + breakTimeLong +
						"Зэ-одиннадцать" + breakTimeLong +
						"И-двенадцать" + breakTimeLong +
						"И-краткое - тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Р—19, С—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-девятнадцать, Эс-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ф—22, Х—23 и Ц—24\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эф-двадцать два" + breakTimeLong +
						"Ха-двадцать три" + breakTimeLong +
						"И - Цэ-двадцать четыре" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel3:
		switch color {
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 13:
				resp.Text = "От В—2 до В—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-два до Вэ-пять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 12:
				resp.Text = "От Г—3 до Г—6, от Г—10 до Г—12 и от Г—16 до Г—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-три до Гэ-шесть" + breakTimeLong +
						"От Гэ-десять до Гэ-двенадцать" + breakTimeLong +
						"И - от Гэ-шестнадцать до Гэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 11:
				resp.Text = "От Д—5 до Д—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 10:
				resp.Text = "От Е—11 до Е—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-одиннадцать до Е-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 9:
				resp.Text = "От Ё—11 до Ё—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-одиннадцать до Йо-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "Ж—9, Ж—10 и от Ж—14 до Ж—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-девять" + breakTimeLong +
						"Жэ-десять" + breakTimeLong +
						"И - от Жэ-четырнадцать до Жэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "От И—16 до И—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-шестнадцать до И-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "От Й—8 до Й—10 и от Й—14 до Й—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - восемь до И-краткое - десять" + breakTimeLong +
						"И - от И-краткое - четырнадцать до И-краткое - двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "К—10, К—11 и от К—13 до К—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-десять" + breakTimeLong +
						"Ка-одиннадцать" + breakTimeLong +
						"И - от Ка-тринадцать до Ка-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "Л—6 и от Л—11 до Л—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-шесть" + breakTimeLong +
						"И - от Эль-одиннадцать до Эль-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От М—5 до М—22, кроме: М—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-пять до Эм-двадцать два" + breakTimeMiddle +
						"Кроме: Эм-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От Н—3 до Н—6, от Н—10 до Н—12 и от Н—16 до Н—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-три до Эн-шесть" + breakTimeLong +
						"От Эн-десять до Эн-двенадцать" + breakTimeLong +
						"И - от Эн-шестнадцать до Эн-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От О—2 до О—5\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-два до О-пять" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 18:
				resp.Text = "А—13, А—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-тринадцать, А-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "От Б—1 до Б—5 и от Б—12 до Б—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-один до Бэ-пять" + breakTimeLong +
						"И - от Бэ-двенадцать до Бэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "В—1, В—6 и от В—10 до В—19, кроме: В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-один" + breakTimeLong +
						"Вэ-шесть" + breakTimeLong +
						"И - от Вэ-десять до Вэ-девятнадцать" + breakTimeMiddle +
						"Кроме: Вэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "Г—2, от Г—7 до Г—9, от Г—13 до Г—15 и от Г—20 до Г—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два" + breakTimeLong +
						"От Гэ-семь до Гэ-девять" + breakTimeLong +
						"От Гэ-тринадцать до Гэ-пятнадцать" + breakTimeLong +
						"И - от Гэ-двадцать до Гэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Д—3, Д—4 и от Д—23 до Д—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-три" + breakTimeLong +
						"Дэ-четыре" + breakTimeLong +
						"И - от Дэ-двадцать три до Дэ-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "От Е—5 до Е—10, Е—28 и Е—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пять до Е-десять" + breakTimeLong +
						"Е-двадцать восемь" + breakTimeLong +
						"И - Е-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "От Ё—5 до Ё—10, Ё—25, Ё—26 и Ё—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пять до Йо-десять" + breakTimeLong +
						"Йо-двадцать пять" + breakTimeLong +
						"Йо-двадцать шесть" + breakTimeLong +
						"И - Йо-тридцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "От Ж—4 до Ж—8, Ж—11, Ж—12, Ж—27, Ж—28 и Ж—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-четыре до Жэ-восемь" + breakTimeLong +
						"Жэ-одиннадцать" + breakTimeLong +
						"Жэ-двенадцать" + breakTimeLong +
						"Жэ-двадцать семь" + breakTimeLong +
						"Жэ-двадцать восемь" + breakTimeLong +
						"И - Жэ-тридцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "От З—4 до З—9, З—13 и З—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четыре до Зэ-девять" + breakTimeLong +
						"Зэ-тринадцать" + breakTimeLong +
						"И - Зэ-тридцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "От И—4 до И—11, И—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-четыре до И-одиннадцать" + breakTimeLong +
						"И-тридцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От Й—4 до Й—7, Й—11, Й—27, Й—28, Й—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - четыре до И-краткое - семь" + breakTimeLong +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И-краткое - двадцать семь" + breakTimeLong +
						"И-краткое - двадцать восемь" + breakTimeLong +
						"И-краткое - тридцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "От К—5 до К—9, К—12, К—25, К—26 и К—30"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-пять до Ка-девять" + breakTimeLong +
						"Ка-двенадцать" + breakTimeLong +
						"Ка-двадцать пять" + breakTimeLong +
						"Ка-двадцать шесть" + breakTimeLong +
						"И - Ка-тридцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Л—5, от Л—7 до Л—10, Л—28 и Л—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-пять" + breakTimeLong +
						"От Эль-семь до Эль-десять" + breakTimeLong +
						"Эль-двадцать восемь" + breakTimeLong +
						"И - Эль-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "М—3, М—4, М—11 и от М—23 до М—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-три" + breakTimeLong +
						"Эм-четыре" + breakTimeLong +
						"Эм-одиннадцать" + breakTimeLong +
						"И - от Эм-двадцать три до Эм-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Н—2, от Н—7 до Н—9, от Н—13 до Н—15 и от Н—20 до Н—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-два" + breakTimeLong +
						"От Эн-семь до Эн-девять" + breakTimeLong +
						"От Эн-тринадцать до Эн-пятнадцать" + breakTimeLong +
						"И - от Эн-двадцать до Эн-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "О—1, О—6 и от О—10 до О—19, кроме: О—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-один" + breakTimeLong +
						"О-шесть" + breakTimeLong +
						"И - от О-десять до О-девятнадцать" + breakTimeMiddle +
						"Кроме: О-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От П—1 до П—5 и от П—12 до П—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-один до Пэ-пять" + breakTimeLong +
						"И - от Пэ-двенадцать до Пэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Р—13, Р—14\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-тринадцать, Эр-четырнадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel4:
		switch color {
		case Beige:
			switch colorsAndQuantity[Beige] {
			case 8:
				resp.Text = "От Г—9 до Ё—9 и от У—9 до Ц—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-девять до Йо-девять" + breakTimeLong +
						"И - от У-девять до Цэ-девять" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 7:
				resp.Text = "От В—10 до Д—10, Ж—10, Т—10 и от Х—10 до Ч—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-десять до Дэ-десять" + breakTimeLong +
						"Жэ-десять" + breakTimeLong +
						"Тэ-десять" + breakTimeLong +
						"И - от Ха-десять до Чэ-десять" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 6:
				resp.Text = "В—11, Г—11, Ц—11 и Ч—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" + breakTimeLong +
						"Гэ-одиннадцать" + breakTimeLong +
						"Цэ-одиннадцать" + breakTimeLong +
						"И - Чэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 5:
				resp.Text = "От Б—12 до Г—12 и от Ц—12 до Ш—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-двенадцать до Гэ-двенадцать" + breakTimeLong +
						"И - от Цэ-двенадцать до Ша-двенадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 4:
				resp.Text = "От Б—13 до Д—13 и от Х—13 до Ш—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-тринадцать до Дэ-тринадцать" + breakTimeLong +
						"И - от Ха-тринадцать до Ша-тринадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 3:
				resp.Text = "От А—14 до В—14, Е—14, Ё—14, У—14, Ф—14 и от Ч—14 до Щ—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-четырнадцать до Вэ-четырнадцать" + breakTimeLong +
						"Е-четырнадцать" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"У-четырнадцать" + breakTimeLong +
						"Эф-четырнадцать" + breakTimeLong +
						"И - от Чэ-четырнадцать до Ща-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 2:
				resp.Text = "А—15, Б—15, Ш—15 и Щ—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пятнадцать" + breakTimeLong +
						"Бэ-пятнадцать" + breakTimeLong +
						"Ша-пятнадцать" + breakTimeLong +
						"И - Ща-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 1:
				colorsAndQuantity[Beige]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "А—16, Щ—16\nРисование бежевого цвета завершено. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шестнадцать, Ща-шестнадцать" + breakTimeLong +
						"\nРисование бежевого цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Рисование бежевого цвета завершено. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Рисование бежевого цвета завершено. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали бежевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали бежевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 13:
				resp.Text = "Ё—2, У—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-два, У-два" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 12:
				resp.Text = "Ж—3, от Й—3 до П—3 и Т—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"От И-краткое - три до Пэ-три" + breakTimeLong +
						"И - Тэ-три" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 11:
				resp.Text = "От З—4 до С—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четыре до Эс-четыре" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 10:
				resp.Text = "От Ж—5 до Т—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-пять до Тэ-пять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 9:
				resp.Text = "От Ж—6 до Т—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-шесть до Тэ-шесть" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 8:
				resp.Text = "От Ё—7 до Т—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-семь до Тэ-семь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "Ж—8, З—8, от К—8 до О—8, С—8 и Т—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-восемь" + breakTimeLong +
						"Зэ-восемь" + breakTimeLong +
						"От Ка-восемь до О-восемь" + breakTimeLong +
						"Эс-восемь" + breakTimeLong +
						"И - Тэ-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "К—9, О—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-девять, О-девять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "З—15, С—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-пятнадцать, Эс-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "Ё—16, З—16, И—16, Р—16 и С—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шестнадцать" + breakTimeLong +
						"Зэ-шестнадцать" + breakTimeLong +
						"И-шестнадцать" + breakTimeLong +
						"Эр-шестнадцать" + breakTimeLong +
						"И - Эс-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От Ж—17 до И—17 и от Р—17 до Т—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-семнадцать до И-семнадцать" + breakTimeLong +
						"И - от Эр-семнадцать до Тэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "Е—18, от Ж—18 до И—18 и от Р—18 до Т—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемнадцать" + breakTimeLong +
						"От Жэ-восемнадцать до И-восемнадцать" + breakTimeLong +
						"И - от Эр-восемнадцать до Тэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ж—19 до Й—19 и от П—19 до Т—19\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девятнадцать до И-краткое - девятнадцать" + breakTimeLong +
						"И - от Пэ-девятнадцать до Тэ-девятнадцать" + breakTimeLong +
						"\nОранжевый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали оранжевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали оранжевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 3:
				resp.Text = "З—9, С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-девять, Эс-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "И—10, Р—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-десять, Эр-десять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—11, И—11, Р—11 и У—11\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-одиннадцать" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"Эр-одиннадцать" + breakTimeLong +
						"И - У-одиннадцать" + breakTimeLong +
						"\nСветло-розовый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали светло-розовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали светло-розовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—11 и от Л—12 до Н—12\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-одиннадцать" + breakTimeLong +
						"И - от Эль-двенадцать до Эн-двенадцать" + breakTimeLong +
						"\nКрасный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали красный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали красный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 19:
				resp.Text = "Е—1, Ё—1, У—1 и Ф—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-один" + breakTimeLong +
						"Йо-один" + breakTimeLong +
						"У-один" + breakTimeLong +
						"И - Эф-один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 18:
				resp.Text = "Е—2, Ж—2, от Й—2 до П—2, Т—2 и Ф—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-два" + breakTimeLong +
						"Жэ-два" + breakTimeLong +
						"От И-краткое - два до Пэ-два" + breakTimeLong +
						"Тэ-два" + breakTimeLong +
						"И - Эф-два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 17:
				resp.Text = "Е—3, З—3, И—3, Р—3, С—3 и Ф—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" + breakTimeLong +
						"Зэ-три" + breakTimeLong +
						"И-три" + breakTimeLong +
						"Эр-три" + breakTimeLong +
						"Эс-три" + breakTimeLong +
						"И - Эф-три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 16:
				resp.Text = "Е—4, Ф—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четыре, Эф-четыре" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 15:
				resp.Text = "Ё—5, У—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-пять, У-пять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 14:
				resp.Text = "Ё—6, У—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шесть, У-шесть" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 13:
				resp.Text = "Е—7, Ф—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-семь, Эф-семь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 12:
				resp.Text = "От Г—8 до Ё—8 и от У—8 до Ц—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-восемь до Йо-восемь" + breakTimeLong +
						"И - от У-восемь до Цэ-восемь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 11:
				resp.Text = "В—9, Ж—9, Т—9 и Ч—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-девять" + breakTimeLong +
						"Жэ-девять" + breakTimeLong +
						"Тэ-девять" + breakTimeLong +
						"И - Чэ-девять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 10:
				resp.Text = "Б—10, Е—10, Ё—10, З—10, С—10, У—10, Ф—10 и Ш—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" + breakTimeLong +
						"Е-десять" + breakTimeLong +
						"Йо-десять" + breakTimeLong +
						"Зэ-десять" + breakTimeLong +
						"Эс-десять" + breakTimeLong +
						"У-десять" + breakTimeLong +
						"Эф-десять" + breakTimeLong +
						"И - Ша-десять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 9:
				resp.Text = "Б—11, Д—11, Ж—11, З—11, С—11, Т—11, Х—11 и Ш—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-одиннадцать" + breakTimeLong +
						"Дэ-одиннадцать" + breakTimeLong +
						"Жэ-одиннадцать" + breakTimeLong +
						"Зэ-одиннадцать" + breakTimeLong +
						"Эс-одиннадцать" + breakTimeLong +
						"Тэ-одиннадцать" + breakTimeLong +
						"Ха-одиннадцать" + breakTimeLong +
						"И - Ша-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "А—12, Д—12, Х—12 и Щ—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двенадцать" + breakTimeLong +
						"Дэ-двенадцать" + breakTimeLong +
						"Ха-двенадцать" + breakTimeLong +
						"И - Ща-двенадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "А—13, Е—13, Ё—13, У—13, Ф—13 и Щ—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-тринадцать" + breakTimeLong +
						"Е-тринадцать" + breakTimeLong +
						"Йо-тринадцать" + breakTimeLong +
						"У-тринадцать" + breakTimeLong +
						"Эф-тринадцать" + breakTimeLong +
						"И - Ща-тринадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "Г—14, Д—14, Ж—14, З—14, С—14, Т—14, Х—14 и Ц—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четырнадцать" + breakTimeLong +
						"Дэ-четырнадцать" + breakTimeLong +
						"Жэ-четырнадцать" + breakTimeLong +
						"Зэ-четырнадцать" + breakTimeLong +
						"Эс-четырнадцать" + breakTimeLong +
						"Тэ-четырнадцать" + breakTimeLong +
						"Ха-четырнадцать" + breakTimeLong +
						"И - Цэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "В—15, от Е—15 до Ж—15, И—15, Й—15, П—15, Р—15, от Т—15 до Ф—15 и Ч—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-пятнадцать" + breakTimeLong +
						"От Е-пятнадцать до Жэ-пятнадцать" + breakTimeLong +
						"И-пятнадцать" + breakTimeLong +
						"И-краткое - пятнадцать" + breakTimeLong +
						"Пэ-пятнадцать" + breakTimeLong +
						"Эр-пятнадцать" + breakTimeLong +
						"От Тэ-пятнадцать до Эф-пятнадцать" + breakTimeLong +
						"И - Чэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "Б—16, Д—16, Ж—16, Т—16 и Ш—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-шестнадцать" + breakTimeLong +
						"Дэ-шестнадцать" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"Тэ-шестнадцать" + breakTimeLong +
						"И - Ша-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "Б—17, Д—17, Ё—17, Й—17, П—17, У—17 и Ш—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-семнадцать" + breakTimeLong +
						"Дэ-семнадцать" + breakTimeLong +
						"Йо-семнадцать" + breakTimeLong +
						"И-краткое - семнадцать" + breakTimeLong +
						"Пэ-семнадцать" + breakTimeLong +
						"У-семнадцать" + breakTimeLong +
						"И - Ша-семнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "Д—18, Ё—18, Й—18, П—18 и У—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-восемнадцать" + breakTimeLong +
						"Йо-восемнадцать" + breakTimeLong +
						"И-краткое - восемнадцать" + breakTimeLong +
						"Пэ-восемнадцать" + breakTimeLong +
						"И - У-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—19, Ё—19, К—19, О—19 и У—19\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девятнадцать" + breakTimeLong +
						"Йо-девятнадцать" + breakTimeLong +
						"Ка-девятнадцать" + breakTimeLong +
						"О-девятнадцать" + breakTimeLong +
						"И - У-девятнадцать" + breakTimeLong +
						"\nКоричневый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 6:
				resp.Text = "И—8, Й—8, П—8 и Р—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-восемь" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"Пэ-восемь" + breakTimeLong +
						"И - Эр-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "От Л—9 до Н—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-девять до Эн-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Й—10, М—10 и П—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - десять" + breakTimeLong +
						"Эм-десять" + breakTimeLong +
						"И - Пэ-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От К—11 до О—11, кроме: М—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-одиннадцать до О-одиннадцать" + breakTimeMiddle +
						"Кроме: Эм-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "К—12, О—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-двенадцать, О-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Л—13 до Н—13\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-тринадцать до Эн-тринадцать" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel5:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 7:
				resp.Text = "А—2, А—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-два, А-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "От Б—1 до Б—3 и от Б—26 до Б—28"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-один до Бэ-три" + breakTimeLong +
						"И - от Бэ-двадцать шесть до Бэ-двадцать восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "В—2, В—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-два, Вэ-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "Щ—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ща-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "Ъ—5 и от Ъ—21 до Ъ—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Твердый-знак - пять" + breakTimeLong +
						"И - от Твердый-знак - двадцать один до Твердый-знак - двадцать три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От Ы—4 до Ы—6 и Ы—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ы-четыре до Ы-шесть" + breakTimeLong +
						"И - Ы-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ь—5\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Мягкий-знак - пять" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 12:
				resp.Text = "Б—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 11:
				resp.Text = "В—11, Г—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать, Гэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 10:
				resp.Text = "А—12 и от В—12 до Д—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двенадцать" + breakTimeLong +
						"И - от Вэ-двенадцать до Дэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 9:
				resp.Text = "От В—13 до Е—13, М—13 и Н—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-тринадцать до Е-тринадцать" + breakTimeLong +
						"Эм-тринадцать" + breakTimeLong +
						"И - Эн-тринадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 8:
				resp.Text = "Б—14, от Д—14 до Ё—14 и от Л—14 до Н—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четырнадцать" + breakTimeLong +
						"От Дэ-четырнадцать до Йо-четырнадцать" + breakTimeLong +
						"И - от Эль-четырнадцать до Эн-четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 7:
				resp.Text = "А—15, от Е—15 до И—15 и от Л—15 до Н—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пятнадцать" + breakTimeLong +
						"От Е-пятнадцать до И-пятнадцать" + breakTimeLong +
						"И - от Эль-пятнадцать до Эн-пятнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 6:
				resp.Text = "М—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-шестнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 5:
				resp.Text = "Н—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 4:
				resp.Text = "Ж—27 и от И—27 до О—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать семь" + breakTimeLong +
						"И - от И-двадцать семь до О-двадцать семь" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 3:
				resp.Text = "Н—28"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-двадцать восемь" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "И—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-двадцать девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Л—30\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-тридцать" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 20:
				resp.Text = "К—3, Л—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-три, Эль-три" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 19:
				resp.Text = "От З—4 до К—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четыре до Ка-четыре" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 18:
				resp.Text = "От З—5 до Й—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-пять до И-краткое - пять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 17:
				resp.Text = "От Ж—6 до И—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-шесть до И-шесть" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 16:
				resp.Text = "И—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-семь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 15:
				resp.Text = "От Ц—10 до Ш—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-десять до Ша-десять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 14:
				resp.Text = "Ц—11, Ч—11 и Щ—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Цэ-одиннадцать" + breakTimeLong +
						"Чэ-одиннадцать" + breakTimeLong +
						"И - Ща-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 13:
				resp.Text = "Ц—12, Ч—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Цэ-двенадцать, Чэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 12:
				resp.Text = "От Х—13 до Ш—13 и Ы—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ха-тринадцать до Ша-тринадцать" + breakTimeLong +
						"И - Ы-тринадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 11:
				resp.Text = "От Ф—14 до Ч—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-четырнадцать до Чэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 10:
				resp.Text = "От Т—15 до Ш—15, кроме: Ч—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-пятнадцать до Ша-пятнадцать" + breakTimeMiddle +
						"Кроме: Чэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 9:
				resp.Text = "В—16, Е—16, Ё—16, от У—16 до Ч—16 и Ы—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-шестнадцать" + breakTimeLong +
						"Е-шестнадцать" + breakTimeLong +
						"Йо-шестнадцать" + breakTimeLong +
						"От У-шестнадцать до Чэ-шестнадцать" + breakTimeLong +
						"И - Ы-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 8:
				resp.Text = "Д—17, Е—17 и от У—17 до Ц—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семнадцать" + breakTimeLong +
						"Е-семнадцать" + breakTimeLong +
						"И - от У-семнадцать до Цэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 7:
				resp.Text = "Г—18, Д—18, У—18 и Ф—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-восемнадцать" + breakTimeLong +
						"Дэ-восемнадцать" + breakTimeLong +
						"У-восемнадцать" + breakTimeLong +
						"И - Эф-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 6:
				resp.Text = "Д—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "Д—22, Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двадцать два, Е-двадцать два" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "От Д—23 до Ё—23, Х—23 и Ц—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-двадцать три до Йо-двадцать три" +  breakTimeLong +
						"Ха-двадцать три" + breakTimeLong +
						"И - Цэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "Г—24, Е—24, Ё—24 и от Ф—24 до Ц—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-двадцать четыре" + breakTimeLong +
						"Е-двадцать четыре" + breakTimeLong +
						"Йо-двадцать четыре" + breakTimeLong +
						"И - от Эф-двадцать четыре до Цэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "Ё—25, Ж—25 и от Х—25 до Ч—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать пять" + breakTimeLong +
						"Жэ-двадцать пять" + breakTimeLong +
						"И - от Ха-двадцать пять до Чэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Д—26 и от Ф—26 до Ц—26\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двадцать шесть" + breakTimeLong +
						"И - от Эф-двадцать шесть до Цэ-двадцать шесть" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали голубой цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали голубой цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Purple:
			switch colorsAndQuantity[Purple] {
			case 19:
				resp.Text = "Г—4, Е—4 и Ж—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четыре" + breakTimeLong +
						"Е-четыре" + breakTimeLong +
						"И - Жэ-четыре" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 18:
				resp.Text = "От Е—5 до Ж—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пять до Жэ-пять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 17:
				resp.Text = "А—6, В—6 и от Д—6 до Ё—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шесть" + breakTimeLong +
						"Вэ-шесть" + breakTimeLong +
						"И - от Дэ-шесть до Йо-шесть" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 16:
				resp.Text = "От Г—7 до Ё—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-семь до Йо-семь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 15:
				resp.Text = "От Г—8 до Е—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-восемь до Е-восемь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 14:
				resp.Text = "От В—9 до Д—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-девять до Дэ-девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 13:
				resp.Text = "Г—10, Д—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-десять, Дэ-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 12:
				resp.Text = "Д—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 11:
				resp.Text = "У—9, Ф—19 и Ч—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"У-девятнадцать" + breakTimeLong +
						"Эф-девятнадцать" + breakTimeLong +
						"И - Чэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 10:
				resp.Text = "В—20, от У—20 до Х—20 и Щ—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать" + breakTimeLong +
						"От У-двадцать до Ха-двадцать" + breakTimeLong +
						"И - Ща-двадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 9:
				resp.Text = "От В—21 до Д—21, от У—21 до Х—21 и Ш—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-двадцать один до Дэ-двадцать один" + breakTimeLong +
						"От У-двадцать один до Ха-двадцать один" + breakTimeLong +
						"И - Ша-двадцать один" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 8:
				resp.Text = "Г—22 и от У—22 до Ц—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-двадцать два" + breakTimeLong +
						"И - от У-двадцать два до Цэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 7:
				resp.Text = "Ф—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эф-двадцать три" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 6:
				resp.Text = "О—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 5:
				resp.Text = "От О—26 до С—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-двадцать шесть до Эс-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 4:
				resp.Text = "От П—27 до Т—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-двадцать семь до Тэ-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 3:
				resp.Text = "С—28"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-двадцать восемь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 2:
				resp.Text = "П—29, Т—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-двадцать девять, Тэ-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 1:
				colorsAndQuantity[Purple]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "С—30\nФиолетовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эс-тридцать" + breakTimeLong +
						"Фиолетовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Фиолетовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Фиолетовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали фиолетовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали фиолетовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case DarkGray:
			switch colorsAndQuantity[DarkGray] {
			case 20:
				resp.Text = "От Ё—9 до Ё—12 и от Ё—18 до Ё—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девять до Йо-двенадцать" + breakTimeLong +
						"И - от Йо-восемнадцать до Йо-двадцать один" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 19:
				resp.Text = "От Ж—8 до Ж—13 и от Ж—17 до Ж—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-восемь до Жэ-тринадцать" + breakTimeLong +
						"И - от Жэ-семнадцать до Жэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 18:
				resp.Text = "От З—8 до З—13 и от З—17 до З—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-восемь до Зэ-тринадцать" + breakTimeLong +
						"И - от Зэ-семнадцать до Зэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 17:
				resp.Text = "От И—11 до И—13 и от И—17 до И—25, кроме: И—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-одиннадцать до И-тринадцать" + breakTimeLong +
						"И - от И-семнадцать до И-двадцать пять" + breakTimeMiddle +
						"Кроме: И-двадцать один" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 16:
				resp.Text = "Й—9, Й—10, Й—12 и от Й—17 до Й—25, кроме: Й—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - девять" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"И-краткое - двенадцать" + breakTimeLong +
						"И - от И-краткое - семнадцать до И-краткое - двадцать пять" + breakTimeMiddle +
						"Кроме: И-краткое - двадцать один" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 15:
				resp.Text = "От К—6 до К—11, от К—14 до К—16, от К—18 до К—21 и от К—23 до К—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-шесть до Ка-одиннадцать" + breakTimeLong +
						"От Ка-четырнадцать до Ка-шестнадцать" + breakTimeLong +
						"От Ка-восемнадцать до Ка-двадцать один" + breakTimeLong +
						"И - от Ка-двадцать три до Ка-двадцать пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 14:
				resp.Text = "От Л—6 до Л—13, Л—16, Л—17 и от Л—19 до Л—25, кроме: Л—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-шесть до Эль-тринадцать" + breakTimeLong +
						"Эль-шестнадцать" + breakTimeLong +
						"Эль-семнадцать" + breakTimeLong +
						"И - от Эль-девятнадцать до Эль-двадцать пять" + breakTimeMiddle +
						"Кроме: Эль-двадцать два" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 13:
				resp.Text = "От М—5 до М—12, М—17, от М—19 до М—21 и М—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-пять до Эм-двенадцать" + breakTimeLong +
						"Эм-семнадцать" + breakTimeLong +
						"От Эм-девятнадцать до Эм-двадцать один" + breakTimeLong +
						"И - Эм-двадцать пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 12:
				resp.Text = "От Н—5 до Н—12 и от Н—16 до Н—21, кроме: Н—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-пять до Эн-двенадцать" + breakTimeLong +
						"И - от Эн-шестнадцать до Эн-двадцать один" + breakTimeMiddle +
						"Кроме: Эн-восемнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 11:
				resp.Text = "От О—5 до О—22, кроме: О—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-пять до О-двадцать два" + breakTimeMiddle +
						"Кроме: О-восемнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 10:
				resp.Text = "От П—5 до П—12 и от П—18 до П—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-пять до Пэ-двенадцать" + breakTimeLong +
						"И - от Пэ-восемнадцать до Пэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 9:
				resp.Text = "От Р—5 до Р—24, кроме: Р—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эр-пять до Эр-двадцать четыре" + breakTimeMiddle +
						"Кроме: Эр-пятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 8:
				resp.Text = "От С—5 до С—24, кроме: С—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-пять до Эс-двадцать четыре" + breakTimeMiddle +
						"Кроме: Эс-пятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 7:
				resp.Text = "От Т—6 до Т—8, от Т—10 до Т—13 и от Т—23 до Т—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-шесть до Тэ-восемь" + breakTimeLong +
						"От Тэ-десять до Тэ-тринадцать" + breakTimeLong +
						"И - от Тэ-двадцать три до Тэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 6:
				resp.Text = "От У—7 до У—13 и У—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От У-семь до У-тринадцать" + breakTimeLong +
						"И - У-двадцать пять" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 5:
				resp.Text = "От Ф—6 до Ф—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-шесть до Эф-двенадцать" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 4:
				resp.Text = "Х—6, Х—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ха-шесть, Ха-семь" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 3:
				resp.Text = "От Ц—5 до Ц—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-пять до Цэ-семь" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 2:
				resp.Text = "От Ч—3 до Ч—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Чэ-три до Чэ-шесть" +
						"</speak>"
				colorsAndQuantity[DarkGray]--
			case 1:
				colorsAndQuantity[DarkGray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ш—2\nВы нарисовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ша-два" + breakTimeLong +
						"Вы нарисовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали темно-серый цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали темно-серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Темно-серый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Темно-серый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 26:
				resp.Text = "Ш—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ша-один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 25:
				resp.Text = "Ч—2, Щ—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Чэ-два, Ща-два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 24:
				resp.Text = "Ц—3, Ш—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Цэ-три, Ша-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 23:
				resp.Text = "От Л—4 до С—4, Ц—4 и Ш—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-четыре до Эс-четыре" + breakTimeLong +
						"Цэ-четыре" + breakTimeLong +
						"И - Ша-четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 22:
				resp.Text = "К—5, Л—5, Т—5, Ф—5, Х—5 и Ш—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-пять" + breakTimeLong +
						"Эль-пять" + breakTimeLong +
						"Тэ-пять" + breakTimeLong +
						"Эф-пять" + breakTimeLong +
						"Ха-пять" + breakTimeLong +
						"И - Ша-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 21:
				resp.Text = "Й—6, У—6 и Ш—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - шесть" + breakTimeLong +
						"У-шесть" + breakTimeLong +
						"И - Ша-шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 20:
				resp.Text = "Ж—7, З—7, Й—7 и Ч—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-семь" + breakTimeLong +
						"Зэ-семь" + breakTimeLong +
						"И-краткое - семь" + breakTimeLong +
						"И - Чэ-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 19:
				resp.Text = "Ё—8, И—8, Й—8, Х—8 и Ц—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-восемь" + breakTimeLong +
						"И-восемь" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"Ха-восемь" + breakTimeLong +
						"И - Цэ-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 18:
				resp.Text = "Е—9, И—9, Т—9 и Х—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девять" + breakTimeLong +
						"И-девять" + breakTimeLong +
						"Тэ-девять" + breakTimeLong +
						"И - Ха-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "Е—10, И—10 и Х—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-десять" + breakTimeLong +
						"И-десять" + breakTimeLong +
						"И - Ха-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "Е—11, Й—11 и Х—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-одиннадцать" + breakTimeLong +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И - Ха-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "Е—12, К—12 и Х—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двенадцать" + breakTimeLong +
						"Ка-двенадцать" + breakTimeLong +
						"И - Ха-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Ё—13, Й—13, К—13, П—13 и Ф—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-тринадцать" + breakTimeLong +
						"И-краткое - тринадцать" + breakTimeLong +
						"Ка-тринадцать" + breakTimeLong +
						"Пэ-тринадцать" + breakTimeLong +
						"И - Эф-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "От Ж—14 до Й—14, П—14, Т—14 и У—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-четырнадцать до И-краткое - четырнадцать" + breakTimeLong +
						"Пэ-четырнадцать" + breakTimeLong +
						"Тэ-четырнадцать" + breakTimeLong +
						"И - У-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Й—15 и от П—15 до С—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - пятнадцать" + breakTimeLong +
						"И - от Пэ-пятнадцать до Эс-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "А—16, Б—16, от Ж—16 до Й—16, П—16 и Т—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шестнадцать" + breakTimeLong +
						"Бэ-шестнадцать" + breakTimeLong +
						"От Жэ-шестнадцать до И-краткое - шестнадцать" + breakTimeLong +
						"Пэ-шестнадцать" + breakTimeLong +
						"И - Тэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "Б—17, Ё—17, К—17, П—17 и Т—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-семнадцать" + breakTimeLong +
						"Йо-семнадцать" + breakTimeLong +
						"Ка-семнадцать" + breakTimeLong +
						"Пэ-семнадцать" + breakTimeLong +
						"И - Тэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "В—18, Е—18, от Л—18 до О—18 и Т—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-восемнадцать" + breakTimeLong +
						"Е-восемнадцать" + breakTimeLong +
						"От Эль-восемнадцать до О-восемнадцать" + breakTimeLong +
						"И - Тэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "В—19, Г—19, Е—19 и Т—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-девятнадцать" + breakTimeLong +
						"Гэ-девятнадцать" + breakTimeLong +
						"Е-девятнадцать" + breakTimeLong +
						"И - Тэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "От Г—20 до Е—20 и Т—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двадцать до Е-двадцать" + breakTimeLong +
						"И - Тэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Е—21, И—21, Й—21 и Т—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать один" + breakTimeLong +
						"И-двадцать один" + breakTimeLong +
						"И-краткое - двадцать один" + breakTimeLong +
						"И - Тэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Ё—22, от К—22 до Н—22 и Т—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать два" + breakTimeLong +
						"От Ка-двадцать два до Эн-двадцать два" + breakTimeLong +
						"И - Тэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Ж—23, от М—23 до О—23 и У—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать три" + breakTimeLong +
						"От Эм-двадцать три до О-двадцать три" + breakTimeLong +
						"И - У-двадцать три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Ж—24, от М—24 до О—24 и У—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать четыре" + breakTimeLong +
						"От Эм-двадцать четыре до О-двадцать четыре" + breakTimeLong +
						"И - У-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "З—25, Н—25, от П—25 до С—25 и Ф—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двадцать пять" + breakTimeLong +
						"Эн-двадцать пять" + breakTimeLong +
						"От Пэ-двадцать пять до Эс-двадцать пять" + breakTimeLong +
						"И - Эф-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От И—26 до М—26, Т—26 и У—26\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-двадцать шесть до Эм-двадцать шесть" + breakTimeLong +
						"Тэ-двадцать шесть" + breakTimeLong +
						"И - У-двадцать шесть" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel6:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 15:
				resp.Text = "От Е—15 до Е—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-пятнадцать до Е-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 14:
				resp.Text = "От Ё—16 до Ё—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-шестнадцать до Йо-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 13:
				resp.Text = "И—21, И—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-двадцать один, И-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 12:
				resp.Text = "Й—7, Й—8, от Й—13 до Й—15 и от Й—20 до Й—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - семь" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"От И-краткое - тринадцать до И-краткое - пятнадцать" + breakTimeLong +
						"И - от И-краткое - двадцать до И-краткое - двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 11:
				resp.Text = "От К—7 до К—9, от К—12 до К—16 и от К—19 до К—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-семь до Ка-девять" + breakTimeLong +
						"От Ка-двенадцать до Ка-шестнадцать" + breakTimeLong +
						"И - от Ка-девятнадцать до Ка-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 10:
				resp.Text = "От Л—7 до Л—9, от Л—12 до Л—16 и от Л—19 до Л—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-семь до Эль-девять" + breakTimeLong +
						"От Эль-двенадцать до Эль-шестнадцать" + breakTimeLong +
						"И - от Эль-девятнадцать до Эль-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 9:
				resp.Text = "От М—6 до М—9, от М—13 до М—15 и от М—19 до М—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-шесть до Эм-девять" + breakTimeLong +
						"От Эм-тринадцать до Эм-пятнадцать" + breakTimeLong +
						"И - от Эм-девятнадцать до Эм-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "От Н—6 до Н—8 и от Н—20 до Н—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-шесть до Эн-восемь" + breakTimeLong +
						"И - от Эн-двадцать до Эн-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "От П—13 до П—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-тринадцать до Пэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "От Р—12 до Р—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эр-двенадцать до Эр-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "От С—6 до С—8 и от С—11 до С—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-шесть до Эс-восемь" + breakTimeLong +
						"И - от Эс-одиннадцать до Эс-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "От Т—7 до Т—9 и от Т—11 до Т—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-семь до Тэ-девять" + breakTimeLong +
						"И - от Тэ-одиннадцать до Тэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "У—8, У—9, У—12 и У—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"У-восемь" + breakTimeLong +
						"У-девять" + breakTimeLong +
						"У-двенадцать" + breakTimeLong +
						"И - У-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "Ф—8, Ф—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эф-восемь, Эф-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Х—9\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ха-девять" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Green:
			switch colorsAndQuantity[Green] {
			case 25:
				resp.Text = "От У—1 до Ш—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От У-один до Ша-один" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 24:
				resp.Text = "От Т—2 до Щ—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-два до Ща-два" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 23:
				resp.Text = "От И—3 до Л—3, от С—3 до Ц—3, Щ—3 и Ъ—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-три до Эль-три" + breakTimeLong +
						"От Эс-три до Цэ-три" + breakTimeLong +
						"Ща-три" + breakTimeLong +
						"И - Твердый-знак - три" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 22:
				resp.Text = "От Ж—4 до Н—4, от С—4 до Х—4, Щ—4 и Ъ—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
					"От Жэ-четыре до Эн-четыре" + breakTimeLong +
						"От Эс-четыре до Ха-четыре" + breakTimeLong +
						"Ща-четыре" + breakTimeLong +
						"И - Твердый-знак - четыре" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 21:
				resp.Text = "От Ё—5 до О—5, от С—5 до Х—5, Щ—5 и Ъ—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
					"От Йо-пять до О-пять" + breakTimeLong +
						"От Эс-пять до Ха-пять" + breakTimeLong +
						"Ща-пять" + breakTimeLong +
						"И - Твердый-знак - пять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 20:
				resp.Text = "От Ё—6 до Л—6 и от Т—6 до Ъ—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-шесть до Эль-шесть" + breakTimeLong +
						"И - от Тэ-шесть до Твердый-знак - шесть" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 19:
				resp.Text = "От Ж—7 до И—7 и от У—7 до Ъ—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-семь до И-семь" + breakTimeLong +
						"И - от У-семь до Твердый-знак - семь" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 18:
				resp.Text = "От Х—8 до Ъ—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ха-восемь до Твердый-знак - восемь" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 17:
				resp.Text = "От Ц—9 до Щ—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-девять до Ща-девять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 16:
				resp.Text = "Ч—10, Ш—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Чэ-десять, Ша-десять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 15:
				resp.Text = "От В—13 до Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-тринадцать до Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 14:
				resp.Text = "От Б—14 до Д—14 и от С—14 до Ф—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-четырнадцать до Дэ-четырнадцать" + breakTimeLong +
						"И - от Эс-четырнадцать до Эф-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 13:
				resp.Text = "От А—15 до Д—15 и от Р—15 до Х—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-пятнадцать до Дэ-пятнадцать" + breakTimeLong +
						"И - от Эр-пятнадцать до Ха-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 12:
				resp.Text = "От А—16 до Д—16 и от Р—16 до Ц—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-шестнадцать до Дэ-шестнадцать" + breakTimeLong +
						"И - от Эр-шестнадцать до Цэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 11:
				resp.Text = "От А—17 до Д—17 и от Р—17 до Ч—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-семнадцать до Дэ-семнадцать" + breakTimeLong +
						"И - от Эр-семнадцать до Чэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 10:
				resp.Text = "От А—18 до Г—18 и от С—18 до Ч—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-восемнадцать до Гэ-восемнадцать" + breakTimeLong +
						"И - от Эс-восемнадцать до Чэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 9:
				resp.Text = "Б—19, В—19 и от У—19 до Ч—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-девятнадцать" + breakTimeLong +
						"Вэ-девятнадцать" + breakTimeLong +
						"И - от У-девятнадцать до Чэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 8:
				resp.Text = "От Ф—20 до Ч—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-двадцать до Чэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 7:
				resp.Text = "Х—21, Ц—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ха-двадцать один, Цэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 6:
				resp.Text = "От И—23 до М—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-двадцать три до Эм-двадцать три" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 5:
				resp.Text = "От З—24 до М—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-двадцать четыре до Эм-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "От Ж—25 до М—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двадцать пять до Эм-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "От Ж—26 до Л—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двадцать шесть до Эль-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "От Ж—27 до К—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двадцать семь до Ка-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—28, И—28\nЗеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двадцать восемь, И-двадцать восемь" + breakTimeLong +
						"Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Зеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Зеленый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Зеленый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зеленый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 17:
				resp.Text = "От О—6 до Р—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-шесть до Эр-шесть" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 16:
				resp.Text = "От О—7 до Р—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-семь до Эр-семь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 15:
				resp.Text = "З—8, И—8 и от О—8 до Р—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-восемь" + breakTimeLong +
						"И-восемь" + breakTimeLong +
						"И - от О-восемь до Эр-восемь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 14:
				resp.Text = "От Ж—9 до Й—9 и от Н—9 до С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девять до И-краткое - девять" + breakTimeLong +
						"И - от Эн-девять до Эс-девять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 13:
				resp.Text = "От Ж—10 до Ц—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-десять до Цэ-десять" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 12:
				resp.Text = "От Ё—11 до Р—11 и от У—11 до Ц—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-одиннадцать до Эр-одиннадцать" + breakTimeLong +
						"И - от У-одиннадцать до Цэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 11:
				resp.Text = "От Ё—12 до Й—12, от М—12 до П—12 и от Ф—12 до Ц—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двенадцать до И-краткое - двенадцать" + breakTimeLong +
						"От Эм-двенадцать до Пэ-двенадцать" + breakTimeLong +
						"И - от Эф-двенадцать до Цэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 10:
				resp.Text = "От Е—13 до И—13, Н—13, О—13 и от Ф—13 до Ц—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-тринадцать до И-тринадцать" + breakTimeMiddle +
						"Эн-тринадцать" + breakTimeLong +
						"О-тринадцать" + breakTimeLong +
						"И - от Эф-тринадцать до Цэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 9:
				resp.Text = "От Е—14 до И—14, Н—14, О—14, Х—14 и Ц—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четырнадцать до И-четырнадцать" + breakTimeMiddle +
						"Эн-четырнадцать" + breakTimeLong +
						"О-четырнадцать" + breakTimeLong +
						"Ха-четырнадцать" + breakTimeLong +
						"И - Цэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 8:
				resp.Text = "От Ё—15 до И—15, Н—15 и О—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пятнадцать до И-пятнадцать" + breakTimeLong +
						"Эн-пятнадцать" + breakTimeLong +
						"И - О-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "От Ж—16 до Й—16 и от М—16 до П—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-шестнадцать до И-краткое - шестнадцать" + breakTimeLong +
						"И - От Эм-шестнадцать до Пэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "От Ж—17 до П—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-семнадцать до Пэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "От Ж—18 до Р—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-восемнадцать до Эр-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "От Ё—19 до Й—19 и от Н—19 до Т—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девятнадцать до И-краткое - девятнадцать" + breakTimeLong +
						"И - от Эн-девятнадцать до Тэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От Ё—20 до И—20 и от О—20 до С—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать до И-двадцать" + breakTimeLong +
						"И - от О-двадцать до Эс-двадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "Ж—21, З—21 и от О—21 до Р—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать один" + breakTimeLong +
						"Зэ-двадцать один" + breakTimeLong +
						"И - от О-двадцать один до Эр-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—22, О—22\nКоричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двадцать два, О-двадцать два" + breakTimeLong +
						"Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 2:
				resp.Text = "Ч—4, Ш—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Чэ-четыре, Ша-четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ч—5, Ш—5\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Чэ-пять, Ша-пять" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel7:
		switch color {
		case Beige:
			switch colorsAndQuantity[Beige] {
			case 5:
				resp.Text = "К—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 4:
				resp.Text = "От Ё—12 до Н—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двенадцать до Эн-двенадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 3:
				resp.Text = "От Ж—13 до М—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-тринадцать до Эм-тринадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 2:
				resp.Text = "От З—14 до Л—14, кроме: К—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четырнадцать до Эль-четырнадцать" + breakTimeMiddle +
						"Кроме: Ка-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 1:
				colorsAndQuantity[Beige]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От И—15 до Л—15\nБежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-пятнадцать до Эль-пятнадцать" + breakTimeLong +
						"Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Бежевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Бежевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали бежевый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали бежевый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Red:
			switch colorsAndQuantity[Red] {
			case 17:
				resp.Text = "От С—11 до Х—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-одиннадцать до Ха-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 16:
				resp.Text = "П—12, Р—12, Ц—12 и Ч—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-двенадцать" + breakTimeLong +
						"Эр-двенадцать" + breakTimeLong +
						"Цэ-двенадцать" + breakTimeLong +
						"И - Чэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 15:
				resp.Text = "О—13, от Т—13 до Ф—13 и Ш—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-тринадцать" + breakTimeLong +
						"От Тэ-тринадцать до Эф-тринадцать" + breakTimeLong +
						"И - Ша-тринадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 14:
				resp.Text = "Н—14, от Р—14 до Ц—14 и Щ—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-четырнадцать" + breakTimeLong +
						"От Эр-четырнадцать до Цэ-четырнадцать" + breakTimeLong +
						"И - Ща-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 13:
				resp.Text = "Н—15, от П—15 до С—15, от Х—15 до Ч—15 и Щ—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-пятнадцать" + breakTimeLong +
						"От Пэ-пятнадцать до Эс-пятнадцать" + breakTimeLong +
						"От Ха-пятнадцать до Чэ-пятнадцать" + breakTimeLong +
						"И - Ща-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 12:
				resp.Text = "М—16, П—16, Р—16, Ц—16, Ч—16 и Ъ—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-шестнадцать" + breakTimeLong +
						"Пэ-шестнадцать" + breakTimeLong +
						"Эр-шестнадцать" + breakTimeLong +
						"Цэ-шестнадцать" + breakTimeLong +
						"Чэ-шестнадцать" + breakTimeLong +
						"И - Твердый-знак - шестнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 11:
				resp.Text = "М—17, О—17, П—17, Ч—17, Ш—17 и Ъ—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-семнадцать" + breakTimeLong +
						"О-семнадцать" + breakTimeLong +
						"Пэ-семнадцать" + breakTimeLong +
						"Чэ-семнадцать" + breakTimeLong +
						"Ша-семнадцать" + breakTimeLong +
						"И - Твердый-знак - семнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 10:
				resp.Text = "М—18, О—18, П—18, Ч—18, Ш—18 и Ъ—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-восемнадцать" + breakTimeLong +
						"О-восемнадцать" + breakTimeLong +
						"Пэ-восемнадцать" + breakTimeLong +
						"Чэ-восемнадцать" + breakTimeLong +
						"Ша-восемнадцать" + breakTimeLong +
						"И - Твердый-знак - восемнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 9:
				resp.Text = "М—19, О—19, П—19, Ч—19, Ш—19 и Ъ—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-девятнадцать" + breakTimeLong +
						"О-девятнадцать" + breakTimeLong +
						"Пэ-девятнадцать" + breakTimeLong +
						"Чэ-девятнадцать" + breakTimeLong +
						"Ша-девятнадцать" + breakTimeLong +
						"И - Твердый-знак - девятнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 8:
				resp.Text = "В—20, Г—20, М—20, П—20, Р—20, Ц—20, Ч—20 и Ъ—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать" + breakTimeLong +
						"Гэ-двадцать" + breakTimeLong +
						"Эм-двадцать" + breakTimeLong +
						"Пэ-двадцать" + breakTimeLong +
						"Эр-двадцать" + breakTimeLong +
						"Цэ-двадцать" + breakTimeLong +
						"Чэ-двадцать" + breakTimeLong +
						"И - Твердый-знак - двадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "От Б—21 до Д—21, И—21, К—21, Н—21, от П—21 до С—21, от Х—21 до Ч—21 и Щ—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-двадцать один до Дэ-двадцать один" + breakTimeLong +
						"И-двадцать один" + breakTimeLong +
						"Ка-двадцать один" + breakTimeLong +
						"Эн-двадцать один" + breakTimeLong +
						"От Пэ-двадцать один до Эс-двадцать один" + breakTimeLong +
						"От Ха-двадцать один до Чэ-двадцать один" + breakTimeLong +
						"И - Ща-двадцать один" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "Б—22, В—22, И—22, К—22, Н—22, от Р—22 до Ц—22 и Щ—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать два" + breakTimeLong +
						"Вэ-двадцать два" + breakTimeLong +
						"И-двадцать два" + breakTimeLong +
						"Ка-двадцать два" + breakTimeLong +
						"Эн-двадцать два" + breakTimeLong +
						"От Эр-двадцать два до Цэ-двадцать два" + breakTimeLong +
						"И - Ща-двадцать два" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "О—23, от Т—23 до Ф—23 и Ш—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-двадцать три" + breakTimeLong +
						"От Тэ-двадцать три до Эф-двадцать три" + breakTimeLong +
						"И - Ша-двадцать три" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "Ж—24, П—24, Ц—24 и Ч—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать четыре" + breakTimeLong +
						"Пэ-двадцать четыре" + breakTimeLong +
						"Цэ-двадцать четыре" + breakTimeLong +
						"И - Чэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "От Ё—25 до З—25, Н—25, О—25 и от С—25 до Х—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать пять до Зэ-двадцать пять" + breakTimeLong +
						"Эн-двадцать пять" + breakTimeLong +
						"О-двадцать пять" + breakTimeLong +
						"И - от Эс-двадцать пять до Ха-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "От Ё—26 до З—26 и от Н—26 до П—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать шесть до Зэ-двадцать шесть" + breakTimeLong +
						"И - от Эн-двадцать шесть до Пэ-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ё—27 до З—27 и от Н—27 до П—27\nКрасный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать семь до Зэ-двадцать семь" + breakTimeLong +
						"И - от Эн-двадцать семь до Пэ-двадцать семь" + breakTimeLong +
						"Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Красный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Красный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Красный цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Красный цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 19:
				resp.Text = "Е—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 18:
				resp.Text = "От Ё—4 до Ё—11 и от Ё—17 до Ё—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четыре до Йо-одиннадцать" + breakTimeLong +
						"И - от Йо-семнадцать до Йо-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 17:
				resp.Text = "От Ж—3 до Ж—7, от Ж—9 до Ж—11, от Ж—17 до Ж—20 и Ж—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-три до Жэ-семь" + breakTimeLong +
						"От Жэ-девять до Жэ-одиннадцать" + breakTimeLong +
						"От Жэ-семнадцать до Жэ-двадцать" + breakTimeLong +
						"И - Жэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 16:
				resp.Text = "От З—3 до З—7, З—10, З—11, от З—16 до З—20, З—23 и З—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-три до Зэ-семь" + breakTimeLong +
						"Зэ-десять" + breakTimeLong +
						"Зэ-одиннадцать" + breakTimeLong +
						"От Зэ-шестнадцать до Зэ-двадцать" + breakTimeLong +
						"Зэ-двадцать три" + breakTimeLong +
						"И - Зэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 15:
				resp.Text = "От И—2 до И—8, И—11, от И—16 до И—20 и от И—23 до И—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-два до И-восемь" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"От И-шестнадцать до И-двадцать" + breakTimeLong +
						"И - от И-двадцать три до И-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 14:
				resp.Text = "От Й—2 до Й—4, Й—7, Й—8, Й—10, Й—11, Й—17, Й—19, Й—20, Й—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - два до И-краткое - четыре" + breakTimeLong +
						"И-краткое - семь" + breakTimeLong +
						"И-краткое - восемь" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И-краткое - семнадцать" + breakTimeLong +
						"И-краткое - девятнадцать" + breakTimeLong +
						"И-краткое - двадцать" + breakTimeLong +
						"И-краткое - двадцать три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 13:
				resp.Text = "К—2, К—3, от К—6 до К—10, К—20 и К—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-два" + breakTimeLong +
						"Ка-три" + breakTimeLong +
						"От Ка-шесть до Ка-десять" + breakTimeLong +
						"Ка-двадцать" + breakTimeLong +
						"И - Ка-двадцать три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 12:
				resp.Text = "От Л—2 до Л—4, Л—7, Л—8, Л—10, Л—11 и Л—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-два до Эль-четыре" + breakTimeLong +
						"Эль-семь" + breakTimeLong +
						"Эль-восемь" + breakTimeLong +
						"Эль-десять" + breakTimeLong +
						"Эль-одиннадцать" + breakTimeLong +
						"И - Эль-двадцать три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 11:
				resp.Text = "От М—2 до М—8, М—11 и от М—23 до М—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-два до Эм-восемь" + breakTimeLong +
						"Эм-одиннадцать" + breakTimeLong +
						"И - от Эм-двадцать три до Эм-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 10:
				resp.Text = "От Н—3 до Н—7, Н—10, Н—11 и Н—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-три до Эн-семь" + breakTimeLong +
						"Эн-десять" + breakTimeLong +
						"Эн-одиннадцать" + breakTimeLong +
						"И - Эн-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 9:
				resp.Text = "От О—3 до О—11, кроме: О—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-три до О-одиннадцать" + breakTimeMiddle +
						"Кроме: О-восемь" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 8:
				resp.Text = "От П—4 до П—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-четыре до Пэ-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 7:
				resp.Text = "От Р—17 до Р—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эр-семнадцать до Эр-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 6:
				resp.Text = "От С—16 до С—20, кроме: С—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-шестнадцать до Эс-двадцать" + breakTimeMiddle +
						"Кроме: Эс-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "Т—15, Т—16, Т—20 и Т—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Тэ-пятнадцать" + breakTimeLong +
						"Тэ-шестнадцать" + breakTimeLong +
						"Тэ-двадцать" + breakTimeLong +
						"И - Тэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "У—15, У—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"У-пятнадцать, У-двадцать один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "Ф—15, Ф—16, Ф—20 и Ф—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эф-пятнадцать" + breakTimeLong +
						"Эф-шестнадцать" + breakTimeLong +
						"Эф-двадцать" + breakTimeLong +
						"И - Эф-двадцать один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "От Х—16 до Х—20, кроме: Х—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ха-шестнадцать до Ха-двадцать" + breakTimeMiddle +
						"Кроме: Ха-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ц—17 до Ц—19\nСиний цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-семнадцать до Цэ-девятнадцать" + breakTimeLong +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали синий цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали синий цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 27:
				resp.Text = "От И—1 до М—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-один до Эм-один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 26:
				resp.Text = "Ж—2, З—2, Н—2 и О—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-два" + breakTimeLong +
						"Зэ-два" + breakTimeLong +
						"Эн-два" + breakTimeLong +
						"И - О-два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 25:
				resp.Text = "Ё—3, П—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-три, Пэ-три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 24:
				resp.Text = "Е—4, Р—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четыре, Эр-четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 23:
				resp.Text = "Е—5, Р—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-пять, Эр-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 22:
				resp.Text = "Е—7, Р—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-семь, Эр-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 21:
				resp.Text = "Е—8, Ж—8, З—8, Н—8, О—8 и Р—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемь" + breakTimeLong +
						"Жэ-восемь" + breakTimeLong +
						"Зэ-восемь" + breakTimeLong +
						"Эн-восемь" + breakTimeLong +
						"О-восемь" + breakTimeLong +
						"И - Эр-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 20:
				resp.Text = "Е—9, от З—9 до Й—9, от Л—9 до Н—9 и Р—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девять" + breakTimeLong +
						"От Зэ-девять до И-краткое - девять" + breakTimeLong +
						"От Эль-девять до Эн-девять" + breakTimeLong +
						"И - Эр-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 19:
				resp.Text = "Е—10, И—10, М—10 и от Р—10 до Х—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-десять" + breakTimeLong +
						"И-десять" + breakTimeLong +
						"Эм-десять" + breakTimeLong +
						"И - от Эр-десять до Ха-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 18:
				resp.Text = "Е—11, П—11, Р—11, Ц—11 и Ч—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-одиннадцать" + breakTimeLong +
						"Пэ-одиннадцать" + breakTimeLong +
						"Эр-одиннадцать" + breakTimeLong +
						"Цэ-одиннадцать" + breakTimeLong +
						"И - Чэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 17:
				resp.Text = "Е—12, О—12 и Ш—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двенадцать" + breakTimeLong +
						"О-двенадцать" + breakTimeLong +
						"И - Ша-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "Ё—13, Н—13 и Щ—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-тринадцать" + breakTimeLong +
						"Эн-тринадцать" + breakTimeLong +
						"И - Ща-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 15:
				resp.Text = "Ж—14, К—14, М—14 и Ъ—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четырнадцать" + breakTimeLong +
						"Ка-четырнадцать" + breakTimeLong +
						"Эм-четырнадцать" + breakTimeLong +
						"И - Твердый-знак - четырнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "З—15, М—15 и Ъ—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-пятнадцать" + breakTimeLong +
						"Эм-пятнадцать" + breakTimeLong +
						"И - Твердый-знак - пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 13:
				resp.Text = "Ё—16, Ж—16, от Й—16 до Л—16 и Ы—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шестнадцать" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"От И-краткое - шестнадцать до Эль-шестнадцать" + breakTimeLong +
						"И - Ы-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "Д—17, Е—17, Л—17 и Ы—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семнадцать" + breakTimeLong +
						"Е-семнадцать" + breakTimeLong +
						"Эль-семнадцать" + breakTimeLong +
						"И - Ы-семнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Г—18, Л—18 и Ы—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-восемнадцать" + breakTimeLong +
						"Эль-восемнадцать" + breakTimeLong +
						"И - Ы-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "В—19, Л—19 и Ы—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-девятнадцать" + breakTimeLong +
						"Эль-девятнадцать" + breakTimeLong +
						"И - Ы-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "Б—20, Ё—20, Л—20 и Ы—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать" + breakTimeLong +
						"Йо-двадцать" + breakTimeLong +
						"Эль-двадцать" + breakTimeLong +
						"И - Ы-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "А—21, Е—21, Ж—21, М—21 и Ъ—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать один" + breakTimeLong +
						"Е-двадцать один" + breakTimeLong +
						"Жэ-двадцать один" + breakTimeLong +
						"Эм-двадцать один" + breakTimeLong +
						"И - Твердый-знак - двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "А—22, Г—22, Д—22, Ж—22, М—22 и Ъ—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать два" + breakTimeLong +
						"Гэ-двадцать два" + breakTimeLong +
						"Дэ-двадцать два" + breakTimeLong +
						"Жэ-двадцать два" + breakTimeLong +
						"Эм-двадцать два" + breakTimeLong +
						"И - Твердый-знак - двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Б—23, В—23, Ё—23, Н—23 и Щ—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать три" + breakTimeLong +
						"Вэ-двадцать три" + breakTimeLong +
						"Йо-двадцать три" + breakTimeLong +
						"Эн-двадцать три" + breakTimeLong +
						"И - Ща-двадцать три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "Ё—24, от Й—24 до Л—24, О—24 и Ш—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать четыре" + breakTimeLong +
						"От И-краткое - двадцать четыре до Эль-двадцать четыре" + breakTimeLong +
						"О-двадцать четыре" + breakTimeLong +
						"И - Ша-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "Е—25, Й—25, Л—25, П—25, Р—25, Ц—25 и Ч—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать пять" + breakTimeLong +
						"И-краткое - двадцать пять" + breakTimeLong +
						"Эль-двадцать пять" + breakTimeLong +
						"Пэ-двадцать пять" + breakTimeLong +
						"Эр-двадцать пять" + breakTimeLong +
						"Цэ-двадцать пять" + breakTimeLong +
						"И - Чэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Е—26, И—26, М—26 и от Р—26 до Х—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать шесть" + breakTimeLong +
						"И-двадцать шесть" + breakTimeLong +
						"Эм-двадцать шесть" + breakTimeLong +
						"И - от Эр-двадцать шесть до Ха-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Е—27, И—27, М—27 и Р—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать семь" + breakTimeLong +
						"И-двадцать семь" + breakTimeLong +
						"Эм-двадцать семь" + breakTimeLong +
						"И - Эр-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Ё—28 до З—28 и от Н—28 до П—28\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-двадцать восемь до Зэ-двадцать восемь" + breakTimeLong +
						"И - от Эн-двадцать восемь до Пэ-двадцать восемь" +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel8:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 3:
				resp.Text = "От И—6 до К—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-шесть до Ка-шесть" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От Й—7 до Л—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - семь до Эль-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "К—8\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-восемь" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 12:
				resp.Text = "Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 11:
				resp.Text = "Д—9, Д—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять, Дэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 10:
				resp.Text = "Е—8, Е—11 и Е—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемь" + breakTimeLong +
						"Е-одиннадцать" + breakTimeLong +
						"И - Е-двенадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 9:
				resp.Text = "От Ё—10 до Ё—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-десять до Йо-тринадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 8:
				resp.Text = "От Ж—9 до Ж—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девять до Жэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 7:
				resp.Text = "От З—9 до З—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-девять до Зэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 6:
				resp.Text = "От И—10 до И—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-десять до И-четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 5:
				resp.Text = "От Й—10 до Й—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - десять до И-краткое - четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 4:
				resp.Text = "К—2, К—3 и от К—11 до К—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-два" + breakTimeLong +
						"Ка-три" + breakTimeLong +
						"И - от Ка-одиннадцать до Ка-тринадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 3:
				resp.Text = "Л—2, Л—11 и Л—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-два" + breakTimeLong +
						"Эль-одиннадцать" + breakTimeLong +
						"И - Эль-двенадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "П—5, П—6, П—9 и П—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-пять" + breakTimeLong +
						"Пэ-шесть" + breakTimeLong +
						"Пэ-девять" + breakTimeLong +
						"И - Пэ-десять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Р—7, Р—8\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-семь, Эр-восемь" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Светло-розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Светло-розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case DarkPink:
			switch colorsAndQuantity[DarkPink] {
			case 17:
				resp.Text = "От К—1 до О—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-один до О-один" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 16:
				resp.Text = "От Й—2 до П—2, кроме: К—2, Л—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - два до Пэ-два" + breakTimeMiddle +
						"Кроме: Ка-два, Эль-два" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 15:
				resp.Text = "И—3, Й—3, Л—3 и от О—3 до Р—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-три" + breakTimeLong +
						"И-краткое - три" + breakTimeLong +
						"Эль-три" + breakTimeLong +
						"И - от О-три до Эр-три" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 14:
				resp.Text = "От З—4 до К—4, П—4 и Р—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четыре до Ка-четыре" + breakTimeLong +
						"Пэ-четыре" + breakTimeLong +
						"И - Эр-четыре" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 13:
				resp.Text = "З—5, Й—5, К—5 и Р—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-пять" + breakTimeLong +
						"И-краткое - пять" + breakTimeLong +
						"Ка-пять" + breakTimeLong +
						"И - Эр-пять" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 12:
				resp.Text = "Р—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эр-шесть" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 11:
				resp.Text = "П—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-семь" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 10:
				resp.Text = "От Ё—8 до З—8 и П—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-восемь до Зэ-восемь" + breakTimeLong +
						"И - Пэ-восемь" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 9:
				resp.Text = "Г—9, Е—9, Ё—9, И—9, Й—9, О—9 и Р—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-девять" + breakTimeLong +
						"Е-девять" + breakTimeLong +
						"Йо-девять" + breakTimeLong +
						"И-девять" + breakTimeLong +
						"И-краткое - девять" + breakTimeLong +
						"О-девять" + breakTimeLong +
						"И - Эр-девять" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 8:
				resp.Text = "От В—10 до Е—10 и от К—10 до Р—10, кроме: П—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-десять до Е-десять" + breakTimeLong +
						"И - от Ка-десять до Эр-десять" + breakTimeMiddle +
						"Кроме: Пэ-десять" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 7:
				resp.Text = "От Б—11 до Г—11 и от М—11 до П—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-одиннадцать до Гэ-одиннадцать" + breakTimeLong +
						"И - от Эм-одиннадцать до Пэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 6:
				resp.Text = "А—12, от В—12 до Д—12 и от М—12 до П—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двенадцать" + breakTimeLong +
						"От Вэ-двенадцать до Дэ-двенадцать" + breakTimeLong +
						"И - от Эм-двенадцать до Пэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 5:
				resp.Text = "Д—13, Е—13 и от Л—13 до О—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-тринадцать" + breakTimeLong +
						"Е-тринадцать" + breakTimeLong +
						"И - от Эль-тринадцать до О-тринадцать" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 4:
				resp.Text = "Е—14, Ё—14 и от К—14 до Н—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четырнадцать" + breakTimeLong +
						"Йо-четырнадцать" + breakTimeLong +
						"И - от Ка-четырнадцать до Эн-четырнадцать" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 3:
				resp.Text = "От Ё—15 до М—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пятнадцать до Эм-пятнадцать" +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 2:
				resp.Text = "От З—16 до Л—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-шестнадцать до Эль-шестнадцать" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[DarkPink]--
			case 1:
				colorsAndQuantity[DarkPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Й—17, Л—17\nТемно-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - семнадцать, Эль-семнадцать" + breakTimeLong +
						"Темно-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Темно-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Темно-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали темно-розовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали темно-розовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 10:
				resp.Text = "Д—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "От Е—19 до Е—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-девятнадцать до Е-двадцать один" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "От Ё—19 до Ё—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девятнадцать до Йо-двадцать - два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "Ж—18, Ж—19, Ж—22 и Ж—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-восемнадцать" + breakTimeLong +
						"Жэ-девятнадцать" + breakTimeLong +
						"Жэ-двадцать два" + breakTimeLong +
						"И - Жэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "З—17, З—18, З—23 и З—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-семнадцать" + breakTimeLong +
						"Зэ-восемнадцать" + breakTimeLong +
						"Зэ-двадцать три" + breakTimeLong +
						"И - Зэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "И—5, И—17, И—24, И—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-пять" + breakTimeLong +
						"И-семнадцать" + breakTimeLong +
						"И-двадцать четыре" + breakTimeLong +
						"И-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "От Й—18 до Й—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - восемнадцать до И-краткое - двадцать девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "К—17, К—18, К—24, К—28 и К—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-семнадцать" + breakTimeLong +
						"Ка-восемнадцать" + breakTimeLong +
						"Ка-двадцать четыре" + breakTimeLong +
						"Ка-двадцать восемь" + breakTimeLong +
						"И - Ка-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Л—8, от Л—24 до Л—26 и Л—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-восемь" + breakTimeLong +
						"От Эль-двадцать четыре до Эль-двадцать шесть" + breakTimeLong +
						"И - Эль-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—8, от М—25 до М—27 и М—29\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-восемь" + breakTimeLong +
						"От Эм-двадцать пять до Эм-двадцать семь" + breakTimeLong +
						"И - Эм-двадцать девять" + breakTimeLong +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel9:
		switch color {
		case Yellow:
			switch colorsAndQuantity[Yellow] {
			case 24:
				resp.Text = "В—1, Г—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-один, Гэ-один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 23:
				resp.Text = "От Б—2 до Д—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-два до Дэ-два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 22:
				resp.Text = "От В—3 до Д—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-три до Дэ-три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 21:
				resp.Text = "Г—4, Д—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четыре, Дэ-четыре" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 20:
				resp.Text = "Г—5, Д—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-пять, Дэ-пять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 19:
				resp.Text = "От Г—6 до Е—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-шесть до Е-шесть" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 18:
				resp.Text = "От Г—7 до Ё—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-семь до Йо-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 17:
				resp.Text = "От В—8 до Ж—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-восемь до Жэ-восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 16:
				resp.Text = "От Г—9 до Ё—9 и от П—9 до С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-девять до Йо-девять" + breakTimeLong +
						"И - от Пэ-девять до Эс-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 15:
				resp.Text = "От Л—10 до У—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-десять до У-десять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 14:
				resp.Text = "От Й—11 до О—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - одиннадцать до О-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 13:
				resp.Text = "От И—12 до М—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-двенадцать до Эм-двенадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 12:
				resp.Text = "От Ё—13 до Л—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-тринадцать до Эль-тринадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 11:
				resp.Text = "От Ж—14 до К—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-четырнадцать до Ка-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 10:
				resp.Text = "От Ж—15 до К—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-пятнадцать до Ка-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 9:
				resp.Text = "От З—16 до К—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-шестнадцать до Ка-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "От З—17 до Л—17 и от Т—17 до Ч—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-семнадцать до Эль-семнадцать" + breakTimeLong +
						"И - от Тэ-семнадцать до Чэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "От З—18 до М—18 и от Р—18 до Ш—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-восемнадцать до Эм-восемнадцать" + breakTimeLong +
						"И - от Эр-восемнадцать до Ша-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "От И—19 до Ш—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-девятнадцать до Ша-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "От И—20 до Ш—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-двадцать до Ша-двадцать" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "От Й—21 до Ш—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - двадцать один до Ша-двадцать один" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От Й—22 до Ч—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - двадцать два до Чэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "От К—23 до Ц—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать три до Цэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 1:
				colorsAndQuantity[Yellow]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Н—24 до Х—24\nЖелтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-двадцать четыре до Ха-двадцать четыре" + breakTimeLong +
						"Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Желтый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Желтый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали желтый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали желтый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Orange:
			switch colorsAndQuantity[Orange] {
			case 21:
				resp.Text = "Б—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-три" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 20:
				resp.Text = "От А—4 до В—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-четыре до Вэ-четыре" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 19:
				resp.Text = "А—5, В—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пять, Вэ-пять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 18:
				resp.Text = "А—6, В—6 и О—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шесть" + breakTimeLong +
						"Вэ-шесть" + breakTimeLong +
						"И - О-шесть" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 17:
				resp.Text = "От А—7 до В—7, Н—7 и О—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-семь до Вэ-семь" + breakTimeLong +
						"Эн-семь" + breakTimeLong +
						"И - О-семь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 16:
				resp.Text = "Б—8, О—8 и П—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемь" + breakTimeLong +
						"О-восемь" + breakTimeLong +
						"И - Пэ-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 15:
				resp.Text = "Н—9, О—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-девять, О-девять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 14:
				resp.Text = "А—15, Б—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пятнадцать, Бэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 13:
				resp.Text = "А—16, Б—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шестнадцать, Бэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 12:
				resp.Text = "А—17, Б—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-семнадцать, Бэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 11:
				resp.Text = "Б—18, В—18 и Ж—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемнадцать" + breakTimeLong +
						"Вэ-восемнадцать" + breakTimeLong +
						"И - Жэ-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 10:
				resp.Text = "В—19, Г—19, Ё—19 и Ж—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-девятнадцать" + breakTimeLong +
						"Гэ-девятнадцать" + breakTimeLong +
						"Йо-девятнадцать" + breakTimeLong +
						"И - Жэ-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 9:
				resp.Text = "От Г—20 до Ё—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-двадцать до Йо-двадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 8:
				resp.Text = "Д—21, Е—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-двадцать один, Е-двадцать один" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 7:
				resp.Text = "Е—22, Ё—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать два, Йо-двадцать два" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 6:
				resp.Text = "Ё—23, Ж—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать три, Жэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "Ж—24, З—24, Л—24, Щ—24 и Ъ—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать четыре" + breakTimeLong +
						"Зэ-двадцать четыре" + breakTimeLong +
						"Эль-двадцать четыре" + breakTimeLong +
						"Ща-двадцать четыре" + breakTimeLong +
						"И - Твердый-знак - двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "З—25, И—25, Л—25, от Х—25 до Ч—25, Щ—25 и Ъ—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двадцать пять" + breakTimeLong +
						"И-двадцать пять" + breakTimeLong +
						"Эль-двадцать пять" + breakTimeLong +
						"От Ха-двадцать пять до Чэ-двадцать пять" + breakTimeLong +
						"Ща-двадцать пять" + breakTimeLong +
						"И - Твердый-знак - двадцать пять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От И—26 до Л—26, П—26, Р—26 и от Ч—26 до Щ—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-двадцать шесть до Эль-двадцать шесть" + breakTimeLong +
						"Пэ-двадцать шесть" + breakTimeLong +
						"Эр-двадцать шесть" + breakTimeLong +
						"И - от Чэ-двадцать шесть до Ща-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "От К—27 до Р—27, кроме: Н—27, О—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать семь до Эр-двадцать семь" + breakTimeMiddle +
						"Кроме: Эн-двадцать семь, О-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 1:
				colorsAndQuantity[Orange]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От М—28 до П—28\nОранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-двадцать восемь до Пэ-двадцать восемь" + breakTimeLong +
						"Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Оранжевый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Оранжевый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Оранжевый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Оранжевый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 8:
				resp.Text = "От П—11 до Ф—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-одиннадцать до Эф-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 7:
				resp.Text = "От Н—12 до Х—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-двенадцать до Ха-двенадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 6:
				resp.Text = "От М—13 до Ц—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-тринадцать до Цэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "От Л—14 до Ц—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-четырнадцать до Цэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "От Л—15 до Ц—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-пятнадцать до Цэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "От Л—16 до Ц—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-шестнадцать до Цэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "От М—17 до С—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-семнадцать до Эс-семнадцать" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Н—18 до П—18\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-восемнадцать до Пэ-восемнадцать" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали голубой цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали голубой цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 12:
				resp.Text = "Ш—2, Щ—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ша-два, Ща-два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "От Ц—3 до Ъ—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-три до Твердый-знак - три" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "От Ф—4 до Ъ—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-четыре до Твердый-знак - четыре" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "Б—5 и от Т—5 до Щ—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-пять" + breakTimeLong +
						"И - от Тэ-пять до Ща-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 8:
				resp.Text = "Б—6 и от Р—6 до Ц—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-шесть" + breakTimeLong +
						"И - от Эр-шесть до Цэ-шесть" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 7:
				resp.Text = "От П—7 до У—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-семь до У-семь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "От К—8 до Н—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-восемь до Эн-восемь" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 5:
				resp.Text = "От Ж—9 до М—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девять до Эм-девять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 4:
				resp.Text = "От Г—10 до К—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-десять до Ка-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "От В—11 до И—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-одиннадцать до И-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От В—12 до З—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-двенадцать до Зэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Г—13 до Е—13\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-тринадцать до Е-тринадцать" +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	case pictureIDHardLevel10:
		switch color {
		case LightPink:
			switch colorsAndQuantity[LightPink] {
			case 22:
				resp.Text = "И—3, У—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-три, У-три" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 21:
				resp.Text = "И—4, К—4, Л—4, Р—4, С—4 и У—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-четыре" + breakTimeLong +
						"Ка-четыре" + breakTimeLong +
						"Эль-четыре" + breakTimeLong +
						"Эр-четыре" + breakTimeLong +
						"Эс-четыре" + breakTimeLong +
						"И - У-четыре" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 20:
				resp.Text = "З—5, И—5 от К—5 до М—5, от П—5 до С—5, У—5 и Ф—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-пять" + breakTimeLong +
						"И-пять" + breakTimeLong +
						"От Ка-пять до Эм-пять" + breakTimeLong +
						"От Пэ-пять до Эс-пять" + breakTimeLong +
						"У-пять" + breakTimeLong +
						"И - Эф-пять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 19:
				resp.Text = "От И—6 до У—6, кроме: Н—6, О—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-шесть до У-шесть" + breakTimeMiddle +
						"Кроме: Эн-шесть, О-шесть" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 18:
				resp.Text = "От З—7 до Ф—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-семь до Эф-семь" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 17:
				resp.Text = "От З—8 до Ф—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-восемь до Эф-восемь" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 16:
				resp.Text = "От Ж—9 до И—9, от Л—9 до Р—9 и от У—9 до Х—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девять до И-девять" + breakTimeLong +
						"От Эль-девять до Эр-девять" + breakTimeLong +
						"И - от У-девять до Ха-девять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 15:
				resp.Text = "Ж—10, З—10, от М—10 до П—10, Ф—10 и Х—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-десять" + breakTimeLong +
						"Зэ-десять" + breakTimeLong +
						"От Эм-десять до Пэ-десять" + breakTimeLong +
						"Эф-десять" + breakTimeLong +
						"И - Ха-десять" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 14:
				resp.Text = "Ж—11, З—11, Ф—11 и Х—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-одиннадцать" + breakTimeLong +
						"Зэ-одиннадцать" + breakTimeLong +
						"Эф-одиннадцать" + breakTimeLong +
						"И - Ха-одиннадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 13:
				resp.Text = "От Ж—12 до И—12, от Л—12 до Р—12 и от У—12 до Х—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-двенадцать до И-двенадцать" + breakTimeLong +
						"От Эль-двенадцать до Эр-двенадцать" + breakTimeLong +
						"И - от У-двенадцать до Ха-двенадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 12:
				resp.Text = "От Ж—13 до Х—13, кроме: М—13, П—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-тринадцать до Ха-тринадцать" + breakTimeMiddle +
						"Кроме: Эм-тринадцать, Пэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 11:
				resp.Text = "От З—14 до Ф—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-четырнадцать до Эф-четырнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 10:
				resp.Text = "От Й—15 до Т—15, кроме: К—15, С—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - пятнадцать до Тэ-пятнадцать" + breakTimeMiddle +
						"Кроме: Ка-пятнадцать, Эс-пятнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 9:
				resp.Text = "Е—16, от М—16 до П—16 и Ч—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-шестнадцать" + breakTimeLong +
						"От Эм-шестнадцать до Пэ-шестнадцать" + breakTimeLong +
						"И - Чэ-шестнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 8:
				resp.Text = "Е—17, Ё—17, Л—17, Р—17, Ц—17 и Ч—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-семнадцать" + breakTimeLong +
						"Йо-семнадцать" + breakTimeLong +
						"Эль-семнадцать" + breakTimeLong +
						"Эр-семнадцать" + breakTimeLong +
						"Цэ-семнадцать" + breakTimeLong +
						"И - Чэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 7:
				resp.Text = "Б—18, Е—18, Ж—18, от К—18 до С—18, Х—18, Ч—18 и Ы—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемнадцать" + breakTimeLong +
						"Е-восемнадцать" + breakTimeLong +
						"Жэ-восемнадцать" + breakTimeLong +
						"От Ка-восемнадцать до Эс-восемнадцать" + breakTimeLong +
						"Ха-восемнадцать" + breakTimeLong +
						"Чэ-восемнадцать" + breakTimeLong +
						"И - Ы-восемнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 6:
				resp.Text = "Б—19, В—19, от Й—19 до Л—19, от Р—19 до Т—19, Ъ—19 и Ы—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-девятнадцать" + breakTimeLong +
						"Вэ-девятнадцать" + breakTimeLong +
						"От И-краткое - девятнадцать до Эль-девятнадцать" + breakTimeLong +
						"От Эр-девятнадцать до Тэ-девятнадцать" + breakTimeLong +
						"Твердый-знак - девятнадцать" + breakTimeLong +
						"И - Ы-девятнадцать" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 5:
				resp.Text = "Б—20, Г—20, Й—20, К—20, С—20, Т—20, Щ—20 и Ы—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать" + breakTimeLong +
						"Гэ-двадцать" + breakTimeLong +
						"И-краткое - двадцать" + breakTimeLong +
						"Ка-двадцать" + breakTimeLong +
						"Эс-двадцать" + breakTimeLong +
						"Тэ-двадцать" + breakTimeLong +
						"Ща-двадцать" + breakTimeLong +
						"И - Ы-двадцать" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 4:
				resp.Text = "Б—21, от Ё—21 до З—21, Й—21, К—21, С—21, Т—21, от Ф—21 до Ц—21 и Ы—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать один" + breakTimeLong +
						"От Йо-двадцать один до Зэ-двадцать один" + breakTimeLong +
						"И-краткое - двадцать один" + breakTimeLong +
						"Ка-двадцать один" + breakTimeLong +
						"Эс-двадцать один" + breakTimeLong +
						"Тэ-двадцать один" + breakTimeLong +
						"От Эф-двадцать один до Цэ-двадцать один" + breakTimeLong +
						"И - Ы-двадцать один" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 3:
				resp.Text = "Б—22, Ё—22, И—22, К—22, С—22, У—22, Ц—22 и от Ш—22 до Ъ—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать два" + breakTimeLong +
						"Йо-двадцать два" + breakTimeLong +
						"И-двадцать два" + breakTimeLong +
						"Ка-двадцать два" + breakTimeLong +
						"Эс-двадцать два" + breakTimeLong +
						"У-двадцать два" + breakTimeLong +
						"Цэ-двадцать два" + breakTimeLong +
						"И - от Ша-двадцать два до Твердый-знак двадцать два" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 2:
				resp.Text = "Ё—23, К—23, Л—23, Р—23, С—23, Ц—23 и от Ш—23 до Ь—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать три" + breakTimeLong +
						"Ка-двадцать три" + breakTimeLong +
						"Эль-двадцать три" + breakTimeLong +
						"Эр-двадцать три" + breakTimeLong +
						"Эс-двадцать три" + breakTimeLong +
						"Цэ-двадцать три" + breakTimeLong +
						"И - от Ша-двадцать три до Мягкий-знак - двадцать три" +
						"</speak>"
				colorsAndQuantity[LightPink]--
			case 1:
				colorsAndQuantity[LightPink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От К—24 до С—24, кроме: Н—24, О—24\nСветло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-двадцать четыре до Эс-двадцать четыре" + breakTimeMiddle +
						"Кроме: Эн-двадцать четыре, О-двадцать четыре" + breakTimeLong +
						"Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Светло-розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Светло-розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали светло-розовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали светло-розовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 24:
				resp.Text = "З—2, Н—2, О—2 и Ф—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-два" + breakTimeLong +
						"Эн-два" + breakTimeLong +
						"О-два" + breakTimeLong +
						"И - Эф-два" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 23:
				resp.Text = "З—3, Н—3, О—3 и Ф—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-три" + breakTimeLong +
						"Эн-три" + breakTimeLong +
						"О-три" + breakTimeLong +
						"И - Эф-три" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 22:
				resp.Text = "Ж—4, З—4, Н—4, О—4, Ф—4 и Х—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-четыре" + breakTimeLong +
						"Зэ-четыре" + breakTimeLong +
						"Эн-четыре" + breakTimeLong +
						"О-четыре" + breakTimeLong +
						"Эф-четыре" + breakTimeLong +
						"И - Ха-четыре" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 21:
				resp.Text = "Ж—5, Х—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-пять, Ха-пять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 20:
				resp.Text = "Ж—6, Х—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-шесть, Ха-шесть" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 19:
				resp.Text = "Б—8, В—8, Ъ—8 и Ы—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемь" + breakTimeLong +
						"Вэ-восемь" + breakTimeLong +
						"Твердый-знак - восемь" + breakTimeLong +
						"И - Ы-восемь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 18:
				resp.Text = "В—9, Г—9, Щ—9 и Ъ—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-девять" + breakTimeLong +
						"Гэ-девять" + breakTimeLong +
						"Ща-девять" + breakTimeLong +
						"И - Твердый-знак - девять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 17:
				resp.Text = "От В—10 до Д—10 и от Ш—10 до Ъ—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-десять до Дэ-десять" + breakTimeLong +
						"И - От Ша-десять до Твердый-знак - десять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 16:
				resp.Text = "А—11, Б—11, Г—11, Д—11, Ш—11, Щ—11, Ы—11 и Ь—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-одиннадцать" + breakTimeLong +
						"Бэ-одиннадцать" + breakTimeLong +
						"Гэ-одиннадцать" + breakTimeLong +
						"Дэ-одиннадцать" + breakTimeLong +
						"Ша-одиннадцать" + breakTimeLong +
						"Ща-одиннадцать" + breakTimeLong +
						"Ы-одиннадцать" + breakTimeLong +
						"И - Мягкий-знак - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 15:
				resp.Text = "В—12, Д—12, Ш—12 и Ъ—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двенадцать" + breakTimeLong +
						"Дэ-двенадцать" + breakTimeLong +
						"Ша-двенадцать" + breakTimeLong +
						"И - Твердый-знак - двенадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 14:
				resp.Text = "А—13, Б—13, Д—13, Ш—13, Ы—13 и Ь—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-тринадцать" + breakTimeLong +
						"Бэ-тринадцать" + breakTimeLong +
						"Дэ-тринадцать" + breakTimeLong +
						"Ша-тринадцать" + breakTimeLong +
						"Ы-тринадцать" + breakTimeLong +
						"И - Мягкий-знак - тринадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 13:
				resp.Text = "А—14, Б—14, Г—14, Е—14, Ч—14, Щ—14, Ы—14 и Ь—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-четырнадцать" + breakTimeLong +
						"Бэ-четырнадцать" + breakTimeLong +
						"Гэ-четырнадцать" + breakTimeLong +
						"Е-четырнадцать" + breakTimeLong +
						"Чэ-четырнадцать" + breakTimeLong +
						"Ща-четырнадцать" + breakTimeLong +
						"Ы-четырнадцать" + breakTimeLong +
						"И - Мягкий-знак - четырнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 12:
				resp.Text = "А—15, Б—15, Г—15, Ё—15, Ц—15, Щ—15, Ы—15 и Ь—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пятнадцать" + breakTimeLong +
						"Бэ-пятнадцать" + breakTimeLong +
						"Гэ-пятнадцать" + breakTimeLong +
						"Йо-пятнадцать" + breakTimeLong +
						"Цэ-пятнадцать" + breakTimeLong +
						"Ща-пятнадцать" + breakTimeLong +
						"Ы-пятнадцать" + breakTimeLong +
						"И - Мягкий-знак - пятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 11:
				resp.Text = "А—16, Г—16, Ж—16, Х—16, Щ—16 и Ь—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-шестнадцать" + breakTimeLong +
						"Гэ-шестнадцать" + breakTimeLong +
						"Жэ-шестнадцать" + breakTimeLong +
						"Ха-шестнадцать" + breakTimeLong +
						"Ща-шестнадцать" + breakTimeLong +
						"И - Мягкий-знак - шестнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 10:
				resp.Text = "А—17, Г—17, от З—17 до Й—17, от Т—17 до Ф—17, Щ—17 и Ь—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-семнадцать" + breakTimeLong +
						"Гэ-семнадцать" + breakTimeLong +
						"От Зэ-семнадцать до И-краткое - семнадцать" + breakTimeLong +
						"От Тэ-семнадцать до Эф-семнадцать" + breakTimeLong +
						"Ща-семнадцать" + breakTimeLong +
						"И - Мягкий-знак - семнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 9:
				resp.Text = "А—18, Г—18, Ё—18, Ц—18, Щ—18 и Ь—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-восемнадцать" + breakTimeLong +
						"Гэ-восемнадцать" + breakTimeLong +
						"Йо-восемнадцать" + breakTimeLong +
						"Цэ-восемнадцать" + breakTimeLong +
						"Ща-восемнадцать" + breakTimeLong +
						"И - Мягкий-знак - восемнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 8:
				resp.Text = "А—19, Д—19, Ж—19, Н—19, О—19, Х—19, Ш—19 и Ь—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-девятнадцать" + breakTimeLong +
						"Дэ-девятнадцать" + breakTimeLong +
						"Жэ-девятнадцать" + breakTimeLong +
						"Эн-девятнадцать" + breakTimeLong +
						"О-девятнадцать" + breakTimeLong +
						"Ха-девятнадцать" + breakTimeLong +
						"Ша-девятнадцать" + breakTimeLong +
						"И - Мягкий-знак - девятнадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 7:
				resp.Text = "А—20, В—20, от М—20 до П—20, Ъ—20 и Ь—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать" + breakTimeLong +
						"Вэ-двадцать" + breakTimeLong +
						"От Эм-двадцать до Пэ-двадцать" + breakTimeLong +
						"Твердый-знак - двадцать" + breakTimeLong +
						"И - Мягкий-знак - двадцать" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 6:
				resp.Text = "А—21, В—21, от М—21 до П—21 и Ь—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать один" + breakTimeLong +
						"Вэ-двадцать один" + breakTimeLong +
						"От Эм-двадцать один до Пэ-двадцать один" + breakTimeLong +
						"И - Мягкий-знак - двадцать один" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "А—22, В—22, Ж—22, З—22, Л—22, Р—22, Ф—22 и Х—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать два" + breakTimeLong +
						"Вэ-двадцать два" + breakTimeLong +
						"Жэ-двадцать два" + breakTimeLong +
						"Зэ-двадцать два" + breakTimeLong +
						"Эль-двадцать два" + breakTimeLong +
						"Эр-двадцать два" + breakTimeLong +
						"Эф-двадцать два" + breakTimeLong +
						"И - Ха-двадцать два" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "А—23, от Ж—23 до И—23, М—23, П—23 и от У—23 до Х—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать три" + breakTimeLong +
						"От Жэ-двадцать три до И-двадцать три" + breakTimeLong +
						"Эм-двадцать три" + breakTimeLong +
						"Пэ-двадцать три" + breakTimeLong +
						"И - от У-двадцать три до Ха-двадцать три" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "А—24, от Ж—24 до И—24 и от У—24 до Ь—24, кроме: Ц—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двадцать четыре" + breakTimeLong +
						"От Жэ-двадцать четыре до И-двадцать четыре" + breakTimeLong +
						"И - от У-двадцать четыре до Мягкий-знак - двадцать четыре" + breakTimeMiddle +
						"Кроме: Цэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "Б—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "В—26, Ъ—26\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать шесть, Твердый-знак - двадцать шесть" + breakTimeLong +
						"Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Голубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Голубой цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Голубой цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Голубой цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Purple:
			switch colorsAndQuantity[Purple] {
			case 25:
				resp.Text = "З—1, Н—1, О—1 и Ф—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-один" + breakTimeLong +
						"Эн-один" + breakTimeLong +
						"О-один" + breakTimeLong +
						"И - Эф-один" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 24:
				resp.Text = "Ж—2, И—2, М—2, П—2, У—2 и Х—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-два" + breakTimeLong +
						"И-два" + breakTimeLong +
						"Эм-два" + breakTimeLong +
						"Пэ-два" + breakTimeLong +
						"У-два" + breakTimeLong +
						"И - Ха-два" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 23:
				resp.Text = "Ж—3, от Й—3 до М—3, от П—3 до Т—3 и Х—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"От И-краткое - три до Эм-три" + breakTimeLong +
						"От Пэ-три до Тэ-три" + breakTimeLong +
						"И - Ха-три" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 22:
				resp.Text = "Ё—4, Й—4, М—4, П—4, Т—4 и Ц—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-четыре" + breakTimeLong +
						"И-краткое - четыре" + breakTimeLong +
						"Эм-четыре" + breakTimeLong +
						"Пэ-четыре" + breakTimeLong +
						"Тэ-четыре" + breakTimeLong +
						"И - Цэ-четыре" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 21:
				resp.Text = "Ё—5, Й—5, Н—5, О—5, Т—5 и Ц—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-пять" + breakTimeLong +
						"И-краткое - пять" + breakTimeLong +
						"Эн-пять" + breakTimeLong +
						"О-пять" + breakTimeLong +
						"Тэ-пять" + breakTimeLong +
						"И - Цэ-пять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 20:
				resp.Text = "Ё—6, З—6, Н—6, О—6, Ф—6 и Ц—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-шесть" + breakTimeLong +
						"Зэ-шесть" + breakTimeLong +
						"Эн-шесть" + breakTimeLong +
						"О-шесть" + breakTimeLong +
						"Эф-шесть" + breakTimeLong +
						"И - Цэ-шесть" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 19:
				resp.Text = "Ё—7, Ж—7, Х—7 и Ц—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-семь" + breakTimeLong +
						"Жэ-семь" + breakTimeLong +
						"Ха-семь" + breakTimeLong +
						"И - Цэ-семь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 18:
				resp.Text = "Ж—8, Х—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-восемь, Ха-восемь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 17:
				resp.Text = "Ё—9, Й—9, К—9, С—9, Т—9 и Ц—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-девять" + breakTimeLong +
						"И-краткое - девять" + breakTimeLong +
						"Ка-девять" + breakTimeLong +
						"Эс-девять" + breakTimeLong +
						"Тэ-девять" + breakTimeLong +
						"И - Цэ-девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 16:
				resp.Text = "Ё—10, И—10, Л—10, Р—10, У—10 и Ц—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-десять" + breakTimeLong +
						"И-десять" + breakTimeLong +
						"Эль-десять" + breakTimeLong +
						"Эр-десять" + breakTimeLong +
						"У-десять" + breakTimeLong +
						"И - Цэ-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 15:
				resp.Text = "В—11, Ё—11, И—11, от Л—11 до Р—11, У—11, Ц—11 и Ъ—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" + breakTimeLong +
						"Йо-одиннадцать" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"От Эль-одиннадцать до Эр-одиннадцать" + breakTimeLong +
						"У-одиннадцать" + breakTimeLong +
						"Цэ-одиннадцать" + breakTimeLong +
						"И - Твердый-знак - одиннадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 14:
				resp.Text = "А—12, Б—12, Г—12, Ё—12, Й—12, К—12, С—12, Т—12, Ц—12, Щ—12, Ы—12 и Ь—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двенадцать" + breakTimeLong +
						"Бэ-двенадцать" + breakTimeLong +
						"Гэ-двенадцать" + breakTimeLong +
						"Йо-двенадцать" + breakTimeLong +
						"И-краткое - двенадцать" + breakTimeLong +
						"Ка-двенадцать" + breakTimeLong +
						"Эс-двенадцать" + breakTimeLong +
						"Тэ-двенадцать" + breakTimeLong +
						"Цэ-двенадцать" + breakTimeLong +
						"Ща-двенадцать" + breakTimeLong +
						"Ы-двенадцать" + breakTimeLong +
						"И - Мягкий-знак - двенадцать" + breakTimeLong +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 13:
				resp.Text = "В—13, Г—13, Ё—13, Ц—13, Щ—13 и Ъ—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-тринадцать" + breakTimeLong +
						"Гэ-тринадцать" + breakTimeLong +
						"Йо-тринадцать" + breakTimeLong +
						"Цэ-тринадцать" + breakTimeLong +
						"Ща-тринадцать" + breakTimeLong +
						"И - Твердый-знак - тринадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 12:
				resp.Text = "В—14, Д—14, Ж—14, Х—14, Ш—14 и Ъ—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четырнадцать" + breakTimeLong +
						"Дэ-четырнадцать" + breakTimeLong +
						"Жэ-четырнадцать" + breakTimeLong +
						"Ха-четырнадцать" + breakTimeLong +
						"Ша-четырнадцать" + breakTimeLong +
						"И - Твердый-знак - четырнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 11:
				resp.Text = "В—15, Д—15, Е—15, З—15, И—15, К—15, С—15, У—15, Ф—15, Ч—15, Ш—15 и Ъ—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-пятнадцать" + breakTimeLong +
						"Дэ-пятнадцать" + breakTimeLong +
						"Е-пятнадцать" + breakTimeLong +
						"Зэ-пятнадцать" + breakTimeLong +
						"И-пятнадцать" + breakTimeLong +
						"Ка-пятнадцать" + breakTimeLong +
						"Эс-пятнадцать" + breakTimeLong +
						"У-пятнадцать" + breakTimeLong +
						"Эф-пятнадцать" + breakTimeLong +
						"Чэ-пятнадцать" + breakTimeLong +
						"Ша-пятнадцать" + breakTimeLong +
						"И - Твердый-знак - пятнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 10:
				resp.Text = "Б—16, В—16, Д—16, Ё—16, от Й—16 до Л—16, от Р—16 до Т—16, Ц—16, Ш—16, Ъ—16 и Ы—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-шестнадцать" + breakTimeLong +
						"Вэ-шестнадцать" + breakTimeLong +
						"Дэ-шестнадцать" + breakTimeLong +
						"Йо-шестнадцать" + breakTimeLong +
						"От И-краткое-шестнадцать до Эль-шестнадцать" + breakTimeLong +
						"От Эр-шестнадцать до Тэ-шестнадцать" + breakTimeLong +
						"Цэ-шестнадцать" + breakTimeLong +
						"Ша-шестнадцать" + breakTimeLong +
						"Твердый-знак - шестнадцать" + breakTimeLong +
						"И - Ы-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 9:
				resp.Text = "Б—17, В—17, Д—17, Ж—17, К—17, от М—17 до П—17, С—17, Х—17, Ш—17, Ъ—17 и Ы—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-семнадцать" + breakTimeLong +
						"Вэ-семнадцать" + breakTimeLong +
						"Дэ-семнадцать" + breakTimeLong +
						"Жэ-семнадцать" + breakTimeLong +
						"Ка-семнадцать" + breakTimeLong +
						"От Эм-семнадцать до Пэ-семнадцать" + breakTimeLong +
						"Эс-семнадцать" + breakTimeLong +
						"Ха-семнадцать" + breakTimeLong +
						"Ша-семнадцать" + breakTimeLong +
						"Твердый-знак - семнадцать" + breakTimeLong +
						"И - Ы-семнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 8:
				resp.Text = "В—18, Д—18, от З—18 до Й—18, от Т—18 до Ф—18, Ш—18 и Ъ—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-восемнадцать" + breakTimeLong +
						"Дэ-восемнадцать" + breakTimeLong +
						"От Зэ-восемнадцать до И-краткое - восемнадцать" + breakTimeLong +
						"От Тэ-восемнадцать до Эф-восемнадцать" + breakTimeLong +
						"Ша-восемнадцать" + breakTimeLong +
						"И - Твердый-знак - восемнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 7:
				resp.Text = "Г—19, З—19, И—19, М—19, П—19, У—19, Ф—19 и Щ—19"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-девятнадцать" + breakTimeLong +
						"Зэ-девятнадцать" + breakTimeLong +
						"И-девятнадцать" + breakTimeLong +
						"Эм-девятнадцать" + breakTimeLong +
						"Пэ-девятнадцать" + breakTimeLong +
						"У-девятнадцать" + breakTimeLong +
						"Эф-девятнадцать" + breakTimeLong +
						"И - Ща-девятнадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 6:
				resp.Text = "От Е—20 до И—20, Л—20, Р—20 и от У—20 до Ч—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-двадцать до И-двадцать" + breakTimeLong +
						"Эль-двадцать" + breakTimeLong +
						"Эр-двадцать" + breakTimeLong +
						"И - от У-двадцать до Чэ-двадцать" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 5:
				resp.Text = "Е—21, И—21, Л—21, Р—21, У—21 и от Ч—21 до Ъ—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать один" + breakTimeLong +
						"И-двадцать один" + breakTimeLong +
						"Эль-двадцать один" + breakTimeLong +
						"Эр-двадцать один" + breakTimeLong +
						"У-двадцать один" + breakTimeLong +
						"И - от Чэ-двадцать один до Твердый-знак - двадцать один" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 4:
				resp.Text = "Е—22, Й—22, от М—22 до П—22, Т—22, Ч—22, Ы—22 и Ь—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-двадцать два" + breakTimeLong +
						"И-краткое - двадцать два" + breakTimeLong +
						"От Эм-двадцать два до Пэ-двадцать два" + breakTimeLong +
						"Тэ-двадцать два" + breakTimeLong +
						"Чэ-двадцать два" + breakTimeLong +
						"Ы-двадцать два" + breakTimeLong +
						"И - Мягкий-знак - двадцать два" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 3:
				resp.Text = "Б—23, Е—23, Й—23, Н—23, О—23, Т—23 и Ч—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать три" + breakTimeLong +
						"Е-двадцать три" + breakTimeLong +
						"И-краткое - двадцать три" + breakTimeLong +
						"Эн-двадцать три" + breakTimeLong +
						"О-двадцать три" + breakTimeLong +
						"Тэ-двадцать три" + breakTimeLong +
						"И - Чэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 2:
				resp.Text = "Б—24, Ё—24, Й—24, Н—24, О—24, Т—24 и Ц—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двадцать четыре" + breakTimeLong +
						"Йо-двадцать четыре" + breakTimeLong +
						"И-краткое - двадцать четыре" + breakTimeLong +
						"Эн-двадцать четыре" + breakTimeLong +
						"О-двадцать четыре" + breakTimeLong +
						"Тэ-двадцать четыре" + breakTimeLong +
						"И - Цэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 1:
				colorsAndQuantity[Purple]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "В—25 и от Ж—25 до Ь—25\nФиолетовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-двадцать пять" + breakTimeLong +
						"И - от Жэ-двадцать пять до Мягкий-знак - двадцать пять" + breakTimeLong +
						"Фиолетовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Фиолетовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Фиолетовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали фиолетовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали фиолетовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 3:
				resp.Text = "К—10, Т—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-десять, Тэ-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "Й—11, К—11, С—11 и Т—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - одиннадцать" + breakTimeLong +
						"Ка-одиннадцать" + breakTimeLong +
						"Эс-одиннадцать" + breakTimeLong +
						"И - Тэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "М—13, П—13\nЧерный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-тринадцать, Пэ-тринадцать" +
						"Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже рисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже рисовали черный цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		default:
			resp = processingDefaultColor(resp)
		}
	}
	return resp
}

func processingDefaultColor(resp Response) Response {
	var getRemainingColors = getRemainingColors()
	var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
	resp.Text = "Выберите " + getRemainingColors + "цвет."
	resp.TTSType = "ssml"
	resp.SSML =
		"<speak>" +
			"Выберите " + getRemainingColorsForSpeak + "цвет." +
			"</speak>"
	return resp
}

func mixedLevels(resp Response) Response {
	resp = easyLevels(resp)
	if resp.Text == "" {
		resp = mediumLevels(resp)
		if resp.Text == "" {
			resp = hardLevels(resp)
		}
	}
	return resp
}