FROM scratch
# Copy our static executable.
COPY ./com.software.temp /go/bin/com.software.temp

COPY ./conf/app.conf /go/bin/conf/app.conf
WORKDIR /go/bin/
EXPOSE 8080
# Run the hello binary.
ENTRYPOINT ["/go/bin/com.software.temp"]