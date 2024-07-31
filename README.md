# App Containerization

## Problem

We want to be able to deploy multiple web applications in the same server

## Proposed Stack

- Ubuntu Server: run the applications and store the data
- NGINX proxy server: read the traffic from the outside and direct it to the wanted application
- MySQL: Store Persistent Data
- Golang: App #1
- PHP: App #2
- Docker: Contain both applications

