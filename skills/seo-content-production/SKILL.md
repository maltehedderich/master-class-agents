---
name: seo-content-production
description: "Create SEO content from an opportunity analysis by orchestrating the SEO Brief Architect, SEO Content Drafter, and SEO Content Editor agents. Use when turning a saved or pasted SEO opportunity analysis and optional topic into a brief, draft, editorial review loop, and final Markdown artifact under docs/seo/."
argument-hint: "SEO opportunity analysis or path, optional topic or target query, product/site context, source material, internal links, author or SME details, and publication constraints"
---

# SEO Content Production

## What This Skill Does

This skill turns a completed SEO opportunity analysis into a publishable SEO content draft through a three-agent workflow:

1. `seo-brief-architect` creates the content brief.
2. `seo-content-drafter` writes the draft from the brief.
3. `seo-content-editor` reviews the draft, either fixes minor issues directly or sends actionable feedback back to the drafter.

Briefs are stored in `docs/seo/briefs/`. Final drafts are stored in `docs/seo/drafts/`.

Use this skill for content production after SEO opportunity analysis. Do not use it to create the opportunity analysis itself.

## Required Inputs

Minimum viable inputs:

- `seo_opportunity_analysis` - a pasted report or path to a completed SEO opportunity analysis.
- `topic` - the topic, working title, or target query to produce.

If `topic` is missing, identify the most relevant next topic from the opportunity analysis. Prefer the report's first-piece recommendation. If that is absent, choose the highest-priority opportunity with the strongest overlap of business value, winnability, and search intent clarity, and state the choice in the brief handoff.

If `seo_opportunity_analysis` is missing, stop. Instruct the user to create it with the `seo-opportunity-analysis` skill and ask for the inputs that skill needs:

- target customer
- problem or job to be solved
- solution mechanics
- value proposition
- GTM channels or buying context
- optional founder/operator experience, competitors, domain authority, and keyword data source

Do not continue this skill without an opportunity analysis.

Helpful supporting inputs:

- product or site context not already in the opportunity report
- brand voice, style guide, and formatting requirements
- author or SME credentials
- source pack, proof assets, examples, screenshots, or original data
- required internal links and anchor text
- pages or claims to avoid
- desired publication format or CMS constraints

## Non-Negotiable Constraints

1. Use the installed custom agent roles by name: `seo-brief-architect`, `seo-content-drafter`, and `seo-content-editor`. Do not rely on local filesystem paths for agent definitions.
2. The brief must be completed before the draft starts.
3. The draft must be reviewed by the editor before a final artifact is saved.
4. Run at most two drafter/editor revision cycles after the first editor review.
5. The editor is the final agent in the workflow. If issues are minor, the editor fixes them directly and returns the publishable final draft.
6. Save only durable artifacts to the workspace: briefs in `docs/seo/briefs/`, final drafts in `docs/seo/drafts/`. Do not save intermediate rejected drafts unless the user asks.
7. Do not invent sources, internal links, author credentials, data, screenshots, or firsthand evidence. Use placeholders only when the missing item is unavoidable, and make those placeholders explicit in the final QA note.
8. If live SERP validation is needed, use current web results rather than stale memory.

## Procedure

### Step 1 - Validate Inputs

Confirm that an opportunity analysis is present as pasted text or a readable path.

If the analysis is a path, read it. If the path is unreadable, ask the user to provide the correct path or paste the analysis.

If no topic is supplied:

- scan the opportunity analysis for `First-piece recommendation`, ranked opportunities, cluster candidates, or target queries
- choose the most relevant next topic
- record the inferred topic and reason in the brief handoff

Create the output directories if needed:

- `docs/seo/briefs/`
- `docs/seo/drafts/`

### Step 2 - Spawn the SEO Brief Architect

Spawn a new `seo-brief-architect` subagent.

Pass it:

- the full opportunity analysis
- the selected topic or target query
- any supporting inputs from the user
- the required brief output path: `docs/seo/briefs/<slug>.md`
- the requirement to write the brief file directly

Ask for a production-ready brief that includes:

- intent statement and brief thesis
- SERP snapshot or supplied SERP interpretation
- 10x angle
- title tag, H1, meta description, URL slug, schema, and featured-snippet guidance
- complete H2/H3 hierarchy
- must-cover subpoints and required assets
- E-E-A-T, source, citation, author, and internal-link requirements
- acceptance checklist, assumptions, risks, and missing inputs

After the agent returns, verify that the brief file exists. If the agent could not write files, write its returned brief to the required path yourself.

### Step 3 - Spawn the SEO Content Drafter

Spawn a new `seo-content-drafter` subagent after the brief is complete.

Pass it:

- the brief path and brief content
- the opportunity analysis
- the selected topic
- any source pack, internal links, author details, and publication constraints
- the instruction to return a complete production-ready Markdown draft, not to save the final artifact yet

Ask the drafter to:

- follow the brief exactly
- include metadata, slug, author bio, body copy, citations, and internal links required by the brief
- end with a short QA note covering brief compliance, keyword use, links, citations, assumptions, and unresolved issues

### Step 4 - Spawn the SEO Content Editor

Spawn a new `seo-content-editor` subagent after the drafter returns.

Pass it:

- the opportunity analysis
- the brief
- the drafter's current draft
- the selected topic
- the remaining revision budget
- the final draft path: `docs/seo/drafts/<slug>.md`

Ask for one of three verdicts:

- `approve` - draft is publication-ready; write the final Markdown file.
- `approve with line edits` - apply minor fixes directly; write the final Markdown file.
- `send back to drafter` - return precise blocking feedback for structural, intent, proof, citation, or brief-compliance issues.

The editor must write the final artifact when the verdict is `approve` or `approve with line edits`.

### Step 5 - Revision Loop

If the editor sends the piece back to the drafter:

1. Pass the editor's blocking feedback, the original brief, and the current draft back to the same drafter subagent.
2. Ask for a revised full draft plus a concise change log.
3. Pass the revised draft back to the same editor subagent for another verdict.

Run this loop up to two times.

Stop conditions:

- If the editor approves or approves with line edits, use the editor's final artifact.
- If the editor still has blocking concerns after two revision cycles, ask the editor to make the best possible final publication-gate version, write it to `docs/seo/drafts/<slug>.md`, and include unresolved risks in the QA note.

The editor is always the last agent to touch the final artifact.

## File Naming

Use a short lowercase slug from the selected topic, target query, or brief URL slug.

Default paths:

- Brief: `docs/seo/briefs/<slug>.md`
- Draft: `docs/seo/drafts/<slug>.md`

If a target file already exists, do not overwrite unrelated work. Append `-2`, `-3`, and so on.

## Output Format

After the workflow completes, reply with:

1. Brief path.
2. Final draft path.
3. Selected topic or inferred topic.
4. Editor verdict.
5. Any unresolved assumptions, missing source items, or publication risks.

Do not paste the full brief or draft into chat unless the user asks.

## Quality Checks

The work is ready only when:

- The topic is traceable to the opportunity analysis.
- The brief defines one primary intent and a concrete 10x angle.
- The draft follows the brief's heading hierarchy and on-page requirements.
- Required metadata, citations, internal links, and E-E-A-T signals are present or explicitly flagged as missing.
- The editor has reviewed the latest draft and made the final publication decision.
- The brief and final draft exist at the expected paths.

## Failure Modes To Avoid

- Starting with a draft before the brief exists.
- Treating the opportunity analysis as optional.
- Creating a new opportunity analysis inside this skill instead of routing the user to `seo-opportunity-analysis`.
- Letting the drafter decide the strategy after the brief architect has fixed it.
- Saving the drafter's unreviewed draft as final.
- Running endless revision loops.
- Approving polished prose that misses intent, proof, source, or internal-link requirements.
