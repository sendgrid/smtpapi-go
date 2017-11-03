###Dockerfile for smtpapi-go
FROM golang:1.9

#Author Samuyi <@micheal_sazs>
MAINTAINER Samuyi

#create directory for src
RUN mkdir -p /go/src

#Copy files into conatiner
COPY . /go/src
