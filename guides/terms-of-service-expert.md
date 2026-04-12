# Masterclass Guide to Terms of Service for European SaaS

## Role Framing

Drafting Terms of Service for a European SaaS means writing a contract that is simultaneously commercially protective, legally enforceable across 30+ jurisdictions, and compliant with a dense stack of EU law (GDPR, Unfair Contract Terms, Digital Content Directive, DSA, and increasingly DORA and the AI Act). Excellence is the discipline of saying exactly what you mean, only what the law permits, in language a non-lawyer user can actually understand.

## Best-in-Class Voices

- **Eduardo Ustaran (Hogan Lovells)** — Co-head of the global Privacy and Cybersecurity practice; his writing on EU data protection emphasizes accountability, layered transparency, and treating compliance as design, not decoration.
- **Lilian Edwards (Newcastle University)** — Editor of _Law, Policy and the Internet_; emphasizes how online contracts interact with consumer protection, platform liability, and the limits of "I agree" clickwrap.
- **Max Schrems (noyb)** — Litigator behind _Schrems I/II_; his work is a continual reminder that boilerplate international transfer clauses, dark-pattern consent, and ambiguous legal bases will be challenged and struck down.
- **Dr. Michael Veale (UCL)** — Scholar of GDPR and platform regulation; emphasizes that compliance theatre (long policies, vague purposes) fails under Articles 5 and 12, and that the DSA changes duties for intermediaries materially.
- **Ken Adams (_A Manual of Style for Contract Drafting_)** — Not EU-specific, but the global reference on eliminating ambiguity, categories of contract language (obligation, discretion, policy), and why "shall" is usually wrong. Indispensable for drafting clarity.
- **Heather Burns (tech policy writer, _Understanding Privacy_)** — Pragmatic voice for SMEs; emphasizes proportionality, realistic risk, and writing policies that match actual data practices rather than copying a template.

I'm reasonably confident about these voices' areas of focus; treat specific attributions as pointers, not quotes.

## Core Idea

A European SaaS ToS is not a shield you hide behind — it is a _transparent description of a relationship_ that EU law will rewrite in the user's favor wherever it finds vagueness, imbalance, or overreach. Draft it as if a regulator, a consumer court, and a sophisticated B2B buyer will each read it the same week.

## Essential Best Practices

1. **Separate the three documents: ToS, Privacy Policy, and DPA.**
   _Why:_ They serve different legal functions. The ToS is the commercial contract; the Privacy Policy is a GDPR Article 13/14 transparency notice (not a contract); the DPA is the Article 28 processor agreement with business customers. Conflating them creates contradictions that courts resolve against the drafter.
   _How:_ Cross-reference cleanly, version each independently, and make the DPA a pre-signed annex that B2B customers can incorporate without negotiation where possible.

2. **Decide B2B or B2C — and if both, write two sets of terms.**
   _Why:_ EU consumer protection (Unfair Terms Directive 93/13/EEC, Consumer Rights Directive, Digital Content Directive 2019/770) voids clauses that are perfectly valid in B2B: broad liability caps, unilateral change rights, forum selection away from the consumer's domicile (Brussels I recast), and short limitation periods. A single "one size fits all" ToS loses the B2B protections without gaining consumer enforceability.
   _How:_ Gate consumer sign-up behind a distinct flow with consumer-appropriate terms, withdrawal rights where applicable, and clear pre-contractual information. Keep the B2B terms leaner and more commercially aggressive.

3. **Write for transparency, not obfuscation — it is a legal requirement.**
   _Why:_ GDPR Article 12 requires "concise, transparent, intelligible and easily accessible" language. Article 5 of the Unfair Terms Directive requires "plain and intelligible" drafting and construes ambiguity _contra proferentem_ — against you. Long, defensive, Latin-heavy ToS actively lose cases.
   _How:_ Short sentences. Defined terms only when they earn their keep. Headings that mean something. A summary layer is fine, but the binding text must itself be readable. Ken Adams' categories (language of obligation, discretion, policy, declaration) keep each clause doing one job.

4. **Lock down unilateral change clauses.**
   _Why:_ A right to change terms "at any time for any reason" is presumptively unfair under 93/13/EEC Annex and has been struck down repeatedly. It also undermines the contract's own stability.
   _How:_ Specify valid reasons (legal change, security, new features), give material-change notice (30 days is a common floor), allow termination without penalty on objection, and do not treat continued use alone as acceptance of material changes for consumers.

5. **Build the international-transfer and subprocessor chain correctly.**
   _Why:_ Post-_Schrems II_, transfers outside the EEA require a valid mechanism (adequacy, SCCs with a transfer impact assessment, or BCRs). Vague "we may use third parties globally" language is a compliance failure and a sales blocker for enterprise EU buyers.
   _How:_ Maintain a public, versioned subprocessor list; commit to advance notice of new subprocessors with an objection right; reference the 2021 SCCs in the DPA; document your TIA; and align the ToS with what the list actually says.

6. **Size the liability cap and exclusions to what EU law will actually enforce.**
   _Why:_ Total exclusions of liability for death, personal injury, gross negligence, or willful misconduct are void across virtually every EU member state. Consumer caps are narrowly construed. Germany's AGB regime and French consumer code are particularly strict on standard-form caps.
   _How:_ Exclude indirect/consequential loss and cap direct loss at a defined, proportionate figure (commonly 12 months of fees for B2B). Carve out the unenforceable exclusions explicitly — it strengthens the rest of the clause rather than weakening it.

7. **Get the termination, suspension, and data-exit mechanics right.**
   _Why:_ The Digital Content Directive gives consumers remedies when digital services don't conform; GDPR Article 20 gives portability rights; the DSA imposes statement-of-reasons duties on intermediaries when restricting accounts. An "we can terminate at our sole discretion" clause conflicts with all three.
   _How:_ Define suspension grounds, require notice except in genuine emergencies, commit to an export window (commonly 30–90 days) in a usable format, and provide a reasoned notice on termination for cause.

8. **Check which sector-specific regimes apply on top of the baseline.**
   _Why:_ A SaaS selling to banks now lives under DORA (operational resilience, mandatory contractual clauses for ICT third-party providers). A SaaS acting as an "intermediary" or "hosting service" falls under the DSA. A SaaS deploying or providing general-purpose AI touches the AI Act. Each adds _mandatory_ clauses you cannot draft around.
   _How:_ Map your customer segments and technical role before drafting. If you sell to regulated industries, publish a DORA addendum; if you host user content, add DSA notice-and-action and statement-of-reasons provisions; if you provide AI, align with the Act's transparency duties.

## Common Mistakes

- **Copying a US SaaS template** and leaving in arbitration, jury waivers, Delaware forum, and 100% liability disclaimers that are unenforceable — and signal incompetence to EU legal buyers.
- **Treating the Privacy Policy as part of the contract** by "incorporating it by reference" and then changing it unilaterally, which collapses the distinction between transparency notice and binding terms.
- **Silent auto-renewal and tacit extensions** without complying with national rules (France's Loi Chatel, Germany's Fair Consumer Contracts Act) that require specific notice windows and easy cancellation.

## Quick Start

- **Map your stack first:** who are your customers (B2C/B2B/regulated), what personal data flows where, which subprocessors are outside the EEA, and what is your technical role under the DSA? The answers dictate the document, not the other way around.
- **Split into three documents** — ToS, Privacy Policy, DPA (+ subprocessor list) — and version-control each separately from day one.
- **Run a plain-language pass** after the legal draft: aim for sentences under 25 words, active voice, and defined terms only where they add precision. If a smart non-lawyer can't summarize a clause, rewrite it.
- **Stress-test the high-risk clauses** (liability cap, change-of-terms, termination, transfers) against the Unfair Terms Directive Annex and _Schrems II_ — these are where European enforcement and litigation concentrate.
- **Have it reviewed by EU-qualified counsel in at least one strict jurisdiction** (Germany or France is the usual pressure test) before launch, and re-review annually or on any material product change.
