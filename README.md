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

> - ```go
>   示例
>   c  := NewDefaultCard()
>   fields := []string{"**昵称**\n怪才君","**服务器**\n活动中心","**在线时间**\n9:00-21:00"}
>   c.AddFields(fields)
>   ```

- `func (c *cardMessage) AddColorText(tc TextColor, s string)` - 添加彩色文字，Kook提供四种字体颜色，

> - ```go
>   TextColorSuccess TextColor = "success"
>   TextColorPurple  TextColor = "purple"
>   TextColorWarning TextColor = "warning"
>   TextColorPink    TextColor = "pink"
>   
>   示例
>    c  := NewDefaultCard()
>    c.AddColorText(TextColorSuccess,"Hello! 这是一个测试")
>   ```

- ``func (c *cardMessage) AddTextWithPics(tt TextType, content string, imagePosition Position, imgUrl string, imgSize ImgSize)`` - 添加文本+图片。

> - ```go
>   c  := NewDefaultCard()
>   c.AddTextWithPics(`Kmarkdown`,"Hello! 这是一个测试",Left,"图片链接",ImgSizeSizeLg)
>   ```

太累了。我先B了，改明儿再写。

