---
name: ui-writing
description: "Draft, edit, review, or systematize product interface copy that helps people understand what is happening and take the right next action. Use when working on end-to-end UI flows, buttons and labels, forms and validation, errors, empty and loading states, permissions, confirmations, destructive actions, success messages, onboarding, voice-and-tone guidance, localization-ready strings, accessibility, or copy experiments and audits."
---

# UI Writing

## Purpose

Design the interface with words. Reduce uncertainty at the moment a person must decide or act, using the smallest amount of copy that preserves clarity, trust, and control.

Treat UI writing as product design rather than prose decoration. Shape the whole flow, expose missing states and unclear product behavior, and use evidence and outcomes to settle copy decisions.

## Inputs

Gather only the context needed for the requested scope:

- `goal` — `draft`, `edit`, `review`, `audit`, `voice-chart`, or `experiment`.
- `user_and_job` — who is acting, what they are trying to accomplish, and what they already know.
- `flow` — screens, sequence, entry points, exits, and all known states.
- `decision_or_action` — what each string must help the person understand, choose, or do.
- `system_behavior` — what actually happens, including timing, reversibility, side effects, and failure conditions.
- `evidence` — support tickets, search terms, research notes, usability findings, analytics, or established product vocabulary.
- `voice_constraints` — stable brand traits, preferred and prohibited terms, regulated language, and existing style guidance.
- `delivery_constraints` — component type, character or layout limits, platforms, accessibility requirements, and target languages.
- `success_measure` — task completion, recovery, step drop-off, time on task, or support volume.

When the request is narrow and enough context exists, proceed and state material assumptions. Ask a targeted question only when unknown product behavior, risk, or user intent would change the copy.

## Non-Negotiable Constraints

1. Write the flow, not an isolated string. Check what comes before, what comes after, and every state the user may encounter.
2. Give every string one job. Make it reduce uncertainty, name an action or outcome, prevent a mistake, support recovery, or build warranted trust. Delete copy without a job.
3. Use one user-facing term for each concept. Do not let planning language, database names, or implementation details leak into the interface.
4. Front-load meaning. Make the first words carry the information a scanning reader needs.
5. Name outcomes, not controls. Prefer `Save changes`, `Send invite`, or `Delete 3 files` over `Submit`, `Continue`, or `OK` when the outcome can be named.
6. Cover unhappy and edge states. Include validation, errors, permissions, waiting, emptiness, destructive actions, recovery, and completion where the flow can produce them.
7. Keep voice stable and adapt tone to emotional risk. Use restrained, serious language around money, health, privacy, security, and data loss.
8. Write honestly. Never use confirmshaming, invented urgency, disguised ads, hidden exits, ambiguous destructive actions, or coercive consent language.
9. Write for accessibility, translation, and narrow layouts from the first draft. Do not leave them for a cleanup pass.
10. Attach important changes to evidence or a measurable hypothesis. Do not defend copy by taste alone.

## Workflow

### 1. Frame the user decision

State:

- who the person is in this moment;
- what they are trying to accomplish;
- what they need to know before acting;
- the intended action or decision;
- the consequence, reversibility, and emotional risk;
- the observable outcome that would show the copy worked.

Distinguish a copy problem from a product problem. Flag unclear system behavior, missing recovery, deceptive defaults, or an overloaded decision instead of writing around them.

### 2. Inventory the complete flow

Put all strings in sequence before polishing any of them. Use a table such as:

| Location | State or trigger | User need or decision | Current or draft copy | Constraint |
| --- | --- | --- | --- | --- |

Include every applicable state:

- entry and default;
- instructions and help;
- input, selection, and validation;
- loading and waiting;
- empty and first-use;
- permission and privacy;
- warning and confirmation;
- error and recovery;
- success and next step;
- offline, timeout, expired session, partial completion, and other realistic edges.

Read the inventory aloud in order. Fix missing states, contradictory promises, and vocabulary drift before refining individual strings.

### 3. Establish vocabulary from evidence

Use research and operational evidence before stakeholder preference:

- Mine the words people use in support tickets, search queries, interviews, usability sessions, and session recordings.
- Identify the nouns people look for and the verbs they already associate with the task.
- Create a small glossary of `use`, `avoid`, and `meaning` when the flow introduces recurring concepts.
- Preserve legally or technically necessary distinctions, but explain them in user language.
- Mark missing evidence explicitly. Never invent quotes, findings, or comprehension claims.

### 4. Draft the information structure

Write the hierarchy before adding personality:

- Lead with the user's goal or the object they are looking for.
- Keep headings, labels, instructions, actions, and feedback distinct.
- Use persistent, specific field labels; do not rely on placeholder text as the only label.
- Put preventive guidance before the action or input where it can avert an error.
- Keep one primary action per decision point when the product allows it.
- Make secondary and exit actions clear without hiding or shaming them.

Draft buttons as `verb + object` or `verb + outcome`. Use generic actions only when the outcome is already unmistakable in context.

### 5. Write the critical states

Use these patterns as starting structures, not rigid templates:

| State | Structure | Standard |
| --- | --- | --- |
| Error | What happened. Why, only if it changes the response. What to do next. | Avoid blame and raw error codes as the only explanation. Preserve useful entered data. |
| Empty | What this area is for. How to create or find the first useful item. | Teach the primary action; do not apologize for emptiness. |
| Permission | Concrete benefit in the user's terms. What access is requested and when. | Explain before an operating-system prompt; make refusal understandable. |
| Destructive action | Exact object and consequence. Whether recovery is possible. Explicit action label. | Replace `OK` with the actual outcome and keep a clear safe exit. |
| Loading or waiting | What is happening. Expected duration or progress when known. Available control. | Do not promise a duration the system cannot meet. |
| Success | What completed. Important resulting state. Useful next action, if any. | Keep celebration proportional to the moment. |

Prefer specific recovery copy. For example, replace `Something went wrong` with `We couldn't upload invoice.pdf. Check your connection and try again.` only when those facts and that recovery step are accurate.

### 6. Apply voice and state-sensitive tone

Keep voice stable across the product. Encode it in a compact chart:

| Dimension | Use | Avoid | Before | After |
| --- | --- | --- | --- | --- |
| Concepts | Core ideas the product consistently reinforces | Claims the product cannot sustain | Real weak copy | Real improved copy |
| Vocabulary | Preferred user terms | Internal jargon and prohibited terms | Real weak copy | Real improved copy |
| Verbosity | Typical sentence and component length | Padding and repeated explanation | Real weak copy | Real improved copy |
| Grammar | Person, tense, contractions, capitalization | Inconsistent constructions | Real weak copy | Real improved copy |
| Punctuation | Deliberate punctuation rules | Decorative or excessive punctuation | Real weak copy | Real improved copy |

Then adapt tone to the state:

- Use direct, neutral language for routine tasks.
- Use calm, factual language for errors and interruptions.
- Use restrained confirmation for ordinary success.
- Use plainly serious language for money, health, privacy, security, irreversible actions, and data loss.
- Remove humor, idioms, and brand performance when the person is anxious or blocked.

### 7. Cut and front-load

Edit in this order:

1. Delete the string and see whether the flow still works.
2. Remove greetings, windups, repeated headings, restated buttons, and explanatory padding.
3. Move the key noun, outcome, or problem to the beginning.
4. Replace vague verbs and pronouns with the specific action and object.
5. Break dense paragraphs into scannable units only when each unit adds information.
6. Read the flow aloud and cut again without removing necessary context, consequence, or recovery guidance.

Prefer plain language and sentence case. Optimize for immediate comprehension, not cleverness or maximum brevity.

### 8. Check accessibility and localization

Verify the first draft against these constraints:

- Make buttons and links understandable when a screen reader reads them out of context.
- Do not use standalone `Click here` or `Learn more` when a specific destination can be named.
- Do not rely on position, color, shape, or visual proximity alone to explain an action.
- Provide complete strings to translators; never concatenate fragments into sentences.
- Add context notes for ambiguous nouns, verbs, variables, gender, plurality, and placeholders.
- Avoid idioms, puns, culture-bound humor, and strings that depend on English word order.
- Allow for text expansion and test the tightest supported width and largest relevant text size.
- Avoid long all-caps strings and unnecessary punctuation that screen readers may announce awkwardly.

Do not shorten away meaning merely to fit a component. Recommend a layout or interaction change when the space cannot support honest, usable copy.

### 9. Review the flow end to end

Walk the flow as the intended user and ask at every step:

- What does this person believe is happening now?
- Is the next action self-evident?
- Does the label accurately predict the result?
- Is the same concept named consistently?
- Can the person decline, go back, recover, or understand an irreversible consequence?
- Does feedback arrive at the right time and refer to the correct object?
- Does any string repeat information already visible or fail to answer the next likely question?

Resolve high-risk comprehension and trust problems before polishing low-risk wording.

### 10. Define measurement

Connect the change to the uncertainty it is meant to reduce:

- Use task completion and step drop-off for action clarity.
- Use validation frequency and error recovery rate for form or error copy.
- Use time on task and repeated actions for comprehension problems.
- Use support ticket volume by topic for recurring uncertainty.
- Use usability observation when analytics cannot reveal why a person hesitates.

Record the baseline, exposure, primary metric, and guardrail before comparing variants. Change one meaningful variable at a time when attribution matters. Do not optimize acceptance or completion by making refusal, cancellation, or informed choice harder.

## Branching Logic

- **Net-new flow.** Inventory states and system behavior before drafting. Return a complete string set, not only the happy path shown in a mockup.
- **Existing-flow review.** Diagnose in flow order, prioritize by user risk and frequency, and provide concrete rewrites beside each finding.
- **Single-string request.** Inspect the nearby decision and state. Return the strongest recommendation first, plus alternatives only when they represent meaningfully different product assumptions or tones.
- **Error-heavy surface.** Group failures by cause and recovery action. Standardize structure while keeping each message specific enough to act on.
- **Voice-and-tone request.** Build the voice chart from real product examples, then add a tone map for routine, success, blocked, sensitive, and high-risk states.
- **High-risk or regulated flow.** Preserve required language, make consequences explicit, avoid unsupported reassurance, and recommend legal, compliance, safety, or subject-matter review where necessary.
- **Localization request.** Deliver complete source strings with context notes, variables, character constraints, and plural or grammatical considerations. Do not provide fragment assemblies.
- **No research evidence.** Draft from clearly labeled assumptions and identify the smallest research or support-data check that would reduce uncertainty.
- **Copy cannot solve the problem.** Recommend the product, interaction, hierarchy, or layout change directly and explain which user uncertainty it resolves.

## Output Format

Match the deliverable to the task.

For `draft`:

1. State the user, job, flow boundary, and material assumptions.
2. Provide the complete string inventory in flow order, including applicable non-happy states.
3. Add brief implementation or context notes only where timing, variables, accessibility, or localization affects meaning.
4. List unresolved product questions that block accurate copy.
5. Recommend one success metric and any essential trust or accessibility guardrail.

For `edit`, `review`, or `audit`:

1. Lead with the highest-risk finding.
2. Provide a `location / current copy / recommended copy / reason` table in flow order.
3. List missing states, vocabulary inconsistencies, and product issues separately.
4. Prioritize the changes as `critical`, `important`, or `polish`.
5. Suggest measurement for the highest-leverage change.

For `voice-chart`:

1. Provide the voice dimensions with `use`, `avoid`, and real before/after examples.
2. Provide the state-based tone map.
3. Provide a short preferred/prohibited vocabulary list.
4. Demonstrate the system on a small set of representative components, including an error and a sensitive state.

For a narrow component request, return the recommended copy first. Keep rationale short and surface any system assumption that must be confirmed before shipping.

## Quality Checks

Treat the work as ready only when all applicable checks pass:

1. Identify the user, the immediate decision, and the intended outcome.
2. Review the copy in sequence rather than as disconnected strings.
3. Give every retained string a clear job.
4. Use one consistent, user-understood term per concept.
5. Make action labels predict their outcomes.
6. Cover realistic validation, waiting, empty, permission, error, destructive, recovery, and success states.
7. Explain errors without blame and provide an accurate next step.
8. Match tone to emotional and material risk without changing the underlying voice.
9. Remove redundant copy while retaining necessary consequences and recovery information.
10. Make controls understandable out of context and strings viable for translation and text expansion.
11. Avoid dark patterns, false urgency, hidden exits, and ambiguous consent or destruction.
12. Ground important choices in evidence or label them as hypotheses with a measurement plan.

## Failure Modes

- Writing after the interaction is fixed and merely squeezing words into boxes.
- Polishing the happy path while omitting errors, permissions, emptiness, waiting, and recovery.
- Describing the system instead of answering what the person needs to know to act.
- Using internal feature names, implementation details, or raw codes as user-facing language.
- Choosing vague labels such as `Submit`, `OK`, or `Continue` when the outcome is not obvious.
- Restating a heading or button in nearby helper text.
- Using wit, celebration, or brand personality during failure, payment, health, privacy, or data-loss moments.
- Blaming the person for a validation or system error.
- Hiding consequences, exits, refusal, or cancellation to improve a metric.
- Concatenating fragments, relying on placeholders as labels, or forcing essential copy into an undersized component.
- Producing many cosmetic variants without identifying the decision, evidence, or hypothesis behind them.
- Inventing user research, technical causes, timing promises, or recovery steps.
- Treating a product or interaction problem as if wording alone can repair it.

## Default Stance

Prefer obvious over clever, specific over generic, calm over performative, and evidence over opinion. Write the smallest honest string that helps the person take the right next action, then verify that it still works as part of the complete flow.
