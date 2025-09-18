CC = g++
PYTHON_VERSION = 3
PYTHON_CONFIG = python$(PYTHON_VERSION)-config
CFLAGS = -Wall -fPIC -O2 -std=c++17 -I./include $(shell $(PYTHON_CONFIG) --cflags)
LDFLAGS = $(shell $(PYTHON_CONFIG) --ldflags --embed)
TARGET = lib/libspacy_wrapper.so

SRCS = cpp/spacy_wrapper.cpp
OBJS = $(SRCS:.cpp=.o)

.PHONY: all clean test install-deps

all: $(TARGET)

$(TARGET): $(OBJS)
	@mkdir -p lib
	$(CC) -shared -o $@ $^ $(LDFLAGS)

%.o: %.cpp
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f $(OBJS) $(TARGET)
	rm -rf lib/*.so
	rm -rf *.test

install-deps:
	pip install spacy
	python -m spacy download en_core_web_sm

test: $(TARGET)
	go test -v ./...

build-go: $(TARGET)
	go build -v ./...

run-example: $(TARGET)
	go run example/main.go