# Masterclass Guide to HTML Email Templates

## Role Framing

Excellence in HTML email development means reliably rendering the same message across dozens of rendering engines — from modern webmail like Apple Mail to the Word-based Outlook on Windows — while staying accessible, responsive, dark-mode-aware, and under strict file-size limits. It is a discipline of constraint engineering, not modern web development.

## Best-in-Class Voices

- **Ted Goas** — creator of the _Cerberus_ responsive email patterns; emphasizes bulletproof table-based scaffolding and the hybrid coding approach that degrades gracefully in Outlook.
- **Rémi Parmentier (hteumeuleu.com / "Good Email Code")** — one of the most rigorous technical writers in email; emphasizes minimal-but-correct markup, semantic clarity, and honest testing over cargo-cult hacks.
- **Mark Robbins** — pioneer of interactive/"punched card" email and co-founder of goodemailcode.com; emphasizes progressive enhancement and accessibility-first thinking.
- **Nicole Merlin (Email Wizardry)** — popularized the _fluid hybrid_ ("spongy") layout pattern; emphasizes layouts that work even when media queries don't (notably older Gmail apps on Android).
- **The Litmus and Email on Acid teams** — de facto industry references through their testing tools, CSS-support matrices, and ongoing client rendering research. They emphasize rigorous cross-client testing as non-negotiable.
- **Can I email (caniemail.com, maintained by Parmentier and HTeuMeuLeu)** — the authoritative compatibility reference; emphasizes verifying support before using any CSS feature.

## Core Idea

Design for the **worst renderer you must support** (usually Outlook for Windows), then progressively enhance for modern clients — never the reverse. An email that breaks in Outlook is broken; an email that looks plain there but beautiful elsewhere is correct.

## Essential Best Practices

1. **Build on a table-based, hybrid-fluid skeleton.**
   _Why:_ Outlook on Windows renders with Microsoft Word's engine and ignores flexbox, grid, float, and most modern layout CSS. Tables remain the only layout primitive that works everywhere.
   _How:_ Use nested `<table role="presentation">` elements with `cellpadding="0" cellspacing="0" border="0"`. Combine fixed pixel widths (for Outlook) with `max-width` + percentage widths (for fluid behavior elsewhere) — the "fluid hybrid" pattern. Start from a vetted boilerplate like Cerberus or MJML rather than from scratch.

2. **Inline your CSS, but keep a `<style>` block for what can only live there.**
   _Why:_ Several clients (historically Gmail, and some webmail previews) strip or mishandle `<head>` styles. Inlined styles survive; media queries, pseudo-classes (`:hover`), and `@font-face` cannot be inlined.
   _How:_ Author in a readable embedded stylesheet, then run the file through an inliner (Juice, Premailer, or your ESP's built-in inliner) as a build step. Leave media queries, hover states, and dark-mode rules in the `<style>` block.

3. **Write only CSS that caniemail.com confirms is supported — and feature-detect the rest.**
   _Why:_ "It works in my browser" is meaningless. A property may be supported in Apple Mail but silently fail in Outlook 2016, Yahoo, or Samsung Mail.
   _How:_ Check every property against caniemail.com before shipping. Wrap risky features in `@supports`, MSO conditional comments (`<!--[if mso]>`), or Gmail-specific class hooks (`u + .body`, etc.). Treat modern CSS as a bonus layer, never a foundation.

4. **Handle Outlook explicitly with conditional comments and VML.**
   _Why:_ Outlook for Windows has its own set of failure modes — it ignores background images on most elements, miscalculates widths, and renders differently from every other client.
   _How:_ Use `<!--[if mso]>...<![endif]-->` to serve Outlook-only fixes (e.g., `<v:rect>` VML for background images, ghost tables for hybrid layouts, button rendering via `<v:roundrect>` when you need rounded corners). Give Outlook a plain, functional experience — it does not need parity with Apple Mail.

5. **Design mobile-first and keep it under Gmail's ~102 KB clipping threshold.**
   _Why:_ The majority of opens happen on mobile, and Gmail truncates messages above roughly 102 KB, hiding your unsubscribe link and tracking pixel behind a "View entire message" click. That hurts engagement and deliverability signals.
   _How:_ Single-column layouts by default, touch-friendly tap targets (generally ≥44×44 px), readable body text (typically 14–16 px minimum). Minify HTML, strip comments at build time, compress images, and monitor total size. Use media queries to stack and resize for small viewports.

6. **Treat dark mode as a first-class render target.**
   _Why:_ iOS Mail, Apple Mail, and Outlook.com aggressively invert or re-tint colors in dark mode — logos disappear on dark backgrounds, pure-black text becomes unreadable, transparent PNGs get white halos.
   _How:_ Use `@media (prefers-color-scheme: dark)` where supported, `[data-ol-dm]` and `meta name="color-scheme"` hints, and design logos/icons with a stroke or contrasting outline so they survive inversion. Avoid pure `#FFFFFF` and `#000000` for primary surfaces; test both modes explicitly.

7. **Make it accessible and readable without images.**
   _Why:_ Many clients block images by default until the user trusts the sender, and screen-reader users rely on semantic structure. An image-only email is often an empty email.
   _How:_ Meaningful `alt` text on every image (empty `alt=""` for purely decorative ones), styled alt text as a fallback, real HTML text for all critical content (especially CTAs), `lang` attribute on `<html>`, `role="presentation"` on layout tables, a descriptive `<title>`, and a preheader (hidden snippet text) that complements — not duplicates — the subject line.

8. **Test on real clients before every send — assumptions are expensive.**
   _Why:_ Client behavior changes (Gmail's rendering has shifted multiple times; Outlook releases vary widely; mobile apps update independently). What worked last quarter may break today.
   _How:_ Use Litmus, Email on Acid, or Mailtrap to preview across the client matrix that matches your audience. Test images-off, dark mode, and the major Outlook versions every time. Send seed tests to real inboxes — previews don't catch everything, especially forwarding and "view in browser" cases.

## Common Mistakes

- **Building with divs, flexbox, or CSS grid** because it "works in Gmail" — then discovering Outlook collapses the whole layout.
- **Shipping one light-mode design** and letting iOS/Outlook.com invert it into an unreadable mess with floating white logos.
- **Ignoring file size** until Gmail clips the footer, breaking the unsubscribe link and triggering spam-filter and compliance issues.

## Quick Start

- Fork a vetted starter (Cerberus, MJML, or Foundation for Emails) instead of writing your first template from scratch.
- Set up a build pipeline today: author → CSS inline → HTML minify → size check → Litmus/Email on Acid preview.
- Bookmark **caniemail.com** and consult it before adding any CSS property you haven't personally verified.
- Define a fixed test matrix (e.g., Gmail web + iOS, Apple Mail, Outlook 365 Windows, Outlook.com, one Android client) and run every template through it, in both light and dark mode, with images blocked.
- Write the plain-text version and preheader deliberately — they affect both accessibility and deliverability more than most teams realize.
