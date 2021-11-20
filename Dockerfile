FROM golang:latest

ADD . ./app

WORKDIR /app

RUN while read l; do go get -v "$l"; done < <(go list -f '{{ join .Imports "\n" }}')

RUN go install

CMD ["talha"]
