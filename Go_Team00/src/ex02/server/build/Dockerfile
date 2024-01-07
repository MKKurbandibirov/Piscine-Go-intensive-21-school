FROM ubuntu:22.04

COPY --from=base /build/server/server /app/server

WORKDIR /app

EXPOSE 4000

CMD [ "./server" ]