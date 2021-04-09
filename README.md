# i18n4go
i18n4go is Go i18n library.


## Install
```shell
go get github.com/muzin/i18n4go
```


## Example

```go
i18n := NewI18N()

i18n.LoadLocale("en_US", "./testdata/locale_en_US.ini")
i18n.LoadLocale("zh_CN", "./testdata/locale_zh_CN.ini")

englishLocale, _ := i18n.getLocale("en_US")
chineseLocale, _ := i18n.getLocale("zh_CN")

englishMessage := englishLocale.GetMessage("terminal.print")
chineseMessage := chineseLocale.GetMessage("terminal.print")

t.Logf("print: en:%s zh:%s", englishMessage, chineseMessage)
```