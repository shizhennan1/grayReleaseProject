FROM centos:7
COPY grayReleaseProject /root/server
EXPOSE 8080
CMD /root/server