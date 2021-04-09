package i18n4go

import (
	"testing"
)

func TestI18N_LoadLocale(t *testing.T) {
	t.Run("", func(t *testing.T) {

		i18n := NewI18N()

		i18n.LoadLocale("en_US", "./testdata/locale_en_US.ini")
		i18n.LoadLocale("zh_CN", "./testdata/locale_zh_CN.ini")

		englishLocale, _ := i18n.GetLocale("en_US")
		chineseLocale, _ := i18n.GetLocale("zh_CN")

		englishMessage := englishLocale.GetMessage("terminal.print")
		chineseMessage := chineseLocale.GetMessage("terminal.print")

		t.Logf("print: en:%s zh:%s", englishMessage, chineseMessage)
	})
}
