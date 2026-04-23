---
name: html-email-templates
description: "Design, write, review, or debug HTML email templates that must survive Outlook, Gmail, Apple Mail, dark mode, images-off, and mobile constraints. Use when creating transactional emails, lifecycle emails, newsletters, bulletproof buttons, hybrid-fluid table layouts, VML Outlook fallbacks, inlined CSS, plain-text alternatives, preheaders, or client-specific rendering fixes."
argument-hint: "Goal, email type, copy or draft, brand constraints, ESP/build constraints, and target client matrix"
---

# HTML Email Templates

## What This Skill Does

Build or review HTML email the way email clients require, not the way browsers prefer.

This skill treats email as a constrained rendering target with one controlling rule: design for the worst renderer you must support, then progressively enhance for the rest. In practice that usually means Outlook for Windows decides the floor, while Gmail, Apple Mail, Outlook.com, and mobile clients get the enhancements they can support.

Use this skill to:

- Create a new production-ready HTML email template.
- Adapt a web design into email-safe markup.
- Review or debug an existing template that breaks in Outlook, Gmail, dark mode, or images-off states.
- Produce bulletproof components such as buttons, hero sections, background-image blocks, and hybrid multi-column layouts.
- Return both the HTML version and the supporting delivery artifacts that email work actually needs, such as preheader text, plain-text copy, fallback notes, and a testing checklist.

This skill is for HTML email craft. It is not for landing pages, regular web UI, or browser-only HTML.

## When to Use

- The request involves HTML email, MJML-to-HTML review, ESP template markup, or responsive email layout.
- A design must render acceptably in Outlook for Windows, Gmail, Apple Mail, Outlook.com, or mobile mail clients.
- The user needs table-based layout, inline CSS, dark-mode handling, VML, or image blocking fallbacks.
- The user asks for a transactional email, lifecycle email, marketing email, newsletter, or email component.
- The user has a broken template and needs client-specific diagnosis instead of generic HTML advice.

Do not use this skill for:

- Browser-only HTML where flexbox, grid, and normal CSS support are acceptable.
- Deliverability strategy beyond the template itself.
- Copywriting from scratch when the user has provided no message content at all.

## Inputs To Gather

Ask for whatever is missing before building or fixing the template.

- `goal` - `draft`, `review`, `debug`, or `adapt`.
- `email_type` - `transactional`, `lifecycle`, `marketing`, `newsletter`, or another specific category.
- `copy` - subject, preheader, headline, body copy, CTA text, footer copy, legal copy, and links.
- `brand_constraints` - colors, logo rules, typography limits, tone, spacing, imagery, and whether dark mode has branded requirements.
- `target_clients` - explicit client matrix if known.
- `authoring_constraints` - raw HTML, MJML source, ESP partials, template language, merge tags, or build pipeline requirements.
- `asset_constraints` - image formats, image hosting, SVG allowance, tracking parameters, and background image requirements.
- `size_constraints` - strict payload budget or clipping concerns.

If `target_clients` is missing, default to a conservative matrix:

- Gmail web
- Gmail mobile app
- Apple Mail
- iOS Mail
- Outlook for Windows
- Outlook.com
- one Android mail client

If the request is to review or debug existing email HTML, ask for the exact HTML and the failing clients. If the request is net-new and the copy is incomplete, ask whether placeholders are acceptable before inventing copy.

## Non-Negotiable Constraints

Treat these as defaults unless the user explicitly narrows the support target.

1. Table-based layout is the foundation. Do not build the core layout with divs, flexbox, grid, float, or position-based tricks.
2. Use a hybrid-fluid approach: fixed widths for Outlook compatibility, percentage widths and max-width behavior elsewhere.
3. Inline CSS is the default. Keep a style block only for rules that cannot be inlined, such as media queries, hover states, or dark-mode overrides.
4. Outlook-specific fixes belong in conditional comments or VML when needed.
5. Treat dark mode, images-off, and mobile as first-class targets.
6. Keep critical content as real HTML text, not text baked into images.
7. Watch Gmail clipping. Large decorative markup is a bug, not a style choice.
8. If a CSS feature is not clearly safe for the target matrix, avoid it or make it progressive enhancement only.

## Procedure

Follow the steps in order. Do not skip the renderer-floor decision.

### Step 1 - Establish the job and the renderer floor

Determine:

- whether the task is net-new, review, debug, or adaptation
- the email category and sending context
- the exact clients that must work
- whether the output must be raw final HTML or source for an email build system

Name the worst renderer that must remain fully functional. Usually this is Outlook for Windows. From that point on, every layout choice must remain valid there.

### Step 2 - Define the content hierarchy before touching markup

Map the email into a small number of modules:

- preheader
- logo or brand header
- headline
- supporting copy
- primary CTA
- optional secondary modules
- footer and legal content

Reduce the message to one primary action whenever possible. Transactional email should usually have one CTA and minimal decorative structure. If the hierarchy is muddy, fix the content structure before coding.

### Step 3 - Pick the safest layout pattern

Choose the layout from the support constraints, not from aesthetics.

- Single-column email: default for transactional and lifecycle messages.
- Hybrid two-column or card layout: use only when the content truly needs it.
- Background-image hero: only when the design requires it and you can provide an Outlook fallback.

When multi-column behavior is required, use the fluid-hybrid pattern so the layout still degrades acceptably when media queries fail.

### Step 4 - Build the structural scaffold

Use nested presentation tables with explicit spacing controls.

- Use `table`, `tr`, and `td` as layout primitives.
- Set `role="presentation"` on layout tables.
- Zero out `cellpadding`, `cellspacing`, and `border`.
- Keep widths explicit where Outlook needs them.
- Use predictable spacing with padding on table cells instead of margin-dependent layout.

If authoring in a readable source format first, plan for a real inlining step before final delivery.

### Step 5 - Implement bulletproof components

Build every visible component with fallbacks in mind.

- Buttons: make the CTA readable and clickable without relying on background-only styling.
- Images: always include width, alt text, and an acceptable images-off experience.
- Typography: use safe font fallbacks and readable default sizes.
- Dividers, cards, and feature rows: keep them table-based.
- Background images or rounded CTA treatments: add VML only when the visual requirement justifies the markup weight.

Prefer a plain component that works everywhere over a clever component that fails in one major client.

### Step 6 - Add progressive enhancement carefully

Enhancement is optional. Function is mandatory.

- Add media queries only after the base layout already works.
- Add dark-mode rules only after the light version is structurally sound.
- Add hover states or web-font improvements only as bonuses.
- Use Outlook conditionals for Outlook problems, not for every stylistic preference.

If a feature cannot degrade cleanly, remove it.

### Step 7 - Design for dark mode and images-off

Review the template in hostile viewing conditions.

- Ensure text still has contrast if colors invert.
- Avoid relying on pure black and pure white for main surfaces when a softer equivalent is safer.
- Make logos and icons survive dark backgrounds or inversion.
- Ensure the CTA remains understandable when images are blocked.
- Keep meaningful copy out of hero images.

Dark mode parity is not required. Legibility and brand survival are required.

### Step 8 - Add accessibility and fallback content

Email accessibility is part of the implementation, not a cleanup pass.

- Set the document language when you control the outer HTML.
- Include a descriptive title when relevant.
- Use meaningful alt text, and empty alt text for decorative images.
- Keep reading order aligned with the visual hierarchy.
- Use live text for all essential message content.
- Include a preheader that complements the subject instead of repeating it.
- Provide a plain-text version when the send system expects one.

### Step 9 - Trim weight and remove fragile markup

Before returning the template, reduce unnecessary size and risk.

- Remove unused classes, comments, redundant wrappers, and decorative cruft.
- Watch total size with Gmail clipping in mind.
- Do not keep duplicated structures unless a client-specific fallback genuinely requires them.
- Prefer simple nested tables over sprawling compatibility hacks copied without a reason.

### Step 10 - Test against the client matrix

Do not declare the template ready without a test plan.

At minimum, verify:

- the required client matrix
- mobile rendering
- dark mode where relevant
- images blocked
- click targets and link destinations
- footer and unsubscribe or legal content where applicable

If the user is asking for code only and no preview tooling is available, return the template with explicit testing notes and the highest-risk clients called out.

## Branching Logic

- **Transactional or lifecycle email.** Default to a narrow single-column layout, one main CTA, low markup weight, and maximum resilience.
- **Marketing or newsletter email.** Allow more modules, but police size aggressively and keep decorative sections subordinate to render safety.
- **Adapting a web design.** Rebuild the layout in email-safe structures. Do not transliterate div-based browser markup into email and hope.
- **Existing template fails in Outlook.** Strip the issue back to tables, widths, padding, conditional comments, and VML needs before changing visual design.
- **Existing template fails only in dark mode.** Check inversion behavior, transparent assets, and hardcoded colors before touching layout.
- **Background image is requested.** Provide an Outlook-safe fallback or explicitly recommend dropping the background image.
- **Multi-column layout is requested.** Use fluid-hybrid stacking rather than media-query-only stacking when older Gmail behavior matters.
- **Unknown CSS feature seems useful.** Treat it as unsupported until verified against the target matrix.
- **The HTML is already too large.** Cut modules, wrappers, comments, and decorative code before adding new fallback layers.

## Output Format

Return the result in the form that matches the task.

For `draft` or `adapt`:

1. The email HTML, or the corrected section when only part of the template was requested.
2. Any required supporting notes, such as where CSS must be inlined or where ESP merge tags belong.
3. Preheader text and a plain-text fallback when the request implies a send-ready email.
4. A short compatibility note naming the important fallbacks and the highest-risk clients.

For `review` or `debug`:

1. The concrete rendering problems found.
2. The corrected HTML or corrected snippets.
3. The client-specific reasoning behind the fix.
4. Remaining risks or tests that still need confirmation.

If required inputs are missing, stop and ask only for the minimum missing items.

## Quality Checks

The work is ready only when all of the following are true:

1. The layout is table-based and does not depend on unsupported CSS for structure.
2. Outlook for Windows gets a functional version of the email.
3. The main CTA is obvious, tappable, and understandable without images.
4. The email remains readable on mobile.
5. Dark mode does not destroy legibility or make brand assets disappear.
6. Images have appropriate alt handling.
7. Critical content is real text.
8. The markup is not bloated enough to make Gmail clipping likely without reason.
9. The template includes the content users expect in an email, including preheader and footer needs.
10. The response names any assumptions, unsupported requests, or unverified client behavior.

## Failure Modes To Avoid

- Treating email like normal frontend work.
- Starting from a browser layout and only later thinking about Outlook.
- Using flexbox or grid for the structural layout.
- Returning beautiful code that collapses in one mainstream client.
- Ignoring dark mode until the end.
- Hiding important text in images.
- Shipping markup so large that Gmail clipping becomes likely.
- Adding compatibility hacks that you cannot explain.
- Claiming a template is ready without naming the test matrix.

## Default Stance

Be conservative, explicit, and honest about compatibility. The plain version that works in Outlook is better than the ambitious version that breaks there. Build from table structure, inline styles, and clear fallbacks first. Treat modern CSS as a bonus layer, not as the foundation.

Assume the final deliverable is hand-authored HTML unless the user explicitly asks for MJML or another authoring system. Support both transactional and marketing email, but default toward the stricter transactional standard when tradeoffs are unclear.