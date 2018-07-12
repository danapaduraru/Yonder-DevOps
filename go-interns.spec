Name:	go-interns
Version: 1
Release: 1
Summary: Go-interns RPM package
License: FIXME

%description

%prep

%build
./build.sh

%install
mkdir -p %{buildroot}/opt/go-interns/templates
mkdir -p %{buildroot}/etc/systemd/system
install -m 755 main %{buildroot}/opt/go-interns/main
install -m 755 config.json %{buildroot}/opt/go-interns/config.json
install -m 755 templates/* %{buildroot}/opt/go-interns/templates
install -m 755 go-interns.service %{buildroot}/etc/systemd/system

%files
/opt/go-interns/*
/etc/systemd/system/go-interns.service
%changelog
