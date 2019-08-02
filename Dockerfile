FROM golang

COPY ./deploy-s3-bucket/main .

CMD ["main"]