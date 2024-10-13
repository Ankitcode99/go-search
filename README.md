# Full-Text Search Engine in Go

## Project Overview

This project implements a full-text search engine in Go. It's designed to efficiently index and search through large volumes of text documents, with a focus on performance and scalability.

### Key Features

- Fast document indexing
- Concurrent processing for improved performance
- Support for large datasets (tested with Wikipedia abstracts)
- Simple and efficient search functionality

## Local Setup Instructions

### Prerequisites

- Go 1.15 or later
- Git

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/Ankitcode99/goSearch.git
   cd full-text-search-engine
   ```

2. Build the project:
   ```
   go build
   ```

### Running the Application

1. Prepare your dataset. The application expects an XML file with Wikipedia abstracts. You can download a sample file from [Wikipedia dumps](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz).

2. Run the application:
   ```
   go run main.go
   ```

   The application will start indexing the documents and then wait for search queries.

## Project Structure

- `main.go`: Entry point of the application
- `utils/index.go`: Contains the indexing logic
- `utils/document.go`: Defines the document structure
- `utils/search.go`: Implements the search functionality

## Usage

After running the application:

1. The program will index the documents from the specified XML file.
2. Once indexing is complete, you can enter search queries in the console.
3. The program will return relevant documents based on your search terms.

## Performance Considerations

- The application uses concurrent processing to speed up document indexing.
- Large datasets (e.g., full Wikipedia dumps) may require significant memory and processing time.

## Contributing

Contributions to improve the search engine are welcome. Please feel free to submit issues and pull requests.

