# Go-Spacy: Golang Bindings for Spacy NLP

Go-Spacy provides Golang bindings for the Spacy Natural Language Processing library through a C++ bridge layer. This allows you to use Spacy's powerful NLP capabilities directly from Go applications.

## Features

- **Tokenization**: Break text into individual tokens with linguistic attributes
- **Part-of-Speech (POS) Tagging**: Identify grammatical parts of speech
- **Named Entity Recognition (NER)**: Extract and classify named entities
- **Sentence Splitting**: Segment text into sentences
- **Dependency Parsing**: Analyze grammatical structure
- **Lemmatization**: Get base forms of words
- **Stop Words & Punctuation Detection**: Identify common words and punctuation

## Prerequisites

- Go 1.16 or higher
- Python 3.9+ with Spacy installed
- C++ compiler (g++ or clang++)
- Make

## Installation

### 1. Install Python Dependencies

```bash
pip install spacy
python -m spacy download en_core_web_sm
```

### 2. Build the C++ Wrapper

```bash
make clean
make
```

### 3. Install the Go Package

```bash
go get github.com/yourusername/spacy
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/spacy"
)

func main() {
    // Initialize NLP with a Spacy model
    nlp, err := spacy.NewNLP("en_core_web_sm")
    if err != nil {
        log.Fatal(err)
    }
    defer nlp.Close()

    text := "Apple Inc. was founded by Steve Jobs in California."

    // Tokenization
    tokens := nlp.Tokenize(text)
    for _, token := range tokens {
        fmt.Printf("Token: %s, POS: %s, Lemma: %s\n",
            token.Text, token.POS, token.Lemma)
    }

    // Named Entity Recognition
    entities := nlp.ExtractEntities(text)
    for _, entity := range entities {
        fmt.Printf("Entity: %s [%s]\n", entity.Text, entity.Label)
    }

    // Sentence Splitting
    sentences := nlp.SplitSentences(text)
    for _, sentence := range sentences {
        fmt.Println("Sentence:", sentence)
    }
}
```

## API Reference

### Types

#### Token
```go
type Token struct {
    Text    string  // Original token text
    Lemma   string  // Base form of the word
    POS     string  // Universal POS tag
    Tag     string  // Detailed POS tag
    Dep     string  // Dependency relation
    IsStop  bool    // Is it a stop word?
    IsPunct bool    // Is it punctuation?
}
```

#### Entity
```go
type Entity struct {
    Text  string  // Entity text
    Label string  // Entity type (PERSON, ORG, LOC, etc.)
    Start int     // Start position in text
    End   int     // End position in text
}
```

### Functions

#### NewNLP
```go
func NewNLP(modelName string) (*NLP, error)
```
Creates a new NLP instance with the specified Spacy model.

#### Tokenize
```go
func (n *NLP) Tokenize(text string) []Token
```
Tokenizes the input text and returns detailed token information.

#### ExtractEntities
```go
func (n *NLP) ExtractEntities(text string) []Entity
```
Extracts named entities from the text.

#### SplitSentences
```go
func (n *NLP) SplitSentences(text string) []string
```
Splits text into individual sentences.

#### POSTag
```go
func (n *NLP) POSTag(text string) map[string]string
```
Returns a map of tokens to their POS tags.

#### GetDependencies
```go
func (n *NLP) GetDependencies(text string) map[string]string
```
Returns a map of tokens to their dependency relations.

#### GetLemmas
```go
func (n *NLP) GetLemmas(text string) map[string]string
```
Returns a map of tokens to their lemmas.

#### Close
```go
func (n *NLP) Close()
```
Cleans up resources. Should be called when done using the NLP instance.

## Running Tests

```bash
make test
```

## Running Example

```bash
make run-example
```

## Benchmarks

Run benchmarks to test performance:

```bash
go test -bench=. -benchmem
```

## Project Structure

```
.
├── cpp/
│   └── spacy_wrapper.cpp    # C++ wrapper for Spacy
├── include/
│   └── spacy_wrapper.h      # C interface header
├── lib/
│   └── libspacy_wrapper.so  # Compiled shared library
├── example/
│   └── main.go              # Example usage
├── spacy.go                 # Go bindings
├── spacy_test.go           # Test suite
├── Makefile                # Build configuration
└── README.md               # This file
```

## Supported Spacy Models

This binding works with any Spacy model. Common models include:

- `en_core_web_sm` - Small English model
- `en_core_web_md` - Medium English model
- `en_core_web_lg` - Large English model
- `en_core_web_trf` - Transformer-based English model

Download additional models with:
```bash
python -m spacy download [model_name]
```

## Troubleshooting

### Python/Spacy Not Found

Ensure Python and Spacy are properly installed:
```bash
python --version
python -c "import spacy; print(spacy.__version__)"
```

### Build Errors

Check that Python development headers are installed:
```bash
# Ubuntu/Debian
sudo apt-get install python3-dev

# macOS
brew install python3

# Check python-config
python3.9-config --cflags --ldflags
```

### Model Not Found

Download the required Spacy model:
```bash
python -m spacy download en_core_web_sm
```

## Performance Considerations

- The NLP instance initialization is expensive; reuse instances when possible
- Consider using goroutines for parallel text processing
- The C++ bridge adds minimal overhead compared to pure Python Spacy

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- [Spacy](https://spacy.io/) - Industrial-strength Natural Language Processing
- Built with love for the Go community