# Apple Ads (App Store): Keyword Selection

**Choose keywords by the intent your app can satisfy—not simply by their search popularity.** A useful keyword strategy connects what someone searches for, what your app actually delivers, and what you can afford to pay for that customer.

This guide focuses on **search-results campaigns using Manage Bids**, where you control keywords and bids. Apple also offers **Maximize Conversions**, which uses automated bidding and Search Match to pursue tap-through installs at a target acquisition cost; keyword-level bids are not available in that strategy. ([Apple Ads][1])

The examples below use a hypothetical focus app with Pomodoro sessions, study-session history, and daily focus goals—but **no app blocking, calendar planning, or habit tracking**. The example keywords are research candidates, not verified search-volume recommendations.

## 1. Start with the user’s job—not your app’s category

“Productivity” describes a category. “Study timer” describes something a person wants to do.

For keyword selection, start by completing this sentence:

> Someone searching for **[keyword]** wants **[outcome]**, and our app delivers that through **[specific capability]**.

For the example app:

> Someone searching for **“study timer”** wants to structure study sessions, and our app delivers that through timed focus sessions and study history.

That is a stronger starting hypothesis than:

> Someone searching for **“productivity”** might like our app because it is a productivity app.

Apple considers both relevance and bids when deciding whether an ad can appear. Its documentation explicitly states that an irrelevant app will not be displayed merely because the advertiser is willing to pay more. ([Apple Ads][2])

### Build an intent map

| Keyword family       | Example candidates                              | Selection logic                                                                      |
| -------------------- | ----------------------------------------------- | ------------------------------------------------------------------------------------ |
| Your brand           | Your app name; company name where relevant      | Separate people already seeking your product from people discovering alternatives.   |
| Core function        | `pomodoro timer`, `focus timer`                 | Start here when the phrase directly describes a strong capability.                   |
| Specific use case    | `study timer`, `revision timer`                 | Test when the app genuinely supports that situation.                                 |
| Desired outcome      | `improve concentration`, `stop procrastinating` | Treat as hypotheses: the desired outcome may be broader than the solution you offer. |
| Broad category       | `productivity`, `time management`               | Test cautiously rather than assuming category membership establishes strong intent.  |
| Competitor brand     | Names of directly comparable timer apps         | Evaluate as a separate alternative-selection experiment.                             |
| Unsupported function | `app blocker`, `habit tracker`                  | Exclude when the app cannot deliver the expected function.                           |

For this hypothetical app, I would prioritize core-function and study-use-case terms before broad productivity terms.

**The critical test:** Would a reasonable person searching this phrase feel that your app answered their request?

A keyword should not survive that test merely because it has attractive volume.

## 2. Build your candidate list from evidence

Use several inputs, but keep **candidate generation** separate from **validation**.

### Your product and customer language

Start with your actual feature set, onboarding choices, support messages, reviews, and customer interviews. Extract the language people use to describe the problem—not just the terminology your team uses internally.

For example, “session-based attention management” might be your product language. “Focus timer” is a candidate worth investigating as search language.

Record why each candidate belongs on the list. “A customer used this phrase” and “we brainstormed this phrase” are different levels of evidence.

### Apple’s keyword suggestions and recommendations

Apple provides keyword suggestions when adding keywords, including related suggestions based on terms you enter. Its Recommendations page also provides keyword opportunities with estimated installs, spend, acquisition cost, and search popularity. These estimates are not guarantees. ([Apple Ads][3])

Use those recommendations to expand and investigate your list—not to bypass the relevance test.

A useful research record includes:

**Keyword, country or region, language, intent family, source, product fit, demand signal, proposed test, and current status.**

### Manual App Store investigation

Search your candidate phrases in the intended storefront and inspect the results.

Ask: Are these mostly timers, planners, blockers, or something else? Does the user appear to be looking for a specific brand? Would your screenshots make sense beside the apps being shown?

Treat this as a qualitative intent check, not a search-volume measurement or a complete view of the advertising auction.

### Search Match and broad-match discovery

Search Match can match your ad to searches without an explicit keyword list. Apple says it uses listing metadata, information about similar apps, and other search data. This makes it a useful discovery input, provided you inspect what it finds. ([Apple Ads][4])

Once advertising starts, add observed search terms to your research inventory. A phrase someone actually used to find your ad is more informative than another untested brainstorm.

### Research each market separately

For each country or region, repeat the intent and demand checks rather than treating translation as validation.

For example, investigate whether people use a local-language category term, an English term, a compound word, or a different description of the task. Match that research with localized product-page assets. Apple recommends localizing keywords and product-page content for the markets you serve. ([Apple Developer][5])

**A translated phrase is a candidate. It is not automatically a validated local keyword.**

## 3. Prioritize relevance first, then demand and economics

Avoid a scoring system in which enormous popularity can compensate for a product mismatch.

I recommend a two-stage selection process.

### Stage A: Apply a relevance gate

Reject or hold a candidate when it implies an unsupported feature, an audience you do not serve, or a promise the product cannot fulfill.

For the example app, `app blocker` fails immediately. There is no reason to spend money discovering whether people who want blocking will accept a timer instead.

By contrast, `study planner` deserves a deliberate decision. Daily focus goals do not necessarily satisfy someone expecting timetables, assignments, and calendar planning.

### Stage B: Rank the candidates that pass

Evaluate the remaining terms against four questions:

| Factor                | Question to answer                                                                 |
| --------------------- | ---------------------------------------------------------------------------------- |
| Intent clarity        | How precisely does the phrase describe a job the app performs well?                |
| Demand evidence       | What evidence suggests people use this phrase in the target storefront?            |
| Conversion readiness  | Can the ad, product page, and first-use experience demonstrate the promised value? |
| Economic plausibility | Could the acquisition cost work given the value of the resulting customers?        |

Keep an additional **confidence** label: brainstormed, supported by research, tested, or supported by mature performance data.

That prevents an attractive assumption from being treated as a proven winner.

### Interpret search popularity correctly

Apple currently documents Search Popularity as a **1–5 measure**, with 5 representing the most popular terms, in its reporting and recommendations. It is a popularity indicator—not an exact monthly-search count. ([Apple Ads][6])

Use it to compare opportunities, not to calculate precise market size. Nor should a low score automatically disqualify a highly relevant term.

Apple’s own guidance recommends balancing general and specific keywords: specific terms may improve tap-to-install conversion, but extremely narrow terms can limit reach. ([Apple Ads][7])

### Make the active list fit the learning budget

For a modest initial test, I would consider roughly **10–20 strongly relevant exact-match candidates**, grouped around a few clear intents, before activating a much larger inventory. This is a planning heuristic, not an Apple requirement.

Consider an illustrative €300 test at €1 average cost per tap:

- Across 60 keywords, that averages only five taps per keyword.
- Across 10 keywords, it averages 30 taps per keyword.

Actual delivery will be uneven, but the trade-off is clear: adding more keywords does not necessarily produce more useful learning.

Maintain a broad research inventory. Activate a narrower set you can afford to evaluate.

## 4. Understand what a keyword actually controls

A **keyword** is the targeting input you choose. A **search term** is what the person actually typed. Apple exposes keyword and search-term performance separately in search-results reporting. ([Apple Ads][8])

That distinction matters because one keyword can match multiple searches.

### Exact match

Exact match is the tighter targeting option, but it is **not literal-only matching**. Apple says it may include spelling variations, word rearrangements, translations, and other close variants. ([Apple Ads][9])

### Broad match

Broad match can match relevant variants and related terms. Apple offers these two positive keyword match types; there is no separate phrase-match option. ([Apple Ads][9])

### Search Match

Search Match is automatic query matching—not a third keyword match type. It can find searches without you entering those phrases as keywords. ([Apple Ads][4])

Two practical implications follow.

First, **a generic keyword is not the same thing as broad match**. Setting `productivity` to exact match does not make its underlying intent precise.

Second, review actual search terms even when using exact match. The entered keyword alone does not fully describe the traffic you purchased.

One setup detail is easy to miss: newly added keywords default to broad match. After a keyword is saved, changing its match type requires pausing it and adding a new keyword with the desired type. ([Apple Ads][3])

## 5. Separate discovery from the keywords you want to control

Apple recommends separating brand, category, competitor, and discovery activity. Its suggested discovery structure uses separate broad-match and Search Match ad groups. ([Apple Ads][10])

A practical implementation is:

| Campaign or ad-group purpose | Targeting                                      | Search Match | What you are trying to learn                                   |
| ---------------------------- | ---------------------------------------------- | -----------: | -------------------------------------------------------------- |
| Brand                        | Exact-match brand terms                        |          Off | What happens when people already seek your app?                |
| Category / core intent       | Exact-match function and use-case terms        |          Off | Which nonbrand needs produce valuable users?                   |
| Competitor                   | Exact-match directly comparable app names      |          Off | Can your product win users considering a specific alternative? |
| Discovery: broad             | A controlled set of relevant broad-match seeds |          Off | Which related searches deserve their own test?                 |
| Discovery: Search Match      | No explicit keywords required                  |           On | Which relevant searches have you missed?                       |

The exact-match/Search Match settings above follow Apple’s campaign-structure guidance. ([Apple Ads][11])

You do not need to fund every category immediately. For a new app with little brand demand and a small budget, I would prioritize core intent and a controlled discovery test before an extensive competitor campaign.

Separate campaigns are useful when you need separate budgets. Ad groups can provide thematic separation within a campaign, but the budget is set at campaign level. ([Apple Ads][11])

### Promote discoveries into controlled tests

When a discovery search term shows promising relevance and results:

1. Add it as an exact-match keyword in the appropriate core campaign or ad group.
2. Give it an explicit bid and performance hypothesis.
3. Add that search term as an exact-match negative in discovery.

This is the discovery-to-exact process Apple recommends. ([Apple Ads][11])

Promoting a term means you want more control and cleaner evaluation. It does **not** mean a handful of early installs has proved profitability.

### Keep brand performance separate

A person searching your app’s name may already have intended to download it. Therefore, an attributed branded-search install does not, by itself, demonstrate an additional customer caused by advertising.

For acquisition decisions, compare brand and nonbrand results separately. Otherwise, strong branded-search results can obscure an unprofitable nonbrand strategy.

## 6. Use negative keywords deliberately

Negative keywords serve two different purposes: **excluding unsuitable intent** and **routing known searches away from discovery**. Apple allows negatives at campaign or ad-group level. ([Apple Ads][12])

For the example app, unsuitable-intent candidates might include app-blocking or habit-tracking searches. Routing negatives might include a proven `study timer` term now managed in a dedicated exact-match group.

### Negative matching is not symmetrical with positive matching

An **exact negative** blocks the specified term, not its close variants or reordered versions. A **broad negative** requires all its words to be present, in any order; variants are not automatically covered. ([Apple Ads][9])

Consequently, an exact negative for `app blocker` does not constitute a complete exclusion of every blocking-related search.

Start narrowly when the unwanted intent is uncertain. Use broader exclusions only when you genuinely want to reject the whole represented intent.

Also place routing negatives at the correct level. **When discovery and core targeting share one campaign, put routing negatives on the discovery ad groups—not on the entire campaign.** Otherwise, you can exclude the searches you intended the core groups to receive.

Finally, do not automatically exclude words such as `free`. Decide whether the resulting expectation fits your actual free tier and whether those users convert economically.

## 7. Decide what a keyword is worth before chasing volume

The economically important question is not:

> How cheaply can this keyword generate installs?

It is:

> How much can we pay for the users this keyword brings?

Apple distinguishes tap-through and total acquisition metrics. Its install metrics also include redownloads, while new-download metrics are available separately. Keep those definitions consistent when comparing performance. ([Apple Ads][8])

For a subscription app, a useful planning relationship is:

$$
\text{Affordable average cost per tap}
=
\text{Target cost per paying customer}
\times
P(\text{install}\mid\text{tap})
\times
P(\text{payer}\mid\text{install})
$$

### Illustrative calculation

Suppose your target cost per paying customer is **€40**, your expected tap-to-install rate is **50%**, and your expected install-to-paying-customer rate is **10%**.

Then:

$$
€40 \times 0.50 \times 0.10 = €2.00
$$

Under those assumptions, €2 is the economically tolerable **average** cost per tap. Applying a 25% planning cushion gives €1.50.

This is not a guaranteed bid recommendation. Conversion rates can change, assumptions may be wrong, and the max CPT bid is a ceiling rather than the actual average price paid. Apple notes that actual tap costs can be below the maximum bid. ([Apple Ads][13])

Choose your customer-acquisition target from expected contribution margin, payback requirements, and uncertainty—not simply the headline subscription price.

### Why the cheapest install can be the worse purchase

The following is a **synthetic example**, using comparable, mature new-user cohorts:

| Keyword        | Spend | Taps | New installs | Paying customers | Cost per install | Cost per payer |
| -------------- | ----: | ---: | -----------: | ---------------: | ---------------: | -------------: |
| `focus timer`  |  €120 |  100 |           60 |                6 |            €2.00 |         €20.00 |
| `productivity` |   €80 |  100 |           50 |                1 |            €1.60 |         €80.00 |
| `study timer`  |  €100 |  100 |           40 |                8 |            €2.50 |         €12.50 |

`productivity` has the cheapest installs but the most expensive paying customers.

`study timer` has the most expensive installs but the best payer economics.

For this hypothetical subscription business, I would investigate expanding `study timer` before celebrating the lower install cost of `productivity`.

Use downstream keyword data where your attribution setup supports it. Otherwise, evaluate at the ad-group or campaign-cohort level rather than inventing search-term-level revenue precision.

## 8. Match the keyword to the product page and first-use experience

Treat the test unit as:

**Search intent → ad creative → product page → first useful action**

For `study timer`, the example app should demonstrate studying, timed sessions, and study history.

For `pomodoro timer`, it should demonstrate work intervals, breaks, and session controls.

Apple supports search-results ad variations based on custom product pages and recommends aligning those assets with keyword themes. ([Apple Ads][14])

You do not need a different page for every keyword. I would group terms around a shared promise: one page for study sessions, another for general Pomodoro use, where those differences are meaningful.

Before rejecting a relevant keyword, ask whether you gave it a fair presentation. Someone seeking a study timer should not have to infer that capability from screenshots dominated by unrelated features.

However, better creative cannot repair a nonexistent feature. Do not use a product page to imply app blocking when the product only provides a timer.

**Paid targeting and App Store metadata also have different rules.** Apple’s product-page guidance prohibits competing app names in the App Store keyword metadata field; do not copy a competitor advertising list into that field. ([Apple Developer][5])

## 9. Evaluate evidence—not just elapsed time

“Run every keyword for seven days” is not a sufficient decision rule.

A keyword with three taps and one with 300 taps have produced very different amounts of evidence. A subscription cohort that has not finished its trial period has not yet revealed its full paid conversion.

Use a combination of **traffic, spend, conversion maturity, and uncertainty**.

### No impressions: investigate delivery first

Check relevance, bid competitiveness, budget, targeting, and negative-keyword conflicts before deciding there is no demand.

Apple specifically notes that relevant, popular keywords with few impressions may be losing to higher bids. Increasing a bid can be a diagnostic test—but only within your economic limits. ([Apple Ads][13])

### Impressions but few taps: investigate the initial promise

Review the actual search terms and what your ad communicates.

My diagnostic questions would be: Is the query ambiguous? Does the ad make the relevant function obvious? Is someone seeking a specific alternative that your presentation does not convincingly replace?

### Taps but few installs: investigate expectation mismatch

Inspect the product page, visible value proposition, pricing expectations, and feature fit. Do not assume the keyword itself is the only variable.

### Installs but weak activation or payment: investigate user quality and product experience

Determine whether the keyword attracts the wrong need, whether onboarding delays the promised benefit, or whether monetization expectations are misaligned.

These are diagnostic hypotheses—not conclusions that can be read directly from one conversion-rate number.

### Do not overreact to tiny samples

Suppose a keyword’s true tap-to-install conversion rate were 10%. Under a simplified model of independent taps with a constant conversion probability, the chance of observing zero installs after 10 taps would be:

$$
(1 - 0.10)^{10} \approx 35\%
$$

Zero installs after 10 taps therefore would not be especially surprising.

At the same time, you need not keep spending indefinitely to achieve statistical certainty. Set a test-spend limit in advance and distinguish **“not viable within this budget”** from **“proven incapable of converting.”**

When reviewing subscription performance, compare cohorts at similar ages and avoid treating users still in a trial as confirmed non-payers.

## 10. Maintain a repeatable selection loop

Keep four statuses in your keyword inventory:

**Candidate → Testing → Proven for a defined objective → Paused or excluded**

“Proven” should always include context: country, audience, creative, acquisition objective, cohort window, and observed economics.

A useful recurring review has three passes.

**Protect the test.** Check spend, delivery failures, irrelevant queries, and negative-keyword conflicts. Apple’s dashboards provide keyword and search-term views, filters, and downloadable performance data. ([Apple Ads][15])

**Make targeting decisions.** Promote promising discovery terms, narrow poor broad-match traffic, pause exhausted tests, and choose the next candidates.

**Reassess business value.** Review mature activation, payment, retention, and contribution results. Revisit earlier winners when the product, pricing, or acquisition economics change.

For each decision, write down what changed and why. That creates a record of learning rather than a history of unexplained bid adjustments.

---

## The practical starting point

For a new campaign, I would begin with **one clearly defined market, a small set of exact-match function and use-case terms, a controlled discovery allocation, and an explicit definition of a valuable customer**.

The central discipline is:

> **Use relevance to decide what deserves a test, search-term data to understand what you actually bought, and customer economics to decide what deserves more budget.**

The objective is not the largest keyword list. It is a growing set of search intents your app can satisfy profitably.

[1]: https://ads.apple.com/app-store/help/campaigns/0095-maximize-conversions "Use Maximize Conversions - Help - Apple Ads"
[2]: https://ads.apple.com/app-store/help/ad-placements/0082-search-results "Search Results - Help - Apple Ads"
[3]: https://ads.apple.com/app-store/help/keywords/0014-add-and-manage-keywords "Add and Manage Keywords - Help - Apple Ads"
[4]: https://ads.apple.com/app-store/help/campaigns/0006-understand-search-match "Understand Search Match - Help - Apple Ads"
[5]: https://developer.apple.com/app-store/product-page/ "Creating Your Product Page - App Store - Apple Developer"
[6]: https://ads.apple.com/app-store/help/recommendations/0055-review-recommendations "Review Recommendations - Help - Apple Ads"
[7]: https://ads.apple.com/app-store/best-practices/keywords "Keywords - Best Practices - Apple Ads"
[8]: https://ads.apple.com/app-store/help/reporting/0023-reporting-options-and-definitions "Reporting Options and Definitions - Help - Apple Ads"
[9]: https://ads.apple.com/app-store/help/keywords/0059-understand-keyword-match-types "Understand Match Types - Help - Apple Ads"
[10]: https://ads.apple.com/app-store/help/campaigns/0056-structure-campaigns "Structure Campaigns - Help - Apple Ads"
[11]: https://ads.apple.com/app-store/best-practices/campaign-structure "Campaign Structure - Best Practices - Apple Ads"
[12]: https://ads.apple.com/app-store/help/keywords/0060-use-negative-keywords "Use Negative Keywords - Help - Apple Ads"
[13]: https://ads.apple.com/app-store/best-practices/manual-bidding "Manual Bidding - Best Practices - Apple Ads"
[14]: https://ads.apple.com/app-store/best-practices/ad-variations "Ad Variations - Best Practices - Apple Ads"
[15]: https://ads.apple.com/app-store/help/reporting/0024-view-campaigns-dashboard-metrics "View Campaigns Dashboard Metrics - Help - Apple Ads"
