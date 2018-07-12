Name: myrpm
Version: 1
Release: 1
Summary: RPM for go-interns app
License: FIXME

%description
This is a RPM package for my go-interns application

%prep

%build
./build.sh

%install
mkdir -p %{buildroot}/opt/go-interns/templates
mkdir -p %{buildroot}/etc/systemd/system
install -m 755 main %{buildroot}/opt/go-interns/
install -m 755 templates/* %{buildroot}/opt/go-interns/templates/
install -m 777 conf.json %{buildroot}/opt/go-interns/
install -m 755 gointerns.service %{buildroot}/etc/systemd/system/

%post
useradd go-user
systemctl daemon-reload
systemctl start gointerns.service

%files
/opt/go-interns/*
/etc/systemd/system/gointerns.service

%changelog
