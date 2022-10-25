# CrowdSec

[![CodeFactor](https://www.codefactor.io/repository/github/wojtekxtx/crowdsec/badge)](https://www.codefactor.io/repository/github/wojtekxtx/crowdsec)

CrowdSec is a free, modern & collaborative behavior detection engine, coupled with a global IP reputation network. It stacks on fail2ban's philosophy but is IPV6 compatible and 60x faster (Go vs Python), uses Grok patterns to parse logs and YAML scenario to identify behaviors. CrowdSec is engineered for modern Cloud / Containers / VM based infrastructures (by decoupling detection and remediation). Once detected you can remedy threats with various bouncers (firewall block, nginx http 403, Captchas, etc.) while the aggressive IP can be sent to CrowdSec for curation before being shared among all users to further improve everyone's security.

# Install

Best way is to install `CrowdSec` from source. This way you will always get
the newest version.

Install it by ( **We are not quite there yet**):
```sh
wget https://github.com/wojtekxtx/crowdsec/releases/latest/download/crowdsec-release.tgz
tar xzvf crowdsec-release.tgz
cd crowdsec-v* && sudo ./wizard.sh -i
```




## :information_source: About the CrowdSec project

Crowdsec is an open-source, lightweight software, detecting peers with aggressive 
behaviors to prevent them from accessing your systems. Its user friendly design
and assistance offers a low technical barrier of entry and nevertheless a high security gain.

# Concept

Once an unwanted behavior is detected, deal with it through a [bouncer](https://hub.crowdsec.net/browse/#bouncers). The aggressive IP, scenario triggered and timestamp are sent for curation, to avoid poisoning & false positives. (This can be disabled). If verified, this IP is then redistributed to all CrowdSec users running the same scenario.

# Outnumbering hackers all together

By sharing the threat they faced, all users are protecting each-others (hence the name Crowd-Security). Crowdsec is designed for modern infrastructures, with its "*Detect Here, Remedy There*" approach, letting you analyse logs coming from several sources in one place and block threats at various levels (applicative, system, infrastructural) of your stack.

# :point_right: What it is not

CrowdSec is not a SIEM, storing your logs (neither locally nor remotely). Your data are analyzed locally and forgotten.

Signals sent to the curation platform are limited to the very strict minimum: IP, Scenario, Timestamp. They are only used to allow the system to spot new rogue IPs, rule out false positives or poisoning attempts.
# :tada: Key benefits
## Fast assisted installation, no technical barrier

<details open>
  <summary>Initial configuration is automated, providing functional out-of-the-box setup</summary>
  <img src="https://github.com/wojtekxtx/crowdsec-docs/blob/main/crowdsec-docs/static/img/crowdsec_install.gif?raw=true">
</details>

## Out of the box detection

<details>
  <summary>Baseline detection is effective out-of-the-box, no fine-tuning required (click to expand)</summary>
  <img src="https://github.com/wojtekxtx/crowdsec-docs/blob/main/crowdsec-docs/static/img/out-of-the-box-protection.gif?raw=true">
</details>

## Easy bouncer deployment

<details>
  <summary>It's trivial to add bouncers to enforce decisions of crowdsec (click to expand)</summary>
  <img src="https://github.com/wojtekxtx/crowdsec-docs/blob/main/crowdsec-docs/static/img/blocker-installation.gif?raw=true">
</details>

## Easy dashboard access

<details>
  <summary>It's easy to deploy a metabase interface to view your data simply with cscli (click to expand)</summary>
  <img src="https://github.com/wojtekxtx/crowdsec-docs/blob/main/crowdsec-docs/static/img/cscli-metabase.gif?raw=true">
</details>

## Hot & Cold logs

<details>
  <summary>Process cold logs, for forensic, tests and chasing false-positives & false negatives (click to expand)</summary>
  <img src="https://github.com/wojtekxtx/crowdsec-docs/blob/main/crowdsec-docs/static/img/forensic-mode.gif?raw=true">
</details>


# Contributing

If you wish to contribute to the core of crowdsec, you are welcome to open a PR in this repository.