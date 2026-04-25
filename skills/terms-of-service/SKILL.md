---
name: terms-of-service
description: "Draft, review, or debug EU SaaS terms of service and online contract terms. Use when writing subscription terms, acceptable-use clauses, liability caps and exclusions, change-of-terms provisions, termination and data-exit mechanics, B2B vs B2C contract structure, DSA/DORA/AI Act overlays, renewal and cancellation rules, governing law/forum clauses, and ensuring alignment between ToS, privacy policy, and DPA."
argument-hint: "Goal (draft/review/debug), commercial model, audience (B2B/B2C), product behavior, regulatory overlays, and existing terms"
---

# Terms of Service for EU SaaS

## What This Skill Does

Turn product flows, commercial models, and regulatory constraints into clear, enforceable, buyer-ready terms of service for European SaaS.

This skill treats the ToS as the commercial contract — distinct from the privacy notice and the DPA — and applies one controlling rule: draft for what the product actually does, the audience it actually serves, and what EU consumer and contract law will actually enforce. US-style boilerplate and unilateral discretion are defaults to remove, not import.

Use this skill to:

- Draft new SaaS subscription terms or marketplace terms for an EU-relevant product.
- Review existing terms for fairness, enforceability, and operational alignment.
- Sharpen change-of-terms, suspension, termination, liability, billing, and data-exit clauses.
- Apply DSA, DORA, AI Act, or sector-specific overlays where required.
- Align ToS with the privacy policy and DPA without collapsing the documents.

This skill is for EU SaaS contract craft. It is not for privacy notice drafting (use the privacy-policy skill), cookie banner mechanics (use the cookie-policy skill), or final legal advice.

## When to Use

- The request involves drafting, reviewing, or updating SaaS terms, subscription terms, or acceptable-use policies.
- A product is launching to EU users, B2B customers, or both, and the ToS must hold up under EU consumer and contract law.
- Existing terms include US-style overreach (broad arbitration, unilateral changes, blanket disclaimers) and need EU-fitness review.
- Sector overlays (DSA, DORA, AI Act, regulated industry) need to be reflected in contract language.
- ToS, privacy policy, and DPA need cross-document alignment.

Do not use this skill for:

- Privacy notices or transparency disclosures — use the privacy-policy skill.
- Cookie banners and ePrivacy mechanics — use the cookie-policy skill.
- Final legal advice or representing as legal counsel.

## Inputs To Gather

Ask for whatever is missing before drafting or reviewing.

- `goal` — `draft`, `review`, `debug`, or `update`.
- `commercial_model` — B2B, B2C, both, self-serve, enterprise-only, marketplace-like, user-content-hosting.
- `product_behavior` — billing model, free tiers, trials, renewal mechanics, suspension triggers, data export options.
- `regulatory_overlays` — DSA, DORA, AI Act, financial-services rules, sector-specific add-ons.
- `audience_and_jurisdiction` — primary EU markets, governing-law preference, language requirements.
- `documents_in_scope` — ToS only, ToS plus AUP, ToS plus DPA reference, ToS plus SLA.
- `existing_terms` — current text when reviewing or updating.

If the request is to review or update, ask for the actual current terms. If net-new and both audiences exist, ask whether B2B and B2C terms should be split before drafting.

## Non-Negotiable Constraints

Treat these as defaults unless the user explicitly narrows scope.

1. The ToS is the commercial contract. The privacy notice is a transparency document. The DPA covers processor obligations. Do not blur them.
2. Decide audience before drafting clauses. Blended B2B/B2C terms in Europe weaken enforceability for both sides.
3. Plain, specific language is required. Vague boilerplate, undefined carve-outs, and US-style overreach get struck down or construed against the drafter.
4. Constrain unilateral powers. Change, suspension, price-change, and termination rights need real triggers, notice, and user remedies.
5. Make liability and remedies enforceable. Caps, exclusions, warranties, and indemnities should match what EU law will actually enforce, especially around gross negligence, willful misconduct, and consumer protections.
6. Operational promises must be auditable. If the business cannot operate the clause, do not draft it.
7. Sector overlays (DSA, DORA, AI Act, etc.) are mandatory inputs when relevant, not optional add-ons.
8. Versioning, notice of material changes, cancellation mechanics, and cross-document consistency must be maintainable over time.

## Procedure

Follow the steps in order. Do not skip the audience-classification step.

### Step 1 — Establish the job and the audience

Determine:

- whether the task is draft, review, debug, or update.
- whether the product serves B2B, B2C, or both — and whether to split documents.
- the regulatory overlays that apply (DSA for hosting/intermediary, DORA for financial-services customers, AI Act for AI-system providers, sector add-ons).

If both audiences are in scope, default to separate B2B and B2C tracks, or clearly segmented sections, rather than a blended document.

### Step 2 — Map document boundaries

Decide what belongs where:

- ToS: commercial relationship, fees, term, warranties, liability, governing law.
- Privacy notice: transparency under GDPR.
- DPA: processor obligations under Article 28.
- Acceptable use policy: behavioral rules and enforcement.
- SLA or addenda: uptime, support, and product-specific commitments.

Coordinate cross-references so each document does its own job and points to the others where needed.

### Step 3 — Draft from the highest-risk clauses outward

Start with the clauses that fail most often in enforcement or buyer review:

- Change-of-terms.
- Billing, renewal, and cancellation.
- Suspension and termination.
- Liability caps, exclusions, and remedies.
- Data export and post-termination data handling.
- Governing law and forum.

Refine lower-risk language only after these are solid.

### Step 4 — Tighten unilateral powers

For each discretionary right, specify:

- The trigger (defined, not "at any time for any reason").
- Notice requirements with reasonable lead time.
- The user's remedy (cancel, refund, transition window).
- How disputes about exercise are handled.

Reject "at any time, in our sole discretion" defaults except where genuinely necessary and clearly disclosed.

### Step 5 — Make liability, warranties, and remedies enforceable

Apply EU enforceability constraints:

- Cap liability proportionately, with carve-outs for gross negligence, willful misconduct, and statutory protections.
- Avoid blanket "as is" disclaimers for B2C; align B2B disclaimers with what supreme courts of the relevant jurisdictions will actually enforce.
- Match warranty language to product reality (uptime, support, security commitments).

### Step 6 — Operationalize billing, renewal, and termination

Write clauses you can actually run:

- Billing cadence, dunning, and refunds.
- Auto-renewal mechanics aligned with member-state consumer protections.
- Cancellation paths users can find and use.
- Termination triggers, notice periods, and effects.
- Data export windows and post-termination retention.

If the operation cannot run a clause, change the operation or the clause.

### Step 7 — Apply sector and regulatory overlays

If the product is:

- A hosting or intermediary service: apply DSA notice-and-action, transparency, and trusted-flagger expectations.
- Used by regulated financial customers: apply DORA contractual expectations on incident reporting, exit, and subcontractor management.
- An AI system or providing AI components: apply AI Act obligations relevant to the product role (provider, deployer, GPAI provider).
- Sector-specific (health, financial, public sector): apply the relevant national overlays.

Treat these as required content, not optional.

### Step 8 — Align with privacy policy and DPA

Confirm:

- Privacy disclosures live in the privacy notice, not the ToS.
- Processor obligations live in the DPA, with the ToS pointing to it.
- Subprocessor handling, transfers, and security measures are not duplicated inconsistently across documents.

### Step 9 — Stress-test enforceability

For each clause, ask whether a strict member-state court or consumer authority would view it as transparent, proportionate, and operationally supported. Where the answer is unclear, narrow the clause or add an explicit carve-out.

### Step 10 — Wire versioning and change governance

Specify:

- Effective date and version metadata.
- Material-change notice mechanics.
- How users accept ongoing changes (continued use, in-product confirmation, etc.) compatible with applicable consumer law.
- Where prior versions are accessible.

## Branching Logic

- **B2C consumer SaaS.** Default to consumer-friendly defaults: short notice on changes, clear cancellation, narrow disclaimers, strong refund/withdrawal rights, governing law respecting consumer-protection rules.
- **B2B enterprise SaaS.** Allow stronger limitations and broader disclaimers, but keep gross-negligence and willful-misconduct carve-outs; align with the DPA and SLA.
- **Both B2B and B2C in scope.** Recommend splitting the documents or using clearly segmented sections; do not run both audiences through one blended set of terms.
- **Hosting or intermediary service.** Add DSA-aligned notice-and-action, illegal-content handling, and transparency content.
- **Financial-services customers.** Add DORA-aligned exit, subcontractor, and incident-reporting language; expect customer-side audit rights.
- **AI features.** Add AI Act-aligned role descriptions, intended-use limits, and acceptable-use restrictions on prohibited uses.
- **User-generated content.** Add an acceptable use policy with enforcement mechanics; align suspension/termination with the AUP triggers.
- **Reviewing maximalist US-style terms.** Lead with the clauses most likely to be unenforceable in the EU and provide narrowed, EU-fit replacements.

## Output Format

Match the deliverable to the task.

For `draft`:

1. The terms in clause-by-clause structure with plain headings.
2. A short clause map at the top showing what lives in the ToS vs the privacy notice, DPA, AUP, or SLA.
3. Versioning and change-governance language.
4. "Fact Checks Required" list for business behavior, billing, vendors, and operational practices.
5. "Legal Review Required" note for jurisdiction-sensitive or sector-sensitive issues.

For `review` or `debug`:

1. Risk and enforceability summary at the top, ordered by severity.
2. Findings grouped by commercial scope, fairness/enforceability, operational alignment, and regulatory overlays.
3. Concrete clause replacements or redlines for high-risk language.
4. "Fact Checks Required" and "Legal Review Required" lists.

For `update`:

1. Redline of changed clauses.
2. New version metadata and a short user-facing change summary.
3. Notice mechanics and user options where the change is material.

## Quality Checks

The work is ready only when all of the following are true:

1. The audience is clear, and B2B and B2C are not blended in a way that weakens either.
2. ToS, privacy notice, and DPA boundaries are clean.
3. Change, suspension, and termination rights have triggers, notice, and remedies.
4. Liability caps and exclusions are proportionate and survive likely EU enforcement.
5. Billing, renewal, cancellation, and data-exit clauses match operational reality.
6. Sector and regulatory overlays (DSA, DORA, AI Act, etc.) are reflected when relevant.
7. Cross-document references are consistent.
8. Versioning and material-change communication are operable.

## Failure Modes To Avoid

- Importing US-first defaults (mandatory arbitration, jury waivers, blanket disclaimers, unilateral forum) into EU terms.
- Merging privacy disclosures or processor obligations into the ToS unless coordinating language is explicitly needed.
- Drafting unilateral change, suspension, or termination rights with no triggers or remedies.
- Promising subprocessor, transfer, portability, refund, or export rights that are not operationally implemented.
- Ignoring DSA, DORA, AI Act, or sector overlays where they apply.
- Outputting final legal advice or claiming to act as legal counsel.

## Default Stance

Be specific, fair, and operationally honest. The narrower clause that matches what the business can run beats the maximalist clause that a regulator will strike. Lead with concrete clause language, redlines, and issue lists over abstract legal commentary. When member-state law diverges, default to the stricter realistic EU bar and flag the rest for qualified legal review.
