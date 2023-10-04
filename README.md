# Ascii-Art Generator

## Description
Go program hosts html server, handle user input and display the suitable ascii-art.

## Author

Yuki Kaneko

## Usage: how to run
1. Initiate the html server on command line:
```
go run .
```
2. Go to the server http://localhost:8080/

3. Enter the text you want to generate as ascii-art, select the lendering type and hit "Generate!" button.

## Test (audit)
Follow the instructions in the audit page: https://github.com/01-edu/public/tree/master/subjects/ascii-art-web/audit#does-the-project-run-using-an-api

### Audit texts:
standard render:
```
{123}
<Hello> (World)!
```
```
123??
```
shadow render:
```
$% "=
```
thinkertoy render:
```
123 T/fs#R 
```

## Implementation Details: Algorithm

### main.go

The main.go file defines two HTTP request handlers:

1. DisplayPage Handler:
   - Responsible for rendering the HTML page for generating ASCII art.
   - Parses the HTML template and renders it.
   - Handles template parsing errors and returns a "Not Found (404)" response if the template is missing.

2. GenerateAsciiArtHandler Handler:
   - Processes the user's input data to generate ASCII art.
   - Validates input parameters and returns a "Bad Request (400)" response for invalid input.
   - Uses the GenerateAsciiArt function to create the ASCII art.
   - Handles template parsing errors and returns a "Not Found (404)" response if the template is missing.
   - Returns an "Internal Server Error (500)" response for unhandled errors.

These handlers handle various HTTP status codes to provide meaningful responses to users based on the success or failure of the request.


### ascii-art.go

It includes GenerateAsciiArt function, which is responsible for generating ASCII art based on the provided input text and rendering type. Here is an overview of the algorithm:

1. Input Text Processing:

   The input text is split into individual lines in the `textLines` by the newline character (`\n`).

1. ASCII Art Definitions:

   The ASCII art characters are defined in the artList variable, which includes a list of characters that can be represented in ASCII art.

1. Reading ASCII Art (txt for rendering):

   - The function reads an txt file based on the provided rendering type.
   - The read byte is split into string array by each word of ascii-art.  

1. Creating the ASCII Art by input lines:

   - For each line of the input text (an element of `textLines`), the function looks up the corresponding ASCII art representation for each character and assembles the art lines accordingly.
   - The generated ASCII art lines are formatted to place the characters in the correct order for printing.

1. Concluding the Result:

   - The final ASCII art result is constructed as a string, with special handling for newlines in the input text.

1. Return the Result:
   - The generated ASCII art is returned as a string.

