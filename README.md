# Repo for INGO-MMT Golang Testing Team Coding Event

Please try to follow the Initial Setup before we start the session so that everyone is able to follow through.

- [Initial Setup (PRE)](_docs/INITIAL.md)

## Project Outline

We are building a banking system (believe it!). We have accounts, an account holder can:
- create an account.
- deposit money.
- withdraw money.
- see their balance.

And the interface is a REST API. We are using the Gin framework for building these functionalities. We have the following APIs:
- POST /api/v1/accounts (Creates an account)
- PUT /api/v1/accounts/{accountId}/deposit (Deposit money)
- PUT /api/v1/accounts/{accountId}/withdraw (Withdraw money)
- GET /api/v1/accounts/{accountId} (See the account info, which will also have the balance)

We do have some bugs here and there and code is not tested. We will close these in the meeting. Please go through the
project and familiarise if you get time.
