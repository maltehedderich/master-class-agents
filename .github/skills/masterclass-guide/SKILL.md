---
name: masterclass-guide
description: "Create research-first masterclass guides for roles and disciplines. Use when the user asks for a masterclass, expert guide, role guide, or best-practices synthesis such as backend engineer, product manager, or technical writer. Research best-in-class voices first, then synthesize durable principles into a concise practical guide."
argument-hint: "Role or topic to turn into a masterclass guide"
---

# Masterclass Guide

## What This Skill Does

This skill turns a role, discipline, or narrowly scoped topic into a concise masterclass guide grounded in research rather than generic advice.

## When to Use

- The user asks for a masterclass, expert guide, playbook, or role-specific best practices.
- The user wants a concise synthesis of how top practitioners think and operate.
- The user wants durable principles, judgment, and execution guidance instead of trend-driven tips.

## Inputs

- A role, discipline, or tightly scoped topic.
- Optional audience, industry, context, or named experts to include or exclude.

## Procedure

1. Frame the role or topic.
   - Determine the most likely interpretation.
   - If the request is broad or ambiguous, state the assumption briefly and proceed.
   - Ask a clarifying question only if different interpretations would materially change the guidance.
2. Research best-in-class voices.
   - Identify 3-6 relevant voices.
   - Prefer a mix of elite practitioners, respected teachers, and clear thinkers with demonstrated excellence.
   - Favor published guidance such as books, essays, talks, interviews, or frameworks.
   - Do not invent authorities, quotes, or source claims.
3. Extract durable overlap.
   - Pull the most repeated, high-signal principles across the sources.
   - Prioritize fundamentals, judgment, tradeoffs, and execution.
   - Ignore fashionable tactics unless they are clearly durable.
4. Synthesize the guide.
   - Follow [the masterclass template](./assets/master-class-template.md).
   - Keep the result concise, role-specific, and practical.
   - Lead with the highest-leverage principles.
5. Run a quality check.
   - Ensure the guidance is tailored to the role or topic.
   - Keep research first and synthesis second.
   - Avoid generic recycled advice.
   - If evidence is thin or uncertain, say so briefly instead of overstating.
   - Do not provide harmful, illegal, deceptive, or unethical guidance.

## Output Standard

- Produce the final answer in the exact section order from [the masterclass template](./assets/master-class-template.md).
- Keep every section concise.
- Make the “Best-in-Class Voices” and “Essential Best Practices” sections do the real work.
- Focus on timeless fundamentals and applied judgment.

## Optional File Workflow

- If the user wants the guide saved and the workspace contains a `guides/` directory, write the result to `guides/<topic>.md` using kebab-case.
- If an existing prompt or guide in the repo already defines the voice, stay close to that structure unless the user asks to change it.
