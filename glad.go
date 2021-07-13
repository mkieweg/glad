package glad

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const entry = `
\newglossaryentry{<< .Acronym >>}{
    name={<< .AcronymUpper >>},
	first={<< .Full >> (<< .AcronymUpper >>)},
    description={<< .Description >>}
}
`

type Entry struct {
	Acronym      string
	AcronymUpper string
	Full         string
	Description  string
}

func Add(acronym, full, description string) {
	filename := viper.GetString("glossary-file")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()

	if acronym == "" {
		b := strings.Builder{}
		full := strings.ReplaceAll(full, "-", " ")
		words := strings.Split(full, " ")
		for _, w := range words {
			b.WriteByte(w[0])
		}
		acronym = b.String()
	}
	e := Entry{
		Acronym:      strings.ToLower(acronym),
		AcronymUpper: strings.ToUpper(acronym),
		Full:         full,
		Description:  description,
	}

	t := template.New("entry")
	_, err = t.Delims("<<", ">>").Parse(entry)
	if err != nil {
		log.Error(err)
		return
	}
	if err := t.Execute(f, e); err != nil {
		log.Error(err)
		return
	}
	path, err := filepath.Abs(filename)
	if err != nil {
		log.Error(err)
	}
	log.WithFields(log.Fields{
		"acronym":           e.Acronym,
		"uppercase acronym": e.AcronymUpper,
		"full":              full,
		"description":       description,
		"glossary file":     path,
	}).Info("entry written")
}
