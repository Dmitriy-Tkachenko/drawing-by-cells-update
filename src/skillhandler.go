package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Payload struct {
	Text string
}

type Selected string

const (
	Easy = Selected(rune(iota))
	Medium = Selected(rune(iota))
	Hard = Selected(rune(iota))
	Mixed = Selected(rune(iota))
	None = Selected(rune(iota))
)

var selected Selected
var numbEasyPicture = 5
var numbMediumPicture = 5
var numbHardPicture = 5
var numbMixedPicture = numbEasyPicture + numbMediumPicture + numbHardPicture
var sequenceOfPictures []int
var exercise int
var color = ""
var newSession, pictureSelected, colorSelected, showBtnYesAndNo, pictureSample bool

func processingPhrasesAndClicks(r Request) (resp Response) {
	rand.Seed(time.Now().UnixNano())
	if r.Request.Type == SimpleUtterance {
		resp = processingSimpleUtterance(r)
	}
	if r.Request.Type == ButtonPressed {
		resp = processingButtonPressed(r)
	}
	return
}

func processingSimpleUtterance(r Request) (resp Response) {
	switch r.Request.Command {
	case "запусти навык рисунки по клеточкам", "включи навык рисунки по клеточкам", "открой навык рисунки по клеточкам", "хочу навык рисунки по клеточкам",
		"запусти скилл рисунки по клеточкам", "включи скилл рисунки по клеточкам", "открой скилл рисунки по клеточкам", "хочу скилл рисунки по клеточкам",
		"запусти skill рисунки по клеточкам", "включи skill рисунки по клеточкам", "открой skill рисунки по клеточкам", "хочу skill рисунки по клеточкам":
		resp = processingOnStart(resp)
	case "легкие", "легкий", "легкая":
		resp = processingCommandEasy(resp)
	case "средние", "средний", "среднее", "средняя":
		resp = processingCommandMedium(resp)
	case "сложные", "сложный", "сложная":
		resp = processingCommandHard(resp)
	case "вперемешку", "перемешать", "вперемежку":
		resp = processingCommandMixed(resp)
	case "следующий рисунок", "другой рисунок":
		resp = processingCommandNextPicture(r, resp)
	case "рисовать по картинке", "нарисовать по картинке", "порисовать по картинке", "по картинке":
		resp = processingCommandSample(r, resp)
	case "скачать шаблон", "скачать":
		resp = downloadTemplate(r, resp)
	case "дальше", "далее":
		resp = processingCommandNext(r, resp)
	case "готово", "готова", "готов", "сделано", "сделана", "сделан", "завершено", "завершена", "завершен":
		resp = processingCommandShowResult(r, resp)
	case Black, Brown, Darkblue, Blue, Purple, Red, Pink, Orange, Yellow, Green, Beige, Gray, White:
		resp = processingCommandColor(r, resp, r.Request.Command)
	case OnInterrupt:
		resp = processingCommandOnInterrupt(resp)
	default:
		resp = processingCommandOnDefault(r, resp)
	}
	return
}

func processingButtonPressed(r Request) (resp Response) {
	var p Payload
	err := json.Unmarshal(r.Request.Payload, &p)
	if err != nil {
		resp.Text = "Что-то пошло не так"
		return
	}
	switch p.Text {
	case "легкие":
		resp = processingCommandEasy(resp)
	case "средние":
		resp = processingCommandMedium(resp)
	case "сложные":
		resp = processingCommandHard(resp)
	case "вперемешку":
		resp = processingCommandMixed(resp)
	case "следующий рисунок":
		resp = processingCommandNextPicture(r, resp)
	case "рисовать по картинке":
		resp = processingCommandSample(r, resp)
	case "скачать шаблон":
		resp = downloadTemplate(r, resp)
	case "дальше":
		resp = processingCommandNext(r, resp)
	case "готово":
		resp = processingCommandShowResult(r, resp)
	case "да":
		resp = processingCommandYes(resp)
	case "нет":
		resp = processingCommandNo(r, resp)
	case Black, Brown, Darkblue, Blue, Purple, Red, Pink, Orange, Yellow, Green, Beige, Gray, White:
		resp = processingCommandColor(r, resp, p.Text)
	}
	return
}

func processingOnStart(resp Response) Response {
	resp.Text = "Я подготовила для Вас рисунки трех видов сложности: легкие, средние и сложные. Также можно выбрать вариант «Вперемешку». Для легких рисунков необходимо подготовить клетчатую бумагу размером 15x15 клеток, для средних — 25x25 клеток, а для сложных — 30x30 клеток. Какой вид сложности Вы хотите выбрать?"
	resp.TTS = "Я подготовила для Вас рисунки трех видов сложности:\n легкие - средние - и - сложные. Также можно выбрать вариант ^Вперемешку^. Для легких рисунков необходимо подготовить клетчатую бумагу размером пятнадцать на пятнадцать клеток, для средних - двадцать пять на двадцать пять клеток, а для сложных - тридцать на тридцать клеток.\nКакой вид сложности Вы хотите выбрать?"
	resp.AddButton("Легкие", Payload {
		Text: "легкие",
	})
	resp.AddButton("Средние", Payload {
		Text: "средние",
	})
	resp.AddButton("Сложные", Payload {
		Text: "сложные",
	})
	resp.AddButton("Вперемешку", Payload {
		Text: "вперемешку",
	})
	resp.AddButton("Скачать шаблон", Payload{
		Text: "скачать шаблон",
	})
	exercise = 0
	selected = None
	hintID = 0
	pictureID = 0
	newSession = true
	pictureSelected = false
	colorSelected = false
	pictureSample = false
	showBtnYesAndNo = false
	pictureDraw = false
	color = ""
	return resp
}

func processingCommandEasy(resp Response) Response {
	selected = Easy
	exercise = 0
	hintID = 0
	pictureID = 0
	pictureSelected = false
	colorSelected = false
	pictureSample = false
	showBtnYesAndNo = false
	pictureDraw = false
	color = ""
	makeSetRandomLevels(resp, numbEasyPicture)
	resp = levelSelection(resp)
	if newSession {
		resp.Text = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой.\n" + resp.Text
		resp.TTS = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой. " + resp.TTS
		newSession = false
	}
	return resp
}

func processingCommandMedium(resp Response) Response {
	selected = Medium
	exercise = 0
	hintID = 0
	pictureID = 0
	pictureSelected = false
	colorSelected = false
	pictureSample = false
	showBtnYesAndNo = false
	pictureDraw = false
	color = ""
	makeSetRandomLevels(resp, numbMediumPicture)
	resp = levelSelection(resp)
	if newSession {
		resp.Text = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой.\n" + resp.Text
		resp.TTS = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой. " + resp.TTS
		newSession = false
	}
	return resp
}

func processingCommandHard(resp Response) Response {
	selected = Hard
	exercise = 0
	hintID = 0
	pictureID = 0
	pictureSelected = false
	colorSelected = false
	pictureSample = false
	showBtnYesAndNo = false
	pictureDraw = false
	color = ""
	makeSetRandomLevels(resp, numbHardPicture)
	resp = levelSelection(resp)
	if newSession {
		resp.Text = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой.\n" + resp.Text
		resp.TTS = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой. " + resp.TTS
		newSession = false
	}
	return resp
}

func processingCommandMixed(resp Response) Response {
	selected = Mixed
	exercise = 0
	hintID = 0
	pictureID = 0
	pictureSelected = false
	colorSelected = false
	pictureSample = false
	showBtnYesAndNo = false
	pictureDraw = false
	color = ""
	makeSetRandomLevels(resp, numbMixedPicture)
	resp = levelSelection(resp)
	if newSession {
		resp.Text = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой.\n" + resp.Text
		resp.TTS = "Сейчас Вам предстоит закрашивать клетки, которые я назову. Чтобы перейти к рисованию следующих клеток, просто скажите мне «Дальше». Если вы не хотите рисовать текущий рисунок, то скажите мне «Следующий рисунок», и я подберу для Вас другой. " + resp.TTS
		newSession = false
	}
	return resp
}

func processingCommandSample(r Request, resp Response) Response {
	if selected != None && !pictureDraw {
		resp.Text = pictureType
		resp.Card = NewBigImage(
			hintID,
		)
		resp.AddButton("Готово", Payload{
			Text: "готово",
		})
		resp.AddButton("Следующий рисунок", Payload{
			Text: "следующий рисунок",
		})
		pictureSelected = false
		colorSelected = false
		showBtnYesAndNo = false
		pictureSample = true
	} else if pictureDraw {
		resp = processingCommandOnDefault(r, resp)
	} else  {
		resp = processingCommandOnDefault(r, resp)
	}
	return resp
}

func processingCommandOnInterrupt(resp Response) Response {
	resp.Text = "Возвращайтесь, когда захотите порисовать снова."
	resp.EndSession = true
	return resp
}

func processingCommandOnDefault(r Request, resp Response) Response {
	if showBtnYesAndNo {
		if r.Request.Command == "да" || r.Request.Command == "хочу" {
			return processingCommandYes(resp)
		} else if r.Request.Command == "нет" || r.Request.Command == "не хочу" {
			return processingCommandNo(r, resp)
		}
	}
	resp.Text = "Хотите завершить рисование?"
	resp.AddButton("Да", Payload{
		Text: "да",
	})
	resp.AddButton("Нет", Payload{
		Text: "нет",
	})
	showBtnYesAndNo = true
	return resp
}

func processingCommandYes(resp Response) Response {
	showBtnYesAndNo = false
	return processingCommandOnInterrupt(resp)
}

func processingCommandNo(r Request, resp Response) Response {
	showBtnYesAndNo = false
	if pictureSelected && !colorSelected && !pictureSample && selected != None {
		exercise--
		return levelSelection(resp)
	} else if colorSelected && !pictureSample && selected != None && !pictureDraw {
		return processingCommandNext(r, resp)
	} else if pictureSample && selected != None && !pictureDraw {
		return processingCommandSample(r, resp)
	} else if selected != None && !pictureDraw {
		return processingCommandNextPicture(r, resp)
	} else if selected == None {
		return chooseDifficulty(resp)
	} else if pictureDraw {
		return hintToGoToNextPicture(resp)
	} else {
		return processingOnStart(resp)
	}
}

func makeSetRandomLevels(resp Response, numbPictures int) Response {
	sequenceOfPictures = nil
	exercise = 0
	for i := 0; i < numbPictures; i++ {
		sequenceOfPictures = appendIfMissing(sequenceOfPictures, rand.Intn(numbPictures) + 1, numbPictures)
	}
	return resp
}

func processingCommandNextPicture(r Request, resp Response) Response {
	if (selected == Easy && exercise < numbEasyPicture) ||
		(selected == Medium && exercise < numbMediumPicture) ||
		(selected == Hard && exercise < numbHardPicture) ||
		(selected == Mixed && exercise < numbMixedPicture) {
		color = ""
		resp = levelSelection(resp)
	} else if selected == None {
		resp = processingCommandOnDefault(r, resp)
	} else {
		resp.Text = "Все рисунки данной сложности закончились. Если хотите порисовать снова, выберите сложность."
		resp.TTS = "Все рисунки данной сложности закончились. Если хотите порисовать снова, выберите сложность."
		resp.AddButton("Легкие", Payload {
			Text: "легкие",
		})
		resp.AddButton("Средние", Payload {
			Text: "средние",
		})
		resp.AddButton("Сложные", Payload {
			Text: "сложные",
		})
		resp.AddButton("Вперемешку", Payload {
			Text: "вперемешку",
		})
		selected = None
	}
	return resp
}

func downloadTemplate(r Request, resp Response) Response  {
	if newSession {
		resp.Card = NewLink(
			"https://cloud.mail.ru/public/hdax/G4wAX4pbU",
			"Шаблоны для рисования.",
			"Вы можете скачать необходимы шаблон, распечатать его и рисовать на нем.",
			457239115,
		)
		resp.Text = "Теперь выберите сложность."
		resp.AddButton("Легкие", Payload{
			Text: "легкие",
		})
		resp.AddButton("Средние", Payload{
			Text: "средние",
		})
		resp.AddButton("Сложные", Payload{
			Text: "сложные",
		})
		resp.AddButton("Вперемешку", Payload{
			Text: "вперемешку",
		})
	} else if selected != None && pictureDraw {
		resp = processingCommandOnDefault(r, resp)
	} else if selected == None {
		resp = processingCommandOnDefault(r, resp)
	} else {
		resp = processingCommandOnDefault(r, resp)
	}
	return resp
}

func appendIfMissing(slice []int, i int, numbPicture int) []int {
	for _, ele := range slice {
		if ele == i {
			return appendIfMissing(sequenceOfPictures, rand.Intn(numbPicture) + 1, numbPicture)
		}
	}
	return append(slice, i)
}

func levelSelection(resp Response) Response {
	switch selected {
	case Easy:
		switch sequenceOfPictures[exercise] {
		case 1:
			resp = easyLevel1(resp)
		case 2:
			resp = easyLevel2(resp)
		case 3:
			resp = easyLevel3(resp)
		case 4:
			resp = easyLevel4(resp)
		case 5:
			resp = easyLevel5(resp)
			/*case 6:
				resp = easyLevel6(resp)
				resp = mainButtons(resp)
			case 7:
				resp = easyLevel7(resp)
				resp = mainButtons(resp)
			case 8:
				resp = easyLevel8(resp)
				resp = mainButtons(resp)
			case 9:
				resp = easyLevel9(resp)
				resp = mainButtons(resp)
			case 10:
				resp = easyLevel10(resp)
				resp = mainButtons(resp)
			case 11:
				resp = easyLevel11(resp)
				resp = mainButtons(resp)
			case 12:
				resp = easyLevel12(resp)
				resp = mainButtons(resp)
			case 13:
				resp = easyLevel13(resp)
				resp = mainButtons(resp)
			case 14:
				resp = easyLevel14(resp)
				resp = mainButtons(resp)
			case 15:
				resp = easyLevel15(resp)
				resp = mainButtons(resp)*/
		}
	case Medium:
		switch sequenceOfPictures[exercise] {
		case 1:
			resp = mediumLevel1(resp)
		case 2:
			resp = mediumLevel2(resp)
		case 3:
			resp = mediumLevel3(resp)
		case 4:
			resp = mediumLevel4(resp)
		case 5:
			resp = mediumLevel5(resp)
		}
	case Hard:
		switch sequenceOfPictures[exercise] {
		case 1:
			resp = hardLevel1(resp)
		case 2:
			resp = hardLevel2(resp)
		case 3:
			resp = hardLevel3(resp)
		case 4:
			resp = hardLevel4(resp)
		case 5:
			resp = hardLevel5(resp)
		}
	case Mixed:
		switch sequenceOfPictures[exercise] {
		case 1:
			resp = easyLevel1(resp)
		case 2:
			resp = easyLevel2(resp)
		case 3:
			resp = easyLevel3(resp)
		case 4:
			resp = easyLevel4(resp)
		case 5:
			resp = easyLevel5(resp)
		case 6:
			resp = mediumLevel1(resp)
		case 7:
			resp = mediumLevel2(resp)
		case 8:
			resp = mediumLevel3(resp)
		case 9:
			resp = mediumLevel4(resp)
		case 10:
			resp = mediumLevel5(resp)
		case 11:
			resp = hardLevel1(resp)
		case 12:
			resp = hardLevel2(resp)
		case 13:
			resp = hardLevel3(resp)
		case 14:
			resp = hardLevel4(resp)
		case 15:
			resp = hardLevel5(resp)
		}
	}
	resp = mainButtons(resp)
	exercise++
	pictureSelected = true
	colorSelected = false
	pictureSample = false
	showBtnYesAndNo = false
	return resp
}

//Выбор цвета
func processingCommandColor(r Request, resp Response, command string) Response {
	if selected != None && !pictureDraw {
		color = command
		colorSelected = true
		pictureSelected = false
		pictureSample = false
		showBtnYesAndNo = false
		resp = drawingLevel(resp, color)
	} else if pictureDraw {
		resp = processingCommandOnDefault(r, resp)
	} else {
		resp = processingCommandOnDefault(r, resp)
	}
	return resp
}

//Переход к рисованию следующих клеток с тем же цветом
func processingCommandNext(r Request, resp Response) Response {
	if selected != None && color != "" && !pictureDraw {
		resp = drawingLevel(resp, color)
		colorSelected = true
		pictureSelected = false
		pictureSample = false
		showBtnYesAndNo = false
	} else if pictureDraw {
		resp = processingCommandOnDefault(r, resp)
	} else if selected == None {
		resp = processingCommandOnDefault(r, resp)
	} else {
		resp = processingDefaultColor(resp)
		resp = drawColorButtons(resp)
		resp = mainButtons(resp)
	}
	return resp
}

func processingCommandShowResult(r Request, resp Response) Response {
	if pictureSample {
		resp.Text = resultText
		resp.Card = NewBigImage(
			pictureID,
		)
		resp.AddButton("Следующий рисунок", Payload{
			Text: "следующий рисунок",
		})
		pictureSample = false
		pictureDraw = true
	} else if selected == None {
		resp = processingCommandOnDefault(r, resp)
	} else {
		resp = processingCommandOnDefault(r, resp)
	}

	return resp
}

func hintToGoToNextPicture(resp Response) Response {
	resp.Text = "Скажите мне «Следующий рисунок», если хотите продолжить рисовать."
	resp.AddButton("Следующий рисунок", Payload {
		Text: "следующий рисунок",
	})
	return resp
}

func chooseDifficulty(resp Response) Response {
	resp.Text = "Выберите сложность: легкие, средние, сложные или вариант «Вперемешку»."
	resp.TTS = "Выберите сложность:\n - легкие - средние - сложные или вариант «Вперемешку»."
	resp.AddButton("Легкие", Payload {
		Text: "легкие",
	})
	resp.AddButton("Средние", Payload {
		Text: "средние",
	})
	resp.AddButton("Сложные", Payload {
		Text: "сложные",
	})
	resp.AddButton("Вперемешку", Payload {
		Text: "вперемешку",
	})
	return resp
}

func mainButtons(resp Response) Response {
	resp.AddButton("Рисовать по картинке", Payload {
		Text: "рисовать по картинке",
	})
	resp.AddButton("Следующий рисунок", Payload {
		Text: "следующий рисунок",
	})
	return resp
}

