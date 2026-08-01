## Role Framing

UI writing is interface design performed with words: reducing uncertainty at the exact moment someone must decide or act, in the smallest possible space, in a voice the product can sustain. Excellence requires product judgment and user research literacy more than it requires prose talent.

## Best-in-Class Voices

- **Torrey Podmajersky** — _Strategic Writing for UX_; UX writing at Microsoft, Xbox, Google, OfferUp. Emphasizes UX content as a tool that must serve organizational _and_ user goals simultaneously, the "voice chart" as a scalable decision instrument, and measuring content's effect rather than defending it on taste.
- **Sarah Richards** — founded content design practice at GOV.UK; _Content Design_. Emphasizes evidence over opinion (search data, support logs), plain language as an accessibility requirement, and the idea that the job is answering a user need, not producing words.
- **Kinneret Yifrah** — _Microcopy: The Complete Guide_. Emphasizes the high-friction moments — forms, errors, empty states, waiting, confirmation — as where microcopy earns its keep, and writing that lowers anxiety rather than performing personality.
- **Nicole Fenton & Kate Kiefer Lee** — _Nicely Said_; Kiefer Lee built Mailchimp's content style guide and its voice-and-tone work. Emphasizes the distinction that voice stays constant while tone shifts with the user's emotional state.
- **Ginny Redish** — _Letting Go of the Words_. Emphasizes that people scan rather than read, content as conversation, and writing for the specific moment of need.
- **Steve Krug** — _Don't Make Me Think_. Emphasizes self-evident labels and aggressive omission — his standing advice is to cut half the words, then half again.

For day-to-day reference, the institutional style guides (Apple HIG, Material Design, GOV.UK, Mailchimp) encode most of the above in enforceable form.

## Core Idea

The words _are_ the interface — a person can only act on what they can read and believe. Every string should reduce uncertainty, name the next action, or build trust; anything else is cost paid in attention.

## Essential Best Practices

**1. Write the flow, not the string.**
_Why:_ Strings are read in sequence, but reviewed in isolation. Locally clever copy produces globally incoherent flows — the same concept named three ways across four screens.
_How:_ Before touching a mock, list every string in the flow in one document, including empty, loading, error, permission, success, and edge states. Read it aloud in order. Fix vocabulary drift first; polish second.

**2. Anchor every string to a decision the user is making.**
_Why:_ Most bad microcopy is copy that had no job. Without a decision to serve, writing defaults to describing the system.
_How:_ Ask "what does this person need to know to act right now?" Answer from evidence — support tickets, search queries, session recordings, usability sessions — not from a stakeholder's phrasing preference. If nothing answers the question, delete the string.

**3. Front-load meaning and make labels self-evident.**
_Why:_ People scan; the first one or two words carry the click. A vague label forces a re-read, and a re-read is a small failure.
_How:_ Buttons state the outcome, not the mechanism: _Save changes_, _Delete 3 files_, _Send invite_ — never _Submit_ or _OK_ alone. Put the noun the user is looking for at the front of headings and links. Avoid "Learn more" and "Click here" as standalone targets.

**4. Cut, then cut again.**
_Why:_ Density of meaning is the whole craft. Explanatory padding, greeting text, and restated instructions all dilute the string that matters.
_How:_ Delete each string entirely and see if the flow breaks. Kill introductory "happy talk," redundant help text, and any sentence explaining what the button beneath it already says.

**5. Separate voice from tone, and encode both.**
_Why:_ Voice is who the product is and shouldn't move; tone must move with the user's emotional state. Encoded, these scale to engineers and PMs writing strings you'll never review.
_How:_ Build a one-page voice chart with concrete columns — concepts, vocabulary (words we use / never use), verbosity, grammar, punctuation — with real before/after examples. Then map tone to states: neutral in errors, restrained in success, plainly serious around money, data loss, or health.

**6. Own the unhappy paths.**
_Why:_ Trust is won or lost in errors, permissions, empty states, and waiting. This is the highest-leverage, most-neglected surface in most products.
_How:_ Use a consistent error structure — what happened, plainly and without blaming the user; why, only if it changes their action; exactly what to do next. Never surface a raw code alone. Empty states should teach the primary action, not apologize. Permission prompts should state the concrete benefit in the user's terms before the OS dialog appears.

**7. Write for translation, truncation, and screen readers from the first draft.**
_Why:_ Retrofitting these is expensive and usually produces worse copy than designing for them.
_How:_ Avoid idioms, puns, and culture-bound humor. Never concatenate sentence fragments in code — hand translators whole strings with context notes. Assume many target languages run meaningfully longer than English and test at the tightest widths. Write link and button text that makes sense read out of context. Prefer sentence case; avoid long strings in all caps.

**8. Be honest, and measure the result.**
_Why:_ Manipulative patterns are both a trust liability and, increasingly, a regulatory one — and unmeasured copy is defended by seniority rather than evidence.
_How:_ No confirmshaming ("No thanks, I don't like saving money"), no invented urgency, no hidden exits, no ambiguity about what a destructive action destroys. Then instrument: task completion, error recovery rate, support ticket volume by topic, drop-off at the specific step. Ship changes with a metric attached.

## Common Mistakes

- **Writing last.** Copy handed a finished mock becomes an exercise in fitting words to boxes rather than shaping the flow.
- **Using internal vocabulary for user-facing concepts.** Feature names invented in planning docs, or system language leaking into errors, cost comprehension for no gain.
- **Personality at the wrong moment.** Wit during a failed payment, a lost file, or a health-related task reads as indifference and forces a re-read.

## Quick Start

- Pick one real flow in your product and write out every string, including all failure and empty states, in a single document. Read it aloud.
- Rewrite every button in that flow to name its outcome; delete every sentence that merely restates a button.
- Take your five highest-volume error messages and rewrite them as _what happened / why (if it matters) / what to do next_.
- Draft a one-page voice chart with real before-and-after examples, and hand it to an engineer to use without you.
- Instrument one change — measure drop-off or support tickets for that step before and after, so the next argument is settled with data.
