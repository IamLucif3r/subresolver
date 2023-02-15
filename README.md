# SubResolver
SubResolver is a command line tool for converting a list of subdomains into their corresponding IP addresses. It is written in Go language and can be easily installed on any machine with Go installed.

## Installation

To install SubResolver, follow the procedure:

1. Clone the repository using:

```bash
git clone https://github.com/IamLucif3r/SubResolver.git
```
2. Navigate to the repository directory using: 

```bash
cd SubResolver
```

3. Build the tool using:

```bash 
go build
```

4. Install the tool using:

```bash 
go install
```

## Usage

To use SubResolver, run the following command in your terminal:

```bash
SubResolver <input_file> [options]
```

Where `<input_file>` is the name of the file containing the list of subdomains and `[options]` are optional parameters.


## Options

- **`-o`**: Specify the name of the file to which the output should be saved. If not provided, the output will be printed to the console.
