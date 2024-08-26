FROM golang

COPY . .
RUN go build -o url-shortner

CMD [ "./url-shortner" ]
