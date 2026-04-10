# Masterclass Guide to Site Reliability Engineering

## Role Framing

A Site Reliability Engineer applies software engineering discipline to operations, building systems that are reliable, scalable, and operable under failure — while protecting human attention and engineering velocity from being consumed by toil.

## Best-in-Class Voices

- **Benjamin Treynor Sloss (Google)** — Coined "SRE" and built the discipline at Google. Emphasizes error budgets, treating ops as a software problem, and capping toil.
- **Niall Murphy, Betsy Beyer, Chris Jones, Jennifer Petoff** — Editors of Google's _Site Reliability Engineering_ and _The SRE Workbook_. The canonical texts on SLOs, postmortems, and on-call design.
- **Charity Majors (Honeycomb)** — Pushes observability over monitoring, testing in production, and the idea that software ownership extends through deploy and operation.
- **Liz Fong-Jones (Honeycomb, ex-Google SRE)** — Clear teacher on SLOs, error budgets as a sociotechnical tool, and sustainable on-call.
- **John Allspaw (ex-Etsy, Adaptive Capacity Labs)** — Pioneer of blameless postmortems and resilience engineering; treats incidents as learning opportunities, not failures to eliminate.
- **Tanya Reilly (_The Staff Engineer's Path_, ex-Squarespace/Google SRE)** — Strong on reliability as a cross-cutting concern and the politics of getting reliability work prioritized.

---

## Core Idea

Reliability is a feature with a cost, not an absolute — the SRE's job is to make the right level of reliability cheap, observable, and negotiable, so the organization can move fast without breaking the things that matter.

---

## Essential Best Practices

**1. Define reliability with SLIs, SLOs, and error budgets — then actually use them.**
_Why:_ "Reliable" is meaningless without a target tied to user experience. Error budgets convert reliability from an argument into a math problem and align dev and ops incentives.
_How:_ Pick 2–4 SLIs that reflect what users actually feel (request success, latency at p99, freshness). Set SLOs slightly below current performance. When the budget burns, slow feature launches; when it's healthy, ship aggressively. Review quarterly.

**2. Eliminate toil as engineering work, not as a chore.**
_Why:_ Toil is manual, repetitive, automatable work that scales linearly with the system. Left unchecked, it consumes the team and crowds out the engineering that would prevent it. Google caps SRE toil at ~50% for this reason.
_How:_ Measure toil explicitly. Treat each recurring ticket as a bug against the system. Budget time each sprint for automation, and protect it.

**3. Make systems observable, not just monitored.**
_Why:_ Monitoring tells you when known problems happen; observability lets you ask new questions about novel failures — which is most of what production actually throws at you.
_How:_ Instrument with high-cardinality, structured events (not just metrics and logs). Favor tracing for distributed systems. The test: can you debug a problem you've never seen before without shipping new code?

**4. Run blameless postmortems and treat incidents as the most valuable data you have.**
_Why:_ People don't cause incidents; systems allow them. Blame destroys the honest reporting that makes learning possible, and most "human error" is really a system that made the wrong action easy.
_How:_ Separate the postmortem from performance review entirely. Focus on contributing factors, not root cause singular. Capture what _almost_ went worse. Track action items with the same rigor as feature work.

**5. Design for failure; assume every dependency will fail.**
_Why:_ At scale, rare failures happen constantly. Systems that assume the happy path become brittle in ways that only reveal themselves at 3 a.m.
_How:_ Use timeouts, retries with jitter, circuit breakers, graceful degradation, and load shedding. Practice failure: game days, chaos engineering, dependency-failure drills. Ask "what happens when X is slow?" — slow is worse than down.

**6. Protect on-call as a sustainable, humane practice.**
_Why:_ On-call burnout is the leading cause of SRE attrition and degraded judgment during incidents. A team that dreads its pager will not build reliable systems.
_How:_ Cap pages per shift (Google's heuristic: ≤2 actionable pages per 12-hour shift). Every page should be actionable, novel, and tied to a user-impacting SLO. Compensate on-call. Rotate fairly. Run a weekly review of every page — kill or fix the noisy ones.

**7. Practice capacity planning and change management as core disciplines.**
_Why:_ The two largest sources of outages are changes (deploys, configs, flags) and saturation. Both are predictable with discipline.
_How:_ Progressive rollouts (canary → percentage → global), automated rollback, config changes treated like code. For capacity: load test, model growth, leave headroom, and know your saturation point before traffic finds it.

**8. Make reliability work legible to the rest of the organization.**
_Why:_ SRE that can't explain its value gets cut in the next reorg. Error budgets, SLO dashboards, and postmortem trends are tools for negotiating with product and leadership, not just internal metrics.
_How:_ Publish SLO status where product managers see it. Tie reliability investments to specific incidents and user impact. Frame trade-offs in their language: revenue, churn, launch risk.

---

## Common Mistakes

- **Chasing 100% reliability.** It's infinitely expensive, blocks all change, and users can't tell the difference past a point — their own networks and devices fail more often than your service should.
- **Treating SRE as ops with a new title.** Without engineering authority, error budgets, and the right to push back on launches, it's just a rebranded NOC that burns out faster.
- **Optimizing for MTTR while ignoring change quality.** A team that's great at firefighting often has a deploy pipeline that keeps lighting fires. Prevention beats heroics.

---

## Quick Start

- **Pick one critical user journey this week** and define a single SLI + SLO for it. Don't try to instrument everything.
- **Audit your last 10 pages.** How many were actionable? Delete or fix the rest before adding any new alerts.
- **Run a blameless postmortem on your next incident**, however small — and publish it where engineers outside the team can read it.
- **Identify your top toil source** (the thing your team manually does most often) and budget engineering time to automate it this quarter.
- **Read _The SRE Workbook_, chapters 2 (SLOs) and 5 (Alerting on SLOs)** — they're the most directly applicable starting points in the canon.
