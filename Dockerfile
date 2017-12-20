FROM scratch

ENV PORT 3000
EXPOSE 3000

ADD goMandel goMandel
CMD ["/goMandel"]
