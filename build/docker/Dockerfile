# -----------------------------------------------------------------------------
# Stage: BUILD
# -----------------------------------------------------------------------------
FROM golang:1.13.4-alpine3.10 as BUILD

ARG GITHUB_HEAD_REF="master"

# Add github
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        git

WORKDIR /go/src

# Clone repo
RUN git clone https://github.com/bryan-nice/git-issue-creation.git

# Checkout to target branch and if not master, then locally test
# merging master into branch.
WORKDIR /go/src/git-issue-creation
RUN if [[ "${GITHUB_HEAD_REF}" != "master" ]]; then \
        git checkout ${GITHUB_HEAD_REF}; \
        #git merge master; \
    fi

RUN cd git-issue-creation && \
    go mod download && \
    go get ./... && \
    go install ./...

# -----------------------------------------------------------------------------
# Stage: BUILD
# -----------------------------------------------------------------------------
FROM alpine:3.10

LABEL "com.github.actions.icon"="message-square"
LABEL "com.github.actions.color"="purple"
LABEL "com.github.actions.name"="Git Issue Creation"
LABEL "com.github.actions.description"="Create a Git Issue"

COPY --from=BUILD /go/bin/git-issue-creation /usr/bin
COPY --from=BUILD /go/src/git-issue-creation/LICENSE /
COPY --from=BUILD /go/src/git-issue-creation/README.md /

ENTRYPOINT ["git-issue-creation"]
