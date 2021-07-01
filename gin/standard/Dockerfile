FROM alpine

ENV WD=/rk/service/

RUN mkdir -p $WD
WORKDIR $WD
COPY target/ $WD

CMD ["bin/rk-demo"]