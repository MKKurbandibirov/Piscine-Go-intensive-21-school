FROM ubuntu:22.04

COPY --from=base /build/client/client /app/client

WORKDIR /app

CMD [ "./client" ]