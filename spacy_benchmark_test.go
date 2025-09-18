package spacy

import (
	"fmt"
	"strings"
	"testing"
)

// Benchmark tokenization with different text sizes
func BenchmarkTokenizeSizes(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	sizes := []int{10, 100, 1000, 10000}
	baseText := "The quick brown fox jumps over the lazy dog. "

	for _, size := range sizes {
		text := strings.Repeat(baseText, size/10)
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tokens := nlp.Tokenize(text)
				if len(tokens) == 0 {
					b.Fatal("No tokens returned")
				}
			}
			b.ReportMetric(float64(len(text)), "bytes/op")
		})
	}
}

// Benchmark entity extraction
func BenchmarkEntityExtraction(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	texts := []string{
		"Simple text without entities.",
		"Apple Inc. is located in Cupertino.",
		"Barack Obama was born in Hawaii on August 4, 1961.",
		"Microsoft, Google, Amazon, Facebook, and Apple are major tech companies based in the United States.",
	}

	for i, text := range texts {
		b.Run(fmt.Sprintf("Text_%d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = nlp.ExtractEntities(text)
			}
		})
	}
}

// Benchmark sentence splitting
func BenchmarkSentenceSplitting(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	texts := []string{
		"Single sentence.",
		"First sentence. Second sentence. Third sentence.",
		strings.Repeat("This is a sentence. ", 10),
		strings.Repeat("Complex sentence with multiple clauses, subclauses, and various punctuation marks! ", 10),
	}

	for i, text := range texts {
		b.Run(fmt.Sprintf("Sentences_%d", i+1), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = nlp.SplitSentences(text)
			}
		})
	}
}

// Benchmark POS tagging
func BenchmarkPOSTagging(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	texts := []string{
		"The cat sat.",
		"The quick brown fox jumps over the lazy dog.",
		"Natural language processing enables computers to understand human language.",
		strings.Repeat("Complex sentences require more processing time. ", 10),
	}

	for i, text := range texts {
		b.Run(fmt.Sprintf("Text_%d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = nlp.POSTag(text)
			}
		})
	}
}

// Benchmark dependency parsing
func BenchmarkDependencyParsing(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	text := "The sophisticated algorithm efficiently processes complex linguistic structures."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nlp.GetDependencies(text)
	}
}

// Benchmark lemmatization
func BenchmarkLemmatization(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	text := "The children are playing with their toys and running around happily."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nlp.GetLemmas(text)
	}
}

// Benchmark full pipeline
func BenchmarkFullPipeline(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	text := "Apple Inc., founded by Steve Jobs in 1976, is now one of the world's most valuable companies. " +
		"The company produces consumer electronics, software, and online services."

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nlp.Tokenize(text)
		_ = nlp.ExtractEntities(text)
		_ = nlp.SplitSentences(text)
		_ = nlp.POSTag(text)
		_ = nlp.GetDependencies(text)
		_ = nlp.GetLemmas(text)
	}
}

// Benchmark memory allocation
func BenchmarkMemoryAllocation(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	text := "The quick brown fox jumps over the lazy dog."

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tokens := nlp.Tokenize(text)
		// Force use of tokens to prevent optimization
		if len(tokens) == 0 {
			b.Fatal("No tokens")
		}
	}
}

// Benchmark parallel processing
func BenchmarkParallel(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	texts := []string{
		"First test sentence for parallel processing.",
		"Second test sentence with different content.",
		"Third sentence to increase variety.",
		"Fourth sentence for load distribution.",
	}

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			text := texts[i%len(texts)]
			tokens := nlp.Tokenize(text)
			if len(tokens) == 0 {
				b.Fatal("No tokens returned in parallel")
			}
			i++
		}
	})
}

// Benchmark initialization time
func BenchmarkInitialization(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nlp, err := NewNLP("en_core_web_sm")
		if err != nil {
			b.Fatal(err)
		}
		nlp.Close()
	}
}

// Benchmark with caching effects
func BenchmarkCachingEffects(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	// Same text repeated - should benefit from any caching
	sameText := "The quick brown fox jumps over the lazy dog."

	b.Run("SameText", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = nlp.Tokenize(sameText)
		}
	})

	// Different texts - no caching benefit
	b.Run("DifferentTexts", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			text := fmt.Sprintf("Text number %d with unique content.", i)
			_ = nlp.Tokenize(text)
		}
	})
}

// Benchmark throughput
func BenchmarkThroughput(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	// Create a corpus of text
	corpus := strings.Repeat("The natural language processing system analyzes text efficiently. ", 1000)
	corpusBytes := len(corpus)

	b.ResetTimer()
	start := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tokens := nlp.Tokenize(corpus)
			if len(tokens) == 0 {
				b.Fatal("No tokens")
			}
		}
	})

	elapsed := start.T.Seconds()
	throughput := float64(corpusBytes) * float64(start.N) / elapsed / 1024 / 1024
	b.Logf("Throughput: %.2f MB/s", throughput)
}

// Compare performance of different functions
func BenchmarkCompareFunctions(b *testing.B) {
	nlp, err := NewNLP("en_core_web_sm")
	if err != nil {
		b.Fatalf("Failed to create NLP: %v", err)
	}
	defer nlp.Close()

	text := "Microsoft Corporation was founded by Bill Gates and Paul Allen on April 4, 1975."

	b.Run("Tokenize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = nlp.Tokenize(text)
		}
	})

	b.Run("ExtractEntities", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = nlp.ExtractEntities(text)
		}
	})

	b.Run("SplitSentences", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = nlp.SplitSentences(text)
		}
	})

	b.Run("POSTag", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = nlp.POSTag(text)
		}
	})

	b.Run("GetDependencies", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = nlp.GetDependencies(text)
		}
	})

	b.Run("GetLemmas", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = nlp.GetLemmas(text)
		}
	})
}
