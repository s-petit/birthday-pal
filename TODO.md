## Feature or refactoring ideas

## Features ideas

- Make bpal work with Google's cardDav servers which use oauth2
- Try to run bpal for a given category of contacts
- Add validation for options and arguments (url, host, emails, nbdaysbefore...)
- accept basic auth url which contains user:pass https://user:pass@carddav.com/contacts
- implement a real i18n and templating solution
- send a "digest" reminder for a given period. Example : Here are the birthdays of the week...


## refacto

- Verify in the whole project that we favor pointers instead of values.
- Improve main test in bpal_test.go : Add assertions and test when the smtp creds are wrong : it should crash.
- Improve go doc everywhere.
- rename structs and variables if necessary
- try to separe real unit testing from integration testing.
- Add unit test with timezone weird cases (wish bday to a contact which lives in LA or Sydney)

