---
name: travel-planner
description: "Plan, build, or review travel itineraries with judgment about pace, routing, and meaning. Use when designing a trip from a vague brief, sequencing multi-city travel, choosing destinations and bases, balancing budget against experience, deciding what to book now vs leave flexible, selecting accommodations, or critiquing an existing itinerary for pacing, routing, and trip quality."
argument-hint: "Goal (plan/review/adapt), traveler purpose, dates or season, budget, party size and constraints, must-do experiences, and any existing itinerary"
---

# Travel Itinerary Planning

## What This Skill Does

Turn vague travel ambitions into trips that are coherent, well-paced, and worth taking.

This skill treats itinerary design as a constraint problem with one controlling rule: design for the traveler's purpose, then route, pace, and book around it. Coverage is not the goal. Depth, rhythm, and a trip the traveler will actually enjoy are.

Use this skill to:

- Build an itinerary from a vague trip brief.
- Sequence a multi-city or multi-region trip with realistic transit and recovery.
- Choose between destinations, bases, and pacing options.
- Decide what to book now vs leave flexible.
- Review an existing itinerary for pacing problems, routing waste, and traveler-fit mismatches.
- Adapt a generic "top 10" plan into something specific to the actual traveler.

This skill is for trip design and itinerary craft. It is not for visa, insurance, medical, or legal advice; it is not a real-time booking engine.

## When to Use

- The user has a destination, season, or trip idea and wants help shaping the plan.
- A trip already has rough bookings and needs review for pace, routing, or accommodation fit.
- The traveler is over-scheduling, under-budgeting, or trying to cover too much in too little time.
- The brief is vague ("ten days in Italy in April, mid-budget, two adults") and needs the right structural questions before any specific recommendations.

Do not use this skill for:

- Real-time price comparison or live booking.
- Visa, immigration, vaccination, or insurance advice.
- Detailed hour-by-hour single-day tours where general orientation is what the traveler actually needs.

## Inputs To Gather

Ask for whatever is missing before recommending a plan.

- `goal` — `plan`, `review`, `adapt`, or `decide`.
- `traveler_purpose` — what the trip is for: rest, immersion, family time, food, history, hiking, adventure, recovery, special occasion.
- `who` — number of travelers, ages, mobility, energy level, travel experience.
- `dates_or_season` — fixed dates, flexible window, or season only.
- `duration` — total days available door to door.
- `budget` — total range and where the traveler is willing to splurge or save.
- `pace_tolerance` — relaxed, moderate, ambitious; reaction to early starts and long transit days.
- `must_dos` — one or two non-negotiables.
- `constraints` — dietary, accessibility, language, climate sensitivity, no-fly preference, kids' nap windows.
- `existing_plan` — the current itinerary or bookings when reviewing.

If the brief is "we want to see X, Y, and Z in a week," ask whether reducing scope is acceptable before proposing a transit-heavy plan.

## Non-Negotiable Constraints

Treat these as defaults unless the user explicitly narrows scope.

1. Start with the traveler's purpose, not a destination list. "See everything" is not a useful brief.
2. Optimize for depth and rhythm, not maximum coverage. Fewer bases, longer stays, realistic transition days.
3. Route with geographic and temporal logic. Respect map shape, transit friction, opening days, seasonality. Eliminate backtracking when possible.
4. Protect flexibility where it matters. Book scarce, high-consequence items early; leave the rest open.
5. Choose lodging for location and integration, not luxury signaling.
6. Leave deliberate slack. Unstructured time is where memorable travel happens.
7. Use budget as a design constraint, not a reason to shrink the trip into meaninglessness. Spend where it protects experience or removes major friction.
8. Prepare the traveler, not just the bookings. Context, season, basic language, packing discipline, and contingency thinking matter.

## Procedure

Follow the steps in order. Do not skip the brief-clarification step.

### Step 1 — Clarify the trip brief

Determine:

- the traveler's purpose and what they want to feel, understand, or recover.
- who is going and how their pace, mobility, and interests differ.
- the time window, duration, and budget range.
- one or two must-dos and any hard constraints.

If any of these are missing and would change the plan materially, ask before recommending.

### Step 2 — Set the trip shape before details

Decide the structure before any specific recommendations:

- How many places.
- How many nights per base.
- Single base, hub-and-spoke, or loop.
- Travel-day count and tolerance for transit.

Push back on plans that compress too much. Three nights is usually the floor for "feeling" a place; one-night stops chain transit days unless there is a strong reason.

### Step 3 — Design the route realistically

Sketch the route on a real map, not a wish list:

- Geographic coherence: avoid bouncing back and forth.
- Transit reality: real travel times, including buffer for missed connections.
- Day-of-week and seasonality: what is closed, what is high or low season, what is monsoon or shoulder.
- Arrival and departure timing: do not waste day one on a red-eye and day five on an early checkout.

### Step 4 — Prioritize the non-negotiables

Identify the one or two experiences the traveler will remember most. Protect those. Then keep the rest of the plan intentionally lighter.

If two non-negotiables conflict (peak-season slot vs other plans, scarce reservation vs flexible day), surface the conflict and propose a sequencing solution rather than over-promising.

### Step 5 — Choose the right lodging strategy

Pick lodging by:

- Location and walkability over amenities.
- Neighborhood fit for the trip's purpose.
- Whether a guesthouse, apartment, or hotel best supports the rhythm.
- Whether longer stays in fewer places justify a small upgrade.

Avoid recommending properties without knowing the neighborhood's daily life and the trip's pace.

### Step 6 — Decide what to book now vs leave flexible

Split the plan into:

- Book now: scarce, high-consequence items (intercity transport in peak season, key restaurants, marquee experiences, hard-to-replace lodging in must-do destinations).
- Book later: most restaurants, day trips, secondary lodging when supply is plentiful.
- Leave open: weather-dependent activities, drift days, recovery time.

Make this split explicit so the traveler knows where urgency belongs.

### Step 7 — Stress-test the itinerary

Look for:

- Over-packed days and back-to-back early starts.
- Fragile connections that break the plan if anything slips.
- Lodging in the wrong neighborhood for the day's plan.
- Missing recovery time after long-haul flights or peak-energy days.
- Goals that do not match the proposed plan.

Rework the plan; do not just warn the traveler.

### Step 8 — Prepare the traveler

Round out the plan with:

- Seasonality notes and what to expect.
- Packing principles tied to climate, terrain, and trip type.
- Basic language or cultural context where it changes the experience.
- Contingency thinking: what to do if a key transit leg cancels.
- A short "gotchas" list for the destination or season.

## Branching Logic

- **Vague "see everything" brief.** Refuse to design a transit-heavy trip. Reframe around purpose first, then propose a narrower, deeper plan.
- **Fixed must-dos with limited time.** Build the trip around those anchors and accept reduced coverage elsewhere.
- **Family with young kids.** Optimize for shorter transits, predictable rhythm, kid-tolerant lodging, mid-day downtime.
- **Solo immersive trip.** Prefer fewer bases, longer stays, neighborhood lodging, more drift time.
- **Multi-generational group.** Plan for split tracks within the same base; accept shared lodging plus separate daily plans.
- **Peak-season trip.** Frontload bookings for scarce items; reduce coverage; warn explicitly about queues, prices, and crowd fatigue.
- **Shoulder or low-season trip.** Trade some closures for better pace and price; verify what is actually open.
- **Reviewing an existing itinerary.** Lead with the structural problems (pacing, routing, lodging fit) before commenting on individual restaurants or activities.
- **Budget pressure.** Cut scope or duration before cutting quality of the core experiences. Do not spread a thin plan thinner.

## Output Format

Match the deliverable to the task.

For `plan`:

1. Trip shape: number of bases, nights per base, route logic.
2. Day-by-day or block-by-block itinerary at the right altitude (avoid hour-by-hour unless requested).
3. Lodging recommendations with the reason each fits.
4. "Book now" vs "book later" vs "leave flexible" split.
5. Seasonal, transit, or fatigue risks called out.
6. Short prep list (packing principles, cultural notes, contingency thinking).
7. Assumptions used to drive the plan.

For `review`:

1. Top three structural problems (pacing, routing, lodging fit, mismatch with the brief).
2. Specific changes to fix each.
3. What still works and should be preserved.
4. Lower-priority refinements at the end.

For `adapt`:

1. The reframed plan tied to the traveler's actual purpose.
2. What was cut and why.
3. What was added or expanded and why.

## Quality Checks

The work is ready only when all of the following are true:

1. The plan answers the traveler's purpose, not a generic "things to see" list.
2. Pace matches the travelers' energy, mobility, and family rhythm.
3. The route is geographically and temporally coherent.
4. Lodging is justified by location and fit, not just amenities or brand.
5. Scarce items are flagged for early booking; the rest of the plan stays flexible.
6. Slack and recovery time are explicit, not accidental.
7. Seasonality, closures, and transit reality have been considered.
8. Assumptions and trade-offs are surfaced where they matter.

## Failure Modes To Avoid

- Optimizing for destination count, social-media optics, or bragging rights at the expense of trip quality.
- Recommending one-night-stop chains or overly ambitious transit days without a strong reason.
- Treating every booking with the same urgency.
- Ignoring seasonality, closures, transit reality, walking load, jet lag, or recovery time.
- Recommending accommodations purely on amenities when location matters more.
- Filling every day morning to night with no slack.
- Producing a "top 10" plan rather than one shaped to the actual traveler.

## Default Stance

Plan the trip the traveler will remember in five years, not the one they will brag about on day three. Default to fewer places, better pacing, and geographic coherence. Surface the trade-offs clearly so the traveler can choose with eyes open. When the brief is underspecified, ask only the minimum clarifying questions needed to avoid planning the wrong trip.
