<h1 align="center"><code>Slack That!</code></h1>

<div align="center">
  <sub>Created by <a href="">Jordane Gengo (Titus)</a></sub>
</div>

[![Go Report](https://goreportcard.com/badge/github.com/jgengo/slack_that)](https://goreportcard.com/badge/github.com/jgengo/slack_that) [![Build Status](https://travis-ci.com/jgengo/slack_that.svg?branch=master)](https://travis-ci.com/jgengo/slack_that)


## Work in Progress.

## Description

`Slack That!` is a microservice to deploy a slack posting message gateway designed to also work for multi workspace.

<img src="https://github.com/jgengo/slack_that/raw/master/static/slackthat_diagram.png" />

## Why?

To avoid spreading your slack tokens in all your services and import a slack library when you can just easily send a HTTP POST to your micro service!

At the moment, this micro-service is only meant to be used in a closed network and not exposed to external.

For example, you can run it in a specific docker network and add to this network the other services you want to allow to reach the microservice.

## Installing instructions

You need to create your config.yml with your workspace name as key and the slack token app as value. 

You can copy the sample provided to create your own:

```bash
cp ./configs/config.sample.yml ./configs/config.yml
```

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
