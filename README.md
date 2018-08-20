# birthday-pal

# !!!!!!!!!!! WIP !!!!!!!!!!!!!!!!

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

## Design and limitations

birthday-pal is a simple executable which will perform its task once, as a one-shot. If you want to automate executions in a specific amount of time, let's say, daily or weekly, you have to deploy it on a server which runs permanently, and setup a cron yourself.

It was designed to run inside docker containers or a linux-based system on a server or NAS (like Synology)

Prequisites :

- Go or Docker
