# Masterclass Guide to Cookie Policies for European SaaS Products

## Role Framing

Excellence here means designing cookie compliance that satisfies GDPR and the ePrivacy Directive without degrading UX or analytics quality — translating legal nuance into clear consent flows, defensible documentation, and engineering practices that hold up under DPA scrutiny.

## Best-in-Class Voices

- **Dr. Johnny Ryan (Irish Council for Civil Liberties, formerly Brave)** — Leading enforcement advocate on real-time bidding and consent; emphasizes that "legitimate interest" rarely saves ad-tech cookies and that consent must be genuinely granular.
- **Max Schrems / noyb** — Source of the complaints that reshaped cookie banner enforcement across the EU; emphasizes "reject" must be as easy as "accept," no pre-ticked boxes, no dark patterns, no cookie walls.
- **CNIL (French DPA)** — The most prescriptive regulator on cookies; their published guidelines and sanctions against Google, Amazon, and Meta set the de facto European baseline on banner design and 24-hour logging of consent.
- **EDPB (European Data Protection Board)** — Guidelines 03/2022 on dark patterns and the Cookie Banner Taskforce report define what "valid consent" concretely looks like across member states.
- **Peter Hense (Spirit Legal)** and **Dr. Carlo Piltz** — German practitioners writing clearly on TTDSG/TDDDG and the interplay with GDPR; useful for the strictest member-state interpretation most SaaS vendors should design against.
- **IAB Europe / TCF critics and defenders** — Worth studying both sides: the Belgian DPA's 2022 ruling against TCF v2.0 shows why blindly adopting a CMP framework isn't a shield.

## Core Idea

Under EU law, storing or reading _any_ non-strictly-necessary information on a user's device requires prior, specific, informed, freely given, and revocable consent — and the burden of proof sits entirely on you. Compliance is a UX and engineering problem disguised as a legal one.

## Essential Best Practices

1. **Separate the two legal bases you're actually dealing with.** The ePrivacy Directive (Article 5(3), implemented nationally — TTDSG in Germany, PECR in the UK, etc.) governs the _act of storing/reading_ anything on the device. GDPR governs what you then _do_ with the personal data. Why it matters: many teams conflate them and wrongly invoke "legitimate interest" for analytics cookies — ePrivacy doesn't offer that basis for storage. Apply it well by running two separate analyses per cookie: "Can I place it?" (ePrivacy) and "Can I process the data?" (GDPR).

2. **Consent must be opt-in, granular, and symmetrical.** A valid banner has a "Reject all" button on the first layer, visually equivalent to "Accept all." No pre-ticked boxes, no "by continuing to browse you consent," no cookie walls unless a genuine equivalent alternative exists. Why it matters: this is the single most-enforced point — CNIL fined Google and Facebook €150M and €60M specifically for asymmetric banners. Apply it well by A/B testing only _within_ compliant designs; any pattern that nudges toward acceptance is a dark pattern under EDPB Guidelines 03/2022.

3. **Classify every cookie and SDK before it ships.** Maintain a living inventory with purpose, provenance (first- or third-party), retention, data recipients, and legal basis. Why it matters: Article 13 GDPR transparency requirements and the ePrivacy "informed" requirement both depend on this, and regulators ask for it first in any investigation. Apply it well by gating new third-party scripts in CI — no tag goes live without an entry in the register and a review of whether it's truly strictly necessary.

4. **"Strictly necessary" is narrow — narrower than your marketing team thinks.** Session IDs, load balancing, CSRF tokens, and the user's own consent choices qualify. Analytics, A/B testing, session replay, and "to improve our service" do _not_, regardless of anonymization claims. Why it matters: mis-categorizing analytics as essential is the most common violation in SaaS. Apply it well by defaulting to "requires consent" and only exempting cookies that meet the Article 29 WP 194 criteria.

5. **Log consent like you'd log financial transactions.** Store timestamp, banner version shown, exact choices, policy version, and a mechanism to prove withdrawal was as easy as giving consent. Why it matters: GDPR Article 7(1) puts the burden of proof on the controller; without logs, you cannot defend yourself. Apply it well by versioning your banner and policy text, and keeping an immutable audit trail linked (pseudonymously) to the user or device.

6. **Withdrawal must be as easy as consent — and visible on every page.** A persistent footer link or floating icon that re-opens preferences is now the expected pattern. Why it matters: this is explicit in Article 7(3) and was central to the CNIL decisions. Apply it well by wiring withdrawal to actually delete the cookies and stop the scripts, not just flip a flag.

7. **Don't fire a single non-essential tag before consent.** This includes Google Analytics, Hotjar, Intercom chat, HubSpot, LinkedIn Insight Tag, Meta Pixel, and most CDN-injected scripts. Why it matters: Schrems II and subsequent DPA decisions (notably the Austrian and French DPAs on Google Analytics) compound the problem when US transfers are involved. Apply it well through a tag manager with true consent-mode blocking, server-side tagging where appropriate, and periodic network-tab audits — banners often lie about what's actually loading.

8. **For B2B SaaS, remember the account holder isn't the data subject.** Your customer's employees and _their_ end users have rights too, and your DPA with the customer must reflect processor obligations under Article 28. Why it matters: cookies set via your embedded widgets make you a joint or independent controller in ways most SaaS contracts don't address. Apply it well by documenting roles per data flow and exposing a consent API your customers can integrate with their own CMP.

## Common Mistakes

- Treating a TCF-based CMP as automatic compliance — the Belgian DPA's 2022 ruling on IAB Europe showed it isn't.
- Writing the cookie policy once, then letting the tag inventory drift for 18 months until a marketer adds a new pixel.
- Using "legitimate interest" for tracking cookies — ePrivacy Article 5(3) doesn't permit it for device storage, only GDPR does for the downstream processing.

## Quick Start

- Run a full cookie and tag audit today using browser devtools plus a scanner (OneTrust, Cookiebot, or open-source alternatives like cookie-autodelete); you'll likely find 20–40% more than you expected.
- Redesign your first-layer banner to show "Accept all," "Reject all," and "Customize" as equally prominent buttons — fix this before anything else.
- Block all non-essential tags server-side or via consent mode until an explicit opt-in is logged, and verify in the network tab with cookies cleared.
- Stand up a consent log with banner version, policy version, and choices, retained for the duration of the consent plus a reasonable audit window.
- Publish a plain-language cookie policy with a live, auto-generated table of cookies, purposes, retention, and recipients — and link to it from the banner and the footer.

One caveat: member-state implementations diverge (Germany's TDDDG, France's CNIL guidelines, and the UK's post-Brexit PECR all have quirks), so validate final designs against the strictest jurisdiction you serve — usually France or Germany — and get a qualified EU privacy lawyer to review before launch.
