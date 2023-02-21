# Introduction

Backend for aahar app

# How to build locally

1. Clone the repo
1. Install make, docker, golang-migrate, sqlc
1. Install docker image for postgres 15-alpine
1. Run the following commands

- make postgres
- make createdb
- make migrateup
- make sqlc

## Additional Steps for Windows user

sqlc the package used to generate the go code from sql code doesn't work in windows for postgres for that you need to install wsl and then install homebrew in it and then use the command make sqlc in order to generate the go code from sql query.
