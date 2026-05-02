# Master Class Agents

Custom AI agents and skills built from research-backed masterclass guides.

This repository shares the agents and skills I use, and grows over time. Each one starts as a role-specific masterclass guide created with a research-capable LLM, then gets distilled into either a persona-driven **agent** or a procedural **skill**.

The repository also includes a workspace-scoped GitHub Copilot skill for generating new masterclass guides directly in chat.

## Agents vs Skills

The same masterclass research can land in two different shapes. The choice depends on how the work is actually used.

- **Agent** — a persona you delegate broad judgment to. The agent file describes a role, principles, and an approach the assistant embodies across many situations. Use this when the work spans design, review, debugging, and strategy in a single discipline.
  - _Examples:_ Backend Engineer, Product Manager, UX Designer.
- **Skill** — a procedural workflow that produces one type of artifact. The skill file describes a clear procedure, inputs, constraints, and a quality bar for a specific deliverable. Use this when the work is "given X, produce Y" with a repeatable structure.
  - _Examples:_ Privacy Policy, Cookie Policy, Terms of Service, Travel Itinerary, LinkedIn Post, HTML Email Template.

The frontmatter `description` follows the same split: agent descriptions describe a role and judgment scope; skill descriptions start with action verbs and describe the deliverable.

## Disclaimer

The guides, agents, and skills in this repository are independent syntheses of ideas, practices, and public guidance associated with the people, companies, books, frameworks, and platforms they mention. They are not affiliated with, endorsed by, or official guidance from any of those individuals or organizations.

Any third-party names, product names, and trademarks mentioned in this repository belong to their respective owners and are used only for identification and commentary.

## Install Agents and Skills

If you want to use the included agents and skills, clone the repository and run `./mcagents`:

```sh
git clone https://github.com/maltehedderich/master-class-agents
cd master-class-agents
./mcagents install
```

The first run downloads a prebuilt `mcagents` CLI for your platform (or falls back to `go run` if you already have Go installed). Without flags it opens an interactive picker so you can choose the target tool, the agents, and the skills you want to install.

To install non-interactively — handy for dotfiles, CI, or scripts:

```sh
./mcagents install --tool claude --agents all --skills all
```

You can also install only a subset:

```sh
./mcagents install --tool copilot --agents backend-engineer,frontend-engineer
./mcagents install --tool codex   --skills privacy-policy
```

Install paths:

- **Repo bootstrap** — `./mcagents` (this is the recommended path).
- **`go install`** — `go install github.com/maltehedderich/master-class-agents/cli/cmd/mcagents@latest`. You still need a clone (or `--repo PATH`) for the source content.
- **GitHub Releases** — download a prebuilt binary directly from the [releases page](https://github.com/maltehedderich/master-class-agents/releases).

Run `./mcagents list` to see every available agent and skill, or `./mcagents list --tool claude` to also see where each one would be installed.

## Default Destinations

| Tool           | Agents go to        | Skills go to        |
| -------------- | ------------------- | ------------------- |
| GitHub Copilot | `~/.copilot/agents` | `~/.copilot/skills` |
| Claude Code    | `~/.claude/agents`  | `~/.claude/skills`  |
| Codex          | `~/.codex/skills`   | `~/.codex/skills`   |
| Gemini CLI     | `~/.gemini/skills`  | `~/.gemini/skills`  |

Codex and Gemini CLI install agents as skills, so both kinds land under their `skills/` directory. Pass `--dest /custom/path` to override either destination.

After installing Codex skills, restart Codex so it picks up the new `SKILL.md` files.

After installing Gemini CLI skills, run `/skills reload` if Gemini CLI is already open.

## Current Agents

The repository includes seventeen agents.

| Role                      | Masterclass guide                                                          | Agent                                                                                  |
| ------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| Backend Engineer          | [guides/backend-engineer.md](guides/backend-engineer.md)                   | [agents/backend-engineer.agent.md](agents/backend-engineer.agent.md)                   |
| Frontend Engineer         | [guides/frontend-engineer.md](guides/frontend-engineer.md)                 | [agents/frontend-engineer.agent.md](agents/frontend-engineer.agent.md)                 |
| Growth Hacker             | [guides/growth-hacker.md](guides/growth-hacker.md)                         | [agents/growth-hacker.agent.md](agents/growth-hacker.agent.md)                         |
| LlamaIndex Specialist     | [guides/llamaindex-workflows.md](guides/llamaindex-workflows.md)           | [agents/llamaindex-specialist.agent.md](agents/llamaindex-specialist.agent.md)         |
| Product Manager           | [guides/product-manager.md](guides/product-manager.md)                     | [agents/product-manager.agent.md](agents/product-manager.agent.md)                     |
| Prompt Engineer           | [guides/prompt-engineer.md](guides/prompt-engineer.md)                     | [agents/prompt-engineer.agent.md](agents/prompt-engineer.agent.md)                     |
| Security Engineer         | [guides/security-engineer.md](guides/security-engineer.md)                 | [agents/security-engineer.agent.md](agents/security-engineer.agent.md)                 |
| SEO Brief Architect       | [guides/seo-brief-architect.md](guides/seo-brief-architect.md)             | [agents/seo-brief-architect.agent.md](agents/seo-brief-architect.agent.md)             |
| SEO Content Drafter       | [guides/seo-content-drafter.md](guides/seo-content-drafter.md)             | [agents/seo-content-drafter.agent.md](agents/seo-content-drafter.agent.md)             |
| SEO Content Editor        | [guides/seo-content-editor.md](guides/seo-content-editor.md)               | [agents/seo-content-editor.agent.md](agents/seo-content-editor.agent.md)               |
| SEO Opportunity Analyst   | [guides/seo-opportunity-analyst.md](guides/seo-opportunity-analyst.md)     | [agents/seo-opportunity-analyst.agent.md](agents/seo-opportunity-analyst.agent.md)     |
| Site Reliability Engineer | [guides/site-reliability-engineer.md](guides/site-reliability-engineer.md) | [agents/site-reliability-engineer.agent.md](agents/site-reliability-engineer.agent.md) |
| Terraform Specialist      | [guides/terraform-specialist.md](guides/terraform-specialist.md)           | [agents/terraform-specialist.agent.md](agents/terraform-specialist.agent.md)           |
| Technical Educator        | [guides/technical-educator.md](guides/technical-educator.md)               | [agents/technical-educator.agent.md](agents/technical-educator.agent.md)               |
| Technical Writer          | [guides/technical-writer.md](guides/technical-writer.md)                   | [agents/technical-writer.agent.md](agents/technical-writer.agent.md)                   |
| UI Designer               | [guides/ui-designer.md](guides/ui-designer.md)                             | [agents/ui-designer.agent.md](agents/ui-designer.agent.md)                             |
| UX Designer               | [guides/ux-designer.md](guides/ux-designer.md)                             | [agents/ux-designer.agent.md](agents/ux-designer.agent.md)                             |

## Current Skills

The repository includes seven skills. Each lives under `skills/<name>/SKILL.md` so it can be dropped directly into a tool's skills directory.

| Skill                | Masterclass guide                                                      | Skill                                                                        |
| -------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| Cookie Policy        | [guides/cookie-policy-expert.md](guides/cookie-policy-expert.md)       | [skills/cookie-policy/SKILL.md](skills/cookie-policy/SKILL.md)               |
| HTML Email Templates | [guides/html-email-templates.md](guides/html-email-templates.md)       | [skills/html-email-templates/SKILL.md](skills/html-email-templates/SKILL.md) |
| Idiomatic Go         | [guides/go-language.md](guides/go-language.md)                         | [skills/idiomatic-go/SKILL.md](skills/idiomatic-go/SKILL.md)                 |
| LinkedIn Post        | [guides/linkedin-writer.md](guides/linkedin-writer.md)                 | [skills/linkedin-post/SKILL.md](skills/linkedin-post/SKILL.md)               |
| Privacy Policy       | [guides/privacy-policy-expert.md](guides/privacy-policy-expert.md)     | [skills/privacy-policy/SKILL.md](skills/privacy-policy/SKILL.md)             |
| Terms of Service     | [guides/terms-of-service-expert.md](guides/terms-of-service-expert.md) | [skills/terms-of-service/SKILL.md](skills/terms-of-service/SKILL.md)         |
| Travel Planner       | [guides/travel-planner.md](guides/travel-planner.md)                   | [skills/travel-planner/SKILL.md](skills/travel-planner/SKILL.md)             |

## Create a New Agent or Skill

Use this workflow:

1. Start with [prompts/master-class.md](prompts/master-class.md).
2. Run that prompt in a research-capable model.
3. Generate a masterclass guide for a role or topic and save it in [guides/](guides).
4. Decide whether the result is an agent or a skill (see [Agents vs Skills](#agents-vs-skills)).
5. Distill the guide into the matching artifact:
   - For an agent: write `agents/<role>.agent.md` with role, principles, approach, constraints, and output format.
   - For a skill: write `skills/<name>/SKILL.md` with description, inputs, constraints, procedure, branching logic, output format, quality checks, and failure modes.
6. Refine until the guidance is opinionated, practical, and usable in real work.

The guide is the research artifact. The agent or skill is the working artifact.

Example output for a new role:

```text
Role: Security Engineer
Guide: guides/security-engineer.md
Agent: agents/security-engineer.agent.md
```

Example output for a new procedural skill:

```text
Topic: Cookie Policy
Guide: guides/cookie-policy-expert.md
Skill: skills/cookie-policy/SKILL.md
```

## Create a Masterclass Guide in Copilot

This repo includes a workspace skill at `.github/skills/masterclass-guide/`.

In GitHub Copilot Chat, use `/masterclass-guide <role>` to generate a research-first guide that follows the same structure as [prompts/master-class.md](prompts/master-class.md).

## Why This Exists

Shoutout to [msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents/tree/main). That repository was the starting point for this project and showed how useful specialized agents can be.

After using it, I realized I wanted agents that fit my own needs more closely. Instead of writing agents directly from scratch, this repository uses a research-first workflow: start with a strong masterclass, then turn that synthesis into an agent or a skill — whichever shape fits how the work is actually used.

That extra step matters. It keeps the artifacts grounded in durable practitioner guidance instead of ad hoc prompt writing.

## Repository Structure

```text
agents/    Persona-driven agent definitions
skills/    Procedural skill definitions (one folder per skill, with SKILL.md)
guides/    Source masterclass guides
prompts/   Prompt templates used to generate the guides
.github/   Workspace-scoped GitHub Copilot customizations and CI workflows
cli/       Source for the mcagents CLI (Go module)
mcagents   Bootstrap shell wrapper that downloads or runs the CLI
```

## Add a New Agent or Skill

If you add a new role or topic, keep both artifacts in the repository:

1. Add the source guide to `guides/<name>.md`.
2. Add the derived agent to `agents/<name>.agent.md` _or_ the derived skill to `skills/<name>/SKILL.md`.
3. Keep file and folder names in kebab-case.
4. Make sure the guide and the agent or skill actually match.
5. Update the relevant table in this README.

## License

This repository is available under the [MIT License](LICENSE).
