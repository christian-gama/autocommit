# AutoCommit

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Configuration](#configuration)
5. [Usage](#usage)
6. [Contributing](#contributing)
7. [License](#license)

## Introduction
AutoCommit is a handy command-line tool that simplifies the git commit process by automatically generating meaningful commit messages using AI. Leveraging OpenAI's powerful language model, AutoCommit takes into account the changes made to the codebase and produces concise, descriptive commit messages that reflect the purpose and nature of those changes.

Git commits are an integral part of the software development process. However, crafting meaningful commit messages can sometimes be a time-consuming and even a daunting task. AutoCommit is designed to ease this process and help developers maintain a clean commit history without the overhead of thinking about commit messages for each and every change.

## Features
- **AI-Powered Commit Messages**: AutoCommit uses OpenAI's language model to generate commit messages that are concise and meaningful.
- **Interactive CLI**: The tool comes with an interactive command-line interface that guides you through the commit process.
- **Customizable**: Customize the level of verbosity and other settings to suit your preferences.
- **Clipboard Integration**: Easily copy generated commit messages to the clipboard.

## Installation
1. Clone the repository:
    ```shell
    git clone https://github.com/christian-gama/autocommit.git
    cd autocommit
    ```

2. Build the application:
    ```shell
    make build
    ```

3. Install the application (Linux/macOS):
    ```shell
    make install
    ```

## Configuration
AutoCommit uses OpenAI's API for generating commit messages, so you'll need to have an API key from OpenAI. You can get one by signing up on [OpenAI's website](https://platform.openai.com/account/api-keys).

Once you have your API key, run the `autocommit` command. On the first run, AutoCommit will ask you for your OpenAI API key, preferred language model, and other settings. These settings will be stored locally on your machine for future use.

## Usage
1. Change to your git repository.
2. Make the changes in your repository that you want to commit.
3. Stage the changes.
4. Run the AutoCommit tool:
    ```shell
    autocommit [COMMAND] [OPTIONS]

    Options (Optional)
    -v, --verbose   Enable verbose output (default: false)

    Command (Optional)
    reset           Reset the configuration settings
    ```
5. Follow the interactive command-line interface. Choose whether to commit changes to git, generate a new commit message, copy the commit message to the clipboard, or exit the tool.
6. If you select the commit option, the tool will use the generated message to make a git commit.

## Contributing
Contributions to AutoCommit are welcomed and appreciated. Please follow the standard GitHub flow for contributing:
1. Fork the repository.
2. Create a new feature branch.
3. Make your changes and commit them to your branch.
4. Create a pull request describing your changes.
5. The maintainers will review your pull request before merging.

## License
AutoCommit is released under the [MIT License](https://opensource.org/licenses/MIT).
