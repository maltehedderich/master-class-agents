---
name: apple-ads-keywords
description: Select and prioritize Apple Ads search-results keywords for an app, producing a same-day test shortlist, negative keywords, and explicit console-check instructions. Use for a new keyword plan, market-specific keyword research, or refining a list from search-term performance. Focuses on Manage Bids campaigns, not App Store keyword metadata or web SEO.
---

# Apple Ads Keyword Selection

Reach a concrete keyword selection today. Optimize for search intent the app can satisfy, then demand, conversion readiness, and economics. The deliverable is a justified test selection, not a claim that profitability has been established.

## Working contract

- Use existing product information, reviews, support messages, research, and campaign exports. Customer interviews are useful input when already available; never require new interviews, surveys, paid research tools, or a live advertising test before selecting keywords.
- Fill gaps in audience language, demand hypotheses, and conversion expectations with reasonable, explicitly labeled assumptions. Never invent product capabilities, customer quotes, console observations, search popularity, bids, or performance data.
- Ask only for missing facts that materially change the selection. If the app and its capabilities are completely unknown, request a short description or product link. Infer other inputs from available context and show assumptions for correction.
- Distinguish **selection readiness** from **launch readiness**. Missing account access or economics can leave launch checks pending while a concrete provisional shortlist is delivered today.
- Selection does not authorize launching campaigns, changing bids, or spending money. Respect any existing authorization; otherwise deliver the plan without making account changes.

## 1. Establish the smallest useful brief

Read supplied files, product pages, and account data before asking questions. Capture:

| Input | If unavailable |
| --- | --- |
| App, core capabilities, unsupported functions | Ask for essential product facts; hold terms requiring unconfirmed features. |
| Audience and job to be done | Infer from capabilities; label the inference. |
| Country or region and search language | Propose one initial market from context; mark it assumed. Keep each additional market separate. |
| Pricing, free tier, and valuable customer action | Use known pricing; otherwise leave economics unverified and use a provisional activation objective. |
| Test budget and acceptable acquisition cost | Ask briefly if needed; continue selection with a smaller ranked set and unapproved budget scenarios. |
| Existing campaign strategy, keywords, negatives, and results | Assume a new Manage Bids plan for drafting; request a strategy check before implementation. |

Explain the intended outcome briefly: a ranked first test, exclusions, optional discovery seeds, and remaining launch checks. A useful default is a 60–120 minute working session with one compact round of user checks, not a research project. Adjust to the user's time and available evidence.

This workflow's manual bids and targeting structure apply to **Manage Bids search-results campaigns**. If the user has Maximize Conversions, preserve that choice: finish intent selection, but mark manual-bid setup as inapplicable and verify which keyword controls their ad group supports. Do not silently switch strategies. Apple permits specific keywords in an additional Maximize Conversions ad group, but does not offer keyword-level bids there. See [Apple's strategy documentation](https://ads.apple.com/app-store/help/campaigns/0095-maximize-conversions).

## 2. Generate candidates and apply the relevance gate

For each intent family, complete:

> Someone searching for [phrase] wants [outcome], which this app delivers through [actual capability].

Build candidates from existing customer language, the product, Apple suggestions, local App Store results, and historical search terms when available. Record the source per candidate; a brainstorm is not observed customer language.

Separate own brand, core function, specific use case, desired outcome, broad category, competitor brand, and unsupported function. Start with function and use-case phrases. Competitor candidates must represent credible alternatives; do not import every app in the same category.

Apply relevance before ranking:

- **Pass:** The product directly meets the likely expectation.
- **Hold:** Intent is ambiguous or depends on a capability that needs confirmation.
- **Exclude:** The phrase requires an absent feature, unsupported audience, or misleading promise.

Popularity cannot rescue a failed relevance gate. For example, a timer with session history but no blocking or calendar planning can fit `study timer`; it cannot satisfy `app blocker`, and `study planner` needs scrutiny. These are illustrative candidates, not verified recommendations for the user's app.

Research country and language together. Translations are candidates until local intent is checked. Do not merge markets' popularity or performance into one ranking.

## 3. Gather the checks that can change today's selection

Perform accessible research yourself. When account access or a local storefront requires the user, read [User checks](references/user-checks.md) and give only the relevant tasks, populated with the actual market and candidate phrases.

Every task must say **where to go, what to do, what to return, why it matters, and what happens if unavailable**. Batch a few high-value checks instead of sending the user on an open-ended investigation. Prioritize ambiguous high-ranking terms and likely budget constraints. Continue independent candidate work while awaiting answers.

Use console suggestions and popularity as supporting evidence, and storefront results as qualitative intent evidence. Apple's documented popularity scale is 1–5, not monthly search volume. Missing recommendations or popularity does not establish zero demand. Estimates do not establish future results. See [Apple's recommendations documentation](https://ads.apple.com/app-store/help/recommendations/0055-review-recommendations).

When the user returns checks, update the affected rows and explain additions, removals, and rank changes. If checks cannot be completed today, finalize a provisional selection with `unknown` observations and explicit launch checks. Never end with only instructions to return after doing research.

## 4. Rank a test the budget can teach us about

Rank only passing terms, considering:

1. **Intent clarity:** How directly does the phrase express a supported job?
2. **Demand evidence:** Actual search terms, console signals, or research; unknown where absent.
3. **Conversion readiness:** Can the existing ad, product page, and first-use experience demonstrate that job?
4. **Economic plausibility:** Could the resulting valuable users justify the acquisition cost?

Give a short reason and an evidence label: `assumed`, `research-supported`, `tested`, or `supported by mature performance`. Research-supported does not mean profitable. Prefer qualitative judgments to a weighted score that disguises missing evidence.

Roughly 10–20 exact-match terms can be a starting heuristic, not a quota. Choose fewer for a small budget or narrow product; never pad the list. Keep other plausible terms in reserve. Estimate learning capacity as test budget divided by assumed average CPT, explaining that delivery is uneven and any equal split is only illustrative.

For a subscription objective, use:

`affordable average CPT = target cost per payer × P(install | tap) × P(payer | install)`

Use consistent cohorts and contribution/payback assumptions. If rates or the target are unknown, show labeled conservative/base scenarios or leave the limit unverified; do not manufacture a validated bid. An affordable average CPT is not the max CPT bid setting. Suggested Apple bids are auction inputs, not proof of affordability. Numeric bid and budget proposals remain proposals until authorized.

Make a decision despite uncertainty: select a small relevant test set and state which assumptions could change its size or priority. Do not call low-CPI terms winners without considering the valuable customer action. If downstream attribution is only available by ad group or campaign, evaluate at that level.

## 5. Specify targeting, exclusions, and the promised experience

For a Manage Bids plan, use this structure where the budget supports it:

| Purpose | Positive targeting | Search Match |
| --- | --- | --- |
| Core function / use case | Selected exact-match terms | Off |
| Own brand, if useful | Separate exact-match terms | Off |
| Competitors, if justified | Separate exact-match terms | Off |
| Optional broad discovery | A small set of relevant broad-match seeds | Off |
| Optional Search Match discovery | No explicit keywords required | On |

Do not require all five groups or a discovery spend allocation. Budgets are controlled at campaign level; ad-group themes do not create independent budgets. See [Apple's campaign structure guidance](https://ads.apple.com/app-store/best-practices/campaign-structure).

Specify match type explicitly. Positive exact match includes close variants; broad match reaches related searches. Search Match is automatic matching, not a third keyword match type. Generic intent remains generic even with exact match. New positive keywords default to broad; after saving, changing match type requires pausing and adding a replacement. See [Apple's match-type documentation](https://ads.apple.com/app-store/help/keywords/0059-understand-keyword-match-types).

Give each proposed negative a term, match type, purpose, and precise destination:

- **Unsuitable intent:** Exclude clearly unsupported needs. Use broad negatives only when excluding every query containing those words is intended. Do not automatically exclude `free`; check the offer.
- **Discovery routing:** Add selected core terms as exact negatives in discovery. If discovery shares a campaign with core groups, place these on discovery ad groups, never the shared campaign.

Exact negatives block the specified term without close variants or reordered phrases. Broad negatives require every word, in any order, and do not automatically cover variants. Review positive/negative conflicts and do not promise complete routing isolation. See [negative matching rules](https://ads.apple.com/app-store/help/keywords/0059-understand-keyword-match-types).

Map each selected intent group to its product-page promise and first useful action. Use adequate existing assets; do not require new custom product pages to finish selection. Hold a term if its promise would misrepresent the product. Keep competitor advertising candidates separate from App Store keyword metadata.

## 6. Deliver the selection today

Save `docs/apple-ads-keywords-<app>-<market>.md` in the workspace unless the user requests another format or chat only. Preserve unrelated files. Include:

1. **Decision:** The selected test, market/language, objective, and whether it is ready for setup or provisional.
2. **Selected keywords:** Ranked table with keyword, intent, supporting capability, source/evidence, demand signal or `unknown`, confidence, match type, destination group, and selection reason. Put a copyable exact-keyword list beside it for manual entry; state that Exact must be chosen in the console.
3. **Reserve and rejected terms:** Separate holds from exclusions and explain what would change a hold.
4. **Negatives and optional discovery:** Exact terms, types, destinations, and reasons; explicitly say when discovery is omitted.
5. **Test constraints:** Available budget, economic assumptions, proposed spend limits, product-page fit, and outstanding launch checks with owners.
6. **Evidence and next review:** Checks completed, assumptions still open, and how observed search terms will refine the plan.

Keep the decision table human-readable. Only create an Apple upload CSV when requested, using the current console template and real campaign/ad-group identifiers; a planning table is not an upload file.

Define a follow-up based on traffic, spend limits, and conversion maturity rather than an arbitrary number of days. Separate brand from nonbrand, new downloads from redownloads, and tap-through from total acquisition metrics. Review actual search terms even for exact keywords. Diagnose absent impressions for bid, budget, targeting, relevance, and negative conflicts before concluding there is no demand. Do not treat tiny samples or unfinished trials as proof of failure.

Promote relevant discoveries to controlled exact tests and add appropriately scoped discovery negatives. Keep selection, testing, and proven performance distinct. A term is only proven for a specified objective, market, cohort window, and observed economics.

Before finishing, verify that the output contains actual selected keywords, every selected intent is supported, missing evidence is labeled, exclusions cannot block the intended core test, and no optional research has become a prerequisite. Return the artifact link, the recommended first test, and only the remaining user actions that matter.
