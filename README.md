# AutoCommit

AutoCommit generates git commit messages from your staged changes using AI. It currently supports OpenAI, Ollama 2, Mistral, Groq and Google AI models.

## Features

- AI-driven commit messages based on diffs
- Interactive CLI: commit, copy, regenerate, add instructions, or exit
- Editable instruction template
- Shell completion scripts for Bash, Zsh, Fish, PowerShell

## Installation

### Go

Ensure you have go installed and run:

```sh
go install github.com/christian-gama/autocommit/v2@latest
```

### Manual

```sh
git clone https://github.com/christian-gama/autocommit.git
cd autocommit
make build
make install # Linux/macOS
```

### Enabling Shell Completion (Optional)

Enable shell completion by following the `autocommit completion --help` instructions. If you need help to set up for a specific shell you can run `autocommit completion [shell] --help`. Available shells are:

- bash
- zsh
- fish
- powershell

Example for Zsh (macOS):

```sh
echo "autoload -U compinit; compinit" >> ~/.zshrc
autocommit completion zsh > $(brew --prefix)/share/zsh/site-functions/_autocommit
```

## Configuration

Run `autocommit configure` to choose a provider. Set your credential (such as API Key) and then select your preferred model. Your settings are stored locally.

### Supported Providers

- **OpenAI**

  - Requires an API key from [OpenAI's platform](https://platform.openai.com/api-keys)
  - Supported models:
    - GPT-4o
    - GPT-4.1
    - GPT-4.1-mini
    - GPT-4.1-nano
    - o1
    - o1-mini
    - o3
    - o3-mini
    - o4-mini

- **Ollama**

  - Requires [Ollama](https://ollama.ai/) to be installed and running locally
  - No API key needed, but:
    1. The Ollama service must be active (run `ollama serve` before invoking)
    2. You must have the models you want to use already downloaded locally
       (e.g., run `ollama pull llama4` to download llama4 before selecting it in the autocommit configuration interface)
  - Supported models:
    - gemma:1b, gemma:4b, gemma:12b, gemma:27b
    - qwen3:0.6b, qwen3:1.7b, qwen3:4b, qwen3:8b, qwen3:14b, qwen3:30b, qwen3:32b, qwen3:235b
    - deepseek-r1:1.5b, deepseek-r1:7b, deepseek-r1:8b, deepseek-r1:14b, deepseek-r1:32b, deepseek-r1:70b, deepseek-r1:671b
    - llama4
    - llama3.3

- **Mistral**

  - Requires an API key from [Mistral AI Platform](https://console.mistral.ai/)
  - Supported models:
    - mistral-large-latest
    - mistral-medium-latest
    - mistral-small-latest

- **Google AI**
  - Requires a Google AI API key from [Google AI Studio](https://makersuite.google.com/app/apikey)
  - Supported models:
    - gemini-2.0-flash
    - gemini-2.5-pro-exp-03-25
    - gemini-2.5-pro-preview-05-06
    - gemini-2.5-flash-preview-04-17

- **Groq**
  - Requires a Groq API key from [Groq](https://console.groq.com/keys)
  - Supported models:
    - gemma2-9b-it
    - llama-3.3-70b-versatile
    - llama-3.1-8b-instant
    - llama3-70b-8192
    - llama3-8b-8192

You can change provider settings anytime by running `autocommit configure` again.

## Usage

1. Stage your changes:
   `git add .`
2. Run AutoCommit:
   `autocommit`
3. Follow the prompts to:
   - Commit with the generated message
   - Copy the message to clipboard
   - Regenerate with new instructions
   - Add a custom instruction
   - Exit without committing

## Instructions Template

- Open and edit the AI prompt template:
  `autocommit instructions`
- Restore default template:
  `autocommit instructions restore`

If the template is missing, AutoCommit recreates it automatically.

## Quality of Commit Messages

Message quality depends on the model and context size. Large diffs dilute focus, and older or smaller models often miss instructions or skip key details. Reasoning models tend to follow prompts more faithfully and produce clearer, more accurate messages, but tend to cost more. For best results, keep your commits focused, minimize diff size, and use a reasoning-capable model.

## License

MIT License
See [LICENSE](LICENSE) for details.
