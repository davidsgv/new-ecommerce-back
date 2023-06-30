FROM golang:alpine

ARG git_user
ARG git_mail

RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

#install git
RUN apk update && apk upgrade && apk add git

#config git user and email
RUN git config --global user.name "${git_user}"
RUN git config --global user.email "${git_mail}"


EXPOSE 80