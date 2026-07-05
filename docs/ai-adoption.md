# AI-Assisted Development Guidelines

## Objective

Our goal is to use AI to improve engineering productivity—not to replace engineering judgment. AI should help us spend less time on repetitive implementation work and more time on solving the right technical problems.

Every engineer remains accountable for the code they submit, regardless of whether it was written manually or generated with AI.

---

## 1. Tooling

The team is encouraged to use AI-assisted development tools such as Claude, ChatGPT, and GitHub Copilot throughout the software development lifecycle.

Appropriate use cases include:

* Generating boilerplate code
* Creating unit test scaffolding
* Writing documentation and README files
* Explaining unfamiliar APIs or frameworks
* Refactoring code while preserving behaviour
* Producing SQL queries or migration templates
* Generating Dockerfiles, CI/CD configurations, and infrastructure templates
* Assisting with debugging and root-cause analysis

AI should be treated as an engineering assistant that accelerates implementation, not as the source of truth.

---

## 2. Guardrails

To maintain a production-quality codebase, the following guardrails apply to all AI-assisted development.

### Code Ownership

The engineer submitting the change is fully responsible for the correctness, security, and maintainability of the final implementation.

### Code Review

All AI-generated code follows the same peer review process as handwritten code. Reviewers focus on correctness, readability, performance, security, and alignment with our architecture.

### Testing

No AI-generated code is merged unless it passes:

* Unit tests
* Integration tests
* Static analysis
* Linting and formatting
* Benchmark validation where performance is critical

### Security

AI-generated code must never introduce:

* Hardcoded credentials or secrets
* Weak authentication or authorization
* Insecure dependency usage
* Sensitive data leakage

Security-sensitive changes require explicit human review.

### Maintainability

Generated code should:

* Follow existing coding standards
* Be easy to understand
* Avoid unnecessary complexity
* Include meaningful comments only where needed
* Be supported by appropriate tests

---

## 3. Scope of AI Usage

### AI can be used independently for

Low-risk, repetitive engineering tasks, including:

* Boilerplate implementation
* CRUD endpoints
* Data models
* DTOs
* Test generation
* Documentation
* Refactoring
* SQL generation
* Configuration files
* Benchmark scaffolding

These changes should still pass our normal review and CI process.

### AI-assisted work requiring engineering review

The following areas require deliberate human design and approval before implementation:

* System architecture
* API design
* Database schema design
* Authentication and authorization
* Distributed systems
* Concurrency
* Performance-critical components
* Infrastructure and deployment strategy
* Security-sensitive logic
* Business-critical workflows

AI may assist with implementation, but engineers are responsible for validating the design decisions.

---

## Team Expectations

We will use AI to reduce repetitive work and accelerate delivery while maintaining the same engineering standards we expect from manually written code. Every pull request should remain understandable, maintainable, well-tested, and production-ready. AI is a productivity tool—not a replacement for technical ownership, critical thinking, or sound engineering judgment.

---
---
---
# Translating AI Development Principles into AI Agent Prompts

To ensure consistent, production-quality AI-assisted development across the team, we standardize how AI agents are prompted. Rather than asking AI to "write some code," engineers provide context, constraints, and quality expectations so the generated output aligns with our engineering standards.

## 1. Provide Context

Every prompt should begin with enough project context for the AI to make informed decisions.

**Example Prompt**

> You are working on a Go microservice using Gin, ClickHouse, JWT authentication, and a generic concurrent pipeline library. Follow the existing project architecture and coding conventions. Do not introduce new frameworks or unnecessary dependencies.

---

## 2. Clearly Define the Task

Be explicit about what the AI should build and the expected outcome.

**Example Prompt**

> Implement a repository method to batch insert tracking events into ClickHouse using the existing repository pattern. Reuse existing models and error handling. The implementation should be production-ready.

---

## 3. Specify Engineering Constraints

AI should work within the team's architectural and coding standards.

**Example Prompt**

> Requirements:
>
> * Use Go generics where applicable.
> * Avoid global state.
> * Keep functions small and focused.
> * Return errors instead of panicking.
> * Ensure the implementation is safe for concurrent execution.
> * Minimize allocations where practical.

---

## 4. Define Quality Expectations

The AI should generate code that is complete rather than just functional.

**Example Prompt**

> Along with the implementation, provide:
>
> * Unit tests
> * Error handling
> * Benchmarks (if performance-sensitive)
> * Comments only where necessary
> * Any required interface changes

---

## 5. Preserve Existing Architecture

The AI should extend the system rather than redesign it.

**Example Prompt**

> Reuse the existing repository, service, and handler layers. Do not change public interfaces unless necessary. Follow the current dependency injection pattern.

---

## 6. Encourage Critical Review

AI should explain important decisions rather than acting as an unquestioned authority.

**Example Prompt**

> Explain any design decisions or trade-offs made. If multiple approaches exist, briefly describe why the chosen implementation is preferable.

---

## 7. Security and Performance Guardrails

For sensitive components, include explicit non-functional requirements.

**Example Prompt**

> The implementation must:
>
> * Avoid SQL injection.
> * Never expose secrets or credentials.
> * Be concurrency-safe.
> * Minimize unnecessary allocations.
> * Be suitable for production workloads.

---

## Team Prompt Template

> You are assisting with a production Go application. Follow the existing architecture and coding standards. Reuse existing interfaces and components wherever possible. Produce production-ready code with proper error handling, concurrency safety, and tests where applicable. Explain important design decisions and identify any assumptions. Do not introduce unnecessary dependencies or redesign existing components unless explicitly requested.

By standardizing prompts in this way, AI becomes a consistent engineering assistant rather than an ad hoc code generator. The prompts encode our architectural principles, quality expectations, and review standards, helping every engineer produce code that integrates cleanly into the codebase while maintaining accountability for the final implementation.

