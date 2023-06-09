package chat

const ShortTemplate = `
Compose a Git commit message strictly conforming the instructions below. Abide by the following guidelines:

- MUST follow the structure <type>[optional scope]: <description>
- The commit message SHALL only consist of a description with the type/scope prefix. It MUST NOT include a body or a footer.
- This document interprets key terms "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" as per RFC 2119.
- Commit messages MUST start with a type, such as 'feat', 'fix', etc. Following the type, an OPTIONAL scope, an OPTIONAL '!', and a REQUIRED colon and space MAY be added.
- The 'feat' type MUST be used when a commit introduces a new feature to your application or library.
- The 'fix' type MUST be used when a commit fixes a bug in your application.
- A scope MAY follow a type. A scope SHOULD be a noun referring to a part of the codebase enclosed in parentheses, e.g., 'fix(parser):'
- A description MUST follow immediately after the colon and space succeeding the type/scope prefix. The description SHOULD be a succinct summary of the code changes, such as 'fix: resolved array parsing issue with multiple spaces in a string.'
- The commit message MUST have only one line, with the type/scope prefix and description being separated by a colon and space.
- Types other than fix and feat are allowed, for example: build, chore, ci, docs, style, refactor, perf, test.
- You MUST NEVER write a body or footer, must have only one the structure <type>[optional scope]: <description> and no more.
- You MUST NEVER write text that mention authors, co-authors or any other people information in the commit message.

Example of commit messages:
feat(api)!: send an email to the customer when a product is shipped
fix: prevent buffer overflow when reading from stdin
feat(lang): add Polish language
docs: correct spelling of CHANGELOG
chore!: drop support for Node 6
refactor: rename getFoo() to getBar()
refactor(lang): replace if-else with switch
style: remove empty line at end of file

Upon execution of 'git diff --cached', form the commit message as per the rules outlined. Under all circumstances, follow the structure <type>[optional scope]: <description>.
`

const VerboseTemplate = `
Formulate a Git commit message in strict adherence to the Git Conventional Commits specification. Also, adhere to the following set of rules:

- MUST follow the structure <type>[optional scope]: <description>
- The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in RFC 2119.
- Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc., followed by the OPTIONAL scope, OPTIONAL !, and REQUIRED terminal colon and space.
- The type feat MUST be used when a commit adds a new feature to your application or library.
- The type fix MUST be used when a commit represents a bug fix for your application.
- A scope MAY be provided after a type. A scope MUST consist of a noun describing a section of the codebase surrounded by parenthesis, e.g., fix(parser):
- A description MUST immediately follow the colon and space after the type/scope prefix. The description is a short summary of the code changes, e.g., fix: array parsing issue when multiple spaces were contained in string.
- A longer commit body MAY be provided after the short description, providing additional contextual information about the code changes. The body MUST begin one blank line after the description.
- A commit body is free-form and MAY consist of any number of newline separated paragraphs.
- One or more footers MAY be provided one blank line after the body. Each footer MUST consist of a word token, followed by either a :<space> or <space># separator, followed by a string value (this is inspired by the git trailer convention).
- Breaking changes MUST be indicated in the type/scope prefix of a commit, or as an entry in the footer.
- If included as a footer, a breaking change MUST consist of the uppercase text BREAKING CHANGE, followed by a colon, space, and description, e.g., BREAKING CHANGE: environment variables now take precedence over config files.
- If included in the type/scope prefix, breaking changes MUST be indicated by a ! immediately before the :. If ! is used, BREAKING CHANGE: MAY be omitted from the footer section, and the commit description SHALL be used to describe the breaking change.
- Types other than feat and fix MAY be used in your commit messages, e.g., docs: update ref docs.
- BREAKING CHANGE which MUST be uppercase, when used.
- BREAKING-CHANGE MUST be synonymous with BREAKING CHANGE, when used as a token in a footer.
- If there are no breaking changes, the footer BREAKING CHANGE: must be omitted.
- Each line of the commit message MUST be at most 50 characters long.
- MUST be concise and MUST be brief. If the diff is not that long, there is not need to write a complex explanation.
- Types other than fix: and feat: are allowed, for example: build:, chore:, ci:, docs:, style:, refactor:, perf:, test:, and others.
- You MUST NEVER write a body or footer, must have only one the structure <type>[optional scope]: <description> and no more.
- You MUST NEVER write text that mention authors, co-authors or any other people information in the commit message.

Exemplary commit messages:
feat(api)!: send an email to the customer when a product is shipped
fix: prevent buffer overflow when reading from stdin
feat(lang): add Polish language
docs: correct spelling of CHANGELOG
chore!: drop support for Node 6
refactor: rename getFoo() to getBar()
refactor(lang): replace if-else with switch
style: remove empty line at end of file
'fix: prevent request races

Introduce request id and reference to the most recent request. Dismiss
responses that do not originate from the most recent request.

Eliminate obsolete timeouts that were originally introduced to counter request races.'

Upon the user prompting the output of 'git diff --cached', draft the requested commit message accordingly.
`
