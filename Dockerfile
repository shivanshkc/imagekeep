FROM golang:1.15.3-alpine as builder

# Create and change to the 'code' directory.
WORKDIR /code

# Build Application Native Binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o bin/application

FROM alpine:3

# Create and change to the the source directory.
WORKDIR /usr/src

# Copy the files to the production image from the builder stage.
COPY --from=builder /code/bin ./bin
COPY --from=builder /code/conf ./conf

# Run the web service on container startup.
CMD ["bin/application"]