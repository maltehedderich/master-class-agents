---
name: motion-design-review
description: "Review short SaaS and app marketing motion pieces (finished films, rough cuts, animatics, storyboards, scripts, or written concepts for 6–30 second social ads, product launch clips, and landing-page loops) against a rubric for promise, opening, frame readability, product proof, attention, rhythm, motion personality and sound, and ending. Returns a scorecard and prioritized improvements that fix comprehension before polish. Use when critiquing a product video or ad cut, pressure-testing a motion concept before production, comparing opening variants, or checking a series for a consistent motion style."
argument-hint: "The material (video file, frames, storyboard, script, or concept), placement and aspect ratio, target viewer, the one capability the clip shows, the intended response (CTA or loop), production stage, and any brand motion rules or retention data"
---

# Motion Design Review

## Purpose

Judge whether a short product clip makes one valuable transformation easy to understand and ties it to a memorable product, then say what to change first.

Review comprehension before craft. A polished clip whose result vanishes before it can be read fails. A rough animatic with a clear promise, visible product proof, and one point of attention per beat is ready for polish. Every recommendation should move the clip closer to a first-time viewer answering three questions with the sound off:

1. What does it do?
2. What changed?
3. What would I do next?

The rubric covers 6–30 second SaaS and app marketing clips (social ads, product launch clips, landing-page loops) and the concepts, scripts, storyboards, and animatics that lead to them. It is not a spec for UI micro-interactions, a critique of brand films that show no product, or technical render QC beyond what hurts comprehension.

## Inputs

Gather what the material does not already reveal:

- `material` — video file, exported frames, storyboard, animatic, script or shot list, or written concept.
- `stage` — `concept`, `storyboard`, `animatic`, `rough cut`, or `final`. Stage decides which fixes are cheap.
- `placement` — vertical social feed, in-stream ad, launch post, or landing-page loop, with aspect ratio and duration.
- `viewer` — who should care and what they already know about the product.
- `transformation` — the intended promise: who goes from which problem to which outcome using which capability.
- `intended_response` — the one next action for an ad, or the job the surrounding page does for a loop.
- `sound_context` — whether it mostly plays muted, plus any VO script, music, or sound-effect notes.
- `brand_motion` — existing motion principles, earlier clips in the series, type and color rules.
- `product_truth` — what the product actually does and how long the real action takes, so the demo can be checked for honesty.
- `evidence` — retention graphs, replay spikes, comprehension-test notes, conversion results, or variant comparisons.

Infer placement from aspect ratio and context when unstated (9:16 suggests a vertical feed; a silent 16:9 loop embedded on a page suggests a landing-page loop) and state the assumption. Ask one targeted question only when the answer would change the priorities, most often the intended transformation or the placement. Never hold a review back for brand guidelines or analytics.

## Non-Negotiable Constraints

1. Fix comprehension before polish. Never rank easing, transitions, or sound design above an unclear promise, an unreadable frame, or a result that disappears too fast.
2. Judge as a first-time viewer at real viewing size with the sound off. The maker's familiarity with the product is not evidence of clarity.
3. One transformation per clip. Treat extra features as a cost unless they serve the same transformation, and suggest saving them for other clips.
4. Remove before accelerating. When a beat feels rushed, cut information before speeding everything up.
5. Ground every finding in a timestamp, frame, or script line. "The pacing feels off" is not a finding. "The task list appears at 9.2 s and cuts at 9.8 s, before its four rows can be read" is.
6. Make every fix executable: what to cut, move, hold, enlarge, crop, or reorder, and roughly by how much.
7. Protect product honesty. Flag edits that imply nonexistent capabilities or misrepresent the work or waiting the real product requires. When product truth is unknown, list these as checks for the user to confirm.
8. Say what you could not assess. Sampled frames cannot show easing, a script cannot show legibility, and a muted file cannot show sound. Mark those criteria `n/a` instead of guessing.
9. Never invent evidence: no fabricated retention numbers, viewer quotes, test results, or audio observations. Label your first-viewing check as an expert simulation and recommend a real one.
10. Keep the improvement list short and ordered. Twenty equal-weight notes is a failed review.

## Workflow

### 1. Get the material into reviewable form

- **Video file:** read [Footage intake](references/footage-intake.md) and run its probe, frame-extraction, cut, hold, and silence commands. Review the contact sheet first, then key frames at full and viewing size.
- **Frames, storyboard, or animatic stills:** put them in order, attach intended timings, and note beats that have no frame.
- **Script, shot list, or written concept:** break it into beats with intended durations and describe what each key frame would show.
- **Link only:** if you cannot open it, ask for the file or exported frames. Never review from a title, thumbnail, or description.

Build a beat map before judging anything:

| Time | Beat | On screen | On-screen text | Primary focus | Audio / VO |
| --- | --- | --- | --- | --- | --- |

### 2. Reconstruct the promise

Write the transformation the clip actually communicates, using only the material:

> "[Specific user] goes from [recognizable problem] to [valuable outcome] using [one product capability]."

Compare it with the intended transformation. A gap here (no recognizable user, a vague outcome, three capabilities competing) is usually the root cause of findings further down. Name it first.

### 3. Run the first-viewing check

Walk the beat map with the sound off and answer the three viewer questions as the clip answers them, with the timestamp where each answer first becomes clear:

- **What does it do?** The answer and when it lands, or `unclear`.
- **What changed?** The before and after states the viewer actually saw.
- **What would I do next?** The action the ending invites, or `none`.

Then freeze the first frame: is there a clear subject and a credible reason for the intended viewer to care?

This check finds problems. It does not replace showing the clip to someone unfamiliar with the product.

### 4. Score the rubric

Score each criterion 0–3 and cite evidence for every score. Use `n/a` with a reason when the material cannot support a judgment.

| Score | Meaning |
| --- | --- |
| 0 | Missing, or working against the clip's goal |
| 1 | Present but unclear; a first-time viewer would likely miss it |
| 2 | Works, with friction that slows comprehension or weakens impact |
| 3 | Clear and deliberate; hard to improve without changing the concept |

#### R1. One transformation

**Question:** Is the clip built around one specific, valuable change for a recognizable viewer?

- **Look for:** an obvious user, a recognizable problem, a visible outcome, and one capability doing the work. Every screen serves that story.
- **Strong (3):** the transformation sentence can be written from the clip alone and matches the intent.
- **Weak (0–1):** a feature montage with many screens and no outcome; several capabilities competing; an outcome claimed in text but never shown.

#### R2. Opening relevance

**Question:** Do the first frame and the opening beat give the intended viewer a reason to keep watching?

- **Look for:** a recognizable problem, a desirable result, or a product action whose benefit is immediately apparent. Meaningful action starts promptly. Framing is tight. Brand identity comes through the product, typography, color, or a visible mark.
- **Strong (3):** the frozen first frame has a clear subject and a credible reason to care.
- **Weak (0–1):** a logo reveal, abstract build-up, empty background, or generic headline fills the opening.

#### R3. Frame readability

**Question:** Can the viewer find what matters in each key frame at actual viewing size?

- **Look for:** clear hierarchy between headline, product detail, and supporting material; essential text with enough size, contrast, and space; low information density; composition designed for the target aspect ratio; essential content clear of platform overlays.
- **Strong (3):** every key frame reads at phone size with one obvious focal point.
- **Weak (0–1):** a full-screen UI shrunk until it is illegible; text over busy interface; captions or CTAs under platform UI; a landscape edit letterboxed into a vertical feed.

For safe zones, use the platform's current overlay template rather than remembered percentages. As a default, keep essential text and product detail out of the bottom band and the right-hand column where captions, CTAs, and action buttons sit.

#### R4. Product proof

**Question:** Does the interface show believable evidence that the promise is real?

- **Look for:** a legible sequence of input → action → changed state → useful result; a crop to the relevant area; enlarged essential controls; believable example content; enough context to show where the action happens; a result that is visibly useful.
- **Strong (3):** the viewer sees cause and effect, and the result is something they would want.
- **Weak (0–1):** UI used as decoration; placeholder content like "Project 1"; a result that never appears or flashes past; edits that imply capabilities, speed, or effortlessness the product does not deliver.

#### R5. Attention choreography

**Question:** Does every beat have one primary point of attention, and does attention carry across beats?

- **Look for:** statement → interaction → result sequencing. Several elements can move together when they express one coherent action. Continuity carries across cuts through a shared object, position, or direction, for example a selected note becoming the task card in the next view.
- **Strong (3):** you could draw a single gaze path through the whole clip.
- **Weak (0–1):** a headline animates in while a cursor performs a detailed action; a camera move, a background effect, and a UI change all happen at once; every cut resets the eye to a new corner.

#### R6. Rhythm and holds

**Question:** Does movement get efficiently between ideas and let important moments settle long enough to register?

- **Look for:** demonstration and result getting most of the running time; key words and results held long enough for a first-time read; easing that makes entrances, changes, and arrivals feel intentional; timing that scales with distance, visual weight, and complexity; anticipation, overshoot, and bounce saved for meaningful moments.
- **Strong (3):** nothing important disappears before it is understood, and nothing lingers without purpose.
- **Weak (0–1):** the result cuts away before it can be read; everything moves at one speed; uniform linear or uniformly bouncy motion; a long build-up and a rushed payoff.

There is no universal hold duration. Hold test: read the on-screen text aloud at a normal pace while the hold plays, then allow a beat to take in the change. If the frame is gone first, the hold is too short.

#### R7. Motion personality and sound

**Question:** Does the motion follow a consistent vocabulary that expresses the brand, and does sound reinforce the message without carrying it alone?

- **Look for:** consistent rules for how elements enter, transform, settle, and leave; easing and movement size that match the brand (a calm productivity tool might use short movements and measured settling, a playful creative app more elastic reactions); music for rhythm and selective effects on actions and results; VO and on-screen text saying the same thing; the essential claim understandable muted.
- **Strong (3):** the clip would still feel like this brand with the logo removed, and it works muted.
- **Weak (0–1):** preset-pack transitions; each scene moves differently; VO and on-screen text compete; the key claim exists only in the VO.

#### R8. Ending and intended response

**Question:** Does the ending do what the placement needs?

- **Ad or launch clip:** the demonstrated benefit connects to one clear next action, with the useful result still visible beside the product name and CTA.
- **Landing-page loop:** the useful result settles and the last frame returns smoothly to the first. The page carries the CTA.
- **Strong (3):** the viewer knows exactly what to do next, or the loop seam is invisible.
- **Weak (0–1):** a generic end card replaces the result; several CTAs compete; the loop visibly jumps; the ending introduces a new feature.

### 5. Prioritize improvements

Write each finding as: where (timestamp or frame), criterion, observed problem, which viewer question it blocks or slows, the concrete fix, and effort (`S`, `M`, or `L` at the current stage). Merge findings that share a root cause into one improvement.

Assign a priority:

- **P0 — Fix before anything else.** A first-time viewer cannot answer one of the three questions; the opening gives no reason to watch; essential text or product detail is unreadable or hidden by platform UI; or the demo misrepresents the product.
- **P1 — Fix in this revision.** The answers exist but arrive late, compete for attention, or do not hold long enough; the ending is weak or split; muted comprehension depends on the VO.
- **P2 — Polish once comprehension is confirmed.** Easing refinement, motion-vocabulary consistency, transition continuity, sound design, brand expression.

Within a tier, rank by leverage:

1. Root causes before symptoms. Sharpening the transformation can dissolve three downstream findings.
2. Earlier before later. Viewers who leave in the opening never see a fix at second 12.
3. Cheap at the current stage before expensive.

Adjust for stage:

- **Concept or storyboard:** structural change is cheap. Push hard on the transformation, opening, proof plan, and key frames. Leave easing for later.
- **Animatic or rough cut:** the moment to fix composition, attention, and holds, before anyone polishes.
- **Final or locked cut:** weight effort heavily. Prefer edit-level fixes (trim, re-time, extend a hold, reframe, reorder, swap the end card, add captions) and carry deeper problems into the next clip or a variant. Still name an expensive P0 honestly.

Keep the list actionable: every P0, the strongest P1s (usually no more than five), and at most three P2s. If any P0 is open, say plainly that polish should wait until a revised cut passes a comprehension check.

### 6. Propose a revised structure when needed

When fixes reorder beats, cut a feature, or reallocate time, give a revised beat sheet:

| Time | Viewer takeaway | Visual action |
| --- | --- | --- |

Give the demonstration and its result most of the running time, and label the timings as starting points to test. For concepts and storyboards, also define the four key frames (opening, action, result, ending): what each shows and the one thing the viewer should notice in it.

### 7. Plan validation

Close with the smallest test that would confirm the top fixes:

- Show the revised cut once, sound off, to someone unfamiliar with the product. Ask what it does, what changed, and what they would do next.
- If the opening is the main doubt, cut two opening variants and hold the demonstration and ending constant.
- Once the clip is live, use the retention graph to find drop points and replays, and pair it with the business outcome (for example qualified trial starts) and a comprehension check. A replay alone does not show the message was understood.

## Branching Logic

- **Written concept or one-liner.** Score R1, R2, R4, and R8 from intent and mark the rest `n/a`. Deliver a sharpened transformation sentence, four key frames, a beat sheet, and the risks to check in the animatic.
- **Script or VO-led concept.** Check whether the claim survives without the VO. List the lines that need an on-screen equivalent and the moments that ask the viewer to read while following an interaction.
- **Storyboard or key frames.** Weight R1–R5 and R8. Review frames at viewing size. Judge R6 from intended timings and R7 from motion notes only.
- **Animatic or rough cut.** Full rubric. Holds and attention matter most here. Recommend holding polish until P0s and P1s are resolved.
- **Final or locked cut.** Full rubric. Separate "fix in this cut" from "carry into the next clip."
- **Vertical social ad.** Raise the bar on the first frame, safe zones, muted comprehension, and a single CTA. Composition should be native to the vertical frame, not a reframed landscape edit.
- **Landing-page loop.** Assume muted autoplay and drop the in-clip CTA requirement. Compare the first and last frames for the loop seam, check that the result settles before the restart, and judge legibility at the embed's rendered width.
- **Opening variants.** Review each opening against R2 and R1 with the rest held constant. Say which should establish relevance faster and why, and recommend the test rather than declaring a winner.
- **Series or motion system.** Compare the clips' vocabulary for enter, transform, settle, and leave. List inconsistencies and propose a compact motion-vocabulary table.
- **Retention data provided.** Map drops and replays onto the beat map. Treat a drop as a place to investigate, not a diagnosis.
- **Product truth unknown.** Review the demo as presented and list honesty checks for the user: implied speed, skipped steps, capabilities shown.
- **Longer than about 30 seconds.** Apply the rubric per section and say whether a 6–15 second cutdown around one transformation would serve the placement better.

## Output Format

For a film, cut, animatic, or storyboard:

1. **Reviewed** — material, stage, placement, aspect ratio, duration; which criteria could and could not be assessed; assumptions.
2. **Verdict** — one of the following, with one or two sentences on why:
   - `Rethink the concept` — no single transformation; edits will not fix it.
   - `Fix comprehension first` — at least one P0.
   - `Tighten before polish` — no P0, but P1s slow comprehension.
   - `Polish and test` — comprehension holds; remaining work is craft and validation.
3. **Promise check** — the transformation as the clip communicates it, beside the intended one.
4. **First-viewing check** — the three questions with answers and timestamps, plus the first-frame verdict, labeled as a simulation.
5. **Keep** — up to three strengths the revision must not lose.
6. **Scorecard** — `criterion | score | evidence` for R1–R8.
7. **Prioritized improvements** — numbered, P0 first, each in this shape:

   `1. P0 · 0:00–0:02 · R2 Opening` — What is wrong, as observed. **Fix:** the concrete change. **Unblocks:** which viewer question or criterion. **Effort:** S.

8. **Revised beat sheet** — only when the structure changes.
9. **Validation** — the next test and the result that would confirm the fixes.

For a concept or script:

1. **Reviewed** and assumptions.
2. **Verdict.**
3. **Transformation** — as written, then sharpened.
4. **Scorecard** for the assessable criteria.
5. **Prioritized improvements.**
6. **Four key frames and beat sheet.**
7. **Check in the animatic** — the risks a concept cannot settle, such as legibility at size or whether the result holds long enough.

For opening variants: a comparison table of `variant | first-frame subject | reason to care | time to first meaningful action | risk`, then the recommended test.

For a series: a one-line verdict per clip, the motion vocabulary observed (or its absence), inconsistencies, a proposed vocabulary table of `action | easing and character | typical duration and distance | used for` covering enter, transform, settle, and leave, and the top improvements across the series.

## Quality Checks

Treat the review as ready only when all applicable checks pass:

1. The transformation is written down both as communicated and as intended.
2. The three viewer questions are answered with timestamps or marked `unclear`.
3. Every score cites a timestamp, frame, or script line; unassessable criteria are `n/a` with a reason.
4. No P2 outranks an open P0 or P1.
5. Every improvement names a location and an executable fix.
6. Pacing fixes remove information before speeding things up.
7. Recommendations fit the production stage and the placement.
8. Product-honesty concerns are raised, or their absence is stated.
9. The list is short enough to act on in one revision.
10. Strengths worth keeping are named.
11. The review ends with a concrete validation step.
12. Nothing is invented: no fabricated data, quotes, or claims about motion or audio you did not observe.

## Failure Modes

- Leading with easing, transitions, or color while the result disappears before it can be read.
- Reviewing as someone who already knows the product instead of as a muted first-time viewer at phone size.
- Rewarding a feature montage because every screen looks polished.
- Prescribing "speed it up" for a rushed sequence instead of removing information.
- Notes like "make it pop", "pacing feels off", or "more dynamic" with no timestamp or fix.
- Judging easing from sampled frames, or sound from a muted file, without saying so.
- Treating a logo reveal as a required opening.
- Asking for a rebuild of a locked final when a re-time, crop, hold, or caption would solve the problem.
- Asking for more features, text, or scenes in a clip whose problem is overload.
- Accepting demo edits that imply capabilities or speed the product does not have.
- Treating a replay spike or strong retention as proof the message landed.
- Delivering a long list of equal-weight notes.

## Default Stance

Clarity before craft, proof before polish, one transformation per clip. Watch the clip the way a stranger in a muted feed would, find the earliest moment comprehension breaks, and fix that first, usually by removing something. Then confirm it with a real first-time viewer.
