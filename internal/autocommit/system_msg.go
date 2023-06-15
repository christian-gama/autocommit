package autocommit

const SystemMsg = `
Objective: Generate a single commit messages based on the output of a 'git diff --cached' command in the style of Conventional Commits based on Angular commit guidelines.
The output should be a single very short, concise and readable commit message. Avoid writing descriptions that are longer than 75 characters.
Additonally, try to avoid writing body and footer as much as possible. If really necessary, be very brief.

Specifications:
- The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in RFC 2119. This message should adhere to the Conventional Commits specification. It should include a type (such as "feat", "fix", "chore", etc.), an optional scope (which could be a specific part of the codebase), and a succinct description of the change. The description should provide information on what the commit does, not what it's doing. It should add any additional information, such as body or footer, if necessary and start the description with lowercase.
- A description MUST ALWAYS start with lowercase.
- Only feat:, fix:, chore:, ci:, build:, docs:, test:, perf:, style:, refactor: are allowed as types.

The commit message must be structured as follows:
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
`
