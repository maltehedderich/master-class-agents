---
name: cookie-policy
description: "Draft, review, or debug EU cookie policies, banners, and consent flows for SaaS products. Use when classifying cookies and SDKs, mapping strictly necessary vs consent-based tracking, blocking tags before consent, configuring a CMP, logging consent and withdrawal, designing first-layer banner UX, or aligning ePrivacy and GDPR disclosures for European users."
argument-hint: "Goal (draft/review/debug), product context, tag inventory, banner design, target jurisdictions, and CMP details"
---

# Cookie Policy and Consent

## What This Skill Does

Turn tag inventories, banner UX, and runtime behavior into clear, defensible, regulator-ready cookie programs for European SaaS.

This skill treats cookies and similar technologies as a constrained compliance target with one controlling rule: design for the strictest realistic EU interpretation, then prove it with operational evidence. In practice that usually means CNIL-style banner symmetry, EDPB Cookie Banner Taskforce expectations, and noyb-style enforcement pressure.

Use this skill to:

- Draft a new cookie policy and first-layer banner copy.
- Review or debug an existing banner, CMP configuration, or policy that fails CNIL/EDPB expectations.
- Classify cookies, SDKs, and embedded scripts as strictly necessary vs consent-based.
- Audit runtime tag-firing behavior against the claimed banner design.
- Produce consent-logging and withdrawal mechanics that hold up as evidence.

This skill is for European cookie compliance craft. It is not for US-only consent regimes, full DPA drafting, or generic privacy policy work outside cookie/ePrivacy mechanics.

## When to Use

- The request involves cookie banners, consent flows, CMP configuration, cookie policies, or ePrivacy disclosures.
- A SaaS product needs to ship to EU users and must classify or block tags correctly.
- A user has a banner that regulators or noyb-style complaints would likely flag.
- A team needs runtime verification that non-essential tags do not fire before consent.
- Consent records and withdrawal logging must be evidence-grade.

Do not use this skill for:

- Full GDPR privacy policy drafting beyond cookie/tracking sections — use the privacy-policy skill instead.
- Terms of service or DPA drafting — use the terms-of-service skill or a DPA workflow.
- Final legal advice or representation as legal counsel.

## Inputs To Gather

Ask for whatever is missing before drafting or reviewing.

- `goal` — `draft`, `review`, `debug`, or `classify`.
- `product_context` — B2B vs B2C, regulated sector, embedded widgets, joint-controller scenarios.
- `tag_inventory` — cookies, local storage, SDKs, third-party scripts, pixels, chat widgets, embedded resources.
- `banner_design` — first-layer copy, button labels, symmetry, layered preferences, withdrawal path.
- `cmp_details` — vendor, TCF use, custom integration, blocking strategy.
- `runtime_behavior` — what fires before consent, what changes after opt-in or rejection, withdrawal effects.
- `target_jurisdictions` — explicit member-state focus if known. Default to the strictest realistic EU bar (often FR/DE).
- `existing_policy_or_banner` — the exact text or screenshots when reviewing.

If the request is to review or debug, ask for the actual banner, the CMP configuration, and a list of tags or a network trace. If net-new, ask whether the user has a documented tag inventory before drafting policy text.

## Non-Negotiable Constraints

Treat these as defaults unless the user explicitly narrows scope.

1. Separate device access (ePrivacy) from downstream processing (GDPR). Do not collapse them into one vague legal basis.
2. Consent must be prior, opt-in, granular, and symmetrical. "Accept all" and "Reject all" must be equally available on the first layer.
3. "Strictly necessary" is a narrow exemption. Analytics, attribution, A/B testing, session replay, personalization, and product improvement usually do not qualify.
4. Block first, then prove it. Non-essential technologies must not fire before valid consent. Withdrawal must stop processing.
5. Inventory every cookie, SDK, tag, and embedded script before it ships. Purpose, provenance, duration, recipients, category, and triggering conditions must be documented.
6. Treat consent records as evidence. Log banner version, policy version, timestamp, user/device pseudonym, exact choices, and withdrawal events.
7. A CMP or TCF implementation is not automatic compliance.
8. Do not use "legitimate interest" to justify placing non-essential cookies or similar device identifiers.

## Procedure

Follow the steps in order. Do not skip the inventory step.

### Step 1 — Establish the job and the regulatory floor

Determine:

- whether the task is draft, review, debug, or classify.
- the product audience (B2B, B2C, embedded), the sector, and the realistic EU bar.
- whether the deliverable is policy text, banner UX, classification table, or runtime audit findings.

Name the strictest jurisdiction the product must hold up under. Usually FR (CNIL) or DE sets the practical floor. Every later decision must remain valid there.

### Step 2 — Map the technology before drafting

Build or validate a concrete inventory:

- Cookies and local storage usage by domain and purpose.
- SDKs, third-party scripts, pixels, chat widgets, embedded resources.
- Who sets each item, what data it touches, and what triggers it.

If the inventory is missing, stop and request it. Do not draft policy text against unknown tags.

### Step 3 — Classify each item

For each entry, decide:

- Strictly necessary vs consent-based, with a concrete exemption analysis when claiming necessity.
- Purpose category (e.g., security, session, analytics, advertising, personalization).
- Lawful basis under GDPR for the resulting processing, separately from the ePrivacy access basis.
- Recipients and any international transfer implications.
- Retention period or expiry logic.

If a tag cannot be classified concretely, it should not go live.

### Step 4 — Audit runtime behavior

Verify what actually happens, not what the banner claims:

- What loads before consent.
- What changes after opt-in vs reject-all.
- How withdrawal behaves and whether identifiers are removed or expired.
- Whether the network behavior matches the banner and policy text.

Flag any mismatch as a compliance defect, not a style issue.

### Step 5 — Design the first-layer banner

Apply CNIL-style expectations:

- Equal-prominence "Accept all" and "Reject all" on layer one when consent is required.
- Granular controls reachable in no more layers than required to consent.
- No pre-ticked toggles, no continued-browsing consent, no dark patterns.
- Withdrawal must be at least as easy as consent (persistent footer link or equivalent).

### Step 6 — Draft the cookie policy and detailed preferences

Layer the disclosures:

- Short, plain-language explanation on the first layer.
- Detailed preferences screen with categories and toggles.
- A cookie table listing each item with purpose, provenance, duration, recipients, and category.
- A "manage your choices" entry point that survives session end.

Use concrete language. Avoid generic placeholders that cannot be verified against operations.

### Step 7 — Wire consent logging and versioning

Make consent auditable:

- Log banner version, policy version, timestamp, pseudonymous identifier, exact choices, and withdrawal events.
- Version the banner and policy independently and record which version the user saw.
- Make logs immutable enough to defend in a complaint or audit.

### Step 8 — Stress-test against complaint scenarios

Before finalizing, test against the most common failure patterns:

- Asymmetric banner choices.
- Tags firing before consent.
- "Legitimate interest" claimed for non-essential tags.
- Withdrawal that does not actually stop processing.
- Vendor or transfer disclosures missing from the cookie table.
- Joint-controller exposure for embedded widgets.

### Step 9 — Flag legal-review boundaries

Identify clearly:

- Jurisdiction-sensitive interpretations (member-state variation).
- Sector-specific obligations (e.g., financial, health, children).
- Anything where qualified EU counsel should sign off before publication.

## Branching Logic

- **Net-new SaaS.** Default to a strict CNIL-style banner, narrow exemptions, and a small initial tag set. Push back on shipping non-essential tags before consent infrastructure is real.
- **Existing banner under complaint pressure.** Start with runtime behavior and symmetry, not policy wording. Most complaints succeed because tags fire before consent or because reject-all is harder than accept-all.
- **B2B SaaS with embedded widgets.** Treat the host site, the customer, and the end user as separate data subjects. Document role boundaries explicitly.
- **TCF / IAB framework in use.** Verify that vendor purposes and special features map to the user's actual disclosures. TCF is a transport, not a compliance certificate.
- **Analytics-only product.** Do not assume analytics is exempt. Either obtain consent or use a narrow, documented analytics exemption mode (no cross-site tracking, no personal data, no third parties).
- **Tag inventory is unknown or stale.** Stop and rebuild the inventory before producing policy text or a classification table.

## Output Format

Match the deliverable to the task.

For `draft`:

1. First-layer banner copy and button labels.
2. Cookie policy text in layered structure.
3. Cookie table with purpose, provenance, duration, recipients, category.
4. Consent logging and versioning notes.
5. A "Fact Checks Required" list for anything that depends on internal implementation details.
6. A "Legal Review Required" note for jurisdiction-sensitive issues.

For `review` or `debug`:

1. Concrete compliance gaps grouped by banner UX, classification, blocking behavior, logging and evidence, disclosures, and role boundaries.
2. Exact remediation steps and revised banner or policy language where useful.
3. Highest-priority risks called out first.
4. "Fact Checks Required" and "Legal Review Required" lists.

For `classify`:

1. Per-item classification with exemption reasoning.
2. Items where strictly-necessary status is contested or unsupported.
3. Recommended consent treatment for each contested item.

## Quality Checks

The work is ready only when all of the following are true:

1. Every tag in the inventory has a documented purpose, provenance, duration, and category.
2. Strictly-necessary claims are narrow and defensible, not convenience-driven.
3. The first-layer banner is symmetric and free of dark patterns.
4. Non-essential tags do not fire before consent, and withdrawal actually stops processing.
5. Consent logs capture banner version, policy version, choice, and withdrawal events.
6. The cookie table reflects what the product actually does today.
7. International transfers, recipients, and joint-controller scenarios are disclosed where they exist.
8. Member-state-sensitive points are flagged for qualified legal review.

## Failure Modes To Avoid

- Treating a CMP or TCF integration as automatic compliance.
- Labeling analytics, experimentation, session replay, personalization, or marketing tags as essential.
- Using "legitimate interest" to justify placing non-essential cookies.
- Asymmetric banners, pre-ticked toggles, or harder-to-find withdrawal paths.
- Drafting policy text against an unknown or stale tag inventory.
- Describing categories, retention periods, or controls that are not implemented.
- Outputting final legal advice or claiming to act as legal counsel.

## Default Stance

Be conservative, explicit, and honest about what the product actually does. The narrow, defensible reading of an exemption is better than the convenient one. Lead with concrete cookie tables, consent logs, and network-level verification over abstract compliance language. When member-state interpretations diverge, optimize for the stricter realistic EU bar.
