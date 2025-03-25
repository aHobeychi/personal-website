---
id: Copilot Presentation
aliases: 
tags:
  - InnovateU
---

# Introduction
  
## Edit and Agents

Show the API, that it's working and call for the share price of CGI
Change the api calls from callbacks to async await

```text
Update the api calls to use async/await instead of callbacks
```

Generate Todo App using Claude 3.7

```text
I wan't a desktop webapp to manage a todo list. It needs to be written using VueJS 
```

Use image to modify the app. Doesn't work with Sonnet 3.7, switch to GPT 4o

Go back and show adding test to copilot demo

```tesxt
Can test this application using unit tests, add the necessary dependencies and mocks and run me through it
```

## Extensions

* Show Market Place, talk about different available extensions: Docker, Mermaid, StackOverflow, SonarQube
* Demo mermaid example
* Show Models Extensions

## Models

* Allows for different Models

## Cli

* Two main features
  * Suggest
  * Explain

```bash
gh copilot 
gh copilot suggest "create a python environment with numpy and pandas"
gh copilot explain "curl"
```

Use with Extensions, ex: Models

```bash
gh models
gh models list
gh models run Phi-3-mini-4k-instruct "Create me a python virtual environment containing numpy and pandas"
```

Important to note that It's like any CLI tool, they can be chained together using the pipe command

```bash
git log -n 10 | gh models run Phi-3-mini-4k-instruct "can you summarize this last few commits
```

Alias, you can setup alias to prevent yourself from writing the whole text command out ghce/ghcs

Demo for the pull request, From main branch create a new branch:

```bash
git checkout -b copilot-demo
git push 
```

Go back to main

```bash
git checkout main
ghcs undo the last 25 commits, and push the branch to github
```

## Pull Request

Go back to our Gatling demo, navigate to github create a PR and

