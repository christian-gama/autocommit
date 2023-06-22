
[autocommit.webm](https://github.com/christian-gama/autocommit/assets/85251411/92b41c89-730c-4dfe-b6ee-05b6dd8e7264)

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

PS: Commits from this repository were generated using AutoCommit 🤖

## Features
- **AI-Powered Commit Messages**: AutoCommit uses OpenAI's language model to generate commit messages that are concise and meaningful.
- **Interactive CLI**: The tool comes with an interactive command-line interface that guides you through the commit process.
- **Customizable**: Customize the level of verbosity and other settings to suit your preferences.
- **Clipboard Integration**: Easily copy generated commit messages to the clipboard.
- **Completion Scripts**: AutoCommit comes with completion scripts for Bash, Zsh, Fish, and PowerShell.

## Installation
1. Clone the repository:
    ```shell
    git clone https://github.com/christian-gama/autocommit.git
    cd autocommit
    ```

2. Build the application (Make sure you have Go installed):
    ```shell
    make build
    ```

3. Install the application (Linux/macOS):
    ```shell
    make install
    ```

4. Add completion to your shell (Optional):
    Fish:
    ```shell
    autocommit completion fish > ~/.config/fish/completions/autocommit.fish
    exec fish
    ```

    Bash:
    ```shell
    autocommit completion bash | sudo tee /etc/bash_completion.d/autocommit > /dev/null
    exec bash
    ```

    Zsh:
    ```shell
    mkdir -p ~/.zsh/completions
    autocommit completion zsh > ~/.zsh/completions/_autocommit
    echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
    echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
    source ~/.zshrc
    ```

    PowerShell:
    ```shell
    autocommit completion powershell > autocommit.ps1
    # Then source autocommit.ps1 in your PowerShell profile
    ```

## Configuration
AutoCommit uses OpenAI's API for generating commit messages, so you'll need to have an API key from OpenAI. You can get one by signing up on [OpenAI's website](https://platform.openai.com/account/api-keys).

Once you have your API key, run the `autocommit` command. On the first run, AutoCommit will ask you for your OpenAI API key, preferred language model, and other settings. These settings will be stored locally on your machine for future use.

## Usage
1. Change to your git repository.
2. Make the changes in your repository that you want to commit.
3. Stage the changes.
4. Run the AutoCommit tool:
    ```text
    Autocommit is a CLI tool that uses OpenAI's models to generate commit messages based on the changes made in the repository.

    Usage:
    autocommit [flags]
    autocommit [command]

    Available Commands:
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    reset       Reset the configuration file
    set         Set configuration configs

    Flags:
    -h, --help   help for autocommit

    Use "autocommit [command] --help" for more information about a command.
    ```
5. Follow the interactive command-line interface. Choose whether to commit changes to git, generate a new commit message, copy the commit message to the clipboard, or exit the tool.
6. If you select the commit option, the tool will use the generated message to make a git commit.

## About Conventional Commits
AutoCommit generates commit messages that follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. It is a good practice to make small, incremental commits that are easy to understand and follow. If you're trying to make a commit that contains multiple changes, consider splitting it into multiple commits - it will make your commit history much cleaner and helps the AI generate better commit messages.

## Known Issues
- The tool currently only supports the English language model.
- Sometimes the tool ignores some instructions (gpt-3.5 model).
- Commits with a lot of changes may produce a meaningless commit message.

## License
AutoCommit is released under the [MIT License](https://opensource.org/licenses/MIT).
