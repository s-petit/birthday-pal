# birthday-pal

[![Build Status](https://travis-ci.com/s-petit/birthday-pal.svg?branch=master)](https://travis-ci.com/s-petit/birthday-pal)
[![GoDoc](https://godoc.org/github.com/github.com/s-petit/birthday-pal?status.svg)](https://godoc.org/github.com/s-petit/birthday-pal)
[![Coverage Status](https://coveralls.io/repos/github/s-petit/birthday-pal/badge.svg?branch=master)](https://coveralls.io/github/s-petit/birthday-pal?branch=master)



* * *

## What is birthday-pal

`birthday-pal` is a birthday reminder.

The goal of this app is to remind you birthdays of your contacts, by sending email notifications.

The best way to store and sync contact data between all your devices is to use a contact provider, which uses standard protocols.

`CardDAV` protocol and `vCard` files are the standard for many contact providers. `Google Contacts` is also widely spread.

Notifications are sent inside an email, because everyone consults his mailbox regularly.

`birthday-pal` is written in `golang`, and can be built as a single executable, compatible for [most systems](https://golang.org/doc/install/source#environment)

## Why birthday-pal

`birthday-pal` was made for a simple purpose: Never forget to wish birthdays for people we care.

As far as i know, there are plenty of contact providers, used to store and sync contacts data, for a personal or professional purpose.

Many of them do not provide push notifications or reminders for birthdays, like calendars do for meetings and events.

Some people use their agenda app and create periodic events for each birthday, but it is pretty tedious and it duplicates the birthdate information which should be stored only inside the contact card.
This workaround makes your contacts management harder to maintain.

## Use cases

You may find `birthday-pal` useful when you want to :

- know the birthdays of the day in order to make your wishes
- be reminded several days before, in order to think about a gift
- have a periodic digest. For example, the birthdays of the week or the month
- ...and so on

Emails are available in English and French.


## Prerequisites

In order to use `birthday-pal` in an optimal way, you will need :

- A `contact provider`, with up-to-date data (at least contacts names and birthdates)
- A `SMTP` server, usually available via [email providers](https://en.wikipedia.org/wiki/Comparison_of_webmail_providers) or hosting solutions.
- Host URLs, ports, and credentials (only `BasicAuth` and `OAuth2` are supported)


## Targeted public (for now)

Because `birthday-pal` is a **CLI** tool which needs to connect to **external providers** and should be **scheduled** in order to be fully usable, we can't say that this tool is user friendly and accessible for everyone.

For now, this app is mainly intended for tech people like developers or sysadmins who want a simple way to remind birthdays for them or their acquaintances.


## How does it work

`birthday-pal` reads the name and birth date of your contacts, then send an email to notify you upcoming birthdays,
depending on several criteria.

This app uses standard / popular protocols to achieve this :

- `CardDAV/VCard` and `Google People API` in order to read contacts
- `BasicAuth` and `OAuth2` for authentication
- `SMTP` for email reminders

You just have to provide :

- A `CardDAV` or `Google` server URL with credentials (BasicAuth or OAuth2)
- A `SMTP` server URL with credentials
- The email recipients

`birthday-pal` is standalone and stateless. It does not use or store any data at all.
It just reads contacts, sends emails then shutdown.

You can execute the app manually, but in most cases, you will want to elaborate a routine,
which automatically launch the app in a given time interval (every day, every week...)

Depending on the executing OS, you have to setup a cron-like process in order to execute `birthday-pal` automatically (more details below).

## How to get it

There are 2 ways to get `birthday-pal`:

1. If you want to run the app on an amd64 macOS, Linux or Windows, download it on [github](https://github.com/s-petit/birthday-pal/releases)

2. Otherwise, build it yourself ([go](https://golang.org/doc/install) is required)

Just clone the project `git clone git@github.com:s-petit/birthday-pal.git`

Then build the application with [GOOS and GOARCH](https://golang.org/doc/install/source#environment) depending on the executing system :

For example :

Linux amd64
```
	GOARCH=amd64 GOOS=linux go build
```
macOS amd64
```
	GOARCH=amd64 GOOS=darwin go build
```

## How to run it

Once your executable is installed, just execute it:

For example, in Unix systems :

```
	./birthday-pal
```

## How to use it

`birthday-pal` is a stateless executable app which needs :

- Contacts (`CardDAV` or `Google`) server HTTP URL, with credentials (`BasicAuth` or `OAuth2`)
- `SMTP` server information : host, port, username and password
- Your recipients list
- Other Parameters:
    - days-before (number), reminders will be sent in a given number days before birthdays
    - less-than-or-equal (flag), Activates `<=` operator instead of `=`  to trigger reminders regarding the `days-before` option.
    - lang: `EN` (english), or `FR` (french)

You can have more details on syntax and default values with :

```
	./birthday-pal --help
```


Here are a few examples :

#### CardDAV and BasicAuth

> Given a contacts provider accessible via CardDAV protocol and BasicAuth

> Using a SMTP provider

> Remind Susan and John, who amongst their acquaintances, will celebrate their birthday **tomorrow** (ie. exactly in one day)


```
birthday-pal
  --smtp-host smtp.PROVIDER.com \
  --smtp-port 587 \
  --smtp-user smtpUser \
  --smtp-pass smtpPass \
  --lang FR \
  --days-before 1 \
  carddav \
  --url https://carddav.PROVIDER.com/URL \
  --user zeUser \
  --pass zePass \
  john@mail.com susan@mail.com
```

Please note that the app was designed to support `CardDAV` with `OAuth2` but i never tried it.

#### Google and OAuth2

`Google APIs` use `OAuth2` for authentication. Before using `birthday-pal`, you have to authorize the app for calling the APIs.

1. **Authentication**

Log in to the [Google Dev Console](https://console.developers.google.com) and create a new project.

Next, enable Google People API` inside the API Library menu, then create an `OAuth Client ID` via the Credentials menu.

Once the credential is created, download the json file. You can now perform the authentication:

`./birthday-pal oauth perform MYPROFILE google-credential.json`

Follow the instructions, and you're all set.

The authentication is linked to a profile id (MYPROFILE in the example), so you can perform and store as much authentications as you need.

Authentication data is stored in a dedicated folder: `.birthday-pal`

You can view the authorized profiles like this :

`./birthday-pal oauth list`

2. **Run birthday-pal towards Google People API with a authorized profile**

> Assuming today is Monday.

> Given my Google contacts provider accessible via People API and OAuth2

> Using my SMTP provider

> Remind me who will celebrate his birthday amongst my acquaintances, during the incoming week (ie. in 6 days or less, from Monday to Sunday)


```
./birthday-pal \
  --smtp-host smtp.PROVIDER.com \
  --smtp-port 587 \
  --smtp-user smtpUser \
  --smtp-pass smtpPass \
  --days-before 6 \
  --less-than-or-equal \
  google \
  MYPROFILE \
  me@mail.com

```

By default, `birthday-pal` calls `People API` like this : `https://people.googleapis.com/v1/people/me/connections?requestMask.includeField=person.names%2Cperson.birthdays&pageSize=50`

Depending on your needs, you may want to override this URL. You can do it with the `--url` option or a dedicated environment variable (see below)


## Using environment variables


As you can see, there are plenty of parameters and it can be annoying to retype them for each call. So you can export `CardDAV` and `SMTP` values inside
environment variables :

```
export BPAL_SMTP_HOST=smtp.PROVIDER.com
export BPAL_SMTP_PORT=587
export BPAL_SMTP_USERNAME=smtpUser
export BPAL_SMTP_PASSWORD=smtpPass

export BPAL_CARDDAV_URL=https://carddav.PROVIDER.com/URL
export BPAL_CARDDAV_USERNAME=zeUser
export BPAL_CARDDAV_PASSWORD=zePass

export BPAL_LANG=FR
export BPAL_GOOGLE_API_URL=https://people.googleapis.com/YOURQUERY

```

Thanks to these variables, you can shorten your call like this :


```
birthday-pal
  --days-before 3
  foo@mail.com bar@mail.com
```

or even shorter :

`
birthday-pal -d 3 foo@mail.com bar@mail.com
`

Please note that the exported environments variables are overridable. The CLI option has a higher priority than the env variable.


## Scheduling

`birthday-pal` is a simple executable which will perform its task once, as a one-shot. If you want to automate executions in a specific amount of time, let's say, daily or weekly, you have to deploy it on a server which runs permanently, and setup a `cron` yourself.

For example :

- `crontab` for linux systems
- `AT` for Windows systems
- Some NAS like `Synology` have a dedicated GUI to setup crons
- ...and so on.


## Docker

A `Dockerfile` is available is you need to run it inside a container.

To do it, run the following:

```shell
docker build -t=birthday-pal-image .
docker run birthday-pal-image birthday-pal [ARGS]
```

If you do not have the docker command on your system, you need to [Install Docker](https://docs.docker.com/install/) first


### Timezone issues

If you encounter timezone issues with your container, please read this:

Mac users: https://github.com/docker/for-mac/issues/17#issuecomment-290667509
Linux users: https://stackoverflow.com/questions/24551592/how-to-make-sure-dockers-time-syncs-with-that-of-the-host


# Package Documentation

<!-- Do NOT edit past here. This is replaced by the contents of the package documentation -->




## License
This work is published under the MIT license.

Please see the `LICENSE` file for details.

* * *
Automatically generated by [autoreadme](https://github.com/jimmyfrasche/autoreadme) on 2019.03.22
