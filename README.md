# Barbershop

Barbershop est un simple générateur de site statique qui utilise le
moteur de gabarit [Mustache](https://mustache.github.io).

## Installation

Pour installer Barbershop, il suffit de placer le script `barbershop`
à la racine de votre projet ou dans un répertoire sur votre `PATH`.

## Utilisation

Par défaut, Barbershop utilise la commande `mustache` qui se trouve
sur votre `PATH`. Si Mustache n'est pas installé globalement sur
votre ordinateur, il faut exécuter Barbershop avec comme argument le
chemin d'accès au fichier exécutable Mustache. Par exemple :

```sh
barbershop src site ./mustache
```

## Organisation du contenu

Barbershop localise les gabarits Mustache qui se trouvent dans un
répertoire source (par défaut, `src`), les transforme en document
HTML, et place ceux-ci dans un répertoire de sortie (par défaut,
`site`). À l'exception de la page d'accueil, qui doit se trouver
à la racine du répertoire source, chaque page doit être dans son
propre répertoire.

Barbershop cherche pour un fichier de données `data.json` dans le
même répertoire que le gabarit de la page. Ces données seront
disponibles seulement pour cette page. Si un `global.json` se trouve
à la racine du répertoire source, Barbershop rend ses données
disponibles à toutes les pages.

Les gabarits qui se trouvent dans des sous-répertoires dont le
nom commence avec un tiret bas sont ignorés. Si le répertoire source
contient un sous-répertoire nommé `assets`, celui-ci est copié tel
quel dans répertoire de sortie.

La structure suivante est suggérée :

```sh
.
├── site
└── src
    ├── _partials
    │   └── header.mustache
    │   └── ...
    ├── about # et tout autre page
    │   ├── data.json # données de la page about
    │   └── index.mustache # gabarit de la page about
    ├── assets
    │   ├── css
    │   └── ...
    ├── global.json # données globales pour toutes les pages
    ├── data.json # données de la page d'accueil
    └── index.mustache # gabarit de la page d'accueil
```