# Masterclass Guide to Prompt Engineering

## Role Framing

A prompt engineer designs the inputs, context, and constraints that make large language models produce reliable, high-quality outputs at scale. Excellence requires equal parts empirical experimentation, clear writing, and systems thinking — it's closer to scientific debugging than to "magic words."

## Best-in-Class Voices

- **Anthropic's Applied AI & research teams** (e.g., publications by Amanda Askell, Alex Albert) — authors of Claude's prompting documentation; emphasize clarity, structure, examples, and treating the model like a smart new hire who needs context.
- **OpenAI's prompting and cookbook team** — maintainers of the OpenAI prompting guide and cookbook; emphasize specificity, delimiters, decomposition, and reference texts.
- **Riley Goodside** (Scale AI) — one of the earliest practitioners to publicly document prompt techniques and failure modes; emphasizes empirical testing and adversarial edge cases.
- **Simon Willison** — independent researcher and prolific blogger on LLM behavior, prompt injection, and tool use; emphasizes treating prompts as code and being skeptical of "tricks."
- **Jason Wei** — researcher behind chain-of-thought prompting; emphasizes that reasoning improvements often come from letting the model "think" before answering.
- **Ethan Mollick** (Wharton) — popularizer of practical prompting for knowledge work; emphasizes iteration, persona/context-setting, and treating the model as a collaborator.

## Core Idea

A great prompt is not a clever incantation — it's a clear, well-scoped _specification_ of the task, the context, and the desired output, refined through empirical iteration against real examples.

## Essential Best Practices

1. **Be specific about the task, audience, and output format.**
   _Why:_ Models default to generic, average-of-the-internet responses when under-specified. Most "bad outputs" trace back to under-specified prompts.
   _How:_ State who the output is for, what success looks like, the format (JSON schema, markdown, length), and any non-obvious constraints. Ask yourself: "Could a smart new hire do this task from my instructions alone?"

2. **Give the model context and reference material instead of relying on its memory.**
   _Why:_ Models hallucinate when forced to recall; they reason much more reliably when given source text to ground in.
   _How:_ Paste in the relevant document, data, or examples directly. For factual tasks, instruct it to answer _only_ from the provided context and to say "I don't know" otherwise.

3. **Use few-shot examples for anything with a specific style or structure.**
   _Why:_ One or two well-chosen examples convey tacit requirements that paragraphs of instructions can't. The model pattern-matches strongly to demonstrations.
   _How:_ Show 2–5 input/output pairs that mirror real cases, including edge cases. Make sure your examples are _correct_ — the model will faithfully reproduce mistakes in them.

4. **Let the model think before it answers (when reasoning matters).**
   _Why:_ Chain-of-thought and similar techniques meaningfully improve accuracy on multi-step problems by giving the model "scratch space."
   _How:_ Ask it to work through the problem step by step, or to draft and then critique. For structured outputs, have it reason in a `<thinking>` section first, then produce the final answer separately. With reasoning models, lean on clear instructions rather than forcing explicit CoT.

5. **Decompose complex tasks into smaller steps or chained calls.**
   _Why:_ One mega-prompt asking for ten things at once is brittle; a pipeline of focused prompts is debuggable and more accurate.
   _How:_ Split into stages (extract → classify → summarize → format). Each stage has one job, clear inputs, and clear outputs. This also makes evaluation tractable.

6. **Use structure and delimiters to separate instructions, context, and data.**
   _Why:_ Models parse structured prompts more reliably and are less likely to confuse user data with instructions (which also reduces prompt injection risk).
   _How:_ Use XML tags, markdown headers, or triple backticks to wrap distinct sections (`<instructions>`, `<context>`, `<examples>`, `<input>`). Be consistent.

7. **Treat prompts as code: version, test, and evaluate them.**
   _Why:_ Prompts that "feel good" on three examples often fail on the fourth. Without evals, you're guessing.
   _How:_ Build a small eval set of 20–100 representative inputs with expected outputs or quality criteria. Re-run on every prompt change. Track regressions. For production, automate it.

8. **Iterate empirically — don't trust folklore.**
   _Why:_ Prompting "tips" go stale quickly as models change, and many viral tricks (threats, tips, "take a deep breath") have inconsistent effects across models and tasks.
   _How:_ Run A/B comparisons on your own data. Trust your evals over Twitter advice. Re-test prompts when you upgrade models.

## Common Mistakes

- **Vague instructions plus high expectations** — asking for "a good summary" without defining length, audience, or what to emphasize.
- **Overloading one prompt** with many tasks instead of chaining or splitting them.
- **No evaluation loop** — judging prompt quality from a handful of cherry-picked examples and shipping.

## Quick Start

- Write your next prompt as if briefing a competent stranger: role, task, context, format, constraints, and one example.
- Build a 20-example eval set for your most important prompt this week and re-run it on every change.
- Wrap inputs and instructions in clear delimiters (XML tags or markdown sections).
- For any multi-step task, add "Think step by step before giving your final answer" — or split it into two prompts.
- Read the official prompting guide for the model you actually use (Anthropic's or OpenAI's docs) — they're short, current, and model-specific.
