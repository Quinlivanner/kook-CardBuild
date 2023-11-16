# 欢迎使用 `kook-CardBuild`

## 如何使用

你只需要专注于你的卡片编辑，无需担心其他问题，在卡片编辑结束后，调用`GenerateCardMessageContent`方法直接生成数据。

**创建卡片**：我们提供两种初始化卡片

- **`NewDefaultCard()`**  初始化一张默认的卡片。
- **`NewCardWithOption(t CardTheme, s CardSize, c string)`**  初始化一张自定义参数的卡片，在当前规则下，我建议使用`c string`即`color`参数来控制卡片消息的左边框颜色。

> - **`SetCardTheme()`**
> - **`SetCardColor`()**
> - **`SetCardSize()`**
>
> > - 以上均可作为`cardMessage`的方法。

**添加内容** - 我们提供以下内容模块构造。

- `func (c *cardMessage) AddText(s string)`  - 添加纯文本。仅需要一个文本参数。
- `func (c *cardMessage) AddKmarkdown(s string)` - 添加`Kmarkdown`文本。仅需要一个文本参数。
- `func (c *cardMessage) AddFields(s []string)` - 添加多列文本。接受`[]string`参数，其中sting为文本。请注意，多列文本的上下展示效果由`\n`实现。

```go
package main

import (
	"github.com/Quinlivanner/kook-CardBuild"
)

func CardMessageCreate() (error, string) {

	c := kook_CardBuild.NewDefaultCard()
	fields := []string{"**昵称**\n怪才君", "**服务器**\n活动中心", "**在线时间**\n9:00-21:00"}
	c.AddFields(fields)

	res, err := kook_CardBuild.GenerateCardMessageContent(c)
	if err != nil {
		return err, ""
	}
	return nil, res
}
```

- `func (c *cardMessage) AddColorText(tc TextColor, s string)` - 添加彩色文字，Kook提供四种字体颜色，

> - ```go
>   TextColorSuccess TextColor = "success"
>   TextColorPurple  TextColor = "purple"
>   TextColorWarning TextColor = "warning"
>   TextColorPink    TextColor = "pink"
>   
>   示例
>    c  := kook_CardBuild.NewDefaultCard()
>    c.AddColorText(kook_CardBuild.TextColorSuccess,"Hello! 这是一个测试")
>   ```

- ``func (c *cardMessage) AddTextWithPics(tt TextType, content string, imagePosition Position, imgUrl string, imgSize ImgSize)`` - 添加文本+图片。

> - ```go
>   c  := kook_CardBuild.NewDefaultCard()
>   c.AddTextWithPics(`Kmarkdown`,"Hello! 这是一个测试",Left,"图片链接",ImgSizeSizeLg)
>   ```

**其他方法以此类推**
**最后调用`kook_CardBuild.GenerateCardMessageContent()` 生成 `Post` 参数中的 `Content`**

```go
package main

import (
	"github.com/Quinlivanner/kook-CardBuild"
)

func CardMessageCreate() (error, string) {

	c := kook_CardBuild.NewDefaultCard()

	c.SetCardSize(kook_CardBuild.CardSizeLg)
	err := c.SetCardColor("#FFFFFF")
	if err != nil {
		return err, ""
	}

	c.AddText("Hello! 这是一个测试")
	c.AddKmarkdown("Hello! 这是一个测试")

	fields := []string{"**昵称**\n怪才君", "**服务器**\n活动中心", "**在线时间**\n9:00-21:00"}
	c.AddFields(fields)

	c.AddColorText(kook_CardBuild.TextColorSuccess, "Hello! 这是一个测试")

	c.AddTextWithPics(`Kmarkdown`, "Hello! 这是一个测试", kook_CardBuild.Left, "https://img.kookapp.cn/assets/avatar_2.jpg", kook_CardBuild.ImgSizeSizeLg)

	c.AddTextWithButton(kook_CardBuild.Kmarkdown, "点击按钮参与测试", kook_CardBuild.ButtonThemeDanger, "测试按钮")

	c.AddImage("https://img.kookapp.cn/assets/avatar_2.jpg")

	imgUrls := []string{"https://img.kookapp.cn/assets/avatar_2.jpg", "https://img.kookapp.cn/assets/avatar_2.jpg", "https://img.kookapp.cn/assets/avatar_2.jpg"}
	c.AddImageGroup(imgUrls)

	c.AddHeader("Hello！这是一个测试")

	//分割线，无需任何内容
	c.AddDivider()

	tb := []kook_CardBuild.ButtonCreate{}
	//添加跳转按钮，点击按钮后，将会跳转到Value中的链接
	tb = append(tb, kook_CardBuild.ButtonCreate{ButtonType: kook_CardBuild.LinkButton, Value: "https://img.kookapp.cn/assets/avatar_2.jpg",
		Theme: kook_CardBuild.ButtonThemeDanger, ButtonText: "测试按钮"})
	//添加返回值按钮，点击按钮后，返回的value参数的值。
	tb = append(tb, kook_CardBuild.ButtonCreate{ButtonType: kook_CardBuild.ReturnBUtton, Value: "Clicked",
		Theme: kook_CardBuild.ButtonThemeDanger, ButtonText: "测试按钮"})
	c.AddButtonGroup(tb)

	tn := []kook_CardBuild.NotesCreate{}
	//文字
	tn = append(tn, kook_CardBuild.NotesCreate{Types: kook_CardBuild.TextNoteType, Value: "这是一个测试"})
	//图片
	tn = append(tn, kook_CardBuild.NotesCreate{Types: kook_CardBuild.ImgNotetype, Value: "https://img.kookapp.cn/assets/avatar_2.jpg"})
	c.AddNotes(tn)

	//读秒倒计时，倒计时中，仅有读秒倒计时的starttime需要具体参数
	c.AddCountdown(kook_CardBuild.SecondCountdown, 1700149391000, 1700149391000)
	//小时倒计时
	c.AddCountdown(kook_CardBuild.HourCountdown, 1700149391000, 0)
	//常规倒计时
	c.AddCountdown(kook_CardBuild.DayCountdown, 1700149391000, 0)

	c1 := kook_CardBuild.NewDefaultCard()
	c1.AddColorText(kook_CardBuild.TextColorSuccess, "Hello! 这是一个测试")

	res, err := kook_CardBuild.GenerateCardMessageContent(c, c1)
	if err != nil {
		return err, ""
	}
	return nil, res
}
```

太累了。我先B了，改明儿再写。

