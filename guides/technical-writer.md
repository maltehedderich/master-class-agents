# Masterclass Guide to Technical Writing for Software Applications

## Role Framing

Excellence in software technical writing means making complex systems usable through clear, accurate, task-oriented prose — serving readers who are usually impatient, mid-task, and looking for a specific answer. The best technical writers are part teacher, part editor, part UX designer, and part developer advocate.

## Best-in-Class Voices

- **Daniele Procida** — Creator of the Diátaxis framework (adopted by Django, Cloudflare, Gatsby, and others). Emphasizes that documentation has four distinct modes — tutorials, how-to guides, reference, and explanation — and conflating them is the root cause of bad docs.
- **Tom Johnson (_I'd Rather Be Writing_)** — Long-running practitioner-blogger on API documentation. Emphasizes docs-as-code workflows, developer empathy, and treating documentation as a product.
- **Google Developer Documentation Style Guide team** — Sets a widely copied house standard. Emphasizes plain language, second person, active voice, and accessibility.
- **Write the Docs community** (Eric Holscher, Mikey Harper, et al.) — A practitioner community around Read the Docs and Sphinx. Emphasizes docs-as-code, versioning, and writing for skimmers.
- **Mike Pope / Microsoft Writing Style Guide contributors** — Long-standing influence on enterprise software docs. Emphasize consistency, terminology discipline, and global English.
- **Kathy Sierra (_Badass: Making Users Awesome_)** — Not strictly a tech writer, but hugely influential on the philosophy. Emphasizes that the goal of docs is not to explain the product but to make the user competent and confident.

## Core Idea

Good documentation is not about your software — it is about the reader's task. Every sentence should either help a specific reader accomplish a specific goal or get out of the way.

## Essential Best Practices

1. **Separate the four documentation modes (Diátaxis).**
   _Why it matters:_ Tutorials, how-to guides, reference, and explanation serve different needs and reading patterns. Mixing them produces docs that are simultaneously too long for lookups and too shallow for learning.
   _How to apply:_ Before writing, decide which mode you're in. Tutorials teach a beginner by doing; how-tos solve a specific problem for someone who already knows the basics; reference is dry, complete, and structured; explanation provides context and rationale. Keep them in separate pages or clearly separated sections.

2. **Write for the reader's task, not the system's architecture.**
   _Why it matters:_ Engineers naturally document the way the code is organized. Readers arrive with a goal ("How do I authenticate a request?"), not a mental model of your module hierarchy.
   _How to apply:_ Title pages and sections with the user's verb-led goal ("Authenticate a request"), not the component name ("AuthProvider"). Lead with what the reader wants to accomplish, then show how.

3. **Front-load the answer.**
   _Why it matters:_ Developers scan; they don't read. If they have to wade through three paragraphs of context to find a code snippet, you've lost them.
   _How to apply:_ Put the working example, the key command, or the direct answer in the first screen. Save background, caveats, and "why" for after the payoff. Use the "inverted pyramid" from journalism.

4. **Make every code example runnable, minimal, and copy-pasteable.**
   _Why it matters:_ Broken or incomplete examples destroy trust faster than anything else in docs. Readers will copy your snippet verbatim — assume it.
   _How to apply:_ Test every example in CI if possible. Include imports and setup. Avoid placeholder pseudocode unless clearly marked. Show expected output. Use real values, not `foo`/`bar`, when it aids understanding.

5. **Use plain, direct language: second person, active voice, present tense.**
   _Why it matters:_ It is shorter, less ambiguous, easier to translate, and more accessible for non-native English readers (likely a majority of your audience).
   _How to apply:_ "Run the command" beats "The command should be run." "You can configure" beats "Users may configure." Cut hedges ("simply," "just," "easily") — they patronize and add nothing.

6. **Treat docs as code, and version them with the software.**
   _Why it matters:_ Docs that drift from the code are worse than no docs. Stale instructions waste time and erode credibility.
   _How to apply:_ Store docs in the same repo as the code when possible. Require doc updates in PRs that change behavior. Use linters (Vale, markdownlint). Build docs in CI. Tag versions.

7. **Structure for skimming: headings, lists, tables, and short paragraphs.**
   _Why it matters:_ Readers find information by scanning headings and code blocks, not by reading prose.
   _How to apply:_ Use descriptive headings that work as a table of contents. Keep paragraphs to 2–4 sentences. Use tables for parameter references. Bold key terms sparingly. Make the page navigable without reading it.

8. **For READMEs specifically: lead with what it is, who it's for, and how to run it in 60 seconds.**
   _Why it matters:_ A README is a landing page. Its job is to help a stranger decide whether to invest more time, and to get the curious to "hello world" fast.
   _How to apply:_ A strong README opens with a one-sentence description, then a short "what problem this solves," then install + minimal usage example, then links to deeper docs. Include badges, license, and contribution pointers near the bottom — not the top.

## Common Mistakes

- Documenting the code's structure (modules, classes) instead of the user's goals (tasks, outcomes).
- Writing one sprawling page that tries to be tutorial, reference, and explanation at once — serving none of them well.
- Letting examples rot: untested snippets, outdated flags, or screenshots from three versions ago.

## Quick Start

- Open your current doc and label each section as tutorial, how-to, reference, or explanation. Split anything that's wearing two hats.
- Rewrite your top page's title and first paragraph to lead with the reader's goal and a working example.
- Run a "60-second test" on your README: can a stranger get from landing on it to running the software in under a minute?
- Add a docs lint step (Vale or similar) and a CI check that builds the docs on every PR.
- Pick your three most-visited pages and cut every instance of "simply," "just," "easy," and passive voice.
