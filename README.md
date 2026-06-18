# Forum

# 📁 Forum — Arborescence du projet

Projet de forum web développé en Go avec SQLite.

---

```text
.
├── main.go                  # Point d'entrée de l'application (initialisation et routage)
├── go.mod / go.sum          # Gestion des dépendances Go
├── README.md                # Documentation du projet
│
├── database/                # Persistance des données
│   ├── database.go          # Initialisation de SQLite3 et création des tables
│   └── forum.db             # Base de données SQLite locale
│
├── security/                # Couche Sécurité & Chiffrement (DevSecOps)
│   ├── password.go          # Hachage (bcrypt) et vérification des mots de passe
│   ├── uuid.go              # Génération d'identifiants uniques sécurisés
│   ├── validate.go          # Validation des inputs (formats email, sécurité entrées)
│   └── errors.go            # Gestion centralisée des logs et erreurs de sécurité
│
├── middleware/              # Intercepteurs HTTP
│   ├── auth.go              # Vérification et restriction d'accès aux routes protégées
│   └── session.go           # Gestion du cycle de vie des sessions utilisateurs
│
├── models/                  # Structures de données (Data Transfer Objects)
│   ├── user.go / post.go    # Modèles pour les Utilisateurs et Publications
│   ├── comment.go / likes.go# Modèles pour les Commentaires et Interactions (votes)
│   ├── category.go          # Modèles pour les catégories globales
│   └── postCategory.go      # Table de liaison pour l'association Posts ↔ Catégories
│
├── handlers/                # Contrôleurs (Logique métier HTTP)
│   ├── auth.go              # Inscription, Connexion, Déconnexion
│   ├── posts.go             # Affichage du feed (Index) et gestion des posts uniques
│   ├── comments.go          # Soumission et traitement des commentaires
│   ├── likes.go             # Logique d'upvote / downvote
│   ├── filter.go            # Filtrage des posts par tags ou popularité
│   ├── handlers.go          # Utilitaires de handlers génériques
│   ├── RecentCategorie.go   # Gestion des catégories récemment utilisées
│   ├── Triepop.go           # Logique de tri par popularité (algorithme de score)
│   └── errors.go            # Routage des vues d'erreurs (404, 500)
│
├── templates/               # Vues HTML (Go Templates)
│   ├── index.html           # Page d'accueil (Fil d'actualité global)
│   ├── post.html            # Vue détaillée d'une publication spécifique
│   ├── login / register.html# Formulaires d'authentification
│   ├── category / error.html# Gestion des listes par catégorie et pages d'erreurs
│   └── poly/                # Composants réutilisables (Fragments)
│       ├── header.html      # Barre de navigation supérieure (avec statut de session)
│       ├── footer.html      # Copyright et mentions de la barre latérale
│       └── modal.html       # Fenêtres modales contextuelles
│
├── static/                  # Actifs Statiques (Servis directement par Go)
│   ├── css/                 # Feuilles de style modulaires (base, index, post, header...)
│   ├── js/                  # Logique Front-end
│   │   └── main.js          # Écouteurs globaux (clics cartes, requêtes asynchrones)
│   └── img/                 # Ressources graphiques (logos, icônes SVG d'upvote/downvote)
│
└── tests/                   # Suite de tests automatisés
    ├── auth_test.go         # Tests unitaires sur l'authentification
    ├── posts_test.go        # Fixtures et injections de jeux de données en base
    ├── filter_test.go       # Tests sur les algorithmes de filtrage
    └── security_test.go     # Tests d'intrusion / robustesse des fonctions de sécurité

```

---

### 🎨 Gestion des Styles et Performance

Chaque fichier HTML possède sa propre feuille de style dédiée (ex: `category.css`, `login.css`).

Cependant, pour éviter de surcharger les templates qui nécessitent de nombreux composants communs, nous avons **"polymérisé"** l'inclusion des fichiers transversaux. Au lieu de lister manuellement toutes les balises `<link>` (base, header, footer, modales, boutons), ces pages appellent uniquement un fichier centralisé : `style.css`. Ce dernier regroupe l'ensemble des modules partagés via des règles `@import`, simplifiant l'architecture et optimisant la mise en cache par le navigateur.

---

## Répartition de l'équipe

| Personne   | Rôle principal                  | Périmètre                                                                    |
| ---------- | ------------------------------- | ---------------------------------------------------------------------------- |
| **Amaury** | Référent Backend + Infra        | `database/`, `models/`, `handlers/`, `tests/`                                |
| **Gabor**  | Référent Frontend + Intégration | `templates/`, `static/css/`, `static/js/`, `static/img/`, `handlers/auth.go` |
| **Kilian** | Spé Sécurité                    | `middleware/`, `security/`, `error.html`, `error.css`, revue globale         |

---

## Packages autorisés

- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [gofrs/uuid](https://github.com/gofrs/uuid)
- Tous les packages standard Go
- JS (animations uniquement)

---

## 🚀 Lancer le projet

```bash
go run .
```

---

## 🧭 Accès direct aux interfaces spécifiques

Saisissez manuellement ces URL dans votre navigateur pour visualiser et tester directement les différentes vues :

```
http://localhost:8080/category
http://localhost:8080/login
http://localhost:8080/register
```
