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
Swagger documentation will be generated also with 'make build-static-binary' 

When server is running swagger documentation endpoint is http://localhost:1323/docs/users

## Endpoints

users: http://localhost:1323/users
user #1: http://localhost:1323/users/1
