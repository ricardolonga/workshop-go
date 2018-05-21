FROM scratch

ADD ./workshop-go /workshop-go

CMD ["/workshop-go"]