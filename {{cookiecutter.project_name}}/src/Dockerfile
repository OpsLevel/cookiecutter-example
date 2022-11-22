FROM golang:alpine as build
RUN apk --no-cache add ca-certificates

FROM {{ cookiecutter.docker_base_image }} AS release
# copy the ca-certificate.crt from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY {{ cookiecutter.project_slug }} /{{ cookiecutter.project_slug }}
ENTRYPOINT ["/{{ cookiecutter.project_slug }}"]