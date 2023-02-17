FROM loads/alpine:3.8

WORKDIR /app

ENV GF_GERROR_BRIEF=true

COPY ./bin/linux_amd64/gf2-demo-api /app

EXPOSE 9000

ENTRYPOINT [ "./gf2-demo-api" ]
