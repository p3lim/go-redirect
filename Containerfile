FROM scratch

COPY --chown=65536:65536 --chmod=0777 bin/redirect /

ENV REDIRECT_SOURCE :8080

EXPOSE 8080/tcp

CMD ["/redirect"]
