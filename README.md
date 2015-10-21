# api-server-example

Work in progress example API Server built with golang.

Planned features:
- [x] Docker based dependencies vendoring
- [ ] Swagger documentation
- [ ] Statistics endpoints
- [ ] Periodically send statistics to StatHat or similar services 
- [ ] Heathcheck endpoints
- [ ] Build scripts for minimalistic Docker deployments 

## Swagger documentation

Swagger documentation can be created with 'make generate-swagger' command
Swagger *must* be generated manually every time code changes. This will be changed to static binary creation stage when ready.

When server is running swagger documentation endpoint is http://localhost:1323/docs/users

## Endpoints

users: http://localhost:1323/users
user #1: http://localhost:1323/users/1
