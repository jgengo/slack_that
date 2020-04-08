<h1 align="center"><code>Slack That!</code></h1>

<div align="center">
  <sub>Created by <a href="">Jordane Gengo (Titus)</a></sub>
</div>
<img src="https://goreportcard.com/badge/github.com/jgengo/slack_that" />

## Description

`Slack That!` is a microservice to deploy a slack posting message gateway designed to also work for multi workspace.

<img src="https://github.com/jgengo/slack_that/raw/master/static/slackthat_diagram.png" />
<img src="https://api.travis-ci.com/jgengo/slack_that.svg" />

## Why?

To avoid spreading your slack tokens in all your services and import a slack library when you can just easily send a HTTP POST to your micro service!

## Work in Progress.

### Tasks 

- [x] yaml parsing - workspace: token
- [x] Create GET / route - doc
- [x] Designing the POST route
- [x] Create POST / - send message
- [x] Throttle the slack postMessage to avoid rate limit
- [ ] Secure the POST route
- [ ] Dockerize or not.
- [ ] make the README.md more professional

### Ideas

Run out of ideas, atm.

### Contributors

- Gustavo Belfort <a href="https://github.com/Gustavobelfort">(Gustavobelfort)</a>
