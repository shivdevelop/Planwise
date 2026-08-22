# Planwise

Terraform plan intelligence and safety analysis.

Planwise is an open-source developer tool that analyzes Terraform plans to help engineers understand infrastructure changes, identify potential risks, and make safer decisions before terraform apply.

> **Build with confidence. Change infrastructure with clarity.**


## Overview

Terraform is excellent at describing and provisioning infrastructure, but a large Terraform plan can be difficult to review.

A plan may contain hundreds of resource changes, replacements, dependencies, and configuration modifications. Understanding the actual impact often requires manually inspecting the plan and knowing the behavior of the underlying cloud resources.

Planwise aims to provide an additional intelligence and safety layer around Terraform.

Instead of simply showing:

```text
Plan: 42 to add, 8 to change, 3 to destroy.
```

## Why Planwise?

Terraform tells you **what will change**.

Planwise aims to help you understand:

- What is changing?
- What resources are being created, modified, or destroyed?
- Which changes could be risky?
- What dependencies or side effects should be considered?
- Could a change cause an outage or unexpected behavior?
- Should this plan be reviewed before applying?

The project will progressively evolve from deterministic Terraform plan analysis into a broader infrastructure intelligence tool.

## Current Status

🚧 **Early development**

The project is currently in the foundation stage.

## Roadmap
- [ ] Terraform plan JSON parser
- [ ] Resource change detection
- [ ] Create / update / destroy analysis
- [ ] Resource replacement detection
- [ ] Risk classification
- [ ] Dependency analysis
- [ ] Security checks
- [ ] Blast-radius analysis
- [ ] Human-readable reports
- [ ] JSON output for CI/CD
- [ ] GitHub Actions integration
- [ ] Configurable policies
- [ ] AI-assisted analysis
- [ ] Automated remediation recommendations

## How It Will Work

```text
Terraform Configuration
          │
          ▼
   terraform plan
          │
          ▼
    Terraform JSON
          │
          ▼
       Planwise
          │
    ┌─────┴─────┐
    ▼           ▼
 Analysis      Rules
    │           │
    └─────┬─────┘
          ▼
    Risk Findings
          │
          ▼
      Report
```

The initial implementation will focus on analyzing Terraform's structured plan output rather than executing infrastructure changes.

## Example — Planned Experience

Eventually, Planwise should be able to turn a Terraform plan into a report such as:

``` text
Planwise
────────────────────────────────────


Terraform Plan Analysis


Resources
  + 12 to add
  ~  4 to change
  -  2 to destroy


Risk Summary
  CRITICAL   1
  HIGH       2
  MEDIUM     3
  LOW        4


Findings


[CRITICAL] production database
Resource replacement detected.


[HIGH] public network access
Potentially unsafe network configuration detected.


[MEDIUM] monitoring dependency
Potential resource availability issue detected.


Overall Risk: HIGH
```

This is a target experience, not functionality currently implemented.

## Design Principles

Planwise will follow a few principles:

- Safety first — never silently make infrastructure changes.
- Explainability — findings should explain why something is considered risky.
- Deterministic before AI — use explicit rules where possible before introducing AI.
- CI/CD friendly — analysis should work locally and inside pipelines.
- Cloud agnostic where possible — avoid coupling the core engine to a single cloud provider.
- Open source first — the project should remain useful without requiring proprietary services.

## Technology

The initial implementation will use:

- Go — CLI and core analysis engine
- Terraform — infrastructure plan source
- Terraform JSON — structured plan input

Additional technologies will be introduced only when they provide a clear benefit.

## Project Status

Planwise is currently an experimental open-source project under active development.

The architecture and APIs may change significantly during early development.

## Contributing

Contributions, ideas, issues, and discussions are welcome.

For larger changes, please open an issue first so the approach can be discussed before implementation.

## License

TBD
