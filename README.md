# Mood Tracker

API to track daily habits and mood.

## Motivation

This is mostly a learning exercise to build something I will use myself, as well as testing out various libraries, design patterns and other Go related things.

## To Do
- [ ] User/Auth endpoints
- [ ] API level tests
- [ ] Build CLI interface to services
- [ ] Clone as a template for future Go APIs
- [ ] Endpoints for show current streak
- [ ] Habit tracking endpoints

## App Setup

```
main
  -| initialise data store
  -| initialise any other dependencies
  -| initialise services
server
  -| takes service
  -| handles http
  -| define handlers
service
  -| takes repository or db
  -| performs business logic
  -| define interface that can be used by any transport
```
