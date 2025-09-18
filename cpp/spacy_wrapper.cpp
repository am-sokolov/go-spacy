#include <Python.h>
#include <string>
#include <vector>
#include <memory>
#include <iostream>
#include <cstring>
#include <mutex>
#include "spacy_wrapper.h"

// Global mutex for Python GIL protection
static std::mutex python_mutex;
static bool python_initialized = false;

class SpacyWrapper {
private:
    PyObject* nlp;
    PyObject* spacyModule;

public:
    SpacyWrapper(const char* model_name) {
        std::lock_guard<std::mutex> lock(python_mutex);

        if (!python_initialized) {
            Py_Initialize();
            // PyEval_InitThreads() is deprecated and automatic in Python 3.7+
            python_initialized = true;

            // Add Python paths
            PyObject* sys = PyImport_ImportModule("sys");
            if (!sys) {
                PyErr_Print();
                throw std::runtime_error("Failed to import sys module");
            }

            PyObject* path = PyObject_GetAttrString(sys, "path");
            if (!path) {
                Py_DECREF(sys);
                throw std::runtime_error("Failed to get sys.path");
            }

            // Add common Python library paths
            PyList_Append(path, PyUnicode_FromString("."));
            PyList_Append(path, PyUnicode_FromString("./.venv/lib/python3.12/site-packages"));
            PyList_Append(path, PyUnicode_FromString("/opt/homebrew/lib/python3.9/site-packages"));
            PyList_Append(path, PyUnicode_FromString("/usr/local/lib/python3.9/site-packages"));

            Py_DECREF(path);
            Py_DECREF(sys);
        }

        spacyModule = PyImport_ImportModule("spacy");
        if (!spacyModule) {
            PyErr_Print();
            throw std::runtime_error("Failed to import spacy");
        }

        PyObject* loadFunc = PyObject_GetAttrString(spacyModule, "load");
        PyObject* args = PyTuple_Pack(1, PyUnicode_FromString(model_name));
        nlp = PyObject_CallObject(loadFunc, args);

        Py_DECREF(args);
        Py_DECREF(loadFunc);

        if (!nlp) {
            PyErr_Print();
            throw std::runtime_error("Failed to load spacy model");
        }
    }

    ~SpacyWrapper() {
        if (nlp) {
            Py_DECREF(nlp);
        }
        if (spacyModule) {
            Py_DECREF(spacyModule);
        }
        // Don't finalize Python as it causes issues with re-initialization
        // Py_Finalize();
    }

    std::vector<Token> tokenize(const std::string& text) {
        std::lock_guard<std::mutex> lock(python_mutex);
        std::vector<Token> tokens;

        PyObject* pyText = PyUnicode_FromString(text.c_str());
        PyObject* doc = PyObject_CallFunctionObjArgs(nlp, pyText, NULL);

        if (!doc) {
            PyErr_Print();
            Py_DECREF(pyText);
            return tokens;
        }

        PyObject* iterator = PyObject_GetIter(doc);
        PyObject* token;

        while ((token = PyIter_Next(iterator))) {
            Token t;

            PyObject* textAttr = PyObject_GetAttrString(token, "text");
            t.text = PyUnicode_AsUTF8(textAttr);
            Py_DECREF(textAttr);

            PyObject* lemmaAttr = PyObject_GetAttrString(token, "lemma_");
            t.lemma = PyUnicode_AsUTF8(lemmaAttr);
            Py_DECREF(lemmaAttr);

            PyObject* posAttr = PyObject_GetAttrString(token, "pos_");
            t.pos = PyUnicode_AsUTF8(posAttr);
            Py_DECREF(posAttr);

            PyObject* tagAttr = PyObject_GetAttrString(token, "tag_");
            t.tag = PyUnicode_AsUTF8(tagAttr);
            Py_DECREF(tagAttr);

            PyObject* depAttr = PyObject_GetAttrString(token, "dep_");
            t.dep = PyUnicode_AsUTF8(depAttr);
            Py_DECREF(depAttr);

            PyObject* isStopAttr = PyObject_GetAttrString(token, "is_stop");
            t.is_stop = PyObject_IsTrue(isStopAttr);
            Py_DECREF(isStopAttr);

            PyObject* isPunctAttr = PyObject_GetAttrString(token, "is_punct");
            t.is_punct = PyObject_IsTrue(isPunctAttr);
            Py_DECREF(isPunctAttr);

            tokens.push_back(t);
            Py_DECREF(token);
        }

        Py_DECREF(iterator);
        Py_DECREF(doc);
        Py_DECREF(pyText);

        return tokens;
    }

    std::vector<Entity> extractEntities(const std::string& text) {
        std::lock_guard<std::mutex> lock(python_mutex);
        std::vector<Entity> entities;

        PyObject* pyText = PyUnicode_FromString(text.c_str());
        PyObject* doc = PyObject_CallFunctionObjArgs(nlp, pyText, NULL);

        if (!doc) {
            PyErr_Print();
            Py_DECREF(pyText);
            return entities;
        }

        PyObject* ents = PyObject_GetAttrString(doc, "ents");
        PyObject* iterator = PyObject_GetIter(ents);
        PyObject* ent;

        while ((ent = PyIter_Next(iterator))) {
            Entity e;

            PyObject* textAttr = PyObject_GetAttrString(ent, "text");
            e.text = PyUnicode_AsUTF8(textAttr);
            Py_DECREF(textAttr);

            PyObject* labelAttr = PyObject_GetAttrString(ent, "label_");
            e.label = PyUnicode_AsUTF8(labelAttr);
            Py_DECREF(labelAttr);

            PyObject* startAttr = PyObject_GetAttrString(ent, "start_char");
            e.start = PyLong_AsLong(startAttr);
            Py_DECREF(startAttr);

            PyObject* endAttr = PyObject_GetAttrString(ent, "end_char");
            e.end = PyLong_AsLong(endAttr);
            Py_DECREF(endAttr);

            entities.push_back(e);
            Py_DECREF(ent);
        }

        Py_DECREF(iterator);
        Py_DECREF(ents);
        Py_DECREF(doc);
        Py_DECREF(pyText);

        return entities;
    }

    std::vector<std::string> splitSentences(const std::string& text) {
        std::lock_guard<std::mutex> lock(python_mutex);
        std::vector<std::string> sentences;

        PyObject* pyText = PyUnicode_FromString(text.c_str());
        PyObject* doc = PyObject_CallFunctionObjArgs(nlp, pyText, NULL);

        if (!doc) {
            PyErr_Print();
            Py_DECREF(pyText);
            return sentences;
        }

        PyObject* sents = PyObject_GetAttrString(doc, "sents");
        PyObject* iterator = PyObject_GetIter(sents);
        PyObject* sent;

        while ((sent = PyIter_Next(iterator))) {
            PyObject* textAttr = PyObject_GetAttrString(sent, "text");
            sentences.push_back(PyUnicode_AsUTF8(textAttr));
            Py_DECREF(textAttr);
            Py_DECREF(sent);
        }

        Py_DECREF(iterator);
        Py_DECREF(sents);
        Py_DECREF(doc);
        Py_DECREF(pyText);

        return sentences;
    }
};

// Use a single global instance with mutex protection
static std::unique_ptr<SpacyWrapper> wrapper;

extern "C" {

int spacy_init(const char* model_name) {
    if (!model_name || strlen(model_name) == 0) {
        std::cerr << "Error: model_name is null or empty" << std::endl;
        return -1;
    }

    try {
        wrapper = std::make_unique<SpacyWrapper>(model_name);
        return 0;
    } catch (const std::exception& e) {
        std::cerr << "Error initializing Spacy: " << e.what() << std::endl;
        return -1;
    }
}

void spacy_cleanup() {
    wrapper.reset();
}

TokenArray spacy_tokenize(const char* text) {
    TokenArray result = {nullptr, 0};

    if (!wrapper) {
        std::cerr << "Error: Spacy not initialized" << std::endl;
        return result;
    }

    if (!text) {
        return result;
    }

    auto tokens = wrapper->tokenize(text);
    if (tokens.empty()) {
        return result;
    }

    result.tokens = (CToken*)malloc(sizeof(CToken) * tokens.size());
    if (!result.tokens) {
        std::cerr << "Error: Failed to allocate memory for tokens" << std::endl;
        return result;
    }

    result.count = tokens.size();

    for (size_t i = 0; i < tokens.size(); ++i) {
        result.tokens[i].text = strdup(tokens[i].text.c_str());
        result.tokens[i].lemma = strdup(tokens[i].lemma.c_str());
        result.tokens[i].pos = strdup(tokens[i].pos.c_str());
        result.tokens[i].tag = strdup(tokens[i].tag.c_str());
        result.tokens[i].dep = strdup(tokens[i].dep.c_str());
        result.tokens[i].is_stop = tokens[i].is_stop;
        result.tokens[i].is_punct = tokens[i].is_punct;
    }

    return result;
}

void free_token_array(TokenArray* arr) {
    if (arr && arr->tokens) {
        for (size_t i = 0; i < arr->count; ++i) {
            free((void*)arr->tokens[i].text);
            free((void*)arr->tokens[i].lemma);
            free((void*)arr->tokens[i].pos);
            free((void*)arr->tokens[i].tag);
            free((void*)arr->tokens[i].dep);
        }
        free(arr->tokens);
        arr->tokens = nullptr;
        arr->count = 0;
    }
}

EntityArray spacy_extract_entities(const char* text) {
    EntityArray result = {nullptr, 0};

    if (!wrapper) {
        std::cerr << "Error: Spacy not initialized" << std::endl;
        return result;
    }

    if (!text) {
        return result;
    }

    auto entities = wrapper->extractEntities(text);
    if (entities.empty()) {
        return result;
    }

    result.entities = (CEntity*)malloc(sizeof(CEntity) * entities.size());
    if (!result.entities) {
        std::cerr << "Error: Failed to allocate memory for entities" << std::endl;
        return result;
    }

    result.count = entities.size();

    for (size_t i = 0; i < entities.size(); ++i) {
        result.entities[i].text = strdup(entities[i].text.c_str());
        result.entities[i].label = strdup(entities[i].label.c_str());
        result.entities[i].start = entities[i].start;
        result.entities[i].end = entities[i].end;
    }

    return result;
}

void free_entity_array(EntityArray* arr) {
    if (arr && arr->entities) {
        for (size_t i = 0; i < arr->count; ++i) {
            free((void*)arr->entities[i].text);
            free((void*)arr->entities[i].label);
        }
        free(arr->entities);
        arr->entities = nullptr;
        arr->count = 0;
    }
}

SentenceArray spacy_split_sentences(const char* text) {
    SentenceArray result = {nullptr, 0};

    if (!wrapper) {
        std::cerr << "Error: Spacy not initialized" << std::endl;
        return result;
    }

    if (!text) {
        return result;
    }

    auto sentences = wrapper->splitSentences(text);
    if (sentences.empty()) {
        return result;
    }

    result.sentences = (char**)malloc(sizeof(char*) * sentences.size());
    if (!result.sentences) {
        std::cerr << "Error: Failed to allocate memory for sentences" << std::endl;
        return result;
    }

    result.count = sentences.size();

    for (size_t i = 0; i < sentences.size(); ++i) {
        result.sentences[i] = strdup(sentences[i].c_str());
    }

    return result;
}

void free_sentence_array(SentenceArray* arr) {
    if (arr && arr->sentences) {
        for (size_t i = 0; i < arr->count; ++i) {
            free(arr->sentences[i]);
        }
        free(arr->sentences);
        arr->sentences = nullptr;
        arr->count = 0;
    }
}

}