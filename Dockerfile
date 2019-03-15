FROM registry.fedoraproject.org/fedora-minimal

RUN mkdir -p /work

COPY ./bin/kbsc /work/kbsc

WORKDIR  /work

ENTRYPOINT [ "/work/kbsc" ]

CMD [ "--help" ]