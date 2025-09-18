#ifndef SPACY_WRAPPER_H
#define SPACY_WRAPPER_H

#include <stdbool.h>
#include <stddef.h>

#ifdef __cplusplus
#include <string>

struct Token {
    std::string text;
    std::string lemma;
    std::string pos;
    std::string tag;
    std::string dep;
    bool is_stop;
    bool is_punct;
};

struct Entity {
    std::string text;
    std::string label;
    int start;
    int end;
};

extern "C" {
#endif

typedef struct {
    const char* text;
    const char* lemma;
    const char* pos;
    const char* tag;
    const char* dep;
    bool is_stop;
    bool is_punct;
} CToken;

typedef struct {
    CToken* tokens;
    size_t count;
} TokenArray;

typedef struct {
    const char* text;
    const char* label;
    int start;
    int end;
} CEntity;

typedef struct {
    CEntity* entities;
    size_t count;
} EntityArray;

typedef struct {
    char** sentences;
    size_t count;
} SentenceArray;

int spacy_init(const char* model_name);
void spacy_cleanup();

TokenArray spacy_tokenize(const char* text);
void free_token_array(TokenArray* arr);

EntityArray spacy_extract_entities(const char* text);
void free_entity_array(EntityArray* arr);

SentenceArray spacy_split_sentences(const char* text);
void free_sentence_array(SentenceArray* arr);

#ifdef __cplusplus
}
#endif

#endif