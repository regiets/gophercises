# Quiz Game

A simple CLI quiz app, allowing custom questions and answers to be passed in in the form of a CSV file.

### Getting started
1. Build the app
  ```sh
  go build ./main.go
  ```
2. Start the quiz [with optional flags]
  ```sh
  ./main [-limit=30] [-csv=problems.csv]
  ```

### Options
- **limit:** time limit in seconds for the completion of the quiz (default: 30s)
- **csv:** questions and answers as a CSV file in the format of "question,answer"
