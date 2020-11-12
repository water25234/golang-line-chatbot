package translate

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

var ti translate.TextInput

// Translate mean aws translate language
func Translate(sentence, sourceLang, targetLang, region string) string {
	client := translate.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithRegion(region),
	)

	ti.SetText(sentence)
	ti.SetSourceLanguageCode(sourceLang)
	ti.SetTargetLanguageCode(targetLang)
	result, err := client.Text(&ti)
	if err != nil {
		log.Printf("%v", err)
	}
	return *result.TranslatedText
}
