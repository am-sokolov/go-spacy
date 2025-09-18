package main

import (
	"fmt"
	"log"

	"spacy"
)

func main() {
	nlp, err := spacy.NewNLP("en_core_web_sm")
	if err != nil {
		log.Fatalf("Failed to initialize Spacy: %v", err)
	}
	defer nlp.Close()

	text := "Apple Inc. was founded by Steve Jobs in Cupertino, California. The company is now worth over $2 trillion!"

	fmt.Println("=== Original Text ===")
	fmt.Println(text)
	fmt.Println()

	fmt.Println("=== Tokenization ===")
	tokens := nlp.Tokenize(text)
	for _, token := range tokens {
		fmt.Printf("Token: %-15s POS: %-10s Lemma: %-15s Tag: %-10s Dep: %-10s Stop: %-5v Punct: %-5v\n",
			token.Text, token.POS, token.Lemma, token.Tag, token.Dep, token.IsStop, token.IsPunct)
	}
	fmt.Println()

	fmt.Println("=== Named Entities ===")
	entities := nlp.ExtractEntities(text)
	for _, entity := range entities {
		fmt.Printf("Entity: %-20s Label: %-10s Position: [%d:%d]\n",
			entity.Text, entity.Label, entity.Start, entity.End)
	}
	fmt.Println()

	fmt.Println("=== Sentences ===")
	sentences := nlp.SplitSentences(text)
	for i, sentence := range sentences {
		fmt.Printf("Sentence %d: %s\n", i+1, sentence)
	}
	fmt.Println()

	fmt.Println("=== POS Tags ===")
	posMap := nlp.POSTag(text)
	for word, pos := range posMap {
		fmt.Printf("%-15s -> %s\n", word, pos)
	}
	fmt.Println()

	fmt.Println("=== Dependencies ===")
	depMap := nlp.GetDependencies(text)
	for word, dep := range depMap {
		fmt.Printf("%-15s -> %s\n", word, dep)
	}
	fmt.Println()

	fmt.Println("=== Lemmas ===")
	lemmaMap := nlp.GetLemmas(text)
	for word, lemma := range lemmaMap {
		if word != lemma {
			fmt.Printf("%-15s -> %s\n", word, lemma)
		}
	}
}
