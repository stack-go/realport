FROM alpine:3
ARG PLUGIN_MODULE=github.com/stack-go/realport
ARG PLUGIN_GIT_REPO=https://github.com/stack-go/realport.git
ARG PLUGIN_GIT_BRANCH=master
RUN apk add --update git && \
    git clone ${PLUGIN_GIT_REPO} /plugins-local/src/${PLUGIN_MODULE} \
      --depth 1 --single-branch --branch ${PLUGIN_GIT_BRANCH}

FROM repo.intranet.pags/docker-hub/traefik:v2.8
COPY --from=0 /plugins-local /plugins-local
