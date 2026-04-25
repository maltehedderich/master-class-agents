---
name: seo-opportunity-analysis
description: "Analyze a product spec to find SEO content opportunities. Use when interpreting a product or company spec, deriving audience search behavior, mapping topical territory, proposing pillar and cluster strategy, prioritizing topics by demand x winnability x business value, choosing the first SEO piece to publish, and analyzing the live SERP for the target query with the SEO Opportunity Analyst workflow."
argument-hint: "Product spec or repo path, optional keyword data source, and any known audience or GTM constraints"
---

# SEO Opportunity Analysis

## What This Skill Does

This skill turns a product spec into an implementable SEO opportunity report.

It uses the `SEO Opportunity Analyst` workflow as the core reasoning model: start from the audience job, derive search behavior from the spec, map a compact topical territory, prioritize by demand x winnability x business value, and validate the chosen opening move against the live SERP.

Use this skill for opportunity analysis and territory design. Do not use it to draft briefs, outlines, or articles.

## When to Use

- The user has a product spec and wants SEO content opportunities rather than a content draft.
- The task is to infer audience search behavior from target customer, problem, solution mechanics, value proposition, and GTM context.
- The user needs a pillar-and-cluster strategy with explicit target queries and intent labels.
- The user wants a defensible first-piece recommendation and a live SERP read before writing anything.
- The product is early-stage or trust-poor and needs an E-E-A-T angle rooted in founder experience or first-person knowledge.

Do not use this skill for:

- Content brief construction, article outlining, or drafting.
- Broader GTM, positioning, or channel planning beyond what the supplied spec already defines.
- Paid-media strategy.

## Inputs To Gather

Collect the minimum viable spec before analysis starts.

- `target_customer` — who the product serves.
- `problem` — the job, pain, or constraint the customer is trying to solve.
- `solution_mechanics` — how the product solves it in practice.
- `value_proposition` — why it is better, different, or more credible.
- `gtm_channels` — existing acquisition context that sharpens audience language and buying context.

Helpful but optional:

- Founder or operator experience that can support a first-person E-E-A-T angle.
- Existing domain authority, content footprint, and known competitors.
- A keyword data source for volume and difficulty. If none is available, proceed qualitatively and say so explicitly.

## Non-Negotiable Constraints

1. Start from the audience job, not a keyword dump.
2. Treat the live SERP as the brief for intent, format, depth, and publisher expectations.
3. Keep the territory compact and defensible. Favor 1-3 pillars the founder can realistically cover.
4. Default to a cluster-first first move with a high-intent query unless the spec clearly supports a pillar-first authority play.
5. Do not recommend a topic if the site cannot credibly satisfy the real search intent.
6. Do not drift into content production deliverables.
7. Do not invent keyword metrics. If no data source exists, use qualitative prioritization from SERP evidence.

## Procedure

### Step 1 — Confirm the job and check for fatal spec gaps

Read the spec and extract:

- target customer
- core problem
- solution mechanics
- value proposition
- buying or evaluation context

If the target customer, problem, or solution mechanics are too thin to infer search behavior, stop and ask for the missing fields explicitly before continuing.

### Step 2 — Decide whether keyword metrics are available

Ask once whether the user has a preferred source for keyword volume or difficulty if that information is not already provided.

- If the user has a source, use it as supporting evidence.
- If the user does not have a source, proceed qualitatively from SERP analysis and state that the prioritization is directional rather than metric-backed.

Do not block the workflow on paid keyword tools.

### Step 3 — Use the SEO Opportunity Analyst workflow

Prefer invoking the `SEO Opportunity Analyst` custom agent for the core analysis when it is available in the environment.

Pass it:

- the product spec
- the in-scope tasks
- the out-of-scope tasks
- whether keyword metrics are available
- the required output sections

If the custom agent is not available, mirror its method directly instead of widening scope: interpret the spec, derive search behavior, map territory, prioritize opportunities, choose the first move, and validate it against the SERP.

### Step 4 — Derive audience search behavior

Translate the audience job into likely search behavior across the funnel.

Produce 5-10 plausible queries split across:

- problem-aware
- solution-aware
- brand-aware or comparison-oriented

Prefer natural search phrasing rooted in the spec over synthetic keyword-tool phrasing.

### Step 5 — Map the topical territory

Create 1-3 pillar candidates. For each pillar, define 5-10 cluster post candidates.

Every candidate must include:

- working title or topic
- target query
- intent classification
- relationship to the pillar

Keep the cluster map implementable. If a pillar requires unrealistic coverage depth for the founder or team, cut it.

### Step 6 — Rank the opportunities

Score the best opportunities using:

- demand
- winnability
- business value

When metrics are unavailable:

- infer demand from repeated SERP patterns, specificity of the query, and evidence of active search behavior
- infer winnability from publisher mix, content quality, topical fit, and gap availability
- infer business value from problem intensity, buying proximity, and fit with the spec

Favor topics where these three factors overlap. Reject vanity topics with traffic but weak fit.

### Step 7 — Recommend the first piece

Select the single best first piece to publish.

State whether the strategy is `cluster-first` or `pillar-first` and justify it in 2-4 sentences.

Default rule:

- choose `cluster-first` with a high-intent target query

Override only when:

- the spec implies a strong authority narrative that benefits from a foundational pillar page first
- or the pillar unlocks several immediately linked cluster pieces with better leverage than a single cluster article

### Step 8 — Run live SERP analysis for the chosen query

Use current web search results. Do not rely on stale knowledge.

For the chosen target query, produce a standard SERP block with:

- intent classification
- top-5 SERP table
- table-stakes inventory
- gap inventory
- go or no-go judgment

Use this top-5 table schema:

| Rank | URL or publisher | Format | Intent fit | Strengths | Weaknesses or gaps |

Use this table-stakes inventory schema:

- dominant format and content type
- publisher mix
- notable SERP features
- recurring subtopics
- credibility signals and E-E-A-T expectations

Use this gap inventory schema:

- underserved angle
- missing specificity
- stale examples or weak evidence
- mismatch between ranking pages and the user's likely spec-based angle

If the SERP is dominated by unmovable incumbents and no credible gap exists, do not force the recommendation. Return 1-2 adjacent territories instead.

### Step 9 — Add the E-E-A-T angle

Identify where the founder's or operator's direct experience can become the differentiator.

Examples:

- a first-person lesson from solving the problem repeatedly
- a decision framework grounded in real implementation tradeoffs
- a contrarian position supported by operating experience from the spec

For new or low-authority sites, treat this as load-bearing rather than decorative.

### Step 10 — Produce the Opportunity Report

Return the report in this order:

1. Opportunity thesis — one short paragraph.
2. Audience search-context summary — 5-10 likely queries across funnel stages.
3. Cluster strategy — 1-3 pillars, each with 5-10 cluster candidates, target query, and intent.
4. Ranked opportunities — explain demand, winnability, and business value.
5. First-piece recommendation — the single piece to write first, plus `cluster-first` or `pillar-first` rationale.
6. Standard SERP block for the first piece's target query.
7. E-E-A-T angle.
8. Assumptions, risks, and missing inputs.

## Quality Checks

The work is ready only when all of the following are true:

1. The audience job is explicit and grounded in the provided spec.
2. The query set spans multiple funnel stages rather than only informational curiosity terms.
3. Pillar and cluster relationships are explicit and implementable.
4. The founder or team could realistically commit to the proposed territory.
5. The first-piece recommendation has a clear, defensible reason tied to intent ROI or strategic seeding.
6. The chosen query has a live SERP analysis, not a guessed one.
7. The E-E-A-T angle is specific to the spec and not generic founder-posturing.

## Failure Modes To Avoid

- Starting from keyword volume without first deriving the audience job.
- Turning the output into a content brief or draft.
- Recommending head terms as the first move when a cluster entry point is more winnable.
- Treating keyword metrics as required when qualitative SERP evidence is sufficient.
- Ignoring obvious intent mismatch in the SERP.
- Proposing sprawling topic maps that the founder cannot realistically cover.
- Pushing through a no-win SERP instead of returning adjacent territories.

## Escalation Rules

- If the spec is too thin to derive search behavior, stop and ask for the exact missing fields.
- If keyword data is unavailable, proceed qualitatively and label the confidence accordingly.
- If every promising topic is blocked by dominant incumbents with no exploitable gap, return 1-2 alternative territories instead of forcing a recommendation.

## Suggested Invocation

Use this skill with:

- a product spec pasted directly into chat
- a path to a spec file in the workspace
- optional keyword data or notes about current domain authority

Typical prompt:

`Use seo-opportunity-analysis on this product spec. Prioritize a compact cluster strategy, recommend the first piece, and include a live SERP read for the target query.`
