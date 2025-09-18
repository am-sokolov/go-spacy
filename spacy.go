package spacy

/*
#cgo CFLAGS: -I./include -I/Users/alexeysokolov/.pyenv/versions/3.12.4/include/python3.12
#cgo LDFLAGS: -L./lib -lspacy_wrapper -L/Users/alexeysokolov/.pyenv/versions/3.12.4/lib -lpython3.12
#include <stdlib.h>
#include "spacy_wrapper.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Token struct {
	Text    string
	Lemma   string
	POS     string
	Tag     string
	Dep     string
	IsStop  bool
	IsPunct bool
}

type Entity struct {
	Text  string
	Label string
	Start int
	End   int
}

type NLP struct {
	model string
}

func NewNLP(modelName string) (*NLP, error) {
	cModel := C.CString(modelName)
	defer C.free(unsafe.Pointer(cModel))

	ret := C.spacy_init(cModel)
	if ret != 0 {
		return nil, fmt.Errorf("failed to initialize Spacy with model: %s", modelName)
	}

	return &NLP{model: modelName}, nil
}

func (n *NLP) Close() {
	C.spacy_cleanup()
}

func (n *NLP) Tokenize(text string) []Token {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	tokenArray := C.spacy_tokenize(cText)
	defer C.free_token_array(&tokenArray)

	tokens := make([]Token, tokenArray.count)

	if tokenArray.count == 0 {
		return tokens
	}

	cTokens := (*[1 << 30]C.CToken)(unsafe.Pointer(tokenArray.tokens))[:tokenArray.count:tokenArray.count]

	for i, cToken := range cTokens {
		tokens[i] = Token{
			Text:    C.GoString(cToken.text),
			Lemma:   C.GoString(cToken.lemma),
			POS:     C.GoString(cToken.pos),
			Tag:     C.GoString(cToken.tag),
			Dep:     C.GoString(cToken.dep),
			IsStop:  bool(cToken.is_stop),
			IsPunct: bool(cToken.is_punct),
		}
	}

	return tokens
}

func (n *NLP) ExtractEntities(text string) []Entity {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	entityArray := C.spacy_extract_entities(cText)
	defer C.free_entity_array(&entityArray)

	entities := make([]Entity, entityArray.count)

	if entityArray.count == 0 {
		return entities
	}

	cEntities := (*[1 << 30]C.CEntity)(unsafe.Pointer(entityArray.entities))[:entityArray.count:entityArray.count]

	for i, cEntity := range cEntities {
		entities[i] = Entity{
			Text:  C.GoString(cEntity.text),
			Label: C.GoString(cEntity.label),
			Start: int(cEntity.start),
			End:   int(cEntity.end),
		}
	}

	return entities
}

func (n *NLP) SplitSentences(text string) []string {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sentArray := C.spacy_split_sentences(cText)
	defer C.free_sentence_array(&sentArray)

	sentences := make([]string, sentArray.count)

	if sentArray.count == 0 {
		return sentences
	}

	cSentences := (*[1 << 30]*C.char)(unsafe.Pointer(sentArray.sentences))[:sentArray.count:sentArray.count]

	for i, cSent := range cSentences {
		sentences[i] = C.GoString(cSent)
	}

	return sentences
}

func (n *NLP) POSTag(text string) map[string]string {
	tokens := n.Tokenize(text)
	posMap := make(map[string]string)

	for _, token := range tokens {
		posMap[token.Text] = token.POS
	}

	return posMap
}

func (n *NLP) GetDependencies(text string) map[string]string {
	tokens := n.Tokenize(text)
	depMap := make(map[string]string)

	for _, token := range tokens {
		depMap[token.Text] = token.Dep
	}

	return depMap
}

func (n *NLP) GetLemmas(text string) map[string]string {
	tokens := n.Tokenize(text)
	lemmaMap := make(map[string]string)

	for _, token := range tokens {
		lemmaMap[token.Text] = token.Lemma
	}

	return lemmaMap
}
