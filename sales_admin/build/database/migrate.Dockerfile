FROM migrate/migrate:latest
RUN apk add postgresql-client
COPY ./build/database/scripts/migrate-wait-for-db.sh /migrate-wait-for-db.sh
ENTRYPOINT [ "/migrate-wait-for-db.sh" ]

