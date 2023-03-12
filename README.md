phishing_detector/
  |- main.go
  |- config/
  |    |- config.toml
  |- domain/
  |    |- domain.go
  |- url/
  |    |- url.go
  |- virustotal/
  |    |- virustotal.go
  |- safebrowsing/
       |- safebrowsing.go

* main.go: the main entry point of the app, which initializes and runs the app.
* config/config.toml: the configuration file for the app, which contains API keys for Virustotal and Google Safe Browsing.
* domain/domain.go: a package that provides functions for extracting the domain from a URL and performing a WHOIS lookup on the domain.
* url/url.go: a package that provides functions for validating a URL and checking if it redirects to another URL.
* virustotal/virustotal.go: a package that provides a function for checking a URL against the Virustotal API for malware and phishing.
* safebrowsing/safebrowsing.go: a package that provides a function for checking a URL against the Google Safe Browsing API for phishing and malware.

creating a phishing scanner 
aim is to integrate with 
1- Google Safe Browsing https://developers.google.com/safe-browsing/v4
2- Chatgpt https://platform.openai.com/docs/guides/chat
3- Virustotal
4- Hybrid-analysis
5- Urlscan 
6- Domaintools



later on we can build browser extention.

UX Flow:
url submission
validate the url
unshorten url if required
follow redirect if required

check doamin
1-creation date
2-subdomain numbers,names,ips
3-dns

show screenshot
google safe browsing

eml flow:
