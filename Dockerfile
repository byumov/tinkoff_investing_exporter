FROM centos:8
COPY ./tcs_go /app/tcs_go
CMD /app/tcs_go
