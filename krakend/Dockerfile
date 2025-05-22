FROM devopsfaith/krakend:latest

COPY krakend.json /etc/krakend/krakend.json

EXPOSE 8080

ENTRYPOINT [ "krakend" ]
CMD [ "run", "-c", "/etc/krakend/krakend.json" ]
