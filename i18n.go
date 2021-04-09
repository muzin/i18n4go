package i18n4go

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type TMsgs map[string]string

type I18N struct {
	Locales map[string]*Locale
}

func NewI18N() *I18N {
	i18n := &I18N{}
	i18n.Locales = make(map[string]*Locale)
	return i18n
}

func (this *I18N) GetLocale(lang string) (*Locale, error) {
	locale, ok := this.Locales[lang]
	if ok {
		return locale, nil
	} else {
		return nil, errors.New("not exists")
	}
}

// 通过 文件 加载 配置信息
func (this *I18N) LoadLocale(lang string, iniPath string) error {
	_, err := os.Lstat(iniPath)
	if !os.IsNotExist(err) {
		fileByte, err := ioutil.ReadFile(iniPath)
		if err != nil {
			return err
		} else {
			return this.LoadLocaleFromStream(lang, fileByte)
		}
	} else {
		return errors.New(iniPath + " is not exists")
	}
}

// 通过 文件 内容 加载 配置信息
func (this *I18N) LoadLocaleFromString(lang string, iniStr string) error {

	tMsgs := parseIniFile([]byte(iniStr))

	locale := &Locale{
		Language: lang,
		TMsgs:    tMsgs,
	}

	this.Locales[lang] = locale

	return nil
}

// 通过 文件 内容 加载 配置信息
func (this *I18N) LoadLocaleFromStream(lang string, iniByte []byte) error {

	tMsgs := parseIniFile(iniByte)

	locale := &Locale{
		Language: lang,
		TMsgs:    tMsgs,
	}

	this.Locales[lang] = locale

	return nil
}

type Locale struct {
	Language string
	TMsgs    TMsgs
}

func (this *Locale) GetMessage(name string) string {
	val, ok := this.TMsgs[name]
	if ok {
		return val
	} else {
		return ""
	}
}

// 解析 ini 的内容 加载 配置信息
func parseIniFile(iniByte []byte) TMsgs {

	iniStr := string(iniByte)
	iniStr = strings.TrimSpace(iniStr)

	// 1. 内容去空格
	// 2. 按照 换行 分割
	// 3. 过滤掉 以 #号开头的内容

	tMsgs := make(TMsgs)

	iniStrArr := strings.Split(iniStr, "\n")

	var iniStrArrWithoutNote []string = make([]string, 0)
	for _, iniStr := range iniStrArr {
		iniStrWithoutSpace := strings.TrimSpace(iniStr)
		if strings.Index(iniStrWithoutSpace, "#") != 0 {
			iniStrArrWithoutNote = append(iniStrArrWithoutNote, iniStrWithoutSpace)
		}
	}

	var title = ""
	for _, str := range iniStrArrWithoutNote {
		strWithoutSpace := strings.TrimSpace(str)

		// 如果 是 标题
		titleStartIdx := strings.Index(strWithoutSpace, "[")
		titleEndIdx := strings.Index(strWithoutSpace, "]")
		if titleStartIdx == 0 &&
			titleEndIdx == (len(strWithoutSpace)-1) {
			title = strings.TrimSpace(strWithoutSpace[(titleStartIdx + 1):titleEndIdx])
			continue
		}

		// 如果 行中 有 =
		splitIdx := strings.Index(str, "=")
		if splitIdx > -1 {
			key := strings.TrimSpace(str[0:splitIdx])
			val := strings.TrimSpace(str[splitIdx+1:])
			if len(title) > 0 {
				key = title + "." + key
			}
			tMsgs[key] = val
		}
	}

	title = ""

	return tMsgs
}
