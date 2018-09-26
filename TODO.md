## Feature or refactoring ideas

## Features ideas

- send a "digest" reminder for a given period. Example : Here are the birthdays of the week...
- un vrai readme pro
- essayer de clean les args du SMTP pour faire mieux: bpal smtp ??
- logger dans un fichier (cf erreurs contacts google)


## Refacto

//TODO google Request : trouver une solution pour la page size... --> pagination des req google ?
/TODO SPE peut on cacher le token de 2 identifications differentes ?
//TODO gerer mieux le cote overridable de l url google ou abandonner
tenter de mutualiser la gestion des recipents
//TODO revoir la godoc
//TODO SPE revoir la visibilite de la plupart des fields et methods
//TODO faire un birthday-pal smtp ? rendre obligatoire le smtp ou alors faire une erreur claire ?
//TODO SPE: mutualiser smtp/reminder voire recipient
//TODO SPE duplication du code dans le CLI
//TODO SPE duplication du code dans les mocks complexes
BUG message erreur pas clair quand smtp non precise
BUG mowcli : confusion entre URL et recipients


