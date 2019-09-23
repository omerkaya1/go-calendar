# This section will be completed in the fifth iteration of the CS
FROM some_registry.blabla.com/blabla:1.0.0

LABEL key="value"

USER root

COPY ./some/dir /opt/app/

RUN echo Fifth iteration, bitch.

ENTRYPOINT ["/opt/app/go-calendar/go-calendar", "-c", "/opt/app/go-calendar/configs/config.yaml"]
