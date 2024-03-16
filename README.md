# Package Calculator

## Description

The Package Calculator is a Go application that efficiently matches order quantities with specific package sizes. It is designed to help businesses determine the optimal packaging for any given number of items, using predefined package sizes.

## Installation

To use this application, you need to have Go installed on your machine. If you haven't already installed Go, you can download it from the official Go website: https://golang.org/dl/

After installing Go, you can set up the Package Calculator application with the following steps:

1. Clone the repository to your local machine:
   `git clone https://github.com/spicy-rice/package-calculator.git`

2. Navigate to the cloned repository's directory:
   `cd package-calculator`

3. Build the Go application:
   `go build -o package-calculator .`

This command compiles the application and generates an executable named `package-calculator` in the current directory. Replace `-o package-calculator` with the appropriate executable name for your operating system if needed, such as `package-calculator.exe` on Windows.

4. Run the Go application:

   `./package-calculator`

On Windows, use: `.\package-calculator.exe`

## Usage

After starting the application, it will be accessible at `http://localhost:8080`. Navigate to this address in your web browser to interact with the application.

To calculate the number of packages for a specific order quantity, enter the total number of items in the input field and submit the form. The application will then display the most efficient combination of package sizes based on the configuration defined.

## Configuration

The package sizes can be configured by editing the `config.json` file in the root directory of the project. This file contains an array of integers that represent the available package sizes.

Example `config.json`:

```json
{
  "packSizes": [5000, 2000, 1000, 500, 250]
}
```
