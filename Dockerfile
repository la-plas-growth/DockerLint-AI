FROM alpine:3.14
RUN apk add --no-cache python3
ENTRYPOINT ["python3"]