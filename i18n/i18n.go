package i18n

//go:generate go-bindata -pkg $GOPACKAGE -o i18n_assets.go resources/

import (
    "fmt"
    "os"

    log "github.com/sirupsen/logrus"
    "gopkg.in/yaml.v2"
)

const (
    LangEN = "en"
    LangID = "id"
)

type ResponseI18n map[string]map[string]string

var ResponseI18nMap map[string]ResponseI18n

func init() {
    ResponseI18nMap = make(map[string]ResponseI18n)

    langs := []string{LangEN, LangID}

    for _, lang := range langs {
        ResponseI18nMap[lang] = loadI18n(lang)
    }
}

func loadI18n(lang string) ResponseI18n {
    errorI18n := ResponseI18n{}
    errorBuf, err := Asset(fmt.Sprintf("resources/errors_%s.yml", lang))
    if err != nil {
        log.Fatalf("cannot load i18n file: %v", err)
        os.Exit(1)
    }
    if err:= yaml.Unmarshal(errorBuf, errorI18n); err != nil {
        log.Panic(err)
        os.Exit(1)
    }
    return errorI18n
}
