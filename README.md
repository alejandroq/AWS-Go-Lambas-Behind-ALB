# Front Ending an Amazon Web Service (AWS) Lambda by an Application Load Balancer (ALB) in-lieu of an API Gateway
*By Alejandro Quesada v0.0.0*

## Goal

Simplify the implementation of an AWS ALB with a Lambda target for reducing operational costs when niche API Gateway features are not being used yet the API(s) within the Lambda still need to be client accessible (likely via the internet).

Cloud operation costs and possible savings may not be worth the effort for mid-to-large sized companies, however for hobbyists and small business' such savings can be highly valuable.

An article to read on a comparison of costs: [https://serverless-training.com/articles/save-money-by-replacing-api-gateway-with-application-load-balancer/](https://serverless-training.com/articles/save-money-by-replacing-api-gateway-with-application-load-balancer/).

A secondary goal is to establish a quick baseline to get started with ATDD via `godog`. See integration testing below.

## Setup

I developed the following example whilst using Go version `1.12.6`. My `GO111MODULE` is set to `on`, though I am not sure if this is still required to use `go mod` or if it is on by default now. Provided `go mod` is in-use, this repo does not need to live in the `$GOPATH` and was in-fact developed in a directory on my desktop. I have vendored the dependencies to simplify re-use, but if the `vendor` directory were to be deleted, I believe `go mod` would simply repopulate said dependencies on the next run of the code.

## Running all Automated Checks

Including, but not limited to: vet, lint, fmt and unit tests via: `make`.

At the end your default web browser will appear showcasing test coverage. In the *Terminal* you ran `make`, a coverage report over functions will appear in `stdout`. 

## Running Unit Test (TDD)

Run `make test`. If you want to view a coverage report, following `make test` run `make coverage`.

Quick shorthand for the latter: `make test && make coverage`.

## Running Integration Tests (ATDD)

In *Terminal A*, run `make start`. This will kickstart a local AWS Lambda RPC on `127.0.0.1:8001`.

In *Terminal B*, run `make integration`. This will run a slew of ATDD tests against *Terminal A's* Lambda RPC process.

## TODO

- Security Groups for ALB
- Other related infrastructure, will likely implement with Terraform. 
