package chat

// MinimalInstructions is a template that generates instructions being as minimal
// as possible to save the usage of tokens. It will be used by models that support
// less tokens, such as GPT-3.5-turbo or GPT-4.
const MinimalInstructions = `Specifications:
- The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in RFC 2119. This message should adhere to the Conventional Commits specification. It should include a type (such as "feat", "fix", "chore", etc.), an optional scope (which could be a specific part of the codebase), and a succinct description of the change. The description should provide information on what the commit does, not what it's doing. It should add any additional information, such as body or footer, if necessary and start the description with lowercase.
- A description must ALWAYS start with lowercase.

The commit message should be structured as follows:
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]

Allowed types:
feat, fix, chore, ci, build, docs, test, perf, style, refactor`

// DetailedInstructions is a template that contains a more detailed explanation of the
// Conventional Commits specification. It will be used by models that support more tokens,
// such as GPT-3.5-turbo-16k or GPT-4-32k.
const DetailedInstructions = `
Description:
The commit message should be structured as follows:
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]

The commit contains the following structural elements, to communicate intent to the consumers of your library:
- fix: This type is used for commits that fix a bug in the code. Similar to feat, this kind of commit is likely to trigger a patch version bump in semantic versioning.
- chore: This type of commit is used for changes that don't modify the src (source) or test files. Examples include routine tasks like updating the build system or package manager configurations. Generally, it does not affect the functionality of the application.
- ci: Stands for Continuous Integration. This is used for commits that are related to setting up, modifying, or fixing CI build systems and services, like Travis, Jenkins, or GitHub Actions.
- build: This type is used for commits that affect the build system or external dependencies. Examples include changes in Makefile, npm, dependencies, and configurations related to the building process.
- docs: Used for commits that update documentation, like README files or comments in the source code. This does not affect the functionality of the application.
- test: This type of commit is used when adding or modifying test or fixing issues in existing tests. This ensures that the code is working as expected without changing any functionality. Should be used over 'feat' if changes only affect tests.
- perf: Short for performance. It is used for commits that improve the runtime performance of the code. This kind of commit indicates optimization in the codebase which makes it run faster or use fewer resources.
- style: This type is used for formatting changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc.). It involves cosmetic changes that do not alter the logic of the code.
- refactor: This type is used when code is being restructured or cleaned up, and this neither fixes a bug nor adds a feature. It is primarily used for reorganizing the code or cleaning up code redundancies.
- feat: This type is short for "features". It is used when a new functionallity is added to the application. This kind of commit will usually trigger a minor version bump (if you follow semantic versioning). Creating or adding tests aren't considered as a new feature.
- BREAKING CHANGE: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
- footers other than BREAKING CHANGE: <description> may be provided and follow a convention similar to git trailer format.
- If there are no BREAKING CHANGES, footers should be omitted.

Examples:
Commit message with description and breaking change footer:
feat: allow provided config object to extend other configs

BREAKING CHANGE: 'extends' key in config file is now used for extending other config files

Commit message with ! to draw attention to breaking change:
feat!: send an email to the customer when a product is shipped

Commit message with scope and ! to draw attention to breaking change:
feat(api)!: send an email to the customer when a product is shipped

Commit message with both ! and BREAKING CHANGE footer:
chore!: drop support for Node 6

Commit message with description when adding a new test:
test: add unit tests for the new parser

Commit message with no body:
docs: correct spelling of CHANGELOG

Commit message with scope:
feat(lang): add Polish language

Commit message with multi-paragraph body and multiple footers:
fix: prevent racing of requests

Introduce a request id and a reference to latest request. Dismiss
incoming responses other than from latest request.

Remove timeouts which were used to mitigate the racing issue but are
obsolete now.

Commit message for a fix in the test suite:
test: make failing test pass on Edge 13

Specifications:
- The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in RFC 2119.
- A description MUST always start with lowercase.
- Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc., followed by the OPTIONAL scope, OPTIONAL !, and REQUIRED terminal colon and space.
- The type feat MUST be used when a commit adds a new feature to your application or library.
- The type fix MUST be used when a commit represents a bug fix for your application.
- A scope MAY be provided after a type. A scope MUST consist of a noun describing a section of the codebase surrounded by parenthesis, e.g., fix(parser):
- A description MUST immediately follow the colon and space after the type/scope prefix. The description is a short summary of the code changes, e.g., fix: array parsing issue when multiple spaces were contained in string.
- A longer commit body MAY be provided after the short description, providing additional contextual information about the code changes. The body MUST begin one blank line after the description.
- A commit body is free-form and MAY consist of any number of newline separated paragraphs.
- One or more footers MAY be provided one blank line after the body. Each footer MUST consist of a word token, followed by either a :<space> or <space># separator, followed by a string value (this is inspired by the git trailer convention).
- A footer’s token MUST use - in place of whitespace characters, e.g., Acked-by (this helps differentiate the footer section from a multi-paragraph body). An exception is made for BREAKING CHANGE, which MAY also be used as a token.
- A footer’s value MAY contain spaces and newlines, and parsing MUST terminate when the next valid footer token/separator pair is observed.
- Breaking changes MUST be indicated in the type/scope prefix of a commit, or as an entry in the footer.
- If included as a footer, a breaking change MUST consist of the uppercase text BREAKING CHANGE, followed by a colon, space, and description, e.g., BREAKING CHANGE: environment variables now take precedence over config files.
- If included in the type/scope prefix, breaking changes MUST be indicated by a ! immediately before the :. If ! is used, BREAKING CHANGE: MAY be omitted from the footer section, and the commit description SHALL be used to describe the breaking change.
- Types other than feat and fix MAY be used in your commit messages, e.g., docs: update ref docs.
- The units of information that make up Conventional Commits MUST NOT be treated as case sensitive by implementors, with the exception of BREAKING CHANGE which MUST be uppercase.
- BREAKING-CHANGE MUST be synonymous with BREAKING CHANGE, when used as a token in a footer.`

// ShortMode is a template that generates a shorter commit message.
const ShortMode = `Objective: 
Generate a single commit messages based on the output of a 'git diff --cached' command in the style of Conventional Commits based on Angular commit guidelines.
The output should be a single very short, concise and readable commit message. Avoid writing descriptions that are longer than 50 characters.
Additonally, try to avoid writing body and footer as much as possible - if really necessary, be very brief.`

// DetailedMode is a template that generates a more detailed commit message.
const DetailedMode = `Objective: 
Generate a single commit messages based on the output of a 'git diff --cached' command in the style of Conventional Commits based on Angular commit guidelines.
The output should be a single readable and meaningful commit message. You must never write descriptions that are longer than 75 characters.`
