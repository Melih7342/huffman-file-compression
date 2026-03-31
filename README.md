# Huffman File Compression 🚀

A high-performance command-line interface (CLI) tool written in Go that compresses and decompresses files using the **Huffman Coding Algorithm**.



## ✨ Features

* **Lossless Compression:** Uses character frequency analysis to reduce file size without losing data.
* **Recursive Mode:** Process entire directory trees with a single command.
* **Performance:** Optimized bit manipulation and parallel processing using a worker pool.
* **Smart Overhead Detection:** Automatically skips files if the metadata overhead (JSON table) exceeds the compression savings.
* **Cross-Platform:** Native binaries available for Windows, Linux, and macOS.
* **Safety First:** Includes a "Roundtrip" validation logic to ensure data integrity.

---

## 🚀 Quick Start

### 1. Installation

Download the latest binary for your operating system from the [Releases](https://github.com/Melih7342/huffman-file-compression/releases) page or build it from source:

```bash
go build -o huffman ./main.go
```
### 2. Basic Usage

**Compress a file:**

```bash
./huffman -c -o archive.huff document.txt
```

**Decompress a file:**
```bash
./huffman -d -o restored.txt archive.huff
```

**Recursively compress a directory:**
./huffman -c -r -o backup_folder/ my_project/

## 🛠 Configuration (Flags)

| Flag | Description |
| :--- | :--- |
| `-c` | **Compress:** Enables compression mode. |
| `-d` | **Decompress:** Enables decompression mode. |
| `-o` | **Output:** Defines the target file path or directory. |
| `-r` | **Recursive:** Processes all files within subdirectories. |
| `-v` | **Verbose:** Displays detailed processing information in the console. |
| `-f` | **Force:** Forces compression even if the output file is larger than the original. |

---

## 🏗 Project Structure

* `main.go`: Entry point of the application.
* `internal/algorithm/`: Core logic (Tree building, Bit-handling, Path determination).
* `internal/models/`: Shared data structures and configuration logic.
* `internal/worker/`: Execution engine managing the file processing flow.

---

## 🧪 Testing

The project is backed by a robust test suite, including unit tests for core logic and integration tests for full file roundtrips.

```bash
# Run all tests
go test ./internal/algorithm/...
```
## 📝 How it Works

1. Frequency Analysis: The tool scans the input file to count the occurrences of each byte.
2. Tree Construction: A binary tree is built where the most frequent characters get the shortest bit-paths.
3. Metadata Header: The frequency table is stored as a JSON object after a unique HUFF file marker.
4. Bitstream Encoding: The content is translated into a bitstream and written to the compressed file.
