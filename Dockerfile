# RSHELL on Centos 7
FROM centos:latest
MAINTAINER Jonathan Mainguy jon@soh.re
ADD rshell /bin/
CMD ["/bin/rshell"]
