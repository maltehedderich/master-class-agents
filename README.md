# Master Class Agents

Custom AI agents built from research-backed masterclass guides.

This repository shares the custom agents I use and will grow over time. Each agent starts as a role-specific masterclass guide created with a research-capable LLM, then gets distilled into a practical GitHub Copilot agent with the built-in `/create-agent` skill.

The repository also includes a workspace-scoped GitHub Copilot skill for generating new masterclass guides directly in chat.

## Install an Agent

If you want to use the included agents, clone the repository and run the installer for your tool:

```sh
git clone https://github.com/maltehedderich/master-class-agents
cd master-class-agents
./scripts/install-copilot-agents.sh
```

That command installs the GitHub Copilot agents into `~/.copilot/agents` by default.

## Choose an Installer

Each installer copies the included agents into the default folder for a specific tool.

| Tool           | Command                               | Default destination |
| -------------- | ------------------------------------- | ------------------- |
| GitHub Copilot | `./scripts/install-copilot-agents.sh` | `~/.copilot/agents` |
| Claude Code    | `./scripts/install-claude-agents.sh`  | `~/.claude/agents`  |
| Codex          | `./scripts/install-codex-agents.sh`   | `~/.codex/agents`   |
| Gemini CLI     | `./scripts/install-gemini-agents.sh`  | `~/.gemini/skills`  |

If you want a different destination, pass it as the first argument:

```sh
./scripts/install-copilot-agents.sh /path/to/agents
```

After installing Gemini CLI skills, run `/skills reload` if Gemini CLI is already open.

## Current Agents

The repository includes ten agents:

| Role                      | Masterclass guide                                                          | Agent                                                                                  |
| ------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| Backend Engineer          | [guides/backend-engineer.md](guides/backend-engineer.md)                   | [agents/backend-engineer.agent.md](agents/backend-engineer.agent.md)                   |
| Frontend Engineer         | [guides/frontend-engineer.md](guides/frontend-engineer.md)                 | [agents/frontend-engineer.agent.md](agents/frontend-engineer.agent.md)                 |
| Growth Hacker             | [guides/growth-hacker.md](guides/growth-hacker.md)                         | [agents/growth-hacker.agent.md](agents/growth-hacker.agent.md)                         |
| LinkedIn Writer           | [guides/linkedin-writer.md](guides/linkedin-writer.md)                     | [agents/linkedin-writer.agent.md](agents/linkedin-writer.agent.md)                     |
| Product Manager           | [guides/product-manager.md](guides/product-manager.md)                     | [agents/product-manager.agent.md](agents/product-manager.agent.md)                     |
| Prompt Engineer           | [guides/prompt-engineer.md](guides/prompt-engineer.md)                     | [agents/prompt-engineer.agent.md](agents/prompt-engineer.agent.md)                     |
| Site Reliability Engineer | [guides/site-reliability-engineer.md](guides/site-reliability-engineer.md) | [agents/site-reliability-engineer.agent.md](agents/site-reliability-engineer.agent.md) |
| Terraform Specialist      | [guides/terraform-specialist.md](guides/terraform-specialist.md)           | [agents/terraform-specialist.agent.md](agents/terraform-specialist.agent.md)           |
| Technical Educator        | [guides/technical-educator.md](guides/technical-educator.md)               | [agents/technical-educator.agent.md](agents/technical-educator.agent.md)               |
| Technical Writer          | [guides/technical-writer.md](guides/technical-writer.md)                   | [agents/technical-writer.agent.md](agents/technical-writer.agent.md)                   |

## Create a New Agent

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

## Create a Masterclass Guide in Copilot

This repo includes a workspace skill at `.github/skills/masterclass-guide/`.

In GitHub Copilot Chat, use `/masterclass-guide <role>` to generate a research-first guide that follows the same structure as [prompts/master-class.md](prompts/master-class.md).

## Why This Exists

Shoutout to [msitarzewski/agency-agents](https://github.com/msitarzewski/agency-agents/tree/main). That repository was the starting point for this project and showed how useful specialized agents can be.

After using it, I realized I wanted agents that fit my own needs more closely. Instead of writing agents directly from scratch, this repository uses a research-first workflow: start with a strong masterclass, then turn that synthesis into an agent.

That extra step matters. It keeps the agents grounded in durable practitioner guidance instead of ad hoc prompt writing.

## Repository Structure

```text
agents/   Ready-to-use agent definitions
guides/   Source masterclass guides
prompts/  Prompt templates used to generate the guides
.github/  Workspace-scoped GitHub Copilot customizations
```

## Add a New Agent

If you add a new role, keep both artifacts in the repository:

1. Add the source guide to `guides/<role>.md`.
2. Add the derived agent to `agents/<role>.agent.md`.
3. Keep file names in kebab-case.
4. Make sure the guide and the agent actually match.

## License

This repository is available under the [MIT License](LICENSE).
