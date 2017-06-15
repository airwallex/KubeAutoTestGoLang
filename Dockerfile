FROM golang

RUN go get github.com/airwallex/KubeAutoTestGoLang
RUN go install github.com/airwallex/KubeAutoTestGoLang

CMD CMD ["/go/bin/KubeAutoTestGoLang"]