```
#Masina nr. 4 MONITOR - Instalam httpd
#monitor.sh
yum update -y
sudo yum install httpd -y
systemctl start httpd
systemctl enable httpd
```
