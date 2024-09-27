
# Lem-in

## Project Overview

Lem-in is a graph-processing project built in Go, designed to simulate pathfinding algorithms on a graph. The project reads input files that define the graph's structure and computes paths between nodes. It includes examples of both valid and invalid input data for testing purposes.

## Key Features

- **Graph Processing**: Simulate and solve graph problems like shortest pathfinding.
- **Input Flexibility**: Accepts various input formats for testing different graph configurations.
- **Testing**: Includes test scripts and examples to validate the solution.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Running Tests](#running-tests)
- [Example Files](#example-files)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

- Go version 1.18+ is required.
- Make sure you have a shell environment (bash or similar) for running scripts.

### Steps

1. Clone the repository:

   ```bash
   git clone <repository-url>
   ```

2. Change to the project directory:

   ```bash
   cd lem-in
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

## Usage

You can run the main Go program as follows:

```bash
go run main.go <input-file>
```

Replace `<input-file>` with the path to one of the provided example files or your custom graph input.

Example:

```bash
go run main.go examples/example01.txt
```

### Input Format

The program expects an input file that describes the graph in a specific format, such as nodes, edges, and paths. Please refer to the provided example files to understand the expected input structure.

## Running Tests

The project includes a script to run tests against predefined input files:

```bash
./test.sh
```

This script automatically checks the program against several test cases, including valid and invalid inputs.

## Example Files

The `examples/` directory contains multiple example input files. These examples are categorized into valid and invalid inputs to help you understand the program's behavior:

- **Valid Examples**:
  - `example.txt`
  - `example01.txt`
  - `example02.txt`
  - ...
  
- **Invalid Examples**:
  - `badexample00.txt`
  - `badexample01.txt`

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a clear explanation of your changes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more information.
