# Barbershop

Barbershop est un simple générateur de site statique qui utilise le
moteur de gabarit Mustache.

## Installation

Pour installer Barbershop, il suffit de placer le script `barbershop`
à la racine de votre projet ou dans un répertoire sur votre `PATH`.

## Utilisation

Par défaut, Barbershop utilise la commande `mustache` qui se trouve
sur votre `PATH`. Si Mustache n'est pas installé globalement sur
votre ordinateur, il faut exécuter Barbershop avec comme argument le
chemin d'accès au fichier exécutable Mustache. Par exemple :

```sh
barbershop ./mustache
```

## Organisation du contenu

Barbershop localise les gabarits Mustache qui se trouvent dans le
répertoire `src`, les transforme en document HTML, et place ceux-ci
dans le répertoire `site`. À l'exception de la page d'accueil, qui
doit se trouver à la racine du répertoire `src`, chaque page doit
être dans son propre répertoire. Barbershop cherche pour un fichier
de données `data.json` dans le même répertoire que le gabarit de la
page.

Les gabarits qui se trouvent dans des sous-répertoires dont le
nom commence avec un tiret bas sont ignorés. Si `src` contient un
sous-répertoire nommé `assets`, celui-ci est copié tel quel dans
`site`.

La structure suivante est suggérée :

```sh
.
├── barbershop.sh
├── site
└── src
    ├── _partials
    │   └── header.mustache
    │   └── ...
    ├── about # et tout autre page
    │   ├── data.json # données de la page about
    │   └── index.mustache # gabarit de la page d'accueil
    ├── assets
    │   ├── css
    │   └── ...
    ├── data.json # données de la page d'accueil
    └── index.mustache # gabarit de la page d'accueil
```