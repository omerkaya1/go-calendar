FROM registry.yadro.com/suse/sles15-x86_64:1.1.9

LABEL key="value"

USER root

#COPY ./build/bin/syr /opt/app/syr/syr
#COPY ./static/dist /opt/app/syr/static/dist
#COPY ./pdf/fonts /opt/app/syr/pdf/fonts
#COPY ./ssl /opt/app/syr/ssl

#RUN zypper --quiet --no-gpg-checks install -y -l --no-recommends \
#	ca-certificates \
#	ca-certificates-mozilla
#
#ENTRYPOINT ["/opt/app/syr/syr", "-stderrthreshold=INFO", "-config", "/opt/app/syr/configuration.json"]
