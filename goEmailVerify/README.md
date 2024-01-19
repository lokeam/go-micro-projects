# Go Email Verifier
[Atom weight](https://en.wikipedia.org/wiki/Atomweight_(MMA)) Command line email verification tool. Powered by Golang.

## Featues
- Domain-based Message Authentication, Reporting & Conformance Checking
- Sender Policy Framework Checking
- Mail Exchange Record Checking

## Installation and Usage
```bash
# Clone the repo
$  git clone git@github.com:lokeam/go-micro-projects.git

# Navigate to the goEmailVerify directory
$  cd goEmailVerify

# Install all app dependencies
$  go mody tidy

# Boot up the app
$  go run main.go

# Input an email domain to verify. Viola.
```

Output Example:
```bash
# This is the comma-separated format that you'll receive results from your query
domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord

# Email domain example:
mailchimp.com

# Result:
mailchimp.com, true, true, v=spf1 ip4:205.201.128.0/20 ip4:198.2.128.0/18 ip4:148.105.0.0/16 ip4:129.145.74.12 include:mail.zendesk.com include:_spf.google.com include:mailsenders.netsuite.com include:_spf2.intuit.com include:_spf.qualtrics.com ip4:199.33.145.1 ip4:199.33.145.32 ip4:35.176.132.251 ip4:52.60.115.116 ~all, true, v=DMARC1; p=reject; rua=mailto:19ezfriw@ag.dmarcian.com,mailto:dmarc_rua@emaildefense.proofpoint.com; ruf=mailto:19ezfriw@fr.dmarcian.com,mailto:dmarc_ruf@emaildefense.proofpoint.com;
```
