# DRS
*drs* is a command line utility for  non-interactive download of files from the Web.  It supports HTTP, HTTPS protocols, as well as retrieval through HTTP proxies.

<br>

## ✨ Features

- Just one command to rule them all.
- Written in Go, and works for SPEED
- Distraction Free Downloads, simple workflow
- A simpler replacement to wget

<br>

## ⚡️ Installation
### Method 1

**For Linux users:**

Execute the following command as **Root User** (sudo su):

```bash
curl https://raw.githubusercontent.com/arghyagod-coder/GoLang/main/drs/linux_install.sh > drs_install.sh
chmod u+x ./drs_install.sh
bash ./drs_install.sh
```


**For MacOS users:**

Execute the following command as **Root User** (sudo su):

```bash
curl https://raw.githubusercontent.com/arghyagod-coder/GoLang/main/drs/mac_install.sh > drs_install.sh
chmod +x ./drs_install.sh
bash ./drs_install.sh
```

**For Windows users:**

Open Powershell **as Admin** and execute the following command:
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; (Invoke-WebRequest -Uri https://raw.githubusercontent.com/arghyagod-coder/GoLang/main/drs/windows_install.ps1 -UseBasicParsing).Content | powershell -
```

To verify the installation of *drs*, open a new shell and execute `drs -v`. You should see output like this:
```
drs 1.2.0

Version: 1.2.0
```
If the output isn't something like this, you need to repeat the above steps carefully.

<br>

## 💡 Usage
`drs` usage is simple

```bash
drs {url} --fn {filename}
```

--fn flag stands for filename. If I want a file named foo from https://foo.bar/foo.sh, I can execute:

```bash
drs https://foo.bar/foo.sh --fn foo
```

<br>

