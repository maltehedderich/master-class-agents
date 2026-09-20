# Guided user checks

Use this reference when the user must inspect Apple Ads or a storefront you cannot access. Send only applicable checks, filled with the app, market, and concrete keywords. A normal round should take about 15–30 minutes; ask for fewer terms when time is tight. If a result would not change selection or setup, skip that check.

Separate observation from implementation. Research checks should not require saving keywords, applying recommendations, switching strategies, or starting spend. UI labels can change: verify current Apple help when navigation differs, or ask the user for the labels they see. Never claim to have observed their console yourself.

## 1. Verify campaign context

**Where:** Apple Ads → Campaigns → the campaign for the app → campaign settings; then inspect the relevant ad group's settings.

**Ask the user to:** Read the placement, bid strategy, countries/regions, campaign budget and currency, ad-group Search Match setting, and any restrictive audience settings. Do not change them.

**Return:** `placement | strategy | countries/regions | budget + currency | ad group | Search Match on/off | audience restrictions`.

**Decision:** Determine whether the manual-keyword setup applies and whether the selected market matches the campaign. In Maximize Conversions, the automated group requires Search Match and keyword-level bids are unavailable; an additional ad group can support selected keywords. Keep the user's strategy unless they authorize a change. [Apple strategy help](https://ads.apple.com/app-store/help/campaigns/0095-maximize-conversions).

**If unavailable:** Draft for an assumed Manage Bids search-results setup. Mark strategy and market as launch checks. Do not make account creation a prerequisite for keyword selection.

## 2. Inspect suggestions and available demand signals

**Where:** An existing search-results ad group → Add Keywords. For available recommendations, use Recommendations from the campaign or keyword dashboard → Keywords for the app, filtered to the intended country/region.

**Ask the user to:** Search the supplied 3–5 seed phrases in the suggestion search field and note related terms that express the same supported job. Read rather than save or apply recommendations. If the interface cannot be inspected without creating or changing a campaign, skip it for now. [Apple keyword help](https://ads.apple.com/app-store/help/keywords/0014-add-and-manage-keywords).

**Return:** `term | country/region | source screen | displayed popularity or unavailable | suggested max CPT + currency or unavailable | useful related terms`. Include any estimates they already see, labeled as estimates; do not require every field.

**Decision:** Reapply the relevance gate to suggestions, then use demand and bid signals to order passing terms. Record popularity on its displayed scale, never convert it to monthly searches. A missing recommendation is not evidence of no demand; new or low-activity accounts may have none. [Apple recommendations help](https://ads.apple.com/app-store/help/recommendations/0055-review-recommendations).

**If unavailable:** Use product-derived candidates and any accessible storefront research. Set console evidence to `unknown` and keep a concrete provisional shortlist.

## 3. Resolve ambiguous local intent

**Where:** App Store search in the intended storefront and language. Ask the user to confirm the storefront they actually used; do not require changing their Apple Account region.

**Ask the user to:** Search the supplied 3–5 ambiguous or priority phrases. Note roughly the first five results: app names, dominant function, brand concentration, and whether the user's app screenshots would answer the same request.

**Return:** `phrase | storefront/language | dominant result type | example apps | fits our app: yes / uncertain / no`.

**Decision:** Drop or hold terms dominated by unsupported functions or brand expectations the app cannot credibly satisfy. Treat the results as qualitative intent evidence, not search volume or a full auction view.

**If unavailable:** Use accessible local listings and label the inference. Hold ambiguous phrases; select clearer function terms. Lack of local access need not block the entire market shortlist.

## 4. Reuse existing campaign evidence, if there is any

**Where:** The relevant keyword dashboard and Search Terms view. Filter to the market and record the date range. Use downloadable reports if available. See [Apple reporting help](https://ads.apple.com/app-store/help/reporting/0024-view-campaigns-dashboard-metrics).

**Ask the user to:** Share existing keyword/search-term rows with campaign/ad group, targeting source or match type, impressions, taps, spend, and installs/new downloads where available. Also collect current campaign and ad-group negatives. If they already have activation or payer data, return it at the level their attribution supports, with cohort age and trial duration.

**Return:** An export or compact pasted table, plus market, currency, date range, metric definitions, negative scope, and any downstream cohort notes. Do not require a new tracking integration.

**Decision:** Identify supported queries to select, irrelevant searches to exclude, exhausted tests, and negative conflicts. Keep keyword inputs separate from the queries people typed. Do not infer query-level revenue from campaign totals or count immature trial cohorts as final non-payers.

**If unavailable:** Treat this as initial selection. Historical spend and a completed advertising test are optional inputs.

## 5. Hand off setup after the selection is complete

This is a launch checklist, not a condition for delivering keywords. If the user is implementing, give them the named destinations and keyword lists from the final plan, then ask them to verify:

- The intended app, market, strategy, budget, and bids match the plan and authorized spend.
- Core keywords are explicitly set to Exact and Search Match is off in those Manage Bids groups. Do not accept default Broad by accident.
- Any discovery groups have the specified matching settings. Routing negatives are confined to discovery; no campaign-level negative blocks a selected core term.
- The selected product page visibly supports the intent. Confirm any held capability before including its terms.

Ask the user to return discrepancies or confirm the settings before they activate spending. If activation was not requested, end with the plan and pending checks. Missing launch approval never prevents delivering the selection itself.
