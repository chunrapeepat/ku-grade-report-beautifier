# Grade report beautifier

## About
`ku-grade-report-beautifier` is a grade report web application for all Kasetsart University students. our goal is to make a report more beautiful, readable and simple than the old one.

## How does it works
This section I will show you how this application work step by step to make you feel save to your own privacy.
1. Receive your account username & password from input (this app will not save your username OR password)
2. Send http request to login on real website
3. Get user information and grade report
4. Display on site. Done~

## Build on local
setup project on `GOPATH`
```bash
mkdir -p $GOPATH/src/github.com/chunza2542
cd $GOPATH/src/github.com/chunza2542
```
clone repository to local and run!
```bash
git clone https://github.com/chunza2542/ku-grade-report-beautifier.git

go run src/main.go # or `make dev`
```

---

The Chun Rapeepat production 2017.
