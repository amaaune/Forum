# Forum
# 📁 Forum — Arborescence du projet

Projet de forum web développé en Go avec SQLite.

---

## Structure du projet

```
forum/
├── main.go                  # Point d'entrée — démarre le serveur HTTP et enregistre les routes
├── go.mod                   # Déclare le module et les dépendances (sqlite3, bcrypt, uuid)
├── go.sum                   # Checksums auto-générés par Go — ne pas éditer
├── Dockerfile               # Recette pour construire l'image Docker
├── .dockerignore            # Fichiers exclus de l'image Docker (forum.db, tests…)
└── README.md                # Documentation du projet

├── database/
│   ├── database.go          # Connexion SQLite + CREATE TABLE au démarrage
│   └── forum.db             # Fichier BDD généré automatiquement

├── models/                  # Structs Go qui représentent les données
│   ├── user.go
│   ├── post.go
│   ├── comment.go
│   └── category.go

├── handlers/                # Logique métier de chaque route HTTP
│   ├── auth.go              # Register / login / logout
│   ├── posts.go             # CRUD posts
│   ├── comments.go          # CRUD commentaires
│   ├── likes.go             # Likes / dislikes
│   ├── filter.go            # Filtres : catégorie, mes posts, mes likes
│   └── errors.go            # Gestion erreurs HTTP 404 / 500

├── middleware/              # Intercepteurs exécutés avant chaque route
│   ├── auth.go              # Bloque les routes si non connecté
│   └── session.go           # Création / lecture / expiration des cookies UUID

├── security/                # Fonctions utilitaires de sécurité
│   ├── password.go          # Hash bcrypt + vérification
│   └── validate.go          # Contrôle email, doublons, inputs vides

├── templates/               # Fichiers HTML rendus côté serveur par Go
│   ├── poly/                # Composants HTML réutilisés sur plusieurs pages
│   │   ├── header.html
│   │   └── footer.html
│   ├── index.html           # Accueil + fil de posts
│   ├── login.html           # Connexion
│   ├── register.html        # Inscription
│   ├── post.html            # Détail d'un post + commentaires
│   ├── category.html        # Sous-forum / catégorie
│   └── error.html           # Pages erreur 404 / 500

├── static/                  # Fichiers servis directement au navigateur
│   ├── css/
│   │   ├── base.css         # Variables CSS, reset, typographie — chargé sur toutes les pages
│   │   ├── header.css       # Navbar — toutes les pages
│   │   ├── footer.css       # Footer — toutes les pages
│   │   ├── modal.css        # Modales — toutes les pages
│   │   ├── buttons.css      # Boutons communs — toutes les pages
│   │   ├── alerts.css       # Messages erreur / succès — toutes les pages
│   │   ├── index.css        # Styles spécifiques à l'accueil
│   │   ├── login.css        # Styles spécifiques au login
│   │   ├── register.css     # Styles spécifiques au register
│   │   ├── post.css         # Styles spécifiques au détail post
│   │   ├── category.css     # Styles spécifiques aux catégories
│   │   └── error.css        # Styles spécifiques aux pages erreur
│   ├── js/
│   │   └── main.js          # Animations et interactions légères
│   └── img/
│       ├── icons/           # Icônes SVG d'interface
│       │   ├── like.svg
│       │   ├── dislike.svg
│       │   ├── user.svg
│       │   └── menu.svg
│       ├── logo/            # Identité visuelle du forum
│       │   ├── logo.svg
│       │   └── favicon.ico
│       └── default/         # Images de remplacement
│           └── avatar.png

└── tests/                   # Tests unitaires Go
    ├── posts_test.go
    ├── auth_test.go
    ├── filter_test.go
    └── security_test.go
```

---

## Convention CSS

Chaque page HTML charge ses feuilles de style dans cet ordre :

```html
<!-- 1. Base obligatoire -->
<link rel="stylesheet" href="/static/css/base.css">

<!-- 2. Composants partagés utilisés sur cette page -->
<link rel="stylesheet" href="/static/css/header.css">
<link rel="stylesheet" href="/static/css/footer.css">
<link rel="stylesheet" href="/static/css/modal.css">
<link rel="stylesheet" href="/static/css/buttons.css">
<link rel="stylesheet" href="/static/css/alerts.css">

<!-- 3. CSS spécifique à la page -->
<link rel="stylesheet" href="/static/css/post.css">
```

---

## Répartition de l'équipe

| Personne | Rôle principal | Périmètre |
|----------|---------------|-----------|
| **P1** | Référent Backend + Infra | `database/`, `models/`, `handlers/`, `tests/`, Docker |
| **P2** | Référent Frontend + Intégration | `templates/`, `static/css/`, `static/js/`, `static/img/`, `handlers/auth.go` |
| **P3** | Spé Sécurité | `middleware/`, `security/`, `error.html`, `error.css`, revue globale |

---

## Packages autorisés

- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [gofrs/uuid](https://github.com/gofrs/uuid)
- Tous les packages standard Go
- JS (animations uniquement)

---
<!-- 
## Lancer le projet

### En local
```bash
go run main.go
```

### Avec Docker
```bash
docker build -t forum .
docker run -p 8080:8080 forum
``` -->