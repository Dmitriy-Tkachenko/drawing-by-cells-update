package main

import (
	"sort"
	"strings"
)

const (
	Black = "черный"
	Brown = "коричневый"
	Darkblue = "синий"
	Blue = "голубой"
	Purple = "фиолетовый"
	Red = "красный"
	Pink = "розовый"
	Orange = "оранжевый"
	Yellow = "желтый"
	Green = "зеленый"
	Beige = "бежевый"
	Gray = "серый"
	White = "белый"
)

const (
	hintIDEasyLevel1 = 457239065
	pictureIDEasyLevel1 = 457239066
	hintIDEasyLevel2 = 457239067
	pictureIDEasyLevel2 = 457239068
	hintIDEasyLevel3 = 457239069
	pictureIDEasyLevel3 = 457239070
	hintIDEasyLevel4 = 457239071
	pictureIDEasyLevel4 = 457239072
	hintIDEasyLevel5 = 457239073
	pictureIDEasyLevel5 = 457239074
	/*hintIDEasyLevel6 = 457239075
	pictureIDEasyLevel6 = 457239076
	hintIDEasyLevel7 = 457239077
	pictureIDEasyLevel7 = 457239078
	hintIDEasyLevel8 = 457239079
	pictureIDEasyLevel8 = 457239080
	hintIDEasyLevel9 = 457239081
	pictureIDEasyLevel9 = 457239082
	hintIDEasyLevel10 = 457239083
	pictureIDEasyLevel10 = 457239084
	hintIDEasyLevel11 = 457239085
	pictureIDEasyLevel11 = 457239086
	hintIDEasyLevel12 = 457239087
	pictureIDEasyLevel12 = 457239088
	hintIDEasyLevel13 = 457239089
	pictureIDEasyLevel13 = 457239090
	hintIDEasyLevel14 = 457239091
	pictureIDEasyLevel14 = 457239092
	hintIDEasyLevel15 = 457239093
	pictureIDEasyLevel15 = 457239094*/
)

const (
	hintIDMediumLevel1 = 457239095
	pictureIDMediumLevel1 = 457239096
	hintIDMediumLevel2 = 457239097
	pictureIDMediumLevel2 = 457239098
	hintIDMediumLevel3 = 457239099
	pictureIDMediumLevel3 = 457239100
	hintIDMediumLevel4 = 457239101
	pictureIDMediumLevel4 = 457239102
	hintIDMediumLevel5 = 457239103
	pictureIDMediumLevel5 = 457239104
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
)

var hintID, pictureID, numbOfColoringRules int
var breakTimeShort = "<break time=\"0.005s\"/>"
var breakTimeMiddle = "<break time=\"0.5s\"/>"
var breakTimeLong = "<break time=\"2s\"/>"
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
	resp.Text = "Давайте нарисуем кита. Для этого Вам необходимо 4 цвета: " + Pink + ", " + Purple + ", " + Darkblue + " и " + Black + ". С какого цвета Вы хотите начать?"
	resp.TTS = "Давайте нарисуем кита. Для этого Вам необходимо четыре цв+ета\n" + Pink + "\n" + Purple + "\n" + Darkblue + " - и " + Black + ". С какого цвета Вы хотите начать?"

	resultText = "Поздравляю! Кит нарисован."
	pictureType = "Образец для рисования {кита}{кит`а}."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Pink] = 14; colorsAndQuantity[Purple] = 8; colorsAndQuantity[Darkblue] = 7; colorsAndQuantity[Black] = 1
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

/*func easyLevel6(resp Response) Response {
	resp.Text = "Нарисуем божью коровку. Для этого вам понадобятся 2 цвета: " + Red + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем божью коровку. Для этого вам понадобятся два цвета: " + Red + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Red), Payload {
		Text: Red,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel6
	pictureID = pictureIDEasyLevel6
	return resp
}

func easyLevel7(resp Response) Response {
	resp.Text = "Нарисуем черепаху. Для этого вам понадобятся 3 цвета: " + Yellow + ", " + Green + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем черепаху. Для этого вам понадобятся три цвета: " + Yellow + ", " + Green + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Yellow), Payload {
		Text: Yellow,
	})
	resp.AddButton(strings.Title(Green), Payload {
		Text: Green,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel7
	pictureID = pictureIDEasyLevel7
	return resp
}

func easyLevel8(resp Response) Response {
	resp.Text = "Нарисуем краба. Для этого вам понадобятся 3 цвета: " + Yellow + ", " + Orange + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем краба. Для этого вам понадобятся три цвета: " + Yellow + ", " + Orange + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Yellow), Payload {
		Text: Yellow,
	})
	resp.AddButton(strings.Title(Orange), Payload {
		Text: Orange,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel8
	pictureID = pictureIDEasyLevel8
	return resp
}

func easyLevel9(resp Response) Response {
	resp.Text = "Нарисуем жирафа. Для этого вам понадобятся 3 цвета: " + Orange + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем жирафа. Для этого вам понадобятся три цвета: " + Orange + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Orange), Payload {
		Text: Orange,
	})
	resp.AddButton(strings.Title(Brown), Payload {
		Text: Brown,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel9
	pictureID = pictureIDEasyLevel9
	return resp
}

func easyLevel10(resp Response) Response {
	resp.Text = "Нарисуем паровоз. Для этого вам понадобятся 5 цветов: " + Orange + ", " + Red + ", " + Green + ", " + Blue + " и " + Darkblue + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем паровоз. Для этого вам понадобятся пять цветов: " + Orange + ", " + Red + ", " + Green + ", " + Blue + " и " + Darkblue + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Orange), Payload {
		Text: Orange,
	})
	resp.AddButton(strings.Title(Red), Payload {
		Text: Red,
	})
	resp.AddButton(strings.Title(Green), Payload {
		Text: Green,
	})
	resp.AddButton(strings.Title(Blue), Payload {
		Text: Blue,
	})
	resp.AddButton(strings.Title(Darkblue), Payload {
		Text: Darkblue,
	})
	hintID = hintIDEasyLevel10
	pictureID = pictureIDEasyLevel10
	return resp
}

func easyLevel11(resp Response) Response {
	resp.Text = "Нарисуем кота. Для этого вам понадобятся 2 цвета: " + Orange + " и " + Red + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем кота. Для этого вам понадобятся два цвета: " + Orange + " и " + Red + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Orange), Payload {
		Text: Orange,
	})
	resp.AddButton(strings.Title(Red), Payload {
		Text: Red,
	})
	hintID = hintIDEasyLevel11
	pictureID = pictureIDEasyLevel11
	return resp
}

func easyLevel12(resp Response) Response {
	resp.Text = "Нарисуем елку. Для этого вам понадобятся 7 цветов: " + Yellow + ", " + Pink + ", " + Red + ", " + Green + ", " + Darkblue + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем елку. Для этого вам понадобятся семь цветов: " + Yellow + ", " + Pink + ", " + Red + ", " + Green + ", " + Darkblue + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Yellow), Payload {
		Text: Yellow,
	})
	resp.AddButton(strings.Title(Pink), Payload {
		Text: Pink,
	})
	resp.AddButton(strings.Title(Red), Payload {
		Text: Red,
	})
	resp.AddButton(strings.Title(Green), Payload {
		Text: Green,
	})
	resp.AddButton(strings.Title(Darkblue), Payload {
		Text: Darkblue,
	})
	resp.AddButton(strings.Title(Brown), Payload {
		Text: Brown,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel12
	pictureID = pictureIDEasyLevel12
	return resp
}

func easyLevel13(resp Response) Response {
	resp.Text = "Нарисуем машину. Для этого вам понадобятся 5 цветов: " + Yellow + ", " + Orange + ", " + Red + ", " + Darkblue + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем машину. Для этого вам понадобятся пять цветов: " + Yellow + ", " + Orange + ", " + Red + ", " + Darkblue + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Yellow), Payload {
		Text: Yellow,
	})
	resp.AddButton(strings.Title(Orange), Payload {
		Text: Orange,
	})
	resp.AddButton(strings.Title(Red), Payload {
		Text: Red,
	})
	resp.AddButton(strings.Title(Darkblue), Payload {
		Text: Darkblue,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel13
	pictureID = pictureIDEasyLevel13
	return resp
}

func easyLevel14(resp Response) Response {
	resp.Text = "Нарисуем собаку. Для этого вам понадобятся 2 цвета: " + Orange + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем собаку. Для этого вам понадобятся два цвета: " + Orange + " и " + Black + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Orange), Payload {
		Text: Orange,
	})
	resp.AddButton(strings.Title(Black), Payload {
		Text: Black,
	})
	hintID = hintIDEasyLevel14
	pictureID = pictureIDEasyLevel14
	return resp
}

func easyLevel15(resp Response) Response {
	resp.Text = "Нарисуем мороженое. Для этого вам понадобятся 5 цветов: " + Pink + ", " + Red + ", " + Blue + ", " + Darkblue + " и " + Brown + ". С какого цвета начнем?"
	resp.TTS = "Нарисуем мороженое. Для этого вам понадобятся пять цветов: " + Pink + ", " + Red + ", " + Blue + ", " + Darkblue + " и " + Brown + ". С какого цвета начнем?"
	resp.AddButton(strings.Title(Pink), Payload {
		Text: Pink,
	})
	resp.AddButton(strings.Title(Red), Payload {
		Text: Red,
	})
	resp.AddButton(strings.Title(Blue), Payload {
		Text: Blue,
	})
	resp.AddButton(strings.Title(Darkblue), Payload {
		Text: Darkblue,
	})
	resp.AddButton(strings.Title(Brown), Payload {
		Text: Brown,
	})
	hintID = hintIDEasyLevel15
	pictureID = pictureIDEasyLevel15
	return resp
}*/

func mediumLevel1(resp Response) Response {
	resp.Text = "Нарисуем морскую свинку. Для этого Вам нужно подготовить 3 цвета: " + Pink + ", " + Brown + " и " + Black + ". С какого цвета начнем рисовать?"
	resp.TTS = "Нарисуем морскую свинку. Для этого Вам нужно подготовить три цв+ета\n" + Pink + "\n" + Brown + " - и " + Black + ". С какого цвета начнем рисовать?"

	resultText = "Поздравляю! Вы нарисовали морскую свинку."
	pictureType = "Образец для рисования морской свинки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Pink] = 3; colorsAndQuantity[Brown] = 24; colorsAndQuantity[Black] = 14
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
	resp.Text = "Сейчас я научу Вас рисовать вертолет. Для этого Вам понадобятся 5 цветов: " + Red + ", " + Blue + ", " + Darkblue + ", " + Brown + " и " + Black + ". Какой цвет хотите выбрать?"
	resp.TTS = "Сейчас я научу Вас рисовать вертолет. Для этого Вам понадобятся пять цвет+ов\n" + Red + "\n" + Blue + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". Какой цвет хотите выбрать?"

	resultText = "Поздравляю с нарисованным вертолетом!"
	pictureType = "Образец для рисования вертолета."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Red] = 11; colorsAndQuantity[Blue] = 6; colorsAndQuantity[Darkblue] = 4; colorsAndQuantity[Brown] = 2; colorsAndQuantity[Black] = 4
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

	resultText = "Поздравляю! Вы нарисовали кита."
	pictureType = "Образец для рисования синего кита."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Blue] = 12; colorsAndQuantity[Darkblue] = 8; colorsAndQuantity[Black] = 12
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
	resp.Text = "Нарисуем персонажа «Миньон» из мультфильма «Гадкий я». Перед рисованием подготовьте 6 цветов: " + Gray + ", " + Yellow + ", " + Red + ", " + Darkblue + ", " + Brown + " и " + Black + ". Какой цвет выберем?"
	resp.TTS = "Нарисуем персонажа - Миньон - из мультфильма «Гадкий я». Перед рисованием подготовьте шесть цвет+ов\n" + Gray + "\n" + Yellow + "\n" + Red + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". Какой цвет выберем?"

	resultText = "Миньон нарисован. Поздравляю!"
	pictureType = "Образец для рисования миньона."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Gray] = 8; colorsAndQuantity[Yellow] = 10; colorsAndQuantity[Red] = 2; colorsAndQuantity[Darkblue] = 6; colorsAndQuantity[Brown] = 1; colorsAndQuantity[Black] = 11
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

func hardLevel1(resp Response) Response {
	resp.Text = "Давайте нарисуем девушку Моана из мультфильма «Моана». Перед тем как приступить к рисованию Вам необходимо подготовить 8 цветов: " + Beige + ", " + Yellow + ", " + Pink + ", " + Red + ", " + Blue + ", " + Darkblue + ", " + Brown + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Давайте нарисуем девушку - Моана - из мультфильма «Моана». Перед тем как приступить к рисованию Вам необходимо подготовить восемь цветов\n" + Beige + "\n" + Yellow + "\n" + Pink + "\n" + Red + "\n" + Blue + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Моана нарисована. Поздравляю!"
	pictureType = "Образец для рисования девушки Моана."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 17; colorsAndQuantity[Yellow] = 3; colorsAndQuantity[Pink] = 2; colorsAndQuantity[Red] = 8; colorsAndQuantity[Blue] = 1; colorsAndQuantity[Darkblue] = 2; colorsAndQuantity[Brown] = 9; colorsAndQuantity[Black] = 22
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
	resp.Text = "Нарисуем девушку Мерида из мультфильма «Храбрая сердцем». Для этого вам понадобятся 8 цветов: " + Beige + ", " + Yellow + ", " + Orange + ", " + Pink + ", " + Blue + ", " + Darkblue + ", " + Brown + " и " + Black + ". Какой цвет выберем первым?"
	resp.TTS = "Нарисуем девушку - Мерида - из мультфильма «Храбрая сердцем». Для этого вам понадобятся восемь цветов\n" + Beige + "\n" + Yellow + "\n" + Orange + "\n" + Pink + "\n" + Blue + "\n" + Darkblue + "\n" + Brown + " - и " + Black + ". Какой цвет выберем первым?"

	resultText = "Мерида нарисована. Поздравляю!"
	pictureType = "Образец для рисования девушки Мерида."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 12; colorsAndQuantity[Yellow] = 7; colorsAndQuantity[Orange] = 19; colorsAndQuantity[Pink] = 2; colorsAndQuantity[Blue] = 1; colorsAndQuantity[Darkblue] = 15; colorsAndQuantity[Brown] = 10; colorsAndQuantity[Black] = 7
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
	resp.Text = "Научимся рисовать собачку. Перед рисованием вам нужно подготовить 6 цветов: " + Beige + ", " + Orange + ", " + Pink + ", " + Red + ", " + Brown + " и " + Black + ". Какой цвет будем рисовать первым?"
	resp.TTS = "Научимся рисовать собачку. Перед рисованием вам нужно подготовить шесть цветов\n" + Beige + "\n" + Orange + "\n" + Pink + "\n" + Red + "\n" + Brown + " - и " + Black + ". Какой цвет будем рисовать первым?"

	resultText = "Собачка нарисована. Поздравляю!"
	pictureType = "Образец для рисования собачки."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Beige] = 8; colorsAndQuantity[Orange] = 13; colorsAndQuantity[Pink] = 3; colorsAndQuantity[Red] = 1; colorsAndQuantity[Brown] = 19; colorsAndQuantity[Black] = 6
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
	resp.Text = "Попробуем изобразить слоненка. Для этого вам нужно 6 цветов: " + Yellow + ", " + Pink + ", " + Blue + ", " + Purple + ", " + Gray + " и " + Black + ". С какого цвета начнем?"
	resp.TTS = "Попробуем изобразить слоненка. Для этого вам нужно шесть цветов\n" + Yellow + "\n" + Pink + "\n" + Blue + "\n" + Purple + "\n" + Gray + " - и " + Black + ". С какого цвета начнем?"

	resultText = "Поздравляю! Слоненок нарисован."
	pictureType = "Образец для рисования слоненка."

	colorsAndQuantity = make(map[string]int)
	colorsAndQuantity[Yellow] = 7; colorsAndQuantity[Pink] = 12; colorsAndQuantity[Blue] = 20; colorsAndQuantity[Purple] = 19; colorsAndQuantity[Gray] = 20; colorsAndQuantity[Black] = 26
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

func sortColorsAndQuantity() {
	colorsSort = nil
	for col := range colorsAndQuantity {
		colorsSort = append(colorsSort, col)
	}
	sort.Strings(colorsSort)
}

func drawingLevel(resp Response, color string) Response {
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
				resp.Text = "От Е—4 до Е—7 и от Е—10 до Е—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-семь" + breakTimeLong +
						"И от Е-десять до Е-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "От Ё—4 до Ё—7 и от Ё—10 до Ё—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четыре до Йо-семь" + breakTimeLong +
						"И от Йо-десять до Йо-одиннадцать" +
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
		case Pink:
			switch colorsAndQuantity[Pink] {
			case 14:
				resp.Text = "От Б—5 до Б—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-пять до Бэ-семь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 13:
				resp.Text = "От В—4 до В—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Вэ-семь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 12:
				resp.Text = "Г—4, Г—5, Г—7 и Г—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-четыре" + breakTimeLong +
						"Гэ-пять" + breakTimeLong +
						"Гэ-семь" + breakTimeLong +
						"И - Гэ-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 11:
				resp.Text = "От Д—4 до Д—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-четыре до Дэ-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 10:
				resp.Text = "От Е—4 до Е—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 9:
				resp.Text = "От Ё—5 до Ё—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пять до Йо-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 8:
				resp.Text = "Ж—3 и от Ж—6 до Ж—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-три" + breakTimeLong +
						"И От Жэ-шесть до Жэ-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 7:
				resp.Text = "З—3, З—4, З—7 и З—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-три" + breakTimeLong +
						"Зэ-четыре" + breakTimeLong +
						"Зэ-семь" + breakTimeLong +
						"И - Зэ-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 6:
				resp.Text = "И—4, И—5, И—7, И—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-четыре" + breakTimeLong +
						"И-пять" + breakTimeLong +
						"И-семь" + breakTimeLong +
						"И-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 5:
				resp.Text = "От Й—5 до Й—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - пять до И-краткое - семь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 4:
				resp.Text = "К—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-пять" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 3:
				resp.Text = "Л—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-четыре" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 2:
				resp.Text = "М—3, М—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-три, Эм-четыре" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 1:
				colorsAndQuantity[Pink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Н—3\nВсе. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-три" + breakTimeLong +
						"Все. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Все. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Все. Выберите " + getRemainingColorsForSpeak + "цвет."
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
		case Purple:
			switch colorsAndQuantity[Purple] {
			case 8:
				resp.Text = "Б—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемь" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 7:
				resp.Text = "В—8, В—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-восемь, Вэ-девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 6:
				resp.Text = "Г—9, Г—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-девять, Гэ-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 5:
				resp.Text = "Д—9, Д—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять, Дэ-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 4:
				resp.Text = "Е—9, Е—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девять, Е-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 3:
				resp.Text = "Ё—9, Ё—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-девять, Йо-десять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 2:
				resp.Text = "Ж—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-девять" +
						"</speak>"
				colorsAndQuantity[Purple]--
			case 1:
				colorsAndQuantity[Purple]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—9\nВсе клетки данным цветом расскрашены. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-девять" + breakTimeLong +
						"Все клетки данным цветом расскрашены. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Все клетки данным цветом расскрашены. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Все клетки данным цветом расскрашены. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Фиолетовый цвет уже использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Фиолетовый цвет уже использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Darkblue:
			switch colorsAndQuantity[Darkblue] {
			case 7:
				resp.Text = "А—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-два" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 6:
				resp.Text = "Б—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 5:
				resp.Text = "В—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "Г—2, Г—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два, Гэ-три" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "Д—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "Е—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-один" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—2\nСиний цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-два" + breakTimeLong +
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
						"Синий цвет закончен. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Г—5\nЧерный цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-пять" + breakTimeLong +
						"Черный цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Черный цвет закончен. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Черный цвет ^закончен^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Черный цвет закончен. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Черный цвет закончен. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"Желтый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от Бэ-пять до Бэ-семь" +
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
						"И от Гэ-шесть до Гэ-девять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "От Д—5 до Д—7 и от Д—9 до Д—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-семь" + breakTimeLong +
						"И от Дэ-девять до Дэ-десять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "Е—4 и от Е—6 до Е—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-четыре" + breakTimeLong +
						"И от Е-шесть до Е-девять" +
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
						"И от Жэ-пять до Жэ-семь" +
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
				resp.Text = "От А—3 до А—4 и от А—9 до А—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-три до А-четыре" + breakTimeLong +
						"И от А-девять до А-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 12:
				resp.Text = "От Б—3 до Б—6 и от Б—8 до Б—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-три до Бэ-шесть" + breakTimeLong +
						"И от Бэ-восемь до Бэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 11:
				resp.Text = "От В—4 до В—6, от В—8 до В—9 и от В—12 до В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-четыре до Вэ-шесть" + breakTimeLong +
						"От Вэ-восемь до Вэ-девять" + breakTimeLong +
						"И от Вэ-двенадцать до Вэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 10:
				resp.Text = "От Г—4 до Г—6, от Г—8 до Г—9 и от Г—12 до Г—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-четыре до Гэ-шесть" + breakTimeLong +
						"От Гэ-восемь до Гэ-девять" + breakTimeLong +
						"И от Гэ-двенадцать до Гэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 9:
				resp.Text = "От Д—5 до Д—6, от Д—8 до Д—10 и от Д—12 до Д—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-пять до Дэ-шесть" + breakTimeLong +
						"От Дэ-восемь до Дэ-десять" + breakTimeLong +
						"И от Дэ-двенадцать до Дэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 8:
				resp.Text = "От Е—4 до Е—6, от Е—9 до Е—10 и от Е—12 до Е—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-шесть" + breakTimeLong +
						"От Е-девять до Е-десять" + breakTimeLong +
						"И от Е-двенадцать до Е-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 7:
				resp.Text = "От Ё—3 до Ё—5, от Ё—9 до Ё—10 и от Ё—12 до Ё—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-три до Йо-пять" + breakTimeLong +
						"От Йо-девять до Йо-десять" + breakTimeLong +
						"И от Йо-двенадцать до Йо-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 6:
				resp.Text = "От Ж—2 до Ж—4, от Ж—9 до Ж—10 и от Ж—12 до Ж—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-два до Жэ-четыре" + breakTimeLong +
						"От Жэ-девять до Жэ-десять" + breakTimeLong +
						"И от Жэ-двенадцать до Жэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 5:
				resp.Text = "От З—1 до З—3 и от З—8 до З—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-один до Зэ-три" + breakTimeLong +
						"И от Зэ-восемь до Зэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 4:
				resp.Text = "От И—1 до И—2, от И—7 до И—9 и от И—11 до И—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-один до И-два" + breakTimeLong +
						"От И-семь до И-девять" + breakTimeLong +
						"И от И-одиннадцать до И-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 3:
				resp.Text = "От Й—1 до Й—3, от Й—6 до Й—8 и от Й—11 до Й—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - один до И-краткое - три" + breakTimeLong +
						"От И-краткое - шесть до И-краткое - восемь" + breakTimeLong +
						"И от И-краткое - одиннадцать до И-краткое - двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 2:
				resp.Text = "От К—2 до К—7 и от К—11 до К—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-два до Ка-семь" + breakTimeLong +
						"И от Ка-одиннадцать до Ка-двенадцать" +
						"</speak>"
				colorsAndQuantity[Green]--
			case 1:
				colorsAndQuantity[Green]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От Л—3 до Л—6 и от Л—12 до Л—13\nЗеленый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-три до Эль-шесть" + breakTimeLong +
						"И от Эль-двенадцать до Эль-тринадцать" + breakTimeLong +
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
						"И от Гэ-шесть до Гэ-девять" +
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
				resp.Text = "Е—1, Е—2 и от Е—4 до Е—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-один" + breakTimeLong +
						"Е-два" + breakTimeLong +
						"И от Е-четыре до Е-шесть" +
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
						"И от Жэ-четыре до Жэ-девять" +
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
						"Синий цвет использован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
		/*case pictureIDEasyLevel6:
		case pictureIDEasyLevel7:
		case pictureIDEasyLevel8:
		case pictureIDEasyLevel9:
		case pictureIDEasyLevel10:
		case pictureIDEasyLevel11:
		case pictureIDEasyLevel12:
		case pictureIDEasyLevel13:
		case pictureIDEasyLevel14:
		case pictureIDEasyLevel15:*/
	}
	return resp
}

func mediumLevels(resp Response) Response {
	switch pictureID {
	case pictureIDMediumLevel1:
		switch color {
		case Pink:
			switch colorsAndQuantity[Pink] {
			case 3:
				resp.Text = "Ё—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 2:
				resp.Text = "Ж—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 1:
				colorsAndQuantity[Pink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—11\nМожем переходить к другому цвету. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-одиннадцать" + breakTimeLong +
						"Можем переходить к другому цвету. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Можем переходить к другому цвету. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Можем переходить к другому цв+ету. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
				resp.Text = "От Е—2 до Е—5 и от Е—21 до Е—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-два до Е-пять" + breakTimeLong +
						"И от Е-двадцать один до Е двадцать-два" +
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
						"И от Эм-семь до Эм-восемнадцать" +
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
				resp.Text = "От С—6 до С—9 и от С—20 до С—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-шесть до Эс-девять" + breakTimeLong +
						"И от Эс-двадцать до Эс-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "От Т—7 до Т—12 и от Т—18 до Т—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-семь до Тэ-двенадцать" + breakTimeLong +
						"И от Тэ-восемнадцать до Тэ-двадцать один" +
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
		case Red:
			switch colorsAndQuantity[Red] {
			case 11:
				resp.Text = "От Ё—5 до К—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-пять до Ка-пять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 10:
				resp.Text = "От Ё—6 до З—6, от К—6 до Л—6 и Х—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-шесть до Зэ-шесть" + breakTimeLong +
						"От Ка-шесть до Эль-шесть" + breakTimeLong +
						"И - Ха-шесть" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 9:
				resp.Text = "От Ё—7 до Ж—7, от К—7 до М—7 и от Ф—7 до Х—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-семь до Же-семь" + breakTimeLong +
						"От Ка-семь до Эм-семь" + breakTimeLong +
						"И от Эф-семь до Ха-семь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 8:
				resp.Text = "От Ё—8 до Ж—8, от К—8 до Н—8 и от У—8 до Х—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-восемь до Же-восемь" + breakTimeLong +
						"От Ка-восемь до Эн-восемь" + breakTimeLong +
						"И от У-восемь до Ха-восемь" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 7:
				resp.Text = "От Ё—9 до Ж—9, от К—9 до Т—9 и от Ф—9 до Х—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девять до Же-девять" + breakTimeLong +
						"От Ка-девять до Тэ-девять" + breakTimeLong +
						"И от Эф-девять до Ха-девять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 6:
				resp.Text = "От Ё—10 до Ж—10, от К—10 до Н—10 и от Ф—10 до Х—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-десять до Же-десять" + breakTimeLong +
						"От Ка-десять до Эн-десять" + breakTimeLong +
						"И от Эф-десять до Ха-десять" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 5:
				resp.Text = "От Ж—11 до К—11, от О—11 до Т—11 и от Ф—11 до Х—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Же-одиннадцать до Ка-одиннадцать" + breakTimeLong +
						"От О-одиннадцать до Тэ-одиннадцать" + breakTimeLong +
						"И от Эф-одиннадцать до Ха-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 4:
				resp.Text = "От А—12 до Ж—12, от Й—12 до Н—12 и от У—12 до Х—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-двенадцать до Жэ-двенадцать" + breakTimeLong +
						"От И-краткое - двенадцать до Эн-двенадцать" + breakTimeLong +
						"И от У-двенадцать до Ха-двенадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 3:
				resp.Text = "От Ж—13 до К—13 и от Ф—13 до Х—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-тринадцать до Ка-тринадцать" + breakTimeLong +
						"И от Эф-тринадцать до Ха-тринадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 2:
				resp.Text = "От Б—14 до Н—14 и Х—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-четырнадцать до Эн-четырнадцать" + breakTimeLong +
						"И - Ха-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Red]--
			case 1:
				colorsAndQuantity[Red]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От В—15 до М—15\nВы расскрасили все клетки красным цветом. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пятнадцать до Эм-пятнадцать" + breakTimeLong +
						"Вы расскрасили все клетки красным цветом. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы расскрасили все клетки красным цветом. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы расскрасили все клетки красным цветом. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Все клетки красным цветом уже раскрашены. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Все клетки красным цветом уже раскрашены. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Blue:
			switch colorsAndQuantity[Blue] {
			case 6:
				resp.Text = "Е—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-пять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "От Г—6 до Е—6 и от И—6 до Й—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-шесть до Е-шесть" + breakTimeLong +
						"И от И-шесть до И-краткое - шесть" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "От В—7 до Е—7 и от З—7 до Й—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-семь до Е-семь" + breakTimeLong +
						"И от Зэ-семь до И-краткое - семь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "От Б—8 до Е—8 и от З—8 до Й—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-восемь до Е-восемь" + breakTimeLong +
						"И от Зэ-восемь до И-краткое - восемь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "От А—9 до Е—9 и от З—9 до Й—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-девять до Е-девять" + breakTimeLong +
						"И от Зэ-девять до И-краткое - девять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—10 до Е—10 и от З—10 до Й—10\nГолубой цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-десять до Е-десять" + breakTimeLong +
						"И от Зэ-десять до И-краткое - десять" + breakTimeLong +
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
			case 4:
				resp.Text = "У—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"У-девять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 3:
				resp.Text = "От О—10 до У—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-десять до У-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 2:
				resp.Text = "От А—11 до Ё—11, от Л—11 до Н—11 и У—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-одиннадцать до Йо-одиннадцать" + breakTimeLong +
						"От Эль-одиннадцать до Эн-одиннадцать" + breakTimeLong +
						"И - У-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 1:
				colorsAndQuantity[Darkblue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—13 до Ё—13 и от Л—13 до Н—13\nСиний цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-тринадцать до Йо-тринадцать" + breakTimeLong +
						"И От Эль-тринадцать до Эн-тринадцать" + breakTimeLong +
						"Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Синий цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Синий цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Синий цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Синий цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Brown:
			switch colorsAndQuantity[Brown] {
			case 2:
				resp.Text = "От Б—1 до П—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-один до Пэ-один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 1:
				colorsAndQuantity[Brown]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От З—2 до З—4\nКоричневый цвет использован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-два до Зэ-четыре" + breakTimeLong +
						"Коричневый цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Коричневый цвет использован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Коричневый цвет ^использован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже использовали коричневый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже использовали коричневый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
			}
		case Black:
			switch colorsAndQuantity[Black] {
			case 4:
				resp.Text = "З—12, И—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-двенадцать, И-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "Д—16, К—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-шестнадцать, Ка-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "А—17, Д—17 и К—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-семнадцать" + breakTimeLong +
						"Дэ-семнадцать" + breakTimeLong +
						"И - Ка-семнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 1:
				colorsAndQuantity[Black]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "От А—18 до О—18\nВы нарисовали черный цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-восемнадцать до О-восемнадцать" + breakTimeLong +
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
			case 12:
				resp.Text = "Д—1, Ё—1"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-один, Йо-один" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 11:
				resp.Text = "Г—2, Е—2 и Ж—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-два" + breakTimeLong +
						"Е-два" + breakTimeLong +
						"И - Жэ-два" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 10:
				resp.Text = "Е—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-три" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 9:
				resp.Text = "От В—6 до З—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-шесть до Зэ-шесть" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 8:
				resp.Text = "От Б—7 до В—7 и От И—7 до Й—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-семь до Вэ-семь" + breakTimeLong +
						"И от И-семь до И-краткое - семь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 7:
				resp.Text = "Б—8, К—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-восемь, Ка-восемь" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 6:
				resp.Text = "Б—9, П—9 и С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-девять" + breakTimeLong +
						"Пэ-девять" + breakTimeLong +
						"И - Эс-девять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 5:
				resp.Text = "Б—10 и от П—10 до С—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" + breakTimeLong +
						"И от Пэ-десять до Эс-десять" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 4:
				resp.Text = "Б—11 и от П—11 до С—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-одиннадцать" + breakTimeLong +
						"И от Пэ-одиннадцать до Эс-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 3:
				resp.Text = "Б—12, Р—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-двенадцать, Эр-двенадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "П—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Пэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 1:
				colorsAndQuantity[Blue]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "З—14, З—15\nВы нарисовали голубой цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-четырнадцать, Зэ-пятнадцать" + breakTimeLong +
						"Вы нарисовали голубой цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали голубой цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали голубой цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
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
				resp.Text = "От В—10 до Г—10 и от Ё—10 до К—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-десять до Гэ-десять" + breakTimeLong +
						"И от Йо-десять до Ка-десять" +
						"</speak>"
				colorsAndQuantity[Darkblue]--
			case 4:
				resp.Text = "От В—11 до Г—11 и от Ё—11 до К—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-одиннадцать до Гэ-одиннадцать" + breakTimeLong +
						"И от Йо-одиннадцать до Ка-одиннадцать" +
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
						"И от Эль-четырнадцать до Пэ-четырнадцать" + breakTimeLong +
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
			case 12:
				resp.Text = "От В—5 до З—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пять до Зэ-пять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "Б—6 и от И—6 до Й—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-шесть" + breakTimeLong +
						"И от И-шесть до И-краткое - шесть" +
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
				resp.Text = "От В—15 до П—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пятнадцать до Пэ-пятнадцать" +
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
						"И от Вэ-пятнадцать до Вэ-двадцать один" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 13:
				resp.Text = "Г—7, Г—13 и от Г—15 до Г—20"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Гэ-семь" + breakTimeLong +
						"Гэ-тринадцать" + breakTimeLong +
						"И от Гэ-пятнадцать до Гэ-двадцать" +
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
						"И от Йо-семнадцать до Йо-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 9:
				resp.Text = "От Ж—9 до Ж—11 и от Ж—16 до Ж—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-девять до Жэ-одиннадцать" + breakTimeLong +
						"И от Жэ-шестнадцать до Жэ-двадцать два" +
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
						"И от И-шестнадцать до И-двадцать два" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 6:
				resp.Text = "Й—8, Й—12 и от Й—17 до Й—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - восемь" + breakTimeLong +
						"И-краткое - двенадцать" + breakTimeLong +
						"И от И-краткое - семнадцать до И-краткое - двадцать два" +
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
						"И от Эм-пятнадцать до Эм-двадцать" +
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
						"И от Эн-пятнадцать до Эн-двадцать один" +
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
						"Бежевый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от И-двадцать четыре до Ка-двадцать четыре" + breakTimeLong +
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
						"И от Бэ-тринадцать до Бэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 15:
				resp.Text = "От В—2 до В—6, от В—13 до В—14 и от В—22 до В—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-два до Вэ-шесть" + breakTimeLong +
						"От Вэ-тринадцать до Вэ-четырнадцать" + breakTimeLong +
						"И от Вэ-двадцать два до Вэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 14:
				resp.Text = "От Г—3 до Г—6, Г—14 и от Г—21 до Г—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Гэ-три до Гэ-шесть" + breakTimeLong +
						"Гэ-четырнадцать" + breakTimeLong +
						"И от Гэ-двадцать один до Гэ-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 13:
				resp.Text = "От Д—4 до Д—6 и от Д—14 до Д—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Дэ-четыре до Дэ-шесть" + breakTimeLong +
						"И от Дэ-четырнадцать до Дэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 12:
				resp.Text = "От Е—4 до Е—6 и от Е—14 до Е—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-четыре до Е-шесть" + breakTimeLong +
						"И от Е-четырнадцать до Е-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 11:
				resp.Text = "От Ё—4 до Ё—7 и от Ё—13 до Ё—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-четыре до Йо-семь" + breakTimeLong +
						"И от Йо-тринадцать до Йо-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 10:
				resp.Text = "От Ж—4 до Ж—8 и от Ж—14 до Ж—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-четыре до Жэ-восемь" + breakTimeLong +
						"И от Жэ-четырнадцать до Жэ-пятнадцать" +
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
				resp.Text = "От И—4 до И—8 и от И—14 до И—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-четыре до И-восемь" + breakTimeLong +
						"И от И-четырнадцать до И-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 7:
				resp.Text = "От Й—4 до Й—7 и от Й—13 до Й—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - четыре до И-краткое - семь" + breakTimeLong +
						"И от И-краткое - тринадцать до И-краткое - шестнадцать" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 6:
				resp.Text = "От К—4 до К—6 и от К—14 до К—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-четыре до Ка-шесть" + breakTimeLong +
						"И от Ка-четырнадцать до Ка-двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "От Л—4 до Л—6 и от Л—14 до Л—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-четыре до Эль-шесть" + breakTimeLong +
						"И от Эль-четырнадцать до Эль-двадцать два" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "От М—3 до М—6, М—14 и от М—21 до М—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эм-три до Эм-шесть" + breakTimeLong +
						"Эм-четырнадцать" + breakTimeLong +
						"И от Эм-двадцать один до Эм-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 3:
				resp.Text = "От Н—2 до Н—6, от Н—13 до Н—14 и от Н—22 до Н—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-два до Эн-шесть" + breakTimeLong +
						"От Эн-тринадцать до Эн-четырнадцать" + breakTimeLong +
						"И от Эн-двадцать два до Эн-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 2:
				resp.Text = "От О—1 до О—4, О—6 и от О—13 до О—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-один до О-четыре" + breakTimeLong +
						"О-шесть" + breakTimeLong +
						"И от О-тринадцать до О-двадцать два" +
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
						"И от Ка-восемь до Эм-восемь" +
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
						"И от Ка-двенадцать до Эм-двенадцать" + breakTimeLong +
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
		case Gray:
			switch colorsAndQuantity[Gray] {
			case 8:
				resp.Text = "Л—2, М—2"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-два, Эм-два" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 7:
				resp.Text = "К—3, Н—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-три, Эн-три" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 6:
				resp.Text = "В—4, Й—4 и О—4"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четыре" + breakTimeLong +
						"И-краткое - четыре" + breakTimeLong +
						"И - О-четыре" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 5:
				resp.Text = "От В—5 до Д—5, Й—5 и О—5"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-пять до Дэ-пять" + breakTimeLong +
						"И-краткое - пять" + breakTimeLong +
						"И - О-пять" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 4:
				resp.Text = "От В—6 до Д—6, К—6 и Н—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-шесть до Дэ-шесть" + breakTimeLong +
						"Ка-шесть" + breakTimeLong +
						"И - Эн-шесть" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 3:
				resp.Text = "От В—7 до Д—7, Л—7 и М—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-семь до Дэ-семь" + breakTimeLong +
						"Эль-семь" + breakTimeLong +
						"И - Эм-семь" +
						"</speak>"
				colorsAndQuantity[Gray]--
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
				colorsAndQuantity[Gray]--
			case 1:
				colorsAndQuantity[Gray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Е—9\nВы нарисовали серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-девять" + breakTimeLong +
						"Вы нарисовали серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали серый цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Серый цвет вы уже нарисовали. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Серый цвет вы уже нарисовали. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от Эн-два до Тэ-два" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 8:
				resp.Text = "Й—3 и от О—3 до Т—3"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - три" + breakTimeLong +
						"И от О-три до Тэ-три" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 7:
				resp.Text = "Й—6 и от О—6 до Т—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - шесть" + breakTimeLong +
						"И от О-шесть до Тэ-шесть" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 6:
				resp.Text = "Й—7, К—7 и от Н—7 до Т—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - семь" + breakTimeLong +
						"Ка-семь" + breakTimeLong +
						"И от Эн-семь до Тэ-семь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 5:
				resp.Text = "От Й—8 до Н—8 и от П—8 до Т—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - восемь до Эн-восемь" + breakTimeLong +
						"И от Пэ-восемь до Тэ-восемь" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 4:
				resp.Text = "От Й—9 до К—9 и от О—9 до Р—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-краткое - девять до Ка-девять" + breakTimeLong +
						"И от О-девять до Эр-девять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 3:
				resp.Text = "От К—10 до О—10 и от С—10 до Т—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-десять до О-десять" + breakTimeLong +
						"И от Эс-десять до Тэ-десять" +
						"</speak>"
				colorsAndQuantity[Yellow]--
			case 2:
				resp.Text = "Й—11 и от П—11 до Т—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - одиннадцать" + breakTimeLong +
						"И от Пэ-одиннадцать до Тэ-одиннадцать" +
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
						"И от Пэ-четыре до Тэ-четыре" +
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
						"И от Пэ-пять до Тэ-пять" +
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
						"И от Эль-девять до Эн-девять" +
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
						"И от О-семнадцать до Эр-семнадцать" + breakTimeLong +
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
						"И от Дэ-двадцать четыре до Дэ-тридцать" +
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
						"И от Е-двадцать пять до Е-тридцать" +
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
						"И от Йо-двадцать шесть до Йо-двадцать девять" +
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
						"И от И-краткое - двадцать шесть до И-краткое - двадцать девять" +
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
						"И от Ка-двадцать пять до Ка-тридцать" +
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
						"И от Эль-двадцать четыре до Эль-тридцать" +
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
		case Pink:
			switch colorsAndQuantity[Pink] {
			case 2:
				resp.Text = "Д—8, Л—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-восемь, Эль-восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 1:
				colorsAndQuantity[Pink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Д—9, Е—9, К—9 и Л—9\nРозовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять" + breakTimeLong +
						"Е-девять" + breakTimeLong +
						"Ка-девять" + breakTimeLong +
						"И - Эль-девять" + breakTimeLong +
						"Розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Розовый цвет уже был нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Розовый цвет уже был нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от Жэ-двадцать один до Жэ-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "З—19, З—20 и от З—22 до З—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-девятнадцать" + breakTimeLong +
						"Зэ-двадцать" + breakTimeLong +
						"И от Зэ-двадцать два до Зэ-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "И—19 и от И—21 до И—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-девятнадцать" + breakTimeLong +
						"И от И-двадцать один до И-двадцать шесть" +
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
						"И от Эль-пять до О-пять" +
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
						"И от Эм-десять до Пэ-десять" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 12:
				resp.Text = "От Б—11 до Д—11 и от Л—11 до Р—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-одиннадцать до Дэ-одиннадцать" + breakTimeLong +
						"И от Эль-одиннадцать до Эр-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 11:
				resp.Text = "От Б—12 до Ё—12, от Й—12 до С—12 и от Ц—8 до Ш—8"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-двенадцать до Йо-двенадцать" + breakTimeLong +
						"От И-краткое - двенадцать до Эс-двенадцать" + breakTimeLong +
						"И от Цэ-двенадцать до Ша-двенадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 10:
				resp.Text = "От Б—13 до Г—13, от М—13 до Т—13 и от Х—13 до Ц—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-тринадцать до Гэ-тринадцать" + breakTimeLong +
						"От Эм-тринадцать до Тэ-тринадцать" + breakTimeLong +
						"И от Ха-тринадцать до Цэ-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 9:
				resp.Text = "В—14 и от Н—14 до Ц—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-четырнадцать" + breakTimeLong +
						"И от Эн-четырнадцать до Цэ-четырнадцать" +
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
						"И от Пэ-шестнадцать до Эс-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 6:
				resp.Text = "Д—17, Т—17, от К—17 до М—17, от Р—17 до Т—17, Ф—17 и Х—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-семнадцать" + breakTimeLong +
						"Тэ-семнадцать" + breakTimeLong +
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
						"И от Эр-двадцать два до У-двадцать два" +
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
						"И от Эр-пятнадцать до Тэ-пятнадцать" +
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
						"И от Тэ-двадцать до Эф-двадцать" +
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
						"И от О-шесть до У-шесть" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 13:
				resp.Text = "От Е—7 до З—7, Н—7 и от С—7 до Ф—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Е-семь до Зэ-семь" + breakTimeLong +
						"Эн-семь" + breakTimeLong +
						"И от Эс-семь до Эф-семь" +
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
						"И от Тэ-восемь до Эф-восемь" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 11:
				resp.Text = "Д—9, Е—9 и от Т—9 до Х—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-девять" + breakTimeLong +
						"Е-девять" + breakTimeLong +
						"И от Тэ-девять до Ха-девять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 10:
				resp.Text = "В—10, от Д—10 до Ё—10 и от Т—10 до Х—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-десять" + breakTimeLong +
						"От Дэ-десять до Йо-десять" + breakTimeLong +
						"И от Тэ-десять до Ха-десять" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 9:
				resp.Text = "В—11, от Д—11 до Ж—11 и от С—11 до Ф—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать" + breakTimeLong +
						"От Дэ-одиннадцать до Жэ-одиннадцать" + breakTimeLong +
						"И от Эс-одиннадцать до Эф-одиннадцать" +
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
						"И от Эр-двенадцать до Ха-двенадцать" +
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
						"И от О-тринадцать до Ха-тринадцать" +
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
						"И от Тэ-четырнадцать до Цэ-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 5:
				resp.Text = "От А—15 до Е—15 и от У—15 до Ч—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От А-пятнадцать до Е-пятнадцать" + breakTimeLong +
						"И от У-пятнадцать до Чэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 4:
				resp.Text = "От Б—16 до Е—16 и от Ф—16 до Ш—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-шестнадцать до Е-шестнадцать" + breakTimeLong +
						"И от Эф-шестнадцать до Ша-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 3:
				resp.Text = "От Б—17 до Е—17 и от Ф—17 до Ш—17"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-семнадцать до Е-семнадцать" + breakTimeLong +
						"И от Эф-семнадцать до Ша-семнадцать" +
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
		case Pink:
			switch colorsAndQuantity[Pink] {
			case 2:
				resp.Text = "И—9, Р—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-девять, Эр-девять" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 1:
				colorsAndQuantity[Pink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "И—10, Й—10, П—10 и Р—10\nРозовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-десять" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"Пэ-десять" + breakTimeLong +
						"И - Эр-десять" + breakTimeLong +
						"Розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от Эс-двадцать один до У-двадцать один" +
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
						"И от О-двадцать три до Эр-двадцать три" +
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
						"От Дэ-двенадцать до Дэ-тринадцать" +
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
						"И от Эль-двадцать три до Эн-двадцать три" +
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
						"И от Эф-двадцать пять до Чэ-двадцать пять" + breakTimeLong +
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
						"И от Гэ-шестнадцать до Гэ-девятнадцать" +
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
						"И от Жэ-четырнадцать до Жэ-двадцать" +
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
						"И от И-краткое - четырнадцать до И-краткое - двадцать один" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 5:
				resp.Text = "К—10, К—11 и от К—13 до К—23"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ка-десять" + breakTimeLong +
						"Ка-одиннадцать" + breakTimeLong +
						"И от Ка-тринадцать до Ка-двадцать три" +
						"</speak>"
				colorsAndQuantity[Brown]--
			case 4:
				resp.Text = "Л—6 и от Л—11 до Л—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-шесть" + breakTimeLong +
						"И от Эль-одиннадцать до Эль-двадцать четыре" +
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
						"И от Эн-шестнадцать до Эн-девятнадцать" +
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
						"И от Бэ-двенадцать до Бэ-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 16:
				resp.Text = "В—1, В—6 и от В—10 до В—19, кроме: В—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-один" + breakTimeLong +
						"Вэ-шесть" + breakTimeLong +
						"И от Вэ-десять до Вэ-девятнадцать" + breakTimeMiddle +
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
						"И от Гэ-двадцать до Гэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 14:
				resp.Text = "Д—3, Д—4 и от Д—23 до Д—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Дэ-три" + breakTimeLong +
						"Дэ-четыре" + breakTimeLong +
						"И от Дэ-двадцать три до Дэ-двадцать семь" +
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
						"И от Эм-двадцать три до Эм-двадцать семь" +
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
						"И от Эн-двадцать до Эн-двадцать два" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 3:
				resp.Text = "О—1, О—6 и от О—10 до О—19, кроме: О—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"О-один" + breakTimeLong +
						"О-шесть" + breakTimeLong +
						"И от О-десять до О-девятнадцать" + breakTimeMiddle +
						"Кроме: О-тринадцать" +
						"</speak>"
				colorsAndQuantity[Black]--
			case 2:
				resp.Text = "От П—1 до П—5 и от П—12 до П—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-один до Пэ-пять" + breakTimeLong +
						"И от Пэ-двенадцать до Пэ-пятнадцать" +
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
						"И от У-девять до Цэ-девять" +
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
						"И от Ха-десять до Чэ-десять" +
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
						"И от Цэ-двенадцать до Ша-двенадцать" +
						"</speak>"
				colorsAndQuantity[Beige]--
			case 4:
				resp.Text = "От Б—13 до Д—13 и от Х—13 до Ш—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Бэ-тринадцать до Дэ-тринадцать" + breakTimeLong +
						"И от Ха-тринадцать до Ша-тринадцать" +
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
						"И от Чэ-четырнадцать до Ща-четырнадцать" +
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
						"И от Эр-семнадцать до Тэ-семнадцать" +
						"</speak>"
				colorsAndQuantity[Orange]--
			case 2:
				resp.Text = "Е—18, от Ж—18 до И—18 и от Р—18 до Т—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Е-восемнадцать" + breakTimeLong +
						"От Жэ-восемнадцать до И-восемнадцать" + breakTimeLong +
						"И от Эр-восемнадцать до Тэ-восемнадцать" +
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
						"И от Пэ-девятнадцать до Тэ-девятнадцать" + breakTimeLong +
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
		case Pink:
			switch colorsAndQuantity[Pink] {
			case 3:
				resp.Text = "З—9, С—9"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Зэ-девять, Эс-девять" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 2:
				resp.Text = "И—10, Р—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-десять, Эр-десять" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 1:
				colorsAndQuantity[Pink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ё—11, И—11, Р—11 и У—11\nРозовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-одиннадцать" + breakTimeLong +
						"И-одиннадцать" + breakTimeLong +
						"Эр-одиннадцать" + breakTimeLong +
						"И - У-одиннадцать" + breakTimeLong +
						"\nРозовый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Розовый цвет нарисован. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Вы уже нарисовали розовый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вы уже нарисовали розовый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от Эль-двенадцать до Эн-двенадцать" + breakTimeLong +
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
						"И от У-восемь до Цэ-восемь" +
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
						"И от Бэ-двадцать шесть до Бэ-двадцать восемь" +
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
						"И от Твердый-знак - двадцать один до Твердый-знак - двадцать три" +
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
		case Pink:
			switch colorsAndQuantity[Pink] {
			case 12:
				resp.Text = "Б—10"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-десять" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 11:
				resp.Text = "В—11, Г—11"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Вэ-одиннадцать, Гэ-одиннадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 10:
				resp.Text = "А—12 и от В—12 до Д—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-двенадцать" + breakTimeLong +
						"И от Вэ-двенадцать до Дэ-двенадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 9:
				resp.Text = "От В—13 до Е—13, М—13 и Н—13"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Вэ-тринадцать до Е-тринадцать" + breakTimeLong +
						"Эм-тринадцать" + breakTimeLong +
						"И - Эн-тринадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 8:
				resp.Text = "Б—14, от Д—14 до Ё—14 и от Л—14 до Н—14"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Бэ-четырнадцать" + breakTimeLong +
						"От Дэ-четырнадцать до Йо-четырнадцать" + breakTimeLong +
						"И от Эль-четырнадцать до Эн-четырнадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 7:
				resp.Text = "А—15, от Е—15 до И—15 и от Л—15 до Н—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"А-пятнадцать" + breakTimeLong +
						"От Е-пятнадцать до И-пятнадцать" + breakTimeLong +
						"И от Эль-пятнадцать до Эн-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 6:
				resp.Text = "М—16"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эм-шестнадцать" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 5:
				resp.Text = "Н—26"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-двадцать шесть" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 4:
				resp.Text = "Ж—27 и от И—27 до О—27"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Жэ-двадцать семь" + breakTimeLong +
						"И от И-двадцать семь до О-двадцать семь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 3:
				resp.Text = "Н—28"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эн-двадцать восемь" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 2:
				resp.Text = "И—29"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-двадцать девять" +
						"</speak>"
				colorsAndQuantity[Pink]--
			case 1:
				colorsAndQuantity[Pink]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Л—30\nРозовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Эль-тридцать" + breakTimeLong +
						"Розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Розовый цвет нарисован. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Розовый цвет ^нарисован^. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Розовый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Розовый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от У-семнадцать до Цэ-семнадцать" +
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
						"И от Эф-двадцать четыре до Цэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Blue]--
			case 2:
				resp.Text = "Ё—25, Ж—25 и от Х—25 до Ч—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Йо-двадцать пять" + breakTimeLong +
						"Жэ-двадцать пять" + breakTimeLong +
						"И от Ха-двадцать пять до Чэ-двадцать пять" +
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
						"И от Эф-двадцать шесть до Цэ-двадцать шесть" + breakTimeLong +
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
						"И от Дэ-шесть до Йо-шесть" +
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
						"И от У-двадцать два до Цэ-двадцать два" +
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
		case Gray:
			switch colorsAndQuantity[Gray] {
			case 20:
				resp.Text = "От Ё—9 до Ё—12 и от Ё—18 до Ё—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Йо-девять до Йо-двенадцать" + breakTimeLong +
						"И от Йо-восемнадцать до Йо-двадцать один" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 19:
				resp.Text = "От Ж—8 до Ж—13 и от Ж—17 до Ж—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Жэ-восемь до Жэ-тринадцать" + breakTimeLong +
						"И от Жэ-семнадцать до Жэ-двадцать два" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 18:
				resp.Text = "От З—8 до З—13 и от З—17 до З—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Зэ-восемь до Зэ-тринадцать" + breakTimeLong +
						"И от Зэ-семнадцать до Зэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 17:
				resp.Text = "От И—11 до И—13 и от И—17 до И—25, кроме: И—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От И-одиннадцать до И-тринадцать" + breakTimeLong +
						"И от И-семнадцать до И-двадцать пять" + breakTimeMiddle +
						"Кроме: И-двадцать один" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 16:
				resp.Text = "Й—9, Й—10, Й—12 и от Й—17 до Й—25, кроме: Й—21"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"И-краткое - девять" + breakTimeLong +
						"И-краткое - десять" + breakTimeLong +
						"И-краткое - двенадцать" + breakTimeLong +
						"И от И-краткое - семнадцать до И-краткое - двадцать пять" + breakTimeMiddle +
						"Кроме: И-краткое - двадцать один" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 15:
				resp.Text = "От К—6 до К—11, от К—14 до К—16, от К—18 до К—21 и от К—23 до К—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Ка-шесть до Ка-одиннадцать" + breakTimeLong +
						"От Ка-четырнадцать до Ка-шестнадцать" + breakTimeLong +
						"От Ка-восемнадцать до Ка-двадцать один" + breakTimeLong +
						"И от Ка-двадцать три до Ка-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 14:
				resp.Text = "От Л—6 до Л—13, Л—16, Л—17 и от Л—19 до Л—25, кроме: Л—22"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эль-шесть до Эль-тринадцать" + breakTimeLong +
						"Эль-шестнадцать" + breakTimeLong +
						"Эль-семнадцать" + breakTimeLong +
						"И от Эль-девятнадцать до Эль-двадцать пять" + breakTimeMiddle +
						"Кроме: Эль-двадцать два" +
						"</speak>"
				colorsAndQuantity[Gray]--
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
				colorsAndQuantity[Gray]--
			case 12:
				resp.Text = "От Н—5 до Н—12 и от Н—16 до Н—21, кроме: Н—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эн-пять до Эн-двенадцать" + breakTimeLong +
						"И от Эн-шестнадцать до Эн-двадцать один" + breakTimeMiddle +
						"Кроме: Эн-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 11:
				resp.Text = "От О—5 до О—22, кроме: О—18"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От О-пять до О-двадцать два" + breakTimeMiddle +
						"Кроме: О-восемнадцать" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 10:
				resp.Text = "От П—5 до П—12 и от П—18 до П—24"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Пэ-пять до Пэ-двенадцать" + breakTimeLong +
						"И от Пэ-восемнадцать до Пэ-двадцать четыре" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 9:
				resp.Text = "От Р—5 до Р—24, кроме: Р—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эр-пять до Эр-двадцать четыре" + breakTimeMiddle +
						"Кроме: Эр-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 8:
				resp.Text = "От С—5 до С—24, кроме: С—15"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эс-пять до Эс-двадцать четыре" + breakTimeMiddle +
						"Кроме: Эс-пятнадцать" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 7:
				resp.Text = "От Т—6 до Т—8, от Т—10 до Т—13 и от Т—23 до Т—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Тэ-шесть до Тэ-восемь" + breakTimeLong +
						"От Тэ-десять до Тэ-тринадцать" + breakTimeLong +
						"И от Тэ-двадцать три до Тэ-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 6:
				resp.Text = "От У—7 до У—13 и У—25"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От У-семь до У-тринадцать" + breakTimeLong +
						"И - У-двадцать пять" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 5:
				resp.Text = "От Ф—6 до Ф—12"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Эф-шесть до Эф-двенадцать" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 4:
				resp.Text = "Х—6, Х—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ха-шесть, Ха-семь" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 3:
				resp.Text = "От Ц—5 до Ц—7"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Цэ-пять до Цэ-семь" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 2:
				resp.Text = "От Ч—3 до Ч—6"
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"От Чэ-три до Чэ-шесть" +
						"</speak>"
				colorsAndQuantity[Gray]--
			case 1:
				colorsAndQuantity[Gray]--
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Ш—2\nВы нарисовали серый цвет. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Ша-два" + breakTimeLong +
						"Вы нарисовали серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет." +
						"</speak>"
				endDrawingText = "Вы нарисовали серый цвет. Выберите " + getRemainingColors + "цвет."
				endDrawingSSML = "Вы нарисовали серый цвет. Выберите " + getRemainingColorsForSpeak + "цвет."
			case 0:
				var getRemainingColors = getRemainingColors()
				var getRemainingColorsForSpeak = getRemainingColorsForSpeak()
				resp.Text = "Серый цвет уже нарисован. Выберите " + getRemainingColors + "цвет."
				resp.TTSType = "ssml"
				resp.SSML =
					"<speak>" +
						"Серый цвет уже нарисован. Выберите " + getRemainingColorsForSpeak + "цвет." +
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
						"И от Пэ-пятнадцать до Эс-пятнадцать" +
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