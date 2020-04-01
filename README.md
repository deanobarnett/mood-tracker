# Mood Tracker

## To Do
- Add user endpoints
- Clone as a template for future Go APIs
- Add habit endpoints

## App Setup

```
main
  -| initialise data store
  -| initialise user service
  -| initialise entry service
server
  -| takes service
  -| handles http
  -| define handlers
service
  -| takes repository or db
  -| performs business logic
  -| define interface that can be used by any transport
```
