FROM scratch
WORKDIR /app/
ADD . /app/
EXPOSE 80
CMD ["./steampipe"]
