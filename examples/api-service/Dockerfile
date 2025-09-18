# Example API service using Go-Spacy
FROM ubuntu:22.04 as builder

ENV DEBIAN_FRONTEND=noninteractive
ENV CGO_ENABLED=1

# Install build dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    python3 \
    python3-dev \
    python3-pip \
    wget \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ARG GO_VERSION=1.21.5
RUN wget -O go.tar.gz "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Python dependencies
RUN pip3 install --no-cache-dir spacy \
    && python3 -m spacy download en_core_web_sm

# Set working directory
WORKDIR /app

# Copy the main project files
COPY go.mod go.sum ./
COPY spacy.go ./
COPY cpp/ ./cpp/
COPY include/ ./include/
COPY Makefile ./

# Build the library
RUN make clean && make BUILD_MODE=release

# Copy example service code
COPY examples/api-service/ ./examples/api-service/

# Build the API service
RUN cd examples/api-service && go mod tidy && go build -o api-server .

# Runtime stage
FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Python runtime dependencies
RUN pip3 install --no-cache-dir spacy \
    && python3 -m spacy download en_core_web_sm

# Create non-root user
RUN groupadd -r apiuser && useradd -r -g apiuser apiuser

# Create app directory
RUN mkdir -p /app/lib && chown -R apiuser:apiuser /app

# Copy built artifacts
COPY --from=builder --chown=apiuser:apiuser /app/lib/ /app/lib/
COPY --from=builder --chown=apiuser:apiuser /app/examples/api-service/api-server /app/

# Set library path
ENV LD_LIBRARY_PATH="/app/lib"

# Switch to non-root user
USER apiuser

WORKDIR /app

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --start-period=15s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

CMD ["./api-server"]