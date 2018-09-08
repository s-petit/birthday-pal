# birthday-pal

[![Build Status](https://travis-ci.com/s-petit/birthday-pal.svg?branch=master)](https://travis-ci.com/s-petit/birthday-pal)
[![GoDoc](https://godoc.org/github.com/github.com/s-petit/birthday-pal?status.svg)](https://godoc.org/github.com/s-petit/birthday-pal)
[![Coverage Status](https://coveralls.io/repos/github/s-petit/birthday-pal/badge.svg?branch=master)](https://coveralls.io/github/s-petit/birthday-pal?branch=master)



* * *

# Disclaimer

*birthday-pal is still under development and not actually ready for use.*

# What is birthday-pal

A simple CardDAV birthday email reminder

The goal of this app is to remind you birthdays of your contacts, by sending emails.

There are many ways to save or consult birthdays of your friends and family, but the simplest way is to store it inside your contacts app.

It depends on your contacts provider, but in most cases, it does not provide reminder of birthdays, like calendars do for events.

Sometimes, birthdays are displayed as a Calendar, which is fine, but you still have to verify frequently the birthdays if you don't want to forget to wish it.

## How does it work

This app uses standard protocols to achieve its simple but necessary task : 

- CardDAV and Vcards for contacts because this is the standard and offers the best compatibily between contact systems.
- Emails / SMTP for reminders - Everyone consults his mailbox daily

You just have to provide :

- A Vcard file or a CardDAV server URL with credentials
- Your SMTP server URL with credentials
- The recipients

## How to use it

birthday-pal is a stateless executable app which needs :

- CardDav server HTTP URL, with Basic Auth credentials.
- SMTP server informations : host, port, username and password
- Recipients list
- Business Parameters:
-- days-before, which send a reminder n days before the birthday
-- remind-everyday, which will send a reminder everyday n days before the birthday, until the b-day.


Here is one exemple :

```
birthday-pal

  --carddav-url https://carddav.PROVIDER.com/URL \
  --carddav-user zeUser \
  --carddav-pass zePass \
  --smtp-host smtp.PROVIDER.com \
  --smtp-port 587 \
  --smtp-user smtpUser \
  --smtp-pass smtpPass \
  --days-before 3 \
  --remind-everyday \
  foo@mail.com bar@mail.com
```

There are plenty of parameters and it can be annoying to retype them for each call. So you can export CardDAV and SMTP values inside
environment variables :

```
export BPAL_SMTP_HOST=smtp.PROVIDER.com
export BPAL_SMTP_PORT=587
export BPAL_SMTP_USERNAME=smtpUser
export BPAL_SMTP_PASSWORD=smtpPass

export BPAL_CARDDAV_URL=https://carddav.PROVIDER.com/URL
export BPAL_CARDDAV_USERNAME=zeUser
export BPAL_CARDDAV_PASSWORD=zePass`
```

Thanks to these variables, you can shorten your call like this :


```
birthday-pal
  --days-before 3
  --remind-everyday
  --recipients foo@mail.com bar@mail.com
```

or even shorter :

`
birthday-pal -d 3 -e foo@mail.com bar@mail.com
`

Please note that the exported environments variables are overridable. The CLI option has a higher priority than the env variable.

## Design and limitations

birthday-pal is a simple executable which will perform its task once, as a one-shot. If you want to automate executions in a specific amount of time, let's say, daily or weekly, you have to deploy it on a server which runs permanently, and setup a cron yourself.

It was designed to run inside docker containers or a linux-based system on a server or NAS (like Synology)

Prequisites :

- Go or Docker

## Installation with Go
To install this package, run the following:

```shell
go get github.com/s-petit/birthday-pal
```

If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install/source) first

## Installation with Docker
To install this package, run the following:

```shell
docker build -t=birthday-pal-image .
docker run birthday-pal-image birthday-pal [ARGS]
```

### Timezone issues

If you encounter timezone issues with your container, please read this:

Mac users: https://github.com/docker/for-mac/issues/17#issuecomment-290667509
Linux users: https://stackoverflow.com/questions/24551592/how-to-make-sure-dockers-time-syncs-with-that-of-the-host

If you do not have the docker command on your system, you need to [Install Docker](https://docs.docker.com/install/) first


# Package Documentation

<!-- Do NOT edit past here. This is replaced by the contents of the package documentation -->




## License
This work is published under the MIT license.

Please see the `LICENSE` file for details.

* * *
Automatically generated by [autoreadme](https://github.com/jimmyfrasche/autoreadme) on 2018.09.08