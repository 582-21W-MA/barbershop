# Barbershop

Barbershop est un simple générateur de site statique qui utilise le
moteur de gabarit [Mustache](https://mustache.github.io).

## Installation

Pour installer Barbershop, il suffit de placer le fichier exécutable
ci-joint (voir *releases*) dans un répertoire sur votre `PATH`.

## Utilisation

La commande `barbershop` prend comme argument le répertoire d'entré
dans lequel se trouve les gabarits Mustache. Par exemple :

```sh
barbershop src
```

Étant donné la commande ci-haut, Barbershop interprètera les gabarits
qui se trouvent dans le répertoire `src`, et placera le résultat dans
un répertoire de sortie `site`.

### Serveur

La sous-commande `serve`, suivie du chemin d'accès vers le répertoire
`site`, permet de lancer un serveur de développement local :

```sh
barbershop serve site
```

### Guet

La sous-commande `watch`, suivie du chemin d'accès vers le
répertoire d'entrée, somme Barbershop de s'exécuter à chaque fois
qu'un changement est détecté. Un serveur de développement local est
également lancé. Par exemple :

```sh
barbershop watch src
```

## Organisation du contenu

### Pages

À l'exception de la page d'accueil, qui doit se trouver
à la racine du répertoire d'entré, chaque page doit être dans son
propre répertoire.

### Données

Barbershop cherche pour un fichier de données `data.json` dans le
même répertoire que le gabarit de la page. Ces données seront
disponibles seulement pour cette page. Si un fichier `global.json`
se trouve à la racine du répertoire d'entrée, Barbershop rend ses
données disponibles à toutes les pages.

### Partiels

Barbershop cherche les partiels dans un répertoire nommé
`_partials`. Par conséquent, il n'est pas nécessaire de fournir
le chemin d'accès de ceux-ci. Par exemple, si nous désirons
importer le partiel `_partials/header.mustache` à partir du gabarit
`about/index.mustache`, il suffit de spécifier `{{> header}}`.

### Ressources

Si le répertoire d'entrée contient un sous-répertoire nommé
`assets`, celui-ci est copié tel quel dans répertoire de sortie.

### Arborescence

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