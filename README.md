# Master Class Agents

Custom AI agents built from research-backed masterclass guides.

This repository shares the custom agents I use and will grow over time. Each agent starts as a role-specific masterclass guide created with a research-capable LLM, then gets distilled into a practical GitHub Copilot agent with the built-in `/create-agent` skill.

## Quick Start

Use this workflow to create a new agent:

1. Start with [prompts/master-class.md](prompts/master-class.md).
2. Run that prompt in a research-capable model.
3. Generate a masterclass guide for a role and save it in [guides/](guides).
4. Use GitHub Copilot's built-in `/create-agent` skill to turn that guide into an agent file in [agents/](agents).
5. Refine the agent until the guidance is opinionated, practical, and usable in real work.

The guide is the research artifact. The agent is the working artifact.

Example output for a new role:

```text
Role: Security Engineer
Guide: guides/security-engineer.md
Agent: agents/security-engineer.agent.md
```

## Current Agents

The repository includes six agents:

| Role                      | Masterclass guide                                                | Agent                                                                        |
| ------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| Backend Engineer          | [guides/backend-engineer.md](guides/backend-engineer.md)         | [agents/backend-engineer.agent.md](agents/backend-engineer.agent.md)         |
| Frontend Engineer         | [guides/frontend-engineer.md](guides/frontend-engineer.md)       | [agents/frontend-engineer.agent.md](agents/frontend-engineer.agent.md)       |
| Product Manager           | [guides/product-manager.md](guides/product-manager.md)           | [agents/product-manager.agent.md](agents/product-manager.agent.md)           |
| Site Reliability Engineer | [guides/sre.md](guides/sre.md)                                   | [agents/sre.agent.md](agents/sre.agent.md)                                   |
| Terraform Specialist      | [guides/terraform-specialist.md](guides/terraform-specialist.md) | [agents/terraform-specialist.agent.md](agents/terraform-specialist.agent.md) |
| Technical Writer          | [guides/technical-writer.md](guides/technical-writer.md)         | [agents/technical-writer.agent.md](agents/technical-writer.agent.md)         |

## Install The Agents

Install the included GitHub Copilot agents into your user-level agents folder:

```sh
./scripts/install-copilot-agents.sh
```

By default this copies the `.agent.md` files into `~/.copilot/agents`.

If you want a different location, pass it explicitly:

```sh
./scripts/install-copilot-agents.sh /path/to/agents
```

There is also a matching Claude Code installer:

```sh
./scripts/install-claude-agents.sh
```

And a Codex installer that writes standalone role files:

```sh
./scripts/install-codex-agents.sh
```

And a Gemini CLI installer that converts each `.agent.md` file into a skill:

```sh
./scripts/install-gemini-agents.sh
```

By default this writes to `~/.gemini/skills`. If Gemini CLI is already open, run
`/skills reload` after installing.

## Why This Exists

Shoutout to [msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents/tree/main). That repository was the starting point for this project and showed how useful specialized agents can be.

After using it, I realized I wanted agents that fit my own needs more closely. Instead of writing agents directly from scratch, this repository uses a research-first workflow: start with a strong masterclass, then turn that synthesis into an agent.

That extra step matters. It keeps the agents grounded in durable practitioner guidance instead of ad hoc prompt writing.

## Repository Structure

```text
agents/   Ready-to-use agent definitions
guides/   Source masterclass guides
prompts/  Prompt templates used to generate the guides
```

## Add a New Agent

If you add a new role, keep both artifacts in the repository:

1. Add the source guide to `guides/<role>.md`.
2. Add the derived agent to `agents/<role>.agent.md`.
3. Keep file names in kebab-case.
4. Make sure the guide and the agent actually match.

## License

This repository is available under the [MIT License](LICENSE).
