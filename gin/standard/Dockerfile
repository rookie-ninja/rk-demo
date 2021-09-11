FROM alpine

ENV WD=/rk/service/

RUN mkdir -p $WD
WORKDIR $WD
COPY target/ $WD

EXPOST 8080

CMD ["bin/rk-demo"]
