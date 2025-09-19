package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	spacy "github.com/am-sokolov/go-spacy"
)

type NLPService struct {
	nlp *spacy.NLP
}

type TokenizeRequest struct {
	Text string `json:"text"`
}

type TokenizeResponse struct {
	Tokens []spacy.Token `json:"tokens"`
	Count  int           `json:"count"`
}

type EntitiesRequest struct {
	Text string `json:"text"`
}

type EntitiesResponse struct {
	Entities []spacy.Entity `json:"entities"`
	Count    int            `json:"count"`
}

type AnalyzeRequest struct {
	Text string `json:"text"`
}

type AnalyzeResponse struct {
	Tokens      []spacy.Token     `json:"tokens"`
	Entities    []spacy.Entity    `json:"entities"`
	Sentences   []string          `json:"sentences"`
	NounChunks  []spacy.NounChunk `json:"noun_chunks"`
	TokenCount  int               `json:"token_count"`
	EntityCount int               `json:"entity_count"`
	ProcessedAt time.Time         `json:"processed_at"`
}

func NewNLPService(modelName string) (*NLPService, error) {
	nlp, err := spacy.NewNLP(modelName)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NLP: %w", err)
	}

	return &NLPService{nlp: nlp}, nil
}

func (s *NLPService) Close() {
	if s.nlp != nil {
		s.nlp.Close()
	}
}

func (s *NLPService) handleTokenize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TokenizeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Text field is required", http.StatusBadRequest)
		return
	}

	tokens := s.nlp.Tokenize(req.Text)
	response := TokenizeResponse{
		Tokens: tokens,
		Count:  len(tokens),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s *NLPService) handleEntities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EntitiesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Text field is required", http.StatusBadRequest)
		return
	}

	entities := s.nlp.ExtractEntities(req.Text)
	response := EntitiesResponse{
		Entities: entities,
		Count:    len(entities),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s *NLPService) handleAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Text field is required", http.StatusBadRequest)
		return
	}

	// Perform comprehensive analysis
	tokens := s.nlp.Tokenize(req.Text)
	entities := s.nlp.ExtractEntities(req.Text)
	sentences := s.nlp.SplitSentences(req.Text)
	nounChunks := s.nlp.GetNounChunks(req.Text)

	response := AnalyzeResponse{
		Tokens:      tokens,
		Entities:    entities,
		Sentences:   sentences,
		NounChunks:  nounChunks,
		TokenCount:  len(tokens),
		EntityCount: len(entities),
		ProcessedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s *NLPService) handleHealth(w http.ResponseWriter, _ *http.Request) {
	// Simple health check
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"service":   "go-spacy-api",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func main() {
	// Get configuration from environment
	modelName := os.Getenv("SPACY_MODEL")
	if modelName == "" {
		modelName = "en_core_web_sm"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize NLP service
	service, err := NewNLPService(modelName)
	if err != nil {
		log.Fatalf("Failed to initialize NLP service: %v", err)
	}
	defer service.Close()

	// Set up routes
	http.HandleFunc("/tokenize", service.handleTokenize)
	http.HandleFunc("/entities", service.handleEntities)
	http.HandleFunc("/analyze", service.handleAnalyze)
	http.HandleFunc("/health", service.handleHealth)

	// Root handler with API documentation
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		docs := map[string]interface{}{
			"service": "Go-Spacy NLP API",
			"version": "1.0.0",
			"model":   modelName,
			"endpoints": map[string]interface{}{
				"POST /tokenize": "Tokenize text into individual tokens",
				"POST /entities": "Extract named entities from text",
				"POST /analyze":  "Comprehensive text analysis",
				"GET /health":    "Service health check",
			},
			"example_request": map[string]string{
				"text": "Apple Inc. was founded by Steve Jobs in California.",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(docs); err != nil {
			log.Printf("Error encoding docs: %v", err)
		}
	})

	log.Printf("Starting Go-Spacy API server on port %s with model %s", port, modelName)
	log.Printf("Endpoints available:")
	log.Printf("  GET  /         - API documentation")
	log.Printf("  POST /tokenize - Tokenize text")
	log.Printf("  POST /entities - Extract entities")
	log.Printf("  POST /analyze  - Comprehensive analysis")
	log.Printf("  GET  /health   - Health check")

	// Create server with timeouts to address security concerns
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil, // Use default ServeMux
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
