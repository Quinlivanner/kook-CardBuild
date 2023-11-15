package kook_CardBuild

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type CardTheme string
type ButtonTheme string
type ButtonType string
type CardSize string
type ImgSize string
type TextColor string
type TextType string
type Position string
type NotesType string
type CountdownMode string

const (
	CardSizeXs CardSize = "xs"
	CardSizeSm CardSize = "sm"
	CardSizeMd CardSize = "md"
	CardSizeLg CardSize = "lg"

	ImgSizeSizeSm ImgSize = "sm"
	ImgSizeSizeLg ImgSize = "lg"

	ImgType    string = "image"
	HeaderType string = "header"

	LinkButton   ButtonType = "link"
	ReturnBUtton ButtonType = "return-val"

	TextNoteType NotesType = "plain-text"
	ImgNotetype  NotesType = "image"

	DayCountdown    CountdownMode = "day"
	HourCountdown   CountdownMode = "hour"
	SecondCountdown CountdownMode = "second"
)
const (
	CardThemePrimary   CardTheme = "primary"
	CardThemeSecondary CardTheme = "secondary"
	CardThemeSuccess   CardTheme = "success"
	CardThemeWarning   CardTheme = "warning"
	CardThemeDanger    CardTheme = "danger"
	CardThemeInfo      CardTheme = "info"
	CardThemeNone      CardTheme = "none"

	ButtonThemePrimary   ButtonTheme = "primary"
	ButtonThemeSecondary ButtonTheme = "secondary"
	ButtonThemeSuccess   ButtonTheme = "success"
	ButtonThemeWarning   ButtonTheme = "warning"
	ButtonThemeDanger    ButtonTheme = "danger"
	ButtonThemeInfo      ButtonTheme = "info"
)
const (
	TextColorSuccess TextColor = "success"
	TextColorPurple  TextColor = "purple"
	TextColorWarning TextColor = "warning"
	TextColorPink    TextColor = "pink"
)
const (
	Kmarkdown TextType = "kmarkdown"
	Plain     TextType = "plain-text"
)
const (
	Left  Position = "left"
	Right Position = "right"
)

type cardMessage struct {
	Type    string        `json:"type"`
	Theme   CardTheme     `json:"theme"`
	Color   string        `json:"color"`
	Size    CardSize      `json:"size"`
	Modules []interface{} `json:"modules"`
}
type Text struct {
	Type string `json:"type"`
	Text struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"text"`
}
type KmarkdownText struct {
	Type string `json:"type"`
	Text struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"text"`
}
type Fields struct {
	Type string      `json:"type"`
	Text FieledsText `json:"text"`
}
type FieledsText struct {
	Type   string          `json:"type"`
	Cols   int             `json:"cols"`
	Fields []FiledsContent `json:"fields"`
}
type FiledsContent struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
type TextWithPics struct {
	Type string `json:"type"`
	Text struct {
		Type    TextType `json:"type"`
		Content string   `json:"content"`
	} `json:"text"`
	Mode      Position `json:"mode"`
	Accessory struct {
		Type string  `json:"type"`
		Src  string  `json:"src"`
		Size ImgSize `json:"size"`
	} `json:"accessory"`
}
type TextWithButton struct {
	Type string `json:"type"`
	Text struct {
		Type    TextType `json:"type"`
		Content string   `json:"content"`
	} `json:"text"`
	Mode      string `json:"mode"`
	Accessory struct {
		Type  string      `json:"type"`
		Theme ButtonTheme `json:"theme"`
		Text  struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"text"`
	} `json:"accessory"`
}
type Image struct {
	Type     string `json:"type"`
	Elements []struct {
		Type string `json:"type"`
		Src  string `json:"src"`
	} `json:"elements"`
}
type ImageGroup struct {
	Type     string            `json:"type"`
	Elements []imageUrlDetails `json:"elements"`
}
type imageUrlDetails struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}
type Header struct {
	Type string `json:"type"`
	Text struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"text"`
}
type Buttons struct {
	Type     string   `json:"type"`
	Elements []Button `json:"elements"`
}
type Button struct {
	Type  string      `json:"type"`
	Click string      `json:"click"`
	Theme ButtonTheme `json:"theme"`
	Value string      `json:"value"`
	Text  struct {
		Type    TextType `json:"type"`
		Content string   `json:"content"`
	} `json:"text"`
}
type ButtonCreate struct {
	ButtonType ButtonType
	Value      string
	Theme      ButtonTheme
	ButtonText string
}
type Notes struct {
	Type     string        `json:"type"`
	Elements []interface{} `json:"elements"`
}
type noteText struct {
	Type    NotesType `json:"type"`
	Content string    `json:"content"`
}
type noteImage struct {
	Type NotesType `json:"type"`
	Src  string    `json:"src"`
}
type NotesCreate struct {
	types NotesType
	Value string
}

type SecondCountDown struct {
	Type      string        `json:"type"`
	Mode      CountdownMode `json:"mode"`
	EndTime   int64         `json:"endTime"`
	StartTime int64         `json:"startTime"`
}
type DayHourCountDown struct {
	Type    string        `json:"type"`
	Mode    CountdownMode `json:"mode"`
	EndTime int64         `json:"endTime"`
}

func NewDefaultCard() *cardMessage {
	return &cardMessage{
		Color:   "#666666",
		Type:    "card",
		Theme:   CardThemeSecondary,
		Size:    CardSizeLg,
		Modules: []interface{}{},
	}
}
func NewCardWithOption(t CardTheme, s CardSize, c string) (*cardMessage, error) {
	if !IsHexCode(c) {
		return nil, errors.New("颜色十六进制代码不正确，请检查！")
	}
	return &cardMessage{
		Type:    "card",
		Theme:   t,
		Color:   c,
		Size:    s,
		Modules: []interface{}{},
	}, nil
}

func (c *cardMessage) SetCardTheme(t CardTheme) {
	c.Theme = t
	return
}
func (c *cardMessage) SetCardColor(col string) error {
	if !IsHexCode(col) {
		return fmt.Errorf("颜色十六进制代码不正确，请检查！")
	}
	c.Color = col
	return nil
}
func (c *cardMessage) SetCardSize(s CardSize) {
	c.Size = CardSizeMd
	return
}

func (c *cardMessage) AddText(s string) {
	c.Modules = append(c.Modules, Text{Type: "section", Text: struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}{Type: "plain-text", Content: s}})
	return
}
func (c *cardMessage) AddKmarkdown(s string) {
	c.Modules = append(c.Modules, KmarkdownText{Type: "section", Text: struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}{Type: "kmarkdown", Content: s}})
	return
}
func (c *cardMessage) AddFields(s []string) {
	var filedsContent = []FiledsContent{}
	for _, v := range s {
		filedsContent = append(filedsContent, FiledsContent{Type: "kmarkdown", Content: v})
	}
	c.Modules = append(c.Modules, Fields{Type: "section",
		Text: FieledsText{
			Type:   "paragraph",
			Cols:   len(s),
			Fields: filedsContent,
		}})
	return
}
func (c *cardMessage) AddColorText(tc TextColor, s string) {
	c.Modules = append(c.Modules, KmarkdownText{Type: "section", Text: struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}{Type: "kmarkdown", Content: fmt.Sprintf("(font)%s(font)[%s]", s, tc)}})
	return
}
func (c *cardMessage) AddTextWithPics(tt TextType, content string, imagePosition Position, imgUrl string, imgSize ImgSize) {
	c.Modules = append(c.Modules, TextWithPics{Type: "section", Text: struct {
		Type    TextType `json:"type"`
		Content string   `json:"content"`
	}{Type: tt, Content: content}, Mode: imagePosition, Accessory: struct {
		Type string  `json:"type"`
		Src  string  `json:"src"`
		Size ImgSize `json:"size"`
	}{Type: ImgType, Src: imgUrl, Size: imgSize}})
	return
}
func (c *cardMessage) AddTextWithButton(tt TextType, content string, buttonTheme ButtonTheme, buttonContent string) {
	c.Modules = append(c.Modules, TextWithButton{Type: "section", Text: struct {
		Type    TextType `json:"type"`
		Content string   `json:"content"`
	}{Type: tt, Content: content}, Mode: "right", Accessory: struct {
		Type  string      `json:"type"`
		Theme ButtonTheme `json:"theme"`
		Text  struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"text"`
	}{Type: "button", Theme: buttonTheme, Text: struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}{Type: "plain-text", Content: buttonContent}}})
	return
}
func (c *cardMessage) AddImage(imgUrl string) {
	c.Modules = append(c.Modules, Image{Type: "container", Elements: []struct {
		Type string `json:"type"`
		Src  string `json:"src"`
	}{{
		Type: ImgType,
		Src:  imgUrl,
	}}})
	return
}
func (c *cardMessage) AddImageGroup(imgUrlGroup []string) {
	var ig = []imageUrlDetails{}
	for _, v := range imgUrlGroup {
		ig = append(ig, imageUrlDetails{Type: ImgType, Src: v})
	}
	c.Modules = append(c.Modules, ImageGroup{Type: "image-group", Elements: ig})
	return
}
func (c *cardMessage) AddHeader(content string) {
	c.Modules = append(c.Modules, Header{Type: HeaderType, Text: struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}{Type: "plain-text", Content: content}})
	return
}
func (c *cardMessage) AddDivider() {
	c.Modules = append(c.Modules, struct {
		Type string `json:"type"`
	}{Type: "divider"})
	return
}
func (c *cardMessage) AddButtonGroup(ButtonGroups []ButtonCreate) {
	var temB = []Button{}
	for _, v := range ButtonGroups {
		if v.ButtonType == LinkButton {
			temB = append(temB, Button{Type: "button", Click: "link", Theme: v.Theme, Value: v.Value, Text: struct {
				Type    TextType `json:"type"`
				Content string   `json:"content"`
			}{Type: Plain, Content: v.ButtonText}})
		}
		if v.ButtonType == ReturnBUtton {
			temB = append(temB, Button{Type: "button", Click: "return-val", Theme: v.Theme, Value: v.Value, Text: struct {
				Type    TextType `json:"type"`
				Content string   `json:"content"`
			}{Type: Plain, Content: v.ButtonText}})
		}
	}
	c.Modules = append(c.Modules, Buttons{Type: "action-group", Elements: temB})
	return
}
func (c *cardMessage) AddNotes(notesDetails []NotesCreate) {
	var temN = []interface{}{}
	for _, v := range notesDetails {
		if v.types == TextNoteType {
			temN = append(temN, noteText{Type: TextNoteType, Content: v.Value})
		}
		if v.types == ImgNotetype {
			temN = append(temN, noteImage{Type: ImgNotetype, Src: v.Value})
		}
	}
	c.Modules = append(c.Modules, Notes{Type: "context", Elements: temN})
	return
}
func (c *cardMessage) AddNormalCountdown(countdownMode CountdownMode, endTime, startTime int64) {
	if countdownMode == SecondCountdown {
		c.Modules = append(c.Modules, SecondCountDown{Type: "countdown", Mode: countdownMode, EndTime: endTime, StartTime: startTime})
	} else {
		c.Modules = append(c.Modules, DayHourCountDown{Type: "countdown", Mode: countdownMode, EndTime: endTime})
	}
	return
}

func IsHexCode(str string) bool {
	// 创建匹配十六进制代码的正则表达式
	regex := regexp.MustCompile(`^#?([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`)
	// 使用正则表达式进行匹配
	return regex.MatchString(str)
}

func GenerateCardMessageContent(cards ...*cardMessage) (res string, err error) {
	var cardMessageContent = []*cardMessage{}
	for _, v := range cards {
		cardMessageContent = append(cardMessageContent, v)
	}
	marshal, err := json.Marshal(cardMessageContent)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
